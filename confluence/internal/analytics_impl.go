package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/confluence"
	"net/http"
	"net/url"
	"strings"
)

func NewAnalyticsService(client service.Connector) *AnalyticsService {

	return &AnalyticsService{
		internalClient: &internalAnalyticsServiceImpl{c: client},
	}
}

type AnalyticsService struct {
	internalClient confluence.AnalyticsConnector
}

// Get gets the total number of views a piece of content has.
//
// GET /wiki/rest/api/analytics/content/{contentId}/views
//
// https://docs.go-atlassian.io/confluence-cloud/analytics#get-views
func (a *AnalyticsService) Get(ctx context.Context, contentId, fromDate string) (*model.ContentViewScheme, *model.ResponseScheme, error) {
	return a.internalClient.Get(ctx, contentId, fromDate)
}

// Distinct get the total number of distinct viewers a piece of content has.
//
// GET /wiki/rest/api/analytics/content/{contentId}/viewers
//
// https://docs.go-atlassian.io/confluence-cloud/analytics#get-viewers
func (a *AnalyticsService) Distinct(ctx context.Context, contentId, fromDate string) (*model.ContentViewScheme, *model.ResponseScheme, error) {
	return a.internalClient.Distinct(ctx, contentId, fromDate)
}

type internalAnalyticsServiceImpl struct {
	c service.Connector
}

func (i *internalAnalyticsServiceImpl) Get(ctx context.Context, contentId, fromDate string) (*model.ContentViewScheme, *model.ResponseScheme, error) {

	if contentId == "" {
		return nil, nil, model.ErrNoContentIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/analytics/content/%v/views", contentId))

	if fromDate != "" {
		query := url.Values{}
		query.Add("fromDate", fromDate)

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	views := new(model.ContentViewScheme)
	response, err := i.c.Call(request, views)
	if err != nil {
		return nil, response, err
	}

	return views, response, nil
}

func (i *internalAnalyticsServiceImpl) Distinct(ctx context.Context, contentId, fromDate string) (*model.ContentViewScheme, *model.ResponseScheme, error) {

	if contentId == "" {
		return nil, nil, model.ErrNoContentIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/analytics/content/%v/viewers", contentId))

	if fromDate != "" {
		query := url.Values{}
		query.Add("fromDate", fromDate)

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	views := new(model.ContentViewScheme)
	response, err := i.c.Call(request, views)
	if err != nil {
		return nil, response, err
	}

	return views, response, nil
}
