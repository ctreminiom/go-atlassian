package bitbucket

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type ProjectReviewerConnector interface {

	// Gets returns a list of all default reviewers for a project.
	//
	// This is a list of users that will be added as default reviewers to pull requests for any repository within the project.
	//
	// GET /2.0/workspaces/{workspace}/projects/{project_key}/default-reviewers
	//
	// 	Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project is located.
	// 	- projectKey: The key of the project to be updated.
	//
	// 	Returns:
	// 	- A pointer to the updated ProjectReviewerPageScheme.
	// 	- A pointer to the response scheme.
	// 	- An error if the update fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Gets(ctx context.Context, workspace, projectKey string) (*models.ProjectReviewerPageScheme, *models.ResponseScheme, error)

	// Get retrieves a specific project reviewer based on the provided parameters.
	//
	// GET /2.0/workspaces/{workspace}/projects/{project_key}/default-reviewers/{selected_user}
	//
	// Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project is located.
	// 	- projectKey: The key of the project to be updated.
	// 	- userSlug: The slug of the user to be retrieved.
	//
	// Returns:
	// 	- A pointer to the ProjectReviewerScheme.
	// 	- A pointer to the response scheme.
	// 	- An error if the retrieval fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Get(ctx context.Context, workspace, projectKey, userSlug string) (*models.ProjectReviewerScheme, *models.ResponseScheme, error)

	// Add adds a specific user as a default reviewer for a project.
	//
	// PUT /2.0/workspaces/{workspace}/projects/{project_key}/default-reviewers/{selected_user}
	//
	// Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project is located.
	// 	- projectKey: The key of the project to be updated.
	// 	- userSlug: The slug of the user to be added as a default reviewer.
	//
	// Returns:
	// 	- A pointer to the ProjectReviewerScheme.
	// 	- A pointer to the response scheme.
	// 	- An error if the addition fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Add(ctx context.Context, workspace, projectKey, userSlug string) (*models.ProjectReviewerScheme, *models.ResponseScheme, error)

	// Remove deletes a specific user from the default reviewers for a project.
	//
	// DELETE /2.0/workspaces/{workspace}/projects/{project_key}/default-reviewers/{selected_user}
	//
	// Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project is located.
	// 	- projectKey: The key of the project to be updated.
	// 	- userSlug: The slug of the user to be removed as a default reviewer.
	//
	// Returns:
	// 	- A pointer to the response scheme.
	// 	- An error if the removal fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Remove(ctx context.Context, workspace, projectKey, userSlug string) (*models.ResponseScheme, error)
}
