package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/agile"
	"net/http"
)

func NewBoardBacklogService(client service.Connector, version string) *BoardBacklogService {

	return &BoardBacklogService{
		internalClient: &internalBoardBacklogImpl{c: client, version: version},
	}
}

type BoardBacklogService struct {
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
// POST /rest/agile/1.0/backlog/{boardId}/issue
//
// https://docs.go-atlassian.io/jira-agile/boards/backlog#move-issues-to-a-board-backlog
func (b *BoardBacklogService) MoveTo(ctx context.Context, boardID int, payload *model.BoardBacklogPayloadScheme) (*model.ResponseScheme, error) {
	return b.internalClient.MoveTo(ctx, boardID, payload)
}

type internalBoardBacklogImpl struct {
	c       service.Connector
	version string
}

func (i *internalBoardBacklogImpl) Move(ctx context.Context, issues []string) (*model.ResponseScheme, error) {

	payload := map[string]interface{}{"issues": issues}

	url := fmt.Sprintf("rest/agile/%v/backlog/issue", i.version)

	req, err := i.c.NewRequest(ctx, http.MethodPost, url, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}

func (i *internalBoardBacklogImpl) MoveTo(ctx context.Context, boardID int, payload *model.BoardBacklogPayloadScheme) (*model.ResponseScheme, error) {

	if boardID == 0 {
		return nil, model.ErrNoBoardIDError
	}

	url := fmt.Sprintf("rest/agile/%v/backlog/%v/issue", i.version, boardID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, url, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}
