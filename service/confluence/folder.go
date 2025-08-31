package confluence

import (
	"context"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// FolderConnector represents the Confluence Cloud Folders.
// Use it to search, get, create, delete, and change folders.
type FolderConnector interface {

	// Get returns a specific folder.
	//
	// GET /wiki/api/v2/folders/{id}
	//
	// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-folder
	Get(ctx context.Context, folderID string) (*models.FolderScheme, *models.ResponseScheme, error)

	// Gets returns all folders that fit the filtering criteria.
	//
	// The number of results is limited by the limit parameter and additional results
	//
	// (if available) will be available through the next cursor
	//
	// GET /wiki/api/v2/folders
	//
	// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-folder
	Gets(ctx context.Context, options *models.FolderOptionsScheme, cursor string, limit int) (*models.FolderChunkScheme, *models.ResponseScheme, error)

	// GetsBySpace returns all folders in a space.
	//
	// The number of results is limited by the limit parameter and additional results (if available)
	//
	// will be available through the next cursor
	//
	// GET /wiki/api/v2/spaces/{id}/folders
	//
	// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-folder
	GetsBySpace(ctx context.Context, spaceID int, cursor string, limit int) (*models.FolderChunkScheme, *models.ResponseScheme, error)

	// GetsByParent returns all child folders of a parent folder.
	//
	// The number of results is limited by the limit parameter and additional results (if available)
	//
	// will be available through the next cursor
	//
	// GET /wiki/api/v2/folders/{id}/children
	//
	// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-folder
	GetsByParent(ctx context.Context, parentID string, cursor string, limit int) (*models.FolderChunkScheme, *models.ResponseScheme, error)

	// Create creates a folder in the space.
	//
	// POST /wiki/api/v2/folders
	//
	// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-folder
	Create(ctx context.Context, payload *models.FolderCreatePayloadScheme) (*models.FolderScheme, *models.ResponseScheme, error)

	// Update updates a folder by id.
	//
	// PUT /wiki/api/v2/folders/{id}
	//
	// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-folder
	Update(ctx context.Context, folderID string, payload *models.FolderUpdatePayloadScheme) (*models.FolderScheme, *models.ResponseScheme, error)

	// Delete deletes a folder by id.
	//
	// DELETE /wiki/api/v2/folders/{id}
	//
	// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-folder
	Delete(ctx context.Context, folderID string) (*models.ResponseScheme, error)
}