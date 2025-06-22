package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"strconv"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewTypeSchemeService creates a new instance of TypeSchemeService.
func NewTypeSchemeService(client service.Connector, version string) (*TypeSchemeService, error) {

	if version == "" {
		return nil, fmt.Errorf("jira: %w", model.ErrNoVersionProvided)
	}

	return &TypeSchemeService{
		internalClient: &internalTypeSchemeImpl{c: client, version: version},
	}, nil
}

// TypeSchemeService provides methods to manage issue type schemes in Jira Service Management.
type TypeSchemeService struct {
	// internalClient is the connector interface for issue type scheme operations.
	internalClient jira.TypeSchemeConnector
}

// Gets returns a paginated list of issue type schemes.
//
// GET /rest/api/{2-3}/issuetypescheme
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#get-all-issue-type-schemes
func (t *TypeSchemeService) Gets(ctx context.Context, issueTypeSchemeIDs []int, startAt, maxResults int) (*model.IssueTypeSchemePageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*TypeSchemeService).Gets")
	defer span.End()

	return t.internalClient.Gets(ctx, issueTypeSchemeIDs, startAt, maxResults)
}

// Create creates an issue type scheme.
//
// POST /rest/api/{2-3}/issuetypescheme
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#create-issue-type-scheme
func (t *TypeSchemeService) Create(ctx context.Context, payload *model.IssueTypeSchemePayloadScheme) (*model.NewIssueTypeSchemeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*TypeSchemeService).Create")
	defer span.End()

	return t.internalClient.Create(ctx, payload)
}

// Items returns a paginated list of issue type scheme items.
//
// GET /rest/api/{2-3}/issuetypescheme/mapping
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#get-issue-type-scheme-items
func (t *TypeSchemeService) Items(ctx context.Context, issueTypeSchemeIDs []int, startAt, maxResults int) (*model.IssueTypeSchemeItemPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*TypeSchemeService).Items")
	defer span.End()

	return t.internalClient.Items(ctx, issueTypeSchemeIDs, startAt, maxResults)
}

// Projects returns a paginated list of issue type schemes and, for each issue type scheme, a list of the projects that use it.
//
// GET /rest/api/{2-3}/issuetypescheme/project
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#get-issue-type-schemes-for-projects
func (t *TypeSchemeService) Projects(ctx context.Context, projectIDs []int, startAt, maxResults int) (*model.ProjectIssueTypeSchemePageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*TypeSchemeService).Projects")
	defer span.End()

	return t.internalClient.Projects(ctx, projectIDs, startAt, maxResults)
}

// Assign assigns an issue type scheme to a project.
//
// PUT /rest/api/{2-3}/issuetypescheme/project
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#assign-issue-type-scheme-to-project
func (t *TypeSchemeService) Assign(ctx context.Context, issueTypeSchemeID, projectID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*TypeSchemeService).Assign")
	defer span.End()

	return t.internalClient.Assign(ctx, issueTypeSchemeID, projectID)
}

// Update updates an issue type scheme.
//
// PUT /rest/api/{2-3}/issuetypescheme/{issueTypeSchemeID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#update-issue-type-scheme
func (t *TypeSchemeService) Update(ctx context.Context, issueTypeSchemeID int, payload *model.IssueTypeSchemePayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*TypeSchemeService).Update")
	defer span.End()

	return t.internalClient.Update(ctx, issueTypeSchemeID, payload)
}

// Delete deletes an issue type scheme.
//
// 1.Only issue type schemes used in classic projects can be deleted.
//
// 2.Any projects assigned to the scheme are reassigned to the default issue type scheme.
//
// DELETE /rest/api/{2-3}/issuetypescheme/{issueTypeSchemeID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#delete-issue-type-scheme
func (t *TypeSchemeService) Delete(ctx context.Context, issueTypeSchemeID int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*TypeSchemeService).Delete")
	defer span.End()

	return t.internalClient.Delete(ctx, issueTypeSchemeID)
}

