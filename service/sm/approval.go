package sm

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type ApprovalConnector interface {

	// Gets returns all approvals on a customer request.
	//
	// GET /rest/servicedeskapi/request/{issueKeyOrID}/approval
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/approval#get-approvals
	Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.CustomerApprovalPageScheme, *model.ResponseScheme, error)

	// Get returns an approval. Use this method to determine the status of an approval and the list of approvers.
	//
	// GET /rest/servicedeskapi/request/{issueKeyOrID}/approval/{approvalID}
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/approval#get-approval-by-id
	Get(ctx context.Context, issueKeyOrID string, approvalID int) (*model.CustomerApprovalScheme, *model.ResponseScheme, error)

	// Answer enables a user to Approve or Decline an approval on a customer request.
	//
	// The approval is assumed to be owned by the user making the call.
	//
	// POST /rest/servicedeskapi/request/{issueKeyOrID}/approval/{approvalID}
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/approval#answer-approval
	Answer(ctx context.Context, issueKeyOrID string, approvalID int, approve bool) (*model.CustomerApprovalScheme, *model.ResponseScheme, error)
}
