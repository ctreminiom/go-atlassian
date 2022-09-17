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

func NewEpicService(client service.Client, version string) (*EpicService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &EpicService{
		internalClient: &internalEpicImpl{c: client, version: version},
	}, nil
}

type EpicService struct {
	internalClient agile.EpicConnector
}

// Get returns the epic for a given epic ID.
//
// This epic will only be returned if the user has permission to view it.
//
// Note: This operation does not work for epics in next-gen projects.
//
// GET /rest/agile/1.0/epic/{epicIdOrKey}
//
// https://docs.go-atlassian.io/jira-agile/epics#get-epic
func (e *EpicService) Get(ctx context.Context, epicIdOrKey string) (*model.EpicScheme, *model.ResponseScheme, error) {
	return e.internalClient.Get(ctx, epicIdOrKey)
}

// Issues returns all issues that belong to the epic, for the given epic ID.
//
// This only includes issues that the user has permission to view.
//
// Issues returned from this resource include Agile fields, like sprint, closedSprints,  flagged, and epic.
//
// By default, the returned issues are ordered by rank.
//
// GET /rest/agile/1.0/epic/{epicIdOrKey}/issue
//
// https://docs.go-atlassian.io/jira-agile/epics#get-issues-for-epic
func (e *EpicService) Issues(ctx context.Context, epicIdOrKey string, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {
	return e.internalClient.Issues(ctx, epicIdOrKey, opts, startAt, maxResults)
}

// Move moves issues to an epic, for a given epic id.
//
// Issues can be only in a single epic at the same time.
// That means that already assigned issues to an epic, will not be assigned to the previous epic anymore.
//
// The user needs to have the edit issue permission for all issue they want to move and to the epic.
//
// The maximum number of issues that can be moved in one operation is 50.
//
// POST /rest/agile/1.0/epic/{epicIdOrKey}/issue
//
// https://docs.go-atlassian.io/jira-agile/epics#move-issues-to-epic
func (e *EpicService) Move(ctx context.Context, epicIdOrKey string, issues []string) (*model.ResponseScheme, error) {
	return e.internalClient.Move(ctx, epicIdOrKey, issues)
}

type internalEpicImpl struct {
	c       service.Client
	version string
}

func (i *internalEpicImpl) Get(ctx context.Context, epicIdOrKey string) (*model.EpicScheme, *model.ResponseScheme, error) {

	if epicIdOrKey == "" {
		return nil, nil, model.ErrNoEpicIDError
	}

	endpoint := fmt.Sprintf("rest/agile/%v/epic/%v", i.version, epicIdOrKey)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	epic := new(model.EpicScheme)
	response, err := i.c.Call(request, epic)
	if err != nil {
		return nil, response, err
	}

	return epic, response, nil
}

func (i *internalEpicImpl) Issues(ctx context.Context, epicIdOrKey string, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {

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

	endpoint := fmt.Sprintf("rest/agile/%v/epic/%v/issue?%v", i.version, epicIdOrKey, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BoardIssuePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalEpicImpl) Move(ctx context.Context, epicIdOrKey string, issues []string) (*model.ResponseScheme, error) {

	if epicIdOrKey == "" {
		return nil, model.ErrNoEpicIDError
	}

	reader, err := i.c.TransformStructToReader(map[string]interface{}{"issues": issues})
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/agile/%v/epic/%v/issue", i.version, epicIdOrKey)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
