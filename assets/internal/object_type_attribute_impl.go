package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/assets"
	"net/http"
)

func NewObjectTypeAttributeService(client service.Client) *ObjectTypeAttributeService {

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
func (o *ObjectTypeAttributeService) Create(ctx context.Context, workspaceID, objectTypeID string, payload *model.ObjectTypeAttributeScheme) (*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {
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
	c service.Client
}

func (i *internalObjectTypeAttributeImpl) Create(ctx context.Context, workspaceID, objectTypeID string, payload *model.ObjectTypeAttributeScheme) (*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectTypeID == "" {
		return nil, nil, model.ErrNoObjectTypeIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objecttypeattribute/%v", workspaceID, objectTypeID)

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	attribute := new(model.ObjectTypeAttributeScheme)
	response, err := i.c.Call(request, attribute)
	if err != nil {
		return nil, response, err
	}

	return attribute, response, nil
}

func (i *internalObjectTypeAttributeImpl) Update(ctx context.Context, workspaceID, objectTypeID, attributeID string, payload *model.ObjectTypeAttributeScheme) (*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {

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

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	attribute := new(model.ObjectTypeAttributeScheme)
	response, err := i.c.Call(request, attribute)
	if err != nil {
		return nil, response, err
	}

	return attribute, response, nil
}

func (i *internalObjectTypeAttributeImpl) Delete(ctx context.Context, workspaceID, attributeID string) (*model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, model.ErrNoWorkspaceIDError
	}

	if attributeID == "" {
		return nil, model.ErrNoObjectTypeAttributeIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objecttypeattribute/%v", workspaceID, attributeID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
