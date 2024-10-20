package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewGroupService creates a new instance of GroupService.
// It takes a service.Connector and a version string as input.
// Returns a pointer to GroupService and an error if the version is not provided.
func NewGroupService(client service.Connector, version string) (*GroupService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &GroupService{
		internalClient: &internalGroupServiceImpl{c: client, version: version},
	}, nil
}

// GroupService provides methods to manage groups in Jira Service Management.
type GroupService struct {
	// internalClient is the connector interface for group operations.
	internalClient jira.GroupConnector
}

// Delete deletes a group.
//
// DELETE /rest/api/{2-3}/group
//
// https://docs.go-atlassian.io/jira-software-cloud/groups#remove-group
func (g *GroupService) Delete(ctx context.Context, groupName string) (*model.ResponseScheme, error) {
	return g.internalClient.Delete(ctx, groupName)
}

// Bulk returns a paginated list of groups.
//
// GET /rest/api/{2-3}/group/bulk
//
// https://docs.go-atlassian.io/jira-software-cloud/groups#bulk-groups
func (g *GroupService) Bulk(ctx context.Context, options *model.GroupBulkOptionsScheme, startAt, maxResults int) (*model.BulkGroupScheme, *model.ResponseScheme, error) {
	return g.internalClient.Bulk(ctx, options, startAt, maxResults)
}

// Members returns a paginated list of all users in a group.
//
// GET /rest/api/{2-3}/group/member
//
// https://docs.go-atlassian.io/jira-software-cloud/groups#get-users-from-groups
func (g *GroupService) Members(ctx context.Context, groupName string, inactive bool, startAt, maxResults int) (*model.GroupMemberPageScheme, *model.ResponseScheme, error) {
	return g.internalClient.Members(ctx, groupName, inactive, startAt, maxResults)
}

// Add adds a user to a group.
//
// POST /rest/api/{2-3}/group/user
//
// https://docs.go-atlassian.io/jira-software-cloud/groups#add-user-to-group
func (g *GroupService) Add(ctx context.Context, groupName, accountID string) (*model.GroupScheme, *model.ResponseScheme, error) {
	return g.internalClient.Add(ctx, groupName, accountID)
}

// Remove removes a user from a group.
//
// DELETE /rest/api/{2-3}/group/user
//
// https://docs.go-atlassian.io/jira-software-cloud/groups#remove-user-from-group
func (g *GroupService) Remove(ctx context.Context, groupName, accountID string) (*model.ResponseScheme, error) {
	return g.internalClient.Remove(ctx, groupName, accountID)
}

// Create creates a group.
//
// POST /rest/api/{2-3}/group
//
// https://docs.go-atlassian.io/jira-software-cloud/groups#create-group
func (g *GroupService) Create(ctx context.Context, groupName string) (*model.GroupScheme, *model.ResponseScheme, error) {
	return g.internalClient.Create(ctx, groupName)
}

type internalGroupServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalGroupServiceImpl) Create(ctx context.Context, groupName string) (*model.GroupScheme, *model.ResponseScheme, error) {

	if groupName == "" {
		return nil, nil, model.ErrNoGroupName
	}

	endpoint := fmt.Sprintf("rest/api/%v/group", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"name": groupName})
	if err != nil {
		return nil, nil, err
	}

	group := new(model.GroupScheme)
	response, err := i.c.Call(request, group)
	if err != nil {
		return nil, response, err
	}

	return group, response, nil
}

func (i *internalGroupServiceImpl) Delete(ctx context.Context, groupName string) (*model.ResponseScheme, error) {

	if groupName == "" {
		return nil, model.ErrNoGroupName
	}

	params := url.Values{}
	params.Add("groupname", groupName)

	endpoint := fmt.Sprintf("rest/api/%v/group?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalGroupServiceImpl) Bulk(ctx context.Context, options *model.GroupBulkOptionsScheme, startAt, maxResults int) (*model.BulkGroupScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		for _, id := range options.GroupIDs {
			params.Add("groupId", id)
		}

		for _, name := range options.GroupNames {
			params.Add("groupName", name)
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/group/bulk?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BulkGroupScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalGroupServiceImpl) Members(ctx context.Context, groupName string, inactive bool, startAt, maxResults int) (*model.GroupMemberPageScheme, *model.ResponseScheme, error) {

	if groupName == "" {
		return nil, nil, model.ErrNoGroupName
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("groupname", groupName)
	params.Add("includeInactiveUsers", fmt.Sprintf("%v", inactive))

	endpoint := fmt.Sprintf("rest/api/%v/group/member?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.GroupMemberPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalGroupServiceImpl) Add(ctx context.Context, groupName, accountID string) (*model.GroupScheme, *model.ResponseScheme, error) {

	if groupName == "" {
		return nil, nil, model.ErrNoGroupName
	}

	if accountID == "" {
		return nil, nil, model.ErrNoAccountID
	}

	params := url.Values{}
	params.Add("groupname", groupName)
	endpoint := fmt.Sprintf("rest/api/%v/group/user?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"accountId": accountID})
	if err != nil {
		return nil, nil, err
	}

	group := new(model.GroupScheme)
	response, err := i.c.Call(request, group)
	if err != nil {
		return nil, response, err
	}

	return group, response, nil
}

func (i *internalGroupServiceImpl) Remove(ctx context.Context, groupName, accountID string) (*model.ResponseScheme, error) {

	if groupName == "" {
		return nil, model.ErrNoGroupName
	}

	if accountID == "" {
		return nil, model.ErrNoAccountID
	}

	params := url.Values{}
	params.Add("groupname", groupName)
	params.Add("accountId", accountID)
	endpoint := fmt.Sprintf("rest/api/%v/group/user?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
