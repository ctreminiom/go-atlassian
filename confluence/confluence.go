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
	"reflect"
	"strings"
)

type Client struct {
	HTTP *http.Client
	Site *url.URL

	Auth    *AuthenticationService
	Content *ContentService
	Space   *SpaceService
	Label   *LabelService
	Search  *SearchService
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

	client.Content = &ContentService{
		client:             client,
		Attachment:         &ContentAttachmentService{client: client},
		ChildrenDescendant: &ContentChildrenDescendantService{client: client},
		Comment:            &ContentCommentService{client: client},
		Permission:         &ContentPermissionService{client: client},
		Label:              &ContentLabelService{client: client},
		Property:           &ContentPropertyService{client: client},
		Restriction: &ContentRestrictionService{
			client: client,
			Operation: &ContentRestrictionOperationService{
				client: client,
				Group:  &ContentRestrictionOperationGroupService{client: client},
				User:   &ContentRestrictionOperationUserService{client: client},
			}},
		Version: &ContentVersionService{client: client},
	}

	client.Space = &SpaceService{client: client}
	client.Label = &LabelService{client: client}
	client.Search = &SearchService{client: client}
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

	if c.Auth.basicAuthProvided {
		request.SetBasicAuth(c.Auth.mail, c.Auth.token)
	}

	if c.Auth.userAgentProvided {
		request.Header.Set("User-Agent", c.Auth.agent)
	}

	return
}

func (c *Client) Call(request *http.Request, structure interface{}) (result *ResponseScheme, err error) {
	response, _ := c.HTTP.Do(request)
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

		if response.StatusCode == http.StatusBadRequest {

			responseAsBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return responseTransformed, err
			}

			var apiError ApiErrorResponseScheme
			if err = json.Unmarshal(responseAsBytes, &apiError); err != nil {
				return responseTransformed, err
			}

			responseTransformed.API = &apiError
			responseTransformed.Bytes.Write(responseAsBytes)

			return responseTransformed, fmt.Errorf(requestFailedError, response.StatusCode)
		}
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
	API      *ApiErrorResponseScheme
	Bytes    bytes.Buffer
	Headers  map[string][]string
}

type ApiErrorResponseScheme struct {
	StatusCode int                         `json:"statusCode"`
	Message    string                      `json:"message"`
	Data       *ApiErrorResponseDataScheme `json:"data"`
}

type ApiErrorResponseDataScheme struct {
	Authorized bool `json:"authorized"`
	Valid      bool `json:"valid"`
	Errors     []struct {
		Message struct {
			Key  string        `json:"key"`
			Args []interface{} `json:"args"`
		} `json:"message"`
	} `json:"errors"`
	Successful bool `json:"successful"`
}

var (
	requestFailedError      = "request failed. Please analyze the request body for more details. Status Code: %d"
	requestCreationError    = "request creation failed: %v"
	urlParsedError          = "URL parsing failed: %v"
	structureNotParsedError = errors.New("failed to parse the interface pointer, please provide a valid one")
)
