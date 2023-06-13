package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/confluence"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// NewPageService returns a new Confluence V2 Page service
func NewPageService(client service.Client) *PageService {
	return &PageService{
		internalClient: &internalPageImpl{c: client},
	}
}

type PageService struct {
	internalClient confluence.PageConnector
}

// Get returns a specific page.
//
// GET /wiki/api/v2/pages/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/v2/page#get-page-by-id
func (p *PageService) Get(ctx context.Context, pageID int, format string, draft bool, version int) (*model.PageScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, pageID, format, draft, version)
}

// Bulk returns all pages.
//
// # The number of results is limited by the limit parameter and additional results
//
// (if available) will be available through the next cursor
//
// GET /wiki/api/v2/pages
//
// https://docs.go-atlassian.io/confluence-cloud/v2/page#get-pages
func (p *PageService) Bulk(ctx context.Context, cursor string, limit int) (*model.PageChunkScheme, *model.ResponseScheme, error) {
	return p.internalClient.Bulk(ctx, cursor, limit)
}

// BulkFiltered returns all pages that fit the filtering criteria
//
// # The number of results is limited by the limit parameter and additional results
//
// (if available) will be available through the next cursor
//
// GET /wiki/api/v2/pages
//
// https://docs.go-atlassian.io/confluence-cloud/v2/page#get-pages
func (p *PageService) BulkFiltered(ctx context.Context, status, format, cursor string, limit int, pageIDs ...int) (*model.PageChunkScheme, *model.ResponseScheme, error) {
	return p.internalClient.BulkFiltered(ctx, status, format, cursor, limit, pageIDs...)
}

// GetsByLabel returns the pages of specified label.
//
// # The number of results is limited by the limit parameter and additional results
//
// (if available) will be available through the next cursor
//
// GET /wiki/api/v2/labels/{id}/pages
//
// https://docs.go-atlassian.io/confluence-cloud/v2/page#get-pages-for-label
func (p *PageService) GetsByLabel(ctx context.Context, labelID int, sort, cursor string, limit int) (*model.PageChunkScheme, *model.ResponseScheme, error) {
	return p.internalClient.GetsByLabel(ctx, labelID, sort, cursor, limit)
}

// GetsBySpace returns all pages in a space.
//
// The number of results is limited by the limit parameter and additional results (if available)
//
// will be available through the next cursor
//
// GET /wiki/api/v2/spaces/{id}/pages
//
// https://docs.go-atlassian.io/confluence-cloud/v2/page#get-pages-in-space
func (p *PageService) GetsBySpace(ctx context.Context, spaceID int, cursor string, limit int) (*model.PageChunkScheme, *model.ResponseScheme, error) {
	return p.internalClient.GetsBySpace(ctx, spaceID, cursor, limit)
}

// Create creates a page in the space.
//
// Pages are created as published by default unless specified as a draft in the status field.
//
// If creating a published page, the title must be specified.
//
// POST /wiki/api/v2/pages
//
// https://docs.go-atlassian.io/confluence-cloud/v2/page#create-page
func (p *PageService) Create(ctx context.Context, payload *model.PageCreatePayloadScheme) (*model.PageScheme, *model.ResponseScheme, error) {
	return p.internalClient.Create(ctx, payload)
}

// Update updates a page by id.
//
// PUT /wiki/api/v2/pages/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/v2/page#update-page
func (p *PageService) Update(ctx context.Context, pageID int, payload *model.PageUpdatePayloadScheme) (*model.PageScheme, *model.ResponseScheme, error) {
	return p.internalClient.Update(ctx, pageID, payload)
}

// Delete deletes a page by id.
//
// DELETE /wiki/api/v2/pages/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/v2/page#delete-page
func (p *PageService) Delete(ctx context.Context, pageID int) (*model.ResponseScheme, error) {
	return p.internalClient.Delete(ctx, pageID)
}

type internalPageImpl struct {
	c service.Client
}

func (i *internalPageImpl) Get(ctx context.Context, pageID int, format string, draft bool, version int) (*model.PageScheme, *model.ResponseScheme, error) {

	if pageID == 0 {
		return nil, nil, model.ErrNoPageIDError
	}

	query := url.Values{}

	if format != "" {
		query.Add("body-format", format)
	}

	if draft {
		query.Add("get-draft", "true")
	}

	if version != 0 {
		query.Add("version", strconv.Itoa(version))
	}

	endpoint := fmt.Sprintf("wiki/api/v2/pages/%v?%v", pageID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.PageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalPageImpl) Bulk(ctx context.Context, cursor string, limit int) (*model.PageChunkScheme, *model.ResponseScheme, error) {
	return i.BulkFiltered(ctx, "", "", cursor, limit)
}

func (i *internalPageImpl) BulkFiltered(ctx context.Context, status, format, cursor string, limit int, pageIDs ...int) (*model.PageChunkScheme, *model.ResponseScheme, error) {

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))

	if status != "" {
		query.Add("status", status)
	}

	if format != "" {
		query.Add("body-format", format)
	}

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	if len(pageIDs) > 0 {
		ids := make([]string, 0, len(pageIDs))
		for _, id := range pageIDs {
			ids = append(ids, strconv.Itoa(id))
		}
		query.Add("id", strings.Join(ids, ","))
	}

	endpoint := fmt.Sprintf("wiki/api/v2/pages?%v", query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	chunk := new(model.PageChunkScheme)
	response, err := i.c.Call(request, chunk)
	if err != nil {
		return nil, response, err
	}

	return chunk, response, nil
}

func (i *internalPageImpl) GetsByLabel(ctx context.Context, labelID int, sort, cursor string, limit int) (*model.PageChunkScheme, *model.ResponseScheme, error) {

	if labelID == 0 {
		return nil, nil, model.ErrNoLabelIDError
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	if sort != "" {
		query.Add("sort", sort)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/labels/%v/pages?%v", labelID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	chunk := new(model.PageChunkScheme)
	response, err := i.c.Call(request, chunk)
	if err != nil {
		return nil, response, err
	}

	return chunk, response, nil
}

func (i *internalPageImpl) GetsBySpace(ctx context.Context, spaceID int, cursor string, limit int) (*model.PageChunkScheme, *model.ResponseScheme, error) {

	if spaceID == 0 {
		return nil, nil, model.ErrNoSpaceIDError
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/spaces/%v/pages?%v", spaceID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	chunk := new(model.PageChunkScheme)
	response, err := i.c.Call(request, chunk)
	if err != nil {
		return nil, response, err
	}

	return chunk, response, nil
}

func (i *internalPageImpl) Create(ctx context.Context, payload *model.PageCreatePayloadScheme) (*model.PageScheme, *model.ResponseScheme, error) {

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := "wiki/api/v2/pages"

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.PageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalPageImpl) Update(ctx context.Context, pageID int, payload *model.PageUpdatePayloadScheme) (*model.PageScheme, *model.ResponseScheme, error) {

	if pageID == 0 {
		return nil, nil, model.ErrNoPageIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("wiki/api/v2/pages/%v", pageID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.PageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalPageImpl) Delete(ctx context.Context, pageID int) (*model.ResponseScheme, error) {

	if pageID == 0 {
		return nil, model.ErrNoPageIDError
	}

	endpoint := fmt.Sprintf("wiki/api/v2/pages/%v", pageID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
