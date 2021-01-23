package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type ProjectVersionService struct{ client *Client }

type ProjectVersionGetsOptions struct {
	OrderBy string
	Query   string
	Status  string
	Expand  []string
}

type ProjectVersionPageScheme struct {
	Self       string                 `json:"self,omitempty"`
	NextPage   string                 `json:"nextPage,omitempty"`
	MaxResults int                    `json:"maxResults,omitempty"`
	StartAt    int                    `json:"startAt,omitempty"`
	Total      int                    `json:"total,omitempty"`
	IsLast     bool                   `json:"isLast,omitempty"`
	Values     []ProjectVersionScheme `json:"values,omitempty"`
}

type ProjectVersionScheme struct {
	Self                      string `json:"self,omitempty"`
	ID                        string `json:"id,omitempty"`
	Description               string `json:"description,omitempty"`
	Name                      string `json:"name,omitempty"`
	Archived                  bool   `json:"archived,omitempty"`
	Released                  bool   `json:"released,omitempty"`
	ReleaseDate               string `json:"releaseDate,omitempty"`
	Overdue                   bool   `json:"overdue,omitempty"`
	UserReleaseDate           string `json:"userReleaseDate,omitempty"`
	ProjectID                 int    `json:"projectId,omitempty"`
	IssuesStatusForFixVersion struct {
		Unmapped   int `json:"unmapped,omitempty"`
		ToDo       int `json:"toDo,omitempty"`
		InProgress int `json:"inProgress,omitempty"`
		Done       int `json:"done,omitempty"`
	} `json:"issuesStatusForFixVersion,omitempty"`
}

// Returns a paginated list of all versions in a project.
// See the Get project versions resource if you want to get a full list of versions without pagination.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-project-projectidorkey-version-get
func (p *ProjectVersionService) Gets(ctx context.Context, projectKeyOrID string, options *ProjectVersionGetsOptions, startAt, maxResults int) (result *ProjectVersionPageScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var expand string
	for index, value := range options.Expand {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	if len(options.Query) != 0 {
		params.Add("query", options.Query)
	}

	if len(options.Status) != 0 {
		params.Add("status", options.Status)
	}

	if len(options.OrderBy) != 0 {
		params.Add("orderBy", options.OrderBy)
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/version?%v", projectKeyOrID, params.Encode())

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectVersionPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Creates a project version.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-version-post
func (p *ProjectVersionService) Create(ctx context.Context, payload *ProjectVersionScheme) (result *ProjectVersionScheme, response *Response, err error) {

	var endpoint = "rest/api/3/version"

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectVersionScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns a project version.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-version-id-get
func (p *ProjectVersionService) Get(ctx context.Context, versionID string) (result *ProjectVersionScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/version/%v", versionID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectVersionScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Updates a project version.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-version-id-put
func (p *ProjectVersionService) Update(ctx context.Context, versionID string, payload *ProjectVersionScheme) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/version/%v", versionID)

	request, err := p.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	return
}
