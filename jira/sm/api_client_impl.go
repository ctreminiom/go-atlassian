package sm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/ctreminiom/go-atlassian/v2/jira/sm/internal"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/oauth2"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

const defaultServiceManagementVersion = "latest"

// ClientOption is a function that configures a Client
type ClientOption func(*Client) error

// WithOAuth configures the client with OAuth 2.0 support
func WithOAuth(config *common.OAuth2Config) ClientOption {
	return func(c *Client) error {
		if config == nil {
			return fmt.Errorf("oauth config cannot be nil")
		}
		
		oauthService, err := oauth2.NewOAuth2Service(c.HTTP, config)
		if err != nil {
			return fmt.Errorf("failed to create OAuth service: %w", err)
		}
		
		c.OAuth = oauthService
		return nil
	}
}

// WithAutoRenewalToken enables automatic OAuth token renewal with the provided token.
// This option requires WithOAuth to be configured first or used together.
func WithAutoRenewalToken(token *common.OAuth2Token) ClientOption {
	return func(c *Client) error {
		if token == nil {
			return fmt.Errorf("token cannot be nil for auto-renewal")
		}
		
		if c.OAuth == nil {
			return fmt.Errorf("OAuth must be configured before enabling auto-renewal (use WithOAuth first)")
		}
		
		// Create token sources with storage support if configured
		_, reuseSource, err := oauth2.SetupTokenSourcesWithStorage(
			context.Background(),
			token,
			c.OAuth,
			c.HTTP,
		)
		if err != nil {
			return fmt.Errorf("failed to setup token sources: %w", err)
		}

		// Extract base transport and restore original HTTP client if wrapped
		base := oauth2.ExtractBaseTransport(c.HTTP)
		if wrapper, ok := oauth2.ExtractWrapper(c.HTTP); ok {
			c.HTTP = wrapper.OriginalClient
		}

		// Create OAuth transport
		c.HTTP = oauth2.CreateOAuthTransport(reuseSource, base, c.Auth)
		
		// Set initial token
		c.Auth.SetBearerToken(token.AccessToken)
		
		return nil
	}
}

// WithOAuthWithAutoRenewal is a convenience option that combines WithOAuth and WithAutoRenewalToken.
// It configures OAuth support and enables automatic token renewal in one step.
func WithOAuthWithAutoRenewal(config *common.OAuth2Config, token *common.OAuth2Token) ClientOption {
	return func(c *Client) error {
		// First configure OAuth
		if err := WithOAuth(config)(c); err != nil {
			return err
		}
		
		// Then enable auto-renewal
		return WithAutoRenewalToken(token)(c)
	}
}

// WithTokenStore configures the client to use external token storage
func WithTokenStore(store oauth2.TokenStore) ClientOption {
	return func(c *Client) error {
		if store == nil {
			return fmt.Errorf("token store cannot be nil")
		}
		
		c.HTTP = oauth2.WrapHTTPClient(c.HTTP).WithStore(store)
		return nil
	}
}

// WithTokenCallback configures the client to use a callback for token refresh events
func WithTokenCallback(callback oauth2.TokenCallback) ClientOption {
	return func(c *Client) error {
		if callback == nil {
			return fmt.Errorf("token callback cannot be nil")
		}
		
		c.HTTP = oauth2.WrapHTTPClient(c.HTTP).WithCallback(callback)
		return nil
	}
}

