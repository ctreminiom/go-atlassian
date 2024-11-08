package confluence

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// AnalyticsConnector the interface for the analytics methods of the Confluence Service.
type AnalyticsConnector interface {

	// Get gets the total number of views a piece of content has.
	//
	// GET /wiki/rest/api/analytics/content/{contentID}/views
	//
	// https://docs.go-atlassian.io/confluence-cloud/analytics#get-views
	Get(ctx context.Context, contentID, fromDate string) (*model.ContentViewScheme, *model.ResponseScheme, error)

	// Distinct get the total number of distinct viewers a piece of content has.
	//
	// GET /wiki/rest/api/analytics/content/{contentID}/viewers
	//
	// https://docs.go-atlassian.io/confluence-cloud/analytics#get-viewers
	Distinct(ctx context.Context, contentID, fromDate string) (*model.ContentViewScheme, *model.ResponseScheme, error)
}
