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
	Property   *ProjectPropertyService
	Role       *ProjectRoleService
	Type       *ProjectTypeService
	Version    *ProjectVersionService
}

type ProjectPayloadScheme struct {
	NotificationScheme  int    `json:"notificationScheme" validate:"required"`
	Description         string `json:"description" validate:"required"`
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

type ProjectSearchOptionsScheme struct {
	OrderBy        string
	Query          string
	ProjectKeyType string
	CategoryID     int
	SearchBy       string
	Action         string
	Expand         []string
}

type ProjectSearchScheme struct {
	Self       string `json:"self"`
	MaxResults int    `json:"maxResults"`
	StartAt    int    `json:"startAt"`
	Total      int    `json:"total"`
	IsLast     bool   `json:"isLast"`
	Values     []struct {
		Expand      string `json:"expand"`
		Self        string `json:"self"`
		ID          string `json:"id"`
		Key         string `json:"key"`
		Description string `json:"description"`
		Lead        struct {
			Self       string `json:"self"`
			AccountID  string `json:"accountId"`
			AvatarUrls struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
		} `json:"lead"`
		IssueTypes []struct {
			Self        string `json:"self"`
			ID          string `json:"id"`
			Description string `json:"description"`
			IconURL     string `json:"iconUrl"`
			Name        string `json:"name"`
			Subtask     bool   `json:"subtask"`
			AvatarID    int    `json:"avatarId"`
		} `json:"issueTypes"`
		Name       string `json:"name"`
		AvatarUrls struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		ProjectKeys    []string `json:"projectKeys"`
		ProjectTypeKey string   `json:"projectTypeKey"`
		Simplified     bool     `json:"simplified"`
		Style          string   `json:"style"`
		IsPrivate      bool     `json:"isPrivate"`
		Properties     struct {
		} `json:"properties"`
		Insight struct {
			TotalIssueCount     int    `json:"totalIssueCount"`
			LastIssueUpdateTime string `json:"lastIssueUpdateTime"`
		} `json:"insight"`
	} `json:"values"`
}

// Returns a paginated list of projects visible to the user.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-search-get
func (p *ProjectService) Search(ctx context.Context, opts *ProjectSearchOptionsScheme, startAt, maxResults int) (result *ProjectSearchScheme, response *Response, err error) {

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

	if len(opts.SearchBy) != 0 {
		params.Add("searchBy", opts.SearchBy)
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
	Expand      string `json:"expand"`
	Self        string `json:"self"`
	ID          string `json:"id"`
	Key         string `json:"key"`
	Description string `json:"description"`
	Lead        struct {
		Self       string `json:"self"`
		AccountID  string `json:"accountId"`
		AvatarUrls struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
	} `json:"lead"`
	Components []interface{} `json:"components"`
	IssueTypes []struct {
		Self        string `json:"self"`
		ID          string `json:"id"`
		Description string `json:"description"`
		IconURL     string `json:"iconUrl"`
		Name        string `json:"name"`
		Subtask     bool   `json:"subtask"`
		AvatarID    int    `json:"avatarId,omitempty"`
	} `json:"issueTypes"`
	AssigneeType string        `json:"assigneeType"`
	Versions     []interface{} `json:"versions"`
	Name         string        `json:"name"`
	Roles        struct {
		AtlassianAddonsProjectAccess string `json:"atlassian-addons-project-access"`
		ServiceDeskTeam              string `json:"Service Desk Team"`
		ServiceDeskCustomers         string `json:"Service Desk Customers"`
		Administrators               string `json:"Administrators"`
	} `json:"roles"`
	AvatarUrls struct {
		Four8X48  string `json:"48x48"`
		Two4X24   string `json:"24x24"`
		One6X16   string `json:"16x16"`
		Three2X32 string `json:"32x32"`
	} `json:"avatarUrls"`
	ProjectKeys    []string `json:"projectKeys"`
	ProjectTypeKey string   `json:"projectTypeKey"`
	Simplified     bool     `json:"simplified"`
	Style          string   `json:"style"`
	IsPrivate      bool     `json:"isPrivate"`
	Properties     struct {
	} `json:"properties"`
}

// Returns the project details for a project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-get
func (p *ProjectService) Get(ctx context.Context, projectKeyOrID string, expands []string) (result *ProjectScheme, response *Response, err error) {

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

	var endpoint = fmt.Sprintf("rest/api/3/project/%v?%v", projectKeyOrID, params.Encode())

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
	Description string `json:"description"`
	Lead        string `json:"lead"`
	URL         string `json:"url"`
	Name        string `json:"name"`
}

// Updates the project details of a project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-put
func (p *ProjectService) Update(ctx context.Context, projectKeyOrID string, payload *ProjectUpdateScheme) (result *ProjectScheme, response *Response, err error) {

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

type ProjectStatus struct {
	Self     string `json:"self"`
	ID       string `json:"id"`
	Name     string `json:"name"`
	Subtask  bool   `json:"subtask"`
	Statuses []struct {
		Self        string `json:"self"`
		Description string `json:"description"`
		IconURL     string `json:"iconUrl"`
		Name        string `json:"name"`
		ID          string `json:"id"`
	} `json:"statuses"`
}

// Returns the valid statuses for a project.
// The statuses are grouped by issue type, as each project has a set of valid issue types and each issue type has a set of valid statuses.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-projects/#api-rest-api-3-project-projectidorkey-statuses-get
func (p *ProjectService) Statuses(ctx context.Context, projectKeyOrID string) (result *[]ProjectStatus, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/statuses", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new([]ProjectStatus)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
