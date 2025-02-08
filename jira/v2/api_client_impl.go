package v2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/ctreminiom/go-atlassian/v2/jira/internal"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

// APIVersion is the version of the Jira API that this client targets.
const APIVersion = "2"

// New creates a new Jira API client.
// If a nil httpClient is provided, http.DefaultClient will be used.
// If the site is empty, an error will be returned.
func New(httpClient common.HTTPClient, site string) (*Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if site == "" {
		return nil, models.ErrNoSite
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

	auditRecordService, err := internal.NewAuditRecordService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	applicationRoleService, err := internal.NewApplicationRoleService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	dashboardService, err := internal.NewDashboardService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	filterShareService, err := internal.NewFilterShareService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	filterService, err := internal.NewFilterService(client, APIVersion, filterShareService)
	if err != nil {
		return nil, err
	}

	groupUserPickerService, err := internal.NewGroupUserPickerService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	groupService, err := internal.NewGroupService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	issueAttachmentService, err := internal.NewIssueAttachmentService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	_, commentService, err := internal.NewCommentService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	fieldConfigurationItemService, err := internal.NewIssueFieldConfigurationItemService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	fieldConfigurationSchemeService, err := internal.NewIssueFieldConfigurationSchemeService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	fieldConfigService, err := internal.NewIssueFieldConfigurationService(client, APIVersion, fieldConfigurationItemService, fieldConfigurationSchemeService)
	if err != nil {
		return nil, err
	}

	optionService, err := internal.NewIssueFieldContextOptionService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	fieldContextService, err := internal.NewIssueFieldContextService(client, APIVersion, optionService)
	if err != nil {
		return nil, err
	}

	fieldTrashService, err := internal.NewIssueFieldTrashService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	issueFieldService, err := internal.NewIssueFieldService(client, APIVersion, fieldConfigService, fieldContextService, fieldTrashService)
	if err != nil {
		return nil, err
	}

	label, err := internal.NewLabelService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	linkType, err := internal.NewLinkTypeService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	remoteLink, err := internal.NewRemoteLinkService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	_, link, err := internal.NewLinkService(client, APIVersion, linkType, remoteLink)
	if err != nil {
		return nil, err
	}

	metadata, err := internal.NewMetadataService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	priority, err := internal.NewPriorityService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	resolution, err := internal.NewResolutionService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	_, search, err := internal.NewSearchService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	typeScheme, err := internal.NewTypeSchemeService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	issueTypeScreenScheme, err := internal.NewTypeScreenSchemeService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	typ, err := internal.NewTypeService(client, APIVersion, typeScheme, issueTypeScreenScheme)
	if err != nil {
		return nil, err
	}

	vote, err := internal.NewVoteService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	watcher, err := internal.NewWatcherService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	worklog, err := internal.NewWorklogRichTextService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	issueProperty, err := internal.NewIssuePropertyService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	issueServices := &internal.IssueServices{
		Attachment:      issueAttachmentService,
		CommentRT:       commentService,
		Field:           issueFieldService,
		Label:           label,
		LinkRT:          link,
		Metadata:        metadata,
		Priority:        priority,
		Resolution:      resolution,
		SearchRT:        search,
		Type:            typ,
		Vote:            vote,
		Watcher:         watcher,
		WorklogRichText: worklog,
		Property:        issueProperty,
	}

	issueService, _, err := internal.NewIssueService(client, APIVersion, issueServices)
	if err != nil {
		return nil, err
	}

	mySelf, err := internal.NewMySelfService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	permissionSchemeGrant, err := internal.NewPermissionSchemeGrantService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	permissionScheme, err := internal.NewPermissionSchemeService(client, APIVersion, permissionSchemeGrant)
	if err != nil {
		return nil, err
	}

	permission, err := internal.NewPermissionService(client, APIVersion, permissionScheme)
	if err != nil {
		return nil, err
	}

	projectCategory, err := internal.NewProjectCategoryService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	projectComponent, err := internal.NewProjectComponentService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	projectFeature, err := internal.NewProjectFeatureService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	projectPermission, err := internal.NewProjectPermissionSchemeService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	projectProperties, err := internal.NewProjectPropertyService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	projectRoleActor, err := internal.NewProjectRoleActorService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	projectRole, err := internal.NewProjectRoleService(client, APIVersion, projectRoleActor)
	if err != nil {
		return nil, err
	}

	projectType, err := internal.NewProjectTypeService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	projectValidator, err := internal.NewProjectValidatorService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	projectVersion, err := internal.NewProjectVersionService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	projectNotificationScheme, err := internal.NewNotificationSchemeService(client, APIVersion)
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

	project, err := internal.NewProjectService(client, APIVersion, projectSubService)
	if err != nil {
		return nil, err
	}

	screenFieldTabField, err := internal.NewScreenTabFieldService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	screenTab, err := internal.NewScreenTabService(client, APIVersion, screenFieldTabField)
	if err != nil {
		return nil, err
	}

	screenScheme, err := internal.NewScreenSchemeService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	screen, err := internal.NewScreenService(client, APIVersion, screenScheme, screenTab)
	if err != nil {
		return nil, err
	}

	task, err := internal.NewTaskService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	server, err := internal.NewServerService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	userSearch, err := internal.NewUserSearchService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	user, err := internal.NewUserService(client, APIVersion, userSearch)
	if err != nil {
		return nil, err
	}

	workflowScheme := internal.NewWorkflowSchemeService(
		client,
		APIVersion,
		internal.NewWorkflowSchemeIssueTypeService(client, APIVersion))

	workflowStatus, err := internal.NewWorkflowStatusService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	workflow, err := internal.NewWorkflowService(client, APIVersion, workflowScheme, workflowStatus)
	if err != nil {
		return nil, err
	}

	jql, err := internal.NewJQLService(client, APIVersion)
	if err != nil {
		return nil, err
	}

	client.Audit = auditRecordService
	client.Permission = permission
	client.MySelf = mySelf
	client.Auth = internal.NewAuthenticationService(client)
	client.Banner = internal.NewAnnouncementBannerService(client, APIVersion)
	client.Role = applicationRoleService
	client.Dashboard = dashboardService
	client.Filter = filterService
	client.GroupUserPicker = groupUserPickerService
	client.Group = groupService
	client.Issue = issueService
	client.Project = project
	client.Screen = screen
	client.Server = server
	client.Task = task
	client.User = user
	client.Workflow = workflow
	client.JQL = jql
	client.NotificationScheme = projectNotificationScheme
	client.Team = internal.NewTeamService(client)

	client.Archive = internal.NewIssueArchivalService(client, APIVersion)

	return client, nil
}

type Client struct {
	HTTP               common.HTTPClient
	Auth               common.Authentication
	Site               *url.URL
	Role               *internal.ApplicationRoleService
	Banner             *internal.AnnouncementBannerService
	Audit              *internal.AuditRecordService
	Dashboard          *internal.DashboardService
	Filter             *internal.FilterService
	Group              *internal.GroupService
	GroupUserPicker    *internal.GroupUserPickerService
	Issue              *internal.IssueRichTextService
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

	Archive *internal.IssueArchivalService
}

// NewRequest creates an API request.
func (c *Client) NewRequest(ctx context.Context, method, urlStr, contentType string, body interface{}) (*http.Request, error) {

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

	if contentType != "" {
		// When the contentType is provided, it means the request needs to be created to handle files
		req.Header.Set("Content-Type", contentType)
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
			return res, models.ErrInternal

		case http.StatusBadRequest:
			return res, models.ErrBadRequest

		default:
			return res, models.ErrInvalidStatusCode
		}
	}

	if structure != nil {
		if err = json.Unmarshal(responseAsBytes, &structure); err != nil {
			return res, err
		}
	}

	return res, nil
}
