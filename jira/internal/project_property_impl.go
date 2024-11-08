package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewProjectPropertyService creates a new instance of ProjectPropertyService.
func NewProjectPropertyService(client service.Connector, version string) (*ProjectPropertyService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ProjectPropertyService{
		internalClient: &internalProjectPropertyImpl{c: client, version: version},
	}, nil
}

// ProjectPropertyService provides methods to manage project properties in Jira Service Management.
type ProjectPropertyService struct {
	// internalClient is the connector interface for project property operations.
	internalClient jira.ProjectPropertyConnector
}

// Gets returns all project property keys for the project.
//
// GET /rest/api/{2-3}/project/{projectKeyOrID}/properties
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/properties#get-project-properties-keys
func (p *ProjectPropertyService) Gets(ctx context.Context, projectKeyOrID string) (*model.PropertyPageScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx, projectKeyOrID)
}

// Get returns the value of a project property.
//
// GET /rest/api/{2-3}/project/{projectKeyOrID}/properties/{propertyKey}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/properties#get-project-property
func (p *ProjectPropertyService) Get(ctx context.Context, projectKeyOrID, propertyKey string) (*model.EntityPropertyScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, projectKeyOrID, propertyKey)
}

// Set sets the value of the project property.
//
// You can use project properties to store custom data against the project.
//
// The value of the request body must be a valid, non-empty JSON blob.
//
// The maximum length is 32768 characters.
//
// PUT /rest/api/{2-3}/project/{projectKeyOrID}/properties/{propertyKey}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/properties#set-project-property
func (p *ProjectPropertyService) Set(ctx context.Context, projectKeyOrID, propertyKey string, payload interface{}) (*model.ResponseScheme, error) {
	return p.internalClient.Set(ctx, projectKeyOrID, propertyKey, payload)
}

// Delete deletes the property from a project.
//
// DELETE /rest/api/{2-3}/project/{projectKeyOrID}/properties/{propertyKey}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/properties#delete-project-property
func (p *ProjectPropertyService) Delete(ctx context.Context, projectKeyOrID, propertyKey string) (*model.ResponseScheme, error) {
	return p.internalClient.Delete(ctx, projectKeyOrID, propertyKey)
}

type internalProjectPropertyImpl struct {
	c       service.Connector
	version string
}

func (i *internalProjectPropertyImpl) Gets(ctx context.Context, projectKeyOrID string) (*model.PropertyPageScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/properties", i.version, projectKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	properties := new(model.PropertyPageScheme)
	response, err := i.c.Call(request, properties)
	if err != nil {
		return nil, response, err
	}

	return properties, response, nil
}

func (i *internalProjectPropertyImpl) Get(ctx context.Context, projectKeyOrID, propertyKey string) (*model.EntityPropertyScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	if propertyKey == "" {
		return nil, nil, model.ErrNoPropertyKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/properties/%v", i.version, projectKeyOrID, propertyKey)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	property := new(model.EntityPropertyScheme)
	response, err := i.c.Call(request, property)
	if err != nil {
		return nil, response, err
	}

	return property, response, nil
}

func (i *internalProjectPropertyImpl) Set(ctx context.Context, projectKeyOrID, propertyKey string, payload interface{}) (*model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, model.ErrNoProjectIDOrKey
	}

	if propertyKey == "" {
		return nil, model.ErrNoPropertyKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/properties/%v", i.version, projectKeyOrID, propertyKey)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalProjectPropertyImpl) Delete(ctx context.Context, projectKeyOrID, propertyKey string) (*model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, model.ErrNoProjectIDOrKey
	}

	if propertyKey == "" {
		return nil, model.ErrNoPropertyKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/properties/%v", i.version, projectKeyOrID, propertyKey)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
