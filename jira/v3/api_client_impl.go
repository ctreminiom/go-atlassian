package v3

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/ctreminiom/go-atlassian/jira/internal"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service/common"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

func New(httpClient common.HttpClient, site string) (*Client, error) {

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

	client := &Client{
		HTTP: httpClient,
		Site: siteAsURL,
	}

	applicationRoleService, err := internal.NewApplicationRoleService(client, "3")
	if err != nil {
		return nil, err
	}

	dashboardService, err := internal.NewDashboardService(client, "3")
	if err != nil {
		return nil, err
	}

	filterShareService, err := internal.NewFilterShareService(client, "3")
	if err != nil {
		return nil, err
	}

	filterService, err := internal.NewFilterService(client, "3", filterShareService)
	if err != nil {
		return nil, err
	}

	groupService, err := internal.NewGroupService(client, "3")
	if err != nil {
		return nil, err
	}

	issueAttachmentService, err := internal.NewIssueAttachmentService(client, "3")
	if err != nil {
		return nil, err
	}

	commentService, _, err := internal.NewCommentService(client, "3")
	if err != nil {
		return nil, err
	}

	fieldConfigurationItemService, err := internal.NewIssueFieldConfigurationItemService(client, "3")
	if err != nil {
		return nil, err
	}

	fieldConfigurationSchemeService, err := internal.NewIssueFieldConfigurationSchemeService(client, "3")
	if err != nil {
		return nil, err
	}

	fieldConfigService, err := internal.NewIssueFieldConfigurationService(client, "3", fieldConfigurationItemService, fieldConfigurationSchemeService)
	if err != nil {
		return nil, err
	}

	optionService, err := internal.NewIssueFieldContextOptionService(client, "3")
	if err != nil {
		return nil, err
	}

	fieldContextService, err := internal.NewIssueFieldContextService(client, "3", optionService)
	if err != nil {
		return nil, err
	}

	fieldTrashService, err := internal.NewIssueFieldTrashService(client, "3")
	if err != nil {
		return nil, err
	}

	issueFieldService, err := internal.NewIssueFieldService(client, "3", fieldConfigService, fieldContextService, fieldTrashService)
	if err != nil {
		return nil, err
	}

	label, err := internal.NewLabelService(client, "3")
	if err != nil {
		return nil, err
	}

	linkType, err := internal.NewLinkTypeService(client, "3")
	if err != nil {
		return nil, err
	}

	link, _, err := internal.NewLinkService(client, "3", linkType)
	if err != nil {
		return nil, err
	}

	metadata, err := internal.NewMetadataService(client, "3")
	if err != nil {
		return nil, err
	}

	priority, err := internal.NewPriorityService(client, "3")
	if err != nil {
		return nil, err
	}

	resolution, err := internal.NewResolutionService(client, "3")
	if err != nil {
		return nil, err
	}

	search, _, err := internal.NewSearchService(client, "3")
	if err != nil {
		return nil, err
	}

	typeScheme, err := internal.NewTypeSchemeService(client, "3")
	if err != nil {
		return nil, err
	}

	issueTypeScreenScheme, err := internal.NewTypeScreenSchemeService(client, "3")
	if err != nil {
		return nil, err
	}

	type_, err := internal.NewTypeService(client, "3", typeScheme, issueTypeScreenScheme)
	if err != nil {
		return nil, err
	}

	vote, err := internal.NewVoteService(client, "3")
	if err != nil {
		return nil, err
	}

	watcher, err := internal.NewWatcherService(client, "3")
	if err != nil {
		return nil, err
	}

	worklog, err := internal.NewWorklogADFService(client, "3")
	if err != nil {
		return nil, err
	}

	issueServices := &internal.IssueServices{
		Attachment: issueAttachmentService,
		CommentADF: commentService,
		Field:      issueFieldService,
		Label:      label,
		LinkADF:    link,
		Metadata:   metadata,
		Priority:   priority,
		Resolution: resolution,
		SearchADF:  search,
		Type:       type_,
		Vote:       vote,
		Watcher:    watcher,
		WorklogAdf: worklog,
	}

	mySelf, err := internal.NewMySelfService(client, "3")
	if err != nil {
		return nil, err
	}

	permissionSchemeGrant, err := internal.NewPermissionSchemeGrantService(client, "3")
	if err != nil {
		return nil, err
	}

	permissionScheme, err := internal.NewPermissionSchemeService(client, "3", permissionSchemeGrant)
	if err != nil {
		return nil, err
	}

	permission, err := internal.NewPermissionService(client, "3", permissionScheme)
	if err != nil {
		return nil, err
	}

	_, issueService, err := internal.NewIssueService(client, "3", issueServices)
	if err != nil {
		return nil, err
	}

	projectCategory, err := internal.NewProjectCategoryService(client, "3")
	if err != nil {
		return nil, err
	}

	projectComponent, err := internal.NewProjectComponentService(client, "3")
	if err != nil {
		return nil, err
	}

	projectFeature, err := internal.NewProjectFeatureService(client, "3")
	if err != nil {
		return nil, err
	}

	projectPermission, err := internal.NewProjectPermissionSchemeService(client, "3")
	if err != nil {
		return nil, err
	}

	projectProperties, err := internal.NewProjectPropertyService(client, "3")
	if err != nil {
		return nil, err
	}

	projectRoleActor, err := internal.NewProjectRoleActorService(client, "3")
	if err != nil {
		return nil, err
	}

	projectRole, err := internal.NewProjectRoleService(client, "3", projectRoleActor)
	if err != nil {
		return nil, err
	}

	projectType, err := internal.NewProjectTypeService(client, "3")
	if err != nil {
		return nil, err
	}

	projectValidator, err := internal.NewProjectValidatorService(client, "3")
	if err != nil {
		return nil, err
	}

	projectVersion, err := internal.NewProjectVersionService(client, "3")
	if err != nil {
		return nil, err
	}

	projectSubService := &internal.ProjectChildServices{
		Category:   projectCategory,
		Component:  projectComponent,
		Feature:    projectFeature,
		Permission: projectPermission,
		Property:   projectProperties,
		Role:       projectRole,
		Type:       projectType,
		Validator:  projectValidator,
		Version:    projectVersion,
	}

	project, err := internal.NewProjectService(client, "3", projectSubService)
	if err != nil {
		return nil, err
	}

	screenFieldTabField, err := internal.NewScreenTabFieldService(client, "3")
	if err != nil {
		return nil, err
	}

	screenTab, err := internal.NewScreenTabService(client, "3", screenFieldTabField)
	if err != nil {
		return nil, err
	}

	screenScheme, err := internal.NewScreenSchemeService(client, "3")
	if err != nil {
		return nil, err
	}

	screen, err := internal.NewScreenService(client, "3", screenScheme, screenTab)
	if err != nil {
		return nil, err
	}

	task, err := internal.NewTaskService(client, "3")
	if err != nil {
		return nil, err
	}

	server, err := internal.NewServerService(client, "3")
	if err != nil {
		return nil, err
	}

	userSearch, err := internal.NewUserSearchService(client, "3")
	if err != nil {
		return nil, err
	}

	user, err := internal.NewUserService(client, "3", userSearch)
	if err != nil {
		return nil, err
	}

	workflowScheme, err := internal.NewWorkflowSchemeService(client, "3")
	if err != nil {
		return nil, err
	}

	workflowStatus, err := internal.NewWorkflowStatusService(client, "3")
	if err != nil {
		return nil, err
	}

	workflow, err := internal.NewWorkflowService(client, "3", workflowScheme, workflowStatus)
	if err != nil {
		return nil, err
	}

	client.Permission = permission
	client.MySelf = mySelf
	client.Auth = internal.NewAuthenticationService(client)
	client.Role = applicationRoleService
	client.Dashboard = dashboardService
	client.Filter = filterService
	client.Group = groupService
	client.Issue = issueService
	client.Project = project
	client.Screen = screen
	client.Task = task
	client.Server = server
	client.User = user
	client.Workflow = workflow

	return client, nil
}

