package assets

import (
	"context"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// ObjectTypeAttributeConnector represents the Assets object type attributes.
// Use it to search, get, create, delete, and change object type attributes.
type ObjectTypeAttributeConnector interface {

	// Create creates a new attribute on the given object type
	//
	// POST /jsm/assets/workspace/{workspaceID}/v1/objecttypeattribute/{objectTypeID}
	//
	// https://docs.go-atlassian.io/jira-assets/object/type/attribute#create-object-type-attribute
	Create(ctx context.Context, workspaceID, objectTypeID string, payload *models.ObjectTypeAttributePayloadScheme) (
		*models.ObjectTypeAttributeScheme, *models.ResponseScheme, error)

	// Update updates an existing object type attribute
	//
	// PUT /jsm/assets/workspace/{workspaceID}/v1/objecttypeattribute/{objectTypeID}/{id}
	//
	// https://docs.go-atlassian.io/jira-assets/object/type/attribute#update-object-type-attribute
	Update(ctx context.Context, workspaceID, objectTypeID, attributeID string, payload *models.ObjectTypeAttributePayloadScheme) (
		*models.ObjectTypeAttributeScheme, *models.ResponseScheme, error)

	// Delete deletes an existing object type attribute
	//
	// DELETE /jsm/assets/workspace/{workspaceID}/v1/objecttypeattribute/{attributeID}
	//
	// https://docs.go-atlassian.io/jira-assets/object/type/attribute#delete-object-type-attribute
	Delete(ctx context.Context, workspaceID, attributeID string) (*models.ResponseScheme, error)
}
