package agile

import (
	"context"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// BoardBacklogConnector represents the board backlogs.
// Use it to search, get, create, delete, and change backlogs.
type BoardBacklogConnector interface {

	// Move moves issues to the backlog.
	//
	// This operation is equivalent to remove future and active sprints from a given set of issues.
	//
	// At most 50 issues may be moved at once.
	//
	// POST /rest/agile/1.0/backlog/issue
	//
	// https://docs.go-atlassian.io/jira-agile/boards/backlog#move-issues-to-backlog
	Move(ctx context.Context, issues []string) (*models.ResponseScheme, error)

	// MoveTo moves issues to the backlog of a particular board (if they are already on that board).
	//
	// This operation is equivalent to remove future and active sprints from a given set of issues if the board has sprints.
	//
	// If the board does not have sprints this will put the issues back into the backlog from the board.
	//
	// At most 50 issues may be moved at once.
	//
	// POST /rest/agile/1.0/backlog/{boardID}/issue
	//
	// https://docs.go-atlassian.io/jira-agile/boards/backlog#move-issues-to-a-board-backlog
	MoveTo(ctx context.Context, boardID int, payload *models.BoardBacklogPayloadScheme) (*models.ResponseScheme, error)
}
