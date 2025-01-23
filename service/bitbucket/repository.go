package bitbucket

import (
	"context"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// RepositoryConnector represents the Bitbucket Cloud repositories.
//
// A Git repository is a virtual storage of your project. It allows you to save versions of your code, which you can access when needed.
//
// The repo resource allows you to access public repos, or repos that belong to a specific workspace.
type RepositoryConnector interface {

	// Get returns the object describing this repository.
	//
	// GET /2.0/repositories/{workspace}/{repo_slug}
	Get(ctx context.Context, workspace, repoSlug string) (*models.RepositoryScheme, *models.ResponseScheme, error)

	// Update updates a Bitbucket repository
	//
	// PUT /2.0/repositories/{workspace}/{repo_slug}
	//
	// Note: Changing the name of the repository will cause the location to be changed.
	//
	// This is because the URL of the repo is derived from the name (a process called slugification).
	//
	// In such a scenario, it is possible for the request to fail if the newly created slug conflicts with an existing repository's slug.
	Update(ctx context.Context, workspace, repoSlug string, payload *models.RepositoryScheme) (*models.RepositoryScheme, *models.ResponseScheme, error)

	// Delete deletes the repository. This is an irreversible operation and this does not affect its forks.
	//
	// DELETE /2.0/repositories/{workspace}/{repo_slug}
	Delete(ctx context.Context, workspace, repoSlug, redirectTo string) (*models.ResponseScheme, error)

	// Create creates a new repository.
	Create(ctx context.Context, workspace, repoSlug string, payload *models.RepositoryScheme) (*models.RepositoryScheme, *models.ResponseScheme, error)

	// Watchers returns a paginated list of all the watchers on the specified repository.
	//
	// GET /2.0/repositories/{workspace}/{repo_slug}/watchers
	Watchers(ctx context.Context, workspace, repoSlug string)
}

// RepositoryForkConnector represents the Bitbucket Cloud repository forks.
type RepositoryForkConnector interface {
	Gets(ctx context.Context)
	Execute(ctx context.Context)
}

// RepositoryWebhookConnector represents the Bitbucket Cloud repository webhooks.
type RepositoryWebhookConnector interface {
	Gets(ctx context.Context)
	Get(ctx context.Context)
	Create(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
}

// RepositorySettingConnector represents the Bitbucket Cloud repository settings.
type RepositorySettingConnector interface {
	Gets(ctx context.Context)
	Set(ctx context.Context)
}

// RepositoryGroupPermissionConnector represents the Bitbucket Cloud repository group permissions.
type RepositoryGroupPermissionConnector interface {
	Gets(ctx context.Context)
	Get(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
}

// RepositoryUserPermissionConnector represents the Bitbucket Cloud repository user permissions.
type RepositoryUserPermissionConnector interface {
	Gets(ctx context.Context)
	Get(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
	Check(ctx context.Context)
}
