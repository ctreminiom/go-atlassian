package sm

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type RequestConnector interface {

	// Create creates a customer request in a service desk.
	//
	// The JSON request must include the service desk and customer request type, as well as any fields that are required for the request type.
	//
	// POST /rest/servicedeskapi/request
	//
	// https://docs.go-atlassian.io/jira-service-management/request#create-customer-request
	Create(ctx context.Context, payload *model.CreateCustomerRequestPayloadScheme) (*model.CustomerRequestScheme, *model.ResponseScheme, error)

	// Gets returns all customer requests for the user executing the query.
	//
	// The returned customer requests are ordered chronologically by the latest activity on each request. For example, the latest status transition or comment.
	//
	// GET /rest/servicedeskapi/request
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request#get-customer-requests
	Gets(ctx context.Context, options *model.ServiceRequestOptionScheme, start, limit int) (*model.CustomerRequestPageScheme, *model.ResponseScheme, error)

	// Get returns a customer request.
	//
	// GET /rest/servicedeskapi/request/{issueKeyOrID}
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request#get-customer-request-by-id-or-key
	Get(ctx context.Context, issueKeyOrID string, expand []string) (*model.CustomerRequestScheme, *model.ResponseScheme, error)

	// Subscribe subscribes the user to receiving notifications from a customer request.
	//
	// PUT /rest/servicedeskapi/request/{issueKeyOrID}/notification
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request#subscribe
	Subscribe(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error)

	// Unsubscribe unsubscribes the user from notifications from a customer request.
	//
	// DELETE /rest/servicedeskapi/request/{issueKeyOrID}/notification
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request#unsubscribe
	Unsubscribe(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error)

	// Transitions returns a list of transitions, the workflow processes that moves a customer request from one status to another, that the user can perform on a request.
	//
	// GET /rest/servicedeskapi/request/{issueKeyOrID}/transition
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request#get-customer-transitions
	Transitions(ctx context.Context, issueKeyOrID string, start, limit int) (*model.CustomerRequestTransitionPageScheme, *model.ResponseScheme, error)

	// Transition performs a customer transition for a given request and transition.
	//
	// An optional comment can be included to provide a reason for the transition.
	//
	// POST /rest/servicedeskapi/request/{issueKeyOrID}/transition
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request#perform-customer-transition
	Transition(ctx context.Context, issueKeyOrID, transitionID, comment string) (*model.ResponseScheme, error)
}
