package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
)

type WatcherService struct{ client *Client }

// Gets returns the watchers for an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/watcher#get-issue-watchers
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-watchers/#api-rest-api-3-issue-issueidorkey-watchers-get
func (w *WatcherService) Gets(ctx context.Context, issueKeyOrID string) (result *models.IssueWatcherScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, models.ErrNoIssueTypeIDError
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
		return nil, models.ErrNoIssueTypeIDError
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
		return nil, models.ErrNoIssueTypeIDError
	}

	if len(accountID) == 0 {
		return nil, models.ErrNoAccountIDError
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
