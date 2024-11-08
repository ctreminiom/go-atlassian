package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/assets"
)

// NewObjectTypeAttributeService creates a new instance of ObjectTypeAttributeService.
// It takes a service.Connector as input and returns a pointer to ObjectTypeAttributeService.
func NewObjectTypeAttributeService(client service.Connector) *ObjectTypeAttributeService {
	return &ObjectTypeAttributeService{
		internalClient: &internalObjectTypeAttributeImpl{c: client},
	}
}

// ObjectTypeAttributeService provides methods to interact with object type attributes in Jira Assets.
type ObjectTypeAttributeService struct {
	// internalClient is the connector interface for object type attribute operations.
	internalClient assets.ObjectTypeAttributeConnector
}

// Create creates a new attribute on the given object type
//
// POST /jsm/assets/workspace/{workspaceId}/v1/objecttypeattribute/{objectTypeId}
//
// https://docs.go-atlassian.io/jira-assets/object/type/attribute#create-object-type-attribute
func (o *ObjectTypeAttributeService) Create(ctx context.Context, workspaceID, objectTypeID string, payload *model.ObjectTypeAttributePayloadScheme) (*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ObjectTypeAttributeService).Create")
	defer span.End()

	return o.internalClient.Create(ctx, workspaceID, objectTypeID, payload)
}

// Update updates an existing object type attribute
//
// PUT /jsm/assets/workspace/{workspaceId}/v1/objecttypeattribute/{objectTypeId}/{id}
//
// https://docs.go-atlassian.io/jira-assets/object/type/attribute#update-object-type-attribute
func (o *ObjectTypeAttributeService) Update(ctx context.Context, workspaceID, objectTypeID, attributeID string, payload *model.ObjectTypeAttributePayloadScheme) (*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ObjectTypeAttributeService).Update")
	defer span.End()

	return o.internalClient.Update(ctx, workspaceID, objectTypeID, attributeID, payload)
}

// Delete deletes an existing object type attribute
//
// DELETE /jsm/assets/workspace/{workspaceId}/v1/objecttypeattribute/{id}
//
// https://docs.go-atlassian.io/jira-assets/object/type/attribute#delete-object-type-attribute
func (o *ObjectTypeAttributeService) Delete(ctx context.Context, workspaceID, attributeID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ObjectTypeAttributeService).Delete")
	defer span.End()

	return o.internalClient.Delete(ctx, workspaceID, attributeID)
}

type internalObjectTypeAttributeImpl struct {
	c service.Connector
}

func (i *internalObjectTypeAttributeImpl) Create(ctx context.Context, workspaceID, objectTypeID string, payload *model.ObjectTypeAttributePayloadScheme) (*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalObjectTypeAttributeImpl).Create")
	defer span.End()

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceID
	}

	if objectTypeID == "" {
		return nil, nil, model.ErrNoObjectTypeID
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
	ctx, span := tracer().Start(ctx, "(*internalObjectTypeAttributeImpl).Update")
	defer span.End()

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceID
	}

	if objectTypeID == "" {
		return nil, nil, model.ErrNoObjectTypeID
	}

	if attributeID == "" {
		return nil, nil, model.ErrNoObjectTypeAttributeID
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
	ctx, span := tracer().Start(ctx, "(*internalObjectTypeAttributeImpl).Delete")
	defer span.End()

	if workspaceID == "" {
		return nil, model.ErrNoWorkspaceID
	}

	if attributeID == "" {
		return nil, model.ErrNoObjectTypeAttributeID
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objecttypeattribute/%v", workspaceID, attributeID)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}
