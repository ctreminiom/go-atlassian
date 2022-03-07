package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
)

type ProjectPropertyService struct{ client *Client }

// Gets returns all project property keys for the project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/properties#get-project-properties-keys
func (p *ProjectPropertyService) Gets(ctx context.Context, projectKeyOrID string) (result *models.ProjectPropertyPageScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, models.ErrNoProjectIDError
	}

	endpoint := fmt.Sprintf("rest/api/3/project/%v/properties", projectKeyOrID)

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

// Get returns the value of a project property.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/properties#get-project-property
func (p *ProjectPropertyService) Get(ctx context.Context, projectKeyOrID, propertyKey string) (result *models.EntityPropertyScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, models.ErrNoProjectIDError
	}

	if len(propertyKey) == 0 {
		return nil, nil, models.ErrNoPropertyKeyError
	}

	endpoint := fmt.Sprintf("rest/api/3/project/%v/properties/%v", projectKeyOrID, propertyKey)

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

// Set sets the value of the project property. You can use project properties to store custom data against the project.
// The value of the request body must be a valid, non-empty JSON blob. The maximum length is 32768 characters.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/properties#set-project-property
func (p *ProjectPropertyService) Set(ctx context.Context, projectKeyOrID, propertyKey string, payload interface{}) (response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, models.ErrNoProjectIDError
	}

	if len(propertyKey) == 0 {
		return nil, models.ErrNoPropertyKeyError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/3/project/%v/properties/%v", projectKeyOrID, propertyKey)

	request, err := p.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Delete deletes the property from a project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/properties#delete-project-property
func (p *ProjectPropertyService) Delete(ctx context.Context, projectKeyOrID, propertyKey string) (response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, models.ErrNoProjectIDError
	}

	if len(propertyKey) == 0 {
		return nil, models.ErrNoPropertyKeyError
	}

	endpoint := fmt.Sprintf("rest/api/3/project/%v/properties/%v", projectKeyOrID, propertyKey)

	request, err := p.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.call(request, nil)
	if err != nil {
		return
	}

	return
}
