package confluence

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

	Auth *AuthenticationService
	Content *ContentService
}

func New(httpClient *http.Client, site string) (client *Client, err error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if !strings.HasSuffix(site, "/") {
		site += "/"
	}

	siteAsURL, err := url.Parse(site)
	if err != nil {
		return
	}

	client = &Client{}
	client.HTTP = httpClient
	client.Site = siteAsURL
	client.Auth = &AuthenticationService{client: client}
	client.Content = &ContentService{client: client}

	return
}

func (c *Client) newRequest(ctx context.Context, method, apiEndpoint string, payload io.Reader) (request *http.Request, err error) {

	if ctx == nil {
		return nil, errors.New("the context param is nil, please provide a valid one")
	}

	relativePath, err := url.Parse(apiEndpoint)
	if err != nil {
		return
	}

	endpoint := c.Site.ResolveReference(relativePath).String()

	request, err = http.NewRequestWithContext(ctx, method, endpoint, payload)
	if err != nil {
		return
	}

	if c.Auth.basicAuthProvided {
		request.SetBasicAuth(c.Auth.mail, c.Auth.token)
	}

	if c.Auth.userAgentProvided {
		request.Header.Set("User-Agent", c.Auth.agent)
	}

	return
}

func (c *Client) Call(request *http.Request, structure interface{}) (result *ResponseScheme, err error) {

	response, err := c.HTTP.Do(request)
	if err != nil {
		return
	}

	return transformTheHTTPResponse(response, structure)
}

func transformTheHTTPResponse(response *http.Response, structure interface{}) (result *ResponseScheme, err error) {

	if response == nil {
		return nil, errors.New("validation failed, please provide a http.Response pointer")
	}

	responseTransformed := &ResponseScheme{}
	responseTransformed.Code = response.StatusCode
	responseTransformed.Endpoint = response.Request.URL.String()
	responseTransformed.Method = response.Request.Method

	responseAsBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return responseTransformed, err
	}

	if err = json.Unmarshal(responseAsBytes, &structure); err != nil {
		return nil, err
	}

	_, err = responseTransformed.Bytes.Write(responseAsBytes)
	if err != nil {
		return responseTransformed, err
	}

	var wasSuccess = response.StatusCode >= 200 && response.StatusCode < 300
	if !wasSuccess {

		if response.StatusCode == http.StatusBadRequest {
			var apiError ApiErrorResponseScheme
			if err = json.Unmarshal(responseAsBytes, &apiError); err != nil {
				return responseTransformed, err
			}
			return responseTransformed, fmt.Errorf(requestFailedError, response.StatusCode)
		}

		return responseTransformed, fmt.Errorf(requestFailedError, response.StatusCode)
	}

	return responseTransformed, nil
}

type ResponseScheme struct {
	Code     int
	Endpoint string
	Method   string
	API      *ApiErrorResponseScheme
	Bytes    bytes.Buffer
	Headers  map[string][]string
}

type ApiErrorResponseScheme struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

var (
	requestFailedError = "request failed. Please analyze the request body for more details. Status Code: %d"
)