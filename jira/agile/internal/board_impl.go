package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/agile"
)

// NewBoardService creates a new instance of BoardService.
// It takes a service.Connector and a version string as input and returns a pointer to BoardService.
func NewBoardService(client service.Connector, version string) *BoardService {
	return &BoardService{
		internalClient: &internalBoardImpl{c: client, version: version},
	}
}

// BoardService provides methods to interact with board operations in Jira Agile.
type BoardService struct {
	// internalClient is the connector interface for board operations.
	internalClient agile.BoardConnector
}

// Get returns the board for the given board ID.
// This board will only be returned if the user has permission to view it.
//
// Admins without the view permission will see the board as a private one,
//
// so will see only a subset of the board's data (board location for instance).
//
// GET /rest/agile/1.0/board/{boardID}
//
// https://docs.go-atlassian.io/jira-agile/boards#get-board
func (b *BoardService) Get(ctx context.Context, boardID int) (*model.BoardScheme, *model.ResponseScheme, error) {
	return b.internalClient.Get(ctx, boardID)
}

// Create creates a new board. Board name, type and filter ID is required.
//
// POST /rest/agile/1.0/board
//
// Docs: https://docs.go-atlassian.io/jira-agile/boards#create-board
func (b *BoardService) Create(ctx context.Context, payload *model.BoardPayloadScheme) (*model.BoardScheme, *model.ResponseScheme, error) {
	return b.internalClient.Create(ctx, payload)
}

// Filter returns any boards which use the provided filter id.
//
// # This method can be executed by users without a valid software license in order
//
// to find which boards are using a particular filter.
//
// GET /rest/agile/1.0/board/filter/{filterID}
//
// https://docs.go-atlassian.io/jira-agile/boards#get-board-by-filter-id
func (b *BoardService) Filter(ctx context.Context, filterID, startAt, maxResults int) (*model.BoardPageScheme, *model.ResponseScheme, error) {
	return b.internalClient.Filter(ctx, filterID, startAt, maxResults)
}

