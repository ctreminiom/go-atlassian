package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

func NewFilterShareService(client service.Client, version string) (jira.FilterShare, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &FilterShareService{client, version}, nil
}

type FilterShareService struct {
	c       service.Client
	version string
}

func (f FilterShareService) Scope(ctx context.Context) (*model.ShareFilterScopeScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/filter/defaultShareScope", f.version)

	request, err := f.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	settings := new(model.ShareFilterScopeScheme)
	response, err := f.c.Call(request, settings)
	if err != nil {
		return nil, response, err
	}

	return settings, response, nil
}

func (f FilterShareService) SetScope(ctx context.Context, scope string) (*model.ResponseScheme, error) {

	reader, err := f.c.TransformStructToReader(&model.ShareFilterScopeScheme{Scope: scope})
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/defaultShareScope", f.version)

	request, err := f.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return f.c.Call(request, nil)
}

func (f FilterShareService) Gets(ctx context.Context, filterId int) ([]*model.SharePermissionScheme, *model.ResponseScheme, error) {

	if filterId == 0 {
		return nil, nil, model.ErrNoFilterIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v/permission", f.version, filterId)

	request, err := f.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var filters []*model.SharePermissionScheme
	response, err := f.c.Call(request, filters)
	if err != nil {
		return nil, response, err
	}

	return filters, response, nil
}

func (f FilterShareService) Add(ctx context.Context, filterId int, payload *model.PermissionFilterPayloadScheme) ([]*model.SharePermissionScheme, *model.ResponseScheme, error) {

	if filterId == 0 {
		return nil, nil, model.ErrNoFilterIDError
	}

	reader, err := f.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v/permission", f.version, filterId)

	request, err := f.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	var permissions []*model.SharePermissionScheme
	response, err := f.c.Call(request, permissions)
	if err != nil {
		return nil, response, err
	}

	return permissions, response, nil
}

func (f FilterShareService) Get(ctx context.Context, filterId, permissionId int) (*model.SharePermissionScheme, *model.ResponseScheme, error) {

	if filterId == 0 {
		return nil, nil, model.ErrNoFilterIDError
	}

	if permissionId == 0 {
		return nil, nil, model.ErrNoPermissionGrantIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v/permission/%v", f.version, filterId, permissionId)

	request, err := f.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	permission := new(model.SharePermissionScheme)
	response, err := f.c.Call(request, permission)
	if err != nil {
		return nil, response, err
	}

	return permission, response, nil
}

func (f FilterShareService) Delete(ctx context.Context, filterId, permissionId int) (*model.ResponseScheme, error) {

	if filterId == 0 {
		return nil, model.ErrNoFilterIDError
	}

	if permissionId == 0 {
		return nil, model.ErrNoPermissionGrantIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v/permission/%v", f.version, filterId, permissionId)

	request, err := f.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return f.c.Call(request, nil)
}
