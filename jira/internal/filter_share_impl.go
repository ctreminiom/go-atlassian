package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewFilterShareService creates a new instance of FilterShareService.
// It takes a service.Connector and a version string as input.
// Returns a pointer to FilterShareService and an error if the version is not provided.
func NewFilterShareService(client service.Connector, version string) (*FilterShareService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &FilterShareService{
		internalClient: &internalFilterShareImpl{c: client, version: version},
	}, nil
}

// FilterShareService provides methods to manage filter sharing in Jira Service Management.
type FilterShareService struct {
	// internalClient is the connector interface for filter sharing operations.
	internalClient jira.FilterSharingConnector
}

// Scope returns the default sharing settings for new filters and dashboards for a user.
//
// GET /rest/api/{2-3}/filter/defaultShareScope
//
// https://docs.go-atlassian.io/jira-software-cloud/filters/sharing#get-default-share-scope
func (f *FilterShareService) Scope(ctx context.Context) (*model.ShareFilterScopeScheme, *model.ResponseScheme, error) {
	return f.internalClient.Scope(ctx)
}

// SetScope sets the default sharing for new filters and dashboards for a user.
//
// PUT /rest/api/{2-3}/filter/defaultShareScope
//
// https://docs.go-atlassian.io/jira-software-cloud/filters/sharing#set-default-share-scope
func (f *FilterShareService) SetScope(ctx context.Context, scope string) (*model.ResponseScheme, error) {
	return f.internalClient.SetScope(ctx, scope)
}

// Gets returns the share permissions for a filter.
//
// 1.A filter can be shared with groups, projects, all logged-in users, or the public.
//
// 2.Sharing with all logged-in users or the public is known as a global share permission.
//
// GET /rest/api/{2-3}/filter/{id}/permission
//
// https://docs.go-atlassian.io/jira-software-cloud/filters/sharing#get-share-permissions
func (f *FilterShareService) Gets(ctx context.Context, filterID int) ([]*model.SharePermissionScheme, *model.ResponseScheme, error) {
	return f.internalClient.Gets(ctx, filterID)
}

// Add a share permissions to a filter.
//
// If you add a global share permission (one for all logged-in users or the public)
//
// it will overwrite all share permissions for the filter.
//
// POST /rest/api/{2-3}/filter/{id}/permission
//
// https://docs.go-atlassian.io/jira-software-cloud/filters/sharing#add-share-permission
func (f *FilterShareService) Add(ctx context.Context, filterID int, payload *model.PermissionFilterPayloadScheme) ([]*model.SharePermissionScheme, *model.ResponseScheme, error) {
	return f.internalClient.Add(ctx, filterID, payload)
}

// Get returns a share permission for a filter.
//
// A filter can be shared with groups, projects, all logged-in users, or the public.
//
// Sharing with all logged-in users or the public is known as a global share permission.
//
// GET /rest/api/{2-3}/filter/{id}/permission/{permissionID}
//
// https://docs.go-atlassian.io/jira-software-cloud/filters/sharing#get-share-permission
func (f *FilterShareService) Get(ctx context.Context, filterID, permissionID int) (*model.SharePermissionScheme, *model.ResponseScheme, error) {
	return f.internalClient.Get(ctx, filterID, permissionID)
}

// Delete deletes a share permission from a filter.
//
// DELETE /rest/api/{2-3}/filter/{id}/permission/{permissionID}
//
// https://docs.go-atlassian.io/jira-software-cloud/filters/sharing#delete-share-permission
func (f *FilterShareService) Delete(ctx context.Context, filterID, permissionID int) (*model.ResponseScheme, error) {
	return f.internalClient.Delete(ctx, filterID, permissionID)
}

type internalFilterShareImpl struct {
	c       service.Connector
	version string
}

func (i *internalFilterShareImpl) Scope(ctx context.Context) (*model.ShareFilterScopeScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/filter/defaultShareScope", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	settings := new(model.ShareFilterScopeScheme)
	response, err := i.c.Call(request, settings)
	if err != nil {
		return nil, response, err
	}

	return settings, response, nil
}

func (i *internalFilterShareImpl) SetScope(ctx context.Context, scope string) (*model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/filter/defaultShareScope", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", &model.ShareFilterScopeScheme{Scope: scope})
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalFilterShareImpl) Gets(ctx context.Context, filterID int) ([]*model.SharePermissionScheme, *model.ResponseScheme, error) {

	if filterID == 0 {
		return nil, nil, model.ErrNoFilterID
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v/permission", i.version, filterID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var filters []*model.SharePermissionScheme
	response, err := i.c.Call(request, filters)
	if err != nil {
		return nil, response, err
	}

	return filters, response, nil
}

func (i *internalFilterShareImpl) Add(ctx context.Context, filterID int, payload *model.PermissionFilterPayloadScheme) ([]*model.SharePermissionScheme, *model.ResponseScheme, error) {

	if filterID == 0 {
		return nil, nil, model.ErrNoFilterID
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v/permission", i.version, filterID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	var permissions []*model.SharePermissionScheme
	response, err := i.c.Call(request, permissions)
	if err != nil {
		return nil, response, err
	}

	return permissions, response, nil
}

func (i *internalFilterShareImpl) Get(ctx context.Context, filterID, permissionID int) (*model.SharePermissionScheme, *model.ResponseScheme, error) {

	if filterID == 0 {
		return nil, nil, model.ErrNoFilterID
	}

	if permissionID == 0 {
		return nil, nil, model.ErrNoPermissionGrantID
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v/permission/%v", i.version, filterID, permissionID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	permission := new(model.SharePermissionScheme)
	response, err := i.c.Call(request, permission)
	if err != nil {
		return nil, response, err
	}

	return permission, response, nil
}

func (i *internalFilterShareImpl) Delete(ctx context.Context, filterID, permissionID int) (*model.ResponseScheme, error) {

	if filterID == 0 {
		return nil, model.ErrNoFilterID
	}

	if permissionID == 0 {
		return nil, model.ErrNoPermissionGrantID
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v/permission/%v", i.version, filterID, permissionID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
