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

type ContentLabelService struct {
	client *Client
}

// Gets returns the labels on a piece of content.
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/labels#get-labels-for-content
func (c *ContentLabelService) Gets(ctx context.Context, contentID, prefix string, startAt, maxResults int) (result *model.ContentLabelPageScheme,
	response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, model.ErrNoCQLError
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if len(prefix) != 0 {
		query.Add("prefix", prefix)
	}

	var endpoint = fmt.Sprintf("wiki/rest/api/content/%v/label?%v", contentID, query.Encode())

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

// Add adds labels to a piece of content. Does not modify the existing labels.
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/labels#add-labels-to-content
func (c *ContentLabelService) Add(ctx context.Context, contentID string, payload []*model.ContentLabelPayloadScheme, want400Response bool) (
	result *model.ContentLabelPageScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, model.ErrNoCQLError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/wiki/rest/api/content/%v/label", contentID))

	query := url.Values{}
	if want400Response {
		query.Add("use-400-error-response", "true")
	}

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
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

// Remove removes a label from a piece of content
// Docs: https://docs.go-atlassian.io/confluence-cloud/content/labels#remove-label-from-content
func (c *ContentLabelService) Remove(ctx context.Context, contentID, labelName string) (response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, model.ErrNoCQLError
	}

	if len(labelName) == 0 {
		return nil, model.ErrNoContentLabelError
	}

	var endpoint = fmt.Sprintf("/wiki/rest/api/content/%v/label/%v", contentID, labelName)

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
