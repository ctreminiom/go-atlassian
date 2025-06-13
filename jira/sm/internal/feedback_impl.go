package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/sm"
	"net/http"
)

// NewFeedbackService creates a new instance of FeedbackService.
// It takes a service.Connector and a version string as input and returns a pointer to FeedbackService.
func NewFeedbackService(client service.Connector, version string) *FeedbackService {
	return &FeedbackService{
		internalClient: &internalServiceRequestFeedbackImpl{c: client, version: version},
	}
}

// FeedbackService provides methods to interact with feedback operations in Jira Service Management.
type FeedbackService struct {
	// internalClient is the connector interface for feedback operations.
	internalClient sm.FeedbackConnector
}

// Get retrieves a feedback of a request using it's requestKey or requestId
//
// GET /rest/servicedeskapi/request/{requestIdOrKey}/feedback
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/feedback#get-feedback
func (s *FeedbackService) Get(ctx context.Context, requestIDOrKey string) (*model.CustomerFeedbackScheme, *model.ResponseScheme, error) {
	return s.internalClient.Get(ctx, requestIDOrKey)
}

// Post adds a feedback on a request using its requestKey or requestId
//
// POST /rest/servicedeskapi/request/{requestIdOrKey}/feedback
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/feedback#post-feedback
func (s *FeedbackService) Post(ctx context.Context, requestIDOrKey string, rating int, comment string) (*model.CustomerFeedbackScheme, *model.ResponseScheme, error) {
	return s.internalClient.Post(ctx, requestIDOrKey, rating, comment)
}

// Delete deletes the feedback of request using its requestKey or requestId
//
// DELETE /rest/servicedeskapi/request/{requestIdOrKey}/feedback
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/feedback#delete-feedback
func (s *FeedbackService) Delete(ctx context.Context, requestIDOrKey string) (*model.ResponseScheme, error) {
	return s.internalClient.Delete(ctx, requestIDOrKey)
}

type internalServiceRequestFeedbackImpl struct {
	c       service.Connector
	version string
}

func (i *internalServiceRequestFeedbackImpl) Get(ctx context.Context, requestIDOrKey string) (*model.CustomerFeedbackScheme, *model.ResponseScheme, error) {

	if requestIDOrKey == "" {
		return nil, nil, fmt.Errorf("sm: %w", model.ErrNoIssueKeyOrID)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/feedback", requestIDOrKey)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	feedback := new(model.CustomerFeedbackScheme)
	res, err := i.c.Call(req, feedback)
	if err != nil {
		return nil, res, err
	}

	return feedback, res, nil
}

func (i *internalServiceRequestFeedbackImpl) Post(ctx context.Context, requestIDOrKey string, rating int, comment string) (*model.CustomerFeedbackScheme, *model.ResponseScheme, error) {

	if requestIDOrKey == "" {
		return nil, nil, fmt.Errorf("sm: %w", model.ErrNoIssueKeyOrID)
	}

	payload := map[string]interface{}{
		"rating": rating,
		"type":   "csat",
		"comment": map[string]interface{}{
			"body": comment,
		},
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/feedback", requestIDOrKey)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	feedback := new(model.CustomerFeedbackScheme)
	res, err := i.c.Call(req, feedback)
	if err != nil {
		return nil, res, err
	}

	return feedback, res, nil
}

func (i *internalServiceRequestFeedbackImpl) Delete(ctx context.Context, requestIDOrKey string) (*model.ResponseScheme, error) {

	if requestIDOrKey == "" {
		return nil, fmt.Errorf("sm: %w", model.ErrNoIssueKeyOrID)
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/feedback", requestIDOrKey)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}
