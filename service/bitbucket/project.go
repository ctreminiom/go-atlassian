package bitbucket

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// ProjectConnector defines an interface for connecting with Bitbucket Cloud project endpoints.
type ProjectConnector interface {

	// Create creates a new project.
	//
	// POST /workspaces/{workspace}/projects
	Create(ctx context.Context, workspace string, payload *models.BitbucketProjectPayloadScheme) (
		*models.BitbucketProjectScheme, *models.ResponseScheme, error)

	// Get returns the requested project.
	//
	// GET /workspaces/{workspace}/projects/{project_key}
	Get(ctx context.Context, workspace, projectKey string) (*models.BitbucketProjectScheme, *models.ResponseScheme, error)

	// Update updates or create a new project.
	// Since this endpoint can be used to both update and to create a project, the request body depends on the intent.
	//
	// 	1.Creation: See the POST documentation for the project collection for an example of the request body.
	//	  Please note the key should not be specified in the body of request (since it is already present in the URL).
	//	  If the name is required, everything else is optional.
	//
	//	2.Update: The key is not required in the body (since it is already in the URL). The key may be specified
	//	  in the body, if the intent is to change the key itself. In such a scenario, the location of the project is
	//	  changed and is returned in the Location header of the response.
	//
	// PUT /workspaces/{workspace}/projects/{project_key}
	Update(ctx context.Context, workspace, projectKey string, payload *models.BitbucketProjectPayloadScheme) (*models.BitbucketProjectScheme,
		*models.ResponseScheme, error)

	// Delete deletes this project. This is an irreversible operation.
	//	1. You cannot delete a project that still contains repositories.
	//	2. To delete the project, delete or transfer the repositories first.
	//
	// DELETE /workspaces/{workspace}/projects/{project_key}
	Delete(ctx context.Context, workspace, projectKey string) (*models.ResponseScheme, error)
}

// ProjectReviewerConnector is an interface for connecting to Bitbucket Cloud project review endpoints.
//
// It provides methods for interacting with project reviews such as retrieving reviews, creating reviews, and updating reviews.
type ProjectReviewerConnector interface {

	// Gets return a list of all default reviewers for a project.
	//
	// This is a list of users that will be added as default reviewers to pull requests for any repository within the project.
	//
	// GET /workspaces/{workspace}/projects/{project_key}/default-reviewers
	Gets(ctx context.Context, workspace, projectKey string) (*models.ProjectReviewersPageScheme, *models.ResponseScheme, error)

	// Get returns the specified default reviewer.
	//
	// GET /workspaces/{workspace}/projects/{project_key}/default-reviewers/{selected_user}
	Get(ctx context.Context, workspace, projectKey, userSlug string) (*models.ProjectReviewerScheme, *models.ResponseScheme,
		error)

	// Add adds the specified user to the project's list of default reviewers.
	//
	//	1.The method is idempotent.
	//	2.Accepts an optional body containing the uuid of the user to be added.
	//
	// PUT /workspaces/{workspace}/projects/{project_key}/default-reviewers/{selected_user}
	Add(ctx context.Context, workspace, projectKey, userSlug string) (*models.ProjectReviewerScheme, *models.ResponseScheme,
		error)

	// Remove removes a default reviewer from the project.
	//
	// DELETE /workspaces/{workspace}/projects/{project_key}/default-reviewers/{selected_user}
	Remove(ctx context.Context, workspace, projectKey, accountUser string) (*models.ResponseScheme, error)
}

// ProjectGroupPermissionConnector represents an interface for connecting to Bitbucket Cloud
// project review endpoints to manage group permissions.
type ProjectGroupPermissionConnector interface {

	// Gets returns a paginated list of explicit group permissions for the given project.
	//	1.This endpoint does not support BBQL features.
	//
	// GET /workspaces/{workspace}/projects/{project_key}/permissions-config/groups
	Gets(ctx context.Context, workspace, projectKey string)

	// Get returns the group permission for a given group and project. Only users with admin permission for the project
	// may access this resource.
	//
	// Permissions can be:
	//
	//	1.admin
	//	2.create-repo
	//	3.write
	//	4.read
	//	5.none
	//
	// GET /workspaces/{workspace}/projects/{project_key}/permissions-config/groups/{group_slug}
	Get(ctx context.Context, workspace, projectKey, groupSlug string)

	// Update updates the group permission, or grants a new permission if one does not already exist.  Only users with
	// admin permission for the project may access this resource.
	//
	// Due to security concerns, the JWT and OAuth authentication methods are unsupported.
	// This is to ensure integrations and add-ons are not allowed to change permissions.
	//
	// Permissions can be:
	//
	//	1.admin
	//	2.create-repo
	//	3.write
	//	4.read
	// PUT /workspaces/{workspace}/projects/{project_key}/permissions-config/groups/{group_slug}
	Update(ctx context.Context, workspace, projectKey, groupSlug, permission string)

	// Delete deletes the project group permission between the requested project and group, if one exists.
	// Only users with admin permission for the project may access this resource.
	//
	// DELETE /workspaces/{workspace}/projects/{project_key}/permissions-config/groups/{group_slug}
	Delete(ctx context.Context, workspace, projectKey, groupSlug string)
}

// ProjectUserPermissionConnector represents an interface for connecting to Bitbucket Cloud
// project review endpoints to manage user permissions.
type ProjectUserPermissionConnector interface {

	// Gets returns a paginated list of explicit user permissions for the given project.
	//
	// This endpoint does not support BBQL features.
	//
	// GET /workspaces/{workspace}/projects/{project_key}/permissions-config/users
	Gets(ctx context.Context, workspace, projectKey string)

	// Get returns the explicit user permission for a given user and project.
	//
	// Only users with admin permission for the project may access this resource.
	//
	// Permissions can be:
	//
	// 	1.admin
	// 	2.create-repo
	// 	3.write
	// 	4.read
	// 	5.none
	//
	// GET /workspaces/{workspace}/projects/{project_key}/permissions-config/users/{selected_user_id}
	Get(ctx context.Context, workspace, projectKey, userSlug string)

	// Update updates the explicit user permission for a given user and project.
	// The selected user must be a member of the workspace, and cannot be the workspace owner.
	//
	// Only users with admin permission for the project may access this resource.
	//
	// Due to security concerns, the JWT and OAuth authentication methods are unsupported.
	// This is to ensure integrations and add-ons are not allowed to change permissions.
	//
	// Permissions can be:
	//
	//	1.admin
	//	2.create-repo
	//	3.write
	//	4.read
	//
	// PUT /workspaces/{workspace}/projects/{project_key}/permissions-config/users/{selected_user_id}
	Update(ctx context.Context, workspace, projectKey, userSlug, permission string)

	//Delete deletes the project user permission between the requested project and user, if one exists.
	//
	//	1.Only users with admin permission for the project may access this resource.
	// 	2.Due to security concerns, the JWT and OAuth authentication methods are unsupported.
	// 	3.This is to ensure integrations and add-ons are not allowed to change permissions.
	//
	// DELETE /workspaces/{workspace}/projects/{project_key}/permissions-config/users/{selected_user_id}
	Delete(ctx context.Context, workspace, projectKey, userSlug string)
}
