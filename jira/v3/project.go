package v3

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ProjectService struct {
	client     *Client
	Category   *ProjectCategoryService
	Component  *ProjectComponentService
	Valid      *ProjectValidationService
	Permission *ProjectPermissionSchemeService
	Role       *ProjectRoleService
	Type       *ProjectTypeService
	Version    *ProjectVersionService
}

type ProjectPayloadScheme struct {
	NotificationScheme  int    `json:"notificationScheme"`
	Description         string `json:"description"`
	LeadAccountID       string `json:"leadAccountId"`
	URL                 string `json:"url"`
	ProjectTemplateKey  string `json:"projectTemplateKey"`
	AvatarID            int    `json:"avatarId"`
	IssueSecurityScheme int    `json:"issueSecurityScheme"`
	Name                string `json:"name"`
	PermissionScheme    int    `json:"permissionScheme"`
	AssigneeType        string `json:"assigneeType"`
	ProjectTypeKey      string `json:"projectTypeKey"`
	Key                 string `json:"key"`
	CategoryID          int    `json:"categoryId"`
}

type NewProjectCreatedScheme struct {
	Self string `json:"self"`
	ID   int    `json:"id"`
	Key  string `json:"key"`
}

const (
	BusinessContentManagement    = "com.atlassian.jira-core-project-templates:jira-core-simplified-content-management"
	BusinessDocumentApproval     = "com.atlassian.jira-core-project-templates:jira-core-simplified-document-approval"
	BusinessLeadTracking         = "com.atlassian.jira-core-project-templates:jira-core-simplified-lead-tracking"
	BusinessProcessControl       = "com.atlassian.jira-core-project-templates:jira-core-simplified-process-control"
	BusinessProcurement          = "com.atlassian.jira-core-project-templates:jira-core-simplified-procurement"
	BusinessProjectManagement    = "com.atlassian.jira-core-project-templates:jira-core-simplified-project-management"
	BusinessRecruitment          = "com.atlassian.jira-core-project-templates:jira-core-simplified-recruitment"
	BusinessTaskTracking         = "com.atlassian.jira-core-project-templates:jira-core-simplified-task-tracking"
	ITSMServiceDesk              = "com.atlassian.servicedesk:simplified-it-service-desk"
	ITSMInternalServiceDesk      = "com.atlassian.servicedesk:simplified-internal-service-desk"
	ITSMExternalServiceDesk      = "com.atlassian.servicedesk:simplified-external-service-desk"
	SoftwareTeamManagedKanban    = "com.pyxis.greenhopper.jira:gh-simplified-agility-kanban"
	SoftwareTeamManagedScrum     = "com.pyxis.greenhopper.jira:gh-simplified-agility-scrum"
	SoftwareCompanyManagedKanban = "com.pyxis.greenhopper.jira:gh-simplified-kanban-classic"
	SoftwareCompanyManagedScrum  = "com.pyxis.greenhopper.jira:gh-simplified-scrum-classic"
)

