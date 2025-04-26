package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/bitbucket"
	"net/http"
)

type ProjectGroupPermission struct {
	internalClient bitbucket.ProjectGroupPermissionConnector
}

func (p *ProjectGroupPermission) Gets(ctx context.Context, workspace, projectKey string) (*model.BitbucketProjectGroupPermissionPageScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx, workspace, projectKey)
}

func (p *ProjectGroupPermission) Get(ctx context.Context, workspace, projectKey, groupSlug string) (*model.ProjectGroupPermissionScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, workspace, projectKey, groupSlug)
}

func (p *ProjectGroupPermission) Update(ctx context.Context, workspace, projectKey, groupSlug, permission string) (*model.ProjectGroupPermissionScheme, *model.ResponseScheme, error) {
	return p.internalClient.Update(ctx, workspace, projectKey, groupSlug, permission)
}

func (p *ProjectGroupPermission) Delete(ctx context.Context, workspace, projectKey, groupSlug string) (*model.ResponseScheme, error) {
	return p.internalClient.Delete(ctx, workspace, projectKey, groupSlug)
}

func NewProjectGroupPermissionService(client service.Connector) *ProjectGroupPermission {
	return &ProjectGroupPermission{
		internalClient: &internalProjectGroupPermissionImpl{c: client},
	}
}

type internalProjectGroupPermissionImpl struct {
	c service.Connector
}

func (i *internalProjectGroupPermissionImpl) Gets(ctx context.Context, workspace, projectKey string) (*model.BitbucketProjectGroupPermissionPageScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, nil, model.ErrNoProjectSlug
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%s/projects/%s/permissions-config/groups", workspace, projectKey)
	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	permissions := new(model.BitbucketProjectGroupPermissionPageScheme)
	response, err := i.c.Call(request, permissions)
	if err != nil {
		return nil, response, err
	}

	return permissions, response, nil
}

func (i *internalProjectGroupPermissionImpl) Get(ctx context.Context, workspace, projectKey, groupSlug string) (*model.ProjectGroupPermissionScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, nil, model.ErrNoProjectSlug
	}

	if groupSlug == "" {
		return nil, nil, model.ErrNoGroupSlug
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%s/projects/%s/permissions-config/groups/%s", workspace, projectKey, groupSlug)
	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	permission := new(model.ProjectGroupPermissionScheme)
	response, err := i.c.Call(request, permission)
	if err != nil {
		return nil, response, err
	}

	return permission, response, nil
}

func (i *internalProjectGroupPermissionImpl) Update(ctx context.Context, workspace, projectKey, groupSlug, permission string) (*model.ProjectGroupPermissionScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, nil, model.ErrNoProjectSlug
	}

	if groupSlug == "" {
		return nil, nil, model.ErrNoGroupSlug
	}

	if permission == "" {
		return nil, nil, model.ErrNoPermissionSlug
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%s/projects/%s/permissions-config/groups/%s", workspace, projectKey, groupSlug)
	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]string{"permission": permission})
	if err != nil {
		return nil, nil, err
	}

	permissionUpdated := new(model.ProjectGroupPermissionScheme)
	response, err := i.c.Call(request, permissionUpdated)
	if err != nil {
		return nil, response, err
	}

	return permissionUpdated, response, nil
}

func (i *internalProjectGroupPermissionImpl) Delete(ctx context.Context, workspace, projectKey, groupSlug string) (*model.ResponseScheme, error) {

	if workspace == "" {
		return nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, model.ErrNoProjectSlug
	}

	if groupSlug == "" {
		return nil, model.ErrNoGroupSlug
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%s/projects/%s/permissions-config/groups/%s", workspace, projectKey, groupSlug)
	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
