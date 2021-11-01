package v3

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type DashboardService struct{ client *Client }

type DashboardPageScheme struct {
	StartAt    int                `json:"startAt,omitempty"`
	MaxResults int                `json:"maxResults,omitempty"`
	Total      int                `json:"total,omitempty"`
	Dashboards []*DashboardScheme `json:"dashboards,omitempty"`
}

type DashboardScheme struct {
	ID               string                   `json:"id,omitempty"`
	IsFavourite      bool                     `json:"isFavourite,omitempty"`
	Name             string                   `json:"name,omitempty"`
	Owner            *UserScheme              `json:"owner,omitempty"`
	Popularity       int                      `json:"popularity,omitempty"`
	Rank             int                      `json:"rank,omitempty"`
	Self             string                   `json:"self,omitempty"`
	SharePermissions []*SharePermissionScheme `json:"sharePermissions,omitempty"`
	EditPermission   []*SharePermissionScheme `json:"editPermissions,omitempty"`
	View             string                   `json:"view,omitempty"`
}

// Gets returns a list of dashboards owned by or shared with the user. The list may be filtered to include only favorite or owned dashboards.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#get-all-dashboards
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-dashboards/#api-rest-api-3-dashboard-get
func (d *DashboardService) Gets(ctx context.Context, startAt, maxResults int, filter string) (result *DashboardPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(filter) != 0 {
		params.Add("filter", filter)
	}

	var endpoint = fmt.Sprintf("rest/api/3/dashboard?%v", params.Encode())
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

type SharePermissionScheme struct {
	ID      int                `json:"id,omitempty"`
	Type    string             `json:"type,omitempty"`
	Project *ProjectScheme     `json:"project,omitempty"`
	Role    *ProjectRoleScheme `json:"role,omitempty"`
	Group   *GroupScheme       `json:"group,omitempty"`
}

type DashboardPayloadScheme struct {
	Name             string                   `json:"name,omitempty"`
	Description      string                   `json:"description,omitempty"`
	SharePermissions []*SharePermissionScheme `json:"sharePermissions,omitempty"`
}

// Create creates a dashboard.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#create-dashboard
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-dashboards/#api-rest-api-3-dashboard-post
// NOTE: Experimental Endpoint
func (d *DashboardService) Create(ctx context.Context, payload *DashboardPayloadScheme) (result *DashboardScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/dashboard"

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
	DashboardName       string
	OwnerAccountID      string
	GroupPermissionName string
	OrderBy             string
	Expand              []string
}

type DashboardSearchPageScheme struct {
	Self       string             `json:"self,omitempty"`
	MaxResults int                `json:"maxResults,omitempty"`
	StartAt    int                `json:"startAt,omitempty"`
	Total      int                `json:"total,omitempty"`
	IsLast     bool               `json:"isLast,omitempty"`
	Values     []*DashboardScheme `json:"values,omitempty"`
}

// Search returns a paginated list of dashboards.
// This operation is similar to Get dashboards except that the results can be refined to include dashboards that have specific attributes.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#search-for-dashboards
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-dashboards/#api-rest-api-3-dashboard-search-get
func (d *DashboardService) Search(ctx context.Context, opts *DashboardSearchOptionsScheme, startAt, maxResults int) (
	result *DashboardSearchPageScheme, response *ResponseScheme, err error) {

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

	var endpoint = fmt.Sprintf("rest/api/3/dashboard/search?%s", params.Encode())
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
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-dashboards/#api-rest-api-3-dashboard-id-get
func (d *DashboardService) Get(ctx context.Context, dashboardID string) (result *DashboardScheme,
	response *ResponseScheme, err error) {

	if len(dashboardID) == 0 {
		return nil, nil, notDashboardIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/dashboard/%v", dashboardID)

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
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-dashboards/#api-rest-api-3-dashboard-id-delete
// NOTE: Experimental Method
func (d *DashboardService) Delete(ctx context.Context, dashboardID string) (response *ResponseScheme, err error) {

	if len(dashboardID) == 0 {
		return nil, notDashboardIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/dashboard/%v", dashboardID)
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
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-dashboards/#api-rest-api-3-dashboard-id-copy-post
// Note: Experimental Method
func (d *DashboardService) Copy(ctx context.Context, dashboardID string, payload *DashboardPayloadScheme) (
	result *DashboardScheme, response *ResponseScheme, err error) {

	if len(dashboardID) == 0 {
		return nil, nil, notDashboardIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/dashboard/%v/copy", dashboardID)

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
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-dashboards/#api-rest-api-3-dashboard-id-put
// NOTE: Experimental Method
func (d *DashboardService) Update(ctx context.Context, dashboardID string, payload *DashboardPayloadScheme) (result *DashboardScheme,
	response *ResponseScheme, err error) {

	if len(dashboardID) == 0 {
		return nil, nil, notDashboardIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/dashboard/%v", dashboardID)

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

var (
	notDashboardIDError = fmt.Errorf("error, please provide a valid dashboardID value")
)
