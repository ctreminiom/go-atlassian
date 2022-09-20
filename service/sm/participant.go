package sm

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type ParticipantConnector interface {

	// Gets returns a list of all the participants on a customer request.
	//
	// GET /rest/servicedeskapi/request/{issueIdOrKey}/participant
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/participants#get-request-participants
	Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.RequestParticipantPageScheme, *model.ResponseScheme, error)

	// Add adds participants to a customer request.
	//
	// POST /rest/servicedeskapi/request/{issueIdOrKey}/participant
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/participants#add-request-participants
	Add(ctx context.Context, issueKeyOrID string, accountIDs []string) (*model.RequestParticipantPageScheme, *model.ResponseScheme, error)

	// Remove removes participants from a customer request.
	//
	// DELETE /rest/servicedeskapi/request/{issueIdOrKey}/participant
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/participants#remove-request-participants
	Remove(ctx context.Context, issueKeyOrID string, accountIDs []string) (*model.RequestParticipantPageScheme, *model.ResponseScheme, error)
}
