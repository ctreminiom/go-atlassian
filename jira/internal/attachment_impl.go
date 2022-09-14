package internal

import (
	"bytes"
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"io"
	"mime/multipart"
	"net/http"
)

func NewIssueAttachmentService(client service.Client, version string) (*IssueAttachmentService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &IssueAttachmentService{
		internalClient: &internalIssueAttachmentServiceImpl{c: client, version: version},
	}, nil
}

type IssueAttachmentService struct {
	internalClient jira.AttachmentConnector
}

// Settings returns the attachment settings, that is, whether attachments are enabled and the maximum attachment size allowed.
//
// GET /rest/api/{2-3}/attachment/meta
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#get-jira-attachment-settings
func (i *IssueAttachmentService) Settings(ctx context.Context) (*model.AttachmentSettingScheme, *model.ResponseScheme, error) {
	return i.internalClient.Settings(ctx)
}

// Metadata returns the metadata for an attachment. Note that the attachment itself is not returned.
//
// GET /rest/api/{2-3}/attachment/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#get-attachment-metadata
func (i *IssueAttachmentService) Metadata(ctx context.Context, attachmentId string) (*model.AttachmentMetadataScheme, *model.ResponseScheme, error) {
	return i.internalClient.Metadata(ctx, attachmentId)
}

// Delete deletes an attachment from an issue.
//
// DELETE /rest/api/{2-3}/attachment/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#delete-attachment
func (i *IssueAttachmentService) Delete(ctx context.Context, attachmentId string) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, attachmentId)
}

// Human returns the metadata for the contents of an attachment, if it is an archive, and metadata for the attachment itself.
//
// For example, if the attachment is a ZIP archive, then information about the files in the archive is returned and metadata for the ZIP archive.
//
// GET /rest/api/{2-3}/attachment/{id}/expand/human
//
// Experimental Endpoint
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#get-all-metadata-for-an-expanded-attachment
func (i *IssueAttachmentService) Human(ctx context.Context, attachmentId string) (*model.AttachmentHumanMetadataScheme, *model.ResponseScheme, error) {
	return i.internalClient.Human(ctx, attachmentId)
}

// Add adds one attachment to an issue. Attachments are posted as multipart/form-data (RFC 1867).
//
// POST /rest/api/{2-3}/issue/{issueIdOrKey}/attachments
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/attachments#add-attachment
func (i *IssueAttachmentService) Add(ctx context.Context, issueKeyOrId, fileName string, file io.Reader) ([]*model.AttachmentScheme, *model.ResponseScheme, error) {
	return i.internalClient.Add(ctx, issueKeyOrId, fileName, file)
}

type internalIssueAttachmentServiceImpl struct {
	c       service.Client
	version string
}

func (i *internalIssueAttachmentServiceImpl) Settings(ctx context.Context) (*model.AttachmentSettingScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/attachment/meta", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	settings := new(model.AttachmentSettingScheme)
	response, err := i.c.Call(request, settings)
	if err != nil {
		return nil, response, err
	}

	return settings, response, nil
}

func (i *internalIssueAttachmentServiceImpl) Metadata(ctx context.Context, attachmentId string) (*model.AttachmentMetadataScheme, *model.ResponseScheme, error) {

	if attachmentId == "" {
		return nil, nil, model.ErrNoAttachmentIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/attachment/%v", i.version, attachmentId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	metadata := new(model.AttachmentMetadataScheme)
	response, err := i.c.Call(request, metadata)
	if err != nil {
		return nil, response, err
	}

	return metadata, response, nil
}

func (i *internalIssueAttachmentServiceImpl) Delete(ctx context.Context, attachmentId string) (*model.ResponseScheme, error) {

	if attachmentId == "" {
		return nil, model.ErrNoAttachmentIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/attachment/%v", i.version, attachmentId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueAttachmentServiceImpl) Human(ctx context.Context, attachmentId string) (*model.AttachmentHumanMetadataScheme, *model.ResponseScheme, error) {

	if attachmentId == "" {
		return nil, nil, model.ErrNoAttachmentIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/attachment/%v/expand/human", i.version, attachmentId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	metadata := new(model.AttachmentHumanMetadataScheme)
	response, err := i.c.Call(request, metadata)
	if err != nil {
		return nil, response, err
	}

	return metadata, response, nil
}

func (i *internalIssueAttachmentServiceImpl) Add(ctx context.Context, issueKeyOrId, fileName string, file io.Reader) ([]*model.AttachmentScheme, *model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if fileName == "" {
		return nil, nil, model.ErrNoAttachmentNameError
	}

	if file == nil {
		return nil, nil, model.ErrNoReaderError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/attachments", i.version, issueKeyOrId)

	reader := &bytes.Buffer{}
	writer := multipart.NewWriter(reader)

	attachment, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, nil, err
	}

	_, err = io.Copy(attachment, file)
	if err != nil {
		return nil, nil, err
	}

	writer.Close()

	request, err := i.c.NewFormRequest(ctx, http.MethodPost, endpoint, writer.FormDataContentType(), reader)
	if err != nil {
		return nil, nil, err
	}

	var attachments []*model.AttachmentScheme
	response, err := i.c.Call(request, attachments)
	if err != nil {
		return nil, response, err
	}

	return attachments, response, nil
}
