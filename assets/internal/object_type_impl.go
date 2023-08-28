package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/assets"
	"net/http"
	"net/url"
	"strings"
)

func NewObjectTypeService(client service.Connector) *ObjectTypeService {

	return &ObjectTypeService{
		internalClient: &internalObjectTypeImpl{c: client},
	}
}

type ObjectTypeService struct {
	internalClient assets.ObjectTypeConnector
}

// Get finds an object type by id
//
// GET /jsm/assets/workspace/{workspaceId}/v1/objecttype/{id}
//
// https://docs.go-atlassian.io/jira-assets/object/type#get-object-type
func (o *ObjectTypeService) Get(ctx context.Context, workspaceID, objectTypeID string) (*model.ObjectTypeScheme, *model.ResponseScheme, error) {
	return o.internalClient.Get(ctx, workspaceID, objectTypeID)
}

// Update updates an existing object type
//
// PUT /jsm/assets/workspace/{workspaceId}/v1/objecttype/{id}
//
// https://docs.go-atlassian.io/jira-assets/object/type#update-object-type
func (o *ObjectTypeService) Update(ctx context.Context, workspaceID, objectTypeID string, payload *model.ObjectTypePayloadScheme) (*model.ObjectTypeScheme, *model.ResponseScheme, error) {
	return o.internalClient.Update(ctx, workspaceID, objectTypeID, payload)
}

// Create creates a new object type
//
// POST /jsm/assets/workspace/{workspaceId}/v1/objecttype/create
//
// https://docs.go-atlassian.io/jira-assets/object/type#create-object-type
func (o *ObjectTypeService) Create(ctx context.Context, workspaceID string, payload *model.ObjectTypePayloadScheme) (*model.ObjectTypeScheme, *model.ResponseScheme, error) {
	return o.internalClient.Create(ctx, workspaceID, payload)
}

// Delete deletes an object type
//
// DELETE /jsm/assets/workspace/{workspaceId}/v1/objecttype/{id}
//
// https://docs.go-atlassian.io/jira-assets/object/type#delete-object-type
func (o *ObjectTypeService) Delete(ctx context.Context, workspaceID, objectTypeID string) (*model.ObjectTypeScheme, *model.ResponseScheme, error) {
	return o.internalClient.Delete(ctx, workspaceID, objectTypeID)
}

// Attributes finds all attributes for this object type
//
// GET /jsm/assets/workspace/{workspaceId}/v1/objecttype/{id}/attributes
//
// https://docs.go-atlassian.io/jira-assets/object/type#get-object-type-attributes
func (o *ObjectTypeService) Attributes(ctx context.Context, workspaceID, objectTypeID string, options *model.ObjectTypeAttributesParamsScheme) ([]*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {
	return o.internalClient.Attributes(ctx, workspaceID, objectTypeID, options)
}

// Position changes the position of this object type
//
// POST /jsm/assets/workspace/{workspaceId}/v1/objecttype/{id}/position
//
// https://docs.go-atlassian.io/jira-assets/object/type#update-object-type-position
func (o *ObjectTypeService) Position(ctx context.Context, workspaceID, objectTypeID string, payload *model.ObjectTypePositionPayloadScheme) (*model.ObjectTypeScheme, *model.ResponseScheme, error) {
	return o.internalClient.Position(ctx, workspaceID, objectTypeID, payload)
}

type internalObjectTypeImpl struct {
	c service.Connector
}

func (i *internalObjectTypeImpl) Get(ctx context.Context, workspaceID, objectTypeID string) (*model.ObjectTypeScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectTypeID == "" {
		return nil, nil, model.ErrNoObjectTypeIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objecttype/%v", workspaceID, objectTypeID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	objectType := new(model.ObjectTypeScheme)
	res, err := i.c.Call(req, objectType)
	if err != nil {
		return nil, res, err
	}

	return objectType, res, nil
}

func (i *internalObjectTypeImpl) Update(ctx context.Context, workspaceID, objectTypeID string, payload *model.ObjectTypePayloadScheme) (*model.ObjectTypeScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectTypeID == "" {
		return nil, nil, model.ErrNoObjectTypeIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objecttype/%v", workspaceID, objectTypeID)

	req, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	objectType := new(model.ObjectTypeScheme)
	res, err := i.c.Call(req, objectType)
	if err != nil {
		return nil, res, err
	}

	return objectType, res, nil
}

func (i *internalObjectTypeImpl) Create(ctx context.Context, workspaceID string, payload *model.ObjectTypePayloadScheme) (*model.ObjectTypeScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objecttype/create", workspaceID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	objectType := new(model.ObjectTypeScheme)
	res, err := i.c.Call(req, objectType)
	if err != nil {
		return nil, res, err
	}

	return objectType, res, nil
}

func (i *internalObjectTypeImpl) Delete(ctx context.Context, workspaceID, objectTypeID string) (*model.ObjectTypeScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectTypeID == "" {
		return nil, nil, model.ErrNoObjectTypeIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objecttype/%v", workspaceID, objectTypeID)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	objectType := new(model.ObjectTypeScheme)
	res, err := i.c.Call(req, objectType)
	if err != nil {
		return nil, res, err
	}

	return objectType, res, nil
}

func (i *internalObjectTypeImpl) Attributes(ctx context.Context, workspaceID, objectTypeID string, options *model.ObjectTypeAttributesParamsScheme) ([]*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectTypeID == "" {
		return nil, nil, model.ErrNoObjectTypeIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("jsm/assets/workspace/%v/v1/objecttype/%v/attributes", workspaceID, objectTypeID))

	query := url.Values{}
	if options != nil {

		if options.OnlyValueEditable {
			query.Add("onlyValueEditable", "true")
		}

		if options.OrderByName {
			query.Add("orderByName", "true")
		}

		if options.Query != "" {
			query.Add("query", options.Query)
		}

		if options.IncludeValuesExist {
			query.Add("includeValuesExist", "true")
		}

		if options.ExcludeParentAttributes {
			query.Add("excludeParentAttributes", "true")
		}

		if options.IncludeChildren {
			query.Add("includeChildren", "true")
		}

		if options.OrderByRequired {
			query.Add("orderByRequired", "true")
		}
	}

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	var attributes []*model.ObjectTypeAttributeScheme
	res, err := i.c.Call(req, &attributes)
	if err != nil {
		return nil, res, err
	}

	return attributes, res, nil
}

func (i *internalObjectTypeImpl) Position(ctx context.Context, workspaceID, objectTypeID string, payload *model.ObjectTypePositionPayloadScheme) (*model.ObjectTypeScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectTypeID == "" {
		return nil, nil, model.ErrNoObjectTypeIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objecttype/%v/position", workspaceID, objectTypeID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	objectType := new(model.ObjectTypeScheme)
	res, err := i.c.Call(req, objectType)
	if err != nil {
		return nil, res, err
	}

	return objectType, res, nil
}
