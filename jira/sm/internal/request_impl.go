package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/sm"
)

// ServiceRequestSubServices holds the sub-services related to service requests in Jira Service Management.
type ServiceRequestSubServices struct {
	// Approval handles approval operations.
	Approval *ApprovalService
	// Attachment handles attachment operations.
	Attachment *AttachmentService
	// Comment handles comment operations.
	Comment *CommentService
	// Participant handles participant operations.
	Participant *ParticipantService
	// SLA handles service level agreement operations.
	SLA *ServiceLevelAgreementService
	// Feedback handles feedback operations.
	Feedback *FeedbackService
	// Type handles request type operations.
	Type *TypeService
}

// NewRequestService creates a new instance of RequestService.
// It takes a service.Connector, a version string, and an optional ServiceRequestSubServices as input.
// Returns a pointer to RequestService and an error if the version is not provided.
func NewRequestService(client service.Connector, version string, subServices *ServiceRequestSubServices) (*RequestService, error) {

	if version == "" {
		return nil, fmt.Errorf("client: %w", model.ErrNoVersionProvided)
	}

	requestService := &RequestService{
		internalClient: &internalServiceRequestImpl{c: client, version: version},
	}

	if subServices != nil {
		requestService.Approval = subServices.Approval
		requestService.Attachment = subServices.Attachment
		requestService.Comment = subServices.Comment
		requestService.Participant = subServices.Participant
		requestService.SLA = subServices.SLA
		requestService.Feedback = subServices.Feedback
		requestService.Type = subServices.Type
	}

	return requestService, nil
}

// RequestService provides methods to interact with service request operations in Jira Service Management.
type RequestService struct {
	// internalClient is the connector interface for service request operations.
	internalClient sm.RequestConnector
	// Approval handles approval operations.
	Approval *ApprovalService
	// Attachment handles attachment operations.
	Attachment *AttachmentService
	// Comment handles comment operations.
	Comment *CommentService
	// Participant handles participant operations.
	Participant *ParticipantService
	// SLA handles service level agreement operations.
	SLA *ServiceLevelAgreementService
	// Feedback handles feedback operations.
	Feedback *FeedbackService
	// Type handles request type operations.
	Type *TypeService
}

// Create creates a customer request in a service desk.
//
// The JSON request must include the service desk and customer request type, as well as any fields that are required for the request type.
//
// POST /rest/servicedeskapi/request
//
// https://docs.go-atlassian.io/jira-service-management/request#create-customer-request
func (s *RequestService) Create(ctx context.Context, payload *model.CreateCustomerRequestPayloadScheme) (*model.CustomerRequestScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*RequestService).Create")
	defer span.End()

	return s.internalClient.Create(ctx, payload)
}

// Gets returns all customer requests for the user executing the query.
//
// The returned customer requests are ordered chronologically by the latest activity on each request. For example, the latest status transition or comment.
//
// GET /rest/servicedeskapi/request
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request#get-customer-requests
func (s *RequestService) Gets(ctx context.Context, options *model.ServiceRequestOptionScheme, start, limit int) (*model.CustomerRequestPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*RequestService).Gets")
	defer span.End()

	return s.internalClient.Gets(ctx, options, start, limit)
}

// Get returns a customer request.
//
// GET /rest/servicedeskapi/request/{issueKeyOrID}
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request#get-customer-request-by-id-or-key
func (s *RequestService) Get(ctx context.Context, issueKeyOrID string, expand []string) (*model.CustomerRequestScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*RequestService).Get")
	defer span.End()

	return s.internalClient.Get(ctx, issueKeyOrID, expand)
}

// Subscribe subscribes the user to receiving notifications from a customer request.
//
// PUT /rest/servicedeskapi/request/{issueKeyOrID}/notification
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request#subscribe
func (s *RequestService) Subscribe(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*RequestService).Subscribe")
	defer span.End()

	return s.internalClient.Subscribe(ctx, issueKeyOrID)
}

// Unsubscribe unsubscribes the user from notifications from a customer request.
//
// DELETE /rest/servicedeskapi/request/{issueKeyOrID}/notification
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request#unsubscribe
func (s *RequestService) Unsubscribe(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*RequestService).Unsubscribe")
	defer span.End()

	return s.internalClient.Unsubscribe(ctx, issueKeyOrID)
}

