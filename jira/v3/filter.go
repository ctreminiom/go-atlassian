package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type FilterService struct {
	client *Client
	Share  *FilterShareService
}

type FilterPayloadScheme struct {
	Name             string                          `json:"name,omitempty"`
	Description      string                          `json:"description,omitempty"`
	JQL              string                          `json:"jql,omitempty"`
	Favorite         bool                            `json:"favourite,omitempty"`
	SharePermissions []*models.SharePermissionScheme `json:"sharePermissions,omitempty"`
}

// Create creates a filter. The filter is shared according to the default share scope.
// The filter is not selected as a favorite.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#create-filter
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filters/#api-rest-api-3-filter-post
func (f *FilterService) Create(ctx context.Context, payload *FilterPayloadScheme) (result *models.FilterScheme,
	response *ResponseScheme, err error) {

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint = "rest/api/3/filter"
	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Favorite returns the visible favorite filters of the user.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#get-favorites
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filters/#api-rest-api-3-filter-favourite-get
func (f *FilterService) Favorite(ctx context.Context) (result []*models.FilterScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/filter/favourite"

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// My returns the filters owned by the user. If includeFavourites is true,
// the user's visible favorite filters are also returned.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#get-my-filters
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filters/#api-rest-api-3-filter-my-get
func (f *FilterService) My(ctx context.Context, favorites bool, expand []string) (result []*models.FilterScheme, response *ResponseScheme, err error) {

	params := url.Values{}

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	if favorites {
		params.Add("includeFavourites", "true")
	}

	var endpoint strings.Builder
	endpoint.WriteString("rest/api/3/filter/my")

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
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

// Search returns a paginated list of filters
// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#search-filters
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filters/#api-rest-api-3-filter-search-get
func (f *FilterService) Search(ctx context.Context, options *FilterSearchOptionScheme, startAt, maxResults int) (
	result *models.FilterSearchPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}

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
	}

	var endpoint = fmt.Sprintf("rest/api/3/filter/search?%v", params.Encode())

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}
	return
}

// Get returns a filter.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#get-filter
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filters/#api-rest-api-3-filter-id-get
func (f *FilterService) Get(ctx context.Context, filterID int, expand []string) (result *models.FilterScheme,
	response *ResponseScheme, err error) {

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpointBuffer strings.Builder
	endpointBuffer.WriteString(fmt.Sprintf("rest/api/3/filter/%v", filterID))

	if params.Encode() != "" {
		endpointBuffer.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := f.client.newRequest(ctx, http.MethodGet, endpointBuffer.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update updates a filter. Use this operation to update a filter's name, description, JQL, or sharing.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#update-filter
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filters/#api-rest-api-3-filter-id-put
func (f *FilterService) Update(ctx context.Context, filterID int, payload *FilterPayloadScheme) (result *models.FilterScheme,
	response *ResponseScheme, err error) {
	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint = fmt.Sprintf("rest/api/3/filter/%v", filterID)

	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete a filter.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/filters#delete-filter
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filters/#api-rest-api-3-filter-id-delete
func (f *FilterService) Delete(ctx context.Context, filterID int) (response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/filter/%v", filterID)

	request, err := f.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = f.client.call(request, nil)
	if err != nil {
		return
	}

	return
}
