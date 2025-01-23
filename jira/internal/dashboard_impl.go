package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewDashboardService creates a new instance of DashboardService.
// It takes a service.Connector and a version string as input.
// Returns a pointer to DashboardService and an error if the version is not provided.
func NewDashboardService(client service.Connector, version string) (*DashboardService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &DashboardService{
		internalClient: &internalDashboardImpl{c: client, version: version},
	}, nil
}

// DashboardService provides methods to interact with dashboard operations in Jira Service Management.
type DashboardService struct {
	// internalClient is the connector interface for dashboard operations.
	internalClient jira.DashboardConnector
}

// Gets returns a list of dashboards owned by or shared with the user.
//
// The list may be filtered to include only favorite or owned dashboards.
//
// GET /rest/api/{3-2}/dashboard
//
// https://docs.go-atlassian.io/jira-software-cloud/dashboards#get-all-dashboards
func (d *DashboardService) Gets(ctx context.Context, startAt, maxResults int, filter string) (*model.DashboardPageScheme, *model.ResponseScheme, error) {
	return d.internalClient.Gets(ctx, startAt, maxResults, filter)
}

// Create creates a dashboard.
//
// POST /rest/api/{3-2}/dashboard
//
// https://docs.go-atlassian.io/jira-software-cloud/dashboards#create-dashboard
func (d *DashboardService) Create(ctx context.Context, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error) {
	return d.internalClient.Create(ctx, payload)
}

// Search returns a paginated list of dashboards.
//
// This operation is similar to Get dashboards except that the results can be refined to include dashboards that have specific attributes.
//
// GET /rest/api/{2-3}/dashboard/search
//
// https://docs.go-atlassian.io/jira-software-cloud/dashboards#search-for-dashboards
func (d *DashboardService) Search(ctx context.Context, options *model.DashboardSearchOptionsScheme, startAt, maxResults int) (*model.DashboardSearchPageScheme, *model.ResponseScheme, error) {
	return d.internalClient.Search(ctx, options, startAt, maxResults)
}

// Get returns a dashboard.
//
// GET /rest/api/{2-3}/dashboard/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/dashboards#get-dashboard
func (d *DashboardService) Get(ctx context.Context, dashboardID string) (*model.DashboardScheme, *model.ResponseScheme, error) {
	return d.internalClient.Get(ctx, dashboardID)
}

// Delete deletes a dashboard.
//
// DELETE /rest/api/{2-3}/dashboard/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/dashboards#delete-dashboard
func (d *DashboardService) Delete(ctx context.Context, dashboardID string) (*model.ResponseScheme, error) {
	return d.internalClient.Delete(ctx, dashboardID)
}

// Copy copies a dashboard.
//
// Any values provided in the dashboard parameter replace those in the copied dashboard.
//
// POST /rest/api/{2-3}/dashboard/{id}/copy
//
// https://docs.go-atlassian.io/jira-software-cloud/dashboards#copy-dashboard
func (d *DashboardService) Copy(ctx context.Context, dashboardID string, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error) {
	return d.internalClient.Copy(ctx, dashboardID, payload)
}

// Update updates a dashboard
//
// PUT /rest/api/{2-3}/dashboard/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/dashboards#update-dashboard
func (d *DashboardService) Update(ctx context.Context, dashboardID string, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error) {
	return d.internalClient.Update(ctx, dashboardID, payload)
}

type internalDashboardImpl struct {
	c       service.Connector
	version string
}

func (i *internalDashboardImpl) Gets(ctx context.Context, startAt, maxResults int, filter string) (*model.DashboardPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(filter) != 0 {
		params.Add("filter", filter)
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.DashboardPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalDashboardImpl) Create(ctx context.Context, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/dashboard", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	dashboard := new(model.DashboardScheme)
	response, err := i.c.Call(request, dashboard)
	if err != nil {
		return nil, response, err
	}

	return dashboard, response, nil
}

func (i *internalDashboardImpl) Search(ctx context.Context, options *model.DashboardSearchOptionsScheme, startAt, maxResults int) (*model.DashboardSearchPageScheme, *model.ResponseScheme, error) {

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

	endpoint := fmt.Sprintf("rest/api/%v/dashboard/search?%s", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.DashboardSearchPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalDashboardImpl) Get(ctx context.Context, dashboardID string) (*model.DashboardScheme, *model.ResponseScheme, error) {

	if dashboardID == "" {
		return nil, nil, model.ErrNoDashboardID
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard/%v", i.version, dashboardID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	dashboard := new(model.DashboardScheme)
	response, err := i.c.Call(request, dashboard)
	if err != nil {
		return nil, response, err
	}

	return dashboard, response, nil
}

func (i *internalDashboardImpl) Delete(ctx context.Context, dashboardID string) (*model.ResponseScheme, error) {

	if dashboardID == "" {
		return nil, model.ErrNoDashboardID
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard/%v", i.version, dashboardID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalDashboardImpl) Copy(ctx context.Context, dashboardID string, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error) {

	if dashboardID == "" {
		return nil, nil, model.ErrNoDashboardID
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard/%v/copy", i.version, dashboardID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	dashboard := new(model.DashboardScheme)
	response, err := i.c.Call(request, dashboard)
	if err != nil {
		return nil, response, err
	}

	return dashboard, response, nil
}

func (i *internalDashboardImpl) Update(ctx context.Context, dashboardID string, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error) {

	if dashboardID == "" {
		return nil, nil, model.ErrNoDashboardID
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard/%v", i.version, dashboardID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	dashboard := new(model.DashboardScheme)
	response, err := i.c.Call(request, dashboard)
	if err != nil {
		return nil, response, err
	}

	return dashboard, response, nil
}
