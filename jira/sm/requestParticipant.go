package sm

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type RequestParticipantService struct{ client *Client }

// Gets returns a list of all the participants on a customer request.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/participants#get-request-participants
func (r *RequestParticipantService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (
	result *RequestParticipantPageScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, notIssueError
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
	result *RequestParticipantPageScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, notIssueError
	}

	if len(accountIDs) == 0 {
		return nil, nil, notAccountsError
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
	result *RequestParticipantPageScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, notIssueError
	}

	if len(accountIDs) == 0 {
		return nil, nil, notAccountsError
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

type RequestParticipantPageScheme struct {
	Size       int                               `json:"size,omitempty"`
	Start      int                               `json:"start,omitempty"`
	Limit      int                               `json:"limit,omitempty"`
	IsLastPage bool                              `json:"isLastPage,omitempty"`
	Values     []*RequestParticipantScheme       `json:"values,omitempty"`
	Expands    []string                          `json:"_expands,omitempty"`
	Links      *RequestParticipantPageLinkScheme `json:"_links,omitempty"`
}

type RequestParticipantPageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type RequestParticipantScheme struct {
	AccountID    string                        `json:"accountId,omitempty"`
	Name         string                        `json:"name,omitempty"`
	Key          string                        `json:"key,omitempty"`
	EmailAddress string                        `json:"emailAddress,omitempty"`
	DisplayName  string                        `json:"displayName,omitempty"`
	Active       bool                          `json:"active,omitempty"`
	TimeZone     string                        `json:"timeZone,omitempty"`
	Links        *RequestParticipantLinkScheme `json:"_links,omitempty"`
}

type RequestParticipantLinkScheme struct {
	Self     string `json:"self,omitempty"`
	JiraRest string `json:"jiraRest,omitempty"`
}

var (
	notAccountsError = fmt.Errorf("error, please provide a valid accountIDs slice value")
)
