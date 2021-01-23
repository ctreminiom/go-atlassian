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
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-user-search/#api-rest-api-3-user-assignable-multiprojectsearch-get
func (u *UserSearchService) Projects(ctx context.Context, accountID string, projectKeys []string, startAt, maxResults int) (result *[]UserScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("accountId", accountID)

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

// Returns a list of users that can be assigned to an issue.
// Use this operation to find the list of users who can be assigned to:
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-user-search/#api-rest-api-3-user-assignable-search-get
func (u *UserSearchService) Issues(ctx context.Context, accountID, issueKey string, startAt, maxResults int) (result *[]UserScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("accountId", accountID)
	params.Add("issueKey", issueKey)

	var endpoint = fmt.Sprintf("rest/api/3/user/assignable/search?%v", params.Encode())

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
