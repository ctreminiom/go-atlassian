package confluence

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type CommentConnector interface {

	// Gets returns the comments on a piece of content.
	//
	// GET /wiki/rest/api/content/{id}/child/comment
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/comments#get-content-comments
	Gets(ctx context.Context, contentID string, expand, location []string, startAt, maxResults int) (*model.ContentPageScheme, *model.ResponseScheme, error)
}
