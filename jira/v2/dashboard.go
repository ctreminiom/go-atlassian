package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type DashboardService struct{ client *Client }

// Gets returns a list of dashboards owned by or shared with the user. The list may be filtered to include only favorite or owned dashboards.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#get-all-dashboards
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-dashboards/#api-rest-api-2-dashboard-get
func (d *DashboardService) Gets(ctx context.Context, startAt, maxResults int, filter string) (result *models.DashboardPageScheme,
	response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(filter) != 0 {
		params.Add("filter", filter)
	}

	var endpoint = fmt.Sprintf("rest/api/2/dashboard?%v", params.Encode())
	request, err := d.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = d.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type DashboardPayloadScheme struct {
	Name             string                          `json:"name,omitempty"`
	Description      string                          `json:"description,omitempty"`
	SharePermissions []*models.SharePermissionScheme `json:"sharePermissions,omitempty"`
}

// Create creates a dashboard.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#create-dashboard
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-dashboards/#api-rest-api-2-dashboard-post
func (d *DashboardService) Create(ctx context.Context, payload *DashboardPayloadScheme) (result *models.DashboardScheme,
	response *ResponseScheme, err error) {

	var endpoint = "rest/api/2/dashboard"

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := d.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = d.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type DashboardSearchOptionsScheme struct {
	DashboardName, OwnerAccountID string
	GroupPermissionName, OrderBy  string
	Expand                        []string
}

// Search returns a paginated list of dashboards.
// This operation is similar to Get dashboards except that the results can be refined to include dashboards that have specific attributes.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#search-for-dashboards
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-dashboards/#api-rest-api-2-dashboard-search-get
func (d *DashboardService) Search(ctx context.Context, opts *DashboardSearchOptionsScheme, startAt, maxResults int) (
	result *models.DashboardSearchPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if len(opts.OwnerAccountID) != 0 {
			params.Add("accountId", opts.OwnerAccountID)
		}

		if len(opts.DashboardName) != 0 {
			params.Add("dashboardName", opts.OwnerAccountID)
		}

		if len(opts.GroupPermissionName) != 0 {
			params.Add("groupname", opts.OwnerAccountID)
		}

		if len(opts.OrderBy) != 0 {
			params.Add("orderBy", opts.OwnerAccountID)
		}

		if len(opts.Expand) != 0 {
			params.Add("expand", strings.Join(opts.Expand, ","))
		}
	}

	var endpoint = fmt.Sprintf("rest/api/2/dashboard/search?%s", params.Encode())
	request, err := d.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = d.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns a dashboard.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#get-dashboard
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-dashboards/#api-rest-api-2-dashboard-id-get
func (d *DashboardService) Get(ctx context.Context, dashboardID string) (result *models.DashboardScheme,
	response *ResponseScheme, err error) {

	if len(dashboardID) == 0 {
		return nil, nil, models.ErrNoDashboardIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/dashboard/%v", dashboardID)

	request, err := d.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = d.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes a dashboard.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#delete-dashboard
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-dashboards/#api-rest-api-2-dashboard-id-delete
func (d *DashboardService) Delete(ctx context.Context, dashboardID string) (response *ResponseScheme, err error) {

	if len(dashboardID) == 0 {
		return nil, models.ErrNoDashboardIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/dashboard/%v", dashboardID)
	request, err := d.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = d.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Copy copies a dashboard.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#copy-dashboard
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-dashboards/#api-rest-api-2-dashboard-id-copy-post
func (d *DashboardService) Copy(ctx context.Context, dashboardID string, payload *DashboardPayloadScheme) (
	result *models.DashboardScheme, response *ResponseScheme, err error) {

	if len(dashboardID) == 0 {
		return nil, nil, models.ErrNoDashboardIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/dashboard/%v/copy", dashboardID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := d.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = d.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update updates a dashboard
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#update-dashboard
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-dashboards/#api-rest-api-2-dashboard-id-put
func (d *DashboardService) Update(ctx context.Context, dashboardID string, payload *DashboardPayloadScheme) (result *models.DashboardScheme,
	response *ResponseScheme, err error) {

	if len(dashboardID) == 0 {
		return nil, nil, models.ErrNoDashboardIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/dashboard/%v", dashboardID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := d.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = d.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
