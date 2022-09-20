package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/sm"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func NewCommentService(client service.Client, version string) (*CommentService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &CommentService{
		internalClient: &internalServiceRequestCommentImpl{c: client, version: version},
	}, nil
}

type CommentService struct {
	internalClient sm.CommentConnector
}

// Gets returns all comments on a customer request.
//
// No permissions error is provided if, for example, the user doesn't have access to the service desk or request,
//
// the method simply returns an empty response.
//
// GET /rest/servicedeskapi/request/{issueIdOrKey}/comment
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#get-request-comments
func (s *CommentService) Gets(ctx context.Context, issueKeyOrID string, public bool, expand []string, start, limit int) (*model.RequestCommentPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Gets(ctx, issueKeyOrID, public, expand, start, limit)
}

// Get returns details of a customer request's comment.
//
// GET /rest/servicedeskapi/request/{issueIdOrKey}/comment/{commentId}
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#get-request-comment-by-id
func (s *CommentService) Get(ctx context.Context, issueKeyOrID string, commentID int, expand []string) (*model.RequestCommentScheme, *model.ResponseScheme, error) {
	return s.internalClient.Get(ctx, issueKeyOrID, commentID, expand)
}

// Create creates a public or private (internal) comment on a customer request, with the comment visibility set by public.
//
// POST /rest/servicedeskapi/request/{issueIdOrKey}/comment
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#create-request-comment
func (s *CommentService) Create(ctx context.Context, issueKeyOrID, body string, public bool) (*model.RequestCommentScheme, *model.ResponseScheme, error) {
	return s.internalClient.Create(ctx, issueKeyOrID, body, public)
}

// Attachments  returns the attachments referenced in a comment.
//
// GET /rest/servicedeskapi/request/{issueIdOrKey}/comment/{commentId}/attachment
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#get-comment-attachments
func (s *CommentService) Attachments(ctx context.Context, issueKeyOrID string, commentID, start, limit int) (*model.RequestAttachmentPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Attachments(ctx, issueKeyOrID, commentID, start, limit)
}

type internalServiceRequestCommentImpl struct {
	c       service.Client
	version string
}

func (i *internalServiceRequestCommentImpl) Gets(ctx context.Context, issueKeyOrID string, public bool, expand []string, start, limit int) (*model.RequestCommentPageScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if !public {
		params.Add("public", "false")
	}

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/comment?%v", issueKeyOrID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.RequestCommentPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalServiceRequestCommentImpl) Get(ctx context.Context, issueKeyOrID string, commentID int, expand []string) (*model.RequestCommentScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if commentID == 0 {
		return nil, nil, model.ErrNoCommentIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/servicedeskapi/request/%v/comment/%v", issueKeyOrID, commentID))

	if expand != nil {
		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	comment := new(model.RequestCommentScheme)
	response, err := i.c.Call(request, comment)
	if err != nil {
		return nil, response, err
	}

	return comment, response, nil
}

func (i *internalServiceRequestCommentImpl) Create(ctx context.Context, issueKeyOrID, body string, public bool) (*model.RequestCommentScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if body == "" {
		return nil, nil, model.ErrNoCommentBodyError
	}

	payload := struct {
		Public bool   `json:"public"`
		Body   string `json:"body"`
	}{
		Public: public,
		Body:   body,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/comment", issueKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	comment := new(model.RequestCommentScheme)
	response, err := i.c.Call(request, comment)
	if err != nil {
		return nil, response, err
	}

	return comment, response, nil
}

func (i *internalServiceRequestCommentImpl) Attachments(ctx context.Context, issueKeyOrID string, commentID, start, limit int) (*model.RequestAttachmentPageScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if commentID == 0 {
		return nil, nil, model.ErrNoContentIDError
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/comment/%v/attachment?%v", issueKeyOrID, commentID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.RequestAttachmentPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}
