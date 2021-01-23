package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProjectPermissionSchemeService struct{ client *Client }

// Returns the issue security scheme associated with the project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-permission-schemes/#api-rest-api-3-project-projectkeyorid-issuesecuritylevelscheme-get
func (p *ProjectPermissionSchemeService) SecurityScheme(ctx context.Context, projectKeyOrID string) (result *IssueSecurityScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/issuesecuritylevelscheme", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueSecurityScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ProjectPermissionSchemeScheme struct {
	ID          int    `json:"id"`
	Self        string `json:"self"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Gets the permission scheme associated with the project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-permission-schemes/#api-rest-api-3-project-projectkeyorid-permissionscheme-get
func (p *ProjectPermissionSchemeService) Get(ctx context.Context, projectKeyOrID string) (result *ProjectPermissionSchemeScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/permissionscheme", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectPermissionSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Assigns a permission scheme with a project.
// See Managing project permissions for more information about permission schemes.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-permission-schemes/#api-rest-api-3-project-projectkeyorid-permissionscheme-put
func (p *ProjectPermissionSchemeService) Assign(ctx context.Context, projectKeyOrID string, permissionSchemeID int) (result *ProjectPermissionSchemeScheme, response *Response, err error) {

	payload := struct {
		ID int `json:"id"`
	}{ID: permissionSchemeID}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/permissionscheme", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectPermissionSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
