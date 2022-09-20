package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/sm"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ServiceRequestSubServices struct {
}

func NewRequestService(client service.Client, version string, subServices *ServiceRequestSubServices) (*RequestService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &RequestService{
		internalClient: &internalServiceRequestImpl{c: client, version: version},
	}, nil
}

type RequestService struct {
	internalClient sm.RequestConnector
}

// Create creates a customer request in a service desk.
//
// The JSON request must include the service desk and customer request type, as well as any fields that are required for the request type.
//
// POST /rest/servicedeskapi/request
//
// https://docs.go-atlassian.io/jira-service-management/request#create-customer-request
func (s *RequestService) Create(ctx context.Context, payload *model.CreateCustomerRequestPayloadScheme, fields *model.CustomerRequestFields) (*model.CustomerRequestScheme, *model.ResponseScheme, error) {
	return s.internalClient.Create(ctx, payload, fields)
}

// Gets returns all customer requests for the user executing the query.
//
// The returned customer requests are ordered chronologically by the latest activity on each request. For example, the latest status transition or comment.
//
// GET /rest/servicedeskapi/request
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request#get-customer-requests
func (s *RequestService) Gets(ctx context.Context, options *model.ServiceRequestOptionScheme, start, limit int) (*model.CustomerRequestPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Gets(ctx, options, start, limit)
}

// Get returns a customer request.
//
// GET /rest/servicedeskapi/request/{issueIdOrKey}
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request#get-customer-request-by-id-or-key
func (s *RequestService) Get(ctx context.Context, issueKeyOrID string, expand []string) (*model.CustomerRequestScheme, *model.ResponseScheme, error) {
	return s.internalClient.Get(ctx, issueKeyOrID, expand)
}

// Subscribe subscribes the user to receiving notifications from a customer request.
//
// PUT /rest/servicedeskapi/request/{issueIdOrKey}/notification
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request#subscribe
func (s *RequestService) Subscribe(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error) {
	return s.internalClient.Subscribe(ctx, issueKeyOrID)
}

// Unsubscribe unsubscribes the user from notifications from a customer request.
//
// DELETE /rest/servicedeskapi/request/{issueIdOrKey}/notification
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request#unsubscribe
func (s *RequestService) Unsubscribe(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error) {
	return s.internalClient.Unsubscribe(ctx, issueKeyOrID)
}

// Transitions returns a list of transitions, the workflow processes that moves a customer request from one status to another, that the user can perform on a request.
//
// GET /rest/servicedeskapi/request/{issueIdOrKey}/transition
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request#get-customer-transitions
func (s *RequestService) Transitions(ctx context.Context, issueKeyOrID string, start, limit int) (*model.CustomerRequestTransitionPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Transitions(ctx, issueKeyOrID, start, limit)
}

// Transition performs a customer transition for a given request and transition.
//
// An optional comment can be included to provide a reason for the transition.
//
// POST /rest/servicedeskapi/request/{issueIdOrKey}/transition
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request#perform-customer-transition
func (s *RequestService) Transition(ctx context.Context, issueKeyOrID, transitionID, comment string) (*model.ResponseScheme, error) {
	return s.internalClient.Transition(ctx, issueKeyOrID, transitionID, comment)
}

type internalServiceRequestImpl struct {
	c       service.Client
	version string
}

func (i *internalServiceRequestImpl) Create(ctx context.Context, payload *model.CreateCustomerRequestPayloadScheme, fields *model.CustomerRequestFields) (*model.CustomerRequestScheme, *model.ResponseScheme, error) {

	if fields == nil || fields.Fields == nil {
		return nil, nil, model.ErrNoCustomRequestFieldsError
	}

	payloadWithFields, err := payload.MergeFields(fields)
	if err != nil {
		return nil, nil, err
	}

	reader, err := i.c.TransformStructToReader(&payloadWithFields)
	if err != nil {
		return nil, nil, err
	}

	endpoint := "rest/servicedeskapi/request"

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	serviceRequest := new(model.CustomerRequestScheme)
	response, err := i.c.Call(request, serviceRequest)
	if err != nil {
		return nil, response, err
	}

	return serviceRequest, response, nil
}

func (i *internalServiceRequestImpl) Gets(ctx context.Context, options *model.ServiceRequestOptionScheme, start, limit int) (*model.CustomerRequestPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if options != nil {

		if len(options.SearchTerm) != 0 {
			params.Add("searchTerm", options.SearchTerm)
		}

		for _, requestOwner := range options.RequestOwnerships {
			params.Add("requestOwnership", requestOwner)
		}

		if len(options.RequestStatus) != 0 {
			params.Add("requestStatus", options.RequestStatus)
		}

		if len(options.ApprovalStatus) != 0 {
			params.Add("approvalStatus", options.ApprovalStatus)
		}

		if options.OrganizationID != 0 {
			params.Add("organizationId", strconv.Itoa(options.OrganizationID))
		}

		if options.ServiceDeskID != 0 {
			params.Add("serviceDeskId", strconv.Itoa(options.ServiceDeskID))
		}

		if options.RequestTypeID != 0 {
			params.Add("requestTypeId", strconv.Itoa(options.RequestTypeID))
		}

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request?%v", params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.CustomerRequestPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalServiceRequestImpl) Get(ctx context.Context, issueKeyOrID string, expand []string) (*model.CustomerRequestScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/servicedeskapi/request/%v", issueKeyOrID))

	if expand != nil {
		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	customerRequest := new(model.CustomerRequestScheme)
	response, err := i.c.Call(request, customerRequest)
	if err != nil {
		return nil, response, err
	}

	return customerRequest, response, nil
}

func (i *internalServiceRequestImpl) Subscribe(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/notification", issueKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalServiceRequestImpl) Unsubscribe(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/notification", issueKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalServiceRequestImpl) Transitions(ctx context.Context, issueKeyOrID string, start, limit int) (*model.CustomerRequestTransitionPageScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/transition?%v", issueKeyOrID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.CustomerRequestTransitionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i internalServiceRequestImpl) Transition(ctx context.Context, issueKeyOrID, transitionID, comment string) (*model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	if transitionID == "" {
		return nil, model.ErrNoTransitionIDError
	}

	payload := struct {
		ID                string `json:"id"`
		AdditionalComment struct {
			Body string `json:"body,omitempty"`
		} `json:"additionalComment,omitempty"`
	}{
		ID: transitionID,
		AdditionalComment: struct {
			Body string `json:"body,omitempty"`
		}{
			Body: comment,
		},
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/transition", issueKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
