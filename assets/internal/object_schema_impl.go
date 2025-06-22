package internal

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/assets"
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
	ctx, span := tracer().Start(ctx, "(*ObjectSchemaService).List", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "list"))

	return o.internalClient.List(ctx, workspaceID)
}

// Create creates a new object schema
//
// POST /jsm/assets/workspace/{workspaceId}/v1/objectschema/create
//
// https://docs.go-atlassian.io/jira-assets/object/schema#create-object-schema
func (o *ObjectSchemaService) Create(ctx context.Context, workspaceID string, payload *model.ObjectSchemaPayloadScheme) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ObjectSchemaService).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "create"))

	return o.internalClient.Create(ctx, workspaceID, payload)
}

// Get returns an object schema by ID
//
// GET /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}
//
// https://docs.go-atlassian.io/jira-assets/object/schema#get-object-schema
func (o *ObjectSchemaService) Get(ctx context.Context, workspaceID, objectSchemaID string) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ObjectSchemaService).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	return o.internalClient.Get(ctx, workspaceID, objectSchemaID)
}

// Update updates an object schema
//
// PUT /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}
//
// https://docs.go-atlassian.io/jira-assets/object/schema#update-object-schema
func (o *ObjectSchemaService) Update(ctx context.Context, workspaceID, objectSchemaID string, payload *model.ObjectSchemaPayloadScheme) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ObjectSchemaService).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update"))

	return o.internalClient.Update(ctx, workspaceID, objectSchemaID, payload)
}

// Delete deletes a schema
//
// DELETE /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}
//
// https://docs.go-atlassian.io/jira-assets/object/schema#delete-object-schema
func (o *ObjectSchemaService) Delete(ctx context.Context, workspaceID, objectSchemaID string) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ObjectSchemaService).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete"))

	return o.internalClient.Delete(ctx, workspaceID, objectSchemaID)
}

// Attributes finds all object type attributes for this object schema
//
// GET /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}/attributes
//
// https://docs.go-atlassian.io/jira-assets/object/schema#get-object-schema-attributes
func (o *ObjectSchemaService) Attributes(ctx context.Context, workspaceID, objectSchemaID string, options *model.ObjectSchemaAttributesParamsScheme) ([]*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ObjectSchemaService).Attributes", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "attributes"))

	return o.internalClient.Attributes(ctx, workspaceID, objectSchemaID, options)
}

// ObjectTypes returns all object types for this object schema
//
// GET /jsm/assets/workspace/{workspaceId}/v1/objectschema/{id}/objecttypes
//
// https://docs.go-atlassian.io/jira-assets/object/schema#get-object-schema-types
func (o *ObjectSchemaService) ObjectTypes(ctx context.Context, workspaceID, objectSchemaID string, excludeAbstract bool) ([]*model.ObjectTypeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ObjectSchemaService).ObjectTypes", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "object_types"))

	return o.internalClient.ObjectTypes(ctx, workspaceID, objectSchemaID, excludeAbstract)
}

type internalObjectSchemaImpl struct {
	c service.Connector
}

func (i *internalObjectSchemaImpl) List(ctx context.Context, workspaceID string) (*model.ObjectSchemaPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalObjectSchemaImpl).List", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "list"))

	if workspaceID == "" {

			return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/list", workspaceID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	page := new(model.ObjectSchemaPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return page, res, nil
}

func (i *internalObjectSchemaImpl) Create(ctx context.Context, workspaceID string, payload *model.ObjectSchemaPayloadScheme) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalObjectSchemaImpl).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "create"))

	if workspaceID == "" {

			return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/create", workspaceID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	schema := new(model.ObjectSchemaScheme)
	res, err := i.c.Call(req, schema)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return schema, res, nil
}

func (i *internalObjectSchemaImpl) Get(ctx context.Context, workspaceID, objectSchemaID string) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalObjectSchemaImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	if workspaceID == "" {

			return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectSchemaID == "" {

			return nil, nil, fmt.Errorf("assets: %w", model.ErrNoObjectSchemaID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/%v", workspaceID, objectSchemaID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	schema := new(model.ObjectSchemaScheme)
	res, err := i.c.Call(req, schema)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return schema, res, nil
}

func (i *internalObjectSchemaImpl) Update(ctx context.Context, workspaceID, objectSchemaID string, payload *model.ObjectSchemaPayloadScheme) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalObjectSchemaImpl).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update"))

	if workspaceID == "" {

			return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectSchemaID == "" {

			return nil, nil, fmt.Errorf("assets: %w", model.ErrNoObjectSchemaID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/%v", workspaceID, objectSchemaID)

	req, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	schema := new(model.ObjectSchemaScheme)
	res, err := i.c.Call(req, schema)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return schema, res, nil
}

func (i *internalObjectSchemaImpl) Delete(ctx context.Context, workspaceID, objectSchemaID string) (*model.ObjectSchemaScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalObjectSchemaImpl).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete"))

	if workspaceID == "" {

			return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectSchemaID == "" {

			return nil, nil, fmt.Errorf("assets: %w", model.ErrNoObjectSchemaID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/%v", workspaceID, objectSchemaID)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	schema := new(model.ObjectSchemaScheme)
	res, err := i.c.Call(req, schema)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return schema, res, nil
}

func (i *internalObjectSchemaImpl) Attributes(ctx context.Context, workspaceID, objectSchemaID string, options *model.ObjectSchemaAttributesParamsScheme) ([]*model.ObjectTypeAttributeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalObjectSchemaImpl).Attributes", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "attributes"))

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
	endpoint.WriteString(fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/%v/attributes", workspaceID, objectSchemaID))

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	var attributes []*model.ObjectTypeAttributeScheme
	res, err := i.c.Call(req, &attributes)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return attributes, res, nil
}

func (i *internalObjectSchemaImpl) ObjectTypes(ctx context.Context, workspaceID, objectSchemaID string, excludeAbstract bool) ([]*model.ObjectTypeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalObjectSchemaImpl).ObjectTypes", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "object_types"))

	if workspaceID == "" {

			return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectSchemaID == "" {

			return nil, nil, fmt.Errorf("assets: %w", model.ErrNoObjectSchemaID)
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("jsm/assets/workspace/%v/v1/objectschema/%v/objecttypes", workspaceID, objectSchemaID))

	if excludeAbstract {
		query := url.Values{}
		query.Add("excludeAbstract", "true")

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	var objectTypes []*model.ObjectTypeScheme
	res, err := i.c.Call(req, &objectTypes)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return objectTypes, res, nil
}
