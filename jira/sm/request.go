package sm

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type RequestService struct {
	client      *Client
	Type        *RequestTypeService
	Approval    *RequestApprovalService
	Attachment  *RequestAttachmentService
	Comment     *RequestCommentService
	Participant *RequestParticipantService
	SLA         *RequestSLAService
	Feedback    *RequestFeedbackService
}

type RequestGetOptionsScheme struct {
	SearchTerm        string
	RequestOwnerships []string
	RequestStatus     string
	ApprovalStatus    string
	OrganizationId    int
	ServiceDeskID     int
	RequestTypeID     int
	Expand            []string
}

// Create creates a customer request in a service desk.
// The JSON request must include the service desk and customer request type, as well as any fields that are required for
// the request type.
// Docs: https://docs.go-atlassian.io/jira-service-management/request#create-customer-request
func (r *RequestService) Create(ctx context.Context, payload *model.CreateCustomerRequestPayloadScheme, fields *model.CustomerRequestFields) (
	result *model.CustomerRequestScheme, response *ResponseScheme, err error) {

	if fields == nil {
		return nil, nil, model.ErrNoCustomRequestFieldsError
	}

	payloadWithCustomFields, err := payload.MergeFields(fields)
	if err != nil {
		return nil, nil, err
	}

	reader, err := transformStructToReader(&payloadWithCustomFields)
	if err != nil {
		return nil, nil, err
	}

	endpoint := "rest/servicedeskapi/request"
	request, err := r.client.newRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Gets returns all customer requests for the user executing the query.
// The returned customer requests are ordered chronologically by the latest activity on each request. For example, the latest status transition or comment.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request#get-customer-requests
func (r *RequestService) Gets(ctx context.Context, options *RequestGetOptionsScheme, start, limit int) (
	result *model.CustomerRequestPageScheme, response *ResponseScheme, err error) {

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

		if options.OrganizationId != 0 {
			params.Add("organizationId", strconv.Itoa(options.OrganizationId))
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

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request?%v", params.Encode())

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns a customer request.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request#get-customer-request-by-id-or-key
func (r *RequestService) Get(ctx context.Context, issueKeyOrID string, expand []string) (result *model.CustomerRequestScheme,
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/servicedeskapi/request/%v", issueKeyOrID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Subscribe subscribes the user to receiving notifications from a customer request.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request#subscribe
func (r *RequestService) Subscribe(ctx context.Context, issueKeyOrID string) (response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/notification", issueKeyOrID)

	request, err := r.client.newRequest(ctx, http.MethodPut, endpoint, nil)
	if err != nil {
		return
	}

	response, err = r.client.Call(request, nil)
	if err != nil {
		return
	}

	return
}

// Unsubscribe unsubscribes the user from notifications from a customer request.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request#unsubscribe
func (r *RequestService) Unsubscribe(ctx context.Context, issueKeyOrID string) (response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/notification", issueKeyOrID)

	request, err := r.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = r.client.Call(request, nil)
	if err != nil {
		return
	}

	return
}

// Transitions returns a list of transitions, the workflow processes that moves a customer request from one status to another,
// that the user can perform on a request.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request#get-customer-transitions
func (r *RequestService) Transitions(ctx context.Context, issueKeyOrID string, start, limit int) (
	result *model.CustomerRequestTransitionPageScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/transition?%v", issueKeyOrID, params.Encode())

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Transition performs a customer transition for a given request and transition.
// An optional comment can be included to provide a reason for the transition.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request#perform-customer-transition
func (r *RequestService) Transition(ctx context.Context, issueKeyOrID, transitionID, comment string) (
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	if len(transitionID) == 0 {
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

	payloadAsReader, _ := transformStructToReader(&payload)
	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/transition", issueKeyOrID)

	request, err := r.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = r.client.Call(request, nil)
	if err != nil {
		return
	}

	return
}
