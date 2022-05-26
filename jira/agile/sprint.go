package agile

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type SprintService struct{ client *Client }

// Get Returns the sprint for a given sprint ID.
// The sprint will only be returned if the user can view the board that the sprint was created on,
// or view at least one of the issues in the sprint.
// Docs: https://docs.go-atlassian.io/jira-agile/sprints#get-sprint
func (s *SprintService) Get(ctx context.Context, sprintID int) (result *models.SprintScheme, response *models.ResponseScheme, err error) {

	if sprintID == 0 {
		return nil, nil, models.ErrNoSprintIDError
	}

	var endpoint = fmt.Sprintf("rest/agile/1.0/sprint/%v", sprintID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Create creates a future sprint.
// Sprint name and origin board id are required.
// Start date, end date, and goal are optional.
// Docs: https://docs.go-atlassian.io/jira-agile/sprints#create-print
func (s *SprintService) Create(ctx context.Context, payload *models.SprintPayloadScheme) (result *models.SprintScheme,
	response *models.ResponseScheme, err error) {

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint = "rest/agile/1.0/sprint"

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Update Performs a full update of a sprint.
// A full update means that the result will be exactly the same as the request body.
// Any fields not present in the request JSON will be set to null.
// Docs: https://docs.go-atlassian.io/jira-agile/sprints#update-sprint
func (s *SprintService) Update(ctx context.Context, sprintID int, payload *models.SprintPayloadScheme) (result *models.SprintScheme,
	response *models.ResponseScheme, err error) {

	if sprintID == 0 {
		return nil, nil, models.ErrNoSprintIDError
	}

	var endpoint = fmt.Sprintf("rest/agile/1.0/sprint/%v", sprintID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := s.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Path Performs a partial update of a sprint.
// A partial update means that fields not present in the request JSON will not be updated.
// Docs: https://docs.go-atlassian.io/jira-agile/sprints#partially-update-sprint
func (s *SprintService) Path(ctx context.Context, sprintID int, payload *models.SprintPayloadScheme) (result *models.SprintScheme,
	response *models.ResponseScheme, err error) {

	if sprintID == 0 {
		return nil, nil, models.ErrNoSprintIDError
	}

	var endpoint = fmt.Sprintf("rest/agile/1.0/sprint/%v", sprintID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Delete deletes a sprint.
// Once a sprint is deleted, all open issues in the sprint will be moved to the backlog.
func (s *SprintService) Delete(ctx context.Context, sprintID int) (response *models.ResponseScheme, err error) {

	if sprintID == 0 {
		return nil, models.ErrNoSprintIDError
	}

	var endpoint = fmt.Sprintf("rest/agile/1.0/sprint/%v", sprintID)

	request, err := s.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = s.client.call(request, nil)
	if err != nil {
		return response, err
	}

	return
}

// Issues returns all issues in a sprint, for a given sprint ID.
// This only includes issues that the user has permission to view.
// By default, the returned issues are ordered by rank.
// Docs: https://docs.go-atlassian.io/jira-agile/sprints#get-issues-for-sprint
func (s *SprintService) Issues(ctx context.Context, sprintID int, opts *models.IssueOptionScheme, startAt, maxResults int) (
	result *models.SprintIssuePageScheme, response *models.ResponseScheme, err error) {

	if sprintID == 0 {
		return nil, nil, models.ErrNoSprintIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if !opts.ValidateQuery {
			params.Add("validateQuery ", "false")
		}

		if len(opts.JQL) != 0 {
			params.Add("jql", opts.JQL)
		}

		if len(opts.Expand) != 0 {
			params.Add("expand", strings.Join(opts.Expand, ","))
		}

		if len(opts.Fields) != 0 {
			params.Add("fields", strings.Join(opts.Fields, ","))
		}
	}

	var endpoint = fmt.Sprintf("rest/agile/1.0/sprint/%v/issue?%v", sprintID, params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = s.client.call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Start initiate the Sprint
// Docs: https://docs.go-atlassian.io/jira-agile/sprints#start-sprint
func (s *SprintService) Start(ctx context.Context, sprintID int) (response *models.ResponseScheme, err error) {

	if sprintID == 0 {
		return nil, models.ErrNoSprintIDError
	}

	payload := models.SprintPayloadScheme{
		State: "Active",
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = fmt.Sprintf("rest/agile/1.0/sprint/%v", sprintID)

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.call(request, nil)
	if err != nil {
		return response, err
	}

	return
}

// Close closes the Sprint
// Docs: https://docs.go-atlassian.io/jira-agile/sprints#close-sprint
func (s *SprintService) Close(ctx context.Context, sprintID int) (response *models.ResponseScheme, err error) {

	if sprintID == 0 {
		return nil, models.ErrNoSprintIDError
	}

	payload := models.SprintPayloadScheme{
		State: "Closed",
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = fmt.Sprintf("rest/agile/1.0/sprint/%v", sprintID)

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.call(request, nil)
	if err != nil {
		return response, err
	}

	return
}
