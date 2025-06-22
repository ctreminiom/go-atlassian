package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewFilterService creates a new instance of FilterService.
// It takes a service.Connector, a version string, and a jira.FilterSharingConnector as input.
// Returns a pointer to FilterService and an error if the version is not provided.
func NewFilterService(client service.Connector, version string, share jira.FilterSharingConnector) (*FilterService, error) {

	if version == "" {
		return nil, fmt.Errorf("jira: %w", model.ErrNoVersionProvided)
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
	ctx, span := tracer().Start(ctx, "(*FilterService).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "create_filter"),
	)

	result, response, err := f.internalClient.Create(ctx, payload)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Favorite returns the visible favorite filters of the user.
//
// GET /rest/api/{2-3}/filter/favourite
//
// https://docs.go-atlassian.io/jira-software-cloud/filters#get-favorites
func (f *FilterService) Favorite(ctx context.Context) ([]*model.FilterScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*FilterService).Favorite", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_favorite_filters"),
	)

	result, response, err := f.internalClient.Favorite(ctx)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// My returns the filters owned by the user. If includeFavourites is true,
//
// The user's visible favorite filters are also returned.
// GET /rest/api/{2-3}/filter/my
//
// https://docs.go-atlassian.io/jira-software-cloud/filters#get-my-filters
func (f *FilterService) My(ctx context.Context, favorites bool, expand []string) ([]*model.FilterScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*FilterService).My", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_my_filters"),
		attribute.Bool("jira.filter.include_favorites", favorites),
		attribute.Int("jira.expand.count", len(expand)),
	)

	result, response, err := f.internalClient.My(ctx, favorites, expand)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Search returns a paginated list of filters
//
// GET /rest/api/{2-3}/filter/search
//
// https://docs.go-atlassian.io/jira-software-cloud/filters#search-filters
func (f *FilterService) Search(ctx context.Context, options *model.FilterSearchOptionScheme, startAt, maxResults int) (*model.FilterSearchPageScheme,
	*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*FilterService).Search", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "search_filters"),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
	)

	result, response, err := f.internalClient.Search(ctx, options, startAt, maxResults)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Get returns a filter.
//
// GET /rest/api/{2-3}/filter/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/filters#get-filter
func (f *FilterService) Get(ctx context.Context, filterID int, expand []string) (*model.FilterScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*FilterService).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_filter"),
		attribute.Int("jira.filter.id", filterID),
		attribute.Int("jira.expand.count", len(expand)),
	)

	result, response, err := f.internalClient.Get(ctx, filterID, expand)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Update updates a filter. Use this operation to update a filter's name, description, JQL, or sharing.
//
// PUT /rest/api/{2-3}/filter/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/filters#update-filter
func (f *FilterService) Update(ctx context.Context, filterID int, payload *model.FilterPayloadScheme) (*model.FilterScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*FilterService).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update_filter"),
		attribute.Int("jira.filter.id", filterID),
	)

	result, response, err := f.internalClient.Update(ctx, filterID, payload)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Delete a filter.
//
// DELETE /rest/api/{2-3}/filter/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/filters#delete-filter
func (f *FilterService) Delete(ctx context.Context, filterID int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*FilterService).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete_filter"),
		attribute.Int("jira.filter.id", filterID),
	)

	response, err := f.internalClient.Delete(ctx, filterID)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

// Change changes the owner of the filter.
//
// PUT /rest/api/{2-3}/filter/{id}/owner
//
// https://docs.go-atlassian.io/jira-software-cloud/filters#change-filter-owner
func (f *FilterService) Change(ctx context.Context, filterID int, accountID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*FilterService).Change", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "change_filter_owner"),
		attribute.Int("jira.filter.id", filterID),
		attribute.String("jira.account.id", accountID),
	)

	response, err := f.internalClient.Change(ctx, filterID, accountID)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

type internalFilterServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalFilterServiceImpl) Create(ctx context.Context, payload *model.FilterPayloadScheme) (*model.FilterScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalFilterServiceImpl).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	endpoint := fmt.Sprintf("rest/api/%v/filter", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	filter := new(model.FilterScheme)
	response, err := i.c.Call(request, filter)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return filter, response, nil
}

func (i *internalFilterServiceImpl) Favorite(ctx context.Context) ([]*model.FilterScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalFilterServiceImpl).Favorite", spanWithKind(trace.SpanKindClient))
	defer span.End()

	endpoint := fmt.Sprintf("rest/api/%v/filter/favourite", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	var filters []*model.FilterScheme
	response, err := i.c.Call(request, filters)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return filters, response, nil
}

func (i *internalFilterServiceImpl) My(ctx context.Context, favorites bool, expand []string) ([]*model.FilterScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalFilterServiceImpl).My", spanWithKind(trace.SpanKindClient))
	defer span.End()

	params := url.Values{}
	params.Add("includeFavourites", fmt.Sprintf("%v", favorites))

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/my?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	var filters []*model.FilterScheme
	response, err := i.c.Call(request, filters)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return filters, response, nil
}

func (i *internalFilterServiceImpl) Search(ctx context.Context, options *model.FilterSearchOptionScheme, startAt, maxResults int) (*model.FilterSearchPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalFilterServiceImpl).Search", spanWithKind(trace.SpanKindClient))
	defer span.End()

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
		recordError(span, err)
		return nil, nil, err
	}

	page := new(model.FilterSearchPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalFilterServiceImpl) Get(ctx context.Context, filterID int, expand []string) (*model.FilterScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalFilterServiceImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if filterID == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoFilterID)
		recordError(span, err)
		return nil, nil, err
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
		recordError(span, err)
		return nil, nil, err
	}

	filter := new(model.FilterScheme)
	response, err := i.c.Call(request, filter)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return filter, response, nil
}

func (i *internalFilterServiceImpl) Update(ctx context.Context, filterID int, payload *model.FilterPayloadScheme) (*model.FilterScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalFilterServiceImpl).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if filterID == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoFilterID)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v", i.version, filterID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	filter := new(model.FilterScheme)
	response, err := i.c.Call(request, filter)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return filter, response, nil
}

func (i *internalFilterServiceImpl) Delete(ctx context.Context, filterID int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalFilterServiceImpl).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if filterID == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoFilterID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v", i.version, filterID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

func (i *internalFilterServiceImpl) Change(ctx context.Context, filterID int, accountID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalFilterServiceImpl).Change", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if filterID == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoFilterID)
		recordError(span, err)
		return nil, err
	}

	if accountID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoAccountID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v/owner", i.version, filterID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]interface{}{"accountId": accountID})
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}
