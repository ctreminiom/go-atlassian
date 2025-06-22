package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewDashboardService creates a new instance of DashboardService.
// It takes a service.Connector and a version string as input.
// Returns a pointer to DashboardService and an error if the version is not provided.
func NewDashboardService(client service.Connector, version string) (*DashboardService, error) {

	if version == "" {
		return nil, fmt.Errorf("jira: %w", model.ErrNoVersionProvided)
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
	ctx, span := tracer().Start(ctx, "(*DashboardService).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_dashboards"),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
		attribute.String("jira.filter", filter),
	)

	result, response, err := d.internalClient.Gets(ctx, startAt, maxResults, filter)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Create creates a dashboard.
//
// POST /rest/api/{3-2}/dashboard
//
// https://docs.go-atlassian.io/jira-software-cloud/dashboards#create-dashboard
func (d *DashboardService) Create(ctx context.Context, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*DashboardService).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "create_dashboard"),
	)

	result, response, err := d.internalClient.Create(ctx, payload)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Search returns a paginated list of dashboards.
//
// This operation is similar to Get dashboards except that the results can be refined to include dashboards that have specific attributes.
//
// GET /rest/api/{2-3}/dashboard/search
//
// https://docs.go-atlassian.io/jira-software-cloud/dashboards#search-for-dashboards
func (d *DashboardService) Search(ctx context.Context, options *model.DashboardSearchOptionsScheme, startAt, maxResults int) (*model.DashboardSearchPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*DashboardService).Search", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "search_dashboards"),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
	)

	result, response, err := d.internalClient.Search(ctx, options, startAt, maxResults)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Get returns a dashboard.
//
// GET /rest/api/{2-3}/dashboard/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/dashboards#get-dashboard
func (d *DashboardService) Get(ctx context.Context, dashboardID string) (*model.DashboardScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*DashboardService).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_dashboard"),
		attribute.String("jira.dashboard.id", dashboardID),
	)

	result, response, err := d.internalClient.Get(ctx, dashboardID)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Delete deletes a dashboard.
//
// DELETE /rest/api/{2-3}/dashboard/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/dashboards#delete-dashboard
func (d *DashboardService) Delete(ctx context.Context, dashboardID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*DashboardService).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete_dashboard"),
		attribute.String("jira.dashboard.id", dashboardID),
	)

	response, err := d.internalClient.Delete(ctx, dashboardID)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

// Copy copies a dashboard.
//
// Any values provided in the dashboard parameter replace those in the copied dashboard.
//
// POST /rest/api/{2-3}/dashboard/{id}/copy
//
// https://docs.go-atlassian.io/jira-software-cloud/dashboards#copy-dashboard
func (d *DashboardService) Copy(ctx context.Context, dashboardID string, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*DashboardService).Copy", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "copy_dashboard"),
		attribute.String("jira.dashboard.id", dashboardID),
	)

	result, response, err := d.internalClient.Copy(ctx, dashboardID, payload)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Update updates a dashboard
//
// PUT /rest/api/{2-3}/dashboard/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/dashboards#update-dashboard
func (d *DashboardService) Update(ctx context.Context, dashboardID string, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*DashboardService).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update_dashboard"),
		attribute.String("jira.dashboard.id", dashboardID),
	)

	result, response, err := d.internalClient.Update(ctx, dashboardID, payload)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

type internalDashboardImpl struct {
	c       service.Connector
	version string
}

func (i *internalDashboardImpl) Gets(ctx context.Context, startAt, maxResults int, filter string) (*model.DashboardPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalDashboardImpl).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_dashboards"),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
		attribute.String("jira.filter", filter),
		attribute.String("api.version", i.version),
	)

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(filter) != 0 {
		params.Add("filter", filter)
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	page := new(model.DashboardPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalDashboardImpl) Create(ctx context.Context, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalDashboardImpl).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "create_dashboard"),
		attribute.String("api.version", i.version),
	)

	endpoint := fmt.Sprintf("rest/api/%v/dashboard", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	dashboard := new(model.DashboardScheme)
	response, err := i.c.Call(request, dashboard)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return dashboard, response, nil
}

func (i *internalDashboardImpl) Search(ctx context.Context, options *model.DashboardSearchOptionsScheme, startAt, maxResults int) (*model.DashboardSearchPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalDashboardImpl).Search", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "search_dashboards"),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
		attribute.String("api.version", i.version),
	)

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
		recordError(span, err)
		return nil, nil, err
	}

	page := new(model.DashboardSearchPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalDashboardImpl) Get(ctx context.Context, dashboardID string) (*model.DashboardScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalDashboardImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_dashboard"),
		attribute.String("jira.dashboard.id", dashboardID),
		attribute.String("api.version", i.version),
	)

	if dashboardID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoDashboardID)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard/%v", i.version, dashboardID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	dashboard := new(model.DashboardScheme)
	response, err := i.c.Call(request, dashboard)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return dashboard, response, nil
}

func (i *internalDashboardImpl) Delete(ctx context.Context, dashboardID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalDashboardImpl).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete_dashboard"),
		attribute.String("jira.dashboard.id", dashboardID),
		attribute.String("api.version", i.version),
	)

	if dashboardID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoDashboardID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard/%v", i.version, dashboardID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

func (i *internalDashboardImpl) Copy(ctx context.Context, dashboardID string, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalDashboardImpl).Copy", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "copy_dashboard"),
		attribute.String("jira.dashboard.id", dashboardID),
		attribute.String("api.version", i.version),
	)

	if dashboardID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoDashboardID)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard/%v/copy", i.version, dashboardID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	dashboard := new(model.DashboardScheme)
	response, err := i.c.Call(request, dashboard)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return dashboard, response, nil
}

func (i *internalDashboardImpl) Update(ctx context.Context, dashboardID string, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalDashboardImpl).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update_dashboard"),
		attribute.String("jira.dashboard.id", dashboardID),
		attribute.String("api.version", i.version),
	)

	if dashboardID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoDashboardID)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/dashboard/%v", i.version, dashboardID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	dashboard := new(model.DashboardScheme)
	response, err := i.c.Call(request, dashboard)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return dashboard, response, nil
}
