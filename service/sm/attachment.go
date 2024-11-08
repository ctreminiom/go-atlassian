package sm

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type AttachmentConnector interface {

	// Gets returns all the attachments for a customer requests.
	//
	// GET /rest/servicedeskapi/request/{issueKeyOrID}/attachment
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/attachment#get-attachments-for-request
	Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.RequestAttachmentPageScheme, *model.ResponseScheme, error)

	// Create adds one or more temporary files (attached to the request's service desk using
	//
	// servicedesk/{serviceDeskID}/attachTemporaryFile) as attachments to a customer request
	//
	// POST /rest/servicedeskapi/request/{issueKeyOrID}/attachment
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/attachment#create-attachment
	Create(ctx context.Context, issueKeyOrID string, payload *model.RequestAttachmentCreationPayloadScheme) (*model.RequestAttachmentCreationScheme, *model.ResponseScheme, error)
}
