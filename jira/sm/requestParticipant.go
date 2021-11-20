package sm

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type RequestParticipantService struct{ client *Client }

// Gets returns a list of all the participants on a customer request.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/participants#get-request-participants
func (r *RequestParticipantService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (
	result *model.RequestParticipantPageScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/participant?%v", issueKeyOrID, params.Encode())

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Add adds participants to a customer request.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/participants#add-request-participants
func (r *RequestParticipantService) Add(ctx context.Context, issueKeyOrID string, accountIDs []string) (
	result *model.RequestParticipantPageScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if len(accountIDs) == 0 {
		return nil, nil, model.ErrNoAccountSliceError
	}

	payload := struct {
		AccountIds []string `json:"accountIds"`
	}{
		AccountIds: accountIDs,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/participant", issueKeyOrID)

	request, err := r.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Remove removes participants from a customer request.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/participants#remove-request-participants
func (r *RequestParticipantService) Remove(ctx context.Context, issueKeyOrID string, accountIDs []string) (
	result *model.RequestParticipantPageScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	if len(accountIDs) == 0 {
		return nil, nil, model.ErrNoAccountSliceError
	}

	payload := struct {
		AccountIds []string `json:"accountIds"`
	}{
		AccountIds: accountIDs,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/participant", issueKeyOrID)

	request, err := r.client.newRequest(ctx, http.MethodDelete, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}
