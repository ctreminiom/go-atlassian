package v3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ctreminiom/go-atlassian/jira/internal"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service/common"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const ApiVersion = "3"

func New(httpClient common.HttpClient, site string) (*Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if site == "" {
		return nil, models.ErrNoSiteError
	}

	if !strings.HasSuffix(site, "/") {
		site += "/"
	}

	u, err := url.Parse(site)
	if err != nil {
		return nil, err
	}

	client := &Client{
		HTTP: httpClient,
		Site: u,
	}

	client.Auth = internal.NewAuthenticationService(client)

	auditRecord, err := internal.NewAuditRecordService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	applicationRoleService, err := internal.NewApplicationRoleService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	dashboardService, err := internal.NewDashboardService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	filterShareService, err := internal.NewFilterShareService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	filterService, err := internal.NewFilterService(client, ApiVersion, filterShareService)
	if err != nil {
		return nil, err
	}

	groupService, err := internal.NewGroupService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	issueAttachmentService, err := internal.NewIssueAttachmentService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	commentService, _, err := internal.NewCommentService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	fieldConfigurationItemService, err := internal.NewIssueFieldConfigurationItemService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	fieldConfigurationSchemeService, err := internal.NewIssueFieldConfigurationSchemeService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	fieldConfigService, err := internal.NewIssueFieldConfigurationService(client, ApiVersion, fieldConfigurationItemService, fieldConfigurationSchemeService)
	if err != nil {
		return nil, err
	}

	optionService, err := internal.NewIssueFieldContextOptionService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	fieldContextService, err := internal.NewIssueFieldContextService(client, ApiVersion, optionService)
	if err != nil {
		return nil, err
	}

	fieldTrashService, err := internal.NewIssueFieldTrashService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	issueFieldService, err := internal.NewIssueFieldService(client, ApiVersion, fieldConfigService, fieldContextService, fieldTrashService)
	if err != nil {
		return nil, err
	}

	label, err := internal.NewLabelService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	linkType, err := internal.NewLinkTypeService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	remoteLink, err := internal.NewRemoteLinkService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	link, _, err := internal.NewLinkService(client, ApiVersion, linkType, remoteLink)
	if err != nil {
		return nil, err
	}

	metadata, err := internal.NewMetadataService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	priority, err := internal.NewPriorityService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	resolution, err := internal.NewResolutionService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	search, _, err := internal.NewSearchService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	typeScheme, err := internal.NewTypeSchemeService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	issueTypeScreenScheme, err := internal.NewTypeScreenSchemeService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	type_, err := internal.NewTypeService(client, ApiVersion, typeScheme, issueTypeScreenScheme)
	if err != nil {
		return nil, err
	}

	vote, err := internal.NewVoteService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	watcher, err := internal.NewWatcherService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	worklog, err := internal.NewWorklogADFService(client, ApiVersion)
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

	mySelf, err := internal.NewMySelfService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	permissionSchemeGrant, err := internal.NewPermissionSchemeGrantService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	permissionScheme, err := internal.NewPermissionSchemeService(client, ApiVersion, permissionSchemeGrant)
	if err != nil {
		return nil, err
	}

	permission, err := internal.NewPermissionService(client, ApiVersion, permissionScheme)
	if err != nil {
		return nil, err
	}

	_, issueService, err := internal.NewIssueService(client, ApiVersion, issueServices)
	if err != nil {
		return nil, err
	}

	projectCategory, err := internal.NewProjectCategoryService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	projectComponent, err := internal.NewProjectComponentService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	projectFeature, err := internal.NewProjectFeatureService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	projectPermission, err := internal.NewProjectPermissionSchemeService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	projectProperties, err := internal.NewProjectPropertyService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	projectRoleActor, err := internal.NewProjectRoleActorService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	projectRole, err := internal.NewProjectRoleService(client, ApiVersion, projectRoleActor)
	if err != nil {
		return nil, err
	}

	projectType, err := internal.NewProjectTypeService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	projectValidator, err := internal.NewProjectValidatorService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	projectVersion, err := internal.NewProjectVersionService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	projectNotificationScheme, err := internal.NewNotificationSchemeService(client, ApiVersion)
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

	project, err := internal.NewProjectService(client, ApiVersion, projectSubService)
	if err != nil {
		return nil, err
	}

	screenFieldTabField, err := internal.NewScreenTabFieldService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	screenTab, err := internal.NewScreenTabService(client, ApiVersion, screenFieldTabField)
	if err != nil {
		return nil, err
	}

	screenScheme, err := internal.NewScreenSchemeService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	screen, err := internal.NewScreenService(client, ApiVersion, screenScheme, screenTab)
	if err != nil {
		return nil, err
	}

	task, err := internal.NewTaskService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	server, err := internal.NewServerService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	userSearch, err := internal.NewUserSearchService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	user, err := internal.NewUserService(client, ApiVersion, userSearch)
	if err != nil {
		return nil, err
	}

	workflowScheme := internal.NewWorkflowSchemeService(
		client,
		ApiVersion,
		internal.NewWorkflowSchemeIssueTypeService(client, ApiVersion))

	workflowStatus, err := internal.NewWorkflowStatusService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	workflow, err := internal.NewWorkflowService(client, ApiVersion, workflowScheme, workflowStatus)
	if err != nil {
		return nil, err
	}

	jql, err := internal.NewJQLService(client, ApiVersion)
	if err != nil {
		return nil, err
	}

	client.Audit = auditRecord
	client.Permission = permission
	client.MySelf = mySelf
	client.Auth = internal.NewAuthenticationService(client)
	client.Banner = internal.NewAnnouncementBannerService(client, ApiVersion)
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
	client.JQL = jql
	client.NotificationScheme = projectNotificationScheme
	client.Team = internal.NewTeamService(client)

	return client, nil
}

