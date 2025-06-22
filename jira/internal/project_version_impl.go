package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewProjectVersionService creates a new instance of ProjectVersionService.
func NewProjectVersionService(client service.Connector, version string) (*ProjectVersionService, error) {

	if version == "" {
		return nil, fmt.Errorf("jira: %w", model.ErrNoVersionProvided)
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
	ctx, span := tracer().Start(ctx, "(*ProjectVersionService).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_project_versions"),
		attribute.String("jira.project.key", projectKeyOrID),
	)

	result, response, err := p.internalClient.Gets(ctx, projectKeyOrID)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Search returns a paginated list of all versions in a project.
//
// GET /rest/api/{2-3}/project/{projectKeyOrID}/version
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-project-versions-paginated
func (p *ProjectVersionService) Search(ctx context.Context, projectKeyOrID string, options *model.VersionGetsOptions, startAt, maxResults int) (*model.VersionPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ProjectVersionService).Search", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "search_project_versions"),
		attribute.String("jira.project.key", projectKeyOrID),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
	)

	result, response, err := p.internalClient.Search(ctx, projectKeyOrID, options, startAt, maxResults)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Create creates a project version.
//
// POST /rest/api/{2-3}/version
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#create-version
func (p *ProjectVersionService) Create(ctx context.Context, payload *model.VersionPayloadScheme) (*model.VersionScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ProjectVersionService).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "create_version"),
	)

	result, response, err := p.internalClient.Create(ctx, payload)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Get returns a project version.
//
// GET /rest/api/{2-3}/version/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-version
func (p *ProjectVersionService) Get(ctx context.Context, versionID string, expand []string) (*model.VersionScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ProjectVersionService).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_version"),
		attribute.String("jira.version.id", versionID),
		attribute.StringSlice("jira.expand", expand),
	)

	result, response, err := p.internalClient.Get(ctx, versionID, expand)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Update updates a project version.
//
// PUT /rest/api/{2-3}/version/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#update-version
func (p *ProjectVersionService) Update(ctx context.Context, versionID string, payload *model.VersionPayloadScheme) (*model.VersionScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ProjectVersionService).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update_version"),
		attribute.String("jira.version.id", versionID),
	)

	result, response, err := p.internalClient.Update(ctx, versionID, payload)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Merge merges two project versions.
//
// # The merge is completed by deleting the version specified in id and replacing any occurrences of
//
// its ID in fixVersion with the version ID specified in moveIssuesTo.
//
// PUT /rest/api/{2-3}/version/{id}/mergeto/{moveIssuesTo}
func (p *ProjectVersionService) Merge(ctx context.Context, versionID, versionMoveIssuesTo string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ProjectVersionService).Merge", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "merge_versions"),
		attribute.String("jira.version.id", versionID),
		attribute.String("jira.version.move_to", versionMoveIssuesTo),
	)

	response, err := p.internalClient.Merge(ctx, versionID, versionMoveIssuesTo)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
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
	ctx, span := tracer().Start(ctx, "(*ProjectVersionService).RelatedIssueCounts", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_version_related_issue_counts"),
		attribute.String("jira.version.id", versionID),
	)

	result, response, err := p.internalClient.RelatedIssueCounts(ctx, versionID)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// UnresolvedIssueCount returns counts of the issues and unresolved issues for the project version.
//
// GET /rest/api/{2-3}/version/{id}/unresolvedIssueCount
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-versions-unresolved-issues-count
func (p *ProjectVersionService) UnresolvedIssueCount(ctx context.Context, versionID string) (*model.VersionUnresolvedIssuesCountScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ProjectVersionService).UnresolvedIssueCount", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_version_unresolved_issue_count"),
		attribute.String("jira.version.id", versionID),
	)

	result, response, err := p.internalClient.UnresolvedIssueCount(ctx, versionID)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

type internalProjectVersionImpl struct {
	c       service.Connector
	version string
}

func (i *internalProjectVersionImpl) Gets(ctx context.Context, projectKeyOrID string) ([]*model.VersionScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalProjectVersionImpl).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_project_versions"),
		attribute.String("jira.project.key", projectKeyOrID),
	)

	if projectKeyOrID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoProjectIDOrKey)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/versions", i.version, projectKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	var versions []*model.VersionScheme
	response, err := i.c.Call(request, &versions)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return versions, response, nil
}

func (i *internalProjectVersionImpl) Search(ctx context.Context, projectKeyOrID string, options *model.VersionGetsOptions, startAt, maxResults int) (*model.VersionPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalProjectVersionImpl).Search", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "search_project_versions"),
		attribute.String("jira.project.key", projectKeyOrID),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
	)

	if projectKeyOrID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoProjectIDOrKey)
		recordError(span, err)
		return nil, nil, err
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
		recordError(span, err)
		return nil, nil, err
	}

	page := new(model.VersionPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalProjectVersionImpl) Create(ctx context.Context, payload *model.VersionPayloadScheme) (*model.VersionScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalProjectVersionImpl).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "create_version"),
	)

	endpoint := fmt.Sprintf("rest/api/%v/version", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	version := new(model.VersionScheme)
	response, err := i.c.Call(request, version)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return version, response, nil
}

func (i *internalProjectVersionImpl) Get(ctx context.Context, versionID string, expand []string) (*model.VersionScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalProjectVersionImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_version"),
		attribute.String("jira.version.id", versionID),
		attribute.StringSlice("jira.expand", expand),
	)

	if versionID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoVersionID)
		recordError(span, err)
		return nil, nil, err
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
		recordError(span, err)
		return nil, nil, err
	}

	version := new(model.VersionScheme)
	response, err := i.c.Call(request, version)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return version, response, nil
}

func (i *internalProjectVersionImpl) Update(ctx context.Context, versionID string, payload *model.VersionPayloadScheme) (*model.VersionScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalProjectVersionImpl).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update_version"),
		attribute.String("jira.version.id", versionID),
	)

	if versionID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoVersionID)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/version/%v", i.version, versionID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	version := new(model.VersionScheme)
	response, err := i.c.Call(request, version)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return version, response, nil
}

func (i *internalProjectVersionImpl) Merge(ctx context.Context, versionID, versionMoveIssuesTo string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalProjectVersionImpl).Merge", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "merge_versions"),
		attribute.String("jira.version.id", versionID),
		attribute.String("jira.version.move_to", versionMoveIssuesTo),
	)

	if versionID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoVersionID)
		recordError(span, err)
		return nil, err
	}

	if versionMoveIssuesTo == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoVersionID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/version/%v/mergeto/%v", i.version, versionID, versionMoveIssuesTo)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

func (i *internalProjectVersionImpl) RelatedIssueCounts(ctx context.Context, versionID string) (*model.VersionIssueCountsScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalProjectVersionImpl).RelatedIssueCounts", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_version_related_issue_counts"),
		attribute.String("jira.version.id", versionID),
	)

	if versionID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoVersionID)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/version/%v/relatedIssueCounts", i.version, versionID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	issues := new(model.VersionIssueCountsScheme)
	response, err := i.c.Call(request, issues)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return issues, response, nil
}

func (i *internalProjectVersionImpl) UnresolvedIssueCount(ctx context.Context, versionID string) (*model.VersionUnresolvedIssuesCountScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalProjectVersionImpl).UnresolvedIssueCount", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_version_unresolved_issue_count"),
		attribute.String("jira.version.id", versionID),
	)

	if versionID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoVersionID)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/version/%v/unresolvedIssueCount", i.version, versionID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	issues := new(model.VersionUnresolvedIssuesCountScheme)
	response, err := i.c.Call(request, issues)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return issues, response, nil
}
