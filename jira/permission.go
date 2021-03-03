package jira

import (
	"context"
	"encoding/json"
	"net/http"
)

type PermissionService struct {
	client *Client
	Scheme *PermissionSchemeService
}

type GlobalPermissionsScheme struct {
	Permissions struct {
		AddComments struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"ADD_COMMENTS"`
		Administer struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"ADMINISTER"`
		AdministerProjects struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"ADMINISTER_PROJECTS"`
		AssignableUser struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"ASSIGNABLE_USER"`
		AssignIssues struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"ASSIGN_ISSUES"`
		BrowseProjects struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"BROWSE_PROJECTS"`
		BulkChange struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"BULK_CHANGE"`
		CloseIssues struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"CLOSE_ISSUES"`
		CreateAttachments struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"CREATE_ATTACHMENTS"`
		CreateIssues struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"CREATE_ISSUES"`
		CreateProject struct {
			Key         string `json:"key"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			Description string `json:"description"`
		} `json:"CREATE_PROJECT"`
	} `json:"permissions"`
}

// Returns a list of permissions indicating which permissions the user has.
// Details of the user's permissions can be obtained in a global, project, or issue context.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions#get-my-permissions
func (p *PermissionService) Gets(ctx context.Context) (result *GlobalPermissionsScheme, response *Response, err error) {

	var endpoint = "rest/api/3/permissions"

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(GlobalPermissionsScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
