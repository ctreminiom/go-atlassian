package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/assets"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// NewObjectService creates a new instance of ObjectService.
// It takes a service.Connector as input and returns a pointer to ObjectService.
func NewObjectService(client service.Connector) *ObjectService {
	return &ObjectService{
		internalClient: &internalObjectImpl{c: client},
	}
}

// ObjectService provides methods to interact with objects in Jira Assets.
type ObjectService struct {
	// internalClient is the connector interface for object operations.
	internalClient assets.ObjectConnector
}

// Get loads one object.
//
// GET /jsm/assets/workspace/{workspaceId}/v1/object/{id}
//
// https://docs.go-atlassian.io/jira-assets/object#get-object-by-id
func (o *ObjectService) Get(ctx context.Context, workspaceID, objectID string) (*model.ObjectScheme, *model.ResponseScheme, error) {
	return o.internalClient.Get(ctx, workspaceID, objectID)
}

// Update updates an existing object in Assets.
//
// PUT /jsm/assets/workspace/{workspaceId}/v1/object/{id}
//
// https://docs.go-atlassian.io/jira-assets/object#update-object-by-id
func (o *ObjectService) Update(ctx context.Context, workspaceID, objectID string, payload *model.ObjectPayloadScheme) (*model.ObjectScheme, *model.ResponseScheme, error) {
	return o.internalClient.Update(ctx, workspaceID, objectID, payload)
}

// Delete deletes the referenced object
//
// DELETE /jsm/assets/workspace/{workspaceId}/v1/object/{id}
//
// https://docs.go-atlassian.io/jira-assets/object#delete-object-by-id
func (o *ObjectService) Delete(ctx context.Context, workspaceID, objectID string) (*model.ResponseScheme, error) {
	return o.internalClient.Delete(ctx, workspaceID, objectID)
}

// Attributes list all attributes for the given object.
//
// GET /jsm/assets/workspace/{workspaceId}/v1/object/{id}/attributes
//
// https://docs.go-atlassian.io/jira-assets/object#get-object-attributes
func (o *ObjectService) Attributes(ctx context.Context, workspaceID, objectID string) ([]*model.ObjectAttributeScheme, *model.ResponseScheme, error) {
	return o.internalClient.Attributes(ctx, workspaceID, objectID)
}

// History retrieves the history entries for this object.
//
// GET /jsm/assets/workspace/{workspaceId}/v1/object/{id}/history
//
// https://docs.go-atlassian.io/jira-assets/object#get-object-changelogs
func (o *ObjectService) History(ctx context.Context, workspaceID, objectID string, ascOrder bool) ([]*model.ObjectHistoryScheme, *model.ResponseScheme, error) {
	return o.internalClient.History(ctx, workspaceID, objectID, ascOrder)
}

// References finds all references for an object.
//
// GET /jsm/assets/workspace/{workspaceId}/v1/object/{id}/referenceinfo
//
// https://docs.go-atlassian.io/jira-assets/object#get-object-references
func (o *ObjectService) References(ctx context.Context, workspaceID, objectID string) ([]*model.ObjectReferenceTypeInfoScheme, *model.ResponseScheme, error) {
	return o.internalClient.References(ctx, workspaceID, objectID)
}

// Create creates a new object in Assets.
//
// POST /jsm/assets/workspace/{workspaceId}/v1/object/create
//
// https://docs.go-atlassian.io/jira-assets/object#create-object
func (o *ObjectService) Create(ctx context.Context, workspaceID string, payload *model.ObjectPayloadScheme) (*model.ObjectScheme, *model.ResponseScheme, error) {
	return o.internalClient.Create(ctx, workspaceID, payload)
}

// Relation returns the relation between Jira issues and Assets objects
//
// GET /jsm/assets/workspace/{workspaceId}/v1/objectconnectedtickets/{objectId}/tickets
//
// https://docs.go-atlassian.io/jira-assets/object#get-object-tickets
func (o *ObjectService) Relation(ctx context.Context, workspaceID, objectID string) (*model.TicketPageScheme, *model.ResponseScheme, error) {
	return o.internalClient.Relation(ctx, workspaceID, objectID)
}

// Filter fetch Objects by AQL.
//
// POST /jsm/assets/workspace/{workspaceId}/v1/object/aql
//
// https://docs.go-atlassian.io/jira-assets/object#filter-objects
func (o *ObjectService) Filter(ctx context.Context, workspaceID, aql string, attributes bool, startAt, maxResults int) (*model.ObjectListResultScheme, *model.ResponseScheme, error) {
	return o.internalClient.Filter(ctx, workspaceID, aql, attributes, startAt, maxResults)
}

