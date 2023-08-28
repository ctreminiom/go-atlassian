package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
	"strings"
)

func NewPermissionSchemeService(client service.Connector, version string, grant *PermissionSchemeGrantService) (*PermissionSchemeService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &PermissionSchemeService{
		internalClient: &internalPermissionSchemeImpl{c: client, version: version},
		Grant:          grant,
	}, nil
}

type PermissionSchemeService struct {
	internalClient jira.PermissionSchemeConnector
	Grant          *PermissionSchemeGrantService
}

// Gets returns all permission schemes.
//
// GET /rest/api/{2-3}/permissionscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#get-all-permission-schemes
func (p *PermissionSchemeService) Gets(ctx context.Context) (*model.PermissionSchemePageScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx)
}

// Get returns a permission scheme.
//
// GET /rest/api/{2-3}/permissionscheme/{schemeId}
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#get-permission-scheme
func (p *PermissionSchemeService) Get(ctx context.Context, permissionSchemeId int, expand []string) (*model.PermissionSchemeScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, permissionSchemeId, expand)
}

// Delete deletes a permission scheme.
//
// DELETE /rest/api/{2-3}/permissionscheme/{schemeId}
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#delete-permission-scheme
func (p *PermissionSchemeService) Delete(ctx context.Context, permissionSchemeId int) (*model.ResponseScheme, error) {
	return p.internalClient.Delete(ctx, permissionSchemeId)
}

// Create creates a new permission scheme.
//
// You can create a permission scheme with or without defining a set of permission grants.
//
// POST /rest/api/{2-3}/permissionscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#create-permission-scheme
func (p *PermissionSchemeService) Create(ctx context.Context, payload *model.PermissionSchemeScheme) (*model.PermissionSchemeScheme, *model.ResponseScheme, error) {
	return p.internalClient.Create(ctx, payload)
}

// Update updates a permission scheme.
// Below are some important things to note when using this resource:
//
// 1. If a permissions list is present in the request, then it is set in the permission scheme, overwriting all existing grants.
//
// 2. If you want to update only the name and description, then do not send a permissions list in the request.
//
// 3. Sending an empty list will remove all permission grants from the permission scheme.
//
// PUT /rest/api/{2-3}/permissionscheme/{schemeId}
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#update-permission-scheme
func (p *PermissionSchemeService) Update(ctx context.Context, permissionSchemeId int, payload *model.PermissionSchemeScheme) (*model.PermissionSchemeScheme, *model.ResponseScheme, error) {
	return p.internalClient.Update(ctx, permissionSchemeId, payload)
}

type internalPermissionSchemeImpl struct {
	c       service.Connector
	version string
}

func (i *internalPermissionSchemeImpl) Gets(ctx context.Context) (*model.PermissionSchemePageScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/permissionscheme", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.PermissionSchemePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalPermissionSchemeImpl) Get(ctx context.Context, permissionSchemeId int, expand []string) (*model.PermissionSchemeScheme, *model.ResponseScheme, error) {

	if permissionSchemeId == 0 {
		return nil, nil, model.ErrNoPermissionSchemeIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/permissionscheme/%v", i.version, permissionSchemeId))

	if expand != nil {

		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	permissionScheme := new(model.PermissionSchemeScheme)
	response, err := i.c.Call(request, permissionScheme)
	if err != nil {
		return nil, response, err
	}

	return permissionScheme, response, nil
}

func (i *internalPermissionSchemeImpl) Delete(ctx context.Context, permissionSchemeId int) (*model.ResponseScheme, error) {

	if permissionSchemeId == 0 {
		return nil, model.ErrNoPermissionSchemeIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/permissionscheme/%v", i.version, permissionSchemeId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalPermissionSchemeImpl) Create(ctx context.Context, payload *model.PermissionSchemeScheme) (*model.PermissionSchemeScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/permissionscheme", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	permissionScheme := new(model.PermissionSchemeScheme)
	response, err := i.c.Call(request, permissionScheme)
	if err != nil {
		return nil, response, err
	}

	return permissionScheme, response, nil
}

func (i *internalPermissionSchemeImpl) Update(ctx context.Context, permissionSchemeId int, payload *model.PermissionSchemeScheme) (*model.PermissionSchemeScheme, *model.ResponseScheme, error) {

	if permissionSchemeId == 0 {
		return nil, nil, model.ErrNoPermissionSchemeIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/permissionscheme/%v", i.version, permissionSchemeId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	permissionScheme := new(model.PermissionSchemeScheme)
	response, err := i.c.Call(request, permissionScheme)
	if err != nil {
		return nil, response, err
	}

	return permissionScheme, response, nil
}
