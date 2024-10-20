package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// GroupConnector is an interface that defines the methods available from GroupConnector API.
type GroupConnector interface {

	// Create creates a group.
	//
	// POST /rest/api/{2-3}/group
	//
	// https://docs.go-atlassian.io/jira-software-cloud/groups#create-group
	Create(ctx context.Context, groupName string) (*model.GroupScheme, *model.ResponseScheme, error)

	// Delete deletes a group.
	//
	// DELETE /rest/api/{2-3}/group
	//
	// https://docs.go-atlassian.io/jira-software-cloud/groups#remove-group
	Delete(ctx context.Context, groupName string) (*model.ResponseScheme, error)

	// Bulk returns a paginated list of groups.
	//
	// GET /rest/api/{2-3}/group/bulk
	//
	// https://docs.go-atlassian.io/jira-software-cloud/groups#bulk-groups
	Bulk(ctx context.Context, options *model.GroupBulkOptionsScheme, startAt, maxResults int) (*model.BulkGroupScheme, *model.ResponseScheme, error)

	// Members returns a paginated list of all users in a group.
	//
	// GET /rest/api/{2-3}/group/member
	//
	// https://docs.go-atlassian.io/jira-software-cloud/groups#get-users-from-groups
	Members(ctx context.Context, groupName string, inactive bool, startAt, maxResults int) (*model.GroupMemberPageScheme, *model.ResponseScheme, error)

	// Add adds a user to a group.
	//
	// POST /rest/api/{2-3}/group/user
	//
	// https://docs.go-atlassian.io/jira-software-cloud/groups#add-user-to-group
	Add(ctx context.Context, groupName, accountID string) (*model.GroupScheme, *model.ResponseScheme, error)

	// Remove removes a user from a group.
	//
	// DELETE /rest/api/{2-3}/group/user
	//
	// https://docs.go-atlassian.io/jira-software-cloud/groups#remove-user-from-group
	Remove(ctx context.Context, groupName, accountID string) (*model.ResponseScheme, error)

	// TODO: GET /rest/api/3/groups/picker needs to be parsed
}
