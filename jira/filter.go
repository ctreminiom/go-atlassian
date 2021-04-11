package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type FilterService struct {
	client *Client
	Share  *FilterShareService
}

type FilterBodyScheme struct {
	Name             string                   `json:"name,omitempty"`
	Description      string                   `json:"description,omitempty"`
	JQL              string                   `json:"jql,omitempty"`
	Favorite         bool                     `json:"favourite,omitempty"`
	SharePermissions []*SharePermissionScheme `json:"sharePermissions,omitempty"`
}

// Creates a filter. The filter is shared according to the default share scope. The filter is not selected as a favorite.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#create-filter
func (f *FilterService) Create(ctx context.Context, payload *FilterBodyScheme) (result *FilterScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, payload value is nil, please provide a valid FilterBodyScheme pointer")
	}

	var endpoint = "rest/api/3/filter"
	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FilterScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Returns the visible favorite filters of the user.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#get-favorites
func (f *FilterService) Favorite(ctx context.Context) (result *[]FilterScheme, response *Response, err error) {

	var endpoint = "rest/api/3/filter/favourite"

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new([]FilterScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Returns the filters owned by the user. If includeFavourites is true, the user's visible favorite filters are also returned.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#get-my-filters
func (f *FilterService) My(ctx context.Context, favorites bool, expands []string) (result *[]FilterScheme, response *Response, err error) {

	params := url.Values{}

	var expand string
	for index, value := range expands {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	if favorites {
		params.Add("includeFavourites", "true")
	}

	var endpoint string
	if params.Encode() != "" {
		endpoint = fmt.Sprintf("rest/api/3/filter/my?%v", params.Encode())
	} else {
		endpoint = "rest/api/3/filter/my"
	}

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new([]FilterScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

type FilterSearchOptionScheme struct {
	Name, AccountID string
	Group, OrderBy  string
	ProjectID       int
	IDs             []int
	Expand          []string
}

// Returns a paginated list of filters. Use this operation to get:
// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#search-filters
func (f *FilterService) Search(ctx context.Context, options *FilterSearchOptionScheme, startAt, maxResults int) (result *FilterPageScheme, response *Response, err error) {

	if options == nil {
		return nil, nil, fmt.Errorf("error, options value is nil, please provide a valid FilterSearchOptionScheme pointer")
	}

	params := url.Values{}

	if options.Name != "" {
		params.Add("filterName", options.Name)
	}

	if options.AccountID != "" {
		params.Add("accountId", options.AccountID)
	}

	if options.Group != "" {
		params.Add("groupname", options.Group)
	}

	if options.ProjectID != 0 {
		params.Add("projectId", strconv.Itoa(options.ProjectID))
	}

	for _, filterID := range options.IDs {
		params.Add("id", strconv.Itoa(filterID))
	}

	if options.OrderBy != "" {
		params.Add("orderBy", options.OrderBy)
	}

	var expand string
	for index, value := range options.Expand {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("rest/api/3/filter/search?%v", params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FilterPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Returns a filter.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#get-filter
func (f *FilterService) Get(ctx context.Context, filterID int, expands []string) (result *FilterScheme, response *Response, err error) {

	params := url.Values{}

	var expand string
	for index, value := range expands {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	var endpoint string
	if params.Encode() != "" {
		endpoint = fmt.Sprintf("rest/api/3/filter/%v?%v", filterID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/api/3/filter/%v", filterID)
	}

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FilterScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Updates a filter. Use this operation to update a filter's name, description, JQL, or sharing.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#update-filter
func (f *FilterService) Update(ctx context.Context, filterID int, payload *FilterBodyScheme) (result *FilterScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, payload value is nil, please provide a valid FilterBodyScheme pointer")
	}

	var endpoint = fmt.Sprintf("rest/api/3/filter/%v", filterID)
	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FilterScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Delete a filter.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#delete-filter
func (f *FilterService) Delete(ctx context.Context, filterID int) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/filter/%v", filterID)
	request, err := f.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	return
}

type FilterPageScheme struct {
	Self       string                  `json:"self,omitempty"`
	MaxResults int                     `json:"maxResults,omitempty"`
	StartAt    int                     `json:"startAt,omitempty"`
	Total      int                     `json:"total,omitempty"`
	IsLast     bool                    `json:"isLast,omitempty"`
	Values     []*FilterPageNodeScheme `json:"values,omitempty"`
}

type FilterScheme struct {
	Self             string                        `json:"self,omitempty"`
	ID               string                        `json:"id,omitempty"`
	Name             string                        `json:"name,omitempty"`
	Owner            *UserScheme                   `json:"owner,omitempty"`
	Jql              string                        `json:"jql,omitempty"`
	ViewURL          string                        `json:"viewUrl,omitempty"`
	SearchURL        string                        `json:"searchUrl,omitempty"`
	Favourite        bool                          `json:"favourite,omitempty"`
	FavouritedCount  int                           `json:"favouritedCount,omitempty"`
	SharePermissions []*SharePermissionScheme      `json:"sharePermissions,omitempty"`
	ShareUsers       *FilterUsersScheme            `json:"sharedUsers,omitempty"`
	Subscriptions    *FilterSubscriptionPageScheme `json:"subscriptions,omitempty"`
}

type FilterPageNodeScheme struct {
	Self             string                   `json:"self,omitempty"`
	ID               string                   `json:"id,omitempty"`
	Name             string                   `json:"name,omitempty"`
	Owner            *UserScheme              `json:"owner,omitempty"`
	Jql              string                   `json:"jql,omitempty"`
	ViewURL          string                   `json:"viewUrl,omitempty"`
	SearchURL        string                   `json:"searchUrl,omitempty"`
	Favourite        bool                     `json:"favourite,omitempty"`
	FavouritedCount  int                      `json:"favouritedCount,omitempty"`
	SharePermissions []*SharePermissionScheme `json:"sharePermissions,omitempty"`
	ShareUsers       *FilterUsersScheme       `json:"sharedUsers,omitempty"`
	Subscriptions    *[]FilterUsersScheme     `json:"subscriptions,omitempty"`
}

type FilterSubscriptionPageScheme struct {
	Size       int                         `json:"size,omitempty"`
	Items      []*FilterSubscriptionScheme `json:"items,omitempty"`
	MaxResults int                         `json:"max-results,omitempty"`
	StartIndex int                         `json:"start-index,omitempty"`
	EndIndex   int                         `json:"end-index,omitempty"`
}

type FilterSubscriptionScheme struct {
	ID    int          `json:"id,omitempty"`
	User  *UserScheme  `json:"user,omitempty"`
	Group *GroupScheme `json:"group,omitempty"`
}

type FilterUsersScheme struct {
	Size       int           `json:"size,omitempty"`
	Items      []*UserScheme `json:"items,omitempty"`
	MaxResults int           `json:"max-results,omitempty"`
	StartIndex int           `json:"start-index,omitempty"`
	EndIndex   int           `json:"end-index,omitempty"`
}
