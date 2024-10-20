package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// NewRestrictionService creates a new instance of RestrictionService.
// It takes a service.Connector and a pointer to RestrictionOperationService as input and returns a pointer to RestrictionService.
func NewRestrictionService(client service.Connector, operation *RestrictionOperationService) *RestrictionService {
	return &RestrictionService{
		internalClient: &internalRestrictionImpl{c: client},
		Operation:      operation,
	}
}

// RestrictionService provides methods to interact with content restriction operations in Confluence.
type RestrictionService struct {
	// internalClient is the connector interface for content restriction operations.
	internalClient confluence.ContentRestrictionConnector
	// Operation is a pointer to RestrictionOperationService for additional restriction operations.
	Operation *RestrictionOperationService
}

// Gets returns the restrictions on a piece of content.
//
// GET /wiki/rest/api/content/{id}/restriction
//
// https://docs.go-atlassian.io/confluence-cloud/content/restrictions#get-restrictions
func (r *RestrictionService) Gets(ctx context.Context, contentID string, expand []string, startAt, maxResults int) (*model.ContentRestrictionPageScheme, *model.ResponseScheme, error) {
	return r.internalClient.Gets(ctx, contentID, expand, startAt, maxResults)
}

// Add adds restrictions to a piece of content. Note, this does not change any existing restrictions on the content.
//
// POST /wiki/rest/api/content/{id}/restriction
//
// https://docs.go-atlassian.io/confluence-cloud/content/restrictions#add-restrictions
func (r *RestrictionService) Add(ctx context.Context, contentID string, payload *model.ContentRestrictionUpdatePayloadScheme, expand []string) (*model.ContentRestrictionPageScheme, *model.ResponseScheme, error) {
	return r.internalClient.Add(ctx, contentID, payload, expand)
}

// Delete removes all restrictions (read and update) on a piece of content.
//
// DELETE /wiki/rest/api/content/{id}/restriction
//
// https://docs.go-atlassian.io/confluence-cloud/content/restrictions#delete-restrictions
func (r *RestrictionService) Delete(ctx context.Context, contentID string, expand []string) (*model.ContentRestrictionPageScheme, *model.ResponseScheme, error) {
	return r.internalClient.Delete(ctx, contentID, expand)
}

// Update updates restrictions for a piece of content. This removes the existing restrictions and replaces them with the restrictions in the request.
//
// PUT /wiki/rest/api/content/{id}/restriction
//
// https://docs.go-atlassian.io/confluence-cloud/content/restrictions#update-restrictions
func (r *RestrictionService) Update(ctx context.Context, contentID string, payload *model.ContentRestrictionUpdatePayloadScheme, expand []string) (*model.ContentRestrictionPageScheme, *model.ResponseScheme, error) {
	return r.internalClient.Update(ctx, contentID, payload, expand)
}

type internalRestrictionImpl struct {
	c service.Connector
}

func (i *internalRestrictionImpl) Gets(ctx context.Context, contentID string, expand []string, startAt, maxResults int) (*model.ContentRestrictionPageScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/restriction?%v", contentID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ContentRestrictionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalRestrictionImpl) Add(ctx context.Context, contentID string, payload *model.ContentRestrictionUpdatePayloadScheme, expand []string) (*model.ContentRestrictionPageScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/restriction", contentID))

	if len(expand) != 0 {
		query := url.Values{}
		query.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ContentRestrictionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalRestrictionImpl) Delete(ctx context.Context, contentID string, expand []string) (*model.ContentRestrictionPageScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/restriction", contentID))

	if len(expand) != 0 {
		query := url.Values{}
		query.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ContentRestrictionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalRestrictionImpl) Update(ctx context.Context, contentID string, payload *model.ContentRestrictionUpdatePayloadScheme, expand []string) (*model.ContentRestrictionPageScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/restriction", contentID))

	if len(expand) != 0 {
		query := url.Values{}
		query.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ContentRestrictionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}
