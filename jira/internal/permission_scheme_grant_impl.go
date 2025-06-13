package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
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
// POST /rest/api/{2-3}/permissionscheme/{permissionSchemeID}/permission
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#create-permission-grant
func (p *PermissionSchemeGrantService) Create(ctx context.Context, permissionSchemeID int, payload *model.PermissionGrantPayloadScheme) (*model.PermissionGrantScheme, *model.ResponseScheme, error) {
	return p.internalClient.Create(ctx, permissionSchemeID, payload)
}

// Gets returns all permission grants for a permission scheme.
//
// GET /rest/api/{2-3}/permissionscheme/{permissionSchemeID}/permission
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#get-permission-scheme-grants
func (p *PermissionSchemeGrantService) Gets(ctx context.Context, permissionSchemeID int, expand []string) (*model.PermissionSchemeGrantsScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx, permissionSchemeID, expand)
}

// Get returns a permission grant.
//
// GET /rest/api/{2-3}/permissionscheme/{permissionSchemeID}/permission/{permissionGrantID}
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#get-permission-scheme-grant
func (p *PermissionSchemeGrantService) Get(ctx context.Context, permissionSchemeID, permissionGrantID int, expand []string) (*model.PermissionGrantScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, permissionSchemeID, permissionGrantID, expand)
}

// Delete deletes a permission grant from a permission scheme. See About permission schemes and grants for more details.
//
// DELETE /rest/api/{2-3}/permissionscheme/{permissionSchemeID}/permission/{permissionGrantID}
//
// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#delete-permission-scheme-grant
func (p *PermissionSchemeGrantService) Delete(ctx context.Context, permissionSchemeID, permissionGrantID int) (*model.ResponseScheme, error) {
	return p.internalClient.Delete(ctx, permissionSchemeID, permissionGrantID)
}

type internalPermissionSchemeGrantImpl struct {
	c       service.Connector
	version string
}

func (i *internalPermissionSchemeGrantImpl) Create(ctx context.Context, permissionSchemeID int, payload *model.PermissionGrantPayloadScheme) (*model.PermissionGrantScheme, *model.ResponseScheme, error) {

	if permissionSchemeID == 0 {
		return nil, nil, model.ErrNoPermissionSchemeID
	}

	endpoint := fmt.Sprintf("rest/api/%v/permissionscheme/%v/permission", i.version, permissionSchemeID)

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

func (i *internalPermissionSchemeGrantImpl) Gets(ctx context.Context, permissionSchemeID int, expand []string) (*model.PermissionSchemeGrantsScheme, *model.ResponseScheme, error) {

	if permissionSchemeID == 0 {
		return nil, nil, model.ErrNoPermissionSchemeID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/permissionscheme/%v/permission", i.version, permissionSchemeID))

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

func (i *internalPermissionSchemeGrantImpl) Get(ctx context.Context, permissionSchemeID, permissionGrantID int, expand []string) (*model.PermissionGrantScheme, *model.ResponseScheme, error) {

	if permissionSchemeID == 0 {
		return nil, nil, model.ErrNoPermissionSchemeID
	}

	if permissionGrantID == 0 {
		return nil, nil, model.ErrNoPermissionGrantID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/permissionscheme/%v/permission/%v", i.version, permissionSchemeID, permissionGrantID))

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

func (i *internalPermissionSchemeGrantImpl) Delete(ctx context.Context, permissionSchemeID, permissionGrantID int) (*model.ResponseScheme, error) {

	if permissionSchemeID == 0 {
		return nil, model.ErrNoPermissionSchemeID
	}

	if permissionGrantID == 0 {
		return nil, model.ErrNoPermissionGrantID
	}

	endpoint := fmt.Sprintf("rest/api/%v/permissionscheme/%v/permission/%v", i.version, permissionSchemeID, permissionGrantID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
