package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/bitbucket"
	"net/http"
)

func NewWorkspaceService(client service.Connector, webhook *WorkspaceHookService, permission *WorkspacePermissionService) *WorkspaceService {

	return &WorkspaceService{
		internalClient: &internalWorkspaceServiceImpl{c: client},
		Webhook:        webhook,
		Permission:     permission,
	}
}

type WorkspaceService struct {
	internalClient bitbucket.WorkspaceConnector
	Webhook        *WorkspaceHookService
	Permission     *WorkspacePermissionService
}

func (w *WorkspaceService) Get(ctx context.Context, workspace string) (*model.WorkspaceScheme, *model.ResponseScheme, error) {
	return w.internalClient.Get(ctx, workspace)
}

func (w *WorkspaceService) Members(ctx context.Context, workspace string) (*model.WorkspaceMembershipPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Members(ctx, workspace)
}

func (w *WorkspaceService) Membership(ctx context.Context, workspace, memberId string) (*model.WorkspaceMembershipScheme, *model.ResponseScheme, error) {
	return w.internalClient.Membership(ctx, workspace, memberId)
}

func (w *WorkspaceService) Projects(ctx context.Context, workspace string) (*model.BitbucketProjectPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Projects(ctx, workspace)
}

type internalWorkspaceServiceImpl struct {
	c service.Connector
}

func (i *internalWorkspaceServiceImpl) Get(ctx context.Context, workspace string) (*model.WorkspaceScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspaceError
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v", workspace)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	result := new(model.WorkspaceScheme)
	response, err := i.c.Call(request, result)
	if err != nil {
		return nil, response, err
	}

	return result, response, nil
}

func (i *internalWorkspaceServiceImpl) Members(ctx context.Context, workspace string) (*model.WorkspaceMembershipPageScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspaceError
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/members", workspace)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
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

func (i *internalWorkspaceServiceImpl) Membership(ctx context.Context, workspace, memberId string) (*model.WorkspaceMembershipScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspaceError
	}

	if memberId == "" {
		return nil, nil, model.ErrNoMemberIDError
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/members/%v", workspace, memberId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	member := new(model.WorkspaceMembershipScheme)
	response, err := i.c.Call(request, member)
	if err != nil {
		return nil, response, err
	}

	return member, response, nil
}

func (i *internalWorkspaceServiceImpl) Projects(ctx context.Context, workspace string) (*model.BitbucketProjectPageScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspaceError
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/projects", workspace)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BitbucketProjectPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}
