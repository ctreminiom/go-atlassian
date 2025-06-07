package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/bitbucket"
)

// NewWorkspaceService handles communication with the workspace related methods of the Bitbucket API.
func NewWorkspaceService(client service.Connector, webhook *WorkspaceHookService, permission *WorkspacePermissionService) *WorkspaceService {

	return &WorkspaceService{
		internalClient: &internalWorkspaceServiceImpl{c: client},
		Hook:           webhook,
		Permission:     permission,
	}
}

// WorkspaceService handles communication with the workspace related methods of the Bitbucket API.
type WorkspaceService struct {
	internalClient bitbucket.WorkspaceConnector
	Hook           *WorkspaceHookService
	Permission     *WorkspacePermissionService
}

// Get returns the requested workspace.
//
// GET /2.0/workspaces/{workspace}
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace#get-a-workspace
func (w *WorkspaceService) Get(ctx context.Context, workspace string) (*model.WorkspaceScheme, *model.ResponseScheme, error) {
	return w.internalClient.Get(ctx, workspace)
}

// Members returns all members of the requested workspace.
//
// GET /2.0/workspaces/{workspace}/members
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace#get-a-workspace
func (w *WorkspaceService) Members(ctx context.Context, workspace string) (*model.WorkspaceMembershipPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Members(ctx, workspace)
}

// Membership returns the workspace membership.
//
// which includes a User object for the member and a Workspace object for the requested workspace.
//
// GET /2.0/workspaces/{workspace}/members/{memberID}
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace#get-member-in-a-workspace
func (w *WorkspaceService) Membership(ctx context.Context, workspace, memberID string) (*model.WorkspaceMembershipScheme, *model.ResponseScheme, error) {
	return w.internalClient.Membership(ctx, workspace, memberID)
}

// Projects returns the list of projects in this workspace.
//
// GET /2.0/workspaces/{workspace}/projects
//
// https://docs.go-atlassian.io/bitbucket-cloud/workspace#get-projects-in-a-workspace
func (w *WorkspaceService) Projects(ctx context.Context, workspace string) (*model.BitbucketProjectPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Projects(ctx, workspace)
}

type internalWorkspaceServiceImpl struct {
	c service.Connector
}

// Get returns the requested workspace.
func (i *internalWorkspaceServiceImpl) Get(ctx context.Context, workspace string) (*model.WorkspaceScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoWorkspace)
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

// Members returns all members of the requested workspace.
func (i *internalWorkspaceServiceImpl) Members(ctx context.Context, workspace string) (*model.WorkspaceMembershipPageScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoWorkspace)
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

// Membership returns the workspace membership.
func (i *internalWorkspaceServiceImpl) Membership(ctx context.Context, workspace, memberID string) (*model.WorkspaceMembershipScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoWorkspace)
	}

	if memberID == "" {
		return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoMemberID)
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%v/members/%v", workspace, memberID)

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

// Projects returns the list of projects in this workspace.
func (i *internalWorkspaceServiceImpl) Projects(ctx context.Context, workspace string) (*model.BitbucketProjectPageScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, fmt.Errorf("bitbucket: %w", model.ErrNoWorkspace)
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
