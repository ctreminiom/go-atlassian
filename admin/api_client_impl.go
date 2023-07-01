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

func New(httpClient common.HttpClient) (*Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	u, err := url.Parse(defaultApiEndpoint)
	if err != nil {
		return nil, err
	}

	client := &Client{
		HTTP: httpClient,
		Site: u,
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

type Client struct {
	HTTP         common.HttpClient
	Site         *url.URL
	Auth         common.Authentication
	Organization *internal.OrganizationService
	User         *internal.UserService
	SCIM         *internal.SCIMService
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr, type_ string, body interface{}) (*http.Request, error) {

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

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")

	if body != nil && type_ == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	if body != nil && type_ != "" {
		req.Header.Set("Content-Type", type_)
	}

	if c.Auth.GetBearerToken() != "" {
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
