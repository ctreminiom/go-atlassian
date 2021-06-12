package jira

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ctreminiom/go-atlassian/jira/agile"
	"github.com/ctreminiom/go-atlassian/jira/sm"
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
	Permission *PermissionService
	Project    *ProjectService
	Screen     *ScreenService
	Server     *ServerService
	Task       *TaskService
	User       *UserService
	MySelf     *MySelfService

	//Service Management Module
	ServiceManagement *sm.Client

	//Cloud Agile Module
	Agile *agile.Client
}

const (
	DateFormatJira = "2006-01-02T15:04:05.999-0700"
)

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

	//Service Management module integration
	serviceManagementClient, _ := sm.New(httpClient, site)

	client.ServiceManagement = serviceManagementClient

	// Agile Module integration
	agileClient, _ := agile.New(httpClient, site)
	client.Agile = agileClient

	client.Role = &ApplicationRoleService{client: client}
	client.Audit = &AuditService{client: client}
	client.Auth = &AuthenticationService{client: client}
	client.Dashboard = &DashboardService{client: client}

	// Admin Endpoints
	client.Server = &ServerService{client: client}
	client.Task = &TaskService{client: client}

	client.Screen = &ScreenService{
		client: client,
		Tab: &ScreenTabService{
			client: client,
			Field:  &ScreenTabFieldService{client: client},
		},
		Scheme: &ScreenSchemeService{client: client},
	}

	client.Filter = &FilterService{
		client: client,
		Share:  &FilterShareService{client: client},
	}

	client.Group = &GroupService{client: client}

	client.Issue = &IssueService{
		client:     client,
		Attachment: &AttachmentService{client: client},
		Comment: &CommentService{
			client: client,
		},
		Field: &FieldService{
			client:        client,
			Configuration: &FieldConfigurationService{client: client},

			Context: &FieldContextService{
				client: client,
				Option: &FieldOptionContextService{client: client},
			},
		},
		Priority:   &PriorityService{client: client},
		Resolution: &ResolutionService{client: client},

		Search: &IssueSearchService{client: client},

		Type: &IssueTypeService{
			client:       client,
			Scheme:       &IssueTypeSchemeService{client: client},
			ScreenScheme: &IssueTypeScreenSchemeService{client: client},
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
		Scheme: &PermissionSchemeService{
			client: client,
			Grant:  &PermissionGrantSchemeService{client: client},
		},
	}

	client.Project = &ProjectService{
		client: client,

		Category:   &ProjectCategoryService{client: client},
		Component:  &ProjectComponentService{client: client},
		Valid:      &ProjectValidationService{client: client},
		Permission: &ProjectPermissionSchemeService{client: client},

		Role: &ProjectRoleService{
			client: client,
			Actor:  &ProjectRoleActorService{client: client},
		},

		Type:    &ProjectTypeService{client: client},
		Version: &ProjectVersionService{client: client},
	}

	client.User = &UserService{
		client: client,
		Search: &UserSearchService{client: client},
	}

	client.MySelf = &MySelfService{client: client}

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
		_ = json.NewEncoder(payloadBuffer).Encode(payload)
	}

	request, err = http.NewRequestWithContext(ctx, method, endpointPath.String(), payloadBuffer)
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
		httpResponseAsBytes, _ = ioutil.ReadAll(http.Body)
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
		httpResponseAsBytes, _ = ioutil.ReadAll(http.Body)
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
