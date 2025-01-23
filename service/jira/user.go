package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type UserConnector interface {

	// Get returns a user
	//
	// GET /rest/api/{2-3}/user
	//
	// https://docs.go-atlassian.io/jira-software-cloud/users#get-user
	Get(ctx context.Context, accountID string, expand []string) (*model.UserScheme, *model.ResponseScheme, error)

	// Create creates a user. This resource is retained for legacy compatibility.
	//
	// As soon as a more suitable alternative is available this resource will be deprecated.
	//
	// The option is provided to set or generate a password for the user.
	//
	// When using the option to generate a password, by omitting password from the request, include "notification": "true" to ensure the user is
	//
	// sent an email advising them that their account is created.
	//
	// This email includes a link for the user to set their password. If the notification isn't sent for a generated password,
	//
	// the user will need to be sent a reset password request from Jira.
	//
	// POST /rest/api/{2-3}user
	//
	// https://docs.go-atlassian.io/jira-software-cloud/users#create-user
	Create(ctx context.Context, payload *model.UserPayloadScheme) (*model.UserScheme, *model.ResponseScheme, error)

	// Delete deletes a user.
	//
	// DELETE /rest/api/{2-3}/user
	//
	// https://docs.go-atlassian.io/jira-software-cloud/users#delete-user
	Delete(ctx context.Context, accountID string) (*model.ResponseScheme, error)

	// Find returns a paginated list of the users specified by one or more account IDs.
	//
	// GET /rest/api/{2-3}/user/bulk
	//
	// https://docs.go-atlassian.io/jira-software-cloud/users#bulk-get-users
	Find(ctx context.Context, accountIDs []string, startAt, maxResults int) (*model.UserSearchPageScheme, *model.ResponseScheme,
		error)

	// Groups returns the groups to which a user belongs.
	//
	// GET /rest/api/{2-3}/user/groups
	//
	// https://docs.go-atlassian.io/jira-software-cloud/users#get-user-groups
	Groups(ctx context.Context, accountIDs string) ([]*model.UserGroupScheme, *model.ResponseScheme, error)

	// Gets returns a list of all (active and inactive) users.
	//
	// GET /rest/api/{2-3}/users/search
	//
	// https://docs.go-atlassian.io/jira-software-cloud/users#get-all-users
	Gets(ctx context.Context, startAt, maxResults int) ([]*model.UserScheme, *model.ResponseScheme, error)
}

type UserSearchConnector interface {

	// Projects returns a list of users who can be assigned issues in one or more projects.
	//
	// The list may be restricted to users whose attributes match a string.
	//
	// GET /rest/api/{2-3}/user/assignable/multiProjectSearch
	//
	// https://docs.go-atlassian.io/jira-software-cloud/users/search#find-users-assignable-to-projects
	Projects(ctx context.Context, accountID string, projectKeys []string, startAt, maxResults int) (
		[]*model.UserScheme, *model.ResponseScheme, error)

	// Do return a list of users that match the search string and property.
	//
	//
	// This operation takes the users in the range defined by startAt and maxResults, up to the thousandth user,
	//
	// and then returns only the users from that range that match the search string and property.
	//
	// This means the operation usually returns fewer users than specified in maxResults
	//
	// GET /rest/api/{2-3}/user/search
	//
	// https://docs.go-atlassian.io/jira-software-cloud/users/search#find-users
	Do(ctx context.Context, accountID, query string, startAt, maxResults int) ([]*model.UserScheme, *model.ResponseScheme, error)

	// Check returns a list of users who fulfill these criteria:
	//
	// 1. their user attributes match a search string.
	// 2. they have a set of permissions for a project or issue.
	//
	//
	// If no search string is provided, a list of all users with the permissions is returned.
	//
	// GET /rest/api/{2-3}/user/permission/search
	//
	// https://docs.go-atlassian.io/jira-software-cloud/users/search#find-users-with-permissions
	Check(ctx context.Context, permission string, options *model.UserPermissionCheckParamsScheme, startAt, maxResults int) ([]*model.UserScheme, *model.ResponseScheme, error)
}