// Create creates a project based on a project type template, as shown in the following table:
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-post
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-post
func (p *ProjectService) Create(ctx context.Context, payload *ProjectPayloadScheme) (result *NewProjectCreatedScheme,
	response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/project"

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type ProjectSearchScheme struct {
	Self       string           `json:"self,omitempty"`
	NextPage   string           `json:"nextPage,omitempty"`
	MaxResults int              `json:"maxResults,omitempty"`
	StartAt    int              `json:"startAt,omitempty"`
	Total      int              `json:"total,omitempty"`
	IsLast     bool             `json:"isLast,omitempty"`
	Values     []*ProjectScheme `json:"values,omitempty"`
}

type ProjectSearchOptionsScheme struct {
	OrderBy        string
	Query          string
	Action         string
	ProjectKeyType string
	CategoryID     int
	Expand         []string
}

// Search returns a paginated list of projects visible to the user.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-search-get
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-search-get
func (p *ProjectService) Search(ctx context.Context, options *ProjectSearchOptionsScheme, startAt, maxResults int) (
	result *ProjectSearchScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}

		if len(options.OrderBy) != 0 {
			params.Add("orderBy", options.OrderBy)
		}

		if len(options.Query) != 0 {
			params.Add("query", options.Query)
		}

		if len(options.ProjectKeyType) != 0 {
			params.Add("typeKey", options.ProjectKeyType)
		}

		if options.CategoryID != 0 {
			params.Add("categoryId", strconv.Itoa(options.CategoryID))
		}

		if len(options.Action) != 0 {
			params.Add("action", options.Action)
		}
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/search?%v", params.Encode())

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type ProjectScheme struct {
	Expand            string                    `json:"expand,omitempty"`
	Self              string                    `json:"self,omitempty"`
	ID                string                    `json:"id,omitempty"`
	Key               string                    `json:"key,omitempty"`
	Description       string                    `json:"description,omitempty"`
	URL               string                    `json:"url,omitempty"`
	Email             string                    `json:"email,omitempty"`
	AssigneeType      string                    `json:"assigneeType,omitempty"`
	Name              string                    `json:"name,omitempty"`
	ProjectTypeKey    string                    `json:"projectTypeKey,omitempty"`
	Simplified        bool                      `json:"simplified,omitempty"`
	Style             string                    `json:"style,omitempty"`
	Favourite         bool                      `json:"favourite,omitempty"`
	IsPrivate         bool                      `json:"isPrivate,omitempty"`
	UUID              string                    `json:"uuid,omitempty"`
	Lead              *UserScheme               `json:"lead,omitempty"`
	Components        []*ProjectComponentScheme `json:"components,omitempty"`
	IssueTypes        []*IssueTypeScheme        `json:"issueTypes,omitempty"`
	Versions          []*ProjectVersionScheme   `json:"versions,omitempty"`
	Roles             *ProjectRolesScheme       `json:"roles,omitempty"`
	AvatarUrls        *AvatarURLScheme          `json:"avatarUrls,omitempty"`
	ProjectKeys       []string                  `json:"projectKeys,omitempty"`
	Insight           *ProjectInsightScheme     `json:"insight,omitempty"`
	Category          *ProjectCategoryScheme    `json:"projectCategory,omitempty"`
	Deleted           bool                      `json:"deleted,omitempty"`
	RetentionTillDate string                    `json:"retentionTillDate,omitempty"`
	DeletedDate       string                    `json:"deletedDate,omitempty"`
	DeletedBy         *UserScheme               `json:"deletedBy,omitempty"`
	Archived          bool                      `json:"archived,omitempty"`
	ArchivedDate      string                    `json:"archivedDate,omitempty"`
	ArchivedBy        *UserScheme               `json:"archivedBy,omitempty"`
}

type ProjectInsightScheme struct {
	TotalIssueCount     int    `json:"totalIssueCount,omitempty"`
	LastIssueUpdateTime string `json:"lastIssueUpdateTime,omitempty"`
}

type AvatarURLScheme struct {
	Four8X48  string `json:"48x48,omitempty"`
	Two4X24   string `json:"24x24,omitempty"`
	One6X16   string `json:"16x16,omitempty"`
	Three2X32 string `json:"32x32,omitempty"`
}

type ProjectRolesScheme struct {
	AtlassianAddonsProjectAccess string `json:"atlassian-addons-project-access,omitempty"`
	ServiceDeskTeam              string `json:"Service Desk Team,omitempty"`
	ServiceDeskCustomers         string `json:"Service Desk Customers,omitempty"`
	Administrators               string `json:"Administrators,omitempty"`
}

// Get returns the project details for a project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-get
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-get
func (p *ProjectService) Get(ctx context.Context, projectKeyOrID string, expand []string) (result *ProjectScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, notProjectIDError
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/3/project/%v", projectKeyOrID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type ProjectUpdateScheme struct {
	NotificationScheme  int    `json:"notificationScheme,omitempty"`
	Description         string `json:"description,omitempty"`
	Lead                string `json:"lead,omitempty"`
	URL                 string `json:"url,omitempty"`
	ProjectTemplateKey  string `json:"projectTemplateKey,omitempty"`
	AvatarID            int    `json:"avatarId,omitempty"`
	IssueSecurityScheme int    `json:"issueSecurityScheme,omitempty"`
	Name                string `json:"name,omitempty"`
	PermissionScheme    int    `json:"permissionScheme,omitempty"`
	AssigneeType        string `json:"assigneeType,omitempty"`
	ProjectTypeKey      string `json:"projectTypeKey,omitempty"`
	Key                 string `json:"key,omitempty"`
	CategoryID          int    `json:"categoryId,omitempty"`
}

// Update updates the project details of a project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-put
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-put
func (p *ProjectService) Update(ctx context.Context, projectKeyOrID string, payload *ProjectUpdateScheme) (result *ProjectScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, notProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v", projectKeyOrID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := p.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes a project.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-delete
func (p *ProjectService) Delete(ctx context.Context, projectKeyOrID string, enableUndo bool) (response *ResponseScheme,
	err error) {

	if len(projectKeyOrID) == 0 {
		return nil, notProjectIDError
	}

	params := url.Values{}
	if enableUndo {
		params.Add("enableUndo", "true")
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/3/project/%v", projectKeyOrID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := p.client.newRequest(ctx, http.MethodDelete, endpoint.String(), nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// DeleteAsynchronously deletes a project asynchronously.
// 1. transactional, that is, if part of the delete fails the project is not deleted.
// 2. asynchronous. Follow the location link in the response to determine the status of the task and use Get task to obtain subsequent updates.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-delete-post
func (p *ProjectService) DeleteAsynchronously(ctx context.Context, projectKeyOrID string) (result *TaskScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, notProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/delete", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Archive archives a project. Archived projects cannot be deleted.
// To delete an archived project, restore the project and then delete it.
// To restore a project, use the Jira UI.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-archive-post
func (p *ProjectService) Archive(ctx context.Context, projectKeyOrID string) (response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, notProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/archive", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Restore restores a project from the Jira recycle bin.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-restore-post
func (p *ProjectService) Restore(ctx context.Context, projectKeyOrID string) (result *ProjectScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, notProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/restore", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type ProjectStatusPageScheme struct {
	Self     string                        `json:"self,omitempty"`
	ID       string                        `json:"id,omitempty"`
	Name     string                        `json:"name,omitempty"`
	Subtask  bool                          `json:"subtask,omitempty"`
	Statuses []*ProjectStatusDetailsScheme `json:"statuses,omitempty"`
}

type ProjectStatusDetailsScheme struct {
	Self           string                `json:"self,omitempty"`
	Description    string                `json:"description,omitempty"`
	IconURL        string                `json:"iconUrl,omitempty"`
	Name           string                `json:"name,omitempty"`
	ID             string                `json:"id,omitempty"`
	StatusCategory *StatusCategoryScheme `json:"statusCategory,omitempty"`
}

// Statuses returns the valid statuses for a project.
// The statuses are grouped by issue type, as each project has a set of valid issue types and each issue type has a set of valid statuses.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-statuses-get
func (p *ProjectService) Statuses(ctx context.Context, projectKeyOrID string) (result []*ProjectStatusPageScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, notProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/statuses", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type ProjectIssueTypeHierarchyScheme struct {
	ProjectID int `json:"projectId"`
	Hierarchy []struct {
		EntityID   string `json:"entityId"`
		Level      int    `json:"level"`
		Name       string `json:"name"`
		IssueTypes []struct {
			ID       int    `json:"id"`
			EntityID string `json:"entityId"`
			Name     string `json:"name"`
			AvatarID int    `json:"avatarId"`
		} `json:"issueTypes"`
	} `json:"hierarchy"`
}

// Hierarchy get the issue type hierarchy for a next-gen project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectid-hierarchy-get
func (p *ProjectService) Hierarchy(ctx context.Context, projectKeyOrID string) (result *ProjectIssueTypeHierarchyScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, notProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/hierarchy", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type NotificationSchemeScheme struct {
	Expand                   string                                  `json:"expand,omitempty"`
	ID                       int                                     `json:"id,omitempty"`
	Self                     string                                  `json:"self,omitempty"`
	Name                     string                                  `json:"name,omitempty"`
	Description              string                                  `json:"description,omitempty"`
	NotificationSchemeEvents []*ProjectNotificationSchemeEventScheme `json:"notificationSchemeEvents,omitempty"`
	Scope                    *TeamManagedProjectScopeScheme          `json:"scope,omitempty"`
}

type ProjectNotificationSchemeEventScheme struct {
	Event         *NotificationEventScheme   `json:"event,omitempty"`
	Notifications []*EventNotificationScheme `json:"notifications,omitempty"`
}

type NotificationEventScheme struct {
	ID            int                      `json:"id,omitempty"`
	Name          string                   `json:"name,omitempty"`
	Description   string                   `json:"description,omitempty"`
	TemplateEvent *NotificationEventScheme `json:"templateEvent,omitempty"`
}

type EventNotificationScheme struct {
	Expand           string             `json:"expand,omitempty"`
	ID               int                `json:"id,omitempty"`
	NotificationType string             `json:"notificationType,omitempty"`
	Parameter        string             `json:"parameter,omitempty"`
	EmailAddress     string             `json:"emailAddress,omitempty"`
	Group            *GroupScheme       `json:"group,omitempty"`
	Field            *IssueFieldScheme  `json:"field,omitempty"`
	ProjectRole      *ProjectRoleScheme `json:"projectRole,omitempty"`
	User             *UserScheme        `json:"user,omitempty"`
}

// NotificationScheme search a notification scheme associated with the project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectkeyorid-notificationscheme-get
func (p *ProjectService) NotificationScheme(ctx context.Context, projectKeyOrID string, expand []string) (
	result *NotificationSchemeScheme, response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, notProjectIDError
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/3/project/%v/notificationscheme", projectKeyOrID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
