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
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-components/#api-rest-api-3-component-post
func (p *ProjectComponentService) Create(ctx context.Context, payload *ProjectComponentPayloadScheme) (result *ProjectComponentScheme, response *Response, err error) {

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
	Self                string     `json:"self"`
	ID                  string     `json:"id"`
	Name                string     `json:"name"`
	Description         string     `json:"description"`
	Lead                UserScheme `json:"lead"`
	AssigneeType        string     `json:"assigneeType"`
	Assignee            UserScheme `json:"assignee"`
	RealAssigneeType    string     `json:"realAssigneeType"`
	RealAssignee        UserScheme `json:"realAssignee"`
	IsAssigneeTypeValid bool       `json:"isAssigneeTypeValid"`
	Project             string     `json:"project"`
	ProjectID           int        `json:"projectId"`
}

// Returns all components in a project.
// See the Get project components paginated resource if you want to get a full list of components with pagination.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-components/#api-rest-api-3-project-projectidorkey-components-get
func (p *ProjectComponentService) Gets(ctx context.Context, projectKeyOrID string) (result *[]ProjectComponentScheme, response *Response, err error) {

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
	Self       string `json:"self"`
	IssueCount int    `json:"issueCount"`
}

// Returns the counts of issues assigned to the component.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-components/#api-rest-api-3-component-id-relatedissuecounts-get
func (p *ProjectComponentService) Count(ctx context.Context, componentID string) (result *ProjectComponentCountScheme, response *Response, err error) {

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
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-components/#api-rest-api-3-component-id-delete
func (p *ProjectComponentService) Delete(ctx context.Context, componentID string) (response *Response, err error) {

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
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-components/#api-rest-api-3-component-id-put
func (p *ProjectComponentService) Update(ctx context.Context, componentID string, payload *ProjectComponentPayloadScheme) (result *ProjectComponentScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/component/%v", componentID)

	request, err := p.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
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

// Returns a component.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-components/#api-rest-api-3-component-id-get
func (p *ProjectComponentService) Get(ctx context.Context, componentID string) (result *ProjectComponentScheme, response *Response, err error) {

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
