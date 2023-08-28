package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"io"
)

type AttachmentConnector interface {

	// Settings returns the attachment settings, that is, whether attachments are enabled and the maximum attachment size allowed.
	//
	// GET /rest/api/{2-3}/attachment/meta
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#get-jira-attachment-settings
	Settings(ctx context.Context) (*model.AttachmentSettingScheme, *model.ResponseScheme, error)

	// Metadata returns the metadata for an attachment. Note that the attachment itself is not returned.
	//
	// GET /rest/api/{2-3}/attachment/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#get-attachment-metadata
	Metadata(ctx context.Context, attachmentId string) (*model.IssueAttachmentMetadataScheme, *model.ResponseScheme, error)

	// Delete deletes an attachment from an issue.
	//
	// DELETE /rest/api/{2-3}/attachment/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#delete-attachment
	Delete(ctx context.Context, attachmentId string) (*model.ResponseScheme, error)

	// Human returns the metadata for the contents of an attachment, if it is an archive, and metadata for the attachment itself.
	//
	// For example, if the attachment is a ZIP archive, then information about the files in the archive is returned and metadata for the ZIP archive.
	//
	// GET /rest/api/{2-3}/attachment/{id}/expand/human
	//
	// Experimental Endpoint
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#get-all-metadata-for-an-expanded-attachment
	Human(ctx context.Context, attachmentId string) (*model.IssueAttachmentHumanMetadataScheme, *model.ResponseScheme, error)

	// Add adds one attachment to an issue. Attachments are posted as multipart/form-data (RFC 1867).
	//
	// POST /rest/api/{2-3}/issue/{issueIdOrKey}/attachments
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#add-attachment
	Add(ctx context.Context, issueKeyOrId, fileName string, file io.Reader) ([]*model.IssueAttachmentScheme, *model.ResponseScheme, error)

	// Download returns the contents of an attachment. A Range header can be set to define a range of bytes within the attachment to download.
	//
	// See the HTTP Range header standard for details.
	//
	// GET /rest/api/{2-3}/attachment/content/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#download-attachment
	Download(ctx context.Context, attachmentID string, redirect bool) (*model.ResponseScheme, error)
}
