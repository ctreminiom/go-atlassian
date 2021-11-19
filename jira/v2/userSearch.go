package v2

import (
	"context"
	"fmt"
	models2 "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type UserSearchService struct{ client *Client }

// Projects returns a list of users who can be assigned issues in one or more projects.
// The list may be restricted to users whose attributes match a string.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/users/search#find-users-assignable-to-projects
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-user-search/#api-rest-api-2-user-assignable-multiprojectsearch-get
func (u *UserSearchService) Projects(ctx context.Context, accountID string, projectKeys []string, startAt, maxResults int) (
	result []*models2.UserScheme, response *ResponseScheme, err error) {

	if len(projectKeys) == 0 {
		return nil, nil, models2.ErrNoProjectKeySliceError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(accountID) != 0 {
		params.Add("accountId", accountID)
	}

	if len(projectKeys) != 0 {
		params.Add("projectKeys", strings.Join(projectKeys, ","))
	}

	var endpoint = fmt.Sprintf("rest/api/2/user/assignable/multiProjectSearch?%v", params.Encode())

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Do return a list of users that match the search string and property.
// This operation takes the users in the range defined by startAt and maxResults, up to the thousandth user,
// and then returns only the users from that range that match the search string and property.
// This means the operation usually returns fewer users than specified in maxResults
// Docs: https://docs.go-atlassian.io/jira-software-cloud/users/search#find-users
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-user-search/#api-rest-api-2-user-search-get
func (u *UserSearchService) Do(ctx context.Context, accountID, query string, startAt, maxResults int) (result []*models2.UserScheme,
	response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(accountID) != 0 {
		params.Add("accountId", accountID)
	}

	if len(query) != 0 {
		params.Add("query", query)
	}

	var endpoint = fmt.Sprintf("rest/api/2/user/search?%v", params.Encode())

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
