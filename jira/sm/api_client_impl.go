package sm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/ctreminiom/go-atlassian/v2/jira/sm/internal"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service/common"
)

const defaultServiceManagementVersion = "latest"

func New(httpClient common.HTTPClient, site string) (*Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if site == "" {
		return nil, model.ErrNoSite
	}

	if !strings.HasSuffix(site, "/") {
		site += "/"
	}

	u, err := url.Parse(site)
	if err != nil {
		return nil, err
	}

	client := &Client{
		HTTP: httpClient,
		Site: u,
	}

	client.Auth = internal.NewAuthenticationService(client)
	client.Customer = internal.NewCustomerService(client, defaultServiceManagementVersion)
	client.Info = internal.NewInfoService(client, defaultServiceManagementVersion)
	client.Knowledgebase = internal.NewKnowledgebaseService(client, defaultServiceManagementVersion)
	client.Organization = internal.NewOrganizationService(client, defaultServiceManagementVersion)
	client.WorkSpace = internal.NewWorkSpaceService(client, defaultServiceManagementVersion)

	requestSubServices := &internal.ServiceRequestSubServices{
		Approval:    internal.NewApprovalService(client, defaultServiceManagementVersion),
		Attachment:  internal.NewAttachmentService(client, defaultServiceManagementVersion),
		Comment:     internal.NewCommentService(client, defaultServiceManagementVersion),
		Participant: internal.NewParticipantService(client, defaultServiceManagementVersion),
		SLA:         internal.NewServiceLevelAgreementService(client, defaultServiceManagementVersion),
		Feedback:    internal.NewFeedbackService(client, defaultServiceManagementVersion),
		Type:        internal.NewTypeService(client, defaultServiceManagementVersion),
	}

	requestService, err := internal.NewRequestService(client, defaultServiceManagementVersion, requestSubServices)
	if err != nil {
		return nil, err
	}
	client.Request = requestService

	serviceDeskService, err := internal.NewServiceDeskService(
		client,
		defaultServiceManagementVersion,
		internal.NewQueueService(client, defaultServiceManagementVersion))

	if err != nil {
		return nil, err
	}
	client.ServiceDesk = serviceDeskService

	return client, nil
}

type Client struct {
	HTTP          common.HTTPClient
	Site          *url.URL
	Auth          common.Authentication
	Customer      *internal.CustomerService
	Info          *internal.InfoService
	Knowledgebase *internal.KnowledgebaseService
	Organization  *internal.OrganizationService
	Request       *internal.RequestService
	ServiceDesk   *internal.ServiceDeskService
	WorkSpace     *internal.WorkSpaceService
}

func (c *Client) NewRequest(ctx context.Context, method, urlStr, contentType string, body interface{}) (*http.Request, error) {
	ctx, span := tracer().Start(ctx, "(*Client).NewRequest")
	defer span.End()

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.Site.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		if err = json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	// If the body interface is a *bytes.Buffer type
	// it means the NewRequest() requires to handle the RFC 1867 ISO
	if attachBuffer, ok := body.(*bytes.Buffer); ok {
		buf = attachBuffer
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	if contentType != "" {
		// When the contentType is provided, it means the request needs to be created to handle files
		req.Header.Set("Content-Type", contentType)
		req.Header.Set("X-Atlassian-Token", "no-check")
	}

	if c.Auth.HasBasicAuth() {
		req.SetBasicAuth(c.Auth.GetBasicAuth())
	}

	if c.Auth.HasSetExperimentalFlag() {
		req.Header.Set("X-ExperimentalApi", "opt-in")
	}

	if c.Auth.GetBearerToken() != "" && !c.Auth.HasBasicAuth() {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", c.Auth.GetBearerToken()))
	}

	if c.Auth.HasUserAgent() {
		req.Header.Set("User-Agent", c.Auth.GetUserAgent())
	}

	return req, nil
}

func (c *Client) Call(request *http.Request, structure interface{}) (*model.ResponseScheme, error) {

	response, err := c.HTTP.Do(request)
	if err != nil {
		return nil, err
	}

	return c.processResponse(response, structure)
}

func (c *Client) processResponse(response *http.Response, structure interface{}) (*model.ResponseScheme, error) {

	defer response.Body.Close()

	res := &model.ResponseScheme{
		Response: response,
		Code:     response.StatusCode,
		Endpoint: response.Request.URL.String(),
		Method:   response.Request.Method,
	}

	responseAsBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return res, err
	}

	res.Bytes.Write(responseAsBytes)

	wasSuccess := response.StatusCode >= 200 && response.StatusCode < 300

	if !wasSuccess {

		switch response.StatusCode {

		case http.StatusNotFound:
			return res, model.ErrNotFound

		case http.StatusUnauthorized:
			return res, model.ErrUnauthorized

		case http.StatusInternalServerError:
			return res, model.ErrInternal

		case http.StatusBadRequest:
			return res, model.ErrBadRequest

		default:
			return res, model.ErrInvalidStatusCode
		}
	}

	if structure != nil {
		if err = json.Unmarshal(responseAsBytes, &structure); err != nil {
			return res, err
		}
	}

	return res, nil
}
