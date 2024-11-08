package assets

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/ctreminiom/go-atlassian/v2/assets/internal"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

const DefaultAssetsSite = "https://api.atlassian.com/"

// New creates a new instance of Client.
// It takes a common.HTTPClient and a site URL as inputs and returns a pointer to Client and an error.
func New(httpClient common.HTTPClient, site string) (*Client, error) {

	// If no HTTP client is provided, use the default HTTP client.
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	// If no site URL is provided, use the default assets site URL.
	if site == "" {
		site = DefaultAssetsSite
	}

	// Parse the site URL.
	u, err := url.Parse(site)
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

	// Initialize the Assets services.
	client.AQL = internal.NewAQLService(client)
	client.Icon = internal.NewIconService(client)
	client.Object = internal.NewObjectService(client)
	client.ObjectSchema = internal.NewObjectSchemaService(client)
	client.ObjectType = internal.NewObjectTypeService(client)
	client.ObjectTypeAttribute = internal.NewObjectTypeAttributeService(client)

	return client, nil
}

// Client represents a client for interacting with the Atlassian Assets API.
type Client struct {
	// HTTP is the HTTP client used for making requests.
	HTTP common.HTTPClient
	// Site is the base URL for the API.
	Site *url.URL
	// Auth is the authentication service.
	Auth common.Authentication
	// AQL is the service for AQL-related operations.
	AQL *internal.AQLService
	// Icon is the service for icon-related operations.
	Icon *internal.IconService
	// Object is the service for object-related operations.
	Object *internal.ObjectService
	// ObjectSchema is the service for object schema-related operations.
	ObjectSchema *internal.ObjectSchemaService
	// ObjectType is the service for object type-related operations.
	ObjectType *internal.ObjectTypeService
	// ObjectTypeAttribute is the service for object type attribute-related operations.
	ObjectTypeAttribute *internal.ObjectTypeAttributeService
}

// NewRequest creates a new HTTP request with the given context, method, URL string, content type, and body.
// It returns an HTTP request and an error.
func (c *Client) NewRequest(ctx context.Context, method, urlStr, contentType string, body interface{}) (*http.Request, error) {

	// Parse the relative URL.
	ctx, span := tracer().Start(ctx, "(*Client).NewRequest")
	defer span.End()

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

	// Add the Basic Authentication header if available.
	if c.Auth.HasBasicAuth() {
		req.SetBasicAuth(c.Auth.GetBasicAuth())
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
			return res, model.ErrInternal

		case http.StatusBadRequest:
			return res, model.ErrBadRequest

		default:
			return res, model.ErrInvalidStatusCode
		}
	}

	if structure != nil {
		if err = json.Unmarshal(responseAsBytes, &structure); err != nil {
			return res, err
		}
	}

	return res, nil
}
