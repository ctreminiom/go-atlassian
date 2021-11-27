package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type GroupService struct{ client *Client }

// Create creates a group.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/groups#create-group
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-post
func (g *GroupService) Create(ctx context.Context, groupName string) (result *models.GroupScheme, response *ResponseScheme, err error) {

	if len(groupName) == 0 {
		return nil, nil, models.ErrNoGroupNameError
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
		return nil, models.ErrNoGroupNameError
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

// Bulk returns a paginated list of groups.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/groups#bulk-groups
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-bulk-get
// NOTE: Experimental Endpoint
func (g *GroupService) Bulk(ctx context.Context, options *models.GroupBulkOptionsScheme, startAt, maxResults int) (
	result *models.BulkGroupScheme, response *ResponseScheme, err error) {

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

// Members returns a paginated list of all users in a group.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/groups#get-users-from-groups
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-groups/#api-rest-api-3-group-member-get
func (g *GroupService) Members(ctx context.Context, groupName string, inactive bool, startAt, maxResults int) (
	result *models.GroupMemberPageScheme, response *ResponseScheme, err error) {

	if len(groupName) == 0 {
		return nil, nil, models.ErrNoGroupNameError
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
func (g *GroupService) Add(ctx context.Context, groupName, accountID string) (result *models.GroupScheme,
	response *ResponseScheme, err error) {

	if len(groupName) == 0 {
		return nil, nil, models.ErrNoGroupNameError
	}

	if len(accountID) == 0 {
		return nil, nil, models.ErrNoGroupIDError
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
		return nil, models.ErrNoGroupNameError
	}

	if len(accountID) == 0 {
		return nil, models.ErrNoGroupIDError
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
