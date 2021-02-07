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
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	JQL         string `json:"jql,omitempty"`
	Favorite    bool   `json:"favourite,omitempty"`
}

// Creates a filter. The filter is shared according to the default share scope. The filter is not selected as a favorite.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filters/#api-rest-api-3-filter-post
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
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filters/#api-rest-api-3-filter-favourite-get
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
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filters/#api-rest-api-3-filter-my-get
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
	Name      string
	AccountID string
	Group     string
	ProjectID int
	IDs       []int
	OrderBy   string
	Expand    []string
}

// Returns a paginated list of filters. Use this operation to get:
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filters/#api-rest-api-3-filter-search-get
func (f *FilterService) Search(ctx context.Context, options *FilterSearchOptionScheme, startAt, maxResults int) (result *FilterSearchScheme, response *Response, err error) {

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

	result = new(FilterSearchScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Returns a filter.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filters/#api-rest-api-3-filter-id-get
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
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filters/#api-rest-api-3-filter-id-put
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
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filters/#api-rest-api-3-filter-id-delete
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

type FilterSearchScheme struct {
	Self       string `json:"self"`
	MaxResults int    `json:"maxResults"`
	StartAt    int    `json:"startAt"`
	Total      int    `json:"total"`
	IsLast     bool   `json:"isLast"`
	Values     []struct {
		Self        string `json:"self"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		Owner       struct {
			Self       string `json:"self"`
			AccountID  string `json:"accountId"`
			AvatarUrls struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
		} `json:"owner"`
		Jql              string        `json:"jql"`
		ViewURL          string        `json:"viewUrl"`
		SearchURL        string        `json:"searchUrl"`
		Favourite        bool          `json:"favourite"`
		FavouritedCount  int           `json:"favouritedCount"`
		SharePermissions []interface{} `json:"sharePermissions"`
		Subscriptions    []interface{} `json:"subscriptions"`
	} `json:"values"`
}

type FilterScheme struct {
	Self  string `json:"self,omitempty"`
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Owner struct {
		Self       string `json:"self,omitempty"`
		AccountID  string `json:"accountId,omitempty"`
		AvatarUrls struct {
			Four8X48  string `json:"48x48,omitempty"`
			Two4X24   string `json:"24x24,omitempty"`
			One6X16   string `json:"16x16,omitempty"`
			Three2X32 string `json:"32x32,omitempty"`
		} `json:"avatarUrls,omitempty"`
		DisplayName string `json:"displayName,omitempty"`
		Active      bool   `json:"active,omitempty"`
	} `json:"owner,omitempty"`
	Jql              string `json:"jql,omitempty"`
	ViewURL          string `json:"viewUrl,omitempty"`
	SearchURL        string `json:"searchUrl,omitempty"`
	Favourite        bool   `json:"favourite,omitempty"`
	FavouritedCount  int    `json:"favouritedCount,omitempty"`
	SharePermissions []struct {
		ID      int    `json:"id,omitempty"`
		Type    string `json:"type,omitempty"`
		Project struct {
			Self         string `json:"self,omitempty"`
			ID           string `json:"id,omitempty"`
			Key          string `json:"key,omitempty"`
			AssigneeType string `json:"assigneeType,omitempty"`
			Name         string `json:"name,omitempty"`
			Roles        struct {
			} `json:"roles,omitempty"`
			AvatarUrls struct {
				Four8X48  string `json:"48x48,omitempty"`
				Two4X24   string `json:"24x24,omitempty"`
				One6X16   string `json:"16x16,omitempty"`
				Three2X32 string `json:"32x32,omitempty"`
			} `json:"avatarUrls,omitempty"`
			ProjectTypeKey string `json:"projectTypeKey,omitempty"`
			Simplified     bool   `json:"simplified,omitempty"`
			Style          string `json:"style,omitempty"`
			Properties     struct {
			} `json:"properties,omitempty"`
		} `json:"project,omitempty"`
	} `json:"sharePermissions,omitempty"`
	SharedUsers struct {
		Size  int `json:"size,omitempty"`
		Items []struct {
			Self       string `json:"self,omitempty"`
			AccountID  string `json:"accountId,omitempty"`
			AvatarUrls struct {
				Four8X48  string `json:"48x48,omitempty"`
				Two4X24   string `json:"24x24,omitempty"`
				One6X16   string `json:"16x16,omitempty"`
				Three2X32 string `json:"32x32,omitempty"`
			} `json:"avatarUrls,omitempty"`
			DisplayName string `json:"displayName,omitempty"`
			Active      bool   `json:"active,omitempty"`
		} `json:"items,omitempty"`
		MaxResults int `json:"max-results,omitempty"`
		StartIndex int `json:"start-index,omitempty"`
		EndIndex   int `json:"end-index,omitempty"`
	} `json:"sharedUsers,omitempty"`
	Subscriptions struct {
		Size  int `json:"size,omitempty"`
		Items []struct {
			ID   int `json:"id,omitempty"`
			User struct {
				Self       string `json:"self,omitempty"`
				AccountID  string `json:"accountId,omitempty"`
				AvatarUrls struct {
					Four8X48  string `json:"48x48,omitempty"`
					Two4X24   string `json:"24x24,omitempty"`
					One6X16   string `json:"16x16,omitempty"`
					Three2X32 string `json:"32x32,omitempty"`
				} `json:"avatarUrls,omitempty"`
				DisplayName string `json:"displayName,omitempty"`
				Active      bool   `json:"active,omitempty"`
			} `json:"user,omitempty"`
		} `json:"items,omitempty"`
		MaxResults int `json:"max-results,omitempty"`
		StartIndex int `json:"start-index,omitempty"`
		EndIndex   int `json:"end-index,omitempty"`
	} `json:"subscriptions,omitempty"`
}
