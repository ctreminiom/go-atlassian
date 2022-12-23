package admin

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

type Client struct {
	HTTP         *http.Client
	Site         *url.URL
	Auth         *AuthenticationService
	Organization *OrganizationService
	User         *UserService
	SCIM         *SCIMService
}

const ApiEndpoint = "https://api.atlassian.com/"

func New(httpClient *http.Client) (client *Client, err error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	siteAsURL, _ := url.Parse(ApiEndpoint)

	client = &Client{}
	client.HTTP = httpClient
	client.Site = siteAsURL

	client.Auth = &AuthenticationService{client: client}
	client.Organization = &OrganizationService{
		client: client,
		Policy: &OrganizationPolicyService{
			client: client,
		},
	}

	client.User = &UserService{
		client: client,
		Token:  &UserTokenService{client: client},
	}

	client.SCIM = &SCIMService{
		client: client,
		User:   &SCIMUserService{client: client},
		Group:  &SCIMGroupService{client: client},
		Scheme: &SCIMSchemeService{client: client},
	}

	return
}

func (c *Client) newRequest(ctx context.Context, method, apiEndpoint string, payload io.Reader) (request *http.Request, err error) {

	relativePath, err := url.Parse(apiEndpoint)
	if err != nil {
		return nil, fmt.Errorf(urlParsedError, err.Error())
	}

	var endpoint = c.Site.ResolveReference(relativePath).String()

	request, err = http.NewRequestWithContext(ctx, method, endpoint, payload)
	if err != nil {
		return nil, fmt.Errorf(requestCreationError, err.Error())
	}

	if c.Auth.beaverToken != "" {
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.Auth.beaverToken))
	}

	if c.Auth.agent != "" {
		request.Header.Set("User-Agent", c.Auth.agent)
	}

	return
}

func (c *Client) call(request *http.Request, structure interface{}) (result *ResponseScheme, err error) {
	response, err := c.HTTP.Do(request)
	if err != nil {
		return nil, err
	}
	return transformTheHTTPResponse(response, structure)
}

func transformStructToReader(structure interface{}) (reader io.Reader, err error) {

	if structure == nil || reflect.ValueOf(structure).IsNil() {
		return nil, structureNotParsedError
	}

	structureAsBodyBytes, err := json.Marshal(structure)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(structureAsBodyBytes), nil
}

func transformTheHTTPResponse(response *http.Response, structure interface{}) (result *ResponseScheme, err error) {

	if response == nil {
		return nil, errors.New("validation failed, please provide a http.Response pointer")
	}

	responseTransformed := &ResponseScheme{}
	responseTransformed.Code = response.StatusCode
	responseTransformed.Endpoint = response.Request.URL.String()
	responseTransformed.Method = response.Request.Method

	var wasSuccess = response.StatusCode >= 200 && response.StatusCode < 300
	if !wasSuccess {

		return responseTransformed, fmt.Errorf(requestFailedError, response.StatusCode)
	}

	responseAsBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return responseTransformed, err
	}

	if structure != nil {
		if err = json.Unmarshal(responseAsBytes, &structure); err != nil {
			return responseTransformed, err
		}
	}

	responseTransformed.Bytes.Write(responseAsBytes)

	return responseTransformed, nil
}

type ResponseScheme struct {
	Code     int
	Endpoint string
	Method   string
	Bytes    bytes.Buffer
	Headers  map[string][]string
}

var (
	requestCreationError    = "request creation failed: %v"
	urlParsedError          = "URL parsing failed: %v"
	requestFailedError      = "request failed. Please analyze the request body for more details. Status Code: %d"
	structureNotParsedError = errors.New("failed to parse the interface pointer, please provide a valid one")
)
