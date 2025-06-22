package internal

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/agile"
)

// NewBoardBacklogService creates a new instance of BoardBacklogService.
// It takes a service.Connector and a version string as input and returns a pointer to BoardBacklogService.
func NewBoardBacklogService(client service.Connector, version string) *BoardBacklogService {
	return &BoardBacklogService{
		internalClient: &internalBoardBacklogImpl{c: client, version: version},
	}
}

// BoardBacklogService provides methods to interact with board backlog operations in Jira Agile.
type BoardBacklogService struct {
	// internalClient is the connector interface for board backlog operations.
	internalClient agile.BoardBacklogConnector
}

// Move moves issues to the backlog.
//
// This operation is equivalent to remove future and active sprints from a given set of issues.
//
// At most 50 issues may be moved at once.
//
// POST /rest/agile/1.0/backlog/issue
//
// https://docs.go-atlassian.io/jira-agile/boards/backlog#move-issues-to-backlog
func (b *BoardBacklogService) Move(ctx context.Context, issues []string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*BoardBacklogService).Move")
	defer span.End()

	return b.internalClient.Move(ctx, issues)
}

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
func (b *BoardBacklogService) MoveTo(ctx context.Context, boardID int, payload *model.BoardBacklogPayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*BoardBacklogService).MoveTo")
	defer span.End()

	return b.internalClient.MoveTo(ctx, boardID, payload)
}

type internalBoardBacklogImpl struct {
	c       service.Connector
	version string
}

func (i *internalBoardBacklogImpl) Move(ctx context.Context, issues []string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalBoardBacklogImpl).Move", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.Int("jira.issue.count", len(issues)),
		attribute.String("operation.name", "move_issues_to_backlog"),
	)

	payload := map[string]interface{}{"issues": issues}

	url := fmt.Sprintf("rest/agile/%v/backlog/issue", i.version)

	req, err := i.c.NewRequest(ctx, http.MethodPost, url, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	res, err := i.c.Call(req, nil)
	if err != nil {
		recordError(span, err)
		return res, err
	}

	setOK(span)
	return res, nil
}

func (i *internalBoardBacklogImpl) MoveTo(ctx context.Context, boardID int, payload *model.BoardBacklogPayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalBoardBacklogImpl).MoveTo", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.Int("jira.board.id", boardID),
		attribute.String("operation.name", "move_issues_to_board_backlog"),
	)

	if boardID == 0 {
		err := fmt.Errorf("agile: %w", model.ErrNoBoardID)
		recordError(span, err)
		return nil, err
	}

	url := fmt.Sprintf("rest/agile/%v/backlog/%v/issue", i.version, boardID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, url, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	res, err := i.c.Call(req, nil)
	if err != nil {
		recordError(span, err)
		return res, err
	}

	setOK(span)
	return res, nil
}
