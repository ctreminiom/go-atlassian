package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProjectComponentService struct{ client *Client }

type ProjectComponentScheme struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Lead        struct {
		Self       string `json:"self"`
		Key        string `json:"key"`
		AccountID  string `json:"accountId"`
		Name       string `json:"name"`
		AvatarUrls struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
	} `json:"lead"`
	AssigneeType string `json:"assigneeType"`
	Assignee     struct {
		Self       string `json:"self"`
		Key        string `json:"key"`
		AccountID  string `json:"accountId"`
		Name       string `json:"name"`
		AvatarUrls struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
	} `json:"assignee"`
	RealAssigneeType string `json:"realAssigneeType"`
	RealAssignee     struct {
		Self       string `json:"self"`
		Key        string `json:"key"`
		AccountID  string `json:"accountId"`
		Name       string `json:"name"`
		AvatarUrls struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
	} `json:"realAssignee"`
	IsAssigneeTypeValid bool   `json:"isAssigneeTypeValid"`
	Project             string `json:"project"`
	ProjectID           int    `json:"projectId"`
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

type ProjectComponentUpdatePayloadScheme struct {
	IsAssigneeTypeValid bool   `json:"isAssigneeTypeValid"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	AssigneeType        string `json:"assigneeType"`
	LeadAccountID       string `json:"leadAccountId"`
}

// Updates a component.
// Any fields included in the request are overwritten.
// If leadAccountId is an empty string ("") the component lead is removed.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-components/#api-rest-api-3-component-id-put
func (p *ProjectComponentService) Update(ctx context.Context, componentID string, payload *ProjectComponentUpdatePayloadScheme) (result *ProjectComponentScheme, response *Response, err error) {

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
