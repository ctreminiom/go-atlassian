package agile

import (
	"context"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type SprintConnector interface {

	// Get Returns the sprint for a given sprint ID.
	//
	// The sprint will only be returned if the user can view the board that the sprint was created on,
	//
	// or view at least one of the issues in the sprint.
	//
	// GET /rest/agile/1.0/sprint/{sprintID}
	//
	// https://docs.go-atlassian.io/jira-agile/sprints#get-sprint
	Get(ctx context.Context, sprintID int) (*models.SprintScheme, *models.ResponseScheme, error)

	// Create creates a future sprint.
	//
	// Sprint name and origin board id are required.
	//
	// Start date, end date, and goal are optional.
	//
	// POST /rest/agile/1.0/sprint
	//
	// https://docs.go-atlassian.io/jira-agile/sprints#create-print
	Create(ctx context.Context, payload *models.SprintPayloadScheme) (*models.SprintScheme, *models.ResponseScheme, error)

	// Update Performs a full update of a sprint.
	//
	// A full update means that the result will be exactly the same as the request body.
	//
	// Any fields not present in the request JSON will be set to null.
	//
	// PUT /rest/agile/1.0/sprint/{sprintID}
	//
	// https://docs.go-atlassian.io/jira-agile/sprints#update-sprint
	Update(ctx context.Context, sprintID int, payload *models.SprintPayloadScheme) (*models.SprintScheme, *models.ResponseScheme, error)

	// Path Performs a partial update of a sprint.
	//
	// A partial update means that fields not present in the request JSON will not be updated.
	//
	// POST /rest/agile/1.0/sprint/{sprintID}
	//
	// https://docs.go-atlassian.io/jira-agile/sprints#partially-update-sprint
	Path(ctx context.Context, sprintID int, payload *models.SprintPayloadScheme) (*models.SprintScheme, *models.ResponseScheme, error)

	// Delete deletes a sprint.
	//
	// Once a sprint is deleted, all open issues in the sprint will be moved to the backlog.
	//
	// DELETE /rest/agile/1.0/sprint/{sprintID}
	//
	// https://docs.go-atlassian.io/jira-agile/sprints#delete-sprint
	Delete(ctx context.Context, sprintID int) (*models.ResponseScheme, error)

	// Issues returns all issues in a sprint, for a given sprint ID.
	//
	// This only includes issues that the user has permission to view.
	//
	// By default, the returned issues are ordered by rank.
	//
	// GET /rest/agile/1.0/sprint/{sprintID}/issue
	//
	// https://docs.go-atlassian.io/jira-agile/sprints#get-issues-for-sprint
	Issues(ctx context.Context, sprintID int, opts *models.IssueOptionScheme, startAt, maxResults int) (*models.SprintIssuePageScheme, *models.ResponseScheme, error)

	// Start initiate the Sprint
	//
	// PUT /rest/agile/1.0/sprint/{sprintID}
	//
	// https://docs.go-atlassian.io/jira-agile/sprints#start-sprint
	Start(ctx context.Context, sprintID int) (*models.ResponseScheme, error)

	// Close closes the Sprint
	//
	// PUT /rest/agile/1.0/sprint/{sprintID}
	//
	// https://docs.go-atlassian.io/jira-agile/sprints#close-sprint
	Close(ctx context.Context, sprintID int) (*models.ResponseScheme, error)

	// Move moves issues to a sprint, for a given sprint ID.
	//
	// Issues can only be moved to open or active sprints.
	//
	// The maximum number of issues that can be moved in one operation is 50.
	//
	// POST /rest/agile/1.0/sprint/{sprintID}/issue
	//
	// https://docs.go-atlassian.io/jira-agile/sprints#move-issues-to-sprint
	Move(ctx context.Context, sprintID int, payload *models.SprintMovePayloadScheme) (*models.ResponseScheme, error)
}
