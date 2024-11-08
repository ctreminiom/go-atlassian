package internal

import (
	"bytes"
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// NewContentAttachmentService creates a new instance of ContentAttachmentService.
// It takes a service.Connector as input and returns a pointer to ContentAttachmentService.
func NewContentAttachmentService(client service.Connector) *ContentAttachmentService {

	return &ContentAttachmentService{
		internalClient: &internalContentAttachmentImpl{c: client},
	}
}

// ContentAttachmentService provides methods to interact with content attachment operations in Confluence.
type ContentAttachmentService struct {
	// internalClient is the connector interface for content attachment operations.
	internalClient confluence.ContentAttachmentConnector
}

// Gets returns the attachments for a piece of content.
//
// By default, the following objects are expanded: metadata.
//
// GET /wiki/rest/api/content/{id}/child/attachment
//
// https://docs.go-atlassian.io/confluence-cloud/content/attachments#get-attachments
func (a *ContentAttachmentService) Gets(ctx context.Context, contentID string, startAt, maxResults int, options *model.GetContentAttachmentsOptionsScheme) (*model.ContentPageScheme, *model.ResponseScheme, error) {
	return a.internalClient.Gets(ctx, contentID, startAt, maxResults, options)
}

// CreateOrUpdate adds an attachment to a piece of content.
//
// If the attachment already exists for the content,
//
// then the attachment is updated (i.e. a new version of the attachment is created).
//
// PUT /wiki/rest/api/content/{id}/child/attachment
//
// https://docs.go-atlassian.io/confluence-cloud/content/attachments#create-or-update-attachment
func (a *ContentAttachmentService) CreateOrUpdate(ctx context.Context, attachmentID, status, fileName string, file io.Reader) (*model.ContentPageScheme, *model.ResponseScheme, error) {
	return a.internalClient.CreateOrUpdate(ctx, attachmentID, status, fileName, file)
}

// Create adds an attachment to a piece of content.
//
// This method only adds a new attachment.
//
// If you want to update an existing attachment, use Create or update attachments.
//
// POST /wiki/rest/api/content/{id}/child/attachment
//
// https://docs.go-atlassian.io/confluence-cloud/content/attachments#create-attachment
func (a *ContentAttachmentService) Create(ctx context.Context, attachmentID, status, fileName string, file io.Reader) (*model.ContentPageScheme, *model.ResponseScheme, error) {
	return a.internalClient.Create(ctx, attachmentID, status, fileName, file)
}

type internalContentAttachmentImpl struct {
	c service.Connector
}

func (i *internalContentAttachmentImpl) Gets(ctx context.Context, contentID string, startAt, maxResults int, options *model.GetContentAttachmentsOptionsScheme) (*model.ContentPageScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	query := url.Values{}
	query.Add("start", strconv.Itoa(startAt))
	query.Add("limit", strconv.Itoa(maxResults))

	if options != nil {

		if len(options.Expand) != 0 {
			query.Add("expand", strings.Join(options.Expand, ","))
		}

		if options.FileName != "" {
			query.Add("filename", options.FileName)
		}

		if options.MediaType != "" {
			query.Add("mediaType", options.MediaType)
		}

	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/child/attachment?%v", contentID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
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

func (i *internalContentAttachmentImpl) CreateOrUpdate(ctx context.Context, attachmentID, status, fileName string, file io.Reader) (*model.ContentPageScheme, *model.ResponseScheme, error) {

	if attachmentID == "" {
		return nil, nil, model.ErrNoContentAttachmentID
	}

	if fileName == "" {
		return nil, nil, model.ErrNoContentAttachmentName
	}

	if file == nil {
		return nil, nil, model.ErrNoContentReader
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/child/attachment", attachmentID))

	if status != "" {
		query := url.Values{}
		query.Add("status", status)

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

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

	if err = writer.WriteField("minorEdit", "true"); err != nil {
		return nil, nil, err
	}

	writer.Close()

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint.String(), writer.FormDataContentType(), reader)
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

func (i *internalContentAttachmentImpl) Create(ctx context.Context, attachmentID, status, fileName string, file io.Reader) (*model.ContentPageScheme, *model.ResponseScheme, error) {

	if attachmentID == "" {
		return nil, nil, model.ErrNoContentAttachmentID
	}

	if fileName == "" {
		return nil, nil, model.ErrNoContentAttachmentName
	}

	if file == nil {
		return nil, nil, model.ErrNoContentReader
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/rest/api/content/%v/child/attachment", attachmentID))

	if status != "" {
		query := url.Values{}
		query.Add("status", status)

		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

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

	if err = writer.WriteField("minorEdit", "true"); err != nil {
		return nil, nil, err
	}

	writer.Close()

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), writer.FormDataContentType(), reader)
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
