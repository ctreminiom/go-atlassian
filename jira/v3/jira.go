package v3

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
	Workflow   *WorkflowService
}

const (
	DateFormatJira = "2006-01-02T15:04:05.999-0700"
)

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
			client: client,
			Configuration: &FieldConfigurationService{
				client: client,
				Item:   &FieldConfigurationItemService{client: client},
				Scheme: &FieldConfigurationSchemeService{client: client},
			},

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
		Worklog:  &IssueWorklogService{client: client},
		Metadata: &IssueMetadataService{client: client},
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

		Type:     &ProjectTypeService{client: client},
		Version:  &ProjectVersionService{client: client},
		Feature:  &ProjectFeatureService{client: client},
		Property: &ProjectPropertyService{client: client},
	}

	client.User = &UserService{
		client: client,
		Search: &UserSearchService{client: client},
	}

	client.MySelf = &MySelfService{client: client}

	client.Workflow = &WorkflowService{
		client: client,
		Scheme: &WorkflowSchemeService{client: client},
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

	if c.Auth.basicAuthProvided {
		request.SetBasicAuth(c.Auth.mail, c.Auth.token)
	}

	if c.Auth.userAgentProvided {
		request.Header.Set("User-Agent", c.Auth.agent)
	}

	return
}

func (c *Client) call(request *http.Request, structure interface{}) (result *ResponseScheme, err error) {
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

		if response.ContentLength != 0 {

			responseAsBytes, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return responseTransformed, err
			}

			responseTransformed.Bytes.Write(responseAsBytes)
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
	Bytes    bytes.Buffer
	Headers  map[string][]string
}

var (
	requestCreationError    = "request creation failed: %v"
	urlParsedError          = "URL parsing failed: %v"
	requestFailedError      = "request failed. Please analyze the request body for more details. Status Code: %d"
	structureNotParsedError = errors.New("failed to parse the interface pointer, please provide a valid one")
)