// Append adds issue types to an issue type scheme.
//
// 1.The added issue types are appended to the issue types list.
//
// 2.If any of the issue types exist in the issue type scheme, the operation fails and no issue types are added.
//
// PUT /rest/api/{2-3}/issuetypescheme/{issueTypeSchemeID}/issuetype
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#add-issue-types-to-issue-type-scheme
func (t *TypeSchemeService) Append(ctx context.Context, issueTypeSchemeID int, issueTypeIDs []int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*TypeSchemeService).Append")
	defer span.End()

	return t.internalClient.Append(ctx, issueTypeSchemeID, issueTypeIDs)
}

// Remove removes an issue type from an issue type scheme, this operation cannot remove:
//
// 1.any issue type used by issues.
//
// 2.any issue types from the default issue type scheme.
//
// 3.the last standard issue type from an issue type scheme.
//
// DELETE /rest/api/{2-3}/issuetypescheme/{issueTypeSchemeID}/issuetype/{issueTypeID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#remove-issue-type-from-issue-type-scheme
func (t *TypeSchemeService) Remove(ctx context.Context, issueTypeSchemeID, issueTypeID int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*TypeSchemeService).Remove")
	defer span.End()

	return t.internalClient.Remove(ctx, issueTypeSchemeID, issueTypeID)
}

// Reorder reorders the issue type scheme by moving one or more issue types after another issue type or to the first/last position.
//
// PUT /rest/api/{2-3}/issuetypescheme/{issueTypeSchemeId}/issuetype/move
func (t *TypeSchemeService) Reorder(ctx context.Context, issueTypeSchemeId string, payload *model.IssueTypeSchemeOrderPayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*TypeSchemeService).Reorder")
	defer span.End()

	return t.internalClient.Reorder(ctx, issueTypeSchemeId, payload)
}

type internalTypeSchemeImpl struct {
	c       service.Connector
	version string
}

func (i *internalTypeSchemeImpl) Gets(ctx context.Context, issueTypeSchemeIDs []int, startAt, maxResults int) (*model.IssueTypeSchemePageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalTypeSchemeImpl).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "get_issue_type_schemes"),
		attribute.IntSlice("issue_type_scheme_ids", issueTypeSchemeIDs),
		attribute.Int("start_at", startAt),
		attribute.Int("max_results", maxResults),
	)

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range issueTypeSchemeIDs {
		params.Add("id", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	page := new(model.IssueTypeSchemePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalTypeSchemeImpl) Create(ctx context.Context, payload *model.IssueTypeSchemePayloadScheme) (*model.NewIssueTypeSchemeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalTypeSchemeImpl).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "create_issue_type_scheme"),
	)

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	issueType := new(model.NewIssueTypeSchemeScheme)
	response, err := i.c.Call(request, issueType)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return issueType, response, nil
}

func (i *internalTypeSchemeImpl) Items(ctx context.Context, issueTypeSchemeIDs []int, startAt, maxResults int) (*model.IssueTypeSchemeItemPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalTypeSchemeImpl).Items", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "get_issue_type_scheme_items"),
		attribute.IntSlice("issue_type_scheme_ids", issueTypeSchemeIDs),
		attribute.Int("start_at", startAt),
		attribute.Int("max_results", maxResults),
	)

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range issueTypeSchemeIDs {
		params.Add("issueTypeSchemeId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/mapping?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	page := new(model.IssueTypeSchemeItemPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalTypeSchemeImpl) Projects(ctx context.Context, projectIDs []int, startAt, maxResults int) (*model.ProjectIssueTypeSchemePageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalTypeSchemeImpl).Projects", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "get_issue_type_schemes_for_projects"),
		attribute.IntSlice("project_ids", projectIDs),
		attribute.Int("start_at", startAt),
		attribute.Int("max_results", maxResults),
	)

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range projectIDs {
		params.Add("projectId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/project?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	page := new(model.ProjectIssueTypeSchemePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalTypeSchemeImpl) Assign(ctx context.Context, issueTypeSchemeID, projectID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalTypeSchemeImpl).Assign", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "assign_issue_type_scheme_to_project"),
		attribute.String("issue_type_scheme_id", issueTypeSchemeID),
		attribute.String("project_id", projectID),
	)

	if issueTypeSchemeID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueTypeSchemeID)
		recordError(span, err)
		return nil, err
	}

	if projectID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoProjectID)
		recordError(span, err)
		return nil, err
	}

	payload := map[string]interface{}{
		"issueTypeSchemeId": issueTypeSchemeID,
		"projectId":         projectID,
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/project", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
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

func (i *internalTypeSchemeImpl) Update(ctx context.Context, issueTypeSchemeID int, payload *model.IssueTypeSchemePayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalTypeSchemeImpl).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "update_issue_type_scheme"),
		attribute.Int("issue_type_scheme_id", issueTypeSchemeID),
	)

	if issueTypeSchemeID == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueTypeSchemeID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/%v", i.version, issueTypeSchemeID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
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

