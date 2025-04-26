package bitbucket

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type ProjectUserPermissionConnector interface {

	// Gets returns a paginated list of explicit user permissions for the given project.
	//
	// GET /2.0/workspaces/{workspace}/projects/{project_key}/permissions-config/users
	//
	// Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project is located.
	// 	- projectKey: The key of the project
	//
	// Returns:
	// 	- A pointer to the updated BitbucketProjectUserPermissionPageScheme.
	// 	- A pointer to the response scheme.
	// 	- An error if the update fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Gets(ctx context.Context, workspace, projectKey string) (*models.BitbucketProjectUserPermissionPageScheme, *models.ResponseScheme, error)

	// Get retrieves the explicit user permissions for a specific user in the given project.
	//
	// Only users with admin permission for the project may access this resource.
	//
	// Permissions can be: "admin", "write", "read", "create-repo" or "none".
	//
	// GET /2.0/workspaces/{workspace}/projects/{project_key}/permissions-config/users/{selected_user_id}
	//
	// Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project is located.
	// 	- projectKey: The key of the project.
	// 	- userSlug: The slug of the user whose permissions are to be retrieved.
	//
	// Returns:
	// 	- A pointer to the BitbucketProjectUserPermissionScheme.
	// 	- A pointer to the response scheme.
	// 	- An error if the retrieval fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Get(ctx context.Context, workspace, projectKey, userSlug string) (*models.BitbucketProjectUserPermissionScheme, *models.ResponseScheme, error)

	// Update updates the explicit user permission for a given user and project.
	//
	// The selected user must be a member of the workspace, and cannot be the workspace
	//
	// Only users with admin permission for the project may access this resource.
	//
	// Due to security concerns, the JWT and OAuth authentication methods are unsupported. This is to ensure integrations and add-ons are not allowed to change permissions.
	//
	// PUT /2.0/workspaces/{workspace}/projects/{project_key}/permissions-config/users/{selected_user_id}
	//
	// Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project is located.
	// 	- projectKey: The key of the project.
	// 	- userSlug: The slug of the user whose permissions are to be updated.
	// 	- permission: The new permission level to be set for the user.
	//
	// Returns:
	// 	- A pointer to the BitbucketProjectUserPermissionScheme.
	// 	- A pointer to the response scheme.
	// 	- An error if the update fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Update(ctx context.Context, workspace, projectKey, userSlug, permission string) (*models.BitbucketProjectUserPermissionScheme, *models.ResponseScheme, error)

	// Delete removes the explicit user permissions for a specific user in the given project.
	//
	// DELETE /2.0/workspaces/{workspace}/projects/{project_key}/permissions-config/users/{selected_user_id}
	//
	// Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project is located.
	// 	- projectKey: The key of the project.
	// 	- userSlug: The slug of the user whose permissions are to be deleted.
	//
	// Returns:
	// 	- A pointer to the response scheme.
	// 	- An error if the deletion fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Delete(ctx context.Context, workspace, projectKey, userSlug string) (*models.ResponseScheme, error)
}
