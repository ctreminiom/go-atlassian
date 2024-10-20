package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type DashboardConnector interface {

	// Gets returns a list of dashboards owned by or shared with the user.
	//
	// The list may be filtered to include only favorite or owned dashboards.
	//
	// GET /rest/api/{3-2}/dashboard
	//
	// https://docs.go-atlassian.io/jira-software-cloud/dashboards#get-all-dashboards
	Gets(ctx context.Context, startAt, maxResults int, filter string) (*model.DashboardPageScheme, *model.ResponseScheme, error)

	// Create creates a dashboard.
	//
	// POST /rest/api/{3-2}/dashboard
	//
	// https://docs.go-atlassian.io/jira-software-cloud/dashboards#create-dashboard
	Create(ctx context.Context, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error)

	// Search returns a paginated list of dashboards.
	//
	// This operation is similar to Get dashboards except that the results can be refined to include dashboards that have specific attributes.
	//
	// GET /rest/api/{2-3}/dashboard/search
	//
	// https://docs.go-atlassian.io/jira-software-cloud/dashboards#search-for-dashboards
	Search(ctx context.Context, options *model.DashboardSearchOptionsScheme, startAt, maxResults int) (*model.DashboardSearchPageScheme, *model.ResponseScheme, error)

	// Get returns a dashboard.
	//
	// GET /rest/api/{2-3}/dashboard/{dashboardID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/dashboards#get-dashboard
	Get(ctx context.Context, dashboardID string) (*model.DashboardScheme, *model.ResponseScheme, error)

	// Delete deletes a dashboard.
	//
	// DELETE /rest/api/{2-3}/dashboard/{dashboardID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/dashboards#delete-dashboard
	Delete(ctx context.Context, dashboardID string) (*model.ResponseScheme, error)

	// Copy copies a dashboard.
	//
	// Any values provided in the dashboard parameter replace those in the copied dashboard.
	//
	// POST /rest/api/{2-3}/dashboard/{dashboardID}/copy
	//
	// https://docs.go-atlassian.io/jira-software-cloud/dashboards#copy-dashboard
	Copy(ctx context.Context, dashboardID string, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error)

	// Update updates a dashboard
	//
	// PUT /rest/api/{2-3}/dashboard/{dashboardID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/dashboards#update-dashboard
	Update(ctx context.Context, dashboardID string, payload *model.DashboardPayloadScheme) (*model.DashboardScheme, *model.ResponseScheme, error)
}
