package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/agile"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// NewSprintService creates a new instance of SprintService.
// It takes a service.Connector and a version string as input and returns a pointer to SprintService.
func NewSprintService(client service.Connector, version string) *SprintService {
	return &SprintService{
		internalClient: &internalSprintImpl{c: client, version: version},
	}
}

// SprintService provides methods to interact with sprint operations in Jira Agile.
type SprintService struct {
	// internalClient is the connector interface for sprint operations.
	internalClient agile.SprintConnector
}

// Get Returns the sprint for a given sprint ID.
//
// The sprint will only be returned if the user can view the board that the sprint was created on,
//
// or view at least one of the issues in the sprint.
//
// GET /rest/agile/1.0/sprint/{sprintId}
//
// https://docs.go-atlassian.io/jira-agile/sprints#get-sprint
func (s *SprintService) Get(ctx context.Context, sprintID int) (*model.SprintScheme, *model.ResponseScheme, error) {
	return s.internalClient.Get(ctx, sprintID)
}

// Create creates a future sprint.
//
// Sprint name and origin board id are required.
//
// Start date, end date, and goal are optional.
//
// POST /rest/agile/1.0/sprint
//
// https://docs.go-atlassian.io/jira-agile/sprints#create-print
func (s *SprintService) Create(ctx context.Context, payload *model.SprintPayloadScheme) (*model.SprintScheme, *model.ResponseScheme, error) {
	return s.internalClient.Create(ctx, payload)
}

// Update Performs a full update of a sprint.
//
// A full update means that the result will be exactly the same as the request body.
//
// Any fields not present in the request JSON will be set to null.
//
// PUT /rest/agile/1.0/sprint/{sprintId}
//
// https://docs.go-atlassian.io/jira-agile/sprints#update-sprint
func (s *SprintService) Update(ctx context.Context, sprintID int, payload *model.SprintPayloadScheme) (*model.SprintScheme, *model.ResponseScheme, error) {
	return s.internalClient.Update(ctx, sprintID, payload)
}

// Path Performs a partial update of a sprint.
//
// A partial update means that fields not present in the request JSON will not be updated.
//
// POST /rest/agile/1.0/sprint/{sprintId}
//
// https://docs.go-atlassian.io/jira-agile/sprints#partially-update-sprint
func (s *SprintService) Path(ctx context.Context, sprintID int, payload *model.SprintPayloadScheme) (*model.SprintScheme, *model.ResponseScheme, error) {
	return s.internalClient.Path(ctx, sprintID, payload)
}

// Delete deletes a sprint.
//
// Once a sprint is deleted, all open issues in the sprint will be moved to the backlog.
//
// DELETE /rest/agile/1.0/sprint/{sprintId}
//
// https://docs.go-atlassian.io/jira-agile/sprints#delete-sprint
func (s *SprintService) Delete(ctx context.Context, sprintID int) (*model.ResponseScheme, error) {
	return s.internalClient.Delete(ctx, sprintID)
}

// Issues returns all issues in a sprint, for a given sprint ID.
//
// This only includes issues that the user has permission to view.
//
// By default, the returned issues are ordered by rank.
//
// GET /rest/agile/1.0/sprint/{sprintId}/issue
//
// https://docs.go-atlassian.io/jira-agile/sprints#get-issues-for-sprint
func (s *SprintService) Issues(ctx context.Context, sprintID int, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.SprintIssuePageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Issues(ctx, sprintID, opts, startAt, maxResults)
}

// Start initiate the Sprint
//
// PUT /rest/agile/1.0/sprint/{sprintId}
//
// https://docs.go-atlassian.io/jira-agile/sprints#start-sprint
func (s *SprintService) Start(ctx context.Context, sprintID int) (*model.ResponseScheme, error) {
	return s.internalClient.Start(ctx, sprintID)
}

// Close closes the Sprint
//
// PUT /rest/agile/1.0/sprint/{sprintId}
//
// https://docs.go-atlassian.io/jira-agile/sprints#close-sprint
func (s *SprintService) Close(ctx context.Context, sprintID int) (*model.ResponseScheme, error) {
	return s.internalClient.Close(ctx, sprintID)
}

// Move moves issues to a sprint, for a given sprint ID.
//
// Issues can only be moved to open or active sprints.
//
// The maximum number of issues that can be moved in one operation is 50.
//
// POST /rest/agile/1.0/sprint/{sprintId}/issue
//
// https://docs.go-atlassian.io/jira-agile/sprints#move-issues-to-sprint
func (s *SprintService) Move(ctx context.Context, sprintID int, payload *model.SprintMovePayloadScheme) (*model.ResponseScheme, error) {
	return s.internalClient.Move(ctx, sprintID, payload)
}

