package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type IssuePropertyService struct{ client *Client }

type IssuePropertyScheme struct {
	Key   string `json:"key"`
	Value struct {
		SystemConversationID string `json:"system.conversation.id"`
		SystemSupportTime    string `json:"system.support.time"`
	} `json:"value"`
}

// Returns the key and value of an issue's property.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-properties/#api-rest-api-3-issue-issueidorkey-properties-propertykey-get
func (i *IssuePropertyService) Get(ctx context.Context, issueKeyOrID, propertyKey string) (result *IssuePropertyScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/properties/%v", issueKeyOrID, propertyKey)
	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssuePropertyScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type IssuePropertiesScheme struct {
	Keys []struct {
		Self string `json:"self"`
		Key  string `json:"key"`
	} `json:"keys"`
}

// Returns the URLs and keys of an issue's properties.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-properties/#api-rest-api-3-issue-issueidorkey-properties-get
func (i *IssuePropertyService) Gets(ctx context.Context, issueKeyOrID string) (result *IssuePropertiesScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/properties", issueKeyOrID)
	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssuePropertiesScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes an issue's property.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-properties/#api-rest-api-3-issue-issueidorkey-properties-propertykey-delete
func (i *IssuePropertyService) Delete(ctx context.Context, issueKeyOrID, propertyKey string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/properties/%v", issueKeyOrID, propertyKey)
	request, err := i.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}
