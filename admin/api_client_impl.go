package admin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/ctreminiom/go-atlassian/v2/admin/internal"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/oauth2"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

const defaultAPIEndpoint = "https://api.atlassian.com/"

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

// New creates a new instance of Client.
// It takes a common.HTTPClient and optional configuration options as input and returns a pointer to Client and an error.
func New(httpClient common.HTTPClient, options ...ClientOption) (*Client, error) {

	// If no HTTP client is provided, use the default HTTP client.
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// Parse the default API endpoint URL.
	u, err := url.Parse(defaultAPIEndpoint)
	if err != nil {
		return nil, err
	}

	// Initialize the Client struct with the provided HTTP client and parsed URL.
	client := &Client{
		HTTP: httpClient,
		Site: u,
	}

	// Initialize the Authentication service.
	client.Auth = internal.NewAuthenticationService(client)

	// Initialize the SCIM service with user, group, and schema services.
	client.SCIM = &internal.SCIMService{
		User:   internal.NewSCIMUserService(client),
		Group:  internal.NewSCIMGroupService(client),
		Schema: internal.NewSCIMSchemaService(client),
	}

	// Initialize the Organization service with policy and directory services.
	client.Organization = internal.NewOrganizationService(
		client,
		internal.NewOrganizationPolicyService(client),
		internal.NewOrganizationDirectoryService(client))

	// Initialize the User service with a user token service.
	client.User = internal.NewUserService(client, internal.NewUserTokenService(client))

	// Apply client options
	for _, option := range options {
		if err := option(client); err != nil {
			return nil, err
		}
	}

	return client, nil
}

// Client represents a client for interacting with the Atlassian Administration API.
type Client struct {
	// HTTP is the HTTP client used for making requests.
	HTTP common.HTTPClient
	// Site is the base URL for the API.
	Site *url.URL
	// Auth is the authentication service.
	Auth common.Authentication
	// OAuth is the OAuth 2.0 service.
	OAuth common.OAuth2Service
	// Organization is the service for organization-related operations.
	Organization *internal.OrganizationService
	// User is the service for user-related operations.
	User *internal.UserService
	// SCIM is the service for SCIM-related operations.
	SCIM *internal.SCIMService
}

// NewRequest creates a new HTTP request with the given context, method, URL string, content type, and body.
// It returns an HTTP request and an error.
func (c *Client) NewRequest(ctx context.Context, method, urlStr, contentType string, body interface{}) (*http.Request, error) {

	// Parse the relative URL.
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// Resolve the relative URL to an absolute URL.
	u := c.Site.ResolveReference(rel)

	// Encode the body to JSON if provided.
	buf := new(bytes.Buffer)
	if body != nil {
		if err = json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	// Create a new HTTP request with the given context, method, and URL.
	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	// Set the Content-Type header if a body is provided.
	if body != nil && contentType == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	if body != nil && contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}

	// Add the Authorization header if a bearer token is available.
	if c.Auth.GetBearerToken() != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.Auth.GetBearerToken()))
	}

	// Set the User-Agent header if available.
	if c.Auth.HasUserAgent() {
		req.Header.Set("User-Agent", c.Auth.GetUserAgent())
	}

	return req, nil
}

// Call sends an HTTP request and processes the response.
// It takes an *http.Request and a structure to unmarshal the response into.
// It returns a pointer to model.ResponseScheme and an error.
func (c *Client) Call(request *http.Request, structure interface{}) (*model.ResponseScheme, error) {

	// Perform the HTTP request.
	response, err := c.HTTP.Do(request)
	if err != nil {
		return nil, err
	}

	// Process the HTTP response.
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