type internalSprintImpl struct {
	c       service.Connector
	version string
}

func (i *internalSprintImpl) Move(ctx context.Context, sprintID int, payload *model.SprintMovePayloadScheme) (*model.ResponseScheme, error) {

	if sprintID == 0 {
		return nil, model.ErrNoSprintID
	}

	url := fmt.Sprintf("/rest/agile/%v/sprint/%v/issue", i.version, sprintID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, url, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}

func (i *internalSprintImpl) Get(ctx context.Context, sprintID int) (*model.SprintScheme, *model.ResponseScheme, error) {

	if sprintID == 0 {
		return nil, nil, model.ErrNoSprintID
	}

	url := fmt.Sprintf("rest/agile/%v/sprint/%v", i.version, sprintID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	sprint := new(model.SprintScheme)
	res, err := i.c.Call(req, sprint)
	if err != nil {
		return nil, res, err
	}

	return sprint, res, nil
}

func (i *internalSprintImpl) Create(ctx context.Context, payload *model.SprintPayloadScheme) (*model.SprintScheme, *model.ResponseScheme, error) {

	url := fmt.Sprintf("rest/agile/%v/sprint", i.version)

	req, err := i.c.NewRequest(ctx, http.MethodPost, url, "", payload)
	if err != nil {
		return nil, nil, err
	}

	sprint := new(model.SprintScheme)
	res, err := i.c.Call(req, sprint)
	if err != nil {
		return nil, res, err
	}

	return sprint, res, nil
}

func (i *internalSprintImpl) Update(ctx context.Context, sprintID int, payload *model.SprintPayloadScheme) (*model.SprintScheme, *model.ResponseScheme, error) {

	if sprintID == 0 {
		return nil, nil, model.ErrNoSprintID
	}

	url := fmt.Sprintf("rest/agile/%v/sprint/%v", i.version, sprintID)

	req, err := i.c.NewRequest(ctx, http.MethodPut, url, "", payload)
	if err != nil {
		return nil, nil, err
	}

	sprint := new(model.SprintScheme)
	res, err := i.c.Call(req, sprint)
	if err != nil {
		return nil, res, err
	}

	return sprint, res, nil
}

func (i *internalSprintImpl) Path(ctx context.Context, sprintID int, payload *model.SprintPayloadScheme) (*model.SprintScheme, *model.ResponseScheme, error) {

	if sprintID == 0 {
		return nil, nil, model.ErrNoSprintID
	}

	url := fmt.Sprintf("rest/agile/%v/sprint/%v", i.version, sprintID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, url, "", payload)
	if err != nil {
		return nil, nil, err
	}

	sprint := new(model.SprintScheme)
	res, err := i.c.Call(req, sprint)
	if err != nil {
		return nil, res, err
	}

	return sprint, res, nil
}

func (i *internalSprintImpl) Delete(ctx context.Context, sprintID int) (*model.ResponseScheme, error) {

	if sprintID == 0 {
		return nil, model.ErrNoSprintID
	}

	url := fmt.Sprintf("rest/agile/%v/sprint/%v", i.version, sprintID)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, url, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}

func (i *internalSprintImpl) Issues(ctx context.Context, sprintID int, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.SprintIssuePageScheme, *model.ResponseScheme, error) {

	if sprintID == 0 {
		return nil, nil, model.ErrNoSprintID
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if !opts.ValidateQuery {
			params.Add("validateQuery", "false")
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

	url := fmt.Sprintf("rest/agile/%v/sprint/%v/issue?%v", i.version, sprintID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.SprintIssuePageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalSprintImpl) Start(ctx context.Context, sprintID int) (*model.ResponseScheme, error) {

	if sprintID == 0 {
		return nil, model.ErrNoSprintID
	}

	url := fmt.Sprintf("rest/agile/%v/sprint/%v", i.version, sprintID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, url, "", &model.SprintPayloadScheme{State: "Active"})
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}

func (i *internalSprintImpl) Close(ctx context.Context, sprintID int) (*model.ResponseScheme, error) {

	if sprintID == 0 {
		return nil, model.ErrNoSprintID
	}

	url := fmt.Sprintf("rest/agile/%v/sprint/%v", i.version, sprintID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, url, "", &model.SprintPayloadScheme{State: "Closed"})
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}
