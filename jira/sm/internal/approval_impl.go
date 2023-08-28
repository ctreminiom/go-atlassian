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
)

func NewApprovalService(client service.Connector, version string) *ApprovalService {

	return &ApprovalService{
		internalClient: &internalServiceRequestApprovalImpl{c: client, version: version},
	}
}

type ApprovalService struct {
	internalClient sm.ApprovalConnector
}

// Gets returns all approvals on a customer request.
//
// GET /rest/servicedeskapi/request/{issueIdOrKey}/approval
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/approval#get-approvals
func (s *ApprovalService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.CustomerApprovalPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Gets(ctx, issueKeyOrID, start, limit)
}

// Get returns an approval. Use this method to determine the status of an approval and the list of approvers.
//
// GET /rest/servicedeskapi/request/{issueIdOrKey}/approval/{approvalId}
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/approval#get-approval-by-id
func (s *ApprovalService) Get(ctx context.Context, issueKeyOrID string, approvalID int) (*model.CustomerApprovalScheme, *model.ResponseScheme, error) {
	return s.internalClient.Get(ctx, issueKeyOrID, approvalID)
}

// Answer enables a user to Approve or Decline an approval on a customer request.
//
// The approval is assumed to be owned by the user making the call.
//
// POST /rest/servicedeskapi/request/{issueIdOrKey}/approval/{approvalId}
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
		return nil, nil, model.ErrNoIssueKeyOrIDError
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
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if approvalID == 0 {
		return nil, nil, model.ErrNoApprovalIDError
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
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if approvalID == 0 {
		return nil, nil, model.ErrNoApprovalIDError
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
