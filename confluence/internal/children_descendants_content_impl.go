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

// NewChildrenDescandantsService creates a new instance of ChildrenDescandantsService.
func NewChildrenDescandantsService(client service.Connector) *ChildrenDescandantsService {
	return &ChildrenDescandantsService{
		internalClient: &internalChildrenDescandantsImpl{c: client},
	}
}

// ChildrenDescandantsService provides methods to interact with children and descendants operations in Confluence.
type ChildrenDescandantsService struct {
	// internalClient is the connector interface for children and descendants operations.
	internalClient confluence.ChildrenDescendantConnector
}

// Children returns a map of the direct children of a piece of content.
//
// A piece of content has different types of child content, depending on its type.
//
// These are the default parent-child content type relationships:
//
// page: child content is page, comment, attachment
//
// blogpost: child content is comment, attachment
//
// attachment: child content is comment
//
// comment: child content is attachment
//
// GET /wiki/rest/api/content/{id}/child
//
// https://docs.go-atlassian.io/confluence-cloud/content/children-descendants#get-content-children
func (c *ChildrenDescandantsService) Children(ctx context.Context, contentID string, expand []string, parentVersion int) (*model.ContentChildrenScheme, *model.ResponseScheme, error) {
	return c.internalClient.Children(ctx, contentID, expand, parentVersion)
}

// Move moves a page from its current location in the hierarchy to another.
//
// Position describes where in the hierarchy the page should be moved to in
// relationship to targetID.
//
// before: page will be a sibling of target but show up just before target in
// the list of children
//
// after: page will be a sibling of target but show up just after target in the
// list of children
//
// append: page will be a child of the target and be appended to targets list of
// children
//
// PUT /wiki/rest/api/content/{pageId}/move/{position}/{targetId}
//
// https://docs.go-atlassian.io/confluence-cloud/content/children-descendants#move
func (c *ChildrenDescandantsService) Move(ctx context.Context, pageID string, position string, targetID string) (*model.ContentMoveScheme, *model.ResponseScheme, error) {
	return c.internalClient.Move(ctx, pageID, position, targetID)
}

// ChildrenByType returns all children of a given type, for a piece of content.
//
// # A piece of content has different types of child content
//
// GET /wiki/rest/api/content/{id}/child/{type}
//
// https://docs.go-atlassian.io/confluence-cloud/content/children-descendants#get-content-children-by-type
func (c *ChildrenDescandantsService) ChildrenByType(ctx context.Context, contentID, contentType string, parentVersion int, expand []string, startAt, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error) {
	return c.internalClient.ChildrenByType(ctx, contentID, contentType, parentVersion, expand, startAt, maxResults)
}

// Descendants returns a map of the descendants of a piece of content.
//
// This is similar to Get content children, except that this method returns child pages at all levels,
//
// rather than just the direct child pages.
//
// GET /wiki/rest/api/content/{id}/descendant
//
// https://docs.go-atlassian.io/confluence-cloud/content/children-descendants#get-content-descendants
func (c *ChildrenDescandantsService) Descendants(ctx context.Context, contentID string, expand []string) (*model.ContentChildrenScheme, *model.ResponseScheme, error) {
	return c.internalClient.Descendants(ctx, contentID, expand)
}

// DescendantsByType returns all descendants of a given type, for a piece of content.
//
// This is similar to Get content children by type,
//
// except that this method returns child pages at all levels, rather than just the direct child pages.
//
// GET /wiki/rest/api/content/{id}/descendant/{type}
//
// https://docs.go-atlassian.io/confluence-cloud/content/children-descendants#get-content-descendants-by-type
func (c *ChildrenDescandantsService) DescendantsByType(ctx context.Context, contentID, contentType, depth string, expand []string, startAt, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error) {
	return c.internalClient.DescendantsByType(ctx, contentID, contentType, depth, expand, startAt, maxResults)
}

// CopyHierarchy copy page hierarchy allows the copying of an entire hierarchy of pages and their associated properties,
//
// permissions and attachments. The id path parameter refers to the content id of the page to copy,
//
// and the new parent of this copied page is defined using the destinationPageId in the request body.
//
// The titleOptions object defines the rules of renaming page titles during the copy;
//
// for example, search and replace can be used in conjunction to rewrite the copied page titles.
//
// RESPONSE =  Use the /longtask/ REST API to get the copy task status.
//
// POST /wiki/rest/api/content/{id}/pagehierarchy/copy
//
// https://docs.go-atlassian.io/confluence-cloud/content/children-descendants#copy-page-hierarchy
func (c *ChildrenDescandantsService) CopyHierarchy(ctx context.Context, contentID string, options *model.CopyOptionsScheme) (*model.TaskScheme, *model.ResponseScheme, error) {
	return c.internalClient.CopyHierarchy(ctx, contentID, options)
}