// Transitions returns a list of transitions, the workflow processes that moves a customer request from one status to another, that the user can perform on a request.
//
// GET /rest/servicedeskapi/request/{issueKeyOrID}/transition
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request#get-customer-transitions
func (s *RequestService) Transitions(ctx context.Context, issueKeyOrID string, start, limit int) (*model.CustomerRequestTransitionPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*RequestService).Transitions")
	defer span.End()

	return s.internalClient.Transitions(ctx, issueKeyOrID, start, limit)
}

// Transition performs a customer transition for a given request and transition.
//
// An optional comment can be included to provide a reason for the transition.
//
// POST /rest/servicedeskapi/request/{issueKeyOrID}/transition
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request#perform-customer-transition
func (s *RequestService) Transition(ctx context.Context, issueKeyOrID, transitionID, comment string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*RequestService).Transition")
	defer span.End()

	return s.internalClient.Transition(ctx, issueKeyOrID, transitionID, comment)
}

type internalServiceRequestImpl struct {
	c       service.Connector
	version string
}

func (i *internalServiceRequestImpl) Create(ctx context.Context, payload *model.CreateCustomerRequestPayloadScheme) (*model.CustomerRequestScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalServiceRequestImpl).Create")
	defer span.End()

	endpoint := "rest/servicedeskapi/request"

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	serviceRequest := new(model.CustomerRequestScheme)
	res, err := i.c.Call(req, serviceRequest)
	if err != nil {
		return nil, res, err
	}

	return serviceRequest, res, nil
}

func (i *internalServiceRequestImpl) Gets(ctx context.Context, options *model.ServiceRequestOptionScheme, start, limit int) (*model.CustomerRequestPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalServiceRequestImpl).Gets")
	defer span.End()

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

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.CustomerRequestPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalServiceRequestImpl) Get(ctx context.Context, issueKeyOrID string, expand []string) (*model.CustomerRequestScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalServiceRequestImpl).Get")
	defer span.End()

	if issueKeyOrID == "" {
		return nil, nil, fmt.Errorf("sm: %w", model.ErrNoIssueKeyOrID)
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/servicedeskapi/request/%v", issueKeyOrID))

	if expand != nil {
		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	customerRequest := new(model.CustomerRequestScheme)
	res, err := i.c.Call(req, customerRequest)
	if err != nil {
		return nil, res, err
	}

	return customerRequest, res, nil
}

func (i *internalServiceRequestImpl) Subscribe(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalServiceRequestImpl).Subscribe")
	defer span.End()

	if issueKeyOrID == "" {
		return nil, fmt.Errorf("sm: %w", model.ErrNoIssueKeyOrID)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/notification", issueKeyOrID)

	req, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}

func (i *internalServiceRequestImpl) Unsubscribe(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalServiceRequestImpl).Unsubscribe")
	defer span.End()

	if issueKeyOrID == "" {
		return nil, fmt.Errorf("sm: %w", model.ErrNoIssueKeyOrID)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/notification", issueKeyOrID)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}

func (i *internalServiceRequestImpl) Transitions(ctx context.Context, issueKeyOrID string, start, limit int) (*model.CustomerRequestTransitionPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalServiceRequestImpl).Transitions")
	defer span.End()

	if issueKeyOrID == "" {
		return nil, nil, fmt.Errorf("sm: %w", model.ErrNoIssueKeyOrID)
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/transition?%v", issueKeyOrID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.CustomerRequestTransitionPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i internalServiceRequestImpl) Transition(ctx context.Context, issueKeyOrID, transitionID, comment string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(internalServiceRequestImpl).Transition")
	defer span.End()

	if issueKeyOrID == "" {
		return nil, fmt.Errorf("sm: %w", model.ErrNoIssueKeyOrID)
	}

	if transitionID == "" {
		return nil, fmt.Errorf("sm: %w", model.ErrNoTransitionID)
	}

	payload := map[string]interface{}{"id": transitionID}

	if comment != "" {
		payload["additionalComment"] = map[string]interface{}{"body": comment}
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/transition", issueKeyOrID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}
