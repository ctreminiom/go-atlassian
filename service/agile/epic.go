package agile

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type Epic interface {

	// Get returns the epic for a given epic ID.
	// This epic will only be returned if the user has permission to view it.
	// Note: This operation does not work for epics in next-gen projects.
	// Docs: https://docs.go-atlassian.io/jira-agile/epics#get-epic
	Get(ctx context.Context, epicIdOrKey string) (*model.EpicScheme, *model.ResponseScheme, error)

	// Issues returns all issues that belong to the epic, for the given epic ID.
	// This only includes issues that the user has permission to view.
	// Issues returned from this resource include Agile fields, like sprint, closedSprints,
	// flagged, and epic.
	// By default, the returned issues are ordered by rank.
	// Docs: https://docs.go-atlassian.io/jira-agile/epics#get-issues-for-epic
	Issues(ctx context.Context, epicIdOrKey string, startAt, maxResults int, opts *model.IssueOptionScheme) (*model.BoardIssuePageScheme,
		*model.ResponseScheme, error)

	// Move moves issues to an epic, for a given epic id.
	// Issues can be only in a single epic at the same time.
	// That means that already assigned issues to an epic, will not be assigned to the previous epic anymore.
	// The user needs to have the edit issue permission for all issue they want to move and to the epic.
	// The maximum number of issues that can be moved in one operation is 50.
	// Docs: TODO: The documentation needs to be created, raise a ticket here: https://github.com/ctreminiom/go-atlassian/issues
	Move(ctx context.Context, epicIdOrKey string, issues []string) (*model.ResponseScheme, error)
}
