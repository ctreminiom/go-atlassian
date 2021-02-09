package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type GroupService struct{ client *Client }

type GroupScheme struct {
	Name  string `json:"name,omitempty"`
	Self  string `json:"self,omitempty"`
	Users struct {
		Size       int          `json:"size,omitempty"`
		Items      []UserScheme `json:"items,omitempty"`
		MaxResults int          `json:"max-results,omitempty"`
		StartIndex int          `json:"start-index,omitempty"`
		EndIndex   int          `json:"end-index,omitempty"`
	} `json:"users,omitempty"`
	Expand string `json:"expand,omitempty"`
}

// Creates a group.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-post
func (g *GroupService) Create(ctx context.Context, name string) (result *GroupScheme, response *Response, err error) {

	payload := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	var endpoint = "rest/api/3/group"
	request, err := g.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = g.client.Do(request)
	if err != nil {
		return
	}

	result = new(GroupScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Deletes a group.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-delete
func (g *GroupService) Delete(ctx context.Context, name string) (response *Response, err error) {

	params := url.Values{}
	params.Add("groupname", name)
	var endpoint = fmt.Sprintf("rest/api/3/group?%v", params.Encode())

	request, err := g.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = g.client.Do(request)
	if err != nil {
		return
	}

	return
}

type BulkGroupScheme struct {
	MaxResults int  `json:"maxResults,omitempty"`
	StartAt    int  `json:"startAt,omitempty"`
	Total      int  `json:"total,omitempty"`
	IsLast     bool `json:"isLast,omitempty"`
	Values     []struct {
		Name    string `json:"name,omitempty"`
		GroupID string `json:"groupId,omitempty"`
	} `json:"values,omitempty"`
}

type GroupBulkOptionsScheme struct {
	GroupIDs   []string
	GroupNames []string
}

// Returns a paginated list of groups.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-bulk-get
func (g *GroupService) Bulk(ctx context.Context, options *GroupBulkOptionsScheme, startAt, maxResults int) (result *BulkGroupScheme, response *Response, err error) {

	if options == nil {
		return nil, nil, fmt.Errorf("error, options value is nil, please provide a valid GroupBulkOptionsScheme pointer")
	}

	params := url.Values{}

	for _, groupID := range options.GroupIDs {
		params.Add("groupId", groupID)
	}

	for _, groupName := range options.GroupNames {
		params.Add("groupName", groupName)
	}

	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("rest/api/3/group/bulk?%v", params.Encode())

	request, err := g.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = g.client.Do(request)
	if err != nil {
		return
	}

	result = new(BulkGroupScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

type GroupUsersScheme struct {
	Self       string `json:"self,omitempty"`
	NextPage   string `json:"nextPage,omitempty"`
	MaxResults int    `json:"maxResults,omitempty"`
	StartAt    int    `json:"startAt,omitempty"`
	Total      int    `json:"total,omitempty"`
	IsLast     bool   `json:"isLast,omitempty"`
	Values     []struct {
		Self         string `json:"self,omitempty"`
		Name         string `json:"name,omitempty"`
		Key          string `json:"key,omitempty"`
		AccountID    string `json:"accountId,omitempty"`
		EmailAddress string `json:"emailAddress,omitempty"`
		AvatarUrls   struct {
		} `json:"avatarUrls,omitempty"`
		DisplayName string `json:"displayName,omitempty"`
		Active      bool   `json:"active,omitempty"`
		TimeZone    string `json:"timeZone,omitempty"`
		AccountType string `json:"accountType,omitempty"`
	} `json:"values,omitempty"`
}

// Returns a paginated list of all users in a group.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-member-get
func (g *GroupService) Members(ctx context.Context, group string, inactive bool, startAt, maxResults int) (result *GroupUsersScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("groupname", group)

	if inactive {
		params.Add("includeInactiveUsers", "true")
	}

	var endpoint = fmt.Sprintf("rest/api/3/group/member?%v", params.Encode())

	request, err := g.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = g.client.Do(request)
	if err != nil {
		return
	}

	result = new(GroupUsersScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Adds a user to a group.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-user-post
func (g *GroupService) Add(ctx context.Context, group, accountID string) (result *GroupScheme, response *Response, err error) {

	payload := struct {
		AccountID string `json:"accountId"`
	}{AccountID: accountID}

	params := url.Values{}
	params.Add("groupname", group)
	var endpoint = fmt.Sprintf("rest/api/3/group/user?%v", params.Encode())

	request, err := g.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = g.client.Do(request)
	if err != nil {
		return
	}

	result = new(GroupScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Removes a user from a group.
// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-user-delete
func (g *GroupService) Remove(ctx context.Context, group, accountID string) (response *Response, err error) {

	params := url.Values{}
	params.Add("groupname", group)
	params.Add("accountId", accountID)
	var endpoint = fmt.Sprintf("rest/api/3/group/user?%v", params.Encode())

	request, err := g.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = g.client.Do(request)
	if err != nil {
		return
	}

	return
}
