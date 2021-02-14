package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ProjectPermissionSchemeService struct{ client *Client }

type ProjectPermissionSchemeScheme struct {
	Expand      string `json:"expand"`
	ID          int    `json:"id"`
	Self        string `json:"self"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Permissions []struct {
		ID     int    `json:"id"`
		Self   string `json:"self"`
		Holder struct {
			Type        string `json:"type"`
			Parameter   string `json:"parameter"`
			ProjectRole struct {
				Self        string `json:"self"`
				Name        string `json:"name"`
				ID          int    `json:"id"`
				Description string `json:"description"`
			} `json:"projectRole"`
			Expand string `json:"expand"`
		} `json:"holder"`
		Permission string `json:"permission"`
	} `json:"permissions"`
}

// Gets the permission scheme associated with the project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-permission-schemes/#api-rest-api-3-project-projectkeyorid-permissionscheme-get
func (p *ProjectPermissionSchemeService) Get(ctx context.Context, projectKeyOrID string, expands []string) (result *ProjectPermissionSchemeScheme, response *Response, err error) {

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
	request.Header.Set("Content-Type", "application/json")

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

type ProjectIssueSecurityLevelsScheme struct {
	Levels []struct {
		Self        string `json:"self"`
		ID          string `json:"id"`
		Description string `json:"description"`
		Name        string `json:"name"`
	} `json:"levels"`
}

// Returns all issue security levels for the project that the user has access to.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-permission-schemes/#api-rest-api-3-project-projectkeyorid-securitylevel-get
func (p *ProjectPermissionSchemeService) SecurityLevels(ctx context.Context, projectKeyOrID string) (result *ProjectIssueSecurityLevelsScheme, response *Response, err error) {

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
