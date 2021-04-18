package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProjectComponentService struct{ client *Client }

type ProjectComponentPayloadScheme struct {
	IsAssigneeTypeValid bool   `json:"isAssigneeTypeValid,omitempty"`
	Name                string `json:"name,omitempty"`
	Description         string `json:"description,omitempty"`
	Project             string `json:"project,omitempty"`
	AssigneeType        string `json:"assigneeType,omitempty"`
	LeadAccountID       string `json:"leadAccountId,omitempty"`
}

// Creates a component.
// Use components to provide containers for issues within a project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/components#create-component
func (p *ProjectComponentService) Create(ctx context.Context, payload *ProjectComponentPayloadScheme) (result *ProjectComponentScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, please provide a valid ProjectComponentPayloadScheme pointer")
	}

	var endpoint = "rest/api/3/component"

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

	result = new(ProjectComponentScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ProjectComponentScheme struct {
	Self                string      `json:"self,omitempty"`
	ID                  string      `json:"id,omitempty"`
	Name                string      `json:"name,omitempty"`
	Description         string      `json:"description,omitempty"`
	Lead                *UserScheme `json:"lead,omitempty"`
	LeadUserName        string      `json:"leadUserName,omitempty"`
	AssigneeType        string      `json:"assigneeType,omitempty"`
	Assignee            *UserScheme `json:"assignee,omitempty"`
	RealAssigneeType    string      `json:"realAssigneeType,omitempty"`
	RealAssignee        *UserScheme `json:"realAssignee,omitempty"`
	IsAssigneeTypeValid bool        `json:"isAssigneeTypeValid,omitempty"`
	Project             string      `json:"project,omitempty"`
	ProjectID           int         `json:"projectId,omitempty"`
}

// Returns all components in a project.
// See the Get project components paginated resource if you want to get a full list of components with pagination.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/components#get-project-components-paginated
func (p *ProjectComponentService) Gets(ctx context.Context, projectKeyOrID string) (result *[]ProjectComponentScheme, response *Response, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid projectKeyOrID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/components", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new([]ProjectComponentScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ProjectComponentCountScheme struct {
	Self       string `json:"self,omitempty"`
	IssueCount int    `json:"issueCount,omitempty"`
}

// Returns the counts of issues assigned to the component.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/components#get-component-issues-count
func (p *ProjectComponentService) Count(ctx context.Context, componentID string) (result *ProjectComponentCountScheme, response *Response, err error) {

	if len(componentID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valida componentID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/component/%v/relatedIssueCounts", componentID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectComponentCountScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes a component.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/components#delete-component
func (p *ProjectComponentService) Delete(ctx context.Context, componentID string) (response *Response, err error) {

	if len(componentID) == 0 {
		return nil, fmt.Errorf("error, please provide a valida componentID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/component/%v", componentID)

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

// Updates a component.
// Any fields included in the request are overwritten.
// If leadAccountId is an empty string ("") the component lead is removed.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/components#update-component
func (p *ProjectComponentService) Update(ctx context.Context, componentID string, payload *ProjectComponentPayloadScheme) (result *ProjectComponentScheme, response *Response, err error) {

	if len(componentID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid componentID value")
	}

	if payload == nil {
		return nil, nil, fmt.Errorf("error, please provide a valid ProjectComponentPayloadScheme pointer")
	}

	var endpoint = fmt.Sprintf("rest/api/3/component/%v", componentID)

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

	result = new(ProjectComponentScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns a component.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/components#get-component
func (p *ProjectComponentService) Get(ctx context.Context, componentID string) (result *ProjectComponentScheme, response *Response, err error) {

	if len(componentID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid componentID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/component/%v", componentID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectComponentScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
