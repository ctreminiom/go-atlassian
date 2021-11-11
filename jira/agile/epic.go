package agile

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models/agile"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type EpicService struct{ client *Client }

// Get returns the epic for a given epic ID.
// This epic will only be returned if the user has permission to view it.
// Note: This operation does not work for epics in next-gen projects.
func (e *EpicService) Get(ctx context.Context, epicIDOrKey string) (result *model.EpicScheme, response *ResponseScheme, err error) {

	if len(epicIDOrKey) == 0 {
		return nil, nil, model.ErrNoEpicIDError
	}

	var endpoint = fmt.Sprintf("/rest/agile/1.0/epic/%v", epicIDOrKey)

	request, err := e.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = e.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Issues returns all issues that belong to the epic, for the given epic ID.
// This only includes issues that the user has permission to view.
// Issues returned from this resource include Agile fields, like sprint, closedSprints,
// flagged, and epic.
// By default, the returned issues are ordered by rank.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/software/rest/api-group-epic/#api-agile-1-0-epic-epicidorkey-issue-get
// Library Docs: N/A
func (e *EpicService) Issues(ctx context.Context, epicIDOrKey string, startAt, maxResults int,
	opts *model.IssueOptionScheme) (result *model.BoardIssuePageScheme, response *ResponseScheme, err error) {

	if len(epicIDOrKey) == 0 {
		return nil, nil, model.ErrNoEpicIDError
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

	var endpoint = fmt.Sprintf("/rest/agile/1.0/epic/%v/issue?%v", epicIDOrKey, params.Encode())

	request, err := e.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = e.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}