type Client struct {
	HTTP       common.HttpClient
	Auth       common.Authentication
	Site       *url.URL
	Role       *internal.ApplicationRoleService
	Dashboard  *internal.DashboardService
	Filter     *internal.FilterService
	Group      *internal.GroupService
	Issue      *internal.IssueADFService
	MySelf     *internal.MySelfService
	Permission *internal.PermissionService
	Project    *internal.ProjectService
	Screen     *internal.ScreenService
	Task       *internal.TaskService
	Server     *internal.ServerService
	User       *internal.UserService
	Workflow   *internal.WorkflowService
}

func (c *Client) NewFormRequest(ctx context.Context, method, apiEndpoint, contentType string, payload io.Reader) (*http.Request, error) {

	relativePath, err := url.Parse(apiEndpoint)
	if err != nil {
		return nil, err
	}

	var endpoint = c.Site.ResolveReference(relativePath).String()

	request, err := http.NewRequestWithContext(ctx, method, endpoint, payload)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept", "application/json")
	request.Header.Set("X-Atlassian-Token", "no-check")

	if c.Auth.HasBasicAuth() {
		request.SetBasicAuth(c.Auth.GetBasicAuth())
	}

	if c.Auth.HasUserAgent() {
		request.Header.Set("User-Agent", c.Auth.GetUserAgent())
	}

	return request, nil
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

	if c.Auth.HasBasicAuth() {
		request.SetBasicAuth(c.Auth.GetBasicAuth())
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

	responseAsBytes, err := ioutil.ReadAll(response.Body)
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
