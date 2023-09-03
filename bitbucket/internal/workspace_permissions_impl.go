package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/bitbucket"
	"net/http"
	"net/url"
	"strings"
)

func NewWorkspacePermissionService(client service.Connector) *WorkspacePermissionService {

	return &WorkspacePermissionService{
		internalClient: &internalWorkspacePermissionServiceImpl{c: client},
	}
}

type WorkspacePermissionService struct {
	internalClient bitbucket.WorkspacePermissionConnector
}

func (w *WorkspacePermissionService) Members(ctx context.Context, workspace, query string) (*model.WorkspaceMembershipPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Members(ctx, workspace, query)
}

func (w *WorkspacePermissionService) Repositories(ctx context.Context, workspace, query, sort string) (*model.RepositoryPermissionPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Repositories(ctx, workspace, query, sort)
}

func (w *WorkspacePermissionService) Repository(ctx context.Context, workspace, repository, query, sort string) (*model.RepositoryPermissionPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Repository(ctx, workspace, repository, query, sort)
}

type internalWorkspacePermissionServiceImpl struct {
	c service.Connector
}

func (i *internalWorkspacePermissionServiceImpl) Members(ctx context.Context, workspace, query string) (*model.WorkspaceMembershipPageScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspaceError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("2.0/workspaces/%v/permissions", workspace))

	if query != "" {

		params := url.Values{}
		params.Add("q", query)

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.WorkspaceMembershipPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalWorkspacePermissionServiceImpl) Repositories(ctx context.Context, workspace, query, sort string) (*model.RepositoryPermissionPageScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspaceError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("2.0/workspaces/%v/permissions/repositories", workspace))

	params := url.Values{}
	if query != "" {
		params.Add("q", query)
	}
	if sort != "" {
		params.Add("sort", sort)
	}

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.RepositoryPermissionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalWorkspacePermissionServiceImpl) Repository(ctx context.Context, workspace, repository, query, sort string) (*model.RepositoryPermissionPageScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspaceError
	}

	if repository == "" {
		return nil, nil, model.ErrNoRepositoryError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("2.0/workspaces/%v/permissions/repositories/%v", workspace, repository))

	params := url.Values{}

	if query != "" {
		params.Add("q", query)
	}
	if sort != "" {
		params.Add("sort", sort)
	}

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.RepositoryPermissionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}
