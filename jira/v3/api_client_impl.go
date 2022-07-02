package v3

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/ctreminiom/go-atlassian/jira/internal"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service/common"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

func NewV2(httpClient common.HttpClient, site string) (*ClientV2, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if !strings.HasSuffix(site, "/") {
		site += "/"
	}

	siteAsURL, err := url.Parse(site)
	if err != nil {
		return nil, err
	}

	client := &ClientV2{
		HTTP: httpClient,
		Site: siteAsURL,
	}

	applicationRoleService, err := internal.NewApplicationRoleService(client, "3")
	if err != nil {
		return nil, err
	}

	client.Authentication = internal.NewAuthenticationService(client)
	client.ApplicationRole = applicationRoleService

	return client, nil
}

type ClientV2 struct {
	HTTP            common.HttpClient
	Site            *url.URL
	Authentication  common.Authentication
	ApplicationRole jira.ApplicationRole
}

func (c *ClientV2) NewJsonRequest(ctx context.Context, method, apiEndpoint string, payload io.Reader) (*http.Request, error) {

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
	request.Header.Set("Content-Type", "application/json")

	if c.Authentication.HasBasicAuth() {
		request.SetBasicAuth(c.Authentication.GetBasicAuth())
	}

	if c.Authentication.HasUserAgent() {
		request.Header.Set("User-Agent", c.Authentication.GetUserAgent())
	}

	return request, nil
}

func (c *ClientV2) NewRequest(ctx context.Context, method, apiEndpoint string, payload io.Reader) (*http.Request, error) {

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

	if c.Authentication.HasBasicAuth() {
		request.SetBasicAuth(c.Authentication.GetBasicAuth())
	}

	if c.Authentication.HasUserAgent() {
		request.Header.Set("User-Agent", c.Authentication.GetUserAgent())
	}

	return request, nil
}

func (c *ClientV2) Call(request *http.Request, structure interface{}) (*models.ResponseScheme, error) {

	response, err := c.HTTP.Do(request)

	if err != nil {
		return nil, err
	}

	responseTransformed := &models.ResponseScheme{
		Response: response,
		Code:     response.StatusCode,
		Endpoint: response.Request.URL.String(),
		Method:   response.Request.Method,
	}

	if !(response.StatusCode >= 200 && response.StatusCode < 300) {
		return responseTransformed, models.ErrInvalidStatusCodeError
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

	_, err = responseTransformed.Bytes.Write(responseAsBytes)
	if err != nil {
		return nil, err
	}

	return responseTransformed, nil
}

func (c *ClientV2) TransformTheHTTPResponse(response *http.Response, structure interface{}) (*models.ResponseScheme, error) {

	if response == nil {
		return nil, errors.New("validation failed, please provide a http.Response pointer")
	}

	responseTransformed := &models.ResponseScheme{}
	responseTransformed.Code = response.StatusCode
	responseTransformed.Endpoint = response.Request.URL.String()
	responseTransformed.Method = response.Request.Method

	var wasSuccess = response.StatusCode >= 200 && response.StatusCode < 300
	if !wasSuccess {

		return responseTransformed, errors.New("TODO")
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

func (c *ClientV2) TransformStructToReader(structure interface{}) (io.Reader, error) {

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