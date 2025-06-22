package internal

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// NewPageService creates a new instance of PageService.
// It takes a service.Connector as input and returns a pointer to PageService.
func NewPageService(client service.Connector) *PageService {
	return &PageService{internalClient: &internalPageImpl{c: client}}
}

// PageService provides methods to interact with page operations in Confluence.
type PageService struct {
	// internalClient is the connector interface for page operations.
	internalClient confluence.PageConnector
}

// Get returns a specific page.
//
// GET /wiki/api/v2/pages/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/v2/page#get-page-by-id
func (p *PageService) Get(ctx context.Context, pageID int, format string, draft bool, version int) (*model.PageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*PageService).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	return p.internalClient.Get(ctx, pageID, format, draft, version)
}

// Gets returns all pages.
//
// GET /wiki/api/v2/pages
//
// https://docs.go-atlassian.io/confluence-cloud/v2/page#get-pages
func (p *PageService) Gets(ctx context.Context, options *model.PageOptionsScheme, cursor string, limit int) (*model.PageChunkScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*PageService).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets"))

	return p.internalClient.Gets(ctx, options, cursor, limit)
}

// Bulk returns all pages.
//
// Deprecated. Please use Page.Gets() instead.
//
// GET /wiki/api/v2/pages
//
// https://docs.go-atlassian.io/confluence-cloud/v2/page#get-pages
func (p *PageService) Bulk(ctx context.Context, cursor string, limit int) (*model.PageChunkScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*PageService).Bulk", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "bulk"))

	return p.internalClient.Bulk(ctx, cursor, limit)
}

// GetsByLabel returns the pages of specified label.
//
// GET /wiki/api/v2/labels/{id}/pages
//
// https://docs.go-atlassian.io/confluence-cloud/v2/page#get-pages-for-label
func (p *PageService) GetsByLabel(ctx context.Context, labelID int, sort, cursor string, limit int) (*model.PageChunkScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*PageService).GetsByLabel", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets_by_label"))

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
	ctx, span := tracer().Start(ctx, "(*PageService).GetsBySpace", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets_by_space"))

	return p.internalClient.GetsBySpace(ctx, spaceID, cursor, limit)
}

// GetsByParent returns all children of a page.
//
// The number of results is limited by the limit parameter and additional results (if available)
//
// will be available through the next cursor
//
// GET /wiki/api/v2/pages/{id}/children
//
// https://docs.go-atlassian.io/confluence-cloud/v2/page#get-pages-by-parent
func (p *PageService) GetsByParent(ctx context.Context, pageID int, cursor string, limit int) (*model.ChildPageChunkScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*PageService).GetsByParent", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets_by_parent"))

	return p.internalClient.GetsByParent(ctx, pageID, cursor, limit)
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
	ctx, span := tracer().Start(ctx, "(*PageService).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "create"))

	return p.internalClient.Create(ctx, payload)
}

// Update updates a page by id.
//
// PUT /wiki/api/v2/pages/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/v2/page#update-page
func (p *PageService) Update(ctx context.Context, pageID int, payload *model.PageUpdatePayloadScheme) (*model.PageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*PageService).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update"))

	return p.internalClient.Update(ctx, pageID, payload)
}

// Delete deletes a page by id.
//
// DELETE /wiki/api/v2/pages/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/v2/page#delete-page
func (p *PageService) Delete(ctx context.Context, pageID int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*PageService).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete"))

	return p.internalClient.Delete(ctx, pageID)
}

type internalPageImpl struct {
	c service.Connector
}

func (i *internalPageImpl) Gets(ctx context.Context, options *model.PageOptionsScheme, cursor string, limit int) (*model.PageChunkScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalPageImpl).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets"))

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	if options != nil {

		if options.Title != "" {
			query.Add("title", options.Title)
		}

		if options.Sort != "" {
			query.Add("sort", options.Sort)
		}

		if options.BodyFormat != "" {
			query.Add("body-format", options.BodyFormat)
		}

		if options.Status != nil {
			query.Add("status", strings.Join(options.Status, ","))
		}

		if len(options.PageIDs) > 0 {

			var pageIDs = make([]string, 0, len(options.PageIDs))
			for _, pageIDAsInt := range options.PageIDs {
				pageIDs = append(pageIDs, strconv.Itoa(pageIDAsInt))
			}

			query.Add("id", strings.Join(pageIDs, ","))
		}

		if len(options.SpaceIDs) > 0 {

			var spaceIDs = make([]string, 0, len(options.SpaceIDs))
			for _, spaceIDAsInt := range options.SpaceIDs {
				spaceIDs = append(spaceIDs, strconv.Itoa(spaceIDAsInt))
			}

			query.Add("space-id", strings.Join(spaceIDs, ","))
		}
	}

	endpoint := fmt.Sprintf("wiki/api/v2/pages?%v", query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	chunk := new(model.PageChunkScheme)
	response, err := i.c.Call(request, chunk)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return chunk, response, nil
}

func (i *internalPageImpl) Get(ctx context.Context, pageID int, format string, draft bool, version int) (*model.PageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalPageImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	if pageID == 0 {

			return nil, nil, fmt.Errorf("confluence: %w", model.ErrNoPageID)
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

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	page := new(model.PageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalPageImpl) Bulk(ctx context.Context, cursor string, limit int) (*model.PageChunkScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalPageImpl).Bulk", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "bulk"))

	return i.Gets(ctx, nil, cursor, limit)
}

func (i *internalPageImpl) GetsByLabel(ctx context.Context, labelID int, sort, cursor string, limit int) (*model.PageChunkScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalPageImpl).GetsByLabel", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets_by_label"))

	if labelID == 0 {

			return nil, nil, fmt.Errorf("confluence: %w", model.ErrNoLabelID)
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

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	chunk := new(model.PageChunkScheme)
	response, err := i.c.Call(request, chunk)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return chunk, response, nil
}

func (i *internalPageImpl) GetsBySpace(ctx context.Context, spaceID int, cursor string, limit int) (*model.PageChunkScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalPageImpl).GetsBySpace", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets_by_space"))

	if spaceID == 0 {

			return nil, nil, fmt.Errorf("confluence: %w", model.ErrNoSpaceID)
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/spaces/%v/pages?%v", spaceID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	chunk := new(model.PageChunkScheme)
	response, err := i.c.Call(request, chunk)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return chunk, response, nil
}

func (i *internalPageImpl) GetsByParent(ctx context.Context, parentID int, cursor string, limit int) (*model.ChildPageChunkScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalPageImpl).GetsByParent", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets_by_parent"))

	if parentID == 0 {

			return nil, nil, fmt.Errorf("confluence: %w", model.ErrNoPageID)
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/pages/%v/children?%v", parentID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	chunk := new(model.ChildPageChunkScheme)
	response, err := i.c.Call(request, chunk)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return chunk, response, nil
}

func (i *internalPageImpl) Create(ctx context.Context, payload *model.PageCreatePayloadScheme) (*model.PageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalPageImpl).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "create"))

	endpoint := "wiki/api/v2/pages"

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	page := new(model.PageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalPageImpl) Update(ctx context.Context, pageID int, payload *model.PageUpdatePayloadScheme) (*model.PageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalPageImpl).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update"))

	if pageID == 0 {

			return nil, nil, fmt.Errorf("confluence: %w", model.ErrNoPageID)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/pages/%v", pageID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	page := new(model.PageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalPageImpl) Delete(ctx context.Context, pageID int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalPageImpl).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete"))

	if pageID == 0 {
		return nil, fmt.Errorf("confluence: %w", model.ErrNoPageID)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/pages/%v", pageID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	return i.c.Call(request, nil)
}
