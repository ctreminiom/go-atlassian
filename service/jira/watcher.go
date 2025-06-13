package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// WatcherConnector is an interface that defines the methods available from WatcherConnector API.
// Use it to get details of users watching an issue as well as start and stop a user watching an issue.
type WatcherConnector interface {

	// Gets returns the watchers for an issue.
	//
	// GET /rest/api/{2-3}/issue/{issueKeyOrID}/watchers
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/watcher#get-issue-watchers
	Gets(ctx context.Context, issueKeyOrID string) (*model.IssueWatcherScheme, *model.ResponseScheme, error)

	// Add adds a user as a watcher of an issue by passing the account ID of the user.
	//
	// For example, "5b10ac8d82e05b22cc7d4ef5". If no user is specified the calling user is added.
	//
	// POST /rest/api/{2-3}/issue/{issueKeyOrID}/watchers
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/watcher#add-watcher
	Add(ctx context.Context, issueKeyOrID string, accountID ...string) (*model.ResponseScheme, error)

	// Delete deletes a user as a watcher of an issue.
	//
	// DELETE /rest/api/{2-3}/issue/{issueKeyOrID}/watchers
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/watcher#delete-watcher
	Delete(ctx context.Context, issueKeyOrID, accountID string) (*model.ResponseScheme, error)
}
