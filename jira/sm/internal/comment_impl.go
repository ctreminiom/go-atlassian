package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/google/go-querystring/query"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/sm"
)

// NewCommentService creates a new instance of CommentService.
// It takes a service.Connector and a version string as input and returns a pointer to CommentService.
func NewCommentService(client service.Connector, version string) *CommentService {
	return &CommentService{
		internalClient: &internalServiceRequestCommentImpl{c: client, version: version},
	}
}

// CommentService provides methods to interact with comment operations in Jira Service Management.
type CommentService struct {
	// internalClient is the connector interface for comment operations.
	internalClient sm.CommentConnector
}

// Gets returns all comments on a customer request.
//
// No permissions error is provided if, for example, the user doesn't have access to the service desk or request,
//
// the method simply returns an empty response.
//
// GET /rest/servicedeskapi/request/{issueKeyOrID}/comment
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#get-request-comments
func (s *CommentService) Gets(ctx context.Context, issueKeyOrID string, options *model.RequestCommentOptionsScheme) (*model.RequestCommentPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*CommentService).Gets")
	defer span.End()

	return s.internalClient.Gets(ctx, issueKeyOrID, options)
}

// Get returns details of a customer request's comment.
//
// GET /rest/servicedeskapi/request/{issueKeyOrID}/comment/{commentID}
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#get-request-comment-by-id
func (s *CommentService) Get(ctx context.Context, issueKeyOrID string, commentID int, expand []string) (*model.RequestCommentScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*CommentService).Get")
	defer span.End()

	return s.internalClient.Get(ctx, issueKeyOrID, commentID, expand)
}

// Create creates a public or private (internal) comment on a customer request, with the comment visibility set by public.
//
// POST /rest/servicedeskapi/request/{issueKeyOrID}/comment
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#create-request-comment
func (s *CommentService) Create(ctx context.Context, issueKeyOrID, body string, public bool) (*model.RequestCommentScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*CommentService).Create")
	defer span.End()

	return s.internalClient.Create(ctx, issueKeyOrID, body, public)
}

// Attachments  returns the attachments referenced in a comment.
//
// GET /rest/servicedeskapi/request/{issueKeyOrID}/comment/{commentID}/attachment
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#get-comment-attachments
func (s *CommentService) Attachments(ctx context.Context, issueKeyOrID string, commentID, start, limit int) (*model.RequestAttachmentPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*CommentService).Attachments")
	defer span.End()

	return s.internalClient.Attachments(ctx, issueKeyOrID, commentID, start, limit)
}

type internalServiceRequestCommentImpl struct {
	c       service.Connector
	version string
}

func (i *internalServiceRequestCommentImpl) Gets(ctx context.Context, issueKeyOrID string, options *model.RequestCommentOptionsScheme) (*model.RequestCommentPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalServiceRequestCommentImpl).Gets")
	defer span.End()

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/comment", issueKeyOrID)

	if options != nil {
		q, err := query.Values(options)
		if err != nil {
			return nil, nil, err
		}

		endpoint += "?" + q.Encode()
	}

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
	ctx, span := tracer().Start(ctx, "(*internalServiceRequestCommentImpl).Get")
	defer span.End()

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	if commentID == 0 {
		return nil, nil, model.ErrNoCommentID
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
	ctx, span := tracer().Start(ctx, "(*internalServiceRequestCommentImpl).Create")
	defer span.End()

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	if body == "" {
		return nil, nil, model.ErrNoCommentBody
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
	ctx, span := tracer().Start(ctx, "(*internalServiceRequestCommentImpl).Attachments")
	defer span.End()

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	if commentID == 0 {
		return nil, nil, model.ErrNoContentID
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
