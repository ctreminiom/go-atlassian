package confluence

import (
	"context"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// DescendantsConnector represents the Confluence Cloud Descendants.
// Use it to get descendants of several resources.
type DescendantsConnector interface {

	// Get descendants of a whiteboard.
	//
	// GET /wiki/api/v2/whiteboards/{id}/descendants
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/descendants#get-descendants-of-a-whiteboard
	GetForWhiteboard(ctx context.Context, whiteboardID int, limit int, depth int, cursor string) (*models.DescendantsScheme, *models.ResponseScheme, error)

	// Get descendants of a database.
	//
	// GET /wiki/api/v2/databases/{id}/descendants
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/descendants#get-descendants-of-a-database
	GetForDatabase(ctx context.Context, databaseID int, limit int, depth int, cursor string) (*models.DescendantsScheme, *models.ResponseScheme, error)

	// Get descendants of a smart link.
	//
	// GET /wiki/api/v2/embeds/{id}/descendants
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/descendants#get-descendants-of-a-smart-link
	GetForSmartLink(ctx context.Context, embedID int, limit int, depth int, cursor string) (*models.DescendantsScheme, *models.ResponseScheme, error)

	// Get descendants of a folder.
	//
	// GET /wiki/api/v2/folders/{id}/descendants
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/descendants#get-descendants-of-a-folder
	GetForFolder(ctx context.Context, folderID int, limit int, depth int, cursor string) (*models.DescendantsScheme, *models.ResponseScheme, error)

	// Get descendants of a page.
	//
	// GET /wiki/api/v2/pages/{id}/descendants
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/descendants#get-descendants-of-a-page
	GetForPage(ctx context.Context, pageID int, limit int, depth int, cursor string) (*models.DescendantsScheme, *models.ResponseScheme, error)
}
