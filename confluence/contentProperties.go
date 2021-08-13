package confluence

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ContentPropertyService struct{ client *Client }

// Gets returns the properties for a piece of content.
// Atlassian Docs: https://developer.atlassian.com/cloud/confluence/rest/api-group-content-properties/#api-wiki-rest-api-content-id-property-get
func (c *ContentPropertyService) Gets(ctx context.Context, contentID string, expand []string, startAt, maxResults int) (
	result *ContentPropertyPageScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	var endpoint = fmt.Sprintf("/wiki/rest/api/content/%v/property?%v", contentID, query.Encode())

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

// Create creates a property for an existing piece of content.
// Atlassian Docs: https://developer.atlassian.com/cloud/confluence/rest/api-group-content-properties/#api-wiki-rest-api-content-id-property-post
func (c *ContentPropertyService) Create(ctx context.Context, contentID string, payload *ContentPropertyPayloadScheme) (
	result *ContentPropertyScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint = fmt.Sprintf("/wiki/rest/api/content/%v/property", contentID)

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

// Get returns a content property for a piece of content.
// Atlassian Docs: https://developer.atlassian.com/cloud/confluence/rest/api-group-content-properties/#api-wiki-rest-api-content-id-property-key-get
func (c *ContentPropertyService) Get(ctx context.Context, contentID, key string) (result *ContentPropertyScheme,
	response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
	}

	if len(key) == 0 {
		return nil, nil, notPropertyKeyError
	}

	var endpoint = fmt.Sprintf("/wiki/rest/api/content/%v/property/%v", contentID, key)

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

// Delete deletes a content property.
// Atlassian Docs: https://developer.atlassian.com/cloud/confluence/rest/api-group-content-properties/#api-wiki-rest-api-content-id-property-key-delete
func (c *ContentPropertyService) Delete(ctx context.Context, contentID, key string) (response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, notContentIDError
	}

	if len(key) == 0 {
		return nil, notPropertyKeyError
	}

	var endpoint = fmt.Sprintf("/wiki/rest/api/content/%v/property/%v", contentID, key)

	request, err := c.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	response, err = c.client.Call(request, nil)
	if err != nil {
		return response, err
	}

	return
}

type ContentPropertyPayloadScheme struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ContentPropertyPageScheme struct {
	Results []*ContentPropertyScheme `json:"results,omitempty"`
	Start   int                      `json:"start,omitempty"`
	Limit   int                      `json:"limit,omitempty"`
	Size    int                      `json:"size,omitempty"`
}

type ContentPropertyScheme struct {
	ID         string                        `json:"id,omitempty"`
	Key        string                        `json:"key,omitempty"`
	Value      interface{}                   `json:"value,omitempty"`
	Version    *ContentPropertyVersionScheme `json:"version,omitempty"`
	Expandable struct {
		Content              string `json:"content,omitempty"`
		AdditionalProperties string `json:"additionalProperties,omitempty"`
	} `json:"_expandable,omitempty"`
}

type ContentPropertyVersionScheme struct {
	When                string `json:"when,omitempty"`
	Message             string `json:"message,omitempty"`
	Number              int    `json:"number,omitempty"`
	MinorEdit           bool   `json:"minorEdit,omitempty"`
	ContentTypeModified bool   `json:"contentTypeModified,omitempty"`
}

var (
	notPropertyKeyError = fmt.Errorf("error!, please provide a valid property key")
)
