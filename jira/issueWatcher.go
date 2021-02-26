package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type WatcherService struct{ client *Client }

type IssueWatcherScheme struct {
	Self       string `json:"self,omitempty"`
	IsWatching bool   `json:"isWatching,omitempty"`
	WatchCount int    `json:"watchCount,omitempty"`
	Watchers   []struct {
		Self        string `json:"self,omitempty"`
		AccountID   string `json:"accountId,omitempty"`
		DisplayName string `json:"displayName,omitempty"`
		Active      bool   `json:"active,omitempty"`
	} `json:"watchers,omitempty"`
}

// Returns the watchers for an issue.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-watchers/#api-rest-api-3-issue-issueidorkey-watchers-get
func (w *WatcherService) Get(ctx context.Context, issueKeyOrID string) (result *IssueWatcherScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/watchers", issueKeyOrID)

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = w.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueWatcherScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Adds a user as a watcher of an issue by passing the account ID of the user.
// For example, "5b10ac8d82e05b22cc7d4ef5". If no user is specified the calling user is added.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-watchers/#api-rest-api-3-issue-issueidorkey-watchers-post
func (w *WatcherService) Add(ctx context.Context, issueKeyOrID string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/watchers", issueKeyOrID)

	request, err := w.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err = w.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Deletes a user as a watcher of an issue.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-watchers/#api-rest-api-3-issue-issueidorkey-watchers-delete
func (w *WatcherService) Delete(ctx context.Context, issueKeyOrID, accountID string) (response *Response, err error) {

	params := url.Values{}
	params.Add("accountId", accountID)

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/watchers?%v", issueKeyOrID, params.Encode())

	request, err := w.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = w.client.Do(request)
	if err != nil {
		return
	}

	return
}
