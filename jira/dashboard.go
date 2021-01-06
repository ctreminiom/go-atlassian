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
func (d *DashboardService) Gets(ctx context.Context, startAt, maxResults int, filter string) (dashboards *DashboardsSchemeResult, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	validateURLParam(&params, "filter", filter)

	var endpoint = fmt.Sprintf("rest/api/3/dashboard?%s", params.Encode())
	request, err := d.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = d.client.Do(request)
	if err != nil {
		return
	}

	result := new(DashboardsSchemeResult)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

/*
func find(listAsString []string, value string) (isAvailable bool) {

	sort.Strings(listAsString)

	//Make the binary string using a custom sort algorithm
	binarySearchIndex := sort.Search(len(listAsString), func(position int) bool { return value <= listAsString[position] })

	if binarySearchIndex < len(listAsString) && listAsString[binarySearchIndex] == value {
		return true
	}

	return false
}
*/
