package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func NewDashboardService(client service.Client, version string) (jira.Dashboard, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &DashboardService{client, version}, nil
}

type DashboardService struct {
	c       service.Client
	version string
}

func (d DashboardService) Gets(ctx context.Context, startAt, maxResults int, filter string) (*model.DashboardPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(filter) != 0 {
		params.Add("filter", filter)
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard?%v", d.version, params.Encode())

	request, err := d.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.DashboardPageScheme)
	response, err := d.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (d DashboardService) Create(ctx context.Context, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error) {

	reader, err := d.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard", d.version)

	request, err := d.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	dashboard := new(model.DashboardScheme)
	response, err := d.c.Call(request, dashboard)
	if err != nil {
		return nil, response, err
	}

	return dashboard, response, nil
}

func (d DashboardService) Search(ctx context.Context, options *model.DashboardSearchOptionsScheme, startAt, maxResults int) (*model.DashboardSearchPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		if len(options.OwnerAccountID) != 0 {
			params.Add("accountId", options.OwnerAccountID)
		}

		if len(options.DashboardName) != 0 {
			params.Add("dashboardName", options.OwnerAccountID)
		}

		if len(options.GroupPermissionName) != 0 {
			params.Add("groupname", options.OwnerAccountID)
		}

		if len(options.OrderBy) != 0 {
			params.Add("orderBy", options.OwnerAccountID)
		}

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard/search?%s", d.version, params.Encode())

	request, err := d.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.DashboardSearchPageScheme)
	response, err := d.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (d DashboardService) Get(ctx context.Context, dashboardId string) (*model.DashboardScheme, *model.ResponseScheme, error) {

	if dashboardId == "" {
		return nil, nil, model.ErrNoDashboardIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard/%v", d.version, dashboardId)

	request, err := d.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	dashboard := new(model.DashboardScheme)
	response, err := d.c.Call(request, dashboard)
	if err != nil {
		return nil, response, err
	}

	return dashboard, response, nil
}

func (d DashboardService) Delete(ctx context.Context, dashboardId string) (*model.ResponseScheme, error) {

	if dashboardId == "" {
		return nil, model.ErrNoDashboardIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard/%v", d.version, dashboardId)

	request, err := d.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	response, err := d.c.Call(request, nil)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (d DashboardService) Copy(ctx context.Context, dashboardId string, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error) {

	if dashboardId == "" {
		return nil, nil, model.ErrNoDashboardIDError
	}

	reader, err := d.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard/%v/copy", d.version, dashboardId)

	request, err := d.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	dashboard := new(model.DashboardScheme)
	response, err := d.c.Call(request, dashboard)
	if err != nil {
		return nil, response, err
	}

	return dashboard, response, nil
}

func (d DashboardService) Update(ctx context.Context, dashboardId string, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error) {

	if dashboardId == "" {
		return nil, nil, model.ErrNoDashboardIDError
	}

	reader, err := d.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard/%v", d.version, dashboardId)
	request, err := d.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	dashboard := new(model.DashboardScheme)
	response, err := d.c.Call(request, dashboard)
	if err != nil {
		return nil, response, err
	}

	return dashboard, response, nil
}
