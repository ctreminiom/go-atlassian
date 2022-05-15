package agile

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type BoardService interface {
	Get(ctx context.Context, boardId int) (
		*models.BoardScheme, *models.ResponseScheme, error)

	/*
		Create(ctx context.Context, payload *models.BoardPayloadScheme) (
			*models.BoardScheme, *models.ResponseScheme, error)
		Filter(ctx context.Context, filterId, startAt, maxResults int) (
			*models.BoardPageScheme, *models.ResponseScheme, error)
		Backlog(ctx context.Context, boardId, startAt, maxResults int, opts *models.IssueOptionScheme) (
			*models.BoardIssuePageScheme, *models.ResponseScheme, error)
		Configuration(ctx context.Context, boardId int) (
			*models.BoardConfigurationScheme, *models.ResponseScheme, error)
		Epics(ctx context.Context, boardId, startAt, maxResults int, done bool) (
			*models.BoardEpicPageScheme, *models.ResponseScheme, error)
		IssuesWithoutEpic(ctx context.Context, boardId, startAt, maxResults int, opts *models.IssueOptionScheme) (
			*models.BoardIssuePageScheme, *models.ResponseScheme, error)
		IssuesByEpic(ctx context.Context, boardId, epicID, startAt, maxResults int, opts *models.IssueOptionScheme) (
			*models.BoardIssuePageScheme, *models.ResponseScheme, error)
		Issues(ctx context.Context, boardId, startAt, maxResults int, opts *models.IssueOptionScheme) (
			*models.BoardIssuePageScheme, *models.ResponseScheme, error)
		Move(ctx context.Context, boardId int, payload *models.BoardMovementPayloadScheme) (
			*models.ResponseScheme, error)
		Projects(ctx context.Context, boardId, startAt, maxResults int) (
			*models.BoardProjectPageScheme, *models.ResponseScheme, error)
		Sprints(ctx context.Context, boardId, startAt, maxResults int, states []string) (
			*models.BoardSprintPageScheme, *models.ResponseScheme, error)
		IssuesBySprint(ctx context.Context, boardId, sprintId, startAt, maxResults int, opts *models.IssueOptionScheme) (
			*models.BoardIssuePageScheme, *models.ResponseScheme, error)
		Versions(ctx context.Context, boardID, startAt, maxResults int, released bool) (
			*models.BoardVersionPageScheme, *models.ResponseScheme, error)
		Delete(ctx context.Context, boardID int) (
			*models.ResponseScheme, error)
		Gets(ctx context.Context, opts *models.GetBoardsOptions, startAt, maxResults int) (
			*models.BoardPageScheme, *models.ResponseScheme, error)
	*/
}
