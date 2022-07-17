package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type RichTextComment interface {
	SharedComment

	// Gets returns all comments for an issue.
	// GET /rest/api/2/issue/{issueIdOrKey}/comment
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comments
	Gets(ctx context.Context, issueKeyOrId, orderBy string, expand []string, startAt, maxResults int) (*model.IssueCommentPageSchemeV2, *model.ResponseScheme, error)

	// Get returns a comment.
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comment
	// GET /rest/api/2/issue/{issueIdOrKey}/comment/{id}
	// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-comments/#api-rest-api-2-comment-list-post
	Get(ctx context.Context, issueKeyOrId, commentId string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error)

	// Add adds a comment to an issue.
	// POST /rest/api/2/issue/{issueIdOrKey}/comment
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#add-comment
	Add(ctx context.Context, issueKeyOrId string, payload *model.CommentPayloadSchemeV2, expand []string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error)
}

type AdfComment interface {
	SharedComment

	// Gets returns all comments for an issue.
	// GET /rest/api/3/issue/{issueIdOrKey}/comment
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comments
	Gets(ctx context.Context, issueKeyOrId, orderBy string, expand []string, startAt, maxResults int) (*model.IssueCommentPageScheme, *model.ResponseScheme, error)

	// Get returns a comment.
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comment
	// GET /rest/api/3/issue/{issueIdOrKey}/comment/{id}
	// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-comments/#api-rest-api-2-comment-list-post
	Get(ctx context.Context, issueKeyOrId, commentId string) (*model.IssueCommentScheme, *model.ResponseScheme, error)

	// Add adds a comment to an issue.
	// POST /rest/api/3/issue/{issueIdOrKey}/comment
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#add-comment
	Add(ctx context.Context, issueKeyOrId string, payload *model.CommentPayloadScheme, expand []string) (*model.IssueCommentScheme, *model.ResponseScheme, error)
}

type SharedComment interface {

	// Delete deletes a comment.
	// DELETE /rest/api/{2-3}/issue/{issueIdOrKey}/comment/{id}
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#delete-comment
	Delete(ctx context.Context, issueKeyOrId, commentId string) (*model.ResponseScheme, error)
}
