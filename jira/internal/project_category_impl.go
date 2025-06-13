package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewProjectCategoryService creates a new instance of ProjectCategoryService.
func NewProjectCategoryService(client service.Connector, version string) (*ProjectCategoryService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ProjectCategoryService{
		internalClient: &internalProjectCategoryImpl{c: client, version: version},
	}, nil
}

// ProjectCategoryService provides methods to manage project categories in Jira Service Management.
type ProjectCategoryService struct {
	// internalClient is the connector interface for project category operations.
	internalClient jira.ProjectCategoryConnector
}

// Gets returns all project categories.
//
// GET /rest/api/{2-3}/projectCategory
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/categories#get-all-project-categories
func (p *ProjectCategoryService) Gets(ctx context.Context) ([]*model.ProjectCategoryScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx)
}

// Get returns a project category.
//
// GET /rest/api/{2-3}/projectCategory/{categoryID}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/categories#get-project-category-by-id
func (p *ProjectCategoryService) Get(ctx context.Context, categoryID int) (*model.ProjectCategoryScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, categoryID)
}

// Create creates a project category.
//
// POST /rest/api/{2-3}/projectCategory
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/categories#create-project-category
func (p *ProjectCategoryService) Create(ctx context.Context, payload *model.ProjectCategoryPayloadScheme) (*model.ProjectCategoryScheme, *model.ResponseScheme, error) {
	return p.internalClient.Create(ctx, payload)
}

// Update updates a project category.
//
// PUT /rest/api/{2-3}/projectCategory/{categoryID}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/categories#update-project-category
func (p *ProjectCategoryService) Update(ctx context.Context, categoryID int, payload *model.ProjectCategoryPayloadScheme) (*model.ProjectCategoryScheme, *model.ResponseScheme, error) {
	return p.internalClient.Update(ctx, categoryID, payload)
}

// Delete deletes a project category.
//
// DELETE /rest/api/{2-3}/projectCategory/{categoryID}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/categories#delete-project-category
func (p *ProjectCategoryService) Delete(ctx context.Context, categoryID int) (*model.ResponseScheme, error) {
	return p.internalClient.Delete(ctx, categoryID)
}

type internalProjectCategoryImpl struct {
	c       service.Connector
	version string
}

func (i *internalProjectCategoryImpl) Gets(ctx context.Context) ([]*model.ProjectCategoryScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/projectCategory", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var categories []*model.ProjectCategoryScheme
	response, err := i.c.Call(request, &categories)
	if err != nil {
		return nil, response, err
	}

	return categories, response, nil
}

func (i *internalProjectCategoryImpl) Get(ctx context.Context, categoryID int) (*model.ProjectCategoryScheme, *model.ResponseScheme, error) {

	if categoryID == 0 {
		return nil, nil, model.ErrNoProjectCategoryID
	}

	endpoint := fmt.Sprintf("rest/api/%v/projectCategory/%v", i.version, categoryID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	category := new(model.ProjectCategoryScheme)
	response, err := i.c.Call(request, category)
	if err != nil {
		return nil, response, err
	}

	return category, response, nil
}

func (i *internalProjectCategoryImpl) Create(ctx context.Context, payload *model.ProjectCategoryPayloadScheme) (*model.ProjectCategoryScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/projectCategory", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	category := new(model.ProjectCategoryScheme)
	response, err := i.c.Call(request, category)
	if err != nil {
		return nil, response, err
	}

	return category, response, nil
}

func (i *internalProjectCategoryImpl) Update(ctx context.Context, categoryID int, payload *model.ProjectCategoryPayloadScheme) (*model.ProjectCategoryScheme, *model.ResponseScheme, error) {

	if categoryID == 0 {
		return nil, nil, model.ErrNoProjectCategoryID
	}

	endpoint := fmt.Sprintf("rest/api/%v/projectCategory/%v", i.version, categoryID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	category := new(model.ProjectCategoryScheme)
	response, err := i.c.Call(request, category)
	if err != nil {
		return nil, response, err
	}

	return category, response, nil
}

func (i *internalProjectCategoryImpl) Delete(ctx context.Context, categoryID int) (*model.ResponseScheme, error) {

	if categoryID == 0 {
		return nil, model.ErrNoProjectCategoryID
	}

	endpoint := fmt.Sprintf("rest/api/%v/projectCategory/%v", i.version, categoryID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
