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
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewProjectVersionService creates a new instance of ProjectVersionService.
func NewProjectVersionService(client service.Connector, version string) (*ProjectVersionService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ProjectVersionService{
		internalClient: &internalProjectVersionImpl{c: client, version: version},
	}, nil
}

// ProjectVersionService provides methods to manage project versions in Jira Service Management.
type ProjectVersionService struct {
	// internalClient is the connector interface for project version operations.
	internalClient jira.ProjectVersionConnector
}

// Gets returns all versions in a project.
//
// The response is not paginated.
//
// Use Search() if you want to get the versions in a project with pagination.
//
// GET /rest/api/{2-3}/project/{projectKeyOrID}/versions
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-project-versions
func (p *ProjectVersionService) Gets(ctx context.Context, projectKeyOrID string) ([]*model.VersionScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx, projectKeyOrID)
}

// Search returns a paginated list of all versions in a project.
//
// GET /rest/api/{2-3}/project/{projectKeyOrID}/version
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-project-versions-paginated
func (p *ProjectVersionService) Search(ctx context.Context, projectKeyOrID string, options *model.VersionGetsOptions, startAt, maxResults int) (*model.VersionPageScheme, *model.ResponseScheme, error) {
	return p.internalClient.Search(ctx, projectKeyOrID, options, startAt, maxResults)
}

// Create creates a project version.
//
// POST /rest/api/{2-3}/version
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#create-version
func (p *ProjectVersionService) Create(ctx context.Context, payload *model.VersionPayloadScheme) (*model.VersionScheme, *model.ResponseScheme, error) {
	return p.internalClient.Create(ctx, payload)
}

// Get returns a project version.
//
// GET /rest/api/{2-3}/version/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-version
func (p *ProjectVersionService) Get(ctx context.Context, versionID string, expand []string) (*model.VersionScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, versionID, expand)
}

// Update updates a project version.
//
// PUT /rest/api/{2-3}/version/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#update-version
func (p *ProjectVersionService) Update(ctx context.Context, versionID string, payload *model.VersionPayloadScheme) (*model.VersionScheme, *model.ResponseScheme, error) {
	return p.internalClient.Update(ctx, versionID, payload)
}

// Merge merges two project versions.
//
// # The merge is completed by deleting the version specified in id and replacing any occurrences of
//
// its ID in fixVersion with the version ID specified in moveIssuesTo.
//
// PUT /rest/api/{2-3}/version/{id}/mergeto/{moveIssuesTo}
func (p *ProjectVersionService) Merge(ctx context.Context, versionID, versionMoveIssuesTo string) (*model.ResponseScheme, error) {
	return p.internalClient.Merge(ctx, versionID, versionMoveIssuesTo)
}

// RelatedIssueCounts returns the following counts for a version:
//
// 1. Number of issues where the fixVersion is set to the version.
//
// 2. Number of issues where the affectedVersion is set to the version.
//
// 3. Number of issues where a version custom field is set to the version.
//
// GET /rest/api/{2-3}/version/{id}/relatedIssueCounts
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-versions-related-issues-count
func (p *ProjectVersionService) RelatedIssueCounts(ctx context.Context, versionID string) (*model.VersionIssueCountsScheme, *model.ResponseScheme, error) {
	return p.internalClient.RelatedIssueCounts(ctx, versionID)
}

// UnresolvedIssueCount returns counts of the issues and unresolved issues for the project version.
//
// GET /rest/api/{2-3}/version/{id}/unresolvedIssueCount
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-versions-unresolved-issues-count
func (p *ProjectVersionService) UnresolvedIssueCount(ctx context.Context, versionID string) (*model.VersionUnresolvedIssuesCountScheme, *model.ResponseScheme, error) {
	return p.internalClient.UnresolvedIssueCount(ctx, versionID)
}

type internalProjectVersionImpl struct {
	c       service.Connector
	version string
}

func (i *internalProjectVersionImpl) Gets(ctx context.Context, projectKeyOrID string) ([]*model.VersionScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/versions", i.version, projectKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var versions []*model.VersionScheme
	response, err := i.c.Call(request, &versions)
	if err != nil {
		return nil, response, err
	}

	return versions, response, nil
}

func (i *internalProjectVersionImpl) Search(ctx context.Context, projectKeyOrID string, options *model.VersionGetsOptions, startAt, maxResults int) (*model.VersionPageScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
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

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/version?%v", i.version, projectKeyOrID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.VersionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalProjectVersionImpl) Create(ctx context.Context, payload *model.VersionPayloadScheme) (*model.VersionScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/version", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	version := new(model.VersionScheme)
	response, err := i.c.Call(request, version)
	if err != nil {
		return nil, response, err
	}

	return version, response, nil
}

func (i *internalProjectVersionImpl) Get(ctx context.Context, versionID string, expand []string) (*model.VersionScheme, *model.ResponseScheme, error) {

	if versionID == "" {
		return nil, nil, model.ErrNoVersionID
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/version/%v", i.version, versionID))

	if expand != nil {

		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	version := new(model.VersionScheme)
	response, err := i.c.Call(request, version)
	if err != nil {
		return nil, response, err
	}

	return version, response, nil
}

func (i *internalProjectVersionImpl) Update(ctx context.Context, versionID string, payload *model.VersionPayloadScheme) (*model.VersionScheme, *model.ResponseScheme, error) {

	if versionID == "" {
		return nil, nil, model.ErrNoVersionID
	}

	endpoint := fmt.Sprintf("rest/api/%v/version/%v", i.version, versionID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	version := new(model.VersionScheme)
	response, err := i.c.Call(request, version)
	if err != nil {
		return nil, response, err
	}

	return version, response, nil
}

func (i *internalProjectVersionImpl) Merge(ctx context.Context, versionID, versionMoveIssuesTo string) (*model.ResponseScheme, error) {

	if versionID == "" {
		return nil, model.ErrNoVersionID
	}

	if versionMoveIssuesTo == "" {
		return nil, model.ErrNoVersionID
	}

	endpoint := fmt.Sprintf("rest/api/%v/version/%v/mergeto/%v", i.version, versionID, versionMoveIssuesTo)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalProjectVersionImpl) RelatedIssueCounts(ctx context.Context, versionID string) (*model.VersionIssueCountsScheme, *model.ResponseScheme, error) {

	if versionID == "" {
		return nil, nil, model.ErrNoVersionID
	}

	endpoint := fmt.Sprintf("rest/api/%v/version/%v/relatedIssueCounts", i.version, versionID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	issues := new(model.VersionIssueCountsScheme)
	response, err := i.c.Call(request, issues)
	if err != nil {
		return nil, response, err
	}

	return issues, response, nil
}

func (i *internalProjectVersionImpl) UnresolvedIssueCount(ctx context.Context, versionID string) (*model.VersionUnresolvedIssuesCountScheme, *model.ResponseScheme, error) {

	if versionID == "" {
		return nil, nil, model.ErrNoVersionID
	}

	endpoint := fmt.Sprintf("rest/api/%v/version/%v/unresolvedIssueCount", i.version, versionID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	issues := new(model.VersionUnresolvedIssuesCountScheme)
	response, err := i.c.Call(request, issues)
	if err != nil {
		return nil, response, err
	}

	return issues, response, nil
}
