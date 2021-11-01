package v3

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
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

// Gets returns a paginated list of all versions in a project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-project-versions-paginated
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-project-projectidorkey-version-get
func (p *ProjectVersionService) Gets(ctx context.Context, projectKeyOrID string, options *ProjectVersionGetsOptions, startAt,
	maxResults int) (result *ProjectVersionPageScheme, response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, notProjectIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
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

	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/version?%v", projectKeyOrID, params.Encode())

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
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

// Create creates a project version.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/versions#create-version
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-version-post
func (p *ProjectVersionService) Create(ctx context.Context, payload *ProjectVersionPayloadScheme) (
	result *ProjectVersionScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/version"

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns a project version.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-version
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-version-id-get
func (p *ProjectVersionService) Get(ctx context.Context, versionID string, expand []string) (result *ProjectVersionScheme,
	response *ResponseScheme, err error) {

	if len(versionID) == 0 {
		return nil, nil, notVersionIDError
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/3/version/%v", versionID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update updates a project version.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/versions#update-version
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-version-id-put
func (p *ProjectVersionService) Update(ctx context.Context, versionID string, payload *ProjectVersionPayloadScheme) (
	result *ProjectVersionScheme, response *ResponseScheme, err error) {

	if len(versionID) == 0 {
		return nil, nil, notVersionIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/version/%v", versionID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := p.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Merge merges two project versions.
// The merge is completed by deleting the version specified in id and replacing any occurrences of its ID in fixVersion with the version ID specified in moveIssuesTo.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/versions#merge-versions
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-version-id-mergeto-moveissuesto-put
func (p *ProjectVersionService) Merge(ctx context.Context, versionID, versionMoveIssuesTo string) (response *ResponseScheme, err error) {

	if len(versionID) == 0 {
		return nil, notVersionIDError
	}

	if len(versionMoveIssuesTo) == 0 {
		return nil, notVersionIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/version/%v/mergeto/%v", versionID, versionMoveIssuesTo)

	request, err := p.client.newRequest(ctx, http.MethodPut, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, nil)
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

// RelatedIssueCounts returns the following counts for a version:
// 1. Number of issues where the fixVersion is set to the version.
// 2. Number of issues where the affectedVersion is set to the version.
// 3. Number of issues where a version custom field is set to the version.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-versions-related-issues-count
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-version-id-relatedissuecounts-get
func (p *ProjectVersionService) RelatedIssueCounts(ctx context.Context, versionID string) (result *VersionIssueCountsScheme,
	response *ResponseScheme, err error) {

	if len(versionID) == 0 {
		return nil, nil, notVersionIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/version/%v/relatedIssueCounts", versionID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type VersionUnresolvedIssuesCountScheme struct {
	Self                  string `json:"self"`
	IssuesUnresolvedCount int    `json:"issuesUnresolvedCount"`
	IssuesCount           int    `json:"issuesCount"`
}

// UnresolvedIssueCount returns counts of the issues and unresolved issues for the project version.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-versions-unresolved-issues-count
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-project-versions/#api-rest-api-3-version-id-unresolvedissuecount-get
func (p *ProjectVersionService) UnresolvedIssueCount(ctx context.Context, versionID string) (
	result *VersionUnresolvedIssuesCountScheme, response *ResponseScheme, err error) {

	if len(versionID) == 0 {
		return nil, nil, notVersionIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/version/%v/unresolvedIssueCount", versionID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

var (
	notVersionIDError = fmt.Errorf("error, please provide a valid versionID value")
)
