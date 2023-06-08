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

func NewObjectSchemaService(client service.Client) *ObjectSchemaService {

	return &ObjectSchemaService{
		internalClient: &internalObjectSchemaImpl{c: client},
	}
}

type ObjectSchemaService struct {
	internalClient assets.ObjectSchemaConnector
}

// List returns all the object schemes available on Assets
//
// GET /jsm/assets/workspace/{workspaceId}/v1/objectschema/list
func (o *ObjectSchemaService) List(ctx context.Context, workspaceID string) (*model.ObjectSchemaPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.List(ctx, workspaceID)
}

// Create creates a new object schema
//
// POST /jsm/assets/workspace/{workspaceId}/v1/objectschema/create
func (o *ObjectSchemaService) Create(ctx context.Context, workspaceID string, payload *model.ObjectSchemaPayloadScheme) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	return o.internalClient.Create(ctx, workspaceID, payload)
}

// Get returns an object scheme by ID
//
// GET /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}
func (o *ObjectSchemaService) Get(ctx context.Context, workspaceID, objectSchemaID string) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	return o.internalClient.Get(ctx, workspaceID, objectSchemaID)
}

// Update updates an object schema
//
// PUT /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}
func (o *ObjectSchemaService) Update(ctx context.Context, workspaceID, objectSchemaID string, payload *model.ObjectSchemaPayloadScheme) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	return o.internalClient.Update(ctx, workspaceID, objectSchemaID, payload)
}

// Delete deletes a schema
//
// DELETE /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}
func (o *ObjectSchemaService) Delete(ctx context.Context, workspaceID, objectSchemaID string) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	return o.internalClient.Delete(ctx, workspaceID, objectSchemaID)
}

// Attributes finds all object type attributes for this object schema
//
// GET /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}/attributes
func (o *ObjectSchemaService) Attributes(ctx context.Context, workspaceID, objectSchemaID string, options *model.ObjectSchemaAttributesParamsScheme) ([]*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {
	return o.internalClient.Attributes(ctx, workspaceID, objectSchemaID, options)
}

// ObjectTypes returns all object types for this object schema
//
// GET /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}/objecttypes
func (o *ObjectSchemaService) ObjectTypes(ctx context.Context, workspaceID, objectSchemaID string, excludeAbstract bool) (*model.ObjectSchemaTypePageScheme, *model.ResponseScheme, error) {
	return o.internalClient.ObjectTypes(ctx, workspaceID, objectSchemaID, excludeAbstract)
}

type internalObjectSchemaImpl struct {
	c service.Client
}

func (i *internalObjectSchemaImpl) List(ctx context.Context, workspaceID string) (*model.ObjectSchemaPageScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/list", workspaceID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ObjectSchemaPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalObjectSchemaImpl) Create(ctx context.Context, workspaceID string, payload *model.ObjectSchemaPayloadScheme) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/create", workspaceID)

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	schema := new(model.ObjectSchemaScheme)
	response, err := i.c.Call(request, schema)
	if err != nil {
		return nil, response, err
	}

	return schema, response, nil
}

func (i *internalObjectSchemaImpl) Get(ctx context.Context, workspaceID, objectSchemaID string) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectSchemaID == "" {
		return nil, nil, model.ErrNoObjectSchemaIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/%v", workspaceID, objectSchemaID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	schema := new(model.ObjectSchemaScheme)
	response, err := i.c.Call(request, schema)
	if err != nil {
		return nil, response, err
	}

	return schema, response, nil
}

func (i *internalObjectSchemaImpl) Update(ctx context.Context, workspaceID, objectSchemaID string, payload *model.ObjectSchemaPayloadScheme) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectSchemaID == "" {
		return nil, nil, model.ErrNoObjectSchemaIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/%v", workspaceID, objectSchemaID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	schema := new(model.ObjectSchemaScheme)
	response, err := i.c.Call(request, schema)
	if err != nil {
		return nil, response, err
	}

	return schema, response, nil
}

func (i *internalObjectSchemaImpl) Delete(ctx context.Context, workspaceID, objectSchemaID string) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectSchemaID == "" {
		return nil, nil, model.ErrNoObjectSchemaIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/%v", workspaceID, objectSchemaID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	schema := new(model.ObjectSchemaScheme)
	response, err := i.c.Call(request, schema)
	if err != nil {
		return nil, response, err
	}

	return schema, response, nil
}

func (i *internalObjectSchemaImpl) Attributes(ctx context.Context, workspaceID, objectSchemaID string, options *model.ObjectSchemaAttributesParamsScheme) ([]*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectSchemaID == "" {
		return nil, nil, model.ErrNoObjectSchemaIDError
	}

	query := url.Values{}
	if options != nil {

		if options.OnlyValueEditable {
			query.Add("onlyValueEditable", "true")
		}

		if options.OnlyValueEditable {
			query.Add("onlyValueEditable", "true")
		}

		if options.Extended {
			query.Add("extended", "true")
		}

		if options.Query != "" {
			query.Add("query", options.Query)
		}
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/%v/attributes", workspaceID, objectSchemaID))

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	var attributes []*model.ObjectTypeAttributeScheme
	response, err := i.c.Call(request, &attributes)
	if err != nil {
		return nil, response, err
	}

	return attributes, response, nil
}

func (i *internalObjectSchemaImpl) ObjectTypes(ctx context.Context, workspaceID, objectSchemaID string, excludeAbstract bool) (*model.ObjectSchemaTypePageScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectSchemaID == "" {
		return nil, nil, model.ErrNoObjectSchemaIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/%v/objecttypes", workspaceID, objectSchemaID))

	if excludeAbstract {
		query := url.Values{}
		query.Add("excludeAbstract", "true")

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ObjectSchemaTypePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}
