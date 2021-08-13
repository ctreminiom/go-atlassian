package confluence

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ContentLabelService struct {
	client *Client
}

// Gets returns the labels on a piece of content.
// Atlassian Docs: https://developer.atlassian.com/cloud/confluence/rest/api-group-content-labels/#api-wiki-rest-api-content-id-label-get
func (c *ContentLabelService) Gets(ctx context.Context, contentID, prefix string, startAt, maxResults int) (result *ContentLabelPageScheme,
	response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
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
// Atlassian Docs: https://developer.atlassian.com/cloud/confluence/rest/api-group-content-labels/#api-wiki-rest-api-content-id-label-post
func (c *ContentLabelService) Add(ctx context.Context, contentID string, payload []*ContentLabelPayloadScheme, want400Response bool) (
	result *ContentLabelPageScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
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
// Atlassian Docs: https://developer.atlassian.com/cloud/confluence/rest/api-group-content-labels/#api-wiki-rest-api-content-id-label-label-delete
func (c *ContentLabelService) Remove(ctx context.Context, contentID, labelName string) (response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, notContentIDError
	}

	if len(labelName) == 0 {
		return nil, notLabelNameError
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

type ContentLabelPayloadScheme struct {
	Prefix string `json:"prefix"`
	Name   string `json:"name"`
}

type ContentLabelPageScheme struct {
	Results []*ContentLabelScheme `json:"results,omitempty"`
	Start   int                   `json:"start,omitempty"`
	Limit   int                   `json:"limit,omitempty"`
	Size    int                   `json:"size,omitempty"`
}

type ContentLabelScheme struct {
	Prefix string `json:"prefix,omitempty"`
	Name   string `json:"name,omitempty"`
	ID     string `json:"id,omitempty"`
	Label  string `json:"label,omitempty"`
}

var (
	notLabelNameError = fmt.Errorf("error!, please provide a label name")
)
