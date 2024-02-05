package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/assets"
	"net/http"
)

func NewObjectTypeAttributeService(client service.Connector) *ObjectTypeAttributeService {

	return &ObjectTypeAttributeService{
		internalClient: &internalObjectTypeAttributeImpl{c: client},
	}
}

type ObjectTypeAttributeService struct {
	internalClient assets.ObjectTypeAttributeConnector
}

// Create creates a new attribute on the given object type
//
// POST /jsm/assets/workspace/{workspaceId}/v1/objecttypeattribute/{objectTypeId}
//
// https://docs.go-atlassian.io/jira-assets/object/type/attribute#create-object-type-attribute
func (o *ObjectTypeAttributeService) Create(ctx context.Context, workspaceID, objectTypeID string, payload *model.ObjectTypeAttributePayloadScheme) (*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {
	return o.internalClient.Create(ctx, workspaceID, objectTypeID, payload)
}

// Update updates an existing object type attribute
//
// PUT /jsm/assets/workspace/{workspaceId}/v1/objecttypeattribute/{objectTypeId}/{id}
//
// https://docs.go-atlassian.io/jira-assets/object/type/attribute#update-object-type-attribute
func (o *ObjectTypeAttributeService) Update(ctx context.Context, workspaceID, objectTypeID, attributeID string, payload *model.ObjectTypeAttributeScheme) (*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {
	return o.internalClient.Update(ctx, workspaceID, objectTypeID, attributeID, payload)
}

// Delete deletes an existing object type attribute
//
// DELETE /jsm/assets/workspace/{workspaceId}/v1/objecttypeattribute/{id}
//
// https://docs.go-atlassian.io/jira-assets/object/type/attribute#delete-object-type-attribute
func (o *ObjectTypeAttributeService) Delete(ctx context.Context, workspaceID, attributeID string) (*model.ResponseScheme, error) {
	return o.internalClient.Delete(ctx, workspaceID, attributeID)
}

type internalObjectTypeAttributeImpl struct {
	c service.Connector
}

func (i *internalObjectTypeAttributeImpl) Create(ctx context.Context, workspaceID, objectTypeID string, payload *model.ObjectTypeAttributePayloadScheme) (*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectTypeID == "" {
		return nil, nil, model.ErrNoObjectTypeIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objecttypeattribute/%v", workspaceID, objectTypeID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	attribute := new(model.ObjectTypeAttributeScheme)
	res, err := i.c.Call(req, attribute)
	if err != nil {
		return nil, res, err
	}

	return attribute, res, nil
}

func (i *internalObjectTypeAttributeImpl) Update(ctx context.Context, workspaceID, objectTypeID, attributeID string, payload *model.ObjectTypeAttributePayloadScheme) (*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectTypeID == "" {
		return nil, nil, model.ErrNoObjectTypeIDError
	}

	if attributeID == "" {
		return nil, nil, model.ErrNoObjectTypeAttributeIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objecttypeattribute/%v/%v", workspaceID, objectTypeID, attributeID)

	req, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	attribute := new(model.ObjectTypeAttributeScheme)
	res, err := i.c.Call(req, attribute)
	if err != nil {
		return nil, res, err
	}

	return attribute, res, nil
}

func (i *internalObjectTypeAttributeImpl) Delete(ctx context.Context, workspaceID, attributeID string) (*model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, model.ErrNoWorkspaceIDError
	}

	if attributeID == "" {
		return nil, model.ErrNoObjectTypeAttributeIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objecttypeattribute/%v", workspaceID, attributeID)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}
