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

// NewPropertyService creates a new instance of PropertyService.
// It takes a service.Connector as input and returns a pointer to PropertyService.
func NewPropertyService(client service.Connector) *PropertyService {
	return &PropertyService{
		internalClient: &internalPropertyImpl{c: client},
	}
}

// PropertyService provides methods to interact with content property operations in Confluence.
type PropertyService struct {
	// internalClient is the connector interface for content property operations.
	internalClient confluence.ContentPropertyConnector
}

// Gets returns the properties for a piece of content.
//
// GET /wiki/rest/api/content/{id}/property
//
// https://docs.go-atlassian.io/confluence-cloud/content/properties#get-content-properties
func (p *PropertyService) Gets(ctx context.Context, contentID string, expand []string, startAt, maxResults int) (*model.ContentPropertyPageScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx, contentID, expand, startAt, maxResults)
}

// Create creates a property for an existing piece of content.
//
// POST /wiki/rest/api/content/{id}/property
//
// https://docs.go-atlassian.io/confluence-cloud/content/properties#create-content-property
func (p *PropertyService) Create(ctx context.Context, contentID string, payload *model.ContentPropertyPayloadScheme) (*model.ContentPropertyScheme, *model.ResponseScheme, error) {
	return p.internalClient.Create(ctx, contentID, payload)
}

// Get returns a content property for a piece of content.
//
// GET /wiki/rest/api/content/{id}/property/{key}
//
// https://docs.go-atlassian.io/confluence-cloud/content/properties#get-content-property
func (p *PropertyService) Get(ctx context.Context, contentID, key string) (*model.ContentPropertyScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, contentID, key)
}

// Delete deletes a content property.
//
// DELETE /wiki/rest/api/content/{id}/property/{key}
//
// https://docs.go-atlassian.io/confluence-cloud/content/properties#delete-content-property
func (p *PropertyService) Delete(ctx context.Context, contentID, key string) (*model.ResponseScheme, error) {
	return p.internalClient.Delete(ctx, contentID, key)
}

type internalPropertyImpl struct {
	c service.Connector
}

func (i *internalPropertyImpl) Gets(ctx context.Context, contentID string, expand []string, startAt, maxResults int) (*model.ContentPropertyPageScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/property?%v", contentID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ContentPropertyPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalPropertyImpl) Create(ctx context.Context, contentID string, payload *model.ContentPropertyPayloadScheme) (*model.ContentPropertyScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/property", contentID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	property := new(model.ContentPropertyScheme)
	response, err := i.c.Call(request, property)
	if err != nil {
		return nil, response, err
	}

	return property, response, nil
}

func (i *internalPropertyImpl) Get(ctx context.Context, contentID, key string) (*model.ContentPropertyScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	if key == "" {
		return nil, nil, model.ErrNoContentProperty
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/property/%v", contentID, key)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	property := new(model.ContentPropertyScheme)
	response, err := i.c.Call(request, property)
	if err != nil {
		return nil, response, err
	}

	return property, response, nil
}

func (i *internalPropertyImpl) Delete(ctx context.Context, contentID, key string) (*model.ResponseScheme, error) {

	if contentID == "" {
		return nil, model.ErrNoContentID
	}

	if key == "" {
		return nil, model.ErrNoContentProperty
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/property/%v", contentID, key)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
