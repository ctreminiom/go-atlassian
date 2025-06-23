package confluence

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/ctreminiom/go-atlassian/v2/confluence/internal"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/oauth2"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

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
		return nil, models.ErrNoSite
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

	contentSubServices := &internal.ContentSubServices{
		Attachment:         internal.NewContentAttachmentService(client),
		ChildrenDescendant: internal.NewChildrenDescandantsService(client),
		Comment:            internal.NewCommentService(client),
		Permission:         internal.NewPermissionService(client),
		Label:              internal.NewContentLabelService(client),
		Property:           internal.NewPropertyService(client),
		Version:            internal.NewVersionService(client),
		Restriction: internal.NewRestrictionService(client,
			internal.NewRestrictionOperationService(client,
				internal.NewRestrictionOperationGroupService(client),
				internal.NewRestrictionOperationUserService(client))),
	}

	client.Auth = internal.NewAuthenticationService(client)
	client.Content = internal.NewContentService(client, contentSubServices)
	client.Space = internal.NewSpaceService(client, internal.NewSpacePermissionService(client))
	client.Label = internal.NewLabelService(client)
	client.Search = internal.NewSearchService(client)
	client.LongTask = internal.NewTaskService(client)
	client.Analytics = internal.NewAnalyticsService(client)
	client.Template = internal.NewTemplateService(client)

	// Apply client options
	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

type Client struct {
	HTTP      common.HTTPClient
	Site      *url.URL
	Auth      common.Authentication
	OAuth     common.OAuth2Service
	Content   *internal.ContentService
	Space     *internal.SpaceService
	Label     *internal.LabelService
	Search    *internal.SearchService
	LongTask  *internal.TaskService
	Analytics *internal.AnalyticsService
	Template  *internal.TemplateService
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

	if c.Auth.HasUserAgent() {
		req.Header.Set("User-Agent", c.Auth.GetUserAgent())
	}

	if c.Auth.GetBearerToken() != "" && !c.Auth.HasBasicAuth() {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.Auth.GetBearerToken()))
	}

	return req, nil
}

func (c *Client) Call(request *http.Request, structure interface{}) (*models.ResponseScheme, error) {

	response, err := c.HTTP.Do(request)
	if err != nil {
		return nil, err
	}

	return c.processResponse(response, structure)
}

func (c *Client) processResponse(response *http.Response, structure interface{}) (*models.ResponseScheme, error) {

	defer response.Body.Close()

	res := &models.ResponseScheme{
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
			return res, models.ErrNotFound

		case http.StatusUnauthorized:
			return res, models.ErrUnauthorized

		case http.StatusInternalServerError:
			return res, models.ErrInternal

		case http.StatusBadRequest:
			return res, models.ErrBadRequest

		default:
			return res, models.ErrInvalidStatusCode
		}
	}

	if structure != nil {
		if err = json.Unmarshal(responseAsBytes, &structure); err != nil {
			return res, err
		}
	}

	return res, nil
}
