package internal

import (
	"bytes"
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/sm"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

func NewServiceDeskService(client service.Connector, version string, queue *QueueService) (*ServiceDeskService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ServiceDeskService{
		internalClient: &internalServiceDeskImpl{c: client, version: version},
		Queue:          queue,
	}, nil
}

type ServiceDeskService struct {
	internalClient sm.ServiceDeskConnector
	Queue          *QueueService
}

// Gets returns all the service desks in the Jira Service Management instance that the user has permission to access.
//
// GET /rest/servicedeskapi/servicedesk
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk#get-service-desks
func (s *ServiceDeskService) Gets(ctx context.Context, start, limit int) (*model.ServiceDeskPageScheme, *model.ResponseScheme, error) {
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
func (s *ServiceDeskService) Get(ctx context.Context, serviceDeskID int) (*model.ServiceDeskScheme, *model.ResponseScheme, error) {
	return s.internalClient.Get(ctx, serviceDeskID)
}

// Attach one temporary attachments to a service desk
//
// POST /rest/servicedeskapi/servicedesk/{serviceDeskId}/attachTemporaryFile
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk#attach-temporary-file
func (s *ServiceDeskService) Attach(ctx context.Context, serviceDeskID int, fileName string, file io.Reader) (*model.ServiceDeskTemporaryFileScheme, *model.ResponseScheme, error) {
	return s.internalClient.Attach(ctx, serviceDeskID, fileName, file)
}

type internalServiceDeskImpl struct {
	c       service.Connector
	version string
}

func (i *internalServiceDeskImpl) Gets(ctx context.Context, start, limit int) (*model.ServiceDeskPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk?%v", params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ServiceDeskPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalServiceDeskImpl) Get(ctx context.Context, serviceDeskID int) (*model.ServiceDeskScheme, *model.ResponseScheme, error) {

	if serviceDeskID == 0 {
		return nil, nil, model.ErrNoServiceDeskIDError
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v", serviceDeskID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	serviceDesk := new(model.ServiceDeskScheme)
	res, err := i.c.Call(req, serviceDesk)
	if err != nil {
		return nil, res, err
	}

	return serviceDesk, res, nil
}

func (i *internalServiceDeskImpl) Attach(ctx context.Context, serviceDeskID int, fileName string, file io.Reader) (*model.ServiceDeskTemporaryFileScheme, *model.ResponseScheme, error) {

	if serviceDeskID == 0 {
		return nil, nil, model.ErrNoServiceDeskIDError
	}

	if fileName == "" {
		return nil, nil, model.ErrNoFileNameError
	}

	if file == nil {
		return nil, nil, model.ErrNoFileReaderError
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/attachTemporaryFile", serviceDeskID)

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

	if err := writer.Close(); err != nil {
		return nil, nil, err
	}

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, writer.FormDataContentType(), reader)
	if err != nil {
		return nil, nil, err
	}

	attachmentID := new(model.ServiceDeskTemporaryFileScheme)
	res, err := i.c.Call(req, attachmentID)
	if err != nil {
		return nil, res, err
	}

	return attachmentID, res, nil
}
