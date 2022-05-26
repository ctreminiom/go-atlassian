package agile

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service/agile"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func newEpicService(client *Client, version string) (agile.Epic, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &EpicService{client, version}, nil
}

type EpicService struct {
	c       *Client
	version string
}

func (e EpicService) Get(ctx context.Context, epic string) (*model.EpicScheme, *model.ResponseScheme, error) {

	if epic == "" {
		return nil, nil, model.ErrNoEpicIDError
	}

	endpoint := fmt.Sprintf("/rest/agile/%v/epic/%v", e.version, epic)

	request, err := e.c.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")
	var epics model.EpicScheme

	response, err := e.c.call(request, &epics)
	if err != nil {
		return nil, response, err
	}

	return &epics, response, nil
}

func (e EpicService) Issues(ctx context.Context, epic string, startAt, maxResults int, opts *model.IssueOptionScheme) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {

	if epic == "" {
		return nil, nil, model.ErrNoEpicIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("validateQuery", fmt.Sprintf("%t", opts.ValidateQuery))

	if opts != nil {

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

	endpoint := fmt.Sprintf("/rest/agile/%v/epic/%v/issue?%v", e.version, epic, params.Encode())

	request, err := e.c.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")
	var issues model.BoardIssuePageScheme

	response, err := e.c.call(request, &issues)
	if err != nil {
		return nil, response, err
	}

	return &issues, response, nil
}

func (e EpicService) Move(ctx context.Context, epic string, issues []string) (*model.ResponseScheme, error) {

	if epic == "" {
		return nil, model.ErrNoEpicIDError
	}

	reader, err := e.c.transformStructToReader(map[string]interface{}{"issues": issues})
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/agile/%v/epic/%v/issue", e.version, epic)

	request, err := e.c.newRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	response, err := e.c.call(request, nil)
	if err != nil {
		return response, err
	}

	return response, nil
}
