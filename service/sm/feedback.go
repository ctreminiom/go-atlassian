package sm

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type FeedbackConnector interface {

	// Get retrieves a feedback of a request using it's request key or request id
	//
	// GET /rest/servicedeskapi/request/{requestIDOrKey}/feedback
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/feedback#get-feedback
	Get(ctx context.Context, requestIDOrKey string) (*model.CustomerFeedbackScheme, *model.ResponseScheme, error)

	// Post adds a feedback on a request using its request key or request id
	//
	// POST /rest/servicedeskapi/request/{requestIDOrKey}/feedback
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/feedback#post-feedback
	Post(ctx context.Context, requestIDOrKey string, rating int, comment string) (*model.CustomerFeedbackScheme, *model.ResponseScheme, error)

	// Delete deletes the feedback of request using its request key or request id
	//
	// DELETE /rest/servicedeskapi/request/{requestIDOrKey}/feedback
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/feedback#delete-feedback
	Delete(ctx context.Context, requestIDOrKey string) (*model.ResponseScheme, error)
}