type Client struct {
	HTTP               common.HttpClient
	Auth               common.Authentication
	Site               *url.URL
	Audit              *internal.AuditRecordService
	Role               *internal.ApplicationRoleService
	Banner             *internal.AnnouncementBannerService
	Dashboard          *internal.DashboardService
	Filter             *internal.FilterService
	Group              *internal.GroupService
	Issue              *internal.IssueADFService
	MySelf             *internal.MySelfService
	Permission         *internal.PermissionService
	Project            *internal.ProjectService
	Screen             *internal.ScreenService
	Task               *internal.TaskService
	Server             *internal.ServerService
	User               *internal.UserService
	Workflow           *internal.WorkflowService
	JQL                *internal.JQLService
	NotificationScheme *internal.NotificationSchemeService
	Team               *internal.TeamService
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

	// If the body interface is a *bytes.Buffer type
	// it means the NewRequest() requires to handle the RFC 1867 ISO
	if attachBuffer, ok := body.(*bytes.Buffer); ok {
		buf = attachBuffer
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if type_ != "" {
		// When the type_ is provided, it means the request needs to be created to handle files
		req.Header.Set("Content-Type", type_)
		req.Header.Set("X-Atlassian-Token", "no-check")
	}

	if c.Auth.HasBasicAuth() {
		req.SetBasicAuth(c.Auth.GetBasicAuth())
	}

	if c.Auth.HasUserAgent() {
		req.Header.Set("User-Agent", c.Auth.GetUserAgent())
	}

	if c.Auth.GetBearerToken() != "" && !c.Auth.HasBasicAuth() {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.Auth.GetBearerToken()))
	}

	return req, nil
}

func (c *Client) Call(request *http.Request, structure interface{}) (*models.ResponseScheme, error) {

	response, err := c.HTTP.Do(request)
	if err != nil {
		return nil, err
	}

	return c.processResponse(response, structure)
}

func (c *Client) processResponse(response *http.Response, structure interface{}) (*models.ResponseScheme, error) {

	defer response.Body.Close()

	res := &models.ResponseScheme{
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
			return res, models.ErrNotFound

		case http.StatusUnauthorized:
			return res, models.ErrUnauthorized

		case http.StatusInternalServerError:
			return res, models.ErrInternalError

		case http.StatusBadRequest:
			return res, models.ErrBadRequestError

		default:
			return res, models.ErrInvalidStatusCodeError
		}
	}

	if structure != nil {
		if err = json.Unmarshal(responseAsBytes, &structure); err != nil {
			return res, err
		}
	}

	return res, nil
}
