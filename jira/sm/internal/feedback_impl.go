package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/sm"
	"net/http"
)

func NewFeedbackService(client service.Client, version string) (*FeedbackService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &FeedbackService{
		internalClient: &internalServiceRequestFeedbackImpl{c: client, version: version},
	}, nil
}

type FeedbackService struct {
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
	c       service.Client
	version string
}

func (i *internalServiceRequestFeedbackImpl) Get(ctx context.Context, requestIDOrKey string) (*model.CustomerFeedbackScheme, *model.ResponseScheme, error) {

	if requestIDOrKey == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/feedback", requestIDOrKey)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	feedback := new(model.CustomerFeedbackScheme)
	response, err := i.c.Call(request, feedback)
	if err != nil {
		return nil, response, err
	}

	return feedback, response, nil
}

func (i *internalServiceRequestFeedbackImpl) Post(ctx context.Context, requestIDOrKey string, rating int, comment string) (*model.CustomerFeedbackScheme, *model.ResponseScheme, error) {

	if requestIDOrKey == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	payload := struct {
		Rating  int `json:"rating,omitempty"`
		Comment struct {
			Body string `json:"body,omitempty"`
		} `json:"comment,omitempty"`
		Type string `json:"type,omitempty"`
	}{
		Rating: rating,
		Comment: struct {
			Body string `json:"body,omitempty"`
		}{
			Body: comment,
		},
		Type: "csat",
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/feedback", requestIDOrKey)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	feedback := new(model.CustomerFeedbackScheme)
	response, err := i.c.Call(request, feedback)
	if err != nil {
		return nil, response, err
	}

	return feedback, response, nil
}

func (i *internalServiceRequestFeedbackImpl) Delete(ctx context.Context, requestIDOrKey string) (*model.ResponseScheme, error) {

	if requestIDOrKey == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/feedback", requestIDOrKey)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
