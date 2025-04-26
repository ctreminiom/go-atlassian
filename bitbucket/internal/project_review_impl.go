package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/bitbucket"
	"net/http"
)

type ProjectReviewService struct {
	internalClient bitbucket.ProjectReviewerConnector
}

func (p *ProjectReviewService) Gets(ctx context.Context, workspace, projectKey string) (*model.ProjectReviewerPageScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx, workspace, projectKey)
}

func (p *ProjectReviewService) Get(ctx context.Context, workspace, projectKey, userSlug string) (*model.ProjectReviewerScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, workspace, projectKey, userSlug)
}

func (p *ProjectReviewService) Add(ctx context.Context, workspace, projectKey, userSlug string) (*model.ProjectReviewerScheme, *model.ResponseScheme, error) {
	return p.internalClient.Add(ctx, workspace, projectKey, userSlug)
}

func (p *ProjectReviewService) Remove(ctx context.Context, workspace, projectKey, userSlug string) (*model.ResponseScheme, error) {
	return p.internalClient.Remove(ctx, workspace, projectKey, userSlug)
}

func NewProjectReviewService(client service.Connector) *ProjectReviewService {
	return &ProjectReviewService{
		internalClient: &internalProjectServiceReviewerImpl{c: client},
	}
}

type internalProjectServiceReviewerImpl struct {
	c service.Connector
}

func (i internalProjectServiceReviewerImpl) Gets(ctx context.Context, workspace, projectKey string) (*model.ProjectReviewerPageScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, nil, model.ErrNoProjectSlug
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%s/projects/%s/default-reviewers", workspace, projectKey)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	reviewers := new(model.ProjectReviewerPageScheme)
	response, err := i.c.Call(request, reviewers)
	if err != nil {
		return nil, response, err
	}

	return reviewers, response, nil
}

func (i internalProjectServiceReviewerImpl) Get(ctx context.Context, workspace, projectKey, userSlug string) (*model.ProjectReviewerScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, nil, model.ErrNoProjectSlug
	}

	if userSlug == "" {
		return nil, nil, model.ErrNoAccountSlug
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%s/projects/%s/default-reviewers/%s", workspace, projectKey, userSlug)
	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	reviewer := new(model.ProjectReviewerScheme)
	response, err := i.c.Call(request, reviewer)
	if err != nil {
		return nil, response, err
	}

	return reviewer, response, nil
}

func (i internalProjectServiceReviewerImpl) Add(ctx context.Context, workspace, projectKey, userSlug string) (*model.ProjectReviewerScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, nil, model.ErrNoProjectSlug
	}

	if userSlug == "" {
		return nil, nil, model.ErrNoAccountSlug
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%s/projects/%s/default-reviewers/%s", workspace, projectKey, userSlug)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	reviewer := new(model.ProjectReviewerScheme)
	response, err := i.c.Call(request, reviewer)
	if err != nil {
		return nil, response, err
	}

	return reviewer, response, nil
}

func (i internalProjectServiceReviewerImpl) Remove(ctx context.Context, workspace, projectKey, userSlug string) (*model.ResponseScheme, error) {

	if workspace == "" {
		return nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, model.ErrNoProjectSlug
	}

	if userSlug == "" {
		return nil, model.ErrNoAccountSlug
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%s/projects/%s/default-reviewers/%s", workspace, projectKey, userSlug)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
