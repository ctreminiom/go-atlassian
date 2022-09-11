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

func NewScreenSchemeService(client service.Client, version string) (*ScreenSchemeService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ScreenSchemeService{
		internalClient: &internalScreenSchemeImpl{c: client, version: version},
	}, nil
}

type ScreenSchemeService struct {
	internalClient jira.ScreenSchemeConnector
}

// Gets returns a paginated list of screen schemes.
//
// Only screen schemes used in classic projects are returned.
//
// GET /rest/api/{2-3}/screenscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#get-screen-schemes
func (s *ScreenSchemeService) Gets(ctx context.Context, options *model.ScreenSchemeParamsScheme, startAt, maxResults int) (*model.ScreenSchemePageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Gets(ctx, options, startAt, maxResults)
}

// Create creates a screen scheme.
//
// POST /rest/api/{2-3}/screenscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#create-screen-scheme
func (s *ScreenSchemeService) Create(ctx context.Context, payload *model.ScreenSchemePayloadScheme) (*model.ScreenSchemeScheme, *model.ResponseScheme, error) {
	return s.internalClient.Create(ctx, payload)
}

// Update updates a screen scheme. Only screen schemes used in classic projects can be updated.
//
// PUT /rest/api/{2-3}/screenscheme/{screenSchemeId}
//
// https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#update-screen-scheme
func (s *ScreenSchemeService) Update(ctx context.Context, screenSchemeId string, payload *model.ScreenSchemePayloadScheme) (*model.ResponseScheme, error) {
	return s.internalClient.Update(ctx, screenSchemeId, payload)
}

// Delete deletes a screen scheme. A screen scheme cannot be deleted if it is used in an issue type screen scheme.
//
// Only screens schemes used in classic projects can be deleted.
//
// DELETE /rest/api/{2-3}/screenscheme/{screenSchemeId}
//
// https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#delete-screen-scheme
func (s *ScreenSchemeService) Delete(ctx context.Context, screenSchemeId string) (*model.ResponseScheme, error) {
	return s.internalClient.Delete(ctx, screenSchemeId)
}

type internalScreenSchemeImpl struct {
	c       service.Client
	version string
}

func (i *internalScreenSchemeImpl) Gets(ctx context.Context, options *model.ScreenSchemeParamsScheme, startAt, maxResults int) (*model.ScreenSchemePageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		for _, id := range options.IDs {
			params.Add("id", strconv.Itoa(id))
		}

		if options.QueryString != "" {
			params.Add("queryString", options.QueryString)
		}

		if options.OrderBy != "orderBy" {
			params.Add("", options.OrderBy)
		}

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/screenscheme?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ScreenSchemePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalScreenSchemeImpl) Create(ctx context.Context, payload *model.ScreenSchemePayloadScheme) (*model.ScreenSchemeScheme, *model.ResponseScheme, error) {

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/screenscheme", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	scheme := new(model.ScreenSchemeScheme)
	response, err := i.c.Call(request, scheme)
	if err != nil {
		return nil, response, err
	}

	return scheme, response, nil
}

func (i *internalScreenSchemeImpl) Update(ctx context.Context, screenSchemeId string, payload *model.ScreenSchemePayloadScheme) (*model.ResponseScheme, error) {

	if screenSchemeId == "" {
		return nil, model.ErrNoScreenSchemeIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/screenscheme/%v", i.version, screenSchemeId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalScreenSchemeImpl) Delete(ctx context.Context, screenSchemeId string) (*model.ResponseScheme, error) {

	if screenSchemeId == "" {
		return nil, model.ErrNoScreenSchemeIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/screenscheme/%v", i.version, screenSchemeId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
