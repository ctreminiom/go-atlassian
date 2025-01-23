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

// ContentSubServices holds references to various sub-services related to content operations in Confluence.
type ContentSubServices struct {
	// Attachment is the service for content attachment operations.
	Attachment *ContentAttachmentService
	// ChildrenDescendant is the service for children and descendants operations.
	ChildrenDescendant *ChildrenDescandantsService
	// Comment is the service for comment operations.
	Comment *CommentService
	// Permission is the service for permission operations.
	Permission *PermissionService
	// Label is the service for content label operations.
	Label *ContentLabelService
	// Property is the service for content property operations.
	Property *PropertyService
	// Restriction is the service for content restriction operations.
	Restriction *RestrictionService
	// Version is the service for content version operations.
	Version *VersionService
}

// NewContentService creates a new instance of ContentService.
// It takes a service.Connector and a ContentSubServices as inputs and returns a pointer to ContentService.
func NewContentService(client service.Connector, subServices *ContentSubServices) *ContentService {
	return &ContentService{
		internalClient:     &internalContentImpl{c: client},
		Attachment:         subServices.Attachment,
		ChildrenDescendant: subServices.ChildrenDescendant,
		Comment:            subServices.Comment,
		Permission:         subServices.Permission,
		Label:              subServices.Label,
		Property:           subServices.Property,
		Restriction:        subServices.Restriction,
		Version:            subServices.Version,
	}
}

// ContentService provides methods to interact with content operations in Confluence.
type ContentService struct {
	// internalClient is the connector interface for content operations.
	internalClient confluence.ContentConnector
	// Attachment is the service for content attachment operations.
	Attachment *ContentAttachmentService
	// ChildrenDescendant is the service for children and descendants operations.
	ChildrenDescendant *ChildrenDescandantsService
	// Comment is the service for comment operations.
	Comment *CommentService
	// Permission is the service for permission operations.
	Permission *PermissionService
	// Label is the service for content label operations.
	Label *ContentLabelService
	// Property is the service for content property operations.
	Property *PropertyService
	// Restriction is the service for content restriction operations.
	Restriction *RestrictionService
	// Version is the service for content version operations.
	Version *VersionService
}

// Gets returns all content in a Confluence instance.
//
// GET /wiki/rest/api/content
//
// https://docs.go-atlassian.io/confluence-cloud/content#get-content
func (c *ContentService) Gets(ctx context.Context, options *model.GetContentOptionsScheme, startAt, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error) {
	return c.internalClient.Gets(ctx, options, startAt, maxResults)
}

// Create creates a new piece of content or publishes an existing draft
//
// To publish a draft, add the id and status properties to the body of the request.
//
// Set the id to the ID of the draft and set the status to 'current'.
//
// When the request is sent, a new piece of content will be created and the metadata from the draft will be transferred into it.
//
// POST /wiki/rest/api/content
//
// https://docs.go-atlassian.io/confluence-cloud/content#create-content
func (c *ContentService) Create(ctx context.Context, payload *model.ContentScheme) (*model.ContentScheme, *model.ResponseScheme, error) {
	return c.internalClient.Create(ctx, payload)
}

// Search returns the list of content that matches a Confluence Query Language (CQL) query
//
// GET /wiki/rest/api/content/search
//
// https://docs.go-atlassian.io/confluence-cloud/content#search-contents-by-cql
func (c *ContentService) Search(ctx context.Context, cql, cqlContext string, expand []string, cursor string, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error) {
	return c.internalClient.Search(ctx, cql, cqlContext, expand, cursor, maxResults)
}

// Get returns a single piece of content, like a page or a blog post.
//
// By default, the following objects are expanded: space, history, version.
//
// GET /wiki/rest/api/content/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/content#get-content
func (c *ContentService) Get(ctx context.Context, contentID string, expand []string, version int) (*model.ContentScheme, *model.ResponseScheme, error) {
	return c.internalClient.Get(ctx, contentID, expand, version)
}

// Update updates a piece of content.
//
// Use this method to update the title or body of a piece of content, change the status, change the parent page, and more.
//
// PUT /wiki/rest/api/content/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/content#update-content
func (c *ContentService) Update(ctx context.Context, contentID string, payload *model.ContentScheme) (*model.ContentScheme, *model.ResponseScheme, error) {
	return c.internalClient.Update(ctx, contentID, payload)
}

// Delete moves a piece of content to the space's trash or purges it from the trash, depending on the content's type and status:
//
// If the content's type is page or blogpost and its status is current, it will be trashed.
//
// If the content's type is page or blogpost and its status is trashed, the content will be purged from the trash and deleted permanently.
//
// === Note, you must also set the status query parameter to trashed in your request. ===
//
// If the content's type is comment or attachment, it will be deleted permanently without being trashed.
//
// DELETE /wiki/rest/api/content/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/content#delete-content
func (c *ContentService) Delete(ctx context.Context, contentID, status string) (*model.ResponseScheme, error) {
	return c.internalClient.Delete(ctx, contentID, status)
}

