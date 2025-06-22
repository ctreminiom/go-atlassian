package internal

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewIssueAttachmentService creates a new instance of IssueAttachmentService.
// It takes a service.Connector and a version string as input.
// Returns a pointer to IssueAttachmentService and an error if the version is not provided.
func NewIssueAttachmentService(client service.Connector, version string) (*IssueAttachmentService, error) {

	if version == "" {
		return nil, fmt.Errorf("jira: %w", model.ErrNoVersionProvided)
	}

	return &IssueAttachmentService{
		internalClient: &internalIssueAttachmentServiceImpl{c: client, version: version},
	}, nil
}

// IssueAttachmentService provides methods to interact with issue attachment operations in Jira Service Management.
type IssueAttachmentService struct {
	// internalClient is the connector interface for issue attachment operations.
	internalClient jira.AttachmentConnector
}

// Settings returns the attachment settings, that is, whether attachments are enabled and the maximum attachment size allowed.
//
// GET /rest/api/{2-3}/attachment/meta
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#get-jira-attachment-settings
func (i *IssueAttachmentService) Settings(ctx context.Context) (*model.AttachmentSettingScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueAttachmentService).Settings", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_attachment_settings"),
	)

	result, response, err := i.internalClient.Settings(ctx)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Metadata returns the metadata for an attachment. Note that the attachment itself is not returned.
//
// GET /rest/api/{2-3}/attachment/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#get-attachment-metadata
func (i *IssueAttachmentService) Metadata(ctx context.Context, attachmentID string) (*model.IssueAttachmentMetadataScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueAttachmentService).Metadata", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_attachment_metadata"),
		attribute.String("jira.attachment.id", attachmentID),
	)

	result, response, err := i.internalClient.Metadata(ctx, attachmentID)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Delete deletes an attachment from an issue.
//
// DELETE /rest/api/{2-3}/attachment/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#delete-attachment
func (i *IssueAttachmentService) Delete(ctx context.Context, attachmentID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueAttachmentService).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete_attachment"),
		attribute.String("jira.attachment.id", attachmentID),
	)

	response, err := i.internalClient.Delete(ctx, attachmentID)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

// Human returns the metadata for the contents of an attachment, if it is an archive, and metadata for the attachment itself.
//
// For example, if the attachment is a ZIP archive, then information about the files in the archive is returned and metadata for the ZIP archive.
//
// GET /rest/api/{2-3}/attachment/{id}/expand/human
//
// # Experimental Endpoint
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#get-all-metadata-for-an-expanded-attachment
func (i *IssueAttachmentService) Human(ctx context.Context, attachmentID string) (*model.IssueAttachmentHumanMetadataScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueAttachmentService).Human", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_attachment_human_metadata"),
		attribute.String("jira.attachment.id", attachmentID),
	)

	result, response, err := i.internalClient.Human(ctx, attachmentID)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Add adds one attachment to an issue. Attachments are posted as multipart/form-data (RFC 1867).
//
// POST /rest/api/{2-3}/issue/{issueKeyOrID}/attachments
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#add-attachment
func (i *IssueAttachmentService) Add(ctx context.Context, issueKeyOrID, fileName string, file io.Reader) ([]*model.IssueAttachmentScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueAttachmentService).Add", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "add_attachment"),
		attribute.String("jira.issue.key", issueKeyOrID),
		attribute.String("jira.attachment.filename", fileName),
	)

	result, response, err := i.internalClient.Add(ctx, issueKeyOrID, fileName, file)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Download returns the contents of an attachment. A Range header can be set to define a range of bytes within the attachment to download.
//
// See the HTTP Range header standard for details.
//
// GET /rest/api/{2-3}/attachment/content/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#download-attachment
func (i *IssueAttachmentService) Download(ctx context.Context, attachmentID string, redirect bool) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueAttachmentService).Download", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "download_attachment"),
		attribute.String("jira.attachment.id", attachmentID),
		attribute.Bool("jira.attachment.redirect", redirect),
	)

	response, err := i.internalClient.Download(ctx, attachmentID, redirect)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

type internalIssueAttachmentServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalIssueAttachmentServiceImpl) Download(ctx context.Context, attachmentID string, redirect bool) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueAttachmentServiceImpl).Download", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "download_attachment"),
		attribute.String("jira.attachment.id", attachmentID),
		attribute.Bool("jira.attachment.redirect", redirect),
		attribute.String("api.version", i.version),
	)

	if attachmentID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoAttachmentID)
		recordError(span, err)
		return nil, err
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/attachment/content/%v", i.version, attachmentID))

	if !redirect {

		params := url.Values{}
		params.Add("redirect", "false") //default: true

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

func (i *internalIssueAttachmentServiceImpl) Settings(ctx context.Context) (*model.AttachmentSettingScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueAttachmentServiceImpl).Settings", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_attachment_settings"),
		attribute.String("api.version", i.version),
	)

	endpoint := fmt.Sprintf("rest/api/%v/attachment/meta", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	settings := new(model.AttachmentSettingScheme)
	response, err := i.c.Call(request, settings)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return settings, response, nil
}

func (i *internalIssueAttachmentServiceImpl) Metadata(ctx context.Context, attachmentID string) (*model.IssueAttachmentMetadataScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueAttachmentServiceImpl).Metadata", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_attachment_metadata"),
		attribute.String("jira.attachment.id", attachmentID),
		attribute.String("api.version", i.version),
	)

	if attachmentID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoAttachmentID)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/attachment/%v", i.version, attachmentID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	metadata := new(model.IssueAttachmentMetadataScheme)
	response, err := i.c.Call(request, metadata)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return metadata, response, nil
}

func (i *internalIssueAttachmentServiceImpl) Delete(ctx context.Context, attachmentID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueAttachmentServiceImpl).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete_attachment"),
		attribute.String("jira.attachment.id", attachmentID),
		attribute.String("api.version", i.version),
	)

	if attachmentID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoAttachmentID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/attachment/%v", i.version, attachmentID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

func (i *internalIssueAttachmentServiceImpl) Human(ctx context.Context, attachmentID string) (*model.IssueAttachmentHumanMetadataScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueAttachmentServiceImpl).Human", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_attachment_human_metadata"),
		attribute.String("jira.attachment.id", attachmentID),
		attribute.String("api.version", i.version),
	)

	if attachmentID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoAttachmentID)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/attachment/%v/expand/human", i.version, attachmentID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	metadata := new(model.IssueAttachmentHumanMetadataScheme)
	response, err := i.c.Call(request, metadata)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return metadata, response, nil
}

func (i *internalIssueAttachmentServiceImpl) Add(ctx context.Context, issueKeyOrID, fileName string, file io.Reader) ([]*model.IssueAttachmentScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueAttachmentServiceImpl).Add", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "add_attachment"),
		attribute.String("jira.issue.key", issueKeyOrID),
		attribute.String("jira.attachment.filename", fileName),
		attribute.String("api.version", i.version),
	)

	if issueKeyOrID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueKeyOrID)
		recordError(span, err)
		return nil, nil, err
	}

	if fileName == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoAttachmentName)
		recordError(span, err)
		return nil, nil, err
	}

	if file == nil {
		err := fmt.Errorf("jira: %w", model.ErrNoReader)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/attachments", i.version, issueKeyOrID)

	reader := &bytes.Buffer{}
	writer := multipart.NewWriter(reader)

	attachment, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	_, err = io.Copy(attachment, file)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	writer.Close()

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, writer.FormDataContentType(), reader)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	var attachments []*model.IssueAttachmentScheme
	response, err := i.c.Call(request, &attachments)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return attachments, response, nil
}
