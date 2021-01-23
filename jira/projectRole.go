package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProjectRoleService struct {
	client *Client
	Actor  *ProjectRoleActorService
}

type ProjectRoleDetailScheme struct {
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
}

// Returns all project roles and the details for each role.
// Note that the list of project roles is common to all projects.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-roles/#api-rest-api-3-project-projectidorkey-roledetails-get
func (p *ProjectRoleService) Details(ctx context.Context, projectKeyOrID string) (result *[]ProjectRoleDetailScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/roledetails", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new([]ProjectRoleDetailScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Gets a list of all project roles, complete with project role details and default actors.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-roles/#api-rest-api-3-role-get
func (p *ProjectRoleService) Global(ctx context.Context) (result *[]ProjectRoleDetailScheme, response *Response, err error) {

	var endpoint = "rest/api/3/role"

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new([]ProjectRoleDetailScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ProjectRoleScheme struct {
	Self        string `json:"self"`
	Name        string `json:"name"`
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// Creates a new project role with no default actors.
// You can use the Add default actors to project role operation to add default actors to the project role after creating it.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-roles/#api-rest-api-3-role-post
func (p *ProjectRoleService) Create(ctx context.Context, name, description string) (result *ProjectRoleScheme, response *Response, err error) {

	payload := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}{
		Name:        name,
		Description: description,
	}

	var endpoint = "rest/api/3/role"

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

	result = new(ProjectRoleScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Gets the project role details and the default actors associated with the role. The list of default actors is sorted by display name.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-roles/#api-rest-api-3-role-id-get
func (p *ProjectRoleService) Get(ctx context.Context, projectRoleID string) (result *ProjectRoleDetailScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/role/%v", projectRoleID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectRoleDetailScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