func (i *internalTypeSchemeImpl) Delete(ctx context.Context, issueTypeSchemeID int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalTypeSchemeImpl).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "delete_issue_type_scheme"),
		attribute.Int("issue_type_scheme_id", issueTypeSchemeID),
	)

	if issueTypeSchemeID == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueTypeSchemeID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/%v", i.version, issueTypeSchemeID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
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

func (i *internalTypeSchemeImpl) Append(ctx context.Context, issueTypeSchemeID int, issueTypeIDs []int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalTypeSchemeImpl).Append", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "append_issue_types_to_scheme"),
		attribute.Int("issue_type_scheme_id", issueTypeSchemeID),
		attribute.IntSlice("issue_type_ids", issueTypeIDs),
	)

	if len(issueTypeIDs) == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueTypes)
		recordError(span, err)
		return nil, err
	}

	var ids []string
	for _, issueTypeID := range issueTypeIDs {
		ids = append(ids, strconv.Itoa(issueTypeID))
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/%v/issuetype", i.version, issueTypeSchemeID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]interface{}{"issueTypeIds": ids})
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

func (i *internalTypeSchemeImpl) Remove(ctx context.Context, issueTypeSchemeID, issueTypeID int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalTypeSchemeImpl).Remove", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "remove_issue_type_from_scheme"),
		attribute.Int("issue_type_scheme_id", issueTypeSchemeID),
		attribute.Int("issue_type_id", issueTypeID),
	)

	if issueTypeSchemeID == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueTypeSchemeID)
		recordError(span, err)
		return nil, err
	}

	if issueTypeID == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueTypeID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/issuetypescheme/%v/issuetype/%v", i.version, issueTypeSchemeID, issueTypeID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
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

func (i *internalTypeSchemeImpl) Reorder(ctx context.Context, issueTypeSchemeId string, payload *model.IssueTypeSchemeOrderPayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalTypeSchemeImpl).Reorder", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "reorder_issue_types_in_scheme"),
		attribute.String("issue_type_scheme_id", issueTypeSchemeId),
	)

	if issueTypeSchemeId == "" {
		err := fmt.Errorf("jira: TypeSchemeService.Reorder: %w", model.ErrNoIssueTypeSchemeID)
		recordError(span, err)
		return nil, err
	}

	if len(payload.IssueTypeIDs) == 0 {
		err := fmt.Errorf("jira: TypeSchemeService.Reorder: %w", model.ErrNoIssueTypes)
		recordError(span, err)
		return nil, err
	}

	if payload.After == "" && payload.Position == "" {
		err := fmt.Errorf("jira: TypeSchemeService.Reorder: %w", model.ErrNoIssueTypeReorderAttr)
		recordError(span, err)
		return nil, err
	}

	if payload.Position != "" && payload.Position != model.SchemePositionFirst && payload.Position != model.SchemePositionLast {
		err := fmt.Errorf("jira: TypeSchemeService.Reorder: %w", model.ErrInvalidIssueTypeSchemePosition)
		recordError(span, err)
		return nil, err
	}

	if slices.Contains(payload.IssueTypeIDs, payload.After) {
		err := fmt.Errorf("jira: TypeSchemeService.Reorder: %w", model.ErrInvalidIssueTypeSchemeAfter)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("/rest/api/%v/issuetypescheme/%v/issuetype/move", i.version, issueTypeSchemeId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		err = fmt.Errorf("jira: TypeSchemeService.Reorder: %w", err)
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
