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
	"strings"
)

type Client struct {
	HTTP *http.Client
	Site *url.URL

	Auth         *AuthenticationService
	Organization *OrganizationService
	User         *UserService

	SCIM *SCIMService
}

const ApiEndpoint = "https://api.atlassian.com/"

//New
func New(httpClient *http.Client) (client *Client, err error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	siteAsURL, err := url.Parse(ApiEndpoint)
	if err != nil {
		return
	}

	client = &Client{}
	client.HTTP = httpClient
	client.Site = siteAsURL

	client.Auth = &AuthenticationService{client: client}
	client.Organization = &OrganizationService{
		client: client,
		Policy: &OrganizationPolicyService{
			client:   client,
			Resource: &OrganizationPolicyResourceService{client: client},
		},
	}

	client.User = &UserService{
		client: client,
		Token:  &UserTokenService{client: client},
	}

	client.SCIM = &SCIMService{
		client:   client,
		User:     &SCIMUserService{client: client},
		Group:    &SCIMGroupService{client: client},
		Scheme:   &SCIMSchemeService{client: client},
		Resource: &SCIMResourceService{client: client},
	}

	return
}

func (c *Client) newRequest(ctx context.Context, method, urlAsString string, payload interface{}) (request *http.Request, err error) {

	if ctx == nil {
		return nil, errors.New("the context param is nil, please provide a valid one")
	}

	relativePath, err := url.Parse(urlAsString)
	if err != nil {
		return
	}

	relativePath.Path = strings.TrimLeft(relativePath.Path, "/")

	endpointPath := c.Site.ResolveReference(relativePath)
	var payloadBuffer io.ReadWriter
	if payload != nil {
		payloadBuffer = new(bytes.Buffer)
		if err = json.NewEncoder(payloadBuffer).Encode(payload); err != nil {
			return
		}
	}

	request, err = http.NewRequestWithContext(ctx, method, endpointPath.String(), payloadBuffer)
	if err != nil {
		return
	}

	if c.Auth.beaverToken != "" {
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.Auth.beaverToken))
	}

	if c.Auth.agent != "" {
		request.Header.Set("User-Agent", c.Auth.agent)
	}

	return
}

func (c *Client) Do(request *http.Request) (response *Response, err error) {

	httpResponse, err := c.HTTP.Do(request)
	if err != nil {
		return
	}

	response, err = checkResponse(httpResponse, request.URL.String())
	if err != nil {
		return
	}

	response, err = newResponse(httpResponse, request.URL.String())
	if err != nil {
		return
	}

	return
}

type Response struct {
	StatusCode  int
	Endpoint    string
	Headers     map[string][]string
	BodyAsBytes []byte
	Method      string
}

func newResponse(http *http.Response, endpoint string) (response *Response, err error) {

	var statusCode = http.StatusCode

	var httpResponseAsBytes []byte
	if http.ContentLength != 0 {
		httpResponseAsBytes, err = ioutil.ReadAll(http.Body)
		if err != nil {
			return
		}
	}

	newResponse := Response{
		StatusCode:  statusCode,
		Headers:     http.Header,
		BodyAsBytes: httpResponseAsBytes,
		Endpoint:    endpoint,
		Method:      http.Request.Method,
	}

	return &newResponse, nil
}

func checkResponse(http *http.Response, endpoint string) (response *Response, err error) {

	var statusCode = http.StatusCode
	if 200 <= statusCode && statusCode <= 299 {
		return
	}

	var httpResponseAsBytes []byte
	if http.ContentLength != 0 {
		httpResponseAsBytes, err = ioutil.ReadAll(http.Body)
		if err != nil {
			return
		}
	}

	newErrorResponse := Response{
		StatusCode:  statusCode,
		Headers:     http.Header,
		BodyAsBytes: httpResponseAsBytes,
		Endpoint:    endpoint,
		Method:      http.Request.Method,
	}

	return &newErrorResponse, fmt.Errorf("request failed. Please analyze the request body for more details. Status Code: %d", statusCode)
}
