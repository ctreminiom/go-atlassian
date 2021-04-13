package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
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
	Popularity       int                      `json:"popularity,omitempty"`
	Self             string                   `json:"self,omitempty"`
	SharePermissions []*SharePermissionScheme `json:"sharePermissions,omitempty"`
	View             string                   `json:"view,omitempty"`
}

// Returns a list of dashboards owned by or shared with the user. The list may be filtered to include only favorite or owned dashboards.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#get-all-dashboards
func (d *DashboardService) Gets(ctx context.Context, startAt, maxResults int, filter string) (result *DashboardPageScheme, response *Response, err error) {

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

	response, err = d.client.Do(request)
	if err != nil {
		return
	}

	result = new(DashboardPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
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

// Creates a dashboard.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#create-dashboard
func (d *DashboardService) Create(ctx context.Context, name, description string, permissions *[]SharePermissionScheme) (result *DashboardScheme, response *Response, err error) {

	if len(name) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid Dashboard name")
	}

	if permissions == nil {
		return nil, nil, fmt.Errorf("error, please provide the permission to user, this param is required by the Atlassian Documentation")
	}

	if len(*permissions) == 0 {
		return nil, nil, fmt.Errorf("error, please provide the permission to user, this param is required by the Atlassian Documentation")
	}

	//Create the payload
	var payload = map[string]interface{}{}
	payload["name"] = name
	payload["description"] = description

	//For each share permission, append the object node
	var sharePermissionsNode []map[string]interface{}
	for _, permission := range *permissions {

		//Convert the SharePermissionScheme struct to map[string]interface{}
		sharePermissionSchemeAsBytes, err := json.Marshal(permission)
		if err != nil {
			return nil, nil, err
		}

		sharePermissionSchemeAsMap := make(map[string]interface{})
		err = json.Unmarshal(sharePermissionSchemeAsBytes, &sharePermissionSchemeAsMap)
		if err != nil {
			return nil, nil, err
		}

		sharePermissionsNode = append(sharePermissionsNode, sharePermissionSchemeAsMap)
	}
	payload["sharePermissions"] = sharePermissionsNode

	var endpoint = "rest/api/3/dashboard"

	request, err := d.client.newRequest(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = d.client.Do(request)
	if err != nil {
		return
	}

	result = new(DashboardScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
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

type DashboardSearchScheme struct {
	Self       string `json:"self"`
	MaxResults int    `json:"maxResults"`
	StartAt    int    `json:"startAt"`
	Total      int    `json:"total"`
	IsLast     bool   `json:"isLast"`
	Values     []struct {
		Description      string                   `json:"description,omitempty"`
		ID               string                   `json:"id"`
		IsFavourite      bool                     `json:"isFavourite"`
		Name             string                   `json:"name"`
		Owner            *UserScheme              `json:"owner,omitempty"`
		Popularity       int                      `json:"popularity"`
		Self             string                   `json:"self"`
		SharePermissions []*SharePermissionScheme `json:"sharePermissions"`
		View             string                   `json:"view"`
		Rank             int                      `json:"rank"`
	} `json:"values"`
}

// Returns a paginated list of dashboards.
// This operation is similar to Get dashboards except that the results can be refined to include dashboards that have specific attributes.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#search-for-dashboards
func (d *DashboardService) Search(ctx context.Context, opts *DashboardSearchOptionsScheme, startAt, maxResults int) (result *DashboardSearchScheme, response *Response, err error) {

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

		var expand string
		for index, value := range opts.Expand {

			if index == 0 {
				expand = value
				continue
			}

			expand += "," + value
		}

		if len(expand) != 0 {
			params.Add("expand", expand)
		}
	}

	var endpoint = fmt.Sprintf("rest/api/3/dashboard/search?%s", params.Encode())
	request, err := d.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = d.client.Do(request)
	if err != nil {
		return
	}

	result = new(DashboardSearchScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns a dashboard.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#get-dashboard
func (d *DashboardService) Get(ctx context.Context, dashboardID string) (result *DashboardScheme, response *Response, err error) {

	if len(dashboardID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid dashboardID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/dashboard/%v", dashboardID)
	request, err := d.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = d.client.Do(request)
	if err != nil {
		return
	}

	result = new(DashboardScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes a dashboard.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#delete-dashboard
func (d *DashboardService) Delete(ctx context.Context, dashboardID string) (response *Response, err error) {

	if len(dashboardID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid dashboardID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/dashboard/%v", dashboardID)
	request, err := d.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = d.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Copies a dashboard.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#copy-dashboard
func (d *DashboardService) Copy(ctx context.Context, dashboardID, newDashboardName, newDashboardDescription string, permissions *[]SharePermissionScheme) (result *DashboardScheme, response *Response, err error) {

	if len(dashboardID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid dashboardID value")
	}

	if len(newDashboardName) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid newDashboardName value")
	}

	if permissions == nil {
		return nil, nil, fmt.Errorf("error, please provide the permission to user, this param is required by the Atlassian Documentation")
	}

	if len(*permissions) == 0 {
		return nil, nil, fmt.Errorf("error, please provide the permission to user, this param is required by the Atlassian Documentation")
	}

	//Create the payload
	var payload = map[string]interface{}{}
	payload["name"] = newDashboardName
	payload["description"] = newDashboardDescription

	//For each share permission, append the object node
	var sharePermissionsNode []map[string]interface{}
	for _, permission := range *permissions {

		//Convert the SharePermissionScheme struct to map[string]interface{}
		sharePermissionSchemeAsBytes, err := json.Marshal(permission)
		if err != nil {
			return nil, nil, err
		}

		sharePermissionSchemeAsMap := make(map[string]interface{})
		err = json.Unmarshal(sharePermissionSchemeAsBytes, &sharePermissionSchemeAsMap)
		if err != nil {
			return nil, nil, err
		}

		sharePermissionsNode = append(sharePermissionsNode, sharePermissionSchemeAsMap)
	}
	payload["sharePermissions"] = sharePermissionsNode

	var endpoint = fmt.Sprintf("rest/api/3/dashboard/%v/copy", dashboardID)

	request, err := d.client.newRequest(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = d.client.Do(request)
	if err != nil {
		return
	}

	result = new(DashboardScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Updates a dashboard
// Docs: https://docs.go-atlassian.io/jira-software-cloud/dashboards#update-dashboard
func (d *DashboardService) Update(ctx context.Context, dashboardID, newDashboardName, newDashboardDescription string, permissions *[]SharePermissionScheme) (result *DashboardScheme, response *Response, err error) {

	if len(dashboardID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid dashboardID value")
	}

	if permissions == nil {
		return nil, nil, fmt.Errorf("error, please provide the permission to user, this param is required by the Atlassian Documentation")
	}

	if len(*permissions) == 0 {
		return nil, nil, fmt.Errorf("error, please provide the permission to user, this param is required by the Atlassian Documentation")
	}

	//Create the payload
	var payload = map[string]interface{}{}
	payload["name"] = newDashboardName

	if len(newDashboardDescription) != 0 {
		payload["description"] = newDashboardDescription
	}

	//For each share permission, append the object node
	var sharePermissionsNode []map[string]interface{}
	for _, permission := range *permissions {

		//Convert the SharePermissionScheme struct to map[string]interface{}
		sharePermissionSchemeAsBytes, err := json.Marshal(permission)
		if err != nil {
			return nil, nil, err
		}

		sharePermissionSchemeAsMap := make(map[string]interface{})
		err = json.Unmarshal(sharePermissionSchemeAsBytes, &sharePermissionSchemeAsMap)
		if err != nil {
			return nil, nil, err
		}

		sharePermissionsNode = append(sharePermissionsNode, sharePermissionSchemeAsMap)
	}
	payload["sharePermissions"] = sharePermissionsNode

	var endpoint = fmt.Sprintf("rest/api/3/dashboard/%v", dashboardID)

	request, err := d.client.newRequest(ctx, http.MethodPut, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = d.client.Do(request)
	if err != nil {
		return
	}

	result = new(DashboardScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
