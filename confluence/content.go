package confluence

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ContentService struct {
	client             *Client
	Attachment         *ContentAttachmentService
	ChildrenDescendant *ContentChildrenDescendantService
	Comment            *ContentCommentService
	Permission         *ContentPermissionService
	Label              *ContentLabelService
	Property           *ContentPropertyService
	Restriction        *ContentRestrictionService
	Version            *ContentVersionService
}

// Gets returns all content in a Confluence instance.
func (c *ContentService) Gets(ctx context.Context, options *model.GetContentOptionsScheme, startAt, maxResults int) (
	result *model.ContentPageScheme, response *ResponseScheme, err error) {

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

	var endpoint = fmt.Sprintf("/wiki/rest/api/content?%v", query.Encode())

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Create creates a new piece of content or publishes an existing draft
// To publish a draft, add the id and status properties to the body of the request.
// Set the id to the ID of the draft and set the status to 'current'.
// When the request is sent, a new piece of content will be created and the metadata from the draft will be transferred into it.
// Docs: https://docs.go-atlassian.io/confluence-cloud/content#create-content
func (c *ContentService) Create(ctx context.Context, payload *model.ContentScheme) (result *model.ContentScheme,
	response *ResponseScheme, err error) {

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint = "/wiki/rest/api/content"

	request, err := c.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Search returns the list of content that matches a Confluence Query Language (CQL) query
// Docs: https://docs.go-atlassian.io/confluence-cloud/content#search-contents-by-cql
func (c *ContentService) Search(ctx context.Context, cql, cqlContext string, expand []string, cursor string, maxResults int) (
	result *model.ContentPageScheme, response *ResponseScheme, err error) {

	if cql == "" {
		return nil, nil, model.ErrNoCQLError
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

	var endpoint = fmt.Sprintf("/wiki/rest/api/content/search?%v", query.Encode())

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Get returns a single piece of content, like a page or a blog post.
// By default, the following objects are expanded: space, history, version.
// Docs: https://docs.go-atlassian.io/confluence-cloud/content#get-content
func (c *ContentService) Get(ctx context.Context, contentID string, expand []string, version int) (result *model.ContentScheme,
	response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, model.ErrNoContentIDError
	}

	query := url.Values{}

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	if version != 0 {
		query.Add("version", strconv.Itoa(version))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/wiki/rest/api/content/%v", contentID))

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Update updates a piece of content.
// Use this method to update the title or body of a piece of content, change the status, change the parent page, and more.
// Docs: https://docs.go-atlassian.io/confluence-cloud/content#update-content
func (c *ContentService) Update(ctx context.Context, contentID string, payload *model.ContentScheme) (result *model.ContentScheme,
	response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, model.ErrNoContentIDError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint = fmt.Sprintf("/wiki/rest/api/content/%v", contentID)

	request, err := c.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Delete moves a piece of content to the space's trash or purges it from the trash,
// depending on the content's type and status:
// If the content's type is page or blogpost and its status is current, it will be trashed.
// If the content's type is page or blogpost and its status is trashed, the content will be purged from the trash and deleted permanently.
// === Note, you must also set the status query parameter to trashed in your request. ===
// If the content's type is comment or attachment, it will be deleted permanently without being trashed.
// Docs: https://docs.go-atlassian.io/confluence-cloud/content#delete-content
func (c *ContentService) Delete(ctx context.Context, contentID, status string) (response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, model.ErrNoContentIDError
	}

	query := url.Values{}
	if status != "" {
		query.Add("status", status)
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/wiki/rest/api/content/%v", contentID))

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := c.client.newRequest(ctx, http.MethodDelete, endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	response, err = c.client.Call(request, nil)
	if err != nil {
		return response, err
	}

	return
}

// History returns the most recent update for a piece of content.
// Docs: https://docs.go-atlassian.io/confluence-cloud/content#get-content-history
func (c *ContentService) History(ctx context.Context, contentID string, expand []string) (result *model.ContentHistoryScheme,
	response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, model.ErrNoContentIDError
	}

	query := url.Values{}
	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/wiki/rest/api/content/%v/history", contentID))

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}
