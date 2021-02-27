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
		Size       int           `json:"size,omitempty"`
		Items      []*UserScheme `json:"items,omitempty"`
		MaxResults int           `json:"max-results,omitempty"`
		StartIndex int           `json:"start-index,omitempty"`
		EndIndex   int           `json:"end-index,omitempty"`
	} `json:"users,omitempty"`
	Expand string `json:"expand,omitempty"`
}

// Creates a group.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/groups#create-group
func (g *GroupService) Create(ctx context.Context, groupName string) (result *GroupScheme, response *Response, err error) {

	if len(groupName) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid groupName value")
	}

	payload := struct {
		Name string `json:"name"`
	}{
		Name: groupName,
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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/groups#remove-group
func (g *GroupService) Delete(ctx context.Context, groupName string) (response *Response, err error) {

	if len(groupName) == 0 {
		return nil, fmt.Errorf("error, please provide a valid groupName value")
	}

	params := url.Values{}
	params.Add("groupname", groupName)
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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/groups#bulk-groups
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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/groups#get-users-from-groups
func (g *GroupService) Members(ctx context.Context, groupName string, inactive bool, startAt, maxResults int) (result *GroupUsersScheme, response *Response, err error) {

	if len(groupName) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid groupName value")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("groupname", groupName)

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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/groups#add-user-to-group
func (g *GroupService) Add(ctx context.Context, groupName, accountID string) (result *GroupScheme, response *Response, err error) {

	if len(groupName) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid groupName value")
	}

	if len(accountID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid accountID value")
	}

	payload := struct {
		AccountID string `json:"accountId"`
	}{AccountID: accountID}

	params := url.Values{}
	params.Add("groupname", groupName)
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
// https://docs.go-atlassian.io/jira-software-cloud/groups#remove-user-from-group
func (g *GroupService) Remove(ctx context.Context, groupName, accountID string) (response *Response, err error) {

	if len(groupName) == 0 {
		return nil, fmt.Errorf("error, please provide a valid groupName value")
	}

	if len(accountID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid accountID value")
	}

	params := url.Values{}
	params.Add("groupname", groupName)
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
