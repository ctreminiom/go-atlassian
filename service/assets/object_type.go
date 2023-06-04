package assets

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// ObjectTypeConnector represents the Assets object types.
// Use it to search, get, create, delete, and change types.
type ObjectTypeConnector interface {

	// Get finds an object type by id
	//
	// GET /jsm/assets/workspace/{workspaceId}/v1/objecttype/{id}
	Get(ctx context.Context, workspaceID, objectTypeID string) (*models.ObjectTypeScheme, *models.ResponseScheme, error)

	// Update updates an existing object type
	//
	// PUT /jsm/assets/workspace/{workspaceId}/v1/objecttype/{id}
	Update(ctx context.Context, workspaceID, objectTypeID string, payload *models.ObjectTypePayloadScheme) (
		*models.ObjectTypeScheme, *models.ResponseScheme, error)

	// Create creates a new object type
	//
	// POST /jsm/assets/workspace/{workspaceId}/v1/objecttype/create
	Create(ctx context.Context, workspaceID string, payload *models.ObjectTypePayloadScheme) (
		*models.ObjectTypeScheme, *models.ResponseScheme, error)

	// Delete deletes an object type
	//
	// DELETE /jsm/assets/workspace/{workspaceId}/v1/objecttype/{id}
	Delete(ctx context.Context, workspaceID, objectTypeID string) (*models.ObjectTypeScheme, *models.ResponseScheme, error)

	// Attributes finds all attributes for this object type
	//
	// GET /jsm/assets/workspace/{workspaceId}/v1/objecttype/{id}/attributes
	Attributes(ctx context.Context, workspaceID, objectTypeID string, options *models.ObjectTypeAttributesParamsScheme) (
		[]*models.ObjectTypeAttributeScheme, *models.ResponseScheme, error)

	// Position changes the position of this object type
	//
	// POST /jsm/assets/workspace/{workspaceId}/v1/objecttype/{id}/position
	Position(ctx context.Context, workspaceID, objectTypeID, toObjectTypeId string, position int) (*models.ObjectTypeScheme,
		*models.ResponseScheme, error)
}
