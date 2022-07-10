package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func NewFilterService(client service.Client, version string, share jira.FilterShare) (*FilterService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &FilterService{
		internalClient: &internalFilterServiceImpl{c: client, version: version},
		Share:          share,
	}, nil
}

type FilterService struct {
	internalClient jira.Filter
	Share          jira.FilterShare
}

func (f *FilterService) Create(ctx context.Context, payload *model.FilterPayloadScheme) (*model.FilterScheme, *model.ResponseScheme, error) {
	return f.internalClient.Create(ctx, payload)
}

func (f *FilterService) Favorite(ctx context.Context) ([]*model.FilterScheme, *model.ResponseScheme, error) {
	return f.internalClient.Favorite(ctx)
}

func (f *FilterService) My(ctx context.Context, favorites bool, expand []string) ([]*model.FilterScheme, *model.ResponseScheme, error) {
	return f.internalClient.My(ctx, favorites, expand)
}

func (f *FilterService) Search(ctx context.Context, options *model.FilterSearchOptionScheme, startAt, maxResults int) (*model.FilterSearchPageScheme,
	*model.ResponseScheme, error) {
	return f.internalClient.Search(ctx, options, startAt, maxResults)
}

func (f *FilterService) Get(ctx context.Context, filterId int, expand []string) (*model.FilterScheme, *model.ResponseScheme, error) {
	return f.internalClient.Get(ctx, filterId, expand)
}

func (f *FilterService) Update(ctx context.Context, filterId int, payload *model.FilterPayloadScheme) (*model.FilterScheme, *model.ResponseScheme, error) {
	return f.internalClient.Update(ctx, filterId, payload)
}

func (f *FilterService) Delete(ctx context.Context, filterId int) (*model.ResponseScheme, error) {
	return f.internalClient.Delete(ctx, filterId)
}

func (f *FilterService) Change(ctx context.Context, filterId int, accountId string) (*model.ResponseScheme, error) {
	return f.internalClient.Change(ctx, filterId, accountId)
}

type internalFilterServiceImpl struct {
	c       service.Client
	version string
}

func (i *internalFilterServiceImpl) Create(ctx context.Context, payload *model.FilterPayloadScheme) (*model.FilterScheme, *model.ResponseScheme, error) {

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter", i.version)

	request, err := i.c.NewJsonRequest(ctx, http.MethodPost, endpoint, reader)
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

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

func (i *internalFilterServiceImpl) Get(ctx context.Context, filterId int, expand []string) (*model.FilterScheme, *model.ResponseScheme, error) {
	if filterId == 0 {
		return nil, nil, model.ErrNoFilterIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/filter/%v", i.version, filterId))

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
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

func (i *internalFilterServiceImpl) Update(ctx context.Context, filterId int, payload *model.FilterPayloadScheme) (*model.FilterScheme, *model.ResponseScheme, error) {

	if filterId == 0 {
		return nil, nil, model.ErrNoFieldIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v", i.version, filterId)

	request, err := i.c.NewJsonRequest(ctx, http.MethodPut, endpoint, reader)
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

func (i *internalFilterServiceImpl) Delete(ctx context.Context, filterId int) (*model.ResponseScheme, error) {
	if filterId == 0 {
		return nil, model.ErrNoFilterIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v", i.version, filterId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		return response, err
	}

	return i.c.Call(request, nil)
}

func (i *internalFilterServiceImpl) Change(ctx context.Context, filterId int, accountId string) (*model.ResponseScheme, error) {
	if filterId == 0 {
		return nil, model.ErrNoFilterIDError
	}

	if accountId == "" {
		return nil, model.ErrNoAccountIDError
	}

	payload := struct {
		AccountID string `json:"accountId"`
	}{
		AccountID: accountId,
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/filter/%v/owner", i.version, filterId)

	request, err := i.c.NewJsonRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
