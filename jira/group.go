package jira

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type GroupService struct{ client *Client }

type GroupScheme struct {
	Name   string               `json:"name,omitempty"`
	Self   string               `json:"self,omitempty"`
	Users  *GroupUserPageScheme `json:"users,omitempty"`
	Expand string               `json:"expand,omitempty"`
}

type GroupUserPageScheme struct {
	Size       int           `json:"size,omitempty"`
	Items      []*UserScheme `json:"items,omitempty"`
	MaxResults int           `json:"max-results,omitempty"`
	StartIndex int           `json:"start-index,omitempty"`
	EndIndex   int           `json:"end-index,omitempty"`
}

// Create creates a group.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/groups#create-group
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-post
func (g *GroupService) Create(ctx context.Context, groupName string) (result *GroupScheme, response *ResponseScheme, err error) {

	if len(groupName) == 0 {
		return nil, nil, notGroupNameError
	}

	payload := struct {
		Name string `json:"name"`
	}{
		Name: groupName,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = "rest/api/3/group"
	request, err := g.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = g.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes a group.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/groups#remove-group
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-delete
func (g *GroupService) Delete(ctx context.Context, groupName string) (response *ResponseScheme, err error) {

	if len(groupName) == 0 {
		return nil, notGroupNameError
	}

	params := url.Values{}
	params.Add("groupname", groupName)
	var endpoint = fmt.Sprintf("rest/api/3/group?%v", params.Encode())

	request, err := g.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = g.client.call(request, nil)
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

// Bulk returns a paginated list of groups.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/groups#bulk-groups
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-bulk-get
// NOTE: Experimental Endpoint
func (g *GroupService) Bulk(ctx context.Context, options *GroupBulkOptionsScheme, startAt, maxResults int) (
	result *BulkGroupScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {
		for _, groupID := range options.GroupIDs {
			params.Add("groupId", groupID)
		}

		for _, groupName := range options.GroupNames {
			params.Add("groupName", groupName)
		}
	}

	var endpoint = fmt.Sprintf("rest/api/3/group/bulk?%v", params.Encode())

	request, err := g.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = g.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type GroupMemberPageScheme struct {
	Self       string                   `json:"self,omitempty"`
	NextPage   string                   `json:"nextPage,omitempty"`
	MaxResults int                      `json:"maxResults,omitempty"`
	StartAt    int                      `json:"startAt,omitempty"`
	Total      int                      `json:"total,omitempty"`
	IsLast     bool                     `json:"isLast,omitempty"`
	Values     []*GroupUserDetailScheme `json:"values,omitempty"`
}

type GroupUserDetailScheme struct {
	Self         string `json:"self"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	AccountID    string `json:"accountId"`
	EmailAddress string `json:"emailAddress"`
	DisplayName  string `json:"displayName"`
	Active       bool   `json:"active"`
	TimeZone     string `json:"timeZone"`
	AccountType  string `json:"accountType"`
}

// Members returns a paginated list of all users in a group.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/groups#get-users-from-groups
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-member-get
func (g *GroupService) Members(ctx context.Context, groupName string, inactive bool, startAt, maxResults int) (
	result *GroupMemberPageScheme, response *ResponseScheme, err error) {

	if len(groupName) == 0 {
		return nil, nil, notGroupNameError
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

	response, err = g.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Add adds a user to a group.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/groups#add-user-to-group
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-user-post
func (g *GroupService) Add(ctx context.Context, groupName, accountID string) (result *GroupScheme,
	response *ResponseScheme, err error) {

	if len(groupName) == 0 {
		return nil, nil, notGroupNameError
	}

	if len(accountID) == 0 {
		return nil, nil, notAccountIDError
	}

	payload := struct {
		AccountID string `json:"accountId"`
	}{
		AccountID: accountID,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	params := url.Values{}
	params.Add("groupname", groupName)
	var endpoint = fmt.Sprintf("rest/api/3/group/user?%v", params.Encode())

	request, err := g.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = g.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Remove removes a user from a group.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/groups#remove-user-from-group
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-user-delete
func (g *GroupService) Remove(ctx context.Context, groupName, accountID string) (response *ResponseScheme, err error) {

	if len(groupName) == 0 {
		return nil, notGroupNameError
	}

	if len(accountID) == 0 {
		return nil, notAccountIDError
	}

	params := url.Values{}
	params.Add("groupname", groupName)
	params.Add("accountId", accountID)
	var endpoint = fmt.Sprintf("rest/api/3/group/user?%v", params.Encode())

	request, err := g.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = g.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

var (
	notGroupNameError = fmt.Errorf("error, please provide a valid groupName value")
	notAccountIDError = fmt.Errorf("error, please provide a valid accountID value")
)
