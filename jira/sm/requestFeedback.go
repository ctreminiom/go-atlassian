package sm

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
)

type RequestFeedbackService struct{ client *Client }

// Get retrieves a feedback of a request using it's requestKey or requestId
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/feedback#get-feedback
func (r *RequestFeedbackService) Get(ctx context.Context, requestIDOrKey string) (result *model.CustomerFeedbackScheme,
	response *ResponseScheme, err error) {

	if len(requestIDOrKey) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/feedback", requestIDOrKey)

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("X-ExperimentalApi", "opt-in")

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Post adds a feedback on a request using its requestKey or requestId
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/feedback#post-feedback
func (r *RequestFeedbackService) Post(ctx context.Context, requestIDOrKey string, rating int, comment string) (
	result *model.CustomerFeedbackScheme, response *ResponseScheme, err error) {

	if len(requestIDOrKey) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	payload := struct {
		Rating  int `json:"rating"`
		Comment struct {
			Body string `json:"body,omitempty"`
		} `json:"comment,omitempty"`
		Type string `json:"type"`
	}{
		Rating: rating,
		Comment: struct {
			Body string `json:"body,omitempty"`
		}{
			Body: comment,
		},
		Type: "csat",
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/feedback", requestIDOrKey)

	request, err := r.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-ExperimentalApi", "opt-in")

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes the feedback of request using its requestKey or requestId
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/feedback#delete-feedback
func (r *RequestFeedbackService) Delete(ctx context.Context, requestIDOrKey string) (response *ResponseScheme, err error) {

	if len(requestIDOrKey) == 0 {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/feedback", requestIDOrKey)

	request, err := r.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("X-ExperimentalApi", "opt-in")

	response, err = r.client.Call(request, nil)
	if err != nil {
		return
	}

	return
}
