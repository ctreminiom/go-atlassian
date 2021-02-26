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
	Self       string                  `json:"self,omitempty"`
	NextPage   string                  `json:"nextPage,omitempty"`
	MaxResults int                     `json:"maxResults,omitempty"`
	StartAt    int                     `json:"startAt,omitempty"`
	Total      int                     `json:"total,omitempty"`
	IsLast     bool                    `json:"isLast,omitempty"`
	Values     []*ProjectVersionScheme `json:"values,omitempty"`
}

type ProjectVersionScheme struct {
	Self                      string                                         `json:"self,omitempty"`
	ID                        string                                         `json:"id,omitempty"`
	Description               string                                         `json:"description,omitempty"`
	Name                      string                                         `json:"name,omitempty"`
	Archived                  bool                                           `json:"archived,omitempty"`
	Released                  bool                                           `json:"released,omitempty"`
	ReleaseDate               string                                         `json:"releaseDate,omitempty"`
	Overdue                   bool                                           `json:"overdue,omitempty"`
	UserReleaseDate           string                                         `json:"userReleaseDate,omitempty"`
	ProjectID                 int                                            `json:"projectId,omitempty"`
	Operations                []*ProjectVersionOperation                     `json:"operations,omitempty"`
	IssuesStatusForFixVersion *ProjectVersionIssuesStatusForFixVersionScheme `json:"issuesStatusForFixVersion,omitempty"`
}

type ProjectVersionOperation struct {
	ID         string `json:"id,omitempty"`
	StyleClass string `json:"styleClass,omitempty"`
	Label      string `json:"label,omitempty"`
	Href       string `json:"href,omitempty"`
	Weight     int    `json:"weight,omitempty"`
}

type ProjectVersionIssuesStatusForFixVersionScheme struct {
	Unmapped   int `json:"unmapped,omitempty"`
	ToDo       int `json:"toDo,omitempty"`
	InProgress int `json:"inProgress,omitempty"`
	Done       int `json:"done,omitempty"`
}

// Returns a paginated list of all versions in a project.
// See the Get project versions resource if you want to get a full list of versions without pagination.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-project-projectidorkey-version-get
func (p *ProjectVersionService) Gets(ctx context.Context, projectKeyOrID string, options *ProjectVersionGetsOptions, startAt, maxResults int) (result *ProjectVersionPageScheme, response *Response, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid projectKeyOrID value")
	}

	if options == nil {
		return nil, nil, fmt.Errorf("error, please provide a valid ProjectVersionGetsOptions pointer")
	}

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

type ProjectVersionPayloadScheme struct {
	Archived    bool   `json:"archived,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ProjectID   int    `json:"projectId,omitempty"`
	Released    bool   `json:"released,omitempty"`
	StartDate   string `json:"startDate,omitempty"`
}

// Creates a project version.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-version-post
func (p *ProjectVersionService) Create(ctx context.Context, payload *ProjectVersionPayloadScheme) (result *ProjectVersionScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, please provide a valid ProjectVersionPayloadScheme pointer")
	}

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
func (p *ProjectVersionService) Get(ctx context.Context, versionID string, expands []string) (result *ProjectVersionScheme, response *Response, err error) {

	params := url.Values{}

	var expand string
	for index, value := range expands {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	var endpoint string
	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("rest/api/3/version/%v?%v", versionID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/api/3/version/%v", versionID)
	}

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
func (p *ProjectVersionService) Update(ctx context.Context, versionID string, payload *ProjectVersionPayloadScheme) (result *ProjectVersionScheme, response *Response, err error) {

	if len(versionID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid versionID value")
	}

	if payload == nil {
		return nil, nil, fmt.Errorf("error, please provide a valid ProjectVersionPayloadScheme pointer")
	}

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

	result = new(ProjectVersionScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Merges two project versions.
// The merge is completed by deleting the version specified in id and replacing any occurrences of its ID in fixVersion with the version ID specified in moveIssuesTo.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-version-id-mergeto-moveissuesto-put
func (p *ProjectVersionService) Merge(ctx context.Context, versionID, moveIssuesTo string) (response *Response, err error) {

	if len(versionID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid versionID value")
	}

	if len(moveIssuesTo) == 0 {
		return nil, fmt.Errorf("error, please provide a valid moveIssuesTo value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/version/%v/mergeto/%v", versionID, moveIssuesTo)

	request, err := p.client.newRequest(ctx, http.MethodPut, endpoint, nil)
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

type VersionIssueCountsScheme struct {
	Self                                     string                                     `json:"self,omitempty"`
	IssuesFixedCount                         int                                        `json:"issuesFixedCount,omitempty"`
	IssuesAffectedCount                      int                                        `json:"issuesAffectedCount,omitempty"`
	IssueCountWithCustomFieldsShowingVersion int                                        `json:"issueCountWithCustomFieldsShowingVersion,omitempty"`
	CustomFieldUsage                         []*VersionIssueCountCustomFieldUsageScheme `json:"customFieldUsage,omitempty"`
}

type VersionIssueCountCustomFieldUsageScheme struct {
	FieldName                          string `json:"fieldName,omitempty"`
	CustomFieldID                      int    `json:"customFieldId,omitempty"`
	IssueCountWithVersionInCustomField int    `json:"issueCountWithVersionInCustomField,omitempty"`
}

// Returns the following counts for a version:
// 1. Number of issues where the fixVersion is set to the version.
// 2. Number of issues where the affectedVersion is set to the version.
// 3. Number of issues where a version custom field is set to the version.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-version-id-relatedissuecounts-get
func (p *ProjectVersionService) RelatedIssueCounts(ctx context.Context, versionID string) (result *VersionIssueCountsScheme, response *Response, err error) {

	if len(versionID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid versionID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/version/%v/relatedIssueCounts", versionID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(VersionIssueCountsScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type VersionUnresolvedIssuesCountScheme struct {
	Self                  string `json:"self"`
	IssuesUnresolvedCount int    `json:"issuesUnresolvedCount"`
	IssuesCount           int    `json:"issuesCount"`
}

// Returns counts of the issues and unresolved issues for the project version.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-version-id-unresolvedissuecount-get
func (p *ProjectVersionService) UnresolvedIssueCount(ctx context.Context, versionID string) (result *VersionUnresolvedIssuesCountScheme, response *Response, err error) {

	if len(versionID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid versionID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/version/%v/unresolvedIssueCount", versionID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(VersionUnresolvedIssuesCountScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
