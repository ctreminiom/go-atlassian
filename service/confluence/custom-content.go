package confluence

import (
	"context"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type CustomContentConnector interface {

	// Gets returns all custom content for a given type.
	//
	// GET /wiki/api/v2/custom-content
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/custom-content#get-custom-content-by-type
	Gets(ctx context.Context, typ string, options *models.CustomContentOptionsScheme, cursor string, limit int) (
		*models.CustomContentPageScheme, *models.ResponseScheme, error)

	// Create creates a new custom content in the given space, page, blogpost or other custom content.
	//
	// POST /wiki/api/v2/custom-content
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/custom-content#create-custom-content
	Create(ctx context.Context, payload *models.CustomContentPayloadScheme) (*models.CustomContentScheme, *models.ResponseScheme, error)

	// Get returns a specific piece of custom content.
	//
	// GET /wiki/api/v2/custom-content/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/custom-content#get-custom-content-by-id
	Get(ctx context.Context, customContentID int, format string, versionID int) (*models.CustomContentScheme, *models.ResponseScheme, error)

	// Update updates a custom content by id.
	//
	// The space id is always required and maximum one of page id, blog post id,
	//
	// or custom content id is allowed in the request body
	//
	// PUT /wiki/api/v2/custom-content/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/custom-content#update-custom-content
	Update(ctx context.Context, customContentID int, payload *models.CustomContentPayloadScheme) (*models.CustomContentScheme, *models.ResponseScheme, error)

	// Delete deletes a custom content by id.
	//
	// DELETE /wiki/api/v2/custom-content/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/custom-content#delete-custom-content
	Delete(ctx context.Context, customContentID int) (*models.ResponseScheme, error)
}
