package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/confluence"
	"net/http"
	"net/url"
	"strconv"
)

// NewAttachmentVersionService returns a new Confluence V2 Attachment Version service
func NewAttachmentVersionService(client service.Connector) *AttachmentVersionService {
	return &AttachmentVersionService{
		internalClient: &internalAttachmentVersionImpl{c: client},
	}
}

type AttachmentVersionService struct {
	internalClient confluence.AttachmentVersionConnector
}

// Gets returns the versions of specific attachment.
//
// GET /wiki/api/v2/attachments/{id}/versions
func (a *AttachmentVersionService) Gets(ctx context.Context, attachmentID, cursor, sort string, limit int) (*model.AttachmentVersionPageScheme, *model.ResponseScheme, error) {
	return a.internalClient.Gets(ctx, attachmentID, cursor, sort, limit)
}

// Get retrieves version details for the specified attachment and version number.
//
// GET /wiki/api/v2/attachments/{attachment-id}/versions/{version-number}
func (a *AttachmentVersionService) Get(ctx context.Context, attachmentID string, versionID int) (*model.DetailedVersionScheme, *model.ResponseScheme, error) {
	return a.internalClient.Get(ctx, attachmentID, versionID)
}

type internalAttachmentVersionImpl struct {
	c service.Connector
}

func (i *internalAttachmentVersionImpl) Gets(ctx context.Context, attachmentID, cursor, sort string, limit int) (*model.AttachmentVersionPageScheme, *model.ResponseScheme, error) {

	if attachmentID == "" {
		return nil, nil, model.ErrNoContentAttachmentIDError
	}

	query := url.Values{}
	query.Add("limit", strconv.Itoa(limit))

	if cursor != "" {
		query.Add("cursor", cursor)
	}

	if sort != "" {
		query.Add("sort", sort)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/attachments/%v/versions?%v", attachmentID, query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.AttachmentVersionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalAttachmentVersionImpl) Get(ctx context.Context, attachmentID string, versionID int) (*model.DetailedVersionScheme, *model.ResponseScheme, error) {

	if attachmentID == "" {
		return nil, nil, model.ErrNoContentAttachmentIDError
	}

	endpoint := fmt.Sprintf("wiki/api/v2/attachments/%v/versions/%v", attachmentID, versionID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	version := new(model.DetailedVersionScheme)
	response, err := i.c.Call(request, version)
	if err != nil {
		return nil, response, err
	}

	return version, response, nil
}
