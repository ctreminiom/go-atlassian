package confluence

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ContentRestrictionService struct {
	client    *Client
	Operation *ContentRestrictionOperationService
}

// Gets returns the restrictions on a piece of content.
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/restrictions#get-restrictions
func (c *ContentRestrictionService) Gets(ctx context.Context, contentID string, expand []string, startAt, maxResults int) (
	result *models.ContentRestrictionPageScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, models.ErrNoContentIDError
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/restriction?%v", contentID, query.Encode())

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

// Add adds restrictions to a piece of content. Note, this does not change any existing restrictions on the content.
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/restrictions#add-restrictions
func (c *ContentRestrictionService) Add(ctx context.Context, contentID string, payload *models.ContentRestrictionUpdatePayloadScheme,
	expand []string) (result *models.ContentRestrictionPageScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, models.ErrNoContentIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/restriction", contentID))

	query := url.Values{}
	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	payloadAsReader, err := transformStructToReader(payload)
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

// Delete removes all restrictions (read and update) on a piece of content.
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/restrictions#delete-restrictions
func (c *ContentRestrictionService) Delete(ctx context.Context, contentID string, expand []string) (
	result *models.ContentRestrictionPageScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, models.ErrNoContentIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/restriction", contentID))

	query := url.Values{}
	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := c.client.newRequest(ctx, http.MethodDelete, endpoint.String(), nil)
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

// Update updates restrictions for a piece of content. This removes the existing restrictions and replaces them with the restrictions in the request.
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/restrictions#update-restrictions
func (c *ContentRestrictionService) Update(ctx context.Context, contentID string, payload *models.ContentRestrictionUpdatePayloadScheme,
	expand []string) (result *models.ContentRestrictionPageScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, models.ErrNoContentIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/restriction", contentID))

	query := url.Values{}
	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := c.client.newRequest(ctx, http.MethodPut, endpoint.String(), payloadAsReader)
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
