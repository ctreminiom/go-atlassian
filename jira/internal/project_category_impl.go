package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

func NewProjectCategoryService(client service.Client, version string) (*ProjectCategoryService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ProjectCategoryService{
		internalClient: &internalProjectCategoryImpl{c: client, version: version},
	}, nil
}

type ProjectCategoryService struct {
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
// GET /rest/api/{2-3}/projectCategory/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/categories#get-project-category-by-id
func (p *ProjectCategoryService) Get(ctx context.Context, categoryId int) (*model.ProjectCategoryScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, categoryId)
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
// PUT /rest/api/{2-3}/projectCategory/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/categories#update-project-category
func (p *ProjectCategoryService) Update(ctx context.Context, categoryId int, payload *model.ProjectCategoryPayloadScheme) (*model.ProjectCategoryScheme, *model.ResponseScheme, error) {
	return p.internalClient.Update(ctx, categoryId, payload)
}

// Delete deletes a project category.
//
// DELETE /rest/api/{2-3}/projectCategory/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/categories#delete-project-category
func (p *ProjectCategoryService) Delete(ctx context.Context, categoryId int) (*model.ResponseScheme, error) {
	return p.internalClient.Delete(ctx, categoryId)
}

type internalProjectCategoryImpl struct {
	c       service.Client
	version string
}

func (i *internalProjectCategoryImpl) Gets(ctx context.Context) ([]*model.ProjectCategoryScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/projectCategory", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

func (i *internalProjectCategoryImpl) Get(ctx context.Context, categoryId int) (*model.ProjectCategoryScheme, *model.ResponseScheme, error) {

	if categoryId == 0 {
		return nil, nil, model.ErrNoProjectCategoryIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/projectCategory/%v", i.version, categoryId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/projectCategory", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
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

func (i *internalProjectCategoryImpl) Update(ctx context.Context, categoryId int, payload *model.ProjectCategoryPayloadScheme) (*model.ProjectCategoryScheme, *model.ResponseScheme, error) {

	if categoryId == 0 {
		return nil, nil, model.ErrNoProjectCategoryIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/projectCategory/%v", i.version, categoryId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
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

func (i *internalProjectCategoryImpl) Delete(ctx context.Context, categoryId int) (*model.ResponseScheme, error) {

	if categoryId == 0 {
		return nil, model.ErrNoProjectCategoryIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/projectCategory/%v", i.version, categoryId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
