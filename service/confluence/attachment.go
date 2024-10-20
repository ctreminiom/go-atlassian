package confluence

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
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

type AttachmentConnector interface {

	// Get returns a specific attachment.
	//
	// GET /wiki/api/v2/attachments/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/attachments#get-attachment-by-id
	Get(ctx context.Context, attachmentID string, versionID int, serializeIDs bool) (*model.AttachmentScheme, *model.ResponseScheme, error)

	// Gets returns the attachments of specific entity type.
	//
	// You can extract the attachments for blog-posts,custom-contents, labels and pages.
	//
	// Depending on the entity type, the endpoint will change based on the entity type.
	//
	// Valid entityType values: blogposts, custom-content, labels, pages.
	//
	// The number of results is limited by the limit parameter and additional results.
	//
	// (if available) will be available through the next URL present in the Link response header.
	//
	// GET /wiki/api/v2/{blogposts,custom-content,labels,pages}/{id}/attachments
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/attachments#get-attachments-by-type
	Gets(ctx context.Context, entityID int, entityType string, options *model.AttachmentParamsScheme, cursor string, limit int) (*model.AttachmentPageScheme, *model.ResponseScheme, error)

	// Delete deletes an attachment by id.
	//
	// DELETE /wiki/api/v2/attachments/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/attachments#delete-attachment
	Delete(ctx context.Context, attachmentID string) (*model.ResponseScheme, error)
}

type AttachmentVersionConnector interface {

	// Gets returns the versions of specific attachment.
	//
	// GET /wiki/api/v2/attachments/{id}/versions
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/attachments/versions#get-attachment-versions
	Gets(ctx context.Context, attachmentID, cursor, sort string, limit int) (*model.AttachmentVersionPageScheme, *model.ResponseScheme, error)

	// Get retrieves version details for the specified attachment and version number.
	//
	// GET /wiki/api/v2/attachments/{attachment-id}/versions/{version-number}
	//
	// https://docs.go-atlassian.io/confluence-cloud/v2/attachments/versions#get-attachment-version
	Get(ctx context.Context, attachmentID string, versionID int) (*model.DetailedVersionScheme, *model.ResponseScheme, error)
}
