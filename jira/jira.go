package jira

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

	Role       *ApplicationRoleService
	Audit      *AuditService
	Auth       *AuthenticationService
	Dashboard  *DashboardService
	Filter     *FilterService
	Group      *GroupService
	Issue      *IssueService
	Search     *SearchService
	Permission *PermissionService
}

//New
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

	client.Role = &ApplicationRoleService{client: client}
	client.Audit = &AuditService{client: client}
	client.Auth = &AuthenticationService{client: client}
	client.Dashboard = &DashboardService{client: client}

	client.Filter = &FilterService{
		client: client,
		Share:  &FilterShareService{client: client},
	}

	client.Group = &GroupService{client: client}
	client.Search = &SearchService{client: client}

	client.Issue = &IssueService{
		client:     client,
		Attachment: &AttachmentService{client: client},
		Comment: &CommentService{
			client:     client,
			Properties: &CommentPropertiesService{client: client},
		},
		Field: &FieldService{
			client:        client,
			Configuration: &FieldConfigurationService{client: client},

			Context: &FieldContextService{
				client: client,
				Option: &FieldOptionContextService{client: client},
			},

			Option: &FieldOptionService{client: client},
		},
		Priority:   &PriorityService{client: client},
		Property:   &IssuePropertyService{client: client},
		Resolution: &ResolutionService{client: client},
		Security: &IssueSecurityService{
			client: client,
			Scheme: &IssueSecuritySchemeService{client: client},
		},
		Type: &IssueTypeService{
			client: client,
			Scheme: &IssueTypeSchemeService{client: client},
		},

		Link: &IssueLinkService{
			client: client,
			Type:   &IssueLinkTypeService{client: client},
		},
		Votes:    &VoteService{client: client},
		Watchers: &WatcherService{client: client},
		Label:    &LabelService{client: client},
	}

	client.Permission = &PermissionService{
		client: client,
		Scheme: &PermissionSchemeService{client: client},
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

	request.SetBasicAuth(c.Auth.mail, c.Auth.token)

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
