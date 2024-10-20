package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/sm"
)

// NewApprovalService creates a new instance of ApprovalService.
// It takes a service.Connector and a version string as input and returns a pointer to ApprovalService.
func NewApprovalService(client service.Connector, version string) *ApprovalService {
	return &ApprovalService{
		internalClient: &internalServiceRequestApprovalImpl{c: client, version: version},
	}
}

// ApprovalService provides methods to interact with approval operations in Jira Service Management.
type ApprovalService struct {
	// internalClient is the connector interface for approval operations.
	internalClient sm.ApprovalConnector
}

// Gets returns all approvals on a customer request.
//
// GET /rest/servicedeskapi/request/{issueKeyOrID}/approval
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/approval#get-approvals
func (s *ApprovalService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.CustomerApprovalPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Gets(ctx, issueKeyOrID, start, limit)
}

// Get returns an approval. Use this method to determine the status of an approval and the list of approvers.
//
// GET /rest/servicedeskapi/request/{issueKeyOrID}/approval/{approvalId}
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/approval#get-approval-by-id
func (s *ApprovalService) Get(ctx context.Context, issueKeyOrID string, approvalID int) (*model.CustomerApprovalScheme, *model.ResponseScheme, error) {
	return s.internalClient.Get(ctx, issueKeyOrID, approvalID)
}

// Answer enables a user to Approve or Decline an approval on a customer request.
//
// The approval is assumed to be owned by the user making the call.
//
// POST /rest/servicedeskapi/request/{issueKeyOrID}/approval/{approvalId}
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/approval#answer-approval
func (s *ApprovalService) Answer(ctx context.Context, issueKeyOrID string, approvalID int, approve bool) (*model.CustomerApprovalScheme, *model.ResponseScheme, error) {
	return s.internalClient.Answer(ctx, issueKeyOrID, approvalID, approve)
}

type internalServiceRequestApprovalImpl struct {
	c       service.Connector
	version string
}

func (i *internalServiceRequestApprovalImpl) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.CustomerApprovalPageScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	url := fmt.Sprintf("rest/servicedeskapi/request/%v/approval?%v", issueKeyOrID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.CustomerApprovalPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalServiceRequestApprovalImpl) Get(ctx context.Context, issueKeyOrID string, approvalID int) (*model.CustomerApprovalScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	if approvalID == 0 {
		return nil, nil, model.ErrNoApprovalID
	}

	url := fmt.Sprintf("rest/servicedeskapi/request/%v/approval/%v", issueKeyOrID, approvalID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	approval := new(model.CustomerApprovalScheme)
	res, err := i.c.Call(req, approval)
	if err != nil {
		return nil, res, err
	}

	return approval, res, nil
}

func (i *internalServiceRequestApprovalImpl) Answer(ctx context.Context, issueKeyOrID string, approvalID int, approve bool) (*model.CustomerApprovalScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	if approvalID == 0 {
		return nil, nil, model.ErrNoApprovalID
	}

	url := fmt.Sprintf("rest/servicedeskapi/request/%v/approval/%v", issueKeyOrID, approvalID)

	payload := make(map[string]interface{})

	if approve {
		payload["decision"] = "approve"
	} else {
		payload["decision"] = "decline"
	}

	req, err := i.c.NewRequest(ctx, http.MethodPost, url, "", payload)
	if err != nil {
		return nil, nil, err
	}

	approval := new(model.CustomerApprovalScheme)
	res, err := i.c.Call(req, approval)
	if err != nil {
		return nil, res, err
	}

	return approval, res, nil
}
