package agile

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type EpicService struct{ client *Client }

type EpicScheme struct {
	ID      int              `json:"id,omitempty"`
	Key     string           `json:"key,omitempty"`
	Self    string           `json:"self,omitempty"`
	Name    string           `json:"name,omitempty"`
	Summary string           `json:"summary,omitempty"`
	Color   *EpicColorScheme `json:"color,omitempty"`
	Done    bool             `json:"done,omitempty"`
}
type EpicColorScheme struct {
	Key string `json:"key,omitempty"`
}

// Get returns the epic for a given epic ID.
// This epic will only be returned if the user has permission to view it.
// Note: This operation does not work for epics in next-gen projects.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/software/rest/api-group-epic/#api-agile-1-0-epic-epicidorkey-get
// Library Docs: N/A
func (e *EpicService) Get(ctx context.Context, epicIDOrKey string) (result *EpicScheme, response *Response, err error) {

	if len(epicIDOrKey) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid epicIDOrKey value")
	}

	var endpoint = fmt.Sprintf("/rest/agile/1.0/epic/%v", epicIDOrKey)

	request, err := e.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = e.client.Do(request)
	if err != nil {
		return
	}

	result = new(EpicScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
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
func (e *EpicService) Issues(ctx context.Context, epicIDOrKey string, startAt, maxResults int, opts *IssueOptionScheme) (result *BoardIssuePageScheme, response *Response, err error) {

	if len(epicIDOrKey) == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid epicIDOrKey value")
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

		var expand string
		for index, value := range opts.Expand {

			if index == 0 {
				expand = value
				continue
			}

			expand += "," + value
		}

		if len(expand) != 0 {
			params.Add("expand", expand)
		}

		var fieldFormatted string
		for index, value := range opts.Fields {

			if index == 0 {
				fieldFormatted = value
				continue
			}
			fieldFormatted += "," + value
		}

		if len(fieldFormatted) != 0 {
			params.Add("fields", fieldFormatted)
		}

	}

	var endpoint = fmt.Sprintf("/rest/agile/1.0/epic/%v/issue?%v", epicIDOrKey, params.Encode())

	request, err := e.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = e.client.Do(request)
	if err != nil {
		return
	}

	result = new(BoardIssuePageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
