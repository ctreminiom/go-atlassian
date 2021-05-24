package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type UserSearchService struct{ client *Client }

// Returns a list of users who can be assigned issues in one or more projects.
// The list may be restricted to users whose attributes match a string.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/users/search#find-users-assignable-to-projects
func (u *UserSearchService) Projects(ctx context.Context, accountID string, projectKeys []string, startAt, maxResults int) (result *[]UserScheme, response *Response, err error) {

	if len(projectKeys) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid projectKeys values")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(accountID) != 0 {
		params.Add("accountId", accountID)
	}

	var keys string
	for index, value := range projectKeys {

		if index == 0 {
			keys = value
			continue
		}

		keys += "," + value
	}

	if len(keys) != 0 {
		params.Add("projectKeys", keys)
	}

	var endpoint = fmt.Sprintf("rest/api/3/user/assignable/multiProjectSearch?%v", params.Encode())

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.Do(request)
	if err != nil {
		return
	}

	result = new([]UserScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns a list of users that match the search string and property.
// This operation takes the users in the range defined by startAt and maxResults, up to the thousandth user,
// and then returns only the users from that range that match the search string and property.
// This means the operation usually returns fewer users than specified in maxResults
// Docs: https://docs.go-atlassian.io/jira-software-cloud/users/search#find-users
func (u *UserSearchService) Do(ctx context.Context, accountID, query string, startAt, maxResults int) (result *[]UserScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(accountID) != 0 {
		params.Add("accountId", accountID)
	}

	if len(query) != 0 {
		params.Add("query", query)
	}

	var endpoint = fmt.Sprintf("rest/api/3/user/search?%v", params.Encode())

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.Do(request)
	if err != nil {
		return
	}

	result = new([]UserScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
