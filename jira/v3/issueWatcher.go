package v3

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/internal/infra/models/jira"
	"net/http"
	"net/url"
)

type WatcherService struct{ client *Client }

type IssueWatcherScheme struct {
	Self       string              `json:"self,omitempty"`
	IsWatching bool                `json:"isWatching,omitempty"`
	WatchCount int                 `json:"watchCount,omitempty"`
	Watchers   []*UserDetailScheme `json:"watchers,omitempty"`
}

type UserDetailScheme struct {
	Self         string `json:"self,omitempty"`
	Name         string `json:"name,omitempty"`
	Key          string `json:"key,omitempty"`
	AccountID    string `json:"accountId,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Active       bool   `json:"active,omitempty"`
	TimeZone     string `json:"timeZone,omitempty"`
	AccountType  string `json:"accountType,omitempty"`
}

// Gets returns the watchers for an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/watcher#get-issue-watchers
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-watchers/#api-rest-api-3-issue-issueidorkey-watchers-get
func (w *WatcherService) Gets(ctx context.Context, issueKeyOrID string) (result *IssueWatcherScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, notIssueKeyOrIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/watchers", issueKeyOrID)

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = w.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Add adds a user as a watcher of an issue by passing the account ID of the user.
// For example, "5b10ac8d82e05b22cc7d4ef5". If no user is specified the calling user is added.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/watcher#add-watcher
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-watchers/#api-rest-api-3-issue-issueidorkey-watchers-post
func (w *WatcherService) Add(ctx context.Context, issueKeyOrID string) (response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, notIssueKeyOrIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/watchers", issueKeyOrID)

	request, err := w.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	response, err = w.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Delete deletes a user as a watcher of an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/watcher#delete-watcher
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-watchers/#api-rest-api-3-issue-issueidorkey-watchers-delete
func (w *WatcherService) Delete(ctx context.Context, issueKeyOrID, accountID string) (response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, notIssueKeyOrIDError
	}

	if len(accountID) == 0 {
		return nil, models.ErrNoGroupNameError
	}

	params := url.Values{}
	params.Add("accountId", accountID)

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/watchers?%v", issueKeyOrID, params.Encode())

	request, err := w.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = w.client.call(request, nil)
	if err != nil {
		return
	}

	return
}
