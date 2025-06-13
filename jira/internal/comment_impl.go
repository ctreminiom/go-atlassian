package internal

import (
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
)

// NewCommentService creates a new instance of CommentADFService and CommentRichTextService.
// It takes a service.Connector and a version string as input.
// Returns pointers to CommentADFService and CommentRichTextService, and an error if the version is not provided.
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
