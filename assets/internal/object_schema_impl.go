package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/assets"
	"net/http"
	"net/url"
	"strings"
)

// NewObjectSchemaService creates a new instance of ObjectSchemaService.
// It takes a service.Connector as input and returns a pointer to ObjectSchemaService.
func NewObjectSchemaService(client service.Connector) *ObjectSchemaService {
	return &ObjectSchemaService{
		internalClient: &internalObjectSchemaImpl{c: client},
	}
}

// ObjectSchemaService provides methods to interact with object schemas in Jira Assets.
type ObjectSchemaService struct {
	// internalClient is the connector interface for object schema operations.
	internalClient assets.ObjectSchemaConnector
}

// List returns all the object schemas available on Assets
//
// GET /jsm/assets/workspace/{workspaceId}/v1/objectschema/list
//
// https://docs.go-atlassian.io/jira-assets/object/schema#get-object-schema-list
func (o *ObjectSchemaService) List(ctx context.Context, workspaceID string) (*model.ObjectSchemaPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.List(ctx, workspaceID)
}

// Create creates a new object schema
//
// POST /jsm/assets/workspace/{workspaceId}/v1/objectschema/create
//
// https://docs.go-atlassian.io/jira-assets/object/schema#create-object-schema
func (o *ObjectSchemaService) Create(ctx context.Context, workspaceID string, payload *model.ObjectSchemaPayloadScheme) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	return o.internalClient.Create(ctx, workspaceID, payload)
}

// Get returns an object schema by ID
//
// GET /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}
//
// https://docs.go-atlassian.io/jira-assets/object/schema#get-object-schema
func (o *ObjectSchemaService) Get(ctx context.Context, workspaceID, objectSchemaID string) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	return o.internalClient.Get(ctx, workspaceID, objectSchemaID)
}

// Update updates an object schema
//
// PUT /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}
//
// https://docs.go-atlassian.io/jira-assets/object/schema#update-object-schema
func (o *ObjectSchemaService) Update(ctx context.Context, workspaceID, objectSchemaID string, payload *model.ObjectSchemaPayloadScheme) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	return o.internalClient.Update(ctx, workspaceID, objectSchemaID, payload)
}

// Delete deletes a schema
//
// DELETE /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}
//
// https://docs.go-atlassian.io/jira-assets/object/schema#delete-object-schema
func (o *ObjectSchemaService) Delete(ctx context.Context, workspaceID, objectSchemaID string) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	return o.internalClient.Delete(ctx, workspaceID, objectSchemaID)
}

// Attributes finds all object type attributes for this object schema
//
// GET /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}/attributes
//
// https://docs.go-atlassian.io/jira-assets/object/schema#get-object-schema-attributes
func (o *ObjectSchemaService) Attributes(ctx context.Context, workspaceID, objectSchemaID string, options *model.ObjectSchemaAttributesParamsScheme) ([]*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {
	return o.internalClient.Attributes(ctx, workspaceID, objectSchemaID, options)
}

// ObjectTypes returns all object types for this object schema
//
// GET /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}/objecttypes
//
// https://docs.go-atlassian.io/jira-assets/object/schema#get-object-schema-types
func (o *ObjectSchemaService) ObjectTypes(ctx context.Context, workspaceID, objectSchemaID string, excludeAbstract bool) ([]*model.ObjectTypeScheme, *model.ResponseScheme, error) {
	return o.internalClient.ObjectTypes(ctx, workspaceID, objectSchemaID, excludeAbstract)
}

type internalObjectSchemaImpl struct {
	c service.Connector
}

func (i *internalObjectSchemaImpl) List(ctx context.Context, workspaceID string) (*model.ObjectSchemaPageScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/list", workspaceID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ObjectSchemaPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalObjectSchemaImpl) Create(ctx context.Context, workspaceID string, payload *model.ObjectSchemaPayloadScheme) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/create", workspaceID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	schema := new(model.ObjectSchemaScheme)
	res, err := i.c.Call(req, schema)
	if err != nil {
		return nil, res, err
	}

	return schema, res, nil
}

func (i *internalObjectSchemaImpl) Get(ctx context.Context, workspaceID, objectSchemaID string) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectSchemaID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoObjectSchemaID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/%v", workspaceID, objectSchemaID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	schema := new(model.ObjectSchemaScheme)
	res, err := i.c.Call(req, schema)
	if err != nil {
		return nil, res, err
	}

	return schema, res, nil
}

func (i *internalObjectSchemaImpl) Update(ctx context.Context, workspaceID, objectSchemaID string, payload *model.ObjectSchemaPayloadScheme) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectSchemaID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoObjectSchemaID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/%v", workspaceID, objectSchemaID)

	req, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	schema := new(model.ObjectSchemaScheme)
	res, err := i.c.Call(req, schema)
	if err != nil {
		return nil, res, err
	}

	return schema, res, nil
}

func (i *internalObjectSchemaImpl) Delete(ctx context.Context, workspaceID, objectSchemaID string) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectSchemaID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoObjectSchemaID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/%v", workspaceID, objectSchemaID)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	schema := new(model.ObjectSchemaScheme)
	res, err := i.c.Call(req, schema)
	if err != nil {
		return nil, res, err
	}

	return schema, res, nil
}

func (i *internalObjectSchemaImpl) Attributes(ctx context.Context, workspaceID, objectSchemaID string, options *model.ObjectSchemaAttributesParamsScheme) ([]*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectSchemaID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoObjectSchemaID)
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
	fmt.Fprintf(&endpoint, "jsm/assets/workspace/%v/v1/objectschema/%v/attributes", workspaceID, objectSchemaID)

	if query.Encode() != "" {
		fmt.Fprintf(&endpoint, "?%v", query.Encode())
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

func (i *internalObjectSchemaImpl) ObjectTypes(ctx context.Context, workspaceID, objectSchemaID string, excludeAbstract bool) ([]*model.ObjectTypeScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectSchemaID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoObjectSchemaID)
	}

	var endpoint strings.Builder
	fmt.Fprintf(&endpoint, "jsm/assets/workspace/%v/v1/objectschema/%v/objecttypes", workspaceID, objectSchemaID)

	if excludeAbstract {
		query := url.Values{}
		query.Add("excludeAbstract", "true")

		fmt.Fprintf(&endpoint, "?%v", query.Encode())
	}

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	var objectTypes []*model.ObjectTypeScheme
	res, err := i.c.Call(req, &objectTypes)
	if err != nil {
		return nil, res, err
	}

	return objectTypes, res, nil
}
