package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type CommentPropertiesService struct{ client *Client }

type PropertiesKeysScheme struct {
	Keys []PropertiesKeyScheme `json:"keys,omitempty"`
}

type PropertiesKeyScheme struct {
	Self  string `json:"self,omitempty"`
	Key   string `json:"key,omitempty"`
	Value struct {
		SystemConversationID string `json:"system.conversation.id,omitempty"`
		SystemSupportTime    string `json:"system.support.time,omitempty"`
	} `json:"value,omitempty"`
}

// Returns the keys of all the properties of a comment.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-comment-properties/#api-rest-api-3-comment-commentid-properties-get
func (c *CommentPropertiesService) Keys(ctx context.Context, commentID string) (result *PropertiesKeysScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/comment/%v/properties", commentID)
	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Do(request)
	if err != nil {
		return
	}

	result = new(PropertiesKeysScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns the value of a comment property.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-comment-properties/#api-rest-api-3-comment-commentid-properties-propertykey-get
func (c *CommentPropertiesService) Get(ctx context.Context, commentID, propertyKey string) (result *PropertiesKeyScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/comment/%v/properties/%v", commentID, propertyKey)
	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Do(request)
	if err != nil {
		return
	}

	result = new(PropertiesKeyScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Creates or updates the value of a property for a comment.
// Use this resource to store custom data against a comment.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-comment-properties/#api-rest-api-3-comment-commentid-properties-propertykey-put
func (c *CommentPropertiesService) Set(ctx context.Context, commentID, propertyKey string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/comment/%v/properties/%v", commentID, propertyKey)
	request, err := c.client.newRequest(ctx, http.MethodPut, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Deletes a comment property.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-comment-properties/#api-rest-api-3-comment-commentid-properties-propertykey-delete
func (c *CommentPropertiesService) Delete(ctx context.Context, commentID, propertyKey string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/comment/%v/properties/%v", commentID, propertyKey)
	request, err := c.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Do(request)
	if err != nil {
		return
	}

	return
}
