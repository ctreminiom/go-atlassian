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

func NewAttachmentService(client service.Client, version string) (*AttachmentService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &AttachmentService{
		internalClient: &internalServiceRequestAttachmentImpl{c: client, version: version},
	}, nil
}

type AttachmentService struct {
	internalClient sm.AttachmentConnector
}

// Gets returns all the attachments for a customer requests.
//
// GET /rest/servicedeskapi/request/{issueIdOrKey}/attachment
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/attachment#get-attachments-for-request
func (s *AttachmentService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.RequestAttachmentPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Gets(ctx, issueKeyOrID, start, limit)
}

// Create adds one or more temporary files (attached to the request's service desk using
//
// servicedesk/{serviceDeskId}/attachTemporaryFile) as attachments to a customer request
//
// POST /rest/servicedeskapi/request/{issueIdOrKey}/attachment
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/attachment#create-attachment
func (s *AttachmentService) Create(ctx context.Context, issueKeyOrID string, temporaryAttachmentIDs []string, public bool) (*model.RequestAttachmentCreationScheme, *model.ResponseScheme, error) {
	return s.internalClient.Create(ctx, issueKeyOrID, temporaryAttachmentIDs, public)
}

type internalServiceRequestAttachmentImpl struct {
	c       service.Client
	version string
}

func (i *internalServiceRequestAttachmentImpl) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.RequestAttachmentPageScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/attachment?%v", issueKeyOrID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.RequestAttachmentPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalServiceRequestAttachmentImpl) Create(ctx context.Context, issueKeyOrID string, temporaryAttachmentIDs []string, public bool) (*model.RequestAttachmentCreationScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if len(temporaryAttachmentIDs) == 0 {
		return nil, nil, model.ErrNoAttachmentIDError
	}

	payload := struct {
		TemporaryAttachmentIds []string `json:"temporaryAttachmentIds,omitempty"`
		Public                 bool     `json:"public,omitempty"`
	}{
		TemporaryAttachmentIds: temporaryAttachmentIDs,
		Public:                 public,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/attachment", issueKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	attachment := new(model.RequestAttachmentCreationScheme)
	response, err := i.c.Call(request, attachment)
	if err != nil {
		return nil, response, err
	}

	return attachment, response, nil
}
