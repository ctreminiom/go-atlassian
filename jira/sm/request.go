package sm

import (
	"context"
	"fmt"
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

// Gets returns all customer requests for the user executing the query.
// The returned customer requests are ordered chronologically by the latest activity on each request. For example, the latest status transition or comment.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request#get-customer-requests
func (r *RequestService) Gets(ctx context.Context, options *RequestGetOptionsScheme, start, limit int) (
	result *CustomerRequestsScheme, response *ResponseScheme, err error) {

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
func (r *RequestService) Get(ctx context.Context, issueKeyOrID string, expand []string) (result *CustomerRequestScheme,
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, notIssueError
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
		return nil, notIssueError
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
		return nil, notIssueError
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
	result *CustomerRequestTransitionPageScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, notIssueError
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
		return nil, notIssueError
	}

	if len(transitionID) == 0 {
		return nil, notTransitionIDError
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

type CustomerRequestTransitionPageScheme struct {
	Size       int                                      `json:"size,omitempty"`
	Start      int                                      `json:"start,omitempty"`
	Limit      int                                      `json:"limit,omitempty"`
	IsLastPage bool                                     `json:"isLastPage,omitempty"`
	Values     []*CustomerRequestTransitionScheme       `json:"values,omitempty"`
	Expands    []string                                 `json:"_expands,omitempty"`
	Links      *CustomerRequestTransitionPageLinkScheme `json:"_links,omitempty"`
}

type CustomerRequestTransitionScheme struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CustomerRequestTransitionPageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type CustomerRequestsScheme struct {
	Size       int                          `json:"size,omitempty"`
	Start      int                          `json:"start,omitempty"`
	Limit      int                          `json:"limit,omitempty"`
	IsLastPage bool                         `json:"isLastPage,omitempty"`
	Values     []*CustomerRequestScheme     `json:"values,omitempty"`
	Expands    []string                     `json:"_expands,omitempty"`
	Links      *CustomerRequestsLinksScheme `json:"_links,omitempty"`
}

type CustomerRequestsLinksScheme struct {
	Self    string `json:"self"`
	Base    string `json:"base"`
	Context string `json:"context"`
	Next    string `json:"next"`
	Prev    string `json:"prev"`
}

type CustomerRequestTypeScheme struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	HelpText      string   `json:"helpText"`
	IssueTypeID   string   `json:"issueTypeId"`
	ServiceDeskID string   `json:"serviceDeskId"`
	GroupIds      []string `json:"groupIds"`
}

type CustomerRequestServiceDeskScheme struct {
	ID          string `json:"id"`
	ProjectID   string `json:"projectId"`
	ProjectName string `json:"projectName"`
	ProjectKey  string `json:"projectKey"`
}

type CustomerRequestDateScheme struct {
	Iso8601     string `json:"iso8601"`
	Jira        string `json:"jira"`
	Friendly    string `json:"friendly"`
	EpochMillis int    `json:"epochMillis"`
}

type CustomerRequestReporterScheme struct {
	AccountID    string `json:"accountId"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	EmailAddress string `json:"emailAddress"`
	DisplayName  string `json:"displayName"`
	Active       bool   `json:"active"`
	TimeZone     string `json:"timeZone"`
}

type CustomerRequestRequestFieldValueScheme struct {
	FieldID string `json:"fieldId"`
	Label   string `json:"label"`
}

type CustomerRequestCurrentStatusScheme struct {
	Status         string `json:"status"`
	StatusCategory string `json:"statusCategory"`
	StatusDate     struct {
	} `json:"statusDate"`
}

type CustomerRequestLinksScheme struct {
	Self     string `json:"self"`
	JiraRest string `json:"jiraRest"`
	Web      string `json:"web"`
	Agent    string `json:"agent"`
}

type CustomerRequestScheme struct {
	IssueID            string                                    `json:"issueId,omitempty"`
	IssueKey           string                                    `json:"issueKey,omitempty"`
	RequestTypeID      string                                    `json:"requestTypeId,omitempty"`
	RequestType        *CustomerRequestTypeScheme                `json:"requestType,omitempty"`
	ServiceDeskID      string                                    `json:"serviceDeskId,omitempty"`
	ServiceDesk        *CustomerRequestServiceDeskScheme         `json:"serviceDesk,omitempty"`
	CreatedDate        *CustomerRequestDateScheme                `json:"createdDate,omitempty"`
	Reporter           *CustomerRequestReporterScheme            `json:"reporter,omitempty"`
	RequestFieldValues []*CustomerRequestRequestFieldValueScheme `json:"requestFieldValues,omitempty"`
	CurrentStatus      *CustomerRequestCurrentStatusScheme       `json:"currentStatus,omitempty"`
	Expands            []string                                  `json:"_expands,omitempty"`
	Links              *CustomerRequestLinksScheme               `json:"_links,omitempty"`
}

var (
	notIssueError        = fmt.Errorf("error, please provide a valid issueKeyOrID value")
	notTransitionIDError = fmt.Errorf("error, please provide a valid transitionID value")
)
