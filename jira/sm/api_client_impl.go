package sm

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/ctreminiom/go-atlassian/jira/sm/internal"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service/common"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

func NewV2(httpClient common.HttpClient, site string) (*ClientV2, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if !strings.HasSuffix(site, "/") {
		site += "/"
	}

	siteAsURL, err := url.Parse(site)
	if err != nil {
		return nil, err
	}

	client := &ClientV2{
		HTTP: httpClient,
		Site: siteAsURL,
	}

	customerService, err := internal.NewCustomerService(client, "latest")
	if err != nil {
		return nil, err
	}
	client.Customer = customerService

	infoService, err := internal.NewInfoService(client, "latest")
	if err != nil {
		return nil, err
	}
	client.Info = infoService

	knowledgebaseService, err := internal.NewKnowledgebaseService(client, "latest")
	if err != nil {
		return nil, err
	}
	client.Knowledgebase = knowledgebaseService

	organizationService, err := internal.NewOrganizationService(client, "latest")
	if err != nil {
		return nil, err
	}
	client.Organization = organizationService

	commentService, err := internal.NewCommentService(client, "latest")
	if err != nil {
		return nil, err
	}

	attachmentService, err := internal.NewAttachmentService(client, "latest")
	if err != nil {
		return nil, err
	}

	approvalService, err := internal.NewApprovalService(client, "latest")
	if err != nil {
		return nil, err
	}

	participantService, err := internal.NewParticipantService(client, "latest")
	if err != nil {
		return nil, err
	}

	slaService, err := internal.NewServiceLevelAgreementService(client, "latest")
	if err != nil {
		return nil, err
	}

	feedbackService, err := internal.NewFeedbackService(client, "latest")
	if err != nil {
		return nil, err
	}

	requestTypeService, err := internal.NewTypeService(client, "latest")
	if err != nil {
		return nil, err
	}

	requestServices := &internal.ServiceRequestSubServices{
		Approval:    approvalService,
		Attachment:  attachmentService,
		Comment:     commentService,
		Participant: participantService,
		SLA:         slaService,
		Feedback:    feedbackService,
		Type:        requestTypeService,
	}

	requestService, err := internal.NewRequestService(client, "latest", requestServices)
	if err != nil {
		return nil, err
	}
	client.Request = requestService

	queueService, err := internal.NewQueueService(client, "latest")
	if err != nil {
		return nil, err
	}

	serviceDeskService, err := internal.NewServiceDeskService(client, "latest", queueService)
	if err != nil {
		return nil, err
	}
	client.ServiceDesk = serviceDeskService
	return client, nil
}

type ClientV2 struct {
	HTTP          common.HttpClient
	Auth          common.Authentication
	Site          *url.URL
	Customer      *internal.CustomerService
	Info          *internal.InfoService
	Knowledgebase *internal.KnowledgebaseService
	Organization  *internal.OrganizationService
	Request       *internal.RequestService
	ServiceDesk   *internal.ServiceDeskService
}

func (c *ClientV2) NewFormRequest(ctx context.Context, method, apiEndpoint, contentType string, payload io.Reader) (*http.Request, error) {

	relativePath, err := url.Parse(apiEndpoint)
	if err != nil {
		return nil, err
	}

	var endpoint = c.Site.ResolveReference(relativePath).String()

	request, err := http.NewRequestWithContext(ctx, method, endpoint, payload)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", contentType)
	request.Header.Add("Accept", "application/json")
	request.Header.Set("X-Atlassian-Token", "no-check")

	if c.Auth.HasBasicAuth() {
		request.SetBasicAuth(c.Auth.GetBasicAuth())
	}

	if c.Auth.HasUserAgent() {
		request.Header.Set("User-Agent", c.Auth.GetUserAgent())
	}

	return request, nil
}

func (c *ClientV2) NewRequest(ctx context.Context, method, apiEndpoint string, payload io.Reader) (*http.Request, error) {

	relativePath, err := url.Parse(apiEndpoint)
	if err != nil {
		return nil, err
	}

	var endpoint = c.Site.ResolveReference(relativePath).String()

	request, err := http.NewRequestWithContext(ctx, method, endpoint, payload)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json")

	if payload != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	if c.Auth.HasBasicAuth() {
		request.SetBasicAuth(c.Auth.GetBasicAuth())
	}

	if c.Auth.HasUserAgent() {
		request.Header.Set("User-Agent", c.Auth.GetUserAgent())
	}

	return request, nil
}

func (c *ClientV2) Call(request *http.Request, structure interface{}) (*models.ResponseScheme, error) {

	response, err := c.HTTP.Do(request)
	if err != nil {
		return nil, err
	}

	return c.TransformTheHTTPResponse(response, structure)
}

func (c *ClientV2) TransformTheHTTPResponse(response *http.Response, structure interface{}) (*models.ResponseScheme, error) {

	responseTransformed := &models.ResponseScheme{
		Response: response,
		Code:     response.StatusCode,
		Endpoint: response.Request.URL.String(),
		Method:   response.Request.Method,
	}

	responseAsBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return responseTransformed, err
	}

	responseTransformed.Bytes.Write(responseAsBytes)

	var wasSuccess = response.StatusCode >= 200 && response.StatusCode < 300
	if !wasSuccess {
		return responseTransformed, models.ErrInvalidStatusCodeError
	}

	if structure != nil {
		if err = json.Unmarshal(responseAsBytes, &structure); err != nil {
			return responseTransformed, err
		}
	}

	return responseTransformed, nil
}

func (c *ClientV2) TransformStructToReader(structure interface{}) (io.Reader, error) {

	if structure == nil {
		return nil, models.ErrNilPayloadError
	}

	if reflect.ValueOf(structure).Type().Kind() == reflect.Struct {
		return nil, models.ErrNonPayloadPointerError
	}

	structureAsBodyBytes, err := json.Marshal(structure)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(structureAsBodyBytes), nil
}
