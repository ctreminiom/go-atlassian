package agile

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type SprintService struct{ client *Client }

// Get Returns the sprint for a given sprint ID.
// The sprint will only be returned if the user can view the board that the sprint was created on,
// or view at least one of the issues in the sprint.
// Docs: N/A
func (s *SprintService) Get(ctx context.Context, sprintID int) (result *SprintScheme, response *Response, err error) {

	if sprintID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid sprint ID")
	}

	var endpoint = fmt.Sprintf("rest/agile/1.0/sprint/%v", sprintID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(SprintScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type SprintScheme struct {
	ID            int       `json:"id,omitempty"`
	Self          string    `json:"self,omitempty"`
	State         string    `json:"state,omitempty"`
	Name          string    `json:"name,omitempty"`
	StartDate     time.Time `json:"startDate,omitempty"`
	EndDate       time.Time `json:"endDate,omitempty"`
	CompleteDate  time.Time `json:"completeDate,omitempty"`
	OriginBoardID int       `json:"originBoardId,omitempty"`
	Goal          string    `json:"goal,omitempty"`
}

// Create creates a future sprint.
// Sprint name and origin board id are required.
// Start date, end date, and goal are optional.
// Docs: N/A
func (s *SprintService) Create(ctx context.Context, payload *SprintPayloadScheme) (result *SprintScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error!, please provide a valid SprintPayloadScheme pointer")
	}

	var endpoint = "rest/agile/1.0/sprint"

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(SprintScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type SprintPayloadScheme struct {
	Name          string `json:"name,omitempty"`
	StartDate     string `json:"startDate,omitempty"`
	EndDate       string `json:"endDate,omitempty"`
	OriginBoardID int    `json:"originBoardId,omitempty"`
	Goal          string `json:"goal,omitempty"`
	State         string `json:"state,omitempty"`
}

// Update Performs a full update of a sprint.
// A full update means that the result will be exactly the same as the request body.
// Any fields not present in the request JSON will be set to null.
// Docs: N/A
func (s *SprintService) Update(ctx context.Context, sprintID int, payload *SprintPayloadScheme) (result *SprintScheme, response *Response, err error) {

	if sprintID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid sprint ID")
	}

	if payload == nil {
		return nil, nil, fmt.Errorf("error!, please provide a valid SprintPayloadScheme pointer")
	}

	var endpoint = fmt.Sprintf("rest/agile/1.0/sprint/%v", sprintID)

	request, err := s.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(SprintScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Path Performs a partial update of a sprint.
// A partial update means that fields not present in the request JSON will not be updated.
// Docs: N/A
func (s *SprintService) Path(ctx context.Context, sprintID int, payload *SprintPayloadScheme) (result *SprintScheme, response *Response, err error) {

	if sprintID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid sprint ID")
	}

	if payload == nil {
		return nil, nil, fmt.Errorf("error!, please provide a valid SprintPayloadScheme pointer")
	}

	var endpoint = fmt.Sprintf("rest/agile/1.0/sprint/%v", sprintID)

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(SprintScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Delete deletes a sprint.
// Once a sprint is deleted, all open issues in the sprint will be moved to the backlog.
// Docs: N/A
func (s *SprintService) Delete(ctx context.Context, sprintID int) (response *Response, err error) {

	if sprintID == 0 {
		return nil, fmt.Errorf("error!, please provide a valid sprint ID")
	}

	var endpoint = fmt.Sprintf("rest/agile/1.0/sprint/%v", sprintID)

	request, err := s.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	return
}

type IssueOptionScheme struct {
	JQL            string
	ValidateQuery  bool
	Fields, Expand []string
}

// Issues returns all issues in a sprint, for a given sprint ID.
// This only includes issues that the user has permission to view.
// By default, the returned issues are ordered by rank.
// Docs: N/A
func (s *SprintService) Issues(ctx context.Context, sprintID int, opts *IssueOptionScheme, startAt, maxResults int) (result *SprintIssuePageScheme, response *Response, err error) {

	if sprintID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid sprint ID")
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

		var expand string
		for index, value := range opts.Expand {

			if index == 0 {
				expand = value
				continue
			}

			expand += "," + value
		}

		if len(expand) != 0 {
			params.Add("expand", expand)
		}

		var fieldFormatted string
		for index, value := range opts.Fields {

			if index == 0 {
				fieldFormatted = value
				continue
			}
			fieldFormatted += "," + value
		}

		if len(fieldFormatted) != 0 {
			params.Add("fields", fieldFormatted)
		}

	}

	var endpoint = fmt.Sprintf("rest/agile/1.0/sprint/%v/issue?%v", sprintID, params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(SprintIssuePageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type SprintIssuePageScheme struct {
	Expand     string               `json:"expand,omitempty"`
	StartAt    int                  `json:"startAt,omitempty"`
	MaxResults int                  `json:"maxResults,omitempty"`
	Total      int                  `json:"total,omitempty"`
	Issues     []*SprintIssueScheme `json:"issues,omitempty"`
}

type SprintIssueScheme struct {
	Expand string `json:"expand,omitempty"`
	ID     string `json:"id,omitempty"`
	Self   string `json:"self,omitempty"`
	Key    string `json:"key,omitempty"`
}

// Start initiate the Sprint
// Docs: N/A

func (s *SprintService) Start(ctx context.Context, sprintID int) (response *Response, err error) {

	if sprintID == 0 {
		return nil, fmt.Errorf("error!, please provide a valid sprint ID")
	}

	payload := SprintPayloadScheme{
		State: "Active",
	}

	var endpoint = fmt.Sprintf("rest/agile/1.0/sprint/%v", sprintID)

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Close closes the Sprint
// Docs: N/A
func (s *SprintService) Close(ctx context.Context, sprintID int) (response *Response, err error) {

	if sprintID == 0 {
		return nil, fmt.Errorf("error!, please provide a valid sprint ID")
	}

	payload := SprintPayloadScheme{
		State: "Closed",
	}

	var endpoint = fmt.Sprintf("rest/agile/1.0/sprint/%v", sprintID)

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	return
}
