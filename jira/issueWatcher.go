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
	Self       string               `json:"self,omitempty"`
	IsWatching bool                 `json:"isWatching,omitempty"`
	WatchCount int                  `json:"watchCount,omitempty"`
	Watchers   []*IssueDetailScheme `json:"watchers,omitempty"`
}

type IssueDetailScheme struct {
	Self         string `json:"self,omitempty"`
	Name         string `json:"name,omitempty"`
	Key          string `json:"key,omitempty"`
	AccountID    string `json:"accountId,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	AvatarUrls   struct {
		One6X16   string `json:"16x16,omitempty"`
		Two4X24   string `json:"24x24,omitempty"`
		Three2X32 string `json:"32x32,omitempty"`
		Four8X48  string `json:"48x48,omitempty"`
	} `json:"avatarUrls,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Active      bool   `json:"active,omitempty"`
	TimeZone    string `json:"timeZone,omitempty"`
	AccountType string `json:"accountType,omitempty"`
}

// Returns the watchers for an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/watcher#get-issue-watchers
func (w *WatcherService) Gets(ctx context.Context, issueKeyOrID string) (result *IssueWatcherScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid issueKeyOrID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/watchers", issueKeyOrID)

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/watcher#add-watcher
func (w *WatcherService) Add(ctx context.Context, issueKeyOrID string) (response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, fmt.Errorf("error!, please provide a valid issueKeyOrID value")
	}

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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/watcher#delete-watcher
func (w *WatcherService) Delete(ctx context.Context, issueKeyOrID, accountID string) (response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, fmt.Errorf("error!, please provide a valid issueKeyOrID value")
	}

	if len(accountID) == 0 {
		return nil, fmt.Errorf("error!, please provide a valid accountID value")
	}

	params := url.Values{}
	params.Add("accountId", accountID)

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/watchers?%v", issueKeyOrID, params.Encode())

	request, err := w.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = w.client.Do(request)
	if err != nil {
		return
	}

	return
}
