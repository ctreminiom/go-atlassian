package confluence

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type CustomContentConnector interface {

	// Gets returns all custom content for a given type.
	//
	// The number of results is limited by the limit parameter and additional results (if available) will be available
	//
	// through the next URL present in the Link response header.
	//
	// GET /wiki/api/v2/custom-content
	Gets(ctx context.Context, type_ string, options *models.CustomContentOptionsScheme, cursor string, limit int) (
		*models.CustomContentPageScheme, *models.ResponseScheme, error)

	// Create creates a new custom content in the given space, page, blogpost or other custom content.
	//
	// POST /wiki/api/v2/custom-content
	Create(ctx context.Context, payload *models.CustomContentPayloadScheme) (*models.CustomContentScheme, *models.ResponseScheme, error)

	// Get returns a specific piece of custom content.
	//
	// GET /wiki/api/v2/custom-content/{id}
	Get(ctx context.Context, customContentID int, format string, versionID int) (*models.CustomContentScheme, *models.ResponseScheme, error)

	// Update updates a custom content by id.
	//
	// The spaceId is always required and maximum one of pageId, blogPostId,
	//
	// or customContentId is allowed in the request body
	//
	// PUT /wiki/api/v2/custom-content/{id}
	Update(ctx context.Context, customContentID int, payload *models.CustomContentPayloadScheme) (*models.CustomContentScheme, *models.ResponseScheme, error)

	// Delete deletes a custom content by id.
	//
	// DELETE /wiki/api/v2/custom-content/{id}
	Delete(ctx context.Context, customContentID int) (*models.ResponseScheme, error)
}
