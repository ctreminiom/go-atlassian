package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/agile"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func NewBoardService(client service.Client, version string) (agile.Board, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &BoardService{client, version}, nil
}

type BoardService struct {
	c       service.Client
	version string
}

func (b BoardService) Get(ctx context.Context, boardId int) (*model.BoardScheme, *model.ResponseScheme, error) {

	if boardId == 0 {
		return nil, nil, model.ErrNoBoardIDError
	}

	endpoint := fmt.Sprintf("/rest/agile/%v/board/%v", b.version, boardId)

	request, err := b.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var boards model.BoardScheme
	response, err := b.c.Call(request, &boards)
	if err != nil {
		return nil, response, err
	}

	return &boards, response, nil
}

func (b BoardService) Create(ctx context.Context, payload *model.BoardPayloadScheme) (*model.BoardScheme, *model.ResponseScheme, error) {

	reader, err := b.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("/rest/agile/%v/board", b.version)

	request, err := b.c.NewJsonRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	var board model.BoardScheme
	response, err := b.c.Call(request, &board)
	if err != nil {
		return nil, response, err
	}

	return &board, response, nil
}

func (b BoardService) Filter(ctx context.Context, filterId, startAt, maxResults int) (*model.BoardPageScheme, *model.ResponseScheme, error) {

	if filterId == 0 {
		return nil, nil, model.ErrNoFilterIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	endpoint := fmt.Sprintf("/rest/agile/%v/board/filter/%v?%v", b.version, filterId, params.Encode())

	request, err := b.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	var page model.BoardPageScheme
	response, err := b.c.Call(request, &page)
	if err != nil {
		return nil, response, err
	}

	return &page, response, nil
}

func (b BoardService) Backlog(ctx context.Context, boardId, startAt, maxResults int, opts *model.IssueOptionScheme) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {

	if boardId == 0 {
		return nil, nil, model.ErrNoBoardIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if !opts.ValidateQuery {
			params.Add("validateQuery", "false")
		} else {
			params.Add("validateQuery", "true")
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

	endpoint := fmt.Sprintf("/rest/agile/%v/board/%v/backlog?%v", b.version, boardId, params.Encode())

	request, err := b.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	var issues model.BoardIssuePageScheme
	response, err := b.c.Call(request, &issues)
	if err != nil {
		return nil, response, err
	}

	return &issues, response, nil
}

func (b BoardService) Configuration(ctx context.Context, boardId int) (*model.BoardConfigurationScheme, *model.ResponseScheme, error) {

	if boardId == 0 {
		return nil, nil, model.ErrNoBoardIDError
	}

	endpoint := fmt.Sprintf("/rest/agile/%v/board/%v/configuration", b.version, boardId)

	request, err := b.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	var configuration model.BoardConfigurationScheme
	response, err := b.c.Call(request, &configuration)
	if err != nil {
		return nil, response, err
	}

	return &configuration, response, nil
}

func (b BoardService) Epics(ctx context.Context, boardId, startAt, maxResults int, done bool) (*model.BoardEpicPageScheme, *model.ResponseScheme, error) {

	if boardId == 0 {
		return nil, nil, model.ErrNoBoardIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("done", fmt.Sprintf("%t", done))

	endpoint := fmt.Sprintf("/rest/agile/%v/board/%v/epic?%v", b.version, boardId, params.Encode())

	request, err := b.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	var epics model.BoardEpicPageScheme
	response, err := b.c.Call(request, &epics)
	if err != nil {
		return nil, response, err
	}

	return &epics, response, nil
}

func (b BoardService) IssuesWithoutEpic(ctx context.Context, boardId, startAt, maxResults int, opts *model.IssueOptionScheme) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {

	if boardId == 0 {
		return nil, nil, model.ErrNoBoardIDError
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

	endpoint := fmt.Sprintf("/rest/agile/%v/board/%v/epic/none/issue?%v", b.version, boardId, params.Encode())

	request, err := b.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	var issues model.BoardIssuePageScheme
	response, err := b.c.Call(request, &issues)
	if err != nil {
		return nil, response, err
	}

	return &issues, response, nil
}

func (b BoardService) IssuesByEpic(ctx context.Context, boardId, epicId, startAt, maxResults int, opts *model.IssueOptionScheme) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {

	if boardId == 0 {
		return nil, nil, model.ErrNoBoardIDError
	}

	if epicId == 0 {
		return nil, nil, model.ErrNoEpicIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if !opts.ValidateQuery {
			params.Add("validateQuery ", "false")
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

	endpoint := fmt.Sprintf("/rest/agile/%v/board/%v/epic/%v/issue?%v", b.version, boardId, epicId, params.Encode())

	request, err := b.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	var issues model.BoardIssuePageScheme
	response, err := b.c.Call(request, &issues)
	if err != nil {
		return nil, response, err
	}

	return &issues, response, nil
}

func (b BoardService) Issues(ctx context.Context, boardId, startAt, maxResults int, opts *model.IssueOptionScheme) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {

	if boardId == 0 {
		return nil, nil, model.ErrNoBoardIDError
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

	endpoint := fmt.Sprintf("/rest/agile/%v/board/%v/issue?%v", b.version, boardId, params.Encode())

	request, err := b.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	var issues model.BoardIssuePageScheme
	response, err := b.c.Call(request, &issues)
	if err != nil {
		return nil, response, err
	}

	return &issues, response, nil
}

func (b BoardService) Move(ctx context.Context, boardId int, payload *model.BoardMovementPayloadScheme) (*model.ResponseScheme, error) {

	if boardId == 0 {
		return nil, model.ErrNoBoardIDError
	}

	reader, err := b.c.TransformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("/rest/agile/%v/board/%v/issue", b.version, boardId)

	request, err := b.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := b.c.Call(request, nil)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (b BoardService) Projects(ctx context.Context, boardId, startAt, maxResults int) (*model.BoardProjectPageScheme, *model.ResponseScheme, error) {

	if boardId == 0 {
		return nil, nil, model.ErrNoBoardIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	endpoint := fmt.Sprintf("/rest/agile/%v/board/%v/project?%v", b.version, boardId, params.Encode())

	request, err := b.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	var projects model.BoardProjectPageScheme
	response, err := b.c.Call(request, &projects)
	if err != nil {
		return nil, response, err
	}

	return &projects, response, nil
}

func (b BoardService) Sprints(ctx context.Context, boardId, startAt, maxResults int, states []string) (*model.BoardSprintPageScheme, *model.ResponseScheme, error) {

	if boardId == 0 {
		return nil, nil, model.ErrNoBoardIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("state", strings.Join(states, ","))

	endpoint := fmt.Sprintf("/rest/agile/%v/board/%v/sprint?%v", b.version, boardId, params.Encode())

	request, err := b.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	var sprints model.BoardSprintPageScheme
	response, err := b.c.Call(request, &sprints)
	if err != nil {
		return nil, response, err
	}

	return &sprints, response, nil
}

func (b BoardService) IssuesBySprint(ctx context.Context, boardId, sprintId, startAt, maxResults int, opts *model.IssueOptionScheme) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {

	if boardId == 0 {
		return nil, nil, model.ErrNoBoardIDError
	}

	if sprintId == 0 {
		return nil, nil, model.ErrNoSprintIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if !opts.ValidateQuery {
			params.Add("validateQuery ", "false")
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

	endpoint := fmt.Sprintf("/rest/agile/%v/board/%v/sprint/%v/issue?%v", b.version, boardId, sprintId, params.Encode())

	request, err := b.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	var issues model.BoardIssuePageScheme
	response, err := b.c.Call(request, &issues)
	if err != nil {
		return nil, response, err
	}

	return &issues, response, nil
}

func (b BoardService) Versions(ctx context.Context, boardID, startAt, maxResults int, released bool) (*model.BoardVersionPageScheme, *model.ResponseScheme, error) {

	if boardID == 0 {
		return nil, nil, model.ErrNoBoardIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("released", fmt.Sprintf("%t", released))

	endpoint := fmt.Sprintf("/rest/agile/%v/board/%v/version?%v", b.version, boardID, params.Encode())

	request, err := b.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	var versions model.BoardVersionPageScheme
	response, err := b.c.Call(request, &versions)
	if err != nil {
		return nil, response, err
	}

	return &versions, response, nil
}

func (b BoardService) Delete(ctx context.Context, boardId int) (*model.ResponseScheme, error) {

	if boardId == 0 {
		return nil, model.ErrNoBoardIDError
	}

	endpoint := fmt.Sprintf("/rest/agile/%v/board/%v", b.version, boardId)

	request, err := b.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	response, err := b.c.Call(request, nil)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (b BoardService) Gets(ctx context.Context, opts *model.GetBoardsOptions, startAt, maxResults int) (*model.BoardPageScheme, *model.ResponseScheme, error) {

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

	endpoint := fmt.Sprintf("/rest/agile/%v/board?%v", b.version, params.Encode())

	request, err := b.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	var boards model.BoardPageScheme
	response, err := b.c.Call(request, &boards)
	if err != nil {
		return nil, response, err
	}

	return &boards, response, nil
}
