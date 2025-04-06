package confluence

import (
	"context"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// FolderConnector represents the Confluence Cloud Folders.
// Use it to search, get, create, delete, and change folders.
type FolderConnector interface {

	// Create creates a folder in the space.
	//
	// Folders are created as published by default unless specified as a draft in the status field.
	//
	// If creating a published folder, the title must be specified.
	//
	// POST /wiki/api/v2/folders
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/folder#create-folder
	Create(ctx context.Context, payload *models.FolderCreatePayloadScheme) (*models.FolderScheme, *models.ResponseScheme, error)

	// Get returns a specific folder.
	//
	// GET /wiki/api/v2/folders/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/folder#get-folder-by-id
	Get(ctx context.Context, folderID int) (*models.FolderScheme, *models.ResponseScheme, error)

	// Delete deletes a folder by id.
	//
	// DELETE /wiki/api/v2/folders/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/folder#delete-folder
	Delete(ctx context.Context, folderID int) (*models.ResponseScheme, error)
}
