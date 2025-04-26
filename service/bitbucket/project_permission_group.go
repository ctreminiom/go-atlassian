package bitbucket

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type ProjectGroupPermissionConnector interface {

	// Gets returns a paginated list of explicit group permissions for the given project.
	//
	// GET /2.0/workspaces/{workspace}/projects/{project_key}/permissions-config/groups
	//
	// Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project is located.
	// 	- projectKey: The key of the project
	//
	// Returns:
	// 	- A pointer to the updated BitbucketProjectGroupPermissionPageScheme.
	// 	- A pointer to the response scheme.
	// 	- An error if the update fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Gets(ctx context.Context, workspace, projectKey string) (*models.BitbucketProjectGroupPermissionPageScheme, *models.ResponseScheme, error)

	// Get retrieves the explicit group permissions for a specific group in the given project.
	//
	// GET /2.0/workspaces/{workspace}/projects/{project_key}/permissions-config/groups/{group_slug}
	//
	// Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project is located.
	// 	- projectKey: The key of the project
	// 	- groupSlug: The slug of the group whose permissions are to be retrieved.
	//
	// Returns:
	// 	- A pointer to the ProjectGroupPermissionScheme.
	// 	- A pointer to the response scheme.
	// 	- An error if the retrieval fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Get(ctx context.Context, workspace, projectKey, groupSlug string) (*models.ProjectGroupPermissionScheme, *models.ResponseScheme, error)

	// Update updates the explicit group permissions for a specific group in the given project.
	//
	// PUT /2.0/workspaces/{workspace}/projects/{project_key}/permissions-config/groups/{group_slug}
	//
	// Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project is located.
	// 	- projectKey: The key of the project.
	// 	- groupSlug: The slug of the group whose permissions are to be updated.
	// 	- permission: The new permission level to be set for the group.
	//
	// Returns:
	// 	- A pointer to the ProjectGroupPermissionScheme.
	// 	- A pointer to the response scheme.
	// 	- An error if the update fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Update(ctx context.Context, workspace, projectKey, groupSlug, permission string) (*models.ProjectGroupPermissionScheme, *models.ResponseScheme, error)

	// Delete removes the explicit group permissions for a specific group in the given project.
	//
	// Only users with admin permission for the project may access this resource.
	//
	// DELETE /2.0/workspaces/{workspace}/projects/{project_key}/permissions-config/groups/{group_slug}
	//
	// Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project is located.
	// 	- projectKey: The key of the project.
	// 	- groupSlug: The slug of the group whose permissions are to be deleted.
	//
	// Returns:
	// 	- A pointer to the response scheme.
	// 	- An error if the deletion fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Delete(ctx context.Context, workspace, projectKey, groupSlug string) (*models.ResponseScheme, error)
}
