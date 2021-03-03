package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ProjectRoleService struct {
	client *Client
	Actor  *ProjectRoleActorService
}

// Returns a list of project roles for the project returning the name and self URL for each role.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-project-roles-for-project
func (p *ProjectRoleService) Gets(ctx context.Context, projectKeyOrID string) (result *map[string]int, response *Response, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid projectKeyOrID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/role", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	var (
		roles          = make(map[string]int)
		resultAsObject map[string]interface{}
	)

	if err = json.Unmarshal(response.BodyAsBytes, &resultAsObject); err != nil {
		return
	}

	for roleName, roleURL := range resultAsObject {

		urlParsed, err := url.Parse(roleURL.(string))
		if err != nil {
			return nil, response, err
		}

		urlPart := strings.Split(urlParsed.Path, "/")

		projectRoleIDAsInt, err := strconv.Atoi(urlPart[len(urlPart)-1])
		if err != nil {
			return nil, response, err
		}

		roles[roleName] = projectRoleIDAsInt
	}

	result = &roles

	return
}

type ProjectRoleScheme struct {
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

// Returns a project role's details and actors associated with the project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-project-role-for-project
func (p *ProjectRoleService) Get(ctx context.Context, projectKeyOrID string, roleID int) (result *ProjectRoleScheme, response *Response, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid projectKeyOrID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/role/%v", projectKeyOrID, roleID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

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

type ProjectRoleDetailScheme struct {
	Self             string `json:"self"`
	Name             string `json:"name"`
	ID               int    `json:"id"`
	Description      string `json:"description"`
	Admin            bool   `json:"admin"`
	Default          bool   `json:"default"`
	RoleConfigurable bool   `json:"roleConfigurable"`
	TranslatedName   string `json:"translatedName"`
}

// Returns all project roles and the details for each role.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-project-role-details
func (p *ProjectRoleService) Details(ctx context.Context, projectKeyOrID string) (result *[]ProjectRoleDetailScheme, response *Response, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid projectKeyOrID value")
	}

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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-all-project-roles
func (p *ProjectRoleService) Global(ctx context.Context) (result *[]ProjectRoleScheme, response *Response, err error) {

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

	result = new([]ProjectRoleScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Creates a new project role with no default actors.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/roles#create-project-role
func (p *ProjectRoleService) Create(ctx context.Context, name, description string) (result *ProjectRoleScheme, response *Response, err error) {

	if len(name) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid projectKeyOrID value")
	}

	payload := struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
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
