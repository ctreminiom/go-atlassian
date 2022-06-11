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

func NewEpicService(client service.Client, version string) (agile.Epic, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &EpicService{client, version}, nil
}

type EpicService struct {
	c       service.Client
	version string
}

func (e EpicService) Get(ctx context.Context, epicIdOrKey string) (*model.EpicScheme, *model.ResponseScheme, error) {

	if epicIdOrKey == "" {
		return nil, nil, model.ErrNoEpicIDError
	}

	endpoint := fmt.Sprintf("/rest/agile/%v/epic/%v", e.version, epicIdOrKey)

	request, err := e.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var epic model.EpicScheme
	response, err := e.c.Call(request, &epic)
	if err != nil {
		return nil, response, err
	}

	return &epic, response, nil
}

func (e EpicService) Issues(ctx context.Context, epicIdOrKey string, startAt, maxResults int, opts *model.IssueOptionScheme) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {

	if epicIdOrKey == "" {
		return nil, nil, model.ErrNoEpicIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		params.Add("validateQuery", fmt.Sprintf("%t", opts.ValidateQuery))

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

	endpoint := fmt.Sprintf("/rest/agile/%v/epic/%v/issue?%v", e.version, epicIdOrKey, params.Encode())

	request, err := e.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var issues model.BoardIssuePageScheme
	response, err := e.c.Call(request, &issues)
	if err != nil {
		return nil, response, err
	}

	return &issues, response, nil
}

func (e EpicService) Move(ctx context.Context, epicIdOrKey string, issues []string) (*model.ResponseScheme, error) {

	if epicIdOrKey == "" {
		return nil, model.ErrNoEpicIDError
	}

	reader, err := e.c.TransformStructToReader(map[string]interface{}{"issues": issues})
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("/rest/agile/%v/epic/%v/issue", e.version, epicIdOrKey)

	request, err := e.c.NewJsonRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, err
	}

	response, err := e.c.Call(request, nil)
	if err != nil {
		return response, err
	}

	return response, nil
}
