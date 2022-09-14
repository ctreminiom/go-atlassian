package agile

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type Sprint interface {

	// Get Returns the sprint for a given sprint ID.
	// The sprint will only be returned if the user can view the board that the sprint was created on,
	// or view at least one of the issues in the sprint.
	// Docs: https://docs.go-atlassian.io/jira-agile/sprints#get-sprint
	Get(ctx context.Context, sprintId int) (*models.SprintScheme, *models.ResponseScheme, error)

	// Create creates a future sprint.
	// Sprint name and origin board id are required.
	// Start date, end date, and goal are optional.
	// Docs: https://docs.go-atlassian.io/jira-agile/sprints#create-print
	Create(ctx context.Context, payload *models.SprintPayloadScheme) (*models.SprintScheme, *models.ResponseScheme, error)

	// Update Performs a full update of a sprint.
	// A full update means that the result will be exactly the same as the request body.
	// Any fields not present in the request JSON will be set to null.
	// Docs: https://docs.go-atlassian.io/jira-agile/sprints#update-sprint
	Update(ctx context.Context, sprintId int, payload *models.SprintPayloadScheme) (*models.SprintScheme, *models.ResponseScheme,
		error)

	// Path Performs a partial update of a sprint.
	// A partial update means that fields not present in the request JSON will not be updated.
	// Docs: https://docs.go-atlassian.io/jira-agile/sprints#partially-update-sprint
	Path(ctx context.Context, sprintId int, payload *models.SprintPayloadScheme) (*models.SprintScheme, *models.ResponseScheme,
		error)

	// Delete deletes a sprint.
	// Once a sprint is deleted, all open issues in the sprint will be moved to the backlog.
	Delete(ctx context.Context, sprintId int) (*models.ResponseScheme, error)

	// Issues returns all issues in a sprint, for a given sprint ID.
	// This only includes issues that the user has permission to view.
	// By default, the returned issues are ordered by rank.
	// Docs: https://docs.go-atlassian.io/jira-agile/sprints#get-issues-for-sprint
	Issues(ctx context.Context, sprintId int, opts *models.IssueOptionScheme, startAt, maxResults int) (*models.SprintIssuePageScheme,
		*models.ResponseScheme, error)

	// Start initiate the Sprint
	// Docs: https://docs.go-atlassian.io/jira-agile/sprints#start-sprint
	Start(ctx context.Context, sprintId int) (*models.ResponseScheme, error)

	// Close closes the Sprint
	// Docs: https://docs.go-atlassian.io/jira-agile/sprints#close-sprint
	Close(ctx context.Context, sprintId int) (*models.ResponseScheme, error)
}
