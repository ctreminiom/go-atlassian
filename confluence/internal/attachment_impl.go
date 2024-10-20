package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
)

// NewAttachmentService creates a new instance of AttachmentService.
// It takes a service.Connector and an AttachmentVersionService as inputs and returns a pointer to AttachmentService.
func NewAttachmentService(client service.Connector, version *AttachmentVersionService) *AttachmentService {
	return &AttachmentService{
		internalClient: &internalAttachmentImpl{c: client},
		Version:        version,
	}
}

// AttachmentService provides methods to interact with attachment operations in Confluence.
type AttachmentService struct {
	// internalClient is the connector interface for attachment operations.
	internalClient confluence.AttachmentConnector
	// Version is the service for attachment version-related operations.
	Version *AttachmentVersionService
}

// Get returns a specific attachment.
//
// GET /wiki/api/v2/attachments/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/v2/attachments#get-attachment-by-id
func (a *AttachmentService) Get(ctx context.Context, attachmentID string, versionID int, serializeIDs bool) (*model.AttachmentScheme, *model.ResponseScheme, error) {
	return a.internalClient.Get(ctx, attachmentID, versionID, serializeIDs)
}

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
func (a *AttachmentService) Gets(ctx context.Context, entityID int, entityType string, options *model.AttachmentParamsScheme, cursor string, limit int) (*model.AttachmentPageScheme, *model.ResponseScheme, error) {
	return a.internalClient.Gets(ctx, entityID, entityType, options, cursor, limit)
}

// Delete deletes an attachment by id.
//
// DELETE /wiki/api/v2/attachments/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/v2/attachments#delete-attachment
func (a *AttachmentService) Delete(ctx context.Context, attachmentID string) (*model.ResponseScheme, error) {
	return a.internalClient.Delete(ctx, attachmentID)
}

type internalAttachmentImpl struct {
	c service.Connector
}

func (i *internalAttachmentImpl) Delete(ctx context.Context, attachmentID string) (*model.ResponseScheme, error) {

	if attachmentID == "" {
		return nil, model.ErrNoContentAttachmentID
	}

	endpoint := fmt.Sprintf("wiki/api/v2/attachments/%v", attachmentID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalAttachmentImpl) Get(ctx context.Context, attachmentID string, versionID int, serializeIDs bool) (*model.AttachmentScheme, *model.ResponseScheme, error) {

	if attachmentID == "" {
		return nil, nil, model.ErrNoContentAttachmentID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("wiki/api/v2/attachments/%v", attachmentID))

	query := url.Values{}
	if versionID != 0 {
		query.Add("version", strconv.Itoa(versionID))
	}

	if serializeIDs {
		query.Add("serialize-ids-as-strings", "true")
	}

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	attachment := new(model.AttachmentScheme)
	response, err := i.c.Call(request, attachment)
	if err != nil {
		return nil, response, err
	}

	return attachment, response, nil
}

func (i *internalAttachmentImpl) Gets(ctx context.Context, entityID int, entityType string, options *model.AttachmentParamsScheme, cursor string, limit int) (*model.AttachmentPageScheme, *model.ResponseScheme, error) {

	if entityID == 0 {
		return nil, nil, model.ErrNoEntityID
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	if options != nil {

		if options.SerializeIDs {
			query.Add("serialize-ids-as-strings", "true")
		}

		if options.MediaType != "" {
			query.Add("mediaType", options.MediaType)
		}

		if options.FileName != "" {
			query.Add("filename", options.FileName)
		}

		if options.Sort != "" {
			query.Add("sort", options.Sort)
		}
	}

	// Checking if the entity type provided is supported by the library
	var isSupported bool
	for _, typ := range model.ValidEntityValues {

		if entityType == typ {
			isSupported = true
			break
		}
	}

	if !isSupported {
		return nil, nil, model.ErrNoEntityValue
	}

	endpoint := fmt.Sprintf("wiki/api/v2/%v/%v/attachments?%v", entityType, entityID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.AttachmentPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}