// History returns the most recent update for a piece of content.
//
// GET /wiki/rest/api/content/{id}/history
//
// https://docs.go-atlassian.io/confluence-cloud/content#get-content-history
func (c *ContentService) History(ctx context.Context, contentID string, expand []string) (*model.ContentHistoryScheme, *model.ResponseScheme, error) {
	return c.internalClient.History(ctx, contentID, expand)
}

// Archive archives a list of pages.
//
// The pages to be archived are specified as a list of content IDs.
//
// This API accepts the archival request and returns a task ID. The archival process happens asynchronously.
//
// Use the /longtask/ REST API to get the copy task status.
//
// POST /wiki/rest/api/content/archive
//
// https://docs.go-atlassian.io/confluence-cloud/content#archive-pages
func (c *ContentService) Archive(ctx context.Context, payload *model.ContentArchivePayloadScheme) (*model.ContentArchiveResultScheme, *model.ResponseScheme, error) {
	return c.internalClient.Archive(ctx, payload)
}

type internalContentImpl struct {
	c service.Connector
}

func (i *internalContentImpl) Gets(ctx context.Context, options *model.GetContentOptionsScheme, startAt, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error) {

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if options != nil {

		if options.ContextType != "" {
			query.Add("type", options.ContextType)
		}

		if options.SpaceKey != "" {
			query.Add("spaceKey", options.SpaceKey)
		}

		if options.Title != "" {
			query.Add("title", options.Title)
		}

		if options.Trigger != "" {
			query.Add("trigger", options.Trigger)
		}

		if options.OrderBy != "" {
			query.Add("orderby", options.OrderBy)
		}

		if !options.PostingDay.IsZero() {
			query.Add("postingDay", options.PostingDay.Format("2006-01-02"))
		}

		if len(options.Status) != 0 {
			query.Add("status", strings.Join(options.Status, ","))
		}

		if len(options.Expand) != 0 {
			query.Add("expand", strings.Join(options.Expand, ","))
		}

	}

	endpoint := fmt.Sprintf("wiki/rest/api/content?%v", query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ContentPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalContentImpl) Create(ctx context.Context, payload *model.ContentScheme) (*model.ContentScheme, *model.ResponseScheme, error) {

	endpoint := "wiki/rest/api/content"

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	content := new(model.ContentScheme)
	response, err := i.c.Call(request, content)
	if err != nil {
		return nil, response, err
	}

	return content, response, nil
}

func (i *internalContentImpl) Search(ctx context.Context, cql, cqlContext string, expand []string, cursor string, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error) {

	if cql == "" {
		return nil, nil, model.ErrNoCQL
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(maxResults))
	query.Add("cql", cql)

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	if cqlContext != "" {
		query.Add("cqlcontext", cqlContext)
	}

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/search?%v", query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ContentPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalContentImpl) Get(ctx context.Context, contentID string, expand []string, version int) (*model.ContentScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	query := url.Values{}
	query.Add("version", strconv.Itoa(version))

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v?%v", contentID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	content := new(model.ContentScheme)
	response, err := i.c.Call(request, content)
	if err != nil {
		return nil, response, err
	}

	return content, response, nil
}

func (i *internalContentImpl) Update(ctx context.Context, contentID string, payload *model.ContentScheme) (*model.ContentScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v", contentID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	content := new(model.ContentScheme)
	response, err := i.c.Call(request, content)
	if err != nil {
		return nil, response, err
	}

	return content, response, nil
}

func (i *internalContentImpl) Delete(ctx context.Context, contentID, status string) (*model.ResponseScheme, error) {

	if contentID == "" {
		return nil, model.ErrNoContentID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v", contentID))

	if status != "" {
		query := url.Values{}
		query.Add("status", status)

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint.String(), "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalContentImpl) History(ctx context.Context, contentID string, expand []string) (*model.ContentHistoryScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/history", contentID))

	if len(expand) != 0 {
		query := url.Values{}
		query.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	history := new(model.ContentHistoryScheme)
	response, err := i.c.Call(request, history)
	if err != nil {
		return nil, response, err
	}

	return history, response, nil
}

func (i *internalContentImpl) Archive(ctx context.Context, payload *model.ContentArchivePayloadScheme) (*model.ContentArchiveResultScheme, *model.ResponseScheme, error) {

	endpoint := "wiki/rest/api/content/archive"

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	result := new(model.ContentArchiveResultScheme)
	response, err := i.c.Call(request, result)
	if err != nil {
		return nil, response, err
	}

	return result, response, nil
}
