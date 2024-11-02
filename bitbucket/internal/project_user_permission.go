package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/bitbucket"
	"net/http"
)

type ProjectUserPermission struct {
	internalClient bitbucket.ProjectUserPermissionConnector
}

func (p *ProjectUserPermission) Gets(ctx context.Context, workspace, projectKey string) (*model.BitbucketProjectUserPermissionPageScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx, workspace, projectKey)
}

func (p *ProjectUserPermission) Get(ctx context.Context, workspace, projectKey, userSlug string) (*model.BitbucketProjectUserPermissionScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, workspace, projectKey, userSlug)
}

func (p *ProjectUserPermission) Update(ctx context.Context, workspace, projectKey, userSlug, permission string) (*model.BitbucketProjectUserPermissionScheme, *model.ResponseScheme, error) {
	return p.internalClient.Update(ctx, workspace, projectKey, userSlug, permission)
}

func (p *ProjectUserPermission) Delete(ctx context.Context, workspace, projectKey, userSlug string) (*model.ResponseScheme, error) {
	return p.internalClient.Delete(ctx, workspace, projectKey, userSlug)
}

func NewProjectUserPermissionService(client service.Connector) *ProjectUserPermission {
	return &ProjectUserPermission{
		internalClient: &internalProjectUserPermissionImpl{c: client},
	}
}

type internalProjectUserPermissionImpl struct {
	c service.Connector
}

func (i *internalProjectUserPermissionImpl) Gets(ctx context.Context, workspace, projectKey string) (*model.BitbucketProjectUserPermissionPageScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, nil, model.ErrNoProjectSlug
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%s/projects/%s/permissions-config/users", workspace, projectKey)
	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	permissions := new(model.BitbucketProjectUserPermissionPageScheme)
	response, err := i.c.Call(request, permissions)
	if err != nil {
		return nil, response, err
	}

	return permissions, response, nil
}

func (i *internalProjectUserPermissionImpl) Get(ctx context.Context, workspace, projectKey, userSlug string) (*model.BitbucketProjectUserPermissionScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, nil, model.ErrNoProjectSlug
	}

	if userSlug == "" {
		return nil, nil, model.ErrNoAccountSlug
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%s/projects/%s/permissions-config/users/%s", workspace, projectKey, userSlug)
	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	permission := new(model.BitbucketProjectUserPermissionScheme)
	response, err := i.c.Call(request, permission)
	if err != nil {
		return nil, response, err
	}

	return permission, response, nil
}

func (i *internalProjectUserPermissionImpl) Update(ctx context.Context, workspace, projectKey, userSlug, permission string) (*model.BitbucketProjectUserPermissionScheme, *model.ResponseScheme, error) {

	if workspace == "" {
		return nil, nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, nil, model.ErrNoProjectSlug
	}

	if userSlug == "" {
		return nil, nil, model.ErrNoAccountSlug
	}

	if permission == "" {
		return nil, nil, model.ErrNoPermissionSlug
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%s/projects/%s/permissions-config/users/%s", workspace, projectKey, userSlug)
	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]string{"permission": permission})
	if err != nil {
		return nil, nil, err
	}

	permissionUpdated := new(model.BitbucketProjectUserPermissionScheme)
	response, err := i.c.Call(request, permissionUpdated)
	if err != nil {
		return nil, response, err
	}

	return permissionUpdated, response, nil
}

func (i *internalProjectUserPermissionImpl) Delete(ctx context.Context, workspace, projectKey, userSlug string) (*model.ResponseScheme, error) {

	if workspace == "" {
		return nil, model.ErrNoWorkspace
	}

	if projectKey == "" {
		return nil, model.ErrNoProjectSlug
	}

	if userSlug == "" {
		return nil, model.ErrNoAccountSlug
	}

	endpoint := fmt.Sprintf("2.0/workspaces/%s/projects/%s/permissions-config/users/%s", workspace, projectKey, userSlug)
	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
