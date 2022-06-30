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

func NewSprintService(client service.Client, version string) (agile.Sprint, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &SprintService{client, version}, nil
}

type SprintService struct {
	c       service.Client
	version string
}

func (s SprintService) Get(ctx context.Context, sprintId int) (*model.SprintScheme, *model.ResponseScheme, error) {

	if sprintId == 0 {
		return nil, nil, model.ErrNoSprintIDError
	}

	endpoint := fmt.Sprintf("rest/agile/%v/sprint/%v", s.version, sprintId)

	request, err := s.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var sprint model.SprintScheme
	response, err := s.c.Call(request, &sprint)
	if err != nil {
		return nil, response, err
	}

	return &sprint, response, nil
}

func (s SprintService) Create(ctx context.Context, payload *model.SprintPayloadScheme) (*model.SprintScheme, *model.ResponseScheme,
	error) {

	reader, err := s.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/agile/%v/sprint", s.version)

	request, err := s.c.NewJsonRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	var sprint model.SprintScheme
	response, err := s.c.Call(request, &sprint)
	if err != nil {
		return nil, response, err
	}

	return &sprint, response, nil
}

func (s SprintService) Update(ctx context.Context, sprintId int, payload *model.SprintPayloadScheme) (*model.SprintScheme,
	*model.ResponseScheme, error) {

	reader, err := s.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/agile/%v/sprint/%v", s.version, sprintId)

	request, err := s.c.NewJsonRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	var sprint model.SprintScheme
	response, err := s.c.Call(request, &sprint)
	if err != nil {
		return nil, response, err
	}

	return &sprint, response, nil
}

func (s SprintService) Path(ctx context.Context, sprintId int, payload *model.SprintPayloadScheme) (*model.SprintScheme,
	*model.ResponseScheme, error) {

	reader, err := s.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/agile/%v/sprint/%v", s.version, sprintId)

	request, err := s.c.NewJsonRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	var sprint model.SprintScheme
	response, err := s.c.Call(request, &sprint)
	if err != nil {
		return nil, response, err
	}

	return &sprint, response, nil
}

func (s SprintService) Delete(ctx context.Context, sprintId int) (*model.ResponseScheme, error) {

	if sprintId == 0 {
		return nil, model.ErrNoSprintIDError
	}

	endpoint := fmt.Sprintf("rest/agile/%v/sprint/%v", s.version, sprintId)

	request, err := s.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	response, err := s.c.Call(request, nil)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (s SprintService) Issues(ctx context.Context, sprintId int, opts *model.IssueOptionScheme, startAt, maxResults int) (
	*model.SprintIssuePageScheme, *model.ResponseScheme, error) {

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

	endpoint := fmt.Sprintf("rest/agile/%v/sprint/%v/issue?%v", s.version, sprintId, params.Encode())

	request, err := s.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var issues model.SprintIssuePageScheme
	response, err := s.c.Call(request, &issues)
	if err != nil {
		return nil, response, err
	}

	return &issues, response, nil
}

func (s SprintService) Start(ctx context.Context, sprintId int) (*model.ResponseScheme, error) {

	if sprintId == 0 {
		return nil, model.ErrNoSprintIDError
	}

	payload := model.SprintPayloadScheme{
		State: "Active",
	}

	reader, err := s.c.TransformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/agile/%v/sprint/%v", s.version, sprintId)

	request, err := s.c.NewJsonRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, err
	}

	response, err := s.c.Call(request, nil)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (s SprintService) Close(ctx context.Context, sprintId int) (*model.ResponseScheme, error) {

	if sprintId == 0 {
		return nil, model.ErrNoSprintIDError
	}

	payload := model.SprintPayloadScheme{
		State: "Closed",
	}

	reader, err := s.c.TransformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/agile/%v/sprint/%v", s.version, sprintId)

	request, err := s.c.NewJsonRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, err
	}

	response, err := s.c.Call(request, nil)
	if err != nil {
		return response, err
	}

	return response, nil
}
