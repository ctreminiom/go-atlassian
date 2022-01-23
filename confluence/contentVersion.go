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

type ContentVersionService struct{ client *Client }

// Gets returns the versions for a piece of content in descending order.
func (c *ContentVersionService) Gets(ctx context.Context, contentID string, expand []string, start, limit int) (
	result *models.ContentVersionPageScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, models.ErrNoContentIDError
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(start))
	query.Add("limit", strconv.Itoa(limit))

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/version?%v", contentID, query.Encode())

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

// Get returns a version for a piece of content.
func (c *ContentVersionService) Get(ctx context.Context, contentID string, versionNumber int, expand []string) (
	result *models.ContentVersionScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, models.ErrNoContentIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/version/%v", contentID, versionNumber))

	if len(expand) != 0 {
		query := url.Values{}
		query.Add("expand", strings.Join(expand, ","))
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

// Restore restores a historical version to be the latest version.
// That is, a new version is created with the content of the historical version.
func (c *ContentVersionService) Restore(ctx context.Context, contentID string, payload *models.ContentRestorePayloadScheme,
	expand []string) (result *models.ContentVersionScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, models.ErrNoContentIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/version", contentID))

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

// Delete deletes a historical version.
// This does not delete the changes made to the content in that version, rather the changes for the deleted version
// are rolled up into the next version. Note, you cannot delete the current version.
func (c *ContentVersionService) Delete(ctx context.Context, contentID string, versionNumber int) (response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, models.ErrNoContentIDError
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/version/%v", contentID, versionNumber)

	request, err := c.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Call(request, nil)
	if err != nil {
		return response, err
	}

	return
}
