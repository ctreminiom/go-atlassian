package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/confluence"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func NewCommentService(client service.Client) *CommentService {

	return &CommentService{
		internalClient: &internalCommentImpl{c: client},
	}
}

type CommentService struct {
	internalClient confluence.CommentConnector
}

// Gets returns the comments on a piece of content.
//
// GET /wiki/rest/api/content/{id}/child/comment
//
// https://docs.go-atlassian.io/confluence-cloud/content/comments#get-content-comments
func (c *CommentService) Gets(ctx context.Context, contentID string, expand, location []string, startAt, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error) {
	return c.internalClient.Gets(ctx, contentID, expand, location, startAt, maxResults)
}

type internalCommentImpl struct {
	c service.Client
}

func (i *internalCommentImpl) Gets(ctx context.Context, contentID string, expand, location []string, startAt, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentIDError
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if len(expand) != 0 {
		query.Add("expand", strings.Join(expand, ","))
	}

	if len(location) != 0 {
		query.Add("location", strings.Join(location, ","))
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/child/comment?%v", contentID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ContentPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}
