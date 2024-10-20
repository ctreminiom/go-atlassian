package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewProjectComponentService creates a new instance of ProjectComponentService.
func NewProjectComponentService(client service.Connector, version string) (*ProjectComponentService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ProjectComponentService{
		internalClient: &internalProjectComponentImpl{c: client, version: version},
	}, nil
}

// ProjectComponentService provides methods to manage project components in Jira Service Management.
type ProjectComponentService struct {
	// internalClient is the connector interface for project component operations.
	internalClient jira.ProjectComponentConnector
}

// Create creates a component. Use components to provide containers for issues within a project.
//
// POST /rest/api/{2-3}/component
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/components#create-component
func (p *ProjectComponentService) Create(ctx context.Context, payload *model.ComponentPayloadScheme) (*model.ComponentScheme, *model.ResponseScheme, error) {
	return p.internalClient.Create(ctx, payload)
}

// Gets returns all components in a project.
//
// GET /rest/api/{2-3}/project/{projectKeyOrID}/components
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/components#get-project-components
func (p *ProjectComponentService) Gets(ctx context.Context, projectKeyOrID string) ([]*model.ComponentScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx, projectKeyOrID)
}

// Count returns the counts of issues assigned to the component.
//
// GET /rest/api/{2-3}/component/{componentID}/relatedIssueCounts
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/components#get-component-issues-count
func (p *ProjectComponentService) Count(ctx context.Context, componentID string) (*model.ComponentCountScheme, *model.ResponseScheme, error) {
	return p.internalClient.Count(ctx, componentID)
}

// Delete deletes a component.
//
// DELETE /rest/api/{2-3}/component/{componentID}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/components#delete-component
func (p *ProjectComponentService) Delete(ctx context.Context, componentID string) (*model.ResponseScheme, error) {
	return p.internalClient.Delete(ctx, componentID)
}

// Update updates a component.
//
// # Any fields included in the request are overwritten
//
// PUT /rest/api/{2-3}/component/{componentID}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/components#update-component
func (p *ProjectComponentService) Update(ctx context.Context, componentID string, payload *model.ComponentPayloadScheme) (*model.ComponentScheme, *model.ResponseScheme, error) {
	return p.internalClient.Update(ctx, componentID, payload)
}

// Get returns a component.
//
// GET /rest/api/{2-3}/component/{componentID}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/components#get-component
func (p *ProjectComponentService) Get(ctx context.Context, componentID string) (*model.ComponentScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, componentID)
}

type internalProjectComponentImpl struct {
	c       service.Connector
	version string
}

func (i *internalProjectComponentImpl) Create(ctx context.Context, payload *model.ComponentPayloadScheme) (*model.ComponentScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/component", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	component := new(model.ComponentScheme)
	response, err := i.c.Call(request, component)
	if err != nil {
		return nil, response, err
	}

	return component, response, nil
}

func (i *internalProjectComponentImpl) Gets(ctx context.Context, projectKeyOrID string) ([]*model.ComponentScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/components", i.version, projectKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var components []*model.ComponentScheme
	response, err := i.c.Call(request, components)
	if err != nil {
		return nil, response, err
	}

	return components, response, nil
}

func (i *internalProjectComponentImpl) Count(ctx context.Context, componentID string) (*model.ComponentCountScheme, *model.ResponseScheme, error) {

	if componentID == "" {
		return nil, nil, model.ErrNoComponentID
	}

	endpoint := fmt.Sprintf("rest/api/%v/component/%v/relatedIssueCounts", i.version, componentID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	count := new(model.ComponentCountScheme)
	response, err := i.c.Call(request, count)
	if err != nil {
		return nil, response, err
	}

	return count, response, nil
}

func (i *internalProjectComponentImpl) Delete(ctx context.Context, componentID string) (*model.ResponseScheme, error) {

	if componentID == "" {
		return nil, model.ErrNoComponentID
	}

	endpoint := fmt.Sprintf("rest/api/%v/component/%v", i.version, componentID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalProjectComponentImpl) Update(ctx context.Context, componentID string, payload *model.ComponentPayloadScheme) (*model.ComponentScheme, *model.ResponseScheme, error) {

	if componentID == "" {
		return nil, nil, model.ErrNoComponentID
	}

	endpoint := fmt.Sprintf("rest/api/%v/component/%v", i.version, componentID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	component := new(model.ComponentScheme)
	response, err := i.c.Call(request, component)
	if err != nil {
		return nil, response, err
	}

	return component, response, nil
}

func (i *internalProjectComponentImpl) Get(ctx context.Context, componentID string) (*model.ComponentScheme, *model.ResponseScheme, error) {

	if componentID == "" {
		return nil, nil, model.ErrNoComponentID
	}

	endpoint := fmt.Sprintf("rest/api/%v/component/%v", i.version, componentID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	component := new(model.ComponentScheme)
	response, err := i.c.Call(request, component)
	if err != nil {
		return nil, response, err
	}

	return component, response, nil
}
