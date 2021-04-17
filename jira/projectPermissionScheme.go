package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ProjectPermissionSchemeService struct{ client *Client }

// Search the permission scheme associated with the project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/permission-schemes#get-assigned-permission-scheme
func (p *ProjectPermissionSchemeService) Get(ctx context.Context, projectKeyOrID string, expands []string) (result *PermissionSchemeScheme, response *Response, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid projectKeyOrID value")
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
		endpoint = fmt.Sprintf("rest/api/3/project/%v/permissionscheme?%v", projectKeyOrID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/api/3/project/%v/permissionscheme", projectKeyOrID)
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

	result = new(PermissionSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Assigns a permission scheme with a project.
// See Managing project permissions for more information about permission schemes.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/permission-schemes#assign-permission-scheme
func (p *ProjectPermissionSchemeService) Assign(ctx context.Context, projectKeyOrID string, permissionSchemeID int) (result *PermissionSchemeScheme, response *Response, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid projectKeyOrID value")
	}

	payload := struct {
		ID int `json:"id"`
	}{ID: permissionSchemeID}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/permissionscheme", projectKeyOrID)

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

	result = new(PermissionSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
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

// Returns all issue security levels for the project that the user has access to.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/permission-schemes#get-project-issue-security-levels
func (p *ProjectPermissionSchemeService) SecurityLevels(ctx context.Context, projectKeyOrID string) (result *ProjectIssueSecurityLevelsScheme, response *Response, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid projectKeyOrID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/securitylevel", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectIssueSecurityLevelsScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
