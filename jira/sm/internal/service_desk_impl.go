package internal

import (
	
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/sm"
)

// NewServiceDeskService creates a new instance of ServiceDeskService.
// It takes a service.Connector, a version string, and a QueueService as input.
// Returns a pointer to ServiceDeskService and an error if the version is not provided.
func NewServiceDeskService(client service.Connector, version string, queue *QueueService) (*ServiceDeskService, error) {

	if version == "" {
		return nil, fmt.Errorf("client: %w", model.ErrNoVersionProvided)
	}

	return &ServiceDeskService{
		internalClient: &internalServiceDeskImpl{c: client, version: version},
		Queue:          queue,
	}, nil
}

// ServiceDeskService provides methods to interact with service desk operations in Jira Service Management.
type ServiceDeskService struct {
	// internalClient is the connector interface for service desk operations.
	internalClient sm.ServiceDeskConnector
	// Queue handles queue operations.
	Queue *QueueService
}

// Gets returns all the service desks in the Jira Service Management instance that the user has permission to access.
//
// GET /rest/servicedeskapi/servicedesk
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk#get-service-desks
func (s *ServiceDeskService) Gets(ctx context.Context, start, limit int) (*model.ServiceDeskPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ServiceDeskService).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets"))

	return s.internalClient.Gets(ctx, start, limit)
}

// Get returns a service desk.
//
// # Use this method to get service desk details whenever your application component is passed a service desk ID
//
// but needs to display other service desk details.
//
// GET /rest/servicedeskapi/servicedesk/{serviceDeskId}
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk#get-service-desk-by-id
func (s *ServiceDeskService) Get(ctx context.Context, serviceDeskID string) (*model.ServiceDeskScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ServiceDeskService).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	return s.internalClient.Get(ctx, serviceDeskID)
}

// Attach one temporary attachments to a service desk
//
// POST /rest/servicedeskapi/servicedesk/{serviceDeskId}/attachTemporaryFile
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk#attach-temporary-file
func (s *ServiceDeskService) Attach(ctx context.Context, serviceDeskID string, fileName string, file io.Reader) (*model.ServiceDeskTemporaryFileScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ServiceDeskService).Attach", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "attach"))

	return s.internalClient.Attach(ctx, serviceDeskID, fileName, file)
}

type internalServiceDeskImpl struct {
	c       service.Connector
	version string
}

func (i *internalServiceDeskImpl) Gets(ctx context.Context, start, limit int) (*model.ServiceDeskPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalServiceDeskImpl).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "gets"))

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk?%v", params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	page := new(model.ServiceDeskPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return page, res, nil
}

func (i *internalServiceDeskImpl) Get(ctx context.Context, serviceDeskID string) (*model.ServiceDeskScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalServiceDeskImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get"))

	if serviceDeskID == "" {

			return nil, nil, fmt.Errorf("sm: %w", model.ErrNoServiceDeskID)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v", serviceDeskID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)

		return nil, nil, err
	}

	serviceDesk := new(model.ServiceDeskScheme)
	res, err := i.c.Call(req, serviceDesk)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return serviceDesk, res, nil
}

func (i *internalServiceDeskImpl) Attach(ctx context.Context, serviceDeskID string, fileName string, file io.Reader) (*model.ServiceDeskTemporaryFileScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalServiceDeskImpl).Attach", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "attach"))

	if serviceDeskID == "" {

			return nil, nil, fmt.Errorf("sm: %w", model.ErrNoServiceDeskID)
	}

	if fileName == "" {

			return nil, nil, fmt.Errorf("sm: %w", model.ErrNoFileName)
	}

	if file == nil {

			return nil, nil, fmt.Errorf("sm: %w", model.ErrNoFileReader)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/attachTemporaryFile", serviceDeskID)

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

	if err := writer.Close(); err != nil {
		return nil, nil, err
	}

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, writer.FormDataContentType(), reader)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	attachmentID := new(model.ServiceDeskTemporaryFileScheme)
	res, err := i.c.Call(req, attachmentID)
	if err != nil {
		recordError(span, err)
		return nil, res, err
	}

	setOK(span)
	return attachmentID, res, nil
}
