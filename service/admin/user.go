package admin

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// UserConnector represents the cloud admin users admin endpoints.
// Use it to search, get, create, delete, and change users.
type UserConnector interface {

	// Permissions returns the set of permissions you have for managing the specified Atlassian account
	//
	// GET /users/{account_id}/manage
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/user#get-user-management-permissions
	Permissions(ctx context.Context, accountID string, privileges []string) (*model.AdminUserPermissionScheme, *model.ResponseScheme, error)

	// Get returns information about a single Atlassian account by ID
	//
	// GET /users/{account_id}/manage/profile
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/user#get-profile
	Get(ctx context.Context, accountID string) (*model.AdminUserScheme, *model.ResponseScheme, error)

	// Update updates fields in a user account. The profile.write privilege details which fields you can change.
	//
	// PATCH /users/{account_id}/manage/profile
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/user#update-profile
	Update(ctx context.Context, accountID string, payload map[string]interface{}) (*model.AdminUserScheme, *model.ResponseScheme, error)

	// Disable deactivate the specified user account.
	//
	// The permission to make use of this resource is exposed by the lifecycle.enablement privilege.
	//
	// You can optionally set a message associated with the block.
	//
	// If none is supplied, a default message will be used.
	//
	// POST /users/{account_id}/manage/lifecycle/disable
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/user#disable-a-user
	Disable(ctx context.Context, accountID, message string) (*model.ResponseScheme, error)

	// Enable activates the specified user account.
	//
	// The permission to make use of this resource is exposed by the lifecycle.enablement privilege.
	//
	// POST /users/{account_id}/manage/lifecycle/enable
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/user#enable-a-user
	Enable(ctx context.Context, accountID string) (*model.ResponseScheme, error)
}
