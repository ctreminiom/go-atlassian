package confluence

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type ContentConnector interface {

	// Gets returns all content in a Confluence instance.
	//
	// GET /wiki/rest/api/content
	//
	// https://docs.go-atlassian.io/confluence-cloud/content#get-content
	Gets(ctx context.Context, options *model.GetContentOptionsScheme, startAt, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error)

	// Create creates a new piece of content or publishes an existing draft
	//
	// To publish a draft, add the id and status properties to the body of the request.
	//
	// Set the id to the ID of the draft and set the status to 'current'.
	//
	// When the request is sent, a new piece of content will be created and the metadata from the draft will be transferred into it.
	//
	// POST /wiki/rest/api/content
	//
	// https://docs.go-atlassian.io/confluence-cloud/content#create-content
	Create(ctx context.Context, payload *model.ContentScheme) (*model.ContentScheme, *model.ResponseScheme, error)

	// Search returns the list of content that matches a Confluence Query Language (CQL) query
	//
	// GET /wiki/rest/api/content/search
	//
	// https://docs.go-atlassian.io/confluence-cloud/content#search-contents-by-cql
	Search(ctx context.Context, cql, cqlContext string, expand []string, cursor string, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error)

	// Get returns a single piece of content, like a page or a blog post.
	//
	// By default, the following objects are expanded: space, history, version.
	//
	// GET /wiki/rest/api/content/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/content#get-content
	Get(ctx context.Context, contentID string, expand []string, version int) (*model.ContentScheme, *model.ResponseScheme, error)

	// Update updates a piece of content.
	//
	// Use this method to update the title or body of a piece of content, change the status, change the parent page, and more.
	//
	// PUT /wiki/rest/api/content/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/content#update-content
	Update(ctx context.Context, contentID string, payload *model.ContentScheme) (*model.ContentScheme, *model.ResponseScheme, error)

	// Delete moves a piece of content to the space's trash or purges it from the trash, depending on the content's type and status:
	//
	// If the content's type is page or blogpost and its status is current, it will be trashed.
	//
	// If the content's type is page or blogpost and its status is trashed, the content will be purged from the trash and deleted permanently.
	//
	// === Note, you must also set the status query parameter to trashed in your request. ===
	//
	// If the content's type is comment or attachment, it will be deleted permanently without being trashed.
	//
	// DELETE /wiki/rest/api/content/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/content#delete-content
	Delete(ctx context.Context, contentID, status string) (*model.ResponseScheme, error)

	// History returns the most recent update for a piece of content.
	//
	// GET /wiki/rest/api/content/{id}/history
	//
	// https://docs.go-atlassian.io/confluence-cloud/content#get-content-history
	History(ctx context.Context, contentID string, expand []string) (*model.ContentHistoryScheme, *model.ResponseScheme, error)

	// Archive archives a list of pages.
	//
	// The pages to be archived are specified as a list of content IDs.
	//
	// This API accepts the archival request and returns a task ID. The archival process happens asynchronously.
	//
	// Use the /longtask/ REST API to get the copy task status.
	//
	// POST /wiki/rest/api/content/archive
	//
	// https://docs.go-atlassian.io/confluence-cloud/content#archive-pages
	Archive(ctx context.Context, payload *model.ContentArchivePayloadScheme) (*model.ContentArchiveResultScheme, *model.ResponseScheme, error)
}
