package jira

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type DashboardService struct{ client *Client }

type DashboardsSchemeResult struct {
	StartAt    int               `json:"startAt,omitempty"`
	MaxResults int               `json:"maxResults,omitempty"`
	Total      int               `json:"total,omitempty"`
	Dashboards []DashboardScheme `json:"dashboards,omitempty"`
}
type SharePermissionsScheme struct {
	ID   int    `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
}
type DashboardScheme struct {
	ID               string                   `json:"id,omitempty"`
	IsFavourite      bool                     `json:"isFavourite,omitempty"`
	Name             string                   `json:"name,omitempty"`
	Popularity       int                      `json:"popularity,omitempty"`
	Self             string                   `json:"self,omitempty"`
	SharePermissions []SharePermissionsScheme `json:"sharePermissions,omitempty"`
	View             string                   `json:"view,omitempty"`
}

// Returns a list of dashboards owned by or shared with the user. The list may be filtered to include only favorite or owned dashboards.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-dashboards/#api-rest-api-3-dashboard-get
func (d *DashboardService) Gets(ctx context.Context, startAt, maxResults int, filter string) (result *DashboardsSchemeResult, response *Response, err error) {

	if ctx == nil {
		return nil, nil, errors.New("the context param is nil, please provide a valid one")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(filter) != 0 {
		params.Add("filter", filter)
	}

	var endpoint = fmt.Sprintf("rest/api/3/dashboard?%s", params.Encode())
	request, err := d.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = d.client.Do(request)
	if err != nil {
		return
	}

	if len(response.BodyAsBytes) == 0 {
		return nil, nil, errors.New("unable to marshall the response body, the HTTP callback did not return any bytes")
	}

	result = new(DashboardsSchemeResult)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