// Search retrieve a list of objects based on an AQL.
//
// Note that the preferred endpoint is /aql
//
// POST /jsm/assets/workspace/{workspaceId}/v1/object/navlist/aql
//
// https://docs.go-atlassian.io/jira-assets/object#search-objects
func (o *ObjectService) Search(ctx context.Context, workspaceID string, payload *model.ObjectSearchParamsScheme) (*model.ObjectListScheme, *model.ResponseScheme, error) {
	return o.internalClient.Search(ctx, workspaceID, payload)
}

type internalObjectImpl struct {
	c service.Connector
}

func (i *internalObjectImpl) Search(ctx context.Context, workspaceID string, payload *model.ObjectSearchParamsScheme) (*model.ObjectListScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/object/navlist/aql", workspaceID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	object := new(model.ObjectListScheme)
	res, err := i.c.Call(req, object)
	if err != nil {
		return nil, res, err
	}

	return object, res, nil

}

func (i *internalObjectImpl) Filter(ctx context.Context, workspaceID, aql string, attributes bool, startAt, maxResults int) (*model.ObjectListResultScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if aql == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoAqlQuery)
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if !attributes {
		params.Add("includeAttributes", "false")
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/object/aql?%v", workspaceID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"qlQuery": aql})
	if err != nil {
		return nil, nil, err
	}

	list := new(model.ObjectListResultScheme)
	res, err := i.c.Call(req, list)
	if err != nil {
		return nil, res, err
	}

	return list, res, nil
}

func (i *internalObjectImpl) Get(ctx context.Context, workspaceID, objectID string) (*model.ObjectScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoObjectID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/object/%v", workspaceID, objectID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	object := new(model.ObjectScheme)
	res, err := i.c.Call(req, object)
	if err != nil {
		return nil, res, err
	}

	return object, res, nil
}

func (i *internalObjectImpl) Update(ctx context.Context, workspaceID, objectID string, payload *model.ObjectPayloadScheme) (*model.ObjectScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoObjectID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/object/%v", workspaceID, objectID)

	req, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	object := new(model.ObjectScheme)
	res, err := i.c.Call(req, object)
	if err != nil {
		return nil, res, err
	}

	return object, res, nil
}

func (i *internalObjectImpl) Delete(ctx context.Context, workspaceID, objectID string) (*model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectID == "" {
		return nil, fmt.Errorf("assets: %w", model.ErrNoObjectID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/object/%v", workspaceID, objectID)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}

func (i *internalObjectImpl) Attributes(ctx context.Context, workspaceID, objectID string) ([]*model.ObjectAttributeScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoObjectID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/object/%v/attributes", workspaceID, objectID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var attributes []*model.ObjectAttributeScheme
	res, err := i.c.Call(req, &attributes)
	if err != nil {
		return nil, res, err
	}

	return attributes, res, nil
}

func (i *internalObjectImpl) History(ctx context.Context, workspaceID, objectID string, ascOrder bool) ([]*model.ObjectHistoryScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoObjectID)
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("jsm/assets/workspace/%v/v1/object/%v/history", workspaceID, objectID))

	if ascOrder {

		query := url.Values{}
		query.Add("asc", "true")

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	var history []*model.ObjectHistoryScheme
	res, err := i.c.Call(req, &history)
	if err != nil {
		return nil, res, err
	}

	return history, res, nil
}

func (i *internalObjectImpl) References(ctx context.Context, workspaceID, objectID string) ([]*model.ObjectReferenceTypeInfoScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoObjectID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/object/%v/referenceinfo", workspaceID, objectID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var references []*model.ObjectReferenceTypeInfoScheme
	res, err := i.c.Call(req, &references)
	if err != nil {
		return nil, res, err
	}

	return references, res, nil
}

func (i *internalObjectImpl) Create(ctx context.Context, workspaceID string, payload *model.ObjectPayloadScheme) (*model.ObjectScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/object/create", workspaceID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	object := new(model.ObjectScheme)
	res, err := i.c.Call(req, object)
	if err != nil {
		return nil, res, err
	}

	return object, res, nil
}

func (i *internalObjectImpl) Relation(ctx context.Context, workspaceID, objectID string) (*model.TicketPageScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if objectID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoObjectID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/objectconnectedtickets/%v/tickets", workspaceID, objectID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.TicketPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}
