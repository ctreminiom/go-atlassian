package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewFilterService creates a new instance of FilterService.
// It takes a service.Connector, a version string, and a jira.FilterSharingConnector as input.
// Returns a pointer to FilterService and an error if the version is not provided.
func NewFilterService(client service.Connector, version string, share jira.FilterSharingConnector) (*FilterService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &FilterService{
		internalClient: &internalFilterServiceImpl{c: client, version: version},
		Share:          share,
	}, nil
}

// FilterService provides methods to manage filters in Jira Service Management.
type FilterService struct {
	// internalClient is the connector interface for filter operations.
	internalClient jira.FilterConnector
	// Share is the service for managing filter sharing.
	Share jira.FilterSharingConnector
}

// Create creates a filter. The filter is shared according to the default share scope.
//
// The filter is not selected as a favorite.
//
// POST /rest/api/{2-3}/filter
//
// https://docs.go-atlassian.io/jira-software-cloud/filters#create-filter
func (f *FilterService) Create(ctx context.Context, payload *model.FilterPayloadScheme) (*model.FilterScheme, *model.ResponseScheme, error) {
	return f.internalClient.Create(ctx, payload)
}

// Favorite returns the visible favorite filters of the user.
//
// GET /rest/api/{2-3}/filter/favourite
//
// https://docs.go-atlassian.io/jira-software-cloud/filters#get-favorites
func (f *FilterService) Favorite(ctx context.Context) ([]*model.FilterScheme, *model.ResponseScheme, error) {
	return f.internalClient.Favorite(ctx)
}

// My returns the filters owned by the user. If includeFavourites is true,
//
// The user's visible favorite filters are also returned.
// GET /rest/api/{2-3}/filter/my
//
// https://docs.go-atlassian.io/jira-software-cloud/filters#get-my-filters
func (f *FilterService) My(ctx context.Context, favorites bool, expand []string) ([]*model.FilterScheme, *model.ResponseScheme, error) {
	return f.internalClient.My(ctx, favorites, expand)
}

// Search returns a paginated list of filters
//
// GET /rest/api/{2-3}/filter/search
//
// https://docs.go-atlassian.io/jira-software-cloud/filters#search-filters
func (f *FilterService) Search(ctx context.Context, options *model.FilterSearchOptionScheme, startAt, maxResults int) (*model.FilterSearchPageScheme,
	*model.ResponseScheme, error) {
	return f.internalClient.Search(ctx, options, startAt, maxResults)
}

// Get returns a filter.
//
// GET /rest/api/{2-3}/filter/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/filters#get-filter
func (f *FilterService) Get(ctx context.Context, filterID int, expand []string) (*model.FilterScheme, *model.ResponseScheme, error) {
	return f.internalClient.Get(ctx, filterID, expand)
}

// Update updates a filter. Use this operation to update a filter's name, description, JQL, or sharing.
//
// PUT /rest/api/{2-3}/filter/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/filters#update-filter
func (f *FilterService) Update(ctx context.Context, filterID int, payload *model.FilterPayloadScheme) (*model.FilterScheme, *model.ResponseScheme, error) {
	return f.internalClient.Update(ctx, filterID, payload)
}

// Delete a filter.
//
// DELETE /rest/api/{2-3}/filter/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/filters#delete-filter
func (f *FilterService) Delete(ctx context.Context, filterID int) (*model.ResponseScheme, error) {
	return f.internalClient.Delete(ctx, filterID)
}

// Change changes the owner of the filter.
//
// PUT /rest/api/{2-3}/filter/{id}/owner
//
// https://docs.go-atlassian.io/jira-software-cloud/filters#change-filter-owner
func (f *FilterService) Change(ctx context.Context, filterID int, accountID string) (*model.ResponseScheme, error) {
	return f.internalClient.Change(ctx, filterID, accountID)
}

type internalFilterServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalFilterServiceImpl) Create(ctx context.Context, payload *model.FilterPayloadScheme) (*model.FilterScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/filter", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	filter := new(model.FilterScheme)
	response, err := i.c.Call(request, filter)
	if err != nil {
		return nil, response, err
	}

	return filter, response, nil
}

func (i *internalFilterServiceImpl) Favorite(ctx context.Context) ([]*model.FilterScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/filter/favourite", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var filters []*model.FilterScheme
	response, err := i.c.Call(request, filters)
	if err != nil {
		return nil, response, err
	}

	return filters, response, nil
}

func (i *internalFilterServiceImpl) My(ctx context.Context, favorites bool, expand []string) ([]*model.FilterScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("includeFavourites", fmt.Sprintf("%v", favorites))

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/my?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var filters []*model.FilterScheme
	response, err := i.c.Call(request, filters)
	if err != nil {
		return nil, response, err
	}

	return filters, response, nil
}

func (i *internalFilterServiceImpl) Search(ctx context.Context, options *model.FilterSearchOptionScheme, startAt, maxResults int) (*model.FilterSearchPageScheme, *model.ResponseScheme, error) {

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

	endpoint := fmt.Sprintf("rest/api/%v/filter/search?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.FilterSearchPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalFilterServiceImpl) Get(ctx context.Context, filterID int, expand []string) (*model.FilterScheme, *model.ResponseScheme, error) {

	if filterID == 0 {
		return nil, nil, model.ErrNoFilterID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/filter/%v", i.version, filterID))

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	filter := new(model.FilterScheme)
	response, err := i.c.Call(request, filter)
	if err != nil {
		return nil, response, err
	}

	return filter, response, nil
}

func (i *internalFilterServiceImpl) Update(ctx context.Context, filterID int, payload *model.FilterPayloadScheme) (*model.FilterScheme, *model.ResponseScheme, error) {

	if filterID == 0 {
		return nil, nil, model.ErrNoFilterID
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v", i.version, filterID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	filter := new(model.FilterScheme)
	response, err := i.c.Call(request, filter)
	if err != nil {
		return nil, response, err
	}

	return filter, response, nil
}

func (i *internalFilterServiceImpl) Delete(ctx context.Context, filterID int) (*model.ResponseScheme, error) {

	if filterID == 0 {
		return nil, model.ErrNoFilterID
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v", i.version, filterID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		return response, err
	}

	return i.c.Call(request, nil)
}

func (i *internalFilterServiceImpl) Change(ctx context.Context, filterID int, accountID string) (*model.ResponseScheme, error) {

	if filterID == 0 {
		return nil, model.ErrNoFilterID
	}

	if accountID == "" {
		return nil, model.ErrNoAccountID
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v/owner", i.version, filterID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]interface{}{"accountId": accountID})
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
