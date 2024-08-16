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

// NewPermissionSchemeGrantService creates a new instance of PermissionSchemeGrantService.
func NewPermissionSchemeGrantService(client service.Connector, version string) (*PermissionSchemeGrantService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &PermissionSchemeGrantService{
		internalClient: &internalPermissionSchemeGrantImpl{c: client, version: version},
	}, nil
}

// PermissionSchemeGrantService provides methods to manage permission scheme grants in Jira Service Management.
type PermissionSchemeGrantService struct {
	// internalClient is the connector interface for permission scheme grant operations.
	internalClient jira.PermissionSchemeGrantConnector
}

// Create creates a permission grant in a permission scheme.
//
// POST /rest/api/{2-3}/permissionscheme/{schemeId}/permission
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#create-permission-grant
func (p *PermissionSchemeGrantService) Create(ctx context.Context, permissionSchemeId int, payload *model.PermissionGrantPayloadScheme) (*model.PermissionGrantScheme, *model.ResponseScheme, error) {
	return p.internalClient.Create(ctx, permissionSchemeId, payload)
}

// Gets returns all permission grants for a permission scheme.
//
// GET /rest/api/{2-3}/permissionscheme/{schemeId}/permission
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#get-permission-scheme-grants
func (p *PermissionSchemeGrantService) Gets(ctx context.Context, permissionSchemeId int, expand []string) (*model.PermissionSchemeGrantsScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx, permissionSchemeId, expand)
}

// Get returns a permission grant.
//
// GET /rest/api/{2-3}/permissionscheme/{schemeId}/permission/{permissionId}
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#get-permission-scheme-grant
func (p *PermissionSchemeGrantService) Get(ctx context.Context, permissionSchemeId, permissionGrantId int, expand []string) (*model.PermissionGrantScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, permissionSchemeId, permissionGrantId, expand)
}

// Delete deletes a permission grant from a permission scheme. See About permission schemes and grants for more details.
//
// DELETE /rest/api/{2-3}/permissionscheme/{schemeId}/permission/{permissionId}
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#delete-permission-scheme-grant
func (p *PermissionSchemeGrantService) Delete(ctx context.Context, permissionSchemeId, permissionGrantId int) (*model.ResponseScheme, error) {
	return p.internalClient.Delete(ctx, permissionSchemeId, permissionGrantId)
}

type internalPermissionSchemeGrantImpl struct {
	c       service.Connector
	version string
}

func (i *internalPermissionSchemeGrantImpl) Create(ctx context.Context, permissionSchemeId int, payload *model.PermissionGrantPayloadScheme) (*model.PermissionGrantScheme, *model.ResponseScheme, error) {

	if permissionSchemeId == 0 {
		return nil, nil, model.ErrNoPermissionSchemeIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/permissionscheme/%v/permission", i.version, permissionSchemeId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	grant := new(model.PermissionGrantScheme)
	response, err := i.c.Call(request, grant)
	if err != nil {
		return nil, response, err
	}

	return grant, response, nil
}

func (i *internalPermissionSchemeGrantImpl) Gets(ctx context.Context, permissionSchemeId int, expand []string) (*model.PermissionSchemeGrantsScheme, *model.ResponseScheme, error) {

	if permissionSchemeId == 0 {
		return nil, nil, model.ErrNoPermissionSchemeIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/permissionscheme/%v/permission", i.version, permissionSchemeId))

	if expand != nil {

		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	grants := new(model.PermissionSchemeGrantsScheme)
	response, err := i.c.Call(request, grants)
	if err != nil {
		return nil, response, err
	}

	return grants, response, nil
}

func (i *internalPermissionSchemeGrantImpl) Get(ctx context.Context, permissionSchemeId, permissionGrantId int, expand []string) (*model.PermissionGrantScheme, *model.ResponseScheme, error) {

	if permissionSchemeId == 0 {
		return nil, nil, model.ErrNoPermissionSchemeIDError
	}

	if permissionGrantId == 0 {
		return nil, nil, model.ErrNoPermissionGrantIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/permissionscheme/%v/permission/%v", i.version, permissionSchemeId, permissionGrantId))

	if expand != nil {

		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	grant := new(model.PermissionGrantScheme)
	response, err := i.c.Call(request, grant)
	if err != nil {
		return nil, response, err
	}

	return grant, response, nil
}

func (i *internalPermissionSchemeGrantImpl) Delete(ctx context.Context, permissionSchemeId, permissionGrantId int) (*model.ResponseScheme, error) {

	if permissionSchemeId == 0 {
		return nil, model.ErrNoPermissionSchemeIDError
	}

	if permissionGrantId == 0 {
		return nil, model.ErrNoPermissionGrantIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/permissionscheme/%v/permission/%v", i.version, permissionSchemeId, permissionGrantId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
