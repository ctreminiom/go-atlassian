package internal

import (
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
)

func NewCommentService(client service.Connector, version string) (*CommentADFService, *CommentRichTextService, error) {

	if version == "" {
		return nil, nil, model.ErrNoVersionProvided
	}

	adfService := &CommentADFService{
		internalClient: &internalAdfCommentImpl{
			c:       client,
			version: version,
		},
	}

	richTextService := &CommentRichTextService{
		internalClient: &internalRichTextCommentImpl{
			c:       client,
			version: version,
		},
	}

	return adfService, richTextService, nil
}
