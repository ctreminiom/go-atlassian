package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
)

func NewWatcherService(client service.Connector, version string) (*WatcherService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &WatcherService{
		internalClient: &internalWatcherImpl{c: client, version: version},
	}, nil
}

type WatcherService struct {
	internalClient jira.WatcherConnector
}

// Gets returns the watchers for an issue.
//
// GET /rest/api/{2-3}/issue/{issueIdOrKey}/watchers
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/watcher#get-issue-watchers
func (w *WatcherService) Gets(ctx context.Context, issueKeyOrId string) (*model.IssueWatcherScheme, *model.ResponseScheme, error) {
	return w.internalClient.Gets(ctx, issueKeyOrId)
}

// Add adds a user as a watcher of an issue by passing the account ID of the user.
//
// For example, "5b10ac8d82e05b22cc7d4ef5". If no user is specified the calling user is added.
//
// POST /rest/api/{2-3}/issue/{issueIdOrKey}/watchers
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/watcher#add-watcher
func (w *WatcherService) Add(ctx context.Context, issueKeyOrId string) (*model.ResponseScheme, error) {
	return w.internalClient.Add(ctx, issueKeyOrId)
}

// Delete deletes a user as a watcher of an issue.
//
// DELETE /rest/api/{2-3}/issue/{issueIdOrKey}/watchers
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/watcher#delete-watcher
func (w *WatcherService) Delete(ctx context.Context, issueKeyOrId, accountId string) (*model.ResponseScheme, error) {
	return w.internalClient.Delete(ctx, issueKeyOrId, accountId)
}

type internalWatcherImpl struct {
	c       service.Connector
	version string
}

func (i *internalWatcherImpl) Gets(ctx context.Context, issueKeyOrId string) (*model.IssueWatcherScheme, *model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/watchers", i.version, issueKeyOrId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	watchers := new(model.IssueWatcherScheme)
	response, err := i.c.Call(request, watchers)
	if err != nil {
		return nil, response, err
	}

	return watchers, response, nil
}

func (i *internalWatcherImpl) Add(ctx context.Context, issueKeyOrId string) (*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/watchers", i.version, issueKeyOrId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalWatcherImpl) Delete(ctx context.Context, issueKeyOrId, accountId string) (*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	if accountId == "" {
		return nil, model.ErrNoAccountIDError
	}

	params := url.Values{}
	params.Add("accountId", accountId)

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/watchers?%v", i.version, issueKeyOrId, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
