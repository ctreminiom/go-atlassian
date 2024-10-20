package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
	"net/http"
)

// NewProjectTypeService creates a new instance of ProjectTypeService.
func NewProjectTypeService(client service.Connector, version string) (*ProjectTypeService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ProjectTypeService{
		internalClient: &internalProjectTypeImpl{c: client, version: version},
	}, nil
}

// ProjectTypeService provides methods to manage project types in Jira Service Management.
type ProjectTypeService struct {
	// internalClient is the connector interface for project type operations.
	internalClient jira.ProjectTypeConnector
}

// Gets returns all project types, whether the instance has a valid license for each type.
//
// GET /rest/api/{2-3}/project/type
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/types#get-all-project-types
func (p *ProjectTypeService) Gets(ctx context.Context) ([]*model.ProjectTypeScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx)
}

// Licensed returns all project types with a valid license.
//
// GET /rest/api/{2-3}/project/type/accessible
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/types#get-licensed-project-types
func (p *ProjectTypeService) Licensed(ctx context.Context) ([]*model.ProjectTypeScheme, *model.ResponseScheme, error) {
	return p.internalClient.Licensed(ctx)
}

// Get returns a project type
//
// GET /rest/api/{2-3}/project/type/{projectTypeKey}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/types#get-project-type-by-key
func (p *ProjectTypeService) Get(ctx context.Context, projectTypeKey string) (*model.ProjectTypeScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, projectTypeKey)
}

// Accessible returns a project type if it is accessible to the user.
//
// GET /rest/api/{2-3}/project/type/{projectTypeKey}/accessible
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/types#get-accessible-project-type-by-key
func (p *ProjectTypeService) Accessible(ctx context.Context, projectTypeKey string) (*model.ProjectTypeScheme, *model.ResponseScheme, error) {
	return p.internalClient.Accessible(ctx, projectTypeKey)
}

type internalProjectTypeImpl struct {
	c       service.Connector
	version string
}

func (i *internalProjectTypeImpl) Gets(ctx context.Context) ([]*model.ProjectTypeScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/project/type", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var _types []*model.ProjectTypeScheme
	response, err := i.c.Call(request, &_types)
	if err != nil {
		return nil, response, err
	}

	return _types, response, nil
}

func (i *internalProjectTypeImpl) Licensed(ctx context.Context) ([]*model.ProjectTypeScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/project/type/accessible", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var _types []*model.ProjectTypeScheme
	response, err := i.c.Call(request, &_types)
	if err != nil {
		return nil, response, err
	}

	return _types, response, nil
}

func (i *internalProjectTypeImpl) Get(ctx context.Context, projectTypeKey string) (*model.ProjectTypeScheme, *model.ResponseScheme, error) {

	if projectTypeKey == "" {
		return nil, nil, model.ErrProjectTypeKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/type/%v", i.version, projectTypeKey)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	_type := new(model.ProjectTypeScheme)
	response, err := i.c.Call(request, _type)
	if err != nil {
		return nil, response, err
	}

	return _type, response, nil
}

func (i *internalProjectTypeImpl) Accessible(ctx context.Context, projectTypeKey string) (*model.ProjectTypeScheme, *model.ResponseScheme, error) {

	if projectTypeKey == "" {
		return nil, nil, model.ErrProjectTypeKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/type/%v/accessible", i.version, projectTypeKey)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	_type := new(model.ProjectTypeScheme)
	response, err := i.c.Call(request, _type)
	if err != nil {
		return nil, response, err
	}

	return _type, response, nil
}
