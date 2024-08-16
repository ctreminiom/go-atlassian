package admin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ctreminiom/go-atlassian/admin/internal"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service/common"
	"io"
	"net/http"
	"net/url"
)

const defaultApiEndpoint = "https://api.atlassian.com/"

// New creates a new instance of Client.
// It takes a common.HttpClient as input and returns a pointer to Client and an error.
func New(httpClient common.HttpClient) (*Client, error) {

	// If no HTTP client is provided, use the default HTTP client.
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// Parse the default API endpoint URL.
	u, err := url.Parse(defaultApiEndpoint)
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

	return client, nil
}

// Client represents a client for interacting with the Atlassian Administration API.
type Client struct {
	// HTTP is the HTTP client used for making requests.
	HTTP common.HttpClient
	// Site is the base URL for the API.
	Site *url.URL
	// Auth is the authentication service.
	Auth common.Authentication
	// Organization is the service for organization-related operations.
	Organization *internal.OrganizationService
	// User is the service for user-related operations.
	User *internal.UserService
	// SCIM is the service for SCIM-related operations.
	SCIM *internal.SCIMService
}

// NewRequest creates a new HTTP request with the given context, method, URL string, content type, and body.
// It returns an HTTP request and an error.
func (c *Client) NewRequest(ctx context.Context, method, urlStr, type_ string, body interface{}) (*http.Request, error) {

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
	if body != nil && type_ == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	if body != nil && type_ != "" {
		req.Header.Set("Content-Type", type_)
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
			return res, model.ErrNotFound

		case http.StatusUnauthorized:
			return res, model.ErrUnauthorized

		case http.StatusInternalServerError:
			return res, model.ErrInternalError

		case http.StatusBadRequest:
			return res, model.ErrBadRequestError

		default:
			return res, model.ErrInvalidStatusCodeError
		}
	}

	if structure != nil {
		if err = json.Unmarshal(responseAsBytes, &structure); err != nil {
			return res, err
		}
	}

	return res, nil
}