// CopyPage copies a single page and its associated properties, permissions, attachments, and custom contents.
//
// The id path parameter refers to the content ID of the page to copy.
//
// The target of the page to be copied is defined using the destination in the request body and can be one of the following types.
//
// 1. space: page will be copied to the specified space as a root page on the space
//
// 2. parent_page: page will be copied as a child of the specified parent page
//
// 3. existing_page: page will be copied and replace the specified page
//
// By default, the following objects are expanded: space, history, version.
//
// POST /wiki/rest/api/content/{id}/copy
//
// https://docs.go-atlassian.io/confluence-cloud/content/children-descendants#copy-single-page
func (c *ChildrenDescandantsService) CopyPage(ctx context.Context, contentID string, expand []string, options *model.CopyOptionsScheme) (*model.ContentScheme, *model.ResponseScheme, error) {
	return c.internalClient.CopyPage(ctx, contentID, expand, options)
}

type internalChildrenDescandantsImpl struct {
	c service.Connector
}

func (i *internalChildrenDescandantsImpl) Children(ctx context.Context, contentID string, expand []string, parentVersion int) (*model.ContentChildrenScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/child", contentID))

	query := url.Values{}

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	if parentVersion != 0 {
		query.Add("parentVersion", strconv.Itoa(parentVersion))
	}

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	children := new(model.ContentChildrenScheme)
	response, err := i.c.Call(request, children)
	if err != nil {
		return nil, response, err
	}

	return children, response, nil
}

func (i *internalChildrenDescandantsImpl) Move(ctx context.Context, pageID string, position string, targetID string) (*model.ContentMoveScheme, *model.ResponseScheme, error) {

	if pageID == "" {
		return nil, nil, model.ErrNoPageID
	}

	if position == "" {
		return nil, nil, model.ErrNoPosition
	}

	if targetID == "" {
		return nil, nil, model.ErrNoTargetID
	}

	_, validPosition := model.ValidPositions[position]
	if !validPosition {
		return nil, nil, model.ErrInvalidPosition
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/move/%s/%v", pageID, position, targetID))

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	movement := new(model.ContentMoveScheme)
	response, err := i.c.Call(request, movement)
	if err != nil {
		return nil, response, err
	}

	return movement, response, nil
}

func (i *internalChildrenDescandantsImpl) ChildrenByType(ctx context.Context, contentID, contentType string, parentVersion int, expand []string, startAt, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	if contentType == "" {
		return nil, nil, model.ErrNoContentType
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	if parentVersion != 0 {
		query.Add("parentVersion", strconv.Itoa(parentVersion))
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/child/%v?%v", contentID, contentType, query.Encode())

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

func (i *internalChildrenDescandantsImpl) Descendants(ctx context.Context, contentID string, expand []string) (*model.ContentChildrenScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/descendant", contentID))

	query := url.Values{}

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	children := new(model.ContentChildrenScheme)
	response, err := i.c.Call(request, children)
	if err != nil {
		return nil, response, err
	}

	return children, response, nil
}

func (i *internalChildrenDescandantsImpl) DescendantsByType(ctx context.Context, contentID, contentType, depth string, expand []string, startAt, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	if contentType == "" {
		return nil, nil, model.ErrNoContentType
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	if len(depth) != 0 {
		query.Add("depth", depth)
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/descendant/%v?%v", contentID, contentType, query.Encode())

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

func (i *internalChildrenDescandantsImpl) CopyHierarchy(ctx context.Context, contentID string, options *model.CopyOptionsScheme) (*model.TaskScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/pagehierarchy/copy", contentID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", options)
	if err != nil {
		return nil, nil, err
	}

	task := new(model.TaskScheme)
	response, err := i.c.Call(request, task)
	if err != nil {
		return nil, response, err
	}

	return task, response, nil
}

func (i *internalChildrenDescandantsImpl) CopyPage(ctx context.Context, contentID string, expand []string, options *model.CopyOptionsScheme) (*model.ContentScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/copy", contentID))

	if len(expand) != 0 {
		query := url.Values{}
		query.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", options)
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
