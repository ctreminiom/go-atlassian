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

func NewCommentService(client service.Connector, version string) *CommentService {

	return &CommentService{
		internalClient: &internalServiceRequestCommentImpl{c: client, version: version},
	}
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
	c       service.Connector
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

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.RequestCommentPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
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

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	comment := new(model.RequestCommentScheme)
	res, err := i.c.Call(req, comment)
	if err != nil {
		return nil, res, err
	}

	return comment, res, nil
}

func (i *internalServiceRequestCommentImpl) Create(ctx context.Context, issueKeyOrID, body string, public bool) (*model.RequestCommentScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if body == "" {
		return nil, nil, model.ErrNoCommentBodyError
	}

	payload := map[string]interface{}{"public": public, "body": body}
	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/comment", issueKeyOrID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	comment := new(model.RequestCommentScheme)
	res, err := i.c.Call(req, comment)
	if err != nil {
		return nil, res, err
	}

	return comment, res, nil
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

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.RequestAttachmentPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}
