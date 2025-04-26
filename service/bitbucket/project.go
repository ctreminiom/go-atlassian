package bitbucket

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// ProjectConnector defines the interface for project-related operations in Bitbucket.
type ProjectConnector interface {

	// Create creates a new project in the specified workspace.
	//
	// 	POST /2.0/workspaces/{workspace}/projects
	//
	// 	Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project will be created.
	// 	- payload: The project details to be created.
	//
	// 	Returns:
	// 	- A pointer to the created BitbucketProjectScheme.
	// 	- A pointer to the response scheme.
	// 	- An error if the creation fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Create(ctx context.Context, workspace string, payload *models.BitbucketProjectScheme) (*models.BitbucketProjectScheme, *models.ResponseScheme, error)

	// Get retrieves a project from the specified workspace using the project key.
	//
	// 	GET /2.0/workspaces/{workspace}/projects/{project_key}
	//
	// 	Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project is located.
	// 	- projectKey: The key of the project to be retrieved.
	//
	// 	Returns:
	// 	- A pointer to the retrieved BitbucketProjectScheme.
	// 	- A pointer to the response scheme.
	// 	- An error if the retrieval fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Get(ctx context.Context, workspace, projectKey string) (*models.BitbucketProjectScheme, *models.ResponseScheme, error)

	// Update updates an existing project in the specified workspace using the project key.
	//
	// Since this endpoint can be used to both update and to create a project, the request body depends on the intent.
	//
	//
	// 	Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project is located.
	// 	- projectKey: The key of the project to be updated.
	// 	- payload: The project details to be updated.
	//
	// 	Returns:
	// 	- A pointer to the updated BitbucketProjectScheme.
	// 	- A pointer to the response scheme.
	// 	- An error if the update fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Update(ctx context.Context, workspace, projectKey string, payload *models.BitbucketProjectScheme) (*models.BitbucketProjectScheme, *models.ResponseScheme, error)

	// Delete removes a project from the specified workspace using the project key.
	//
	// This is an irreversible operation.
	// You cannot delete a project that still contains repositories.
	// To delete the project, delete or transfer the repositories first.
	//
	//	DELETE /2.0/workspaces/{workspace}/projects/{project_key}
	//
	// 	Parameters:
	// 	- ctx: The context for the request.
	// 	- workspace: The workspace identifier where the project is located.
	// 	- projectKey: The key of the project to be deleted.
	//
	// 	Returns:
	// 	- A pointer to the response scheme.
	// 	- An error if the deletion fails.
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace
	Delete(ctx context.Context, workspace, projectKey string) (*models.ResponseScheme, error)
}
