package confluence

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type AnalyticsConnector interface {

	// Get gets the total number of views a piece of content has.
	//
	// GET /wiki/rest/api/analytics/content/{contentId}/views
	//
	// https://docs.go-atlassian.io/confluence-cloud/analytics#get-views
	Get(ctx context.Context, contentId, fromDate string) (*model.ContentViewScheme, *model.ResponseScheme, error)

	// Distinct get the total number of distinct viewers a piece of content has.
	//
	// GET /wiki/rest/api/analytics/content/{contentId}/viewers
	//
	// https://docs.go-atlassian.io/confluence-cloud/analytics#get-viewers
	Distinct(ctx context.Context, contentId, fromDate string) (*model.ContentViewScheme, *model.ResponseScheme, error)
}
