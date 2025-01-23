package agile

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// BoardConnector represents the Jira boards.
// Use it to search, get, create, delete, and change boards.
type BoardConnector interface {

	// Get returns the board for the given board ID.
	// This board will only be returned if the user has permission to view it.
	//
	//
	// Admins without the view permission will see the board as a private one,
	//
	// so will see only a subset of the board's data (board location for instance).
	//
	// GET /rest/agile/1.0/board/{boardID}
	//
	// https://docs.go-atlassian.io/jira-agile/boards#get-board
	Get(ctx context.Context, boardID int) (*model.BoardScheme, *model.ResponseScheme, error)

	// Create creates a new board. Board name, type and filter ID is required.
	//
	// POST /rest/agile/1.0/board
	//
	// Docs: https://docs.go-atlassian.io/jira-agile/boards#create-board
	Create(ctx context.Context, payload *model.BoardPayloadScheme) (*model.BoardScheme, *model.ResponseScheme, error)

	// Filter returns any boards which use the provided filter id.
	//
	// This method can be executed by users without a valid software license in order
	//
	// to find which boards are using a particular filter.
	//
	// GET /rest/agile/1.0/board/filter/{filterID}
	//
	// https://docs.go-atlassian.io/jira-agile/boards#get-board-by-filter-id
	Filter(ctx context.Context, filterID, startAt, maxResults int) (*model.BoardPageScheme, *model.ResponseScheme, error)

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
	Backlog(ctx context.Context, boardID int, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme, *model.ResponseScheme, error)

	// Configuration get the board configuration.
	//
	// GET /rest/agile/1.0/board/{boardID}/configuration
	//
	// https://docs.go-atlassian.io/jira-agile/boards#get-configuration
	Configuration(ctx context.Context, boardID int) (*model.BoardConfigurationScheme, *model.ResponseScheme, error)

	// Epics returns all epics from the board, for the given board ID.
	//
	// This only includes epics that the user has permission to view.
	//
	// Note, if the user does not have permission to view the board, no epics will be returned at all.
	//
	// GET /rest/agile/1.0/board/{boardID}/epic
	//
	// https://docs.go-atlassian.io/jira-agile/boards#get-epics
	Epics(ctx context.Context, boardID, startAt, maxResults int, done bool) (*model.BoardEpicPageScheme, *model.ResponseScheme, error)

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
	IssuesWithoutEpic(ctx context.Context, boardID int, opts *model.IssueOptionScheme, startAt, maxResults int) (
		*model.BoardIssuePageScheme, *model.ResponseScheme, error)

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
	IssuesByEpic(ctx context.Context, boardID, epicID int, opts *model.IssueOptionScheme, startAt, maxResults int) (
		*model.BoardIssuePageScheme, *model.ResponseScheme, error)

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
	Issues(ctx context.Context, boardID int, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme,
		*model.ResponseScheme, error)

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
	Move(ctx context.Context, boardID int, payload *model.BoardMovementPayloadScheme) (*model.ResponseScheme, error)

	// Projects returns all projects that are associated with the board, for the given board ID.
	//
	// If the user does not have permission to view the board, no projects will be returned at all.
	//
	// Returned projects are ordered by the name.
	//
	// GET /rest/agile/1.0/board/{boardID}/project
	//
	// https://docs.go-atlassian.io/jira-agile/boards#get-projects
	Projects(ctx context.Context, boardID, startAt, maxResults int) (*model.BoardProjectPageScheme, *model.ResponseScheme, error)

	// Sprints returns all sprints from a board, for a given board ID.
	//
	// This only includes sprints that the user has permission to view.
	//
	// GET /rest/agile/1.0/board/{boardID}/sprint
	//
	// https://docs.go-atlassian.io/jira-agile/boards#get-all-sprints
	Sprints(ctx context.Context, boardID, startAt, maxResults int, states []string) (*model.BoardSprintPageScheme,
		*model.ResponseScheme, error)

	// IssuesBySprint get all issues you have access to that belong to the sprint from the board.
	//
	// Issue returned from this resource contains additional fields like: sprint, closedSprints, flagged and epic.
	//
	// Issues are returned ordered by rank. JQL order has higher priority than default rank.
	//
	// GET /rest/agile/1.0/board/{boardID}/sprint/{sprintID}/issue
	//
	// https://docs.go-atlassian.io/jira-agile/boards#get-board-issues-for-sprint
	IssuesBySprint(ctx context.Context, boardID, sprintID int, opts *model.IssueOptionScheme, startAt, maxResults int) (
		*model.BoardIssuePageScheme, *model.ResponseScheme, error)

	// Versions returns all versions from a board, for a given board ID.
	//
	// This only includes versions that the user has permission to view.
	//
	// Note, if the user does not have permission to view the board, no versions will be returned at all.
	//
	// Returned versions are ordered by the name of the project from which they belong and then by sequence defined by user.
	//
	// GET /rest/agile/1.0/board/{boardID}/sprint/{sprintID}/issue
	//
	// https://docs.go-atlassian.io/jira-agile/boards#get-all-versions
	Versions(ctx context.Context, boardID, startAt, maxResults int, released bool) (*model.BoardVersionPageScheme,
		*model.ResponseScheme, error)

	// Delete deletes the board. Admin without the view permission can still remove the board.
	//
	// DELETE /rest/agile/1.0/board/{boardID}
	//
	// https://docs.go-atlassian.io/jira-agile/boards#delete-board
	Delete(ctx context.Context, boardID int) (*model.ResponseScheme, error)

	// Gets returns all boards. This only includes boards that the user has permission to view.
	//
	// GET /rest/agile/1.0/board
	//
	// https://docs.go-atlassian.io/jira-agile/boards#get-boards
	Gets(ctx context.Context, opts *model.GetBoardsOptions, startAt, maxResults int) (*model.BoardPageScheme,
		*model.ResponseScheme, error)
}
