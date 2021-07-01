package confluence

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ContentChildrenDescendantService struct {
	client *Client
}

// Children returns a map of the direct children of a piece of content.
// A piece of content has different types of child content, depending on its type.
// These are the default parent-child content type relationships:
// page: child content is page, comment, attachment
// blogpost: child content is comment, attachment
// attachment: child content is comment
// comment: child content is attachment
func (c *ContentChildrenDescendantService) Children(ctx context.Context, contentID string, expand []string,
	parentVersion int) (result *ContentChildrenScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
	}

	query := url.Values{}
	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	if parentVersion != 0 {
		query.Add("parentVersion", strconv.Itoa(parentVersion))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/wiki/rest/api/content/%v/child", contentID))

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

type ContentChildrenScheme struct {
	Attachment *ContentPageScheme `json:"attachment,omitempty"`
	Comments   *ContentPageScheme `json:"comment,omitempty"`
	Page       *ContentPageScheme `json:"page,omitempty"`
	Links      *LinkScheme        `json:"_links,omitempty"`
}

// ChildrenByType returns all children of a given type, for a piece of content.
// A piece of content has different types of child content
func (c *ContentChildrenDescendantService) ChildrenByType(ctx context.Context, contentID, contentType string,
	parentVersion int, expand []string, startAt, maxResults int) (result *ContentPageScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
	}

	if len(contentType) == 0 {
		return nil, nil, notContentTypeError
	}

	var hasValidValue bool
	for _, value := range validContentTypes {

		if value == contentType {
			hasValidValue = true
			break
		}
	}

	if !hasValidValue {
		return nil, nil, invalidContentTypeError
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

	var endpoint = fmt.Sprintf("/wiki/rest/api/content/%v/child/%v?%v", contentID, contentType, query.Encode())

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

// Descendants returns a map of the descendants of a piece of content.
// This is similar to Get content children, except that this method returns child pages at all levels,
// rather than just the direct child pages.
func (c *ContentChildrenDescendantService) Descendants(ctx context.Context, contentID string, expand []string,
) (result *ContentChildrenScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
	}

	query := url.Values{}
	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/wiki/rest/api/content/%v/descendant", contentID))

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

// DescendantsByType returns all descendants of a given type, for a piece of content.
// This is similar to Get content children by type,
// except that this method returns child pages at all levels, rather than just the direct child pages.
func (c *ContentChildrenDescendantService) DescendantsByType(ctx context.Context, contentID, contentType,
	depth string, expand []string, startAt, maxResults int) (result *ContentPageScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
	}

	if len(contentType) == 0 {
		return nil, nil, notContentTypeError
	}

	var hasValidValue bool
	for _, value := range validContentTypes {

		if value == contentType {
			hasValidValue = true
			break
		}
	}

	if !hasValidValue {
		return nil, nil, invalidContentTypeError
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

	var endpoint = fmt.Sprintf("/wiki/rest/api/content/%v/descendant/%v?%v", contentID, contentType, query.Encode())

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

type CopyOptionsScheme struct {
	CopyAttachments    bool                       `json:"copyAttachments,omitempty"`
	CopyPermissions    bool                       `json:"copyPermissions,omitempty"`
	CopyProperties     bool                       `json:"copyProperties,omitempty"`
	CopyLabels         bool                       `json:"copyLabels,omitempty"`
	CopyCustomContents bool                       `json:"copyCustomContents,omitempty"`
	DestinationPageID  string                     `json:"destinationPageId,omitempty"`
	TitleOptions       *CopyTitleOptionScheme     `json:"titleOptions,omitempty"`
	Destination        *CopyPageDestinationScheme `json:"destination,omitempty"`
	PageTitle          string                     `json:"pageTitle,omitempty"`
	Body               *CopyPageBodyScheme        `json:"body,omitempty"`
}

type CopyTitleOptionScheme struct {
	Prefix  string `json:"prefix,omitempty"`
	Replace string `json:"replace,omitempty"`
	Search  string `json:"search,omitempty"`
}

type CopyPageDestinationScheme struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type CopyPageBodyScheme struct {
	Storage *BodyNodeScheme `json:"storage"`
	Editor2 *BodyNodeScheme `json:"editor2"`
}

type TaskScheme struct {
	ID    string          `json:"id,omitempty"`
	Links *TaskLinkScheme `json:"links,omitempty"`
}

type TaskLinkScheme struct {
	Status string `json:"status"`
}

// CopyHierarchy copy page hierarchy allows the copying of an entire hierarchy of pages and their associated properties,
// permissions and attachments. The id path parameter refers to the content id of the page to copy,
// and the new parent of this copied page is defined using the destinationPageId in the request body.
// The titleOptions object defines the rules of renaming page titles during the copy;
// for example, search and replace can be used in conjunction to rewrite the copied page titles.
// RESPONSE =  Use the /longtask/ REST API to get the copy task status.
func (c *ContentChildrenDescendantService) CopyHierarchy(ctx context.Context, contentID string,
	options *CopyOptionsScheme) (result *TaskScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
	}

	payloadAsReader, err := transformStructToReader(options)
	if err != nil {
		return nil, nil, err
	}

	var endpoint = fmt.Sprintf("/wiki/rest/api/content/%v/pagehierarchy/copy", contentID)

	request, err := c.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// CopyPage copies a single page and its associated properties, permissions, attachments, and custom contents.
// The id path parameter refers to the content ID of the page to copy.
// The target of the page to be copied is defined using the destination in the request body and can be one of the following types.
// 1. space: page will be copied to the specified space as a root page on the space
// 2. parent_page: page will be copied as a child of the specified parent page
// 3. existing_page: page will be copied and replace the specified page
// By default, the following objects are expanded: space, history, version.
func (c *ContentChildrenDescendantService) CopyPage(ctx context.Context, contentID string, expand []string,
	options *CopyOptionsScheme) (result *ContentScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
	}

	query := url.Values{}
	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/wiki/rest/api/content/%v/copy", contentID))

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	payloadAsReader, err := transformStructToReader(options)
	if err != nil {
		return nil, nil, err
	}

	request, err := c.client.newRequest(ctx, http.MethodPost, endpoint.String(), payloadAsReader)
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

var (
	notContentTypeError     = fmt.Errorf("error, the content type is required, please provide a valid value")
	invalidContentTypeError = fmt.Errorf("error, the content type provided is not valid, please provide one of "+
		"the following options: %v", strings.Join(validContentTypes, ","))

	validContentTypes = []string{"page", "comment", "attachment"}
)
