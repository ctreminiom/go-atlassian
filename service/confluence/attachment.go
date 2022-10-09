package confluence

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"io"
)

type ContentAttachmentConnector interface {

	// Gets returns the attachments for a piece of content.
	//
	// By default, the following objects are expanded: metadata.
	//
	// GET /wiki/rest/api/content/{id}/child/attachment
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/attachments#get-attachments
	Gets(ctx context.Context, contentID string, startAt, maxResults int, options *model.GetContentAttachmentsOptionsScheme) (*model.ContentPageScheme, *model.ResponseScheme, error)

	// CreateOrUpdate adds an attachment to a piece of content.
	//
	// If the attachment already exists for the content,
	//
	// then the attachment is updated (i.e. a new version of the attachment is created).
	//
	// PUT /wiki/rest/api/content/{id}/child/attachment
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/attachments#create-or-update-attachment
	CreateOrUpdate(ctx context.Context, attachmentID, status, fileName string, file io.Reader) (*model.ContentPageScheme, *model.ResponseScheme, error)

	// Create adds an attachment to a piece of content.
	//
	// This method only adds a new attachment.
	//
	// If you want to update an existing attachment, use Create or update attachments.
	//
	// POST /wiki/rest/api/content/{id}/child/attachment
	//
	// https://docs.go-atlassian.io/confluence-cloud/content/attachments#create-attachment
	Create(ctx context.Context, attachmentID, status, fileName string, file io.Reader) (*model.ContentPageScheme, *model.ResponseScheme, error)
}
