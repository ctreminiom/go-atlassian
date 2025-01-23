package assets

import (
	"context"

	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// ObjectSchemaConnector represents the Assets object schemes.
// Use it to search, get, create, delete, and change schemes.
type ObjectSchemaConnector interface {

	// List returns all the object schemas available on Assets
	//
	// GET /jsm/assets/workspace/{workspaceID}/v1/objectschema/list
	//
	// https://docs.go-atlassian.io/jira-assets/object/schema#get-object-schema-list
	List(ctx context.Context, workspaceID string) (*models.ObjectSchemaPageScheme, *models.ResponseScheme, error)

	// Create creates a new object schema
	//
	// POST /jsm/assets/workspace/{workspaceID}/v1/objectschema/create
	//
	// https://docs.go-atlassian.io/jira-assets/object/schema#create-object-schema
	Create(ctx context.Context, workspaceID string, payload *models.ObjectSchemaPayloadScheme) (*models.ObjectSchemaScheme,
		*models.ResponseScheme, error)

	// Get returns an object schema by ID
	//
	// GET /jsm/assets/workspace/{workspaceID}/v1/objectschema/{objectSchemaID}
	//
	// https://docs.go-atlassian.io/jira-assets/object/schema#get-object-schema
	Get(ctx context.Context, workspaceID, objectSchemaID string) (*models.ObjectSchemaScheme, *models.ResponseScheme, error)

	// Update updates an object schema
	//
	// PUT /jsm/assets/workspace/{workspaceID}/v1/objectschema/{objectSchemaID}
	//
	// https://docs.go-atlassian.io/jira-assets/object/schema#update-object-schema
	Update(ctx context.Context, workspaceID, objectSchemaID string, payload *models.ObjectSchemaPayloadScheme) (
		*models.ObjectSchemaScheme, *models.ResponseScheme, error)

	// Delete deletes a schema
	//
	// DELETE /jsm/assets/workspace/{workspaceID}/v1/objectschema/{objectSchemaID}
	//
	// https://docs.go-atlassian.io/jira-assets/object/schema#delete-object-schema
	Delete(ctx context.Context, workspaceID, objectSchemaID string) (*models.ObjectSchemaScheme, *models.ResponseScheme, error)

	// Attributes finds all object type attributes for this object schema
	//
	// GET /jsm/assets/workspace/{workspaceID}/v1/objectschema/{objectSchemaID}/attributes
	//
	// https://docs.go-atlassian.io/jira-assets/object/schema#get-object-schema-attributes
	Attributes(ctx context.Context, workspaceID, objectSchemaID string, options *models.ObjectSchemaAttributesParamsScheme) (
		[]*models.ObjectTypeAttributeScheme, *models.ResponseScheme, error)

	// ObjectTypes returns all object types for this object schema
	//
	// GET /jsm/assets/workspace/{workspaceID}/v1/objectschema/{objectSchemaID}/objecttypes
	//
	// https://docs.go-atlassian.io/jira-assets/object/schema#get-object-schema-types
	ObjectTypes(ctx context.Context, workspaceID, objectSchemaID string, excludeAbstract bool) (
		[]*models.ObjectTypeScheme, *models.ResponseScheme, error)
}
