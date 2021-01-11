package jira

import (
	"context"
	"fmt"
	"net/http"
)

type IssueLinkService struct {
	client *Client
	Type   *IssueLinkTypeService
}

type issueLinkPayloadScheme struct {
	OutwardIssue struct {
		Key string `json:"key"`
	} `json:"outwardIssue"`
	InwardIssue struct {
		Key string `json:"key"`
	} `json:"inwardIssue"`
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

// Creates a link between two issues. Use this operation to indicate a relationship between two issues
// and optionally add a comment to the from (outward) issue.
// To use this resource the site must have Issue Linking enabled.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-rest-api-3-issuelink-post
func (i *IssueLinkService) Create(ctx context.Context, linkType, inwardIssue, outwardIssue string) (response *Response, err error) {

	payload := issueLinkPayloadScheme{
		OutwardIssue: struct {
			Key string `json:"key"`
		}{Key: outwardIssue},
		InwardIssue: struct {
			Key string `json:"key"`
		}{Key: inwardIssue},
		Type: struct {
			Name string `json:"name"`
		}{Name: linkType},
	}

	var endpoint = "rest/api/3/issueLink"
	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Returns an issue link.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-rest-api-3-issuelink-linkid-get
func (i *IssueLinkService) Get(ctx context.Context, linkID string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issueLink/%v", linkID)

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Deletes an issue link.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-links/#api-rest-api-3-issuelink-linkid-delete
func (i *IssueLinkService) Delete(ctx context.Context, linkID string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issueLink/%v", linkID)

	request, err := i.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}
