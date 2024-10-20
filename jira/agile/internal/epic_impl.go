package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/agile"
)

// NewEpicService creates a new instance of EpicService.
// It takes a service.Connector and a version string as input and returns a pointer to EpicService.
func NewEpicService(client service.Connector, version string) *EpicService {
	return &EpicService{
		internalClient: &internalEpicImpl{c: client, version: version},
	}
}

// EpicService provides methods to interact with epic operations in Jira Agile.
type EpicService struct {
	// internalClient is the connector interface for epic operations.
	internalClient agile.EpicConnector
}

// Get returns the epic for a given epic ID.
//
// This epic will only be returned if the user has permission to view it.
//
// Note: This operation does not work for epics in next-gen projects.
//
// GET /rest/agile/1.0/epic/{epicIDOrKey}
//
// https://docs.go-atlassian.io/jira-agile/epics#get-epic
func (e *EpicService) Get(ctx context.Context, epicIDOrKey string) (*model.EpicScheme, *model.ResponseScheme, error) {
	return e.internalClient.Get(ctx, epicIDOrKey)
}

// Issues returns all issues that belong to the epic, for the given epic ID.
//
// This only includes issues that the user has permission to view.
//
// Issues returned from this resource include Agile fields, like sprint, closedSprints,  flagged, and epic.
//
// By default, the returned issues are ordered by rank.
//
// GET /rest/agile/1.0/epic/{epicIDOrKey}/issue
//
// https://docs.go-atlassian.io/jira-agile/epics#get-issues-for-epic
func (e *EpicService) Issues(ctx context.Context, epicIDOrKey string, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {
	return e.internalClient.Issues(ctx, epicIDOrKey, opts, startAt, maxResults)
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
// POST /rest/agile/1.0/epic/{epicIDOrKey}/issue
//
// https://docs.go-atlassian.io/jira-agile/epics#move-issues-to-epic
func (e *EpicService) Move(ctx context.Context, epicIDOrKey string, issues []string) (*model.ResponseScheme, error) {
	return e.internalClient.Move(ctx, epicIDOrKey, issues)
}

type internalEpicImpl struct {
	c       service.Connector
	version string
}

func (i *internalEpicImpl) Get(ctx context.Context, epicIDOrKey string) (*model.EpicScheme, *model.ResponseScheme, error) {

	if epicIDOrKey == "" {
		return nil, nil, model.ErrNoEpicID
	}

	url := fmt.Sprintf("rest/agile/%v/epic/%v", i.version, epicIDOrKey)

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	epic := new(model.EpicScheme)
	res, err := i.c.Call(req, epic)
	if err != nil {
		return nil, res, err
	}

	return epic, res, nil
}

func (i *internalEpicImpl) Issues(ctx context.Context, epicIDOrKey string, opts *model.IssueOptionScheme, startAt, maxResults int) (*model.BoardIssuePageScheme, *model.ResponseScheme, error) {

	if epicIDOrKey == "" {
		return nil, nil, model.ErrNoEpicID
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

	url := fmt.Sprintf("rest/agile/%v/epic/%v/issue?%v", i.version, epicIDOrKey, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, url, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.BoardIssuePageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalEpicImpl) Move(ctx context.Context, epicIDOrKey string, issues []string) (*model.ResponseScheme, error) {

	if epicIDOrKey == "" {
		return nil, model.ErrNoEpicID
	}

	payload := map[string]interface{}{"issues": issues}
	url := fmt.Sprintf("rest/agile/%v/epic/%v/issue", i.version, epicIDOrKey)

	req, err := i.c.NewRequest(ctx, http.MethodPost, url, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}
