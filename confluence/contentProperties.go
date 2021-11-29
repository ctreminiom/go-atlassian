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

type ContentPropertyService struct{ client *Client }

// Gets returns the properties for a piece of content.
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/properties#get-content-properties
func (c *ContentPropertyService) Gets(ctx context.Context, contentID string, expand []string, startAt, maxResults int) (
	result *model.ContentPropertyPageScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, model.ErrNoContentIDError
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
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/properties#create-content-property
func (c *ContentPropertyService) Create(ctx context.Context, contentID string, payload *model.ContentPropertyPayloadScheme) (
	result *model.ContentPropertyScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, model.ErrNoContentIDError
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
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/properties#get-content-property
func (c *ContentPropertyService) Get(ctx context.Context, contentID, key string) (result *model.ContentPropertyScheme,
	response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, model.ErrNoContentIDError
	}

	if len(key) == 0 {
		return nil, nil, model.ErrNoContentPropertyError
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
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/properties#delete-content-property
func (c *ContentPropertyService) Delete(ctx context.Context, contentID, key string) (response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, model.ErrNoContentIDError
	}

	if len(key) == 0 {
		return nil, model.ErrNoContentPropertyError
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
