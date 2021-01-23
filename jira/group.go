package jira

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type GroupService struct{ client *Client }

type GroupScheme struct {
	Name  string `json:"name"`
	Self  string `json:"self"`
	Users struct {
		Size  int `json:"size"`
		Items []struct {
			Self        string `json:"self"`
			AccountID   string `json:"accountId"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
		} `json:"items"`
		MaxResults int `json:"max-results"`
		StartIndex int `json:"start-index"`
		EndIndex   int `json:"end-index"`
	} `json:"users"`
	Expand string `json:"expand"`
}

// Creates a group.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-post
func (g *GroupService) Create(ctx context.Context, name string) (result *GroupScheme, response *Response, err error) {

	if ctx == nil {
		return nil, nil, errors.New("the context param is nil, please provide a valid one")
	}

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

	if len(response.BodyAsBytes) == 0 {
		return nil, nil, errors.New("unable to marshall the response body, the HTTP callback did not return any bytes")
	}

	result = new(GroupScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes a group.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-delete
func (g *GroupService) Delete(ctx context.Context, name string) (response *Response, err error) {

	if ctx == nil {
		return nil, errors.New("the context param is nil, please provide a valid one")
	}

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
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		Name    string `json:"name"`
		GroupID string `json:"groupId"`
	} `json:"values"`
}

type GroupBulkOptionsScheme struct {
	GroupIDs   []string
	GroupNames []string
}

// Returns a paginated list of groups.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-bulk-get
func (g *GroupService) Bulk(ctx context.Context, options *GroupBulkOptionsScheme, startAt, maxResults int) (result *BulkGroupScheme, response *Response, err error) {

	if ctx == nil {
		return nil, nil, errors.New("the context param is nil, please provide a valid one")
	}

	params := url.Values{}

	var groupIDs string
	for index, value := range options.GroupIDs {

		if index == 0 {
			groupIDs = value
		}

		groupIDs += "," + value
	}

	if len(groupIDs) != 0 {
		params.Add("groupId", groupIDs)
	}

	var groupNames string
	for index, value := range options.GroupNames {

		if index == 0 {
			groupNames = value
		}

		groupNames += "," + value
	}

	if len(groupNames) != 0 {
		params.Add("groupName", groupNames)
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

	if len(response.BodyAsBytes) == 0 {
		return nil, nil, errors.New("unable to marshall the response body, the HTTP callback did not return any bytes")
	}

	result = new(BulkGroupScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type GroupUsersScheme struct {
	Self       string `json:"self"`
	NextPage   string `json:"nextPage"`
	MaxResults int    `json:"maxResults"`
	StartAt    int    `json:"startAt"`
	Total      int    `json:"total"`
	IsLast     bool   `json:"isLast"`
	Values     []struct {
		Self         string `json:"self"`
		Name         string `json:"name"`
		Key          string `json:"key"`
		AccountID    string `json:"accountId"`
		EmailAddress string `json:"emailAddress"`
		AvatarUrls   struct {
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
		TimeZone    string `json:"timeZone"`
		AccountType string `json:"accountType"`
	} `json:"values"`
}

// Returns a paginated list of all users in a group.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-member-get
func (g *GroupService) Members(ctx context.Context, group string, inactive bool, startAt, maxResults int) (result *GroupUsersScheme, response *Response, err error) {

	if ctx == nil {
		return nil, nil, errors.New("the context param is nil, please provide a valid one")
	}

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

	if len(response.BodyAsBytes) == 0 {
		return nil, nil, errors.New("unable to marshall the response body, the HTTP callback did not return any bytes")
	}

	result = new(GroupUsersScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Adds a user to a group.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-user-post
func (g *GroupService) Add(ctx context.Context, group, accountID string) (result *GroupScheme, response *Response, err error) {

	if ctx == nil {
		return nil, nil, errors.New("the context param is nil, please provide a valid one")
	}

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

	if len(response.BodyAsBytes) == 0 {
		return nil, nil, errors.New("unable to marshall the response body, the HTTP callback did not return any bytes")
	}

	result = new(GroupScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Removes a user from a group.
// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-user-delete
func (g *GroupService) Remove(ctx context.Context, group, accountID string) (response *Response, err error) {

	if ctx == nil {
		return nil, errors.New("the context param is nil, please provide a valid one")
	}

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
