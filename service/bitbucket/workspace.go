package bitbucket

import (
	"context"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// WorkspaceConnector is where you create repositories, collaborate on your code,
//
// and organize different streams of work in your Bitbucket Cloud account.
//
// Workspaces replace the use of teams and users in API calls.
type WorkspaceConnector interface {

	// Get returns the requested workspace.
	//
	// GET /2.0/workspaces/{workspace}
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace#get-a-workspace
	Get(ctx context.Context, workspace string) (*models.BitbucketWorkspaceScheme, *models.ResponseScheme, error)

	// Members returns all members of the requested workspace.
	//
	// GET /2.0/workspaces/{workspace}/members
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace#get-a-workspace
	Members(ctx context.Context, workspace string) (*models.WorkspaceMembershipPageScheme, *models.ResponseScheme, error)

	// Membership returns the workspace membership,
	//
	// which includes a User object for the member and a Workspace object for the requested workspace.
	//
	// GET /2.0/workspaces/{workspace}/members/{member}
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace#get-member-in-a-workspace
	Membership(ctx context.Context, workspace, memberID string) (*models.WorkspaceMembershipScheme, *models.ResponseScheme, error)

	// Projects returns the list of projects in this workspace.
	//
	// GET /2.0/workspaces/{workspace}/projects
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace#get-projects-in-a-workspace
	Projects(ctx context.Context, workspace string) (*models.BitbucketProjectPageScheme, *models.ResponseScheme, error)
}

// WorkspaceHookConnector is where you can manage the workspace webhook connector
//
// use it to retrieve, edit and delete webhook on your workspace
type WorkspaceHookConnector interface {

	// Gets returns a paginated list of webhooks installed on this workspace.
	//
	// GET /2.0/workspaces/{workspace}/hooks
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace/webhooks#list-webhooks-for-a-workspace
	Gets(ctx context.Context, workspace string) (*models.WebhookSubscriptionPageScheme, *models.ResponseScheme, error)

	// Create creates a new webhook on the specified workspace.
	//
	// Workspace webhooks are fired for events from all repositories contained by that workspace.
	//
	// POST /2.0/workspaces/{workspace}/hooks
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace/webhooks#create-webhook-for-a-workspace
	Create(ctx context.Context, workspace string, payload *models.WebhookSubscriptionPayloadScheme) (*models.WebhookSubscriptionScheme, *models.ResponseScheme, error)

	// Get returns the webhook with the specified id installed on the given workspace.
	//
	// GET /2.0/workspaces/{workspace}/hooks/{uid}
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace/webhooks#get-webhook-for-a-workspace
	Get(ctx context.Context, workspace, webhookID string) (*models.WebhookSubscriptionScheme, *models.ResponseScheme, error)

	// Update updates the specified webhook subscription.
	//
	// PUT /2.0/workspaces/{workspace}/hooks/{uid}
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace/webhooks#update-webhook-for-a-workspace
	Update(ctx context.Context, workspace, webhookID string, payload *models.WebhookSubscriptionPayloadScheme) (*models.WebhookSubscriptionScheme, *models.ResponseScheme, error)

	// Delete deletes the specified webhook subscription from the given workspace.
	//
	// DELETE /2.0/workspaces/{workspace}/hooks/{uid}
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace/webhooks#delete-webhook-for-a-workspace
	Delete(ctx context.Context, workspace, webhookID string) (*models.ResponseScheme, error)
}

// WorkspacePermissionConnector is where you can manage the workspace permissions
//
// use it to retrieve the workspace permissions and the repositories linked to a workspace.
type WorkspacePermissionConnector interface {

	// Members returns the list of members in a workspace and their permission levels.
	//
	// GET /2.0/workspaces/{workspace}/permissions
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace/permissions#get-user-permissions-in-a-workspace
	Members(ctx context.Context, workspace, query string) (*models.WorkspaceMembershipPageScheme, *models.ResponseScheme, error)

	// Repositories returns an object for each repository permission for all of a workspaces repositories.
	//
	// Permissions returned are effective permissions: the highest level of permission the user has.
	//
	// NOTE: Only users with admin permission for the team may access this resource.
	//
	// GET /2.0/workspaces/{workspace}/permissions/repositories
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace/permissions#gets-all-repository-permissions-in-a-workspace
	Repositories(ctx context.Context, workspace, query, sort string) (*models.RepositoryPermissionPageScheme, *models.ResponseScheme, error)

	// Repository returns an object for the repository permission of each user in the requested repository.
	//
	// Permissions returned are effective permissions: the highest level of permission the user has.
	//
	// GET /2.0/workspaces/{workspace}/permissions/repositories/{repo_slug}
	//
	// https://docs.go-atlassian.io/bitbucket-cloud/workspace/permissions#get-repository-permission-in-a-workspace
	Repository(ctx context.Context, workspace, repository, query, sort string) (*models.RepositoryPermissionPageScheme, *models.ResponseScheme, error)
}
