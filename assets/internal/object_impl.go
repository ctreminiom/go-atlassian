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

func NewObjectService(client service.Client) *ObjectService {

	return &ObjectService{
		internalClient: &internalObjectImpl{c: client},
	}
}

type ObjectService struct {
	internalClient assets.ObjectConnector
}

// Get loads one object.
//
// GET /jsm/assets/workspace/{workspaceId}/v1/object/{id}
func (o *ObjectService) Get(ctx context.Context, workspaceID, objectID string) (*model.ObjectScheme, *model.ResponseScheme, error) {
	return o.internalClient.Get(ctx, workspaceID, objectID)
}

// Update updates an existing object in Assets.
//
// PUT /jsm/assets/workspace/{workspaceId}/v1/object/{id}
func (o *ObjectService) Update(ctx context.Context, workspaceID, objectID string, payload *model.ObjectPayloadScheme) (*model.ObjectScheme, *model.ResponseScheme, error) {
	return o.internalClient.Update(ctx, workspaceID, objectID, payload)
}

// Delete deletes the referenced object
//
// DELETE /jsm/assets/workspace/{workspaceId}/v1/object/{id}
func (o *ObjectService) Delete(ctx context.Context, workspaceID, objectID string) (*model.ResponseScheme, error) {
	return o.internalClient.Delete(ctx, workspaceID, objectID)
}

// Attributes list all attributes for the given object.
//
// GET /jsm/assets/workspace/{workspaceId}/v1/object/{id}/attributes
func (o *ObjectService) Attributes(ctx context.Context, workspaceID, objectID string) ([]*model.ObjectAttributeScheme, *model.ResponseScheme, error) {
	return o.internalClient.Attributes(ctx, workspaceID, objectID)
}

// History retrieves the history entries for this object.
//
// GET /jsm/assets/workspace/{workspaceId}/v1/object/{id}/history
func (o *ObjectService) History(ctx context.Context, workspaceID, objectID string, ascOrder bool) ([]*model.ObjectHistoryScheme, *model.ResponseScheme, error) {
	return o.internalClient.History(ctx, workspaceID, objectID, ascOrder)
}

// References finds all references for an object.
//
// GET /jsm/assets/workspace/{workspaceId}/v1/object/{id}/referenceinfo
func (o *ObjectService) References(ctx context.Context, workspaceID, objectID string) ([]*model.ObjectReferenceTypeInfoScheme, *model.ResponseScheme, error) {
	return o.internalClient.References(ctx, workspaceID, objectID)
}

// Create creates a new object in Assets.
//
// POST /jsm/assets/workspace/{workspaceId}/v1/object/create
func (o *ObjectService) Create(ctx context.Context, workspaceID string, payload *model.ObjectPayloadScheme) (*model.ObjectScheme, *model.ResponseScheme, error) {
	return o.internalClient.Create(ctx, workspaceID, payload)
}

// Relation returns the relation between Jira issues and Assets objects
//
// GET /jsm/assets/workspace/{workspaceId}/v1/objectconnectedtickets/{objectId}/tickets
func (o *ObjectService) Relation(ctx context.Context, workspaceID, objectID string) (*model.TicketPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.Relation(ctx, workspaceID, objectID)
}

type internalObjectImpl struct {
	c service.Client
}

func (i *internalObjectImpl) Get(ctx context.Context, workspaceID, objectID string) (*model.ObjectScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectID == "" {
		return nil, nil, model.ErrNoObjectIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/object/%v", workspaceID, objectID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	object := new(model.ObjectScheme)
	response, err := i.c.Call(request, object)
	if err != nil {
		return nil, response, err
	}

	return object, response, nil
}

func (i *internalObjectImpl) Update(ctx context.Context, workspaceID, objectID string, payload *model.ObjectPayloadScheme) (*model.ObjectScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectID == "" {
		return nil, nil, model.ErrNoObjectIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/object/%v", workspaceID, objectID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	object := new(model.ObjectScheme)
	response, err := i.c.Call(request, object)
	if err != nil {
		return nil, response, err
	}

	return object, response, nil
}

func (i *internalObjectImpl) Delete(ctx context.Context, workspaceID, objectID string) (*model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, model.ErrNoWorkspaceIDError
	}

	if objectID == "" {
		return nil, model.ErrNoObjectIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/object/%v", workspaceID, objectID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalObjectImpl) Attributes(ctx context.Context, workspaceID, objectID string) ([]*model.ObjectAttributeScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectID == "" {
		return nil, nil, model.ErrNoObjectIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/object/%v/attributes", workspaceID, objectID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var attributes []*model.ObjectAttributeScheme
	response, err := i.c.Call(request, &attributes)
	if err != nil {
		return nil, response, err
	}

	return attributes, response, nil
}

func (i *internalObjectImpl) History(ctx context.Context, workspaceID, objectID string, ascOrder bool) ([]*model.ObjectHistoryScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectID == "" {
		return nil, nil, model.ErrNoObjectIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("jsm/assets/workspace/%v/v1/object/%v/history", workspaceID, objectID))

	if ascOrder {

		query := url.Values{}
		query.Add("asc", "true")

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	var history []*model.ObjectHistoryScheme
	response, err := i.c.Call(request, &history)
	if err != nil {
		return nil, response, err
	}

	return history, response, nil
}

func (i *internalObjectImpl) References(ctx context.Context, workspaceID, objectID string) ([]*model.ObjectReferenceTypeInfoScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectID == "" {
		return nil, nil, model.ErrNoObjectIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/object/%v/referenceinfo", workspaceID, objectID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var references []*model.ObjectReferenceTypeInfoScheme
	response, err := i.c.Call(request, &references)
	if err != nil {
		return nil, response, err
	}

	return references, response, nil
}

func (i *internalObjectImpl) Create(ctx context.Context, workspaceID string, payload *model.ObjectPayloadScheme) (*model.ObjectScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/object/create", workspaceID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	object := new(model.ObjectScheme)
	response, err := i.c.Call(request, object)
	if err != nil {
		return nil, response, err
	}

	return object, response, nil
}

func (i *internalObjectImpl) Relation(ctx context.Context, workspaceID, objectID string) (*model.TicketPageScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if objectID == "" {
		return nil, nil, model.ErrNoObjectIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectconnectedtickets/%v/tickets", workspaceID, objectID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.TicketPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}
