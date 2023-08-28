package assets

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// ObjectTypeAttributeConnector represents the Assets object type attributes.
// Use it to search, get, create, delete, and change object type attributes.
type ObjectTypeAttributeConnector interface {

	// Create creates a new attribute on the given object type
	//
	// POST /jsm/assets/workspace/{workspaceId}/v1/objecttypeattribute/{objectTypeId}
	//
	// https://docs.go-atlassian.io/jira-assets/object/type/attribute#create-object-type-attribute
	Create(ctx context.Context, workspaceID, objectTypeID string, payload *models.ObjectTypeAttributeScheme) (
		*models.ObjectTypeAttributeScheme, *models.ResponseScheme, error)

	// Update updates an existing object type attribute
	//
	// PUT /jsm/assets/workspace/{workspaceId}/v1/objecttypeattribute/{objectTypeId}/{id}
	//
	// https://docs.go-atlassian.io/jira-assets/object/type/attribute#update-object-type-attribute
	Update(ctx context.Context, workspaceID, objectTypeID, attributeID string, payload *models.ObjectTypeAttributeScheme) (
		*models.ObjectTypeAttributeScheme, *models.ResponseScheme, error)

	// Delete deletes an existing object type attribute
	//
	// DELETE /jsm/assets/workspace/{workspaceId}/v1/objecttypeattribute/{id}
	//
	// https://docs.go-atlassian.io/jira-assets/object/type/attribute#delete-object-type-attribute
	Delete(ctx context.Context, workspaceID, attributeID string) (*models.ResponseScheme, error)
}
