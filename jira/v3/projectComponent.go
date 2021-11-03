package v3

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"net/http"
)

type ProjectComponentService struct{ client *Client }

// Create creates a component.
// Use components to provide containers for issues within a project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/components#create-component
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-components/#api-rest-api-3-component-post
func (p *ProjectComponentService) Create(ctx context.Context, payload *models.ComponentPayloadScheme) (
	result *models.ComponentScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/component"

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
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

// Gets returns all components in a project.
// See the Get project components paginated resource if you want to get a full list of components with pagination.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/components#get-project-components-paginated
// Atlassian Docs; https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-components/#api-rest-api-3-project-projectidorkey-components-get
func (p *ProjectComponentService) Gets(ctx context.Context, projectKeyOrID string) (result []*models.ComponentScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, models.ErrNoProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/components", projectKeyOrID)

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

// Count returns the counts of issues assigned to the component.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/components#get-component-issues-count
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-components/#api-rest-api-3-component-id-relatedissuecounts-get
func (p *ProjectComponentService) Count(ctx context.Context, componentID string) (result *models.ComponentCountScheme,
	response *ResponseScheme, err error) {

	if len(componentID) == 0 {
		return nil, nil, models.ErrNoComponentIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/component/%v/relatedIssueCounts", componentID)

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

// Delete deletes a component.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/components#delete-component
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-components/#api-rest-api-3-component-id-delete
func (p *ProjectComponentService) Delete(ctx context.Context, componentID string) (response *ResponseScheme, err error) {

	if len(componentID) == 0 {
		return nil, models.ErrNoComponentIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/component/%v", componentID)

	request, err := p.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Update updates a component.
// Any fields included in the request are overwritten.
// If leadAccountId is an empty string ("") the component lead is removed.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/components#update-component
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-components/#api-rest-api-3-component-id-put
func (p *ProjectComponentService) Update(ctx context.Context, componentID string, payload *models.ComponentPayloadScheme) (
	result *models.ComponentScheme, response *ResponseScheme, err error) {

	if len(componentID) == 0 {
		return nil, nil, models.ErrNoComponentIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/component/%v", componentID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

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

// Get returns a component.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/components#get-component
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-components/#api-rest-api-3-component-id-get
func (p *ProjectComponentService) Get(ctx context.Context, componentID string) (result *models.ComponentScheme,
	response *ResponseScheme, err error) {

	if len(componentID) == 0 {
		return nil, nil, models.ErrNoComponentIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/component/%v", componentID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
