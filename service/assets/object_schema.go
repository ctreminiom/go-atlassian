package assets

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// ObjectSchemaConnector represents the Assets object schemes.
// Use it to search, get, create, delete, and change schemes.
type ObjectSchemaConnector interface {

	// List returns all the object schemes available on Assets
	//
	// GET /jsm/assets/workspace/{workspaceId}/v1/objectschema/list
	List(ctx context.Context, workspaceID string) (*models.ObjectSchemaPageScheme, *models.ResponseScheme, error)

	// Create creates a new object schema
	//
	// POST /jsm/assets/workspace/{workspaceId}/v1/objectschema/create
	Create(ctx context.Context, workspaceID string, payload *models.ObjectSchemaPayloadScheme) (*models.ObjectSchemaScheme,
		*models.ResponseScheme, error)

	// Get returns an object scheme by ID
	//
	// GET /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}
	Get(ctx context.Context, workspaceID, objectSchemaID string) (*models.ObjectSchemaScheme, *models.ResponseScheme, error)

	// Update updates an object schema
	//
	// PUT /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}
	Update(ctx context.Context, workspaceID, objectSchemaID string, payload *models.ObjectSchemaPayloadScheme) (
		*models.ObjectSchemaScheme, *models.ResponseScheme, error)

	// Delete deletes a schema
	//
	// DELETE /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}
	Delete(ctx context.Context, workspaceID, objectSchemaID string) (*models.ObjectSchemaScheme, *models.ResponseScheme, error)

	// Attributes finds all object type attributes for this object schema
	//
	// GET /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}/attributes
	Attributes(ctx context.Context, workspaceID, objectSchemaID string, options *models.ObjectSchemaAttributesParamsScheme) (
		[]*models.ObjectTypeAttributeScheme, *models.ResponseScheme, error)

	// ObjectTypes returns all object types for this object schema
	//
	// GET /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}/objecttypes
	ObjectTypes(ctx context.Context, workspaceID, objectSchemaID string, excludeAbstract bool) (
		*models.ObjectSchemaTypePageScheme, *models.ResponseScheme, error)
}
