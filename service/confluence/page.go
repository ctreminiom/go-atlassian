package confluence

import (
	"context"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// PageConnector represents the Confluence Cloud Pages.
// Use it to search, get, create, delete, and change pages.
type PageConnector interface {

	// Get returns a specific page.
	//
	// GET /wiki/api/v2/pages/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/page#get-page-by-id
	Get(ctx context.Context, pageID int, format string, draft bool, version int) (*models.PageScheme, *models.ResponseScheme, error)

	// Bulk returns all pages.
	//
	// The number of results is limited by the limit parameter and additional results
	//
	// (if available) will be available through the next cursor
	//
	// Deprecated. Please use Page.Gets() instead.
	//
	// GET /wiki/api/v2/pages
	//
	Bulk(ctx context.Context, cursor string, limit int) (*models.PageChunkScheme, *models.ResponseScheme, error)

	// Gets returns all pages that fit the filtering criteria.
	//
	// The number of results is limited by the limit parameter and additional results
	//
	// (if available) will be available through the next cursor
	//
	// GET /wiki/api/v2/pages
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/page#get-pages
	Gets(ctx context.Context, options *models.PageOptionsScheme, cursor string, limit int) (*models.PageChunkScheme, *models.ResponseScheme, error)

	// GetsByLabel returns the pages of specified label.
	//
	// The number of results is limited by the limit parameter and additional results
	//
	// (if available) will be available through the next cursor
	//
	// GET /wiki/api/v2/labels/{id}/pages
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/page#get-pages-for-label
	GetsByLabel(ctx context.Context, labelID int, sort, cursor string, limit int) (*models.PageChunkScheme, *models.ResponseScheme, error)

	// GetsBySpace returns all pages in a space.
	//
	// The number of results is limited by the limit parameter and additional results (if available)
	//
	// will be available through the next cursor
	//
	// GET /wiki/api/v2/spaces/{id}/pages
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/page#get-pages-in-space
	GetsBySpace(ctx context.Context, spaceID int, cursor string, limit int) (*models.PageChunkScheme, *models.ResponseScheme, error)

	// GetsByParent returns all children of a page.
	//
	// The number of results is limited by the limit parameter and additional results (if available)
	//
	// will be available through the next cursor
	//
	// GET /wiki/api/v2/pages/{id}/children
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/page#get-pages-by-parent
	GetsByParent(ctx context.Context, spaceID int, cursor string, limit int) (*models.ChildPageChunkScheme, *models.ResponseScheme, error)

	// Create creates a page in the space.
	//
	// Pages are created as published by default unless specified as a draft in the status field.
	//
	// If creating a published page, the title must be specified.
	//
	// POST /wiki/api/v2/pages
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/page#create-page
	Create(ctx context.Context, payload *models.PageCreatePayloadScheme) (*models.PageScheme, *models.ResponseScheme, error)

	// Update updates a page by id.
	//
	// PUT /wiki/api/v2/pages/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/page#update-page
	Update(ctx context.Context, pageID int, payload *models.PageUpdatePayloadScheme) (*models.PageScheme, *models.ResponseScheme, error)

	// Delete deletes a page by id.
	//
	// DELETE /wiki/api/v2/pages/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/page#delete-page
	Delete(ctx context.Context, pageID int) (*models.ResponseScheme, error)
}
