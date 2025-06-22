package internal

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
	"net/http"
	"net/url"
	"strconv"
)

// NewAttachmentVersionService creates a new instance of AttachmentVersionService.
// It takes a service.Connector as input and returns a pointer to AttachmentVersionService.
func NewAttachmentVersionService(client service.Connector) *AttachmentVersionService {
	return &AttachmentVersionService{
		internalClient: &internalAttachmentVersionImpl{c: client},
	}
}

// AttachmentVersionService provides methods to interact with attachment version operations in Confluence.
type AttachmentVersionService struct {
	// internalClient is the connector interface for attachment version operations.
	internalClient confluence.AttachmentVersionConnector
}

// Gets returns the versions of specific attachment.
//
// GET /wiki/api/v2/attachments/{id}/versions
//
// https://docs.go-atlassian.io/confluence-cloud/v2/attachments/versions#get-attachment-versions
func (a *AttachmentVersionService) Gets(ctx context.Context, attachmentID, cursor, sort string, limit int) (*model.AttachmentVersionPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*AttachmentVersionService).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets"))

	return a.internalClient.Gets(ctx, attachmentID, cursor, sort, limit)
}

// Get retrieves version details for the specified attachment and version number.
//
// GET /wiki/api/v2/attachments/{attachment-id}/versions/{version-number}
//
// https://docs.go-atlassian.io/confluence-cloud/v2/attachments/versions#get-attachment-version
func (a *AttachmentVersionService) Get(ctx context.Context, attachmentID string, versionID int) (*model.DetailedVersionScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*AttachmentVersionService).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	return a.internalClient.Get(ctx, attachmentID, versionID)
}

type internalAttachmentVersionImpl struct {
	c service.Connector
}

func (i *internalAttachmentVersionImpl) Gets(ctx context.Context, attachmentID, cursor, sort string, limit int) (*model.AttachmentVersionPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalAttachmentVersionImpl).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets"))

	if attachmentID == "" {

			return nil, nil, fmt.Errorf("confluence: %w", model.ErrNoContentAttachmentID)
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
		recordError(span, err)
		return nil, nil, err
	}

	page := new(model.AttachmentVersionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalAttachmentVersionImpl) Get(ctx context.Context, attachmentID string, versionID int) (*model.DetailedVersionScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalAttachmentVersionImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	if attachmentID == "" {

			return nil, nil, fmt.Errorf("confluence: %w", model.ErrNoContentAttachmentID)
	}

	endpoint := fmt.Sprintf("wiki/api/v2/attachments/%v/versions/%v", attachmentID, versionID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	version := new(model.DetailedVersionScheme)
	response, err := i.c.Call(request, version)
	if err != nil {
		return nil, response, err
	}

	setOK(span)
	return version, response, nil
}
