package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type CommentRichTextConnector interface {
	CommentSharedConnector

	// Gets returns all comments for an issue.
	//
	// GET /rest/api/{2-3}/issue/{issueKeyOrID}/comment
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comments
	Gets(ctx context.Context, issueKeyOrID, orderBy string, expand []string, startAt, maxResults int) (*model.IssueCommentPageSchemeV2, *model.ResponseScheme, error)

	// Get returns a comment.
	//
	// GET /rest/api/{2-3}/issue/{issueKeyOrID}/comment/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comment
	Get(ctx context.Context, issueKeyOrID, commentID string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error)

	// Add adds a comment to an issue.
	//
	// POST /rest/api/{2-3}/issue/{issueKeyOrID}/comment
	//
	//https://docs.go-atlassian.io/jira-software-cloud/issues/comments#add-comment
	Add(ctx context.Context, issueKeyOrID string, payload *model.CommentPayloadSchemeV2, expand []string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error)
}

type CommentADFConnector interface {
	CommentSharedConnector

	// Gets returns all comments for an issue.
	//
	// GET /rest/api/{2-3}/issue/{issueKeyOrID}/comment
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comments
	Gets(ctx context.Context, issueKeyOrID, orderBy string, expand []string, startAt, maxResults int) (*model.IssueCommentPageScheme, *model.ResponseScheme, error)

	// Get returns a comment.
	//
	// GET /rest/api/{2-3}/issue/{issueKeyOrID}/comment/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comment
	Get(ctx context.Context, issueKeyOrID, commentID string) (*model.IssueCommentScheme, *model.ResponseScheme, error)

	// Add adds a comment to an issue.
	//
	// POST /rest/api/{2-3}/issue/{issueKeyOrID}/comment
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/comments#add-comment
	Add(ctx context.Context, issueKeyOrID string, payload *model.CommentPayloadScheme, expand []string) (*model.IssueCommentScheme, *model.ResponseScheme, error)
}

type CommentSharedConnector interface {

	// Delete deletes a comment.
	//
	// DELETE /rest/api/{2-3}/issue/{issueKeyOrID}/comment/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/comments#delete-comment
	Delete(ctx context.Context, issueKeyOrID, commentID string) (*model.ResponseScheme, error)
}
