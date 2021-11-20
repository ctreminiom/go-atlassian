package sm

import (
	"bytes"
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
)

type ServiceDeskService struct {
	client *Client
	Queue  *ServiceDeskQueueService
}

// Gets returns all the service desks in the Jira Service Management instance that the user has permission to access.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk#get-service-desks
func (s *ServiceDeskService) Gets(ctx context.Context, start, limit int) (result *model.ServiceDeskPageScheme,
	response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk?%v", params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns a service desk.
// Use this method to get service desk details whenever your application component is passed a service desk ID
// but needs to display other service desk details.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk#get-service-desk-by-id
func (s *ServiceDeskService) Get(ctx context.Context, serviceDeskID int) (result *model.ServiceDeskScheme,
	response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v", serviceDeskID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Attach one temporary attachments to a service desk
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk#attach-temporary-file
func (s *ServiceDeskService) Attach(ctx context.Context, serviceDeskID int, fileName string, file io.Reader) (
	result *model.ServiceDeskTemporaryFileScheme, response *ResponseScheme, err error) {

	if serviceDeskID == 0 {
		return nil, nil, model.ErrNoServiceDeskIDError
	}

	if len(fileName) == 0 {
		return nil, nil, model.ErrNoFileNameError
	}

	if file == nil {
		return nil, nil, model.ErrNoFileReaderError
	}

	var (
		body             = &bytes.Buffer{}
		attachmentWriter = multipart.NewWriter(body)
	)

	// create the attachment form row
	part, _ := attachmentWriter.CreateFormFile("file", fileName)

	// add the attachment metadata
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, nil, err
	}

	attachmentWriter.Close()

	var endpoint = fmt.Sprintf("/rest/servicedeskapi/servicedesk/%v/attachTemporaryFile", serviceDeskID)

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, body)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Add("Content-Type", attachmentWriter.FormDataContentType())
	request.Header.Set("X-Atlassian-Token", "no-check")
	request.Header.Add("Accept", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}
