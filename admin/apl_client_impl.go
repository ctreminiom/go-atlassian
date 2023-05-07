package admin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ctreminiom/go-atlassian/admin/internal"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service/common"
	"io"
	"net/http"
	"net/url"
	"reflect"
)

const ApiEndpoint = "https://api.atlassian.com/"

type Client struct {
	HTTP         common.HttpClient
	Site         *url.URL
	Auth         common.Authentication
	Organization *internal.OrganizationService
	User         *internal.UserService
	SCIM         *internal.SCIMService
}

func New(httpClient common.HttpClient) (*Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	siteAsURL, err := url.Parse(ApiEndpoint)
	if err != nil {
		return nil, err
	}

	client := &Client{
		HTTP: httpClient,
		Site: siteAsURL,
	}

	client.Auth = internal.NewAuthenticationService(client)

	client.SCIM = &internal.SCIMService{
		User:   internal.NewSCIMUserService(client),
		Group:  internal.NewSCIMGroupService(client),
		Schema: internal.NewSCIMSchemaService(client),
	}

	client.Organization = internal.NewOrganizationService(
		client,
		internal.NewOrganizationPolicyService(client),
		internal.NewOrganizationDirectoryService(client))

	client.User = internal.NewUserService(client, internal.NewUserTokenService(client))

	return client, nil
}

func (c *Client) NewFormRequest(ctx context.Context, method, apiEndpoint, contentType string, payload io.Reader) (*http.Request, error) {
	return nil, nil
}

func (c *Client) NewRequest(ctx context.Context, method, apiEndpoint string, payload io.Reader) (*http.Request, error) {

	relativePath, err := url.Parse(apiEndpoint)
	if err != nil {
		return nil, err
	}

	var endpoint = c.Site.ResolveReference(relativePath).String()

	request, err := http.NewRequestWithContext(ctx, method, endpoint, payload)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json")

	if payload != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	if c.Auth.GetBearerToken() != "" {
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.Auth.GetBearerToken()))
	}

	if c.Auth.HasUserAgent() {
		request.Header.Set("User-Agent", c.Auth.GetUserAgent())
	}

	return request, nil
}

func (c *Client) Call(request *http.Request, structure interface{}) (*models.ResponseScheme, error) {

	response, err := c.HTTP.Do(request)
	if err != nil {
		return nil, err
	}

	return c.TransformTheHTTPResponse(response, structure)
}

func (c *Client) TransformTheHTTPResponse(response *http.Response, structure interface{}) (*models.ResponseScheme, error) {

	responseTransformed := &models.ResponseScheme{
		Response: response,
		Code:     response.StatusCode,
		Endpoint: response.Request.URL.String(),
		Method:   response.Request.Method,
	}

	responseAsBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return responseTransformed, err
	}

	responseTransformed.Bytes.Write(responseAsBytes)

	var wasSuccess = response.StatusCode >= 200 && response.StatusCode < 300
	if !wasSuccess {
		return responseTransformed, models.ErrInvalidStatusCodeError
	}

	if structure != nil {
		if err = json.Unmarshal(responseAsBytes, &structure); err != nil {
			return responseTransformed, err
		}
	}

	return responseTransformed, nil
}

func (c *Client) TransformStructToReader(structure interface{}) (io.Reader, error) {

	if structure == nil {
		return nil, models.ErrNilPayloadError
	}

	if reflect.ValueOf(structure).Type().Kind() == reflect.Struct {
		return nil, models.ErrNonPayloadPointerError
	}

	structureAsBodyBytes, err := json.Marshal(structure)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(structureAsBodyBytes), nil
}
