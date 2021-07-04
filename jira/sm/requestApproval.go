package sm

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type RequestApprovalService struct{ client *Client }

// Gets returns all approvals on a customer request.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/approval#get-approvals
func (r *RequestApprovalService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (result *CustomerApprovalPageScheme,
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, notIssueError
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/approval?%v", issueKeyOrID, params.Encode())

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

// Get returns an approval. Use this method to determine the status of an approval and the list of approvers.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/approval#get-approval-by-id
func (r *RequestApprovalService) Get(ctx context.Context, issueKeyOrID string, approvalID int) (result *CustomerApprovalScheme,
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, notIssueError
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/approval/%v", issueKeyOrID, approvalID)

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

// Answer enables a user to Approve or Decline an approval on a customer request.
// The approval is assumed to be owned by the user making the call.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/approval#answer-approval
func (r *RequestApprovalService) Answer(ctx context.Context, issueKeyOrID string, approvalID int, approve bool) (
	result *CustomerApprovalScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, notIssueError
	}

	var (
		endpoint        = fmt.Sprintf("rest/servicedeskapi/request/%v/approval/%v", issueKeyOrID, approvalID)
		approveAsString string
	)

	if approve {
		approveAsString = "approve"
	} else {
		approveAsString = "decline"
	}

	payload := struct {
		Decision string `json:"decision"`
	}{
		Decision: approveAsString,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	request, err := r.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
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

type CustomerApprovalPageScheme struct {
	Size       int                             `json:"size,omitempty"`
	Start      int                             `json:"start,omitempty"`
	Limit      int                             `json:"limit,omitempty"`
	IsLastPage bool                            `json:"isLastPage,omitempty"`
	Values     []*CustomerApprovalScheme       `json:"values,omitempty"`
	Expands    []string                        `json:"_expands,omitempty"`
	Links      *CustomerApprovalPageLinkScheme `json:"_links,omitempty"`
}

type CustomerApprovalPageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type CustomerApprovalScheme struct {
	ID                string                      `json:"id,omitempty"`
	Name              string                      `json:"name,omitempty"`
	FinalDecision     string                      `json:"finalDecision,omitempty"`
	CanAnswerApproval bool                        `json:"canAnswerApproval,omitempty"`
	Approvers         []*CustomerApproverScheme   `json:"approvers,omitempty"`
	CreatedDate       *CustomerRequestDateScheme  `json:"createdDate,omitempty"`
	CompletedDate     *CustomerRequestDateScheme  `json:"completedDate,omitempty"`
	Links             *CustomerApprovalLinkScheme `json:"_links,omitempty"`
}

type CustomerApprovalLinkScheme struct {
	Self string `json:"self"`
}

type CustomerApproverScheme struct {
	Approver         *ApproverScheme `json:"approver,omitempty"`
	ApproverDecision string          `json:"approverDecision,omitempty"`
}

type ApproverScheme struct {
	AccountID    string              `json:"accountId,omitempty"`
	Name         string              `json:"name,omitempty"`
	Key          string              `json:"key,omitempty"`
	EmailAddress string              `json:"emailAddress,omitempty"`
	DisplayName  string              `json:"displayName,omitempty"`
	Active       bool                `json:"active,omitempty"`
	TimeZone     string              `json:"timeZone,omitempty"`
	Links        *ApproverLinkScheme `json:"_links,omitempty"`
}

type ApproverLinkScheme struct {
	JiraRest   string `json:"jiraRest"`
	AvatarUrls struct {
		Four8X48  string `json:"48x48"`
		Two4X24   string `json:"24x24"`
		One6X16   string `json:"16x16"`
		Three2X32 string `json:"32x32"`
	} `json:"avatarUrls"`
	Self string `json:"self"`
}
