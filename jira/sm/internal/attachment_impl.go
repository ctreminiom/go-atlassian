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

// NewAttachmentService creates a new instance of AttachmentService.
// It takes a service.Connector and a version string as input and returns a pointer to AttachmentService.
func NewAttachmentService(client service.Connector, version string) *AttachmentService {
	return &AttachmentService{
		internalClient: &internalServiceRequestAttachmentImpl{c: client, version: version},
	}
}

// AttachmentService provides methods to interact with attachment operations in Jira Service Management.
type AttachmentService struct {
	// internalClient is the connector interface for attachment operations.
	internalClient sm.AttachmentConnector
}

// Gets returns all the attachments for a customer requests.
//
// GET /rest/servicedeskapi/request/{issueKeyOrID}/attachment
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/attachment#get-attachments-for-request
func (s *AttachmentService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.RequestAttachmentPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Gets(ctx, issueKeyOrID, start, limit)
}

// Create adds one or more temporary files (attached to the request's service desk using
//
// servicedesk/{serviceDeskId}/attachTemporaryFile) as attachments to a customer request
//
// POST /rest/servicedeskapi/request/{issueKeyOrID}/attachment
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/attachment#create-attachment
func (s *AttachmentService) Create(ctx context.Context, issueKeyOrID string, payload *model.RequestAttachmentCreationPayloadScheme) (*model.RequestAttachmentCreationScheme, *model.ResponseScheme, error) {
	return s.internalClient.Create(ctx, issueKeyOrID, payload)
}

type internalServiceRequestAttachmentImpl struct {
	c       service.Connector
	version string
}

func (i *internalServiceRequestAttachmentImpl) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.RequestAttachmentPageScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, fmt.Errorf("sm: %w", model.ErrNoIssueKeyOrID)
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	url := fmt.Sprintf("rest/servicedeskapi/request/%v/attachment?%v", issueKeyOrID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.RequestAttachmentPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalServiceRequestAttachmentImpl) Create(ctx context.Context, issueKeyOrID string, payload *model.RequestAttachmentCreationPayloadScheme) (*model.RequestAttachmentCreationScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, fmt.Errorf("sm: %w", model.ErrNoIssueKeyOrID)
	}

	if len(payload.TemporaryAttachmentIDs) == 0 {
		return nil, nil, fmt.Errorf("sm: %w", model.ErrNoAttachmentID)
	}

	url := fmt.Sprintf("rest/servicedeskapi/request/%v/attachment", issueKeyOrID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, url, "", payload)
	if err != nil {
		return nil, nil, err
	}

	attachment := new(model.RequestAttachmentCreationScheme)
	res, err := i.c.Call(req, attachment)
	if err != nil {
		return nil, res, err
	}

	return attachment, res, nil
}
