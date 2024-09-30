package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/bitbucket"
	"net/http"
)

// ProjectService handles communication with the project related methods of the Bitbucket API.
type ProjectService struct {
	internalClient bitbucket.ProjectConnector
}

// NewProjectService handles communication with the project related methods of the Bitbucket API.
func NewProjectService(client service.Connector) *ProjectService {

	return &ProjectService{
		internalClient: &internalProjectServiceImpl{c: client},
	}
}

type internalProjectServiceImpl struct {
	c service.Connector
}

func (i *internalProjectServiceImpl) Create(ctx context.Context, workspace string, payload *model.BitbucketProjectScheme) (*model.BitbucketProjectScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/projects", workspace)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	projectCreated := new(model.BitbucketProjectScheme)
	response, err := i.c.Call(request, projectCreated)
	if err != nil {
		return nil, response, err
	}

	return projectCreated, response, nil
}

func (i *internalProjectServiceImpl) Get(ctx context.Context, workspace, projectKey string) (*model.BitbucketProjectScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, nil, model.ErrNoProjectSlug
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/projects/%v", workspace, projectKey)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	project := new(model.BitbucketProjectScheme)
	response, err := i.c.Call(request, project)
	if err != nil {
		return nil, response, err
	}

	return project, response, nil
}

func (i *internalProjectServiceImpl) Update(ctx context.Context, workspace, projectKey string, payload *model.BitbucketProjectScheme) (*model.BitbucketProjectScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, nil, model.ErrNoProjectSlug
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/projects/%v", workspace, projectKey)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	projectUpdated := new(model.BitbucketProjectScheme)
	response, err := i.c.Call(request, projectUpdated)
	if err != nil {
		return nil, response, err
	}

	return projectUpdated, response, nil
}

func (i *internalProjectServiceImpl) Delete(ctx context.Context, workspace, projectKey string) (*model.ResponseScheme, error) {

	if workspace == "" {
		return nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, model.ErrNoProjectSlug
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/projects/%v", workspace, projectKey)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
