package v2

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"net/http"
)

type ProjectCategoryService struct{ client *Client }

// Gets returns all project categories.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/categories#get-all-project-categories
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-project-categories/#api-rest-api-2-projectcategory-get
func (p *ProjectCategoryService) Gets(ctx context.Context) (result []*models.ProjectCategoryScheme, response *ResponseScheme,
	err error) {

	var endpoint = "rest/api/2/projectCategory"

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

// Get returns a project category.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/categories#get-project-category-by-id
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-project-categories/#api-rest-api-2-projectcategory-id-get
func (p *ProjectCategoryService) Get(ctx context.Context, projectCategoryID int) (result *models.ProjectCategoryScheme,
	response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/api/2/projectCategory/%v", projectCategoryID)

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

// Create creates a project category.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/categories#create-project-category
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-project-categories/#api-rest-api-2-projectcategory-post
func (p *ProjectCategoryService) Create(ctx context.Context, payload *models.ProjectCategoryPayloadScheme) (result *models.ProjectCategoryScheme,
	response *ResponseScheme, err error) {

	var endpoint = "rest/api/2/projectCategory"

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

// Update updates a project category.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/categories#update-project-category
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-project-categories/#api-rest-api-2-projectcategory-id-put
func (p *ProjectCategoryService) Update(ctx context.Context, projectCategoryID int, payload *models.ProjectCategoryPayloadScheme) (
	result *models.ProjectCategoryScheme, response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/api/2/projectCategory/%v", projectCategoryID)

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

// Delete deletes a project category.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/categories#delete-project-category
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-project-categories/#api-rest-api-2-projectcategory-id-delete
func (p *ProjectCategoryService) Delete(ctx context.Context, projectCategoryID int) (response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/api/2/projectCategory/%v", projectCategoryID)

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
