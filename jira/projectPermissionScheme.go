package jira

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type ProjectPermissionSchemeService struct{ client *Client }

// Get search the permission scheme associated with the project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/permission-schemes#get-assigned-permission-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-permission-schemes/#api-rest-api-3-project-projectkeyorid-permissionscheme-get
func (p *ProjectPermissionSchemeService) Get(ctx context.Context, projectKeyOrID string, expand []string) (
	result *PermissionSchemeScheme, response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, notProjectIDError
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/3/project/%v/permissionscheme", projectKeyOrID))

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

// Assign assigns a permission scheme with a project.
// See Managing project permissions for more information about permission schemes.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/permission-schemes#assign-permission-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-permission-schemes/#api-rest-api-3-project-projectkeyorid-permissionscheme-put
func (p *ProjectPermissionSchemeService) Assign(ctx context.Context, projectKeyOrID string, permissionSchemeID int) (
	result *PermissionSchemeScheme, response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, notProjectIDError
	}

	payload := struct {
		ID int `json:"id"`
	}{
		ID: permissionSchemeID,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/permissionscheme", projectKeyOrID)

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

type ProjectIssueSecurityLevelsScheme struct {
	Levels []struct {
		Self        string `json:"self"`
		ID          string `json:"id"`
		Description string `json:"description"`
		Name        string `json:"name"`
	} `json:"levels"`
}

// SecurityLevels returns all issue security levels for the project that the user has access to.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/permission-schemes#get-project-issue-security-levels
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-permission-schemes/#api-rest-api-3-project-projectkeyorid-securitylevel-get
func (p *ProjectPermissionSchemeService) SecurityLevels(ctx context.Context, projectKeyOrID string) (
	result *ProjectIssueSecurityLevelsScheme, response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, notProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/securitylevel", projectKeyOrID)

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
