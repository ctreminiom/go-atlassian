package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
)

// NewAnalyticsService creates a new instance of AnalyticsService.
// It takes a service.Connector as input and returns a pointer to AnalyticsService.
func NewAnalyticsService(client service.Connector) *AnalyticsService {
	return &AnalyticsService{
		internalClient: &internalAnalyticsServiceImpl{c: client},
	}
}

// AnalyticsService provides methods to interact with analytics operations in Confluence.
type AnalyticsService struct {
	// internalClient is the connector interface for analytics operations.
	internalClient confluence.AnalyticsConnector
}

// Get gets the total number of views a piece of content has.
//
// GET /wiki/rest/api/analytics/content/{contentID}/views
//
// https://docs.go-atlassian.io/confluence-cloud/analytics#get-views
func (a *AnalyticsService) Get(ctx context.Context, contentID, fromDate string) (*model.ContentViewScheme, *model.ResponseScheme, error) {
	return a.internalClient.Get(ctx, contentID, fromDate)
}

// Distinct get the total number of distinct viewers a piece of content has.
//
// GET /wiki/rest/api/analytics/content/{contentID}/viewers
//
// https://docs.go-atlassian.io/confluence-cloud/analytics#get-viewers
func (a *AnalyticsService) Distinct(ctx context.Context, contentID, fromDate string) (*model.ContentViewScheme, *model.ResponseScheme, error) {
	return a.internalClient.Distinct(ctx, contentID, fromDate)
}

type internalAnalyticsServiceImpl struct {
	c service.Connector
}

func (i *internalAnalyticsServiceImpl) Get(ctx context.Context, contentID, fromDate string) (*model.ContentViewScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/analytics/content/%v/views", contentID))

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

func (i *internalAnalyticsServiceImpl) Distinct(ctx context.Context, contentID, fromDate string) (*model.ContentViewScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/analytics/content/%v/viewers", contentID))

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
