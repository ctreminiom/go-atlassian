package bitbucket

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type PullRequestConnector interface {
	Gets(ctx context.Context, workspace, repoSlug, state string) (*models.BitbucketPullRequestPageScheme, *models.ResponseScheme, error)
	Get(ctx context.Context, workspace, repoSlug, pullRequestID string) (*models.BitbucketPullRequestScheme, *models.ResponseScheme, error)
	Create(ctx context.Context, workspace, repoSlug string, payload *models.BitbucketPullRequestScheme) (*models.BitbucketPullRequestScheme, *models.ResponseScheme, error)
	Update(ctx context.Context, workspace, repoSlug, pullRequestID string, payload *models.BitbucketPullRequestScheme) (*models.BitbucketPullRequestScheme, *models.ResponseScheme, error)
	Approve(ctx context.Context, workspace, repoSlug, pullRequestID string) (*models.BitbucketParticipantScheme, *models.ResponseScheme, error)
	Unapprove(ctx context.Context, workspace, repoSlug, pullRequestID string) (*models.ResponseScheme, error)
}
