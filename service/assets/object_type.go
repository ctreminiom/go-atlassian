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
	//
	// https://docs.go-atlassian.io/jira-assets/object/type#get-object-type
	Get(ctx context.Context, workspaceID, objectTypeID string) (*models.ObjectTypeScheme, *models.ResponseScheme, error)

	// Update updates an existing object type
	//
	// PUT /jsm/assets/workspace/{workspaceId}/v1/objecttype/{id}
	//
	// https://docs.go-atlassian.io/jira-assets/object/type#update-object-type
	Update(ctx context.Context, workspaceID, objectTypeID string, payload *models.ObjectTypePayloadScheme) (
		*models.ObjectTypeScheme, *models.ResponseScheme, error)

	// Create creates a new object type
	//
	// POST /jsm/assets/workspace/{workspaceId}/v1/objecttype/create
	//
	// https://docs.go-atlassian.io/jira-assets/object/type#create-object-type
	Create(ctx context.Context, workspaceID string, payload *models.ObjectTypePayloadScheme) (
		*models.ObjectTypeScheme, *models.ResponseScheme, error)

	// Delete deletes an object type
	//
	// DELETE /jsm/assets/workspace/{workspaceId}/v1/objecttype/{id}
	//
	// https://docs.go-atlassian.io/jira-assets/object/type#delete-object-type
	Delete(ctx context.Context, workspaceID, objectTypeID string) (*models.ObjectTypeScheme, *models.ResponseScheme, error)

	// Attributes finds all attributes for this object type
	//
	// GET /jsm/assets/workspace/{workspaceId}/v1/objecttype/{id}/attributes
	//
	// https://docs.go-atlassian.io/jira-assets/object/type#get-object-type-attributes
	Attributes(ctx context.Context, workspaceID, objectTypeID string, options *models.ObjectTypeAttributesParamsScheme) (
		[]*models.ObjectTypeAttributeScheme, *models.ResponseScheme, error)

	// Position changes the position of this object type
	//
	// POST /jsm/assets/workspace/{workspaceId}/v1/objecttype/{id}/position
	//
	// https://docs.go-atlassian.io/jira-assets/object/type#update-object-type-position
	Position(ctx context.Context, workspaceID, objectTypeID string, payload *models.ObjectTypePositionPayloadScheme) (
		*models.ObjectTypeScheme, *models.ResponseScheme, error)
}