// Backlog returns all issues from the board's backlog, for the given board ID.
//
// This only includes issues that the user has permission to view.
//
// The backlog contains incomplete issues that are not assigned to any future or active sprint.
//
// Note, if the user does not have permission to view the board, no issues will be returned at all.
//
// Issues returned from this resource include Agile fields, like sprint, closedSprints, flagged, and epic.
//
// By default, the returned issues are ordered by rank.
//
// GET /rest/agile/1.0/board/{boardID}/backlog
//
// https://docs.go-atlassian.io/jira-agile/boards#get-issues-for-backlog
func (b *BoardService) Backlog(ctx context.Context, boardID int, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {
	return b.internalClient.Backlog(ctx, boardID, opts, startAt, maxResults)
}

// Configuration get the board configuration.
//
// GET /rest/agile/1.0/board/{boardID}/configuration
//
// https://docs.go-atlassian.io/jira-agile/boards#get-configuration
func (b *BoardService) Configuration(ctx context.Context, boardID int) (*model.BoardConfigurationScheme, *model.ResponseScheme, error) {
	return b.internalClient.Configuration(ctx, boardID)
}

// Epics returns all epics from the board, for the given board ID.
//
// This only includes epics that the user has permission to view.
//
// Note, if the user does not have permission to view the board, no epics will be returned at all.
//
// GET /rest/agile/1.0/board/{boardID}/epic
//
// https://docs.go-atlassian.io/jira-agile/boards#get-epics
func (b *BoardService) Epics(ctx context.Context, boardID, startAt, maxResults int, done bool) (*model.BoardEpicPageScheme, *model.ResponseScheme, error) {
	return b.internalClient.Epics(ctx, boardID, startAt, maxResults, done)
}

// IssuesWithoutEpic returns all issues that do not belong to any epic on a board, for a given board ID.
//
// This only includes issues that the user has permission to view.
//
// Issues returned from this resource include Agile fields, like sprint, closedSprints, flagged, and epic.
//
// By default, the returned issues are ordered by rank.
//
// GET /rest/agile/1.0/board/{boardID}/epic/none/issue
//
// Docs: https://docs.go-atlassian.io/jira-agile/boards#get-issues-without-epic-for-board
func (b *BoardService) IssuesWithoutEpic(ctx context.Context, boardID int, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {
	return b.internalClient.IssuesWithoutEpic(ctx, boardID, opts, startAt, maxResults)
}

// IssuesByEpic returns all issues that belong to an epic on the board, for the given epic ID and the board ID.
//
// This only includes issues that the user has permission to view.
//
// Issues returned from this resource include Agile fields, like sprint, closedSprints,
//
// flagged, and epic. By default, the returned issues are ordered by rank.
//
// GET /rest/agile/1.0/board/{boardID}/epic/none/issue
//
// https://docs.go-atlassian.io/jira-agile/boards#get-board-issues-for-epic
func (b *BoardService) IssuesByEpic(ctx context.Context, boardID, epicID int, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {
	return b.internalClient.IssuesByEpic(ctx, boardID, epicID, opts, startAt, maxResults)
}

// Issues returns all issues from a board, for a given board ID.
//
// This only includes issues that the user has permission to view.
//
// An issue belongs to the board if its status is mapped to the board's column.
//
// Epic issues do not belong to the scrum boards. Note, if the user does not have permission to view the board,
//
// no issues will be returned at all.
//
// Issues returned from this resource include Agile fields, like sprint, closedSprints, flagged, and epic.
//
// By default, the returned issues are ordered by rank.
//
// GET /rest/agile/1.0/board/{boardID}/issue
//
// https://docs.go-atlassian.io/jira-agile/boards#get-issues-for-board
func (b *BoardService) Issues(ctx context.Context, boardID int, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {
	return b.internalClient.Issues(ctx, boardID, opts, startAt, maxResults)
}

// Move issues from the backlog to the board (if they are already in the backlog of that board).
//
// This operation either moves an issue(s) onto a board from the backlog (by adding it to the issueList for the board)
//
// Or transitions the issue(s) to the first column for a kanban board with backlog.
//
// At most 50 issues may be moved at once.
//
// POST /rest/agile/1.0/board/{boardID}/issue
//
// https://docs.go-atlassian.io/jira-agile/boards#move-issues-to-backlog-for-board
func (b *BoardService) Move(ctx context.Context, boardID int, payload *model.BoardMovementPayloadScheme) (*model.ResponseScheme, error) {
	return b.internalClient.Move(ctx, boardID, payload)
}

// Projects returns all projects that are associated with the board, for the given board ID.
//
// If the user does not have permission to view the board, no projects will be returned at all.
//
// Returned projects are ordered by the name.
//
// GET /rest/agile/1.0/board/{boardID}/project
//
// https://docs.go-atlassian.io/jira-agile/boards#get-projects
func (b *BoardService) Projects(ctx context.Context, boardID, startAt, maxResults int) (*model.BoardProjectPageScheme, *model.ResponseScheme, error) {
	return b.internalClient.Projects(ctx, boardID, startAt, maxResults)
}

// Sprints returns all sprints from a board, for a given board ID.
//
// This only includes sprints that the user has permission to view.
//
// GET /rest/agile/1.0/board/{boardID}/sprint
//
// https://docs.go-atlassian.io/jira-agile/boards#get-all-sprints
func (b *BoardService) Sprints(ctx context.Context, boardID, startAt, maxResults int, states []string) (*model.BoardSprintPageScheme, *model.ResponseScheme, error) {
	return b.internalClient.Sprints(ctx, boardID, startAt, maxResults, states)
}

// IssuesBySprint get all issues you have access to that belong to the sprint from the board.
//
// Issue returned from this resource contains additional fields like: sprint, closedSprints, flagged and epic.
//
// Issues are returned ordered by rank. JQL order has higher priority than default rank.
//
// GET /rest/agile/1.0/board/{boardID}/sprint/{sprintId}/issue
//
// https://docs.go-atlassian.io/jira-agile/boards#get-board-issues-for-sprint
func (b *BoardService) IssuesBySprint(ctx context.Context, boardID, sprintID int, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {
	return b.internalClient.IssuesBySprint(ctx, boardID, sprintID, opts, startAt, maxResults)
}

// Versions returns all versions from a board, for a given board ID.
//
// This only includes versions that the user has permission to view.
//
// Note, if the user does not have permission to view the board, no versions will be returned at all.
//
// Returned versions are ordered by the name of the project from which they belong and then by sequence defined by user.
//
// GET /rest/agile/1.0/board/{boardID}/sprint/{sprintId}/issue
//
// https://docs.go-atlassian.io/jira-agile/boards#get-all-versions
func (b *BoardService) Versions(ctx context.Context, boardID, startAt, maxResults int, released bool) (*model.BoardVersionPageScheme, *model.ResponseScheme, error) {
	return b.internalClient.Versions(ctx, boardID, startAt, maxResults, released)
}

// Delete deletes the board. Admin without the view permission can still remove the board.
//
// DELETE /rest/agile/1.0/board/{boardID}
//
// https://docs.go-atlassian.io/jira-agile/boards#delete-board
func (b *BoardService) Delete(ctx context.Context, boardID int) (*model.ResponseScheme, error) {
	return b.internalClient.Delete(ctx, boardID)
}

// Gets returns all boards. This only includes boards that the user has permission to view.
//
// GET /rest/agile/1.0/board
//
// https://docs.go-atlassian.io/jira-agile/boards#get-boards
func (b *BoardService) Gets(ctx context.Context, opts *model.GetBoardsOptions, startAt, maxResults int) (*model.BoardPageScheme, *model.ResponseScheme, error) {
	return b.internalClient.Gets(ctx, opts, startAt, maxResults)
}

type internalBoardImpl struct {
	c       service.Connector
	version string
}

func (i *internalBoardImpl) Get(ctx context.Context, boardID int) (*model.BoardScheme, *model.ResponseScheme, error) {

	if boardID == 0 {
		return nil, nil, model.ErrNoBoardID
	}

	url := fmt.Sprintf("rest/agile/%v/board/%v", i.version, boardID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	boards := new(model.BoardScheme)
	res, err := i.c.Call(req, boards)
	if err != nil {
		return nil, res, err
	}

	return boards, res, nil
}

func (i *internalBoardImpl) Create(ctx context.Context, payload *model.BoardPayloadScheme) (*model.BoardScheme, *model.ResponseScheme, error) {

	url := fmt.Sprintf("rest/agile/%v/board", i.version)

	req, err := i.c.NewRequest(ctx, http.MethodPost, url, "", payload)
	if err != nil {
		return nil, nil, err
	}

	board := new(model.BoardScheme)
	res, err := i.c.Call(req, board)
	if err != nil {
		return nil, res, err
	}

	return board, res, nil
}

func (i *internalBoardImpl) Filter(ctx context.Context, filterID, startAt, maxResults int) (*model.BoardPageScheme, *model.ResponseScheme, error) {

	if filterID == 0 {
		return nil, nil, model.ErrNoFilterID
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	url := fmt.Sprintf("rest/agile/%v/board/filter/%v?%v", i.version, filterID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BoardPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalBoardImpl) Backlog(ctx context.Context, boardID int, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {

	if boardID == 0 {
		return nil, nil, model.ErrNoBoardID
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		params.Add("validateQuery", fmt.Sprintf("%v", opts.ValidateQuery))

		if opts.JQL != "" {
			params.Add("jql", opts.JQL)
		}

		if len(opts.Expand) != 0 {
			params.Add("expand", strings.Join(opts.Expand, ","))
		}

		if len(opts.Fields) != 0 {
			params.Add("fields", strings.Join(opts.Fields, ","))
		}
	}

	url := fmt.Sprintf("rest/agile/%v/board/%v/backlog?%v", i.version, boardID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BoardIssuePageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalBoardImpl) Configuration(ctx context.Context, boardID int) (*model.BoardConfigurationScheme, *model.ResponseScheme, error) {

	if boardID == 0 {
		return nil, nil, model.ErrNoBoardID
	}

	url := fmt.Sprintf("rest/agile/%v/board/%v/configuration", i.version, boardID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	conf := new(model.BoardConfigurationScheme)
	res, err := i.c.Call(req, conf)
	if err != nil {
		return nil, res, err
	}

	return conf, res, nil
}

func (i *internalBoardImpl) Epics(ctx context.Context, boardID, startAt, maxResults int, done bool) (*model.BoardEpicPageScheme, *model.ResponseScheme, error) {

	if boardID == 0 {
		return nil, nil, model.ErrNoBoardID
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("done", fmt.Sprintf("%t", done))

	url := fmt.Sprintf("rest/agile/%v/board/%v/epic?%v", i.version, boardID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BoardEpicPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalBoardImpl) IssuesWithoutEpic(ctx context.Context, boardID int, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {

	if boardID == 0 {
		return nil, nil, model.ErrNoBoardID
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

	url := fmt.Sprintf("rest/agile/%v/board/%v/epic/none/issue?%v", i.version, boardID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BoardIssuePageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalBoardImpl) IssuesByEpic(ctx context.Context, boardID, epicID int, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {

	if boardID == 0 {
		return nil, nil, model.ErrNoBoardID
	}

	if epicID == 0 {
		return nil, nil, model.ErrNoEpicID
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

	url := fmt.Sprintf("rest/agile/%v/board/%v/epic/%v/issue?%v", i.version, boardID, epicID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BoardIssuePageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalBoardImpl) Issues(ctx context.Context, boardID int, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {

	if boardID == 0 {
		return nil, nil, model.ErrNoBoardID
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

	url := fmt.Sprintf("rest/agile/%v/board/%v/issue?%v", i.version, boardID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BoardIssuePageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalBoardImpl) Move(ctx context.Context, boardID int, payload *model.BoardMovementPayloadScheme) (*model.ResponseScheme, error) {

	if boardID == 0 {
		return nil, model.ErrNoBoardID
	}

	url := fmt.Sprintf("rest/agile/%v/board/%v/issue", i.version, boardID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, url, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}

func (i *internalBoardImpl) Projects(ctx context.Context, boardID, startAt, maxResults int) (*model.BoardProjectPageScheme, *model.ResponseScheme, error) {

	if boardID == 0 {
		return nil, nil, model.ErrNoBoardID
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	url := fmt.Sprintf("rest/agile/%v/board/%v/project?%v", i.version, boardID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BoardProjectPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalBoardImpl) Sprints(ctx context.Context, boardID, startAt, maxResults int, states []string) (*model.BoardSprintPageScheme, *model.ResponseScheme, error) {

	if boardID == 0 {
		return nil, nil, model.ErrNoBoardID
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("state", strings.Join(states, ","))

	url := fmt.Sprintf("rest/agile/%v/board/%v/sprint?%v", i.version, boardID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BoardSprintPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalBoardImpl) IssuesBySprint(ctx context.Context, boardID, sprintID int, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {

	if boardID == 0 {
		return nil, nil, model.ErrNoBoardID
	}

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

	url := fmt.Sprintf("rest/agile/%v/board/%v/sprint/%v/issue?%v", i.version, boardID, sprintID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BoardIssuePageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalBoardImpl) Versions(ctx context.Context, boardID, startAt, maxResults int, released bool) (*model.BoardVersionPageScheme, *model.ResponseScheme, error) {

	if boardID == 0 {
		return nil, nil, model.ErrNoBoardID
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("released", fmt.Sprintf("%t", released))

	url := fmt.Sprintf("rest/agile/%v/board/%v/version?%v", i.version, boardID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BoardVersionPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalBoardImpl) Delete(ctx context.Context, boardID int) (*model.ResponseScheme, error) {

	if boardID == 0 {
		return nil, model.ErrNoBoardID
	}

	url := fmt.Sprintf("rest/agile/%v/board/%v", i.version, boardID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, url, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalBoardImpl) Gets(ctx context.Context, opts *model.GetBoardsOptions, startAt, maxResults int) (*model.BoardPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if opts.BoardType != "" {
			params.Add("type", opts.BoardType)
		}

		if opts.BoardName != "" {
			params.Add("name", opts.BoardName)
		}

		if opts.ProjectKeyOrID != "" {
			params.Add("projectKeyOrId", opts.ProjectKeyOrID)
		}

		if opts.AccountIDLocation != "" {
			params.Add("accountIdLocation", opts.AccountIDLocation)
		}

		if opts.ProjectIDLocation != "" {
			params.Add("projectLocation", opts.ProjectIDLocation)
		}

		if opts.IncludePrivate {
			params.Add("includePrivate", "true")
		}

		if opts.NegateLocationFiltering {
			params.Add("negateLocationFiltering", "true")
		}

		if opts.OrderBy != "" {
			params.Add("orderBy", opts.OrderBy)
		}

		if opts.Expand != "" {
			params.Add("expand", opts.Expand)
		}

		if opts.FilterID != 0 {
			params.Add("filterId", strconv.Itoa(opts.FilterID))
		}
	}

	url := fmt.Sprintf("rest/agile/%v/board?%v", i.version, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BoardPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}
