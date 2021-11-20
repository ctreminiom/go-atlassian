package sm

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type RequestApprovalService struct{ client *Client }

// Gets returns all approvals on a customer request.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/approval#get-approvals
func (r *RequestApprovalService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (result *model.CustomerApprovalPageScheme,
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
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
func (r *RequestApprovalService) Get(ctx context.Context, issueKeyOrID string, approvalID int) (result *model.CustomerApprovalScheme,
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
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
	result *model.CustomerApprovalScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
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
