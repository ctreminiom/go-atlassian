package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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
	return c.internalClient.Delete(ctx, issueKeyOrID, commentID)
}

// Gets returns all comments for an issue.
//
// GET /rest/api/{2-3}/issue/{issueKeyOrID}/comment
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comments
func (c *CommentRichTextService) Gets(ctx context.Context, issueKeyOrID, orderBy string, expand []string, startAt, maxResults int) (*model.IssueCommentPageSchemeV2, *model.ResponseScheme, error) {
	return c.internalClient.Gets(ctx, issueKeyOrID, orderBy, expand, startAt, maxResults)
}

// Get returns a comment.
//
// GET /rest/api/{2-3}/issue/{issueKeyOrID}/comment/{commentID}
//
// TODO: The documentation needs to be created, raise a ticket here: https://github.com/ctreminiom/go-atlassian/issues
func (c *CommentRichTextService) Get(ctx context.Context, issueKeyOrID, commentID string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error) {
	return c.internalClient.Get(ctx, issueKeyOrID, commentID)
}

// Add adds a comment to an issue.
//
// POST /rest/api/{2-3}/issue/{issueKeyOrID}/comment
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/comments#add-comment
func (c *CommentRichTextService) Add(ctx context.Context, issueKeyOrID string, payload *model.CommentPayloadSchemeV2, expand []string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error) {
	return c.internalClient.Add(ctx, issueKeyOrID, payload, expand)
}

type internalRichTextCommentImpl struct {
	c       service.Connector
	version string
}

func (i *internalRichTextCommentImpl) Delete(ctx context.Context, issueKeyOrID, commentID string) (*model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, model.ErrNoIssueKeyOrID
	}

	if commentID == "" {
		return nil, model.ErrNoCommentID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/comment/%v", i.version, issueKeyOrID, commentID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalRichTextCommentImpl) Gets(ctx context.Context, issueKeyOrID, orderBy string, expand []string, startAt, maxResults int) (*model.IssueCommentPageSchemeV2, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
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
		return nil, nil, err
	}

	comments := new(model.IssueCommentPageSchemeV2)
	response, err := i.c.Call(request, comments)
	if err != nil {
		return nil, response, err
	}

	return comments, response, nil
}

func (i *internalRichTextCommentImpl) Get(ctx context.Context, issueKeyOrID, commentID string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	if commentID == "" {
		return nil, nil, model.ErrNoCommentID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/comment/%v", i.version, issueKeyOrID, commentID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	comment := new(model.IssueCommentSchemeV2)
	response, err := i.c.Call(request, comment)
	if err != nil {
		return nil, response, err
	}

	return comment, response, nil
}

func (i *internalRichTextCommentImpl) Add(ctx context.Context, issueKeyOrID string, payload *model.CommentPayloadSchemeV2, expand []string) (*model.IssueCommentSchemeV2, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
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
		return nil, nil, err
	}

	comment := new(model.IssueCommentSchemeV2)
	response, err := i.c.Call(request, comment)
	if err != nil {
		return nil, response, err
	}

	return comment, response, nil
}
