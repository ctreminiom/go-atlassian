package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewUserSearchService creates a new instance of UserSearchService.
func NewUserSearchService(client service.Connector, version string) (*UserSearchService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &UserSearchService{
		internalClient: &internalUserSearchImpl{c: client, version: version},
	}, nil
}

// UserSearchService provides methods to search for users in Jira Service Management.
type UserSearchService struct {
	// internalClient is the connector interface for user search operations.
	internalClient jira.UserSearchConnector
}

// Projects returns a list of users who can be assigned issues in one or more projects.
//
// The list may be restricted to users whose attributes match a string.
//
// GET /rest/api/{2-3}/user/assignable/multiProjectSearch
//
// https://docs.go-atlassian.io/jira-software-cloud/users/search#find-users-assignable-to-projects
func (u *UserSearchService) Projects(ctx context.Context, accountID string, projectKeys []string, startAt, maxResults int) ([]*model.UserScheme, *model.ResponseScheme, error) {
	return u.internalClient.Projects(ctx, accountID, projectKeys, startAt, maxResults)
}

// Do return a list of users that match the search string and property.
//
// This operation takes the users in the range defined by startAt and maxResults, up to the thousandth user,
//
// and then returns only the users from that range that match the search string and property.
//
// # This means the operation usually returns fewer users than specified in maxResults
//
// GET /rest/api/{2-3}/user/search
//
// https://docs.go-atlassian.io/jira-software-cloud/users/search#find-users
func (u *UserSearchService) Do(ctx context.Context, accountID, query string, startAt, maxResults int) ([]*model.UserScheme, *model.ResponseScheme, error) {
	return u.internalClient.Do(ctx, accountID, query, startAt, maxResults)
}

// Check returns a list of users who fulfill these criteria:
//
// 1. their user attributes match a search string.
// 2. they have a set of permissions for a project or issue.
//
// If no search string is provided, a list of all users with the permissions is returned.
//
// GET /rest/api/{2-3}/user/permission/search
//
// https://docs.go-atlassian.io/jira-software-cloud/users/search#find-users-with-permissions
func (u *UserSearchService) Check(ctx context.Context, permission string, options *model.UserPermissionCheckParamsScheme, startAt, maxResults int) ([]*model.UserScheme, *model.ResponseScheme, error) {
	return u.internalClient.Check(ctx, permission, options, startAt, maxResults)
}

type internalUserSearchImpl struct {
	c       service.Connector
	version string
}

func (i *internalUserSearchImpl) Check(ctx context.Context, permission string, options *model.UserPermissionCheckParamsScheme, startAt, maxResults int) ([]*model.UserScheme, *model.ResponseScheme, error) {

	if permission == "" {
		return nil, nil, model.ErrNoPermissionGrantID
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("permission", permission)

	if options != nil {

		if options.Query != "" {
			params.Add("query", options.Query)
		}

		if options.AccountID != "" {
			params.Add("accountId", options.AccountID)
		}

		if options.IssueKey != "" {
			params.Add("issueKey", options.IssueKey)
		}

		if options.ProjectKey != "" {
			params.Add("projectKey", options.ProjectKey)
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/user/permission/search?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*model.UserScheme
	response, err := i.c.Call(request, &users)
	if err != nil {
		return nil, response, err
	}

	return users, response, nil
}

func (i *internalUserSearchImpl) Projects(ctx context.Context, accountID string, projectKeys []string, startAt, maxResults int) ([]*model.UserScheme, *model.ResponseScheme, error) {

	if len(projectKeys) == 0 {
		return nil, nil, model.ErrNoProjectKeySlice
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if accountID != "" {
		params.Add("accountId", accountID)
	}

	if len(projectKeys) != 0 {
		params.Add("projectKeys", strings.Join(projectKeys, ","))
	}

	endpoint := fmt.Sprintf("rest/api/%v/user/assignable/multiProjectSearch?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*model.UserScheme
	response, err := i.c.Call(request, &users)
	if err != nil {
		return nil, response, err
	}

	return users, response, nil
}

func (i *internalUserSearchImpl) Do(ctx context.Context, accountID, query string, startAt, maxResults int) ([]*model.UserScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if accountID != "" {
		params.Add("accountId", accountID)
	}

	if query != "" {
		params.Add("query", query)
	}

	endpoint := fmt.Sprintf("rest/api/%v/user/search?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var users []*model.UserScheme
	response, err := i.c.Call(request, &users)
	if err != nil {
		return nil, response, err
	}

	return users, response, nil
}