func New(httpClient common.HTTPClient, site string, options ...ClientOption) (*Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if site == "" {
		return nil, fmt.Errorf("client: %w", model.ErrNoSite)
	}

	if !strings.HasSuffix(site, "/") {
		site += "/"
	}

	u, err := url.Parse(site)
	if err != nil {
		return nil, err
	}

	client := &Client{
		HTTP: httpClient,
		Site: u,
	}

	client.Auth = internal.NewAuthenticationService(client)
	client.Customer = internal.NewCustomerService(client, defaultServiceManagementVersion)
	client.Info = internal.NewInfoService(client, defaultServiceManagementVersion)
	client.Knowledgebase = internal.NewKnowledgebaseService(client, defaultServiceManagementVersion)
	client.Organization = internal.NewOrganizationService(client, defaultServiceManagementVersion)
	client.WorkSpace = internal.NewWorkSpaceService(client, defaultServiceManagementVersion)

	requestSubServices := &internal.ServiceRequestSubServices{
		Approval:    internal.NewApprovalService(client, defaultServiceManagementVersion),
		Attachment:  internal.NewAttachmentService(client, defaultServiceManagementVersion),
		Comment:     internal.NewCommentService(client, defaultServiceManagementVersion),
		Participant: internal.NewParticipantService(client, defaultServiceManagementVersion),
		SLA:         internal.NewServiceLevelAgreementService(client, defaultServiceManagementVersion),
		Feedback:    internal.NewFeedbackService(client, defaultServiceManagementVersion),
		Type:        internal.NewTypeService(client, defaultServiceManagementVersion),
	}

	requestService, err := internal.NewRequestService(client, defaultServiceManagementVersion, requestSubServices)
	if err != nil {
		return nil, err
	}
	client.Request = requestService

	serviceDeskService, err := internal.NewServiceDeskService(
		client,
		defaultServiceManagementVersion,
		internal.NewQueueService(client, defaultServiceManagementVersion))

	if err != nil {
		return nil, err
	}
	client.ServiceDesk = serviceDeskService

	// Apply client options
	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

type Client struct {
	HTTP          common.HTTPClient
	Site          *url.URL
	Auth          common.Authentication
	OAuth         common.OAuth2Service
	Customer      *internal.CustomerService
	Info          *internal.InfoService
	Knowledgebase *internal.KnowledgebaseService
	Organization  *internal.OrganizationService
	Request       *internal.RequestService
	ServiceDesk   *internal.ServiceDeskService
	WorkSpace     *internal.WorkSpaceService
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr, contentType string, body interface{}) (*http.Request, error) {

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.Site.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		if err = json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	// If the body interface is a *bytes.Buffer type
	// it means the NewRequest() requires to handle the RFC 1867 ISO
	if attachBuffer, ok := body.(*bytes.Buffer); ok {
		buf = attachBuffer
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if contentType != "" {
		// When the contentType is provided, it means the request needs to be created to handle files
		req.Header.Set("Content-Type", contentType)
		req.Header.Set("X-Atlassian-Token", "no-check")
	}

	if c.Auth.HasBasicAuth() {
		req.SetBasicAuth(c.Auth.GetBasicAuth())
	}

	if c.Auth.HasSetExperimentalFlag() {
		req.Header.Set("X-ExperimentalApi", "opt-in")
	}

	if c.Auth.GetBearerToken() != "" && !c.Auth.HasBasicAuth() {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.Auth.GetBearerToken()))
	}

	if c.Auth.HasUserAgent() {
		req.Header.Set("User-Agent", c.Auth.GetUserAgent())
	}

	return req, nil
}

func (c *Client) Call(request *http.Request, structure interface{}) (*model.ResponseScheme, error) {

	response, err := c.HTTP.Do(request)
	if err != nil {
		return nil, err
	}

	return c.processResponse(response, structure)
}

func (c *Client) Do(request *http.Request) (*http.Response, error) {
	return c.HTTP.Do(request)
}

func (c *Client) processResponse(response *http.Response, structure interface{}) (*model.ResponseScheme, error) {

	defer response.Body.Close()

	res := &model.ResponseScheme{
		Response: response,
		Code:     response.StatusCode,
		Endpoint: response.Request.URL.String(),
		Method:   response.Request.Method,
	}

	responseAsBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return res, err
	}

	res.Bytes.Write(responseAsBytes)

	wasSuccess := response.StatusCode >= 200 && response.StatusCode < 300

	if !wasSuccess {

		switch response.StatusCode {

		case http.StatusNotFound:
			return res, fmt.Errorf("client: %w", model.ErrNotFound)

		case http.StatusUnauthorized:
			return res, fmt.Errorf("client: %w", model.ErrUnauthorized)

		case http.StatusInternalServerError:
			return res, fmt.Errorf("client: %w", model.ErrInternal)

		case http.StatusBadRequest:
			return res, fmt.Errorf("client: %w", model.ErrBadRequest)

		default:
			return res, fmt.Errorf("client: %w", model.ErrInvalidStatusCode)
		}
	}

	if structure != nil {
		if err = json.Unmarshal(responseAsBytes, &structure); err != nil {
			return res, err
		}
	}

	return res, nil
}
