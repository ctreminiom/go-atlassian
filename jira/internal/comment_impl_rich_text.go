package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// CommentRichTextService provides methods to interact with comment operations in Jira Service Management using Rich Text format.
type CommentRichTextService struct {
	// internalClient is the connector interface for Rich Text comment operations.
	internalClient jira.CommentRichTextConnector
}

// Delete deletes a comment.
//
// DELETE /rest/api/{2-3}/issue/{issueKeyOrID}/comment/{commentID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/comments#delete-comment
func (c *CommentRichTextService) Delete(ctx context.Context, issueKeyOrID, commentID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*CommentRichTextService).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete_comment"),
		attribute.String("jira.issue.key", issueKeyOrID),
		attribute.String("jira.comment.id", commentID),
	)

	response, err := c.internalClient.Delete(ctx, issueKeyOrID, commentID)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

// Gets returns all comments for an issue.
//
// GET /rest/api/{2-3}/issue/{issueKeyOrID}/comment
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comments
func (c *CommentRichTextService) Gets(ctx context.Context, issueKeyOrID, orderBy string, expand []string, startAt, maxResults int) (*model.IssueCommentPageSchemeV2, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*CommentRichTextService).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_comments"),
		attribute.String("jira.issue.key", issueKeyOrID),
		attribute.String("jira.order_by", orderBy),
		attribute.StringSlice("jira.expand", expand),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
	)

	result, response, err := c.internalClient.Gets(ctx, issueKeyOrID, orderBy, expand, startAt, maxResults)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Get returns a comment.
//
// GET /rest/api/{2-3}/issue/{issueKeyOrID}/comment/{commentID}
//
// TODO: The documentation needs to be created, raise a ticket here: https://github.com/ctreminiom/go-atlassian/issues
func (c *CommentRichTextService) Get(ctx context.Context, issueKeyOrID, commentID string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*CommentRichTextService).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_comment"),
		attribute.String("jira.issue.key", issueKeyOrID),
		attribute.String("jira.comment.id", commentID),
	)

	result, response, err := c.internalClient.Get(ctx, issueKeyOrID, commentID)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Add adds a comment to an issue.
//
// POST /rest/api/{2-3}/issue/{issueKeyOrID}/comment
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/comments#add-comment
func (c *CommentRichTextService) Add(ctx context.Context, issueKeyOrID string, payload *model.CommentPayloadSchemeV2, expand []string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*CommentRichTextService).Add", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "add_comment"),
		attribute.String("jira.issue.key", issueKeyOrID),
		attribute.StringSlice("jira.expand", expand),
	)

	result, response, err := c.internalClient.Add(ctx, issueKeyOrID, payload, expand)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

type internalRichTextCommentImpl struct {
	c       service.Connector
	version string
}

func (i *internalRichTextCommentImpl) Delete(ctx context.Context, issueKeyOrID, commentID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalRichTextCommentImpl).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete_comment"),
		attribute.String("jira.issue.key", issueKeyOrID),
		attribute.String("jira.comment.id", commentID),
		attribute.String("api.version", i.version),
	)

	if issueKeyOrID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueKeyOrID)
		recordError(span, err)
		return nil, err
	}

	if commentID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoCommentID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/comment/%v", i.version, issueKeyOrID, commentID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

func (i *internalRichTextCommentImpl) Gets(ctx context.Context, issueKeyOrID, orderBy string, expand []string, startAt, maxResults int) (*model.IssueCommentPageSchemeV2, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalRichTextCommentImpl).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_comments"),
		attribute.String("jira.issue.key", issueKeyOrID),
		attribute.String("jira.order_by", orderBy),
		attribute.StringSlice("jira.expand", expand),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
		attribute.String("api.version", i.version),
	)

	if issueKeyOrID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueKeyOrID)
		recordError(span, err)
		return nil, nil, err
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	if orderBy != "" {
		params.Add("orderBy", orderBy)
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/comment?%v", i.version, issueKeyOrID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	comments := new(model.IssueCommentPageSchemeV2)
	response, err := i.c.Call(request, comments)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return comments, response, nil
}

func (i *internalRichTextCommentImpl) Get(ctx context.Context, issueKeyOrID, commentID string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalRichTextCommentImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_comment"),
		attribute.String("jira.issue.key", issueKeyOrID),
		attribute.String("jira.comment.id", commentID),
		attribute.String("api.version", i.version),
	)

	if issueKeyOrID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueKeyOrID)
		recordError(span, err)
		return nil, nil, err
	}

	if commentID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoCommentID)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/comment/%v", i.version, issueKeyOrID, commentID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	comment := new(model.IssueCommentSchemeV2)
	response, err := i.c.Call(request, comment)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return comment, response, nil
}

func (i *internalRichTextCommentImpl) Add(ctx context.Context, issueKeyOrID string, payload *model.CommentPayloadSchemeV2, expand []string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalRichTextCommentImpl).Add", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "add_comment"),
		attribute.String("jira.issue.key", issueKeyOrID),
		attribute.StringSlice("jira.expand", expand),
		attribute.String("api.version", i.version),
	)

	if issueKeyOrID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueKeyOrID)
		recordError(span, err)
		return nil, nil, err
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/issue/%v/comment", i.version, issueKeyOrID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", payload)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	comment := new(model.IssueCommentSchemeV2)
	response, err := i.c.Call(request, comment)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return comment, response, nil
}
