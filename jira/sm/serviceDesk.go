package sm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

type ServiceDeskService struct {
	client *Client
	Queue  *ServiceDeskQueueService
}

// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk#get-service-desks
func (s *ServiceDeskService) Gets(ctx context.Context, start, limit int) (result *ServiceDeskPageScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk?%v", params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ServiceDeskPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk#get-service-desk-by-id
func (s *ServiceDeskService) Get(ctx context.Context, serviceDeskID int) (result *ServiceDeskScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v", serviceDeskID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ServiceDeskScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk#attach-temporary-file
func (s *ServiceDeskService) Attach(ctx context.Context, serviceDeskID int, path string) (result *ServiceDeskTemporaryFileScheme, response *Response, err error) {

	if len(path) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid path value")
	}

	if !filepath.IsAbs(path) {
		return nil, nil, fmt.Errorf("the path provided is not an absolute path, please provide a valid one")
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	filePart, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return
	}

	_, err = io.Copy(filePart, file)
	if err != nil {
		return
	}

	if err = writer.Close(); err != nil {
		return
	}

	var endpoint = fmt.Sprintf("%vrest/servicedeskapi/servicedesk/%v/attachTemporaryFile", s.client.Site.String(), serviceDeskID)
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return
	}
	request.SetBasicAuth(s.client.Auth.mail, s.client.Auth.token)

	request.Header.Add("Content-Type", writer.FormDataContentType())
	request.Header.Add("Accept", "application/json")
	request.Header.Set("X-Atlassian-Token", "no-check")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ServiceDeskTemporaryFileScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

type ServiceDeskTemporaryFileScheme struct {
	TemporaryAttachments []struct {
		TemporaryAttachmentID string `json:"temporaryAttachmentId"`
		FileName              string `json:"fileName"`
	} `json:"temporaryAttachments"`
}

type ServiceDeskPageScheme struct {
	Expands    []interface{} `json:"_expands"`
	Size       int           `json:"size"`
	Start      int           `json:"start"`
	Limit      int           `json:"limit"`
	IsLastPage bool          `json:"isLastPage"`
	Links      struct {
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
	} `json:"_links"`
	Values []*ServiceDeskScheme `json:"values"`
}

type ServiceDeskScheme struct {
	ID          string `json:"id"`
	ProjectID   string `json:"projectId"`
	ProjectName string `json:"projectName"`
	ProjectKey  string `json:"projectKey"`
	Links       struct {
		Self string `json:"self"`
	} `json:"_links"`
}
