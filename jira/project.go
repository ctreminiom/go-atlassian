package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"net/url"
	"strconv"
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
	NotificationScheme  int    `json:"notificationScheme" validate:"required"`
	Description         string `json:"description"`
	LeadAccountID       string `json:"leadAccountId" validate:"required"`
	URL                 string `json:"url"`
	ProjectTemplateKey  string `json:"projectTemplateKey" validate:"required"`
	AvatarID            int    `json:"avatarId" validate:"required"`
	IssueSecurityScheme int    `json:"issueSecurityScheme" validate:"required"`
	Name                string `json:"name" validate:"required"`
	PermissionScheme    int    `json:"permissionScheme" validate:"required"`
	AssigneeType        string `json:"assigneeType" validate:"required"`
	ProjectTypeKey      string `json:"projectTypeKey" validate:"required"`
	Key                 string `json:"key" validate:"required"`
	CategoryID          int    `json:"categoryId" validate:"required"`
}

type NewProjectCreatedScheme struct {
	Self string `json:"self"`
	ID   int    `json:"id"`
	Key  string `json:"key"`
}

// Creates a project based on a project type template, as shown in the following table:
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-post
func (p *ProjectService) Create(ctx context.Context, payload *ProjectPayloadScheme) (result *NewProjectCreatedScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, please provide a ProjectPayloadScheme pointer")
	}

	validate := validator.New()
	if err = validate.Struct(payload); err != nil {
		return
	}

	var endpoint = "rest/api/3/project"

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(NewProjectCreatedScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ProjectSearchScheme struct {
	Self       string          `json:"self"`
	MaxResults int             `json:"maxResults"`
	StartAt    int             `json:"startAt"`
	Total      int             `json:"total"`
	IsLast     bool            `json:"isLast"`
	Values     []ProjectScheme `json:"values"`
}

type ProjectSearchOptionsScheme struct {
	OrderBy        string
	Query          string
	Action         string
	ProjectKeyType string
	CategoryID     int
	Expand         []string
}

// Returns a paginated list of projects visible to the user.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-search-get
func (p *ProjectService) Search(ctx context.Context, opts *ProjectSearchOptionsScheme, startAt, maxResults int) (result *ProjectSearchScheme, response *Response, err error) {

	if opts == nil {
		return nil, nil, fmt.Errorf("error, please provide a valid ProjectSearchOptionsScheme pointer")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var expand string
	for index, value := range opts.Expand {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	if len(opts.OrderBy) != 0 {
		params.Add("orderBy", opts.OrderBy)
	}

	if len(opts.Query) != 0 {
		params.Add("query", opts.Query)
	}

	if len(opts.ProjectKeyType) != 0 {
		params.Add("typeKey", opts.ProjectKeyType)
	}

	if opts.CategoryID != 0 {
		params.Add("categoryId", strconv.Itoa(opts.CategoryID))
	}

	if len(opts.Action) != 0 {
		params.Add("action", opts.Action)
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/search?%v", params.Encode())

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectSearchScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ProjectScheme struct {
	Expand         string                    `json:"expand,omitempty"`
	Self           string                    `json:"self,omitempty"`
	ID             string                    `json:"id,omitempty"`
	Key            string                    `json:"key,omitempty"`
	Description    string                    `json:"description,omitempty"`
	Lead           *UserScheme               `json:"lead,omitempty"`
	Components     []*ProjectComponentScheme `json:"components,omitempty"`
	IssueTypes     []*IssueTypeScheme        `json:"issueTypes,omitempty"`
	AssigneeType   string                    `json:"assigneeType,omitempty"`
	Versions       []*ProjectVersionScheme   `json:"versions,omitempty"`
	Name           string                    `json:"name,omitempty"`
	Roles          *ProjectRolesScheme       `json:"roles,omitempty"`
	AvatarUrls     *AvatarURLScheme          `json:"avatarUrls,omitempty"`
	ProjectKeys    []string                  `json:"projectKeys,omitempty"`
	ProjectTypeKey string                    `json:"projectTypeKey,omitempty"`
	Simplified     bool                      `json:"simplified,omitempty"`
	Style          string                    `json:"style,omitempty"`
	IsPrivate      bool                      `json:"isPrivate,omitempty"`
	Insight        *ProjectInsightScheme     `json:"insight,omitempty"`
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

// Returns the project details for a project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-get
func (p *ProjectService) Get(ctx context.Context, projectKeyOrID string, expands []string) (result *ProjectScheme, response *Response, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a projectKeyOrID value")
	}

	params := url.Values{}
	var expand string
	for index, value := range expands {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	var endpoint string
	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("rest/api/3/project/%v?%v", projectKeyOrID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/api/3/project/%v", projectKeyOrID)
	}

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
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

// Updates the project details of a project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-put
func (p *ProjectService) Update(ctx context.Context, projectKeyOrID string, payload *ProjectUpdateScheme) (result *ProjectScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, please provide a valid ProjectUpdateScheme pointer")
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes a project.
// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-delete
func (p *ProjectService) Delete(ctx context.Context, projectKeyOrID string, enableUndo bool) (response *Response, err error) {

	params := url.Values{}
	if enableUndo {
		params.Add("enableUndo", "true")
	}

	var endpoint string
	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("rest/api/3/project/%v?%v", projectKeyOrID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/api/3/project/%v", projectKeyOrID)
	}

	request, err := p.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Deletes a project asynchronously.
// 1. transactional, that is, if part of the delete fails the project is not deleted.
// 2. asynchronous. Follow the location link in the response to determine the status of the task and use Get task to obtain subsequent updates.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-delete-post
func (p *ProjectService) DeleteAsynchronously(ctx context.Context, projectKeyOrID string) (result *TaskScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/delete", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(TaskScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Archives a project. Archived projects cannot be deleted.
// To delete an archived project, restore the project and then delete it.
// To restore a project, use the Jira UI.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-archive-post
func (p *ProjectService) Archive(ctx context.Context, projectKeyOrID string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/archive", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Restores a project from the Jira recycle bin.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-restore-post
func (p *ProjectService) Restore(ctx context.Context, projectKeyOrID string) (result *ProjectScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/restore", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ProjectStatusScheme struct {
	Self     string `json:"self"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Subtask  bool   `json:"subtask"`
	Statuses []struct {
		Self             string `json:"self"`
		Description      string `json:"description"`
		IconURL          string `json:"iconUrl"`
		Name             string `json:"name"`
		UntranslatedName string `json:"untranslatedName"`
		ID               string `json:"id"`
		StatusCategory   struct {
			Self      string `json:"self"`
			ID        int    `json:"id"`
			Key       string `json:"key"`
			ColorName string `json:"colorName"`
			Name      string `json:"name"`
		} `json:"statusCategory"`
	} `json:"statuses"`
}

// Returns the valid statuses for a project.
// The statuses are grouped by issue type, as each project has a set of valid issue types and each issue type has a set of valid statuses.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-statuses-get
func (p *ProjectService) Statuses(ctx context.Context, projectKeyOrID string) (result *[]ProjectStatusScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/statuses", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new([]ProjectStatusScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
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

// Get the issue type hierarchy for a next-gen project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectid-hierarchy-get
func (p *ProjectService) Hierarchy(ctx context.Context, projectKeyOrID string) (result *ProjectIssueTypeHierarchyScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/hierarchy", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectIssueTypeHierarchyScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type NotificationSchemeScheme struct {
	Expand                   string `json:"expand"`
	ID                       int    `json:"id"`
	Self                     string `json:"self"`
	Name                     string `json:"name"`
	Description              string `json:"description"`
	NotificationSchemeEvents []struct {
		Event struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"event"`

		Notifications []struct {
			ID               int    `json:"id"`
			NotificationType string `json:"notificationType"`
			Parameter        string `json:"parameter,omitempty"`
			Group            struct {
				Name string `json:"name"`
				Self string `json:"self"`
			} `json:"group,omitempty"`
			Expand      string `json:"expand,omitempty"`
			ProjectRole struct {
				Self        string `json:"self"`
				Name        string `json:"name"`
				ID          int    `json:"id"`
				Description string `json:"description"`
				Actors      []struct {
					ID          int    `json:"id"`
					DisplayName string `json:"displayName"`
					Type        string `json:"type"`
					Name        string `json:"name,omitempty"`
					ActorGroup  struct {
						Name        string `json:"name"`
						DisplayName string `json:"displayName"`
					} `json:"actorGroup,omitempty"`
					ActorUser struct {
						AccountID string `json:"accountId"`
					} `json:"actorUser,omitempty"`
				} `json:"actors"`
				Scope struct {
					Type    string `json:"type"`
					Project struct {
						ID   string `json:"id"`
						Key  string `json:"key"`
						Name string `json:"name"`
					} `json:"project"`
				} `json:"scope"`
			} `json:"projectRole,omitempty"`
			EmailAddress string `json:"emailAddress,omitempty"`
			User         struct {
				Self        string `json:"self"`
				AccountID   string `json:"accountId"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
			} `json:"user,omitempty"`
			Field struct {
				ID               string   `json:"id"`
				Key              string   `json:"key"`
				Name             string   `json:"name"`
				UntranslatedName string   `json:"untranslatedName"`
				Custom           bool     `json:"custom"`
				Orderable        bool     `json:"orderable"`
				Navigable        bool     `json:"navigable"`
				Searchable       bool     `json:"searchable"`
				ClauseNames      []string `json:"clauseNames"`
				Schema           struct {
					Type     string `json:"type"`
					Custom   string `json:"custom"`
					CustomID int    `json:"customId"`
				} `json:"schema"`
			} `json:"field,omitempty"`
		} `json:"notifications"`
	} `json:"notificationSchemeEvents"`
}

// Search a notification scheme associated with the project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectkeyorid-notificationscheme-get
func (p *ProjectService) NotificationScheme(ctx context.Context, projectKeyOrID string, expands []string) (result *NotificationSchemeScheme, response *Response, err error) {

	params := url.Values{}
	var expand string

	for index, value := range expands {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	var endpoint string

	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("rest/api/3/project/%v/notificationscheme?%v", projectKeyOrID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/api/3/project/%v/notificationscheme", projectKeyOrID)
	}

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(NotificationSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
