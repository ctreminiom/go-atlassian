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

// NewWorkflowSchemeService creates a new instance of WorkflowSchemeService.
func NewWorkflowSchemeService(client service.Connector, version string, issueType *WorkflowSchemeIssueTypeService) *WorkflowSchemeService {

	return &WorkflowSchemeService{
		internalClient: &internalWorkflowSchemeImpl{c: client, version: version},
		IssueType:      issueType,
	}
}

// WorkflowSchemeService provides methods to manage workflow schemes in Jira Service Management.
type WorkflowSchemeService struct {
	// internalClient is the connector interface for workflow scheme operations.
	internalClient jira.WorkflowSchemeConnector
	// IssueType is the service for managing workflow scheme issue types.
	IssueType *WorkflowSchemeIssueTypeService
}

// Gets returns a paginated list of all workflow schemes, not including draft workflow schemes.
//
// GET /rest/api/{2-3}/workflowscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#gets-workflows-schemes
func (w *WorkflowSchemeService) Gets(ctx context.Context, startAt, maxResults int) (*model.WorkflowSchemePageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkflowSchemeService).Gets")
	defer span.End()

	return w.internalClient.Gets(ctx, startAt, maxResults)
}

// Create creates a workflow scheme.
//
// POST /rest/api/{2-3}/workflowscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#create-workflows-scheme
func (w *WorkflowSchemeService) Create(ctx context.Context, payload *model.WorkflowSchemePayloadScheme) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkflowSchemeService).Create")
	defer span.End()

	return w.internalClient.Create(ctx, payload)
}

// Get returns a workflow scheme
//
// GET /rest/api/{2-3}/workflowscheme/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#get-workflow-scheme
func (w *WorkflowSchemeService) Get(ctx context.Context, schemeID int, returnDraftIfExists bool) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkflowSchemeService).Get")
	defer span.End()

	return w.internalClient.Get(ctx, schemeID, returnDraftIfExists)
}

// Update updates a workflow scheme, including the name, default workflow, issue type to project mappings, and more.
//
// If the workflow scheme is active (that is, being used by at least one project), then a draft workflow scheme is created or updated instead,
//
// provided that updateDraftIfNeeded is set to true.
//
// PUT /rest/api/{2-3}/workflowscheme/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#update-workflow-scheme
func (w *WorkflowSchemeService) Update(ctx context.Context, schemeID int, payload *model.WorkflowSchemePayloadScheme) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkflowSchemeService).Update")
	defer span.End()

	return w.internalClient.Update(ctx, schemeID, payload)
}

// Delete deletes a workflow scheme.
//
// Note that a workflow scheme cannot be deleted if it is active (that is, being used by at least one project).
//
// DELETE /rest/api/{2-3}/workflowscheme/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#delete-workflow-scheme
func (w *WorkflowSchemeService) Delete(ctx context.Context, schemeID int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkflowSchemeService).Delete")
	defer span.End()

	return w.internalClient.Delete(ctx, schemeID)
}

// Associations returns a list of the workflow schemes associated with a list of projects.
//
// Each returned workflow scheme includes a list of the requested projects associated with it.
//
// Any team-managed or non-existent projects in the request are ignored and no errors are returned.
//
// If the project is associated with the Default Workflow Scheme no ID is returned.
//
// This is because the way the Default Workflow Scheme is stored means it has no ID.
//
// GET /rest/api/{2-3}/workflowscheme/project
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#get-workflow-schemes-associations
func (w *WorkflowSchemeService) Associations(ctx context.Context, projectIDs []int) (*model.WorkflowSchemeAssociationPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkflowSchemeService).Associations")
	defer span.End()

	return w.internalClient.Associations(ctx, projectIDs)
}

// Assign assigns a workflow scheme to a project.
//
// This operation is performed only when there are no issues in the project.
//
// Workflow schemes can only be assigned to classic projects.
//
// PUT /rest/api/{2-3}/workflowscheme/project
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#get-workflow-schemes-associations
func (w *WorkflowSchemeService) Assign(ctx context.Context, schemeID, projectID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkflowSchemeService).Assign")
	defer span.End()

	return w.internalClient.Assign(ctx, schemeID, projectID)
}

type internalWorkflowSchemeImpl struct {
	c       service.Connector
	version string
}

func (i *internalWorkflowSchemeImpl) Gets(ctx context.Context, startAt, maxResults int) (*model.WorkflowSchemePageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkflowSchemeImpl).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "get_workflow_schemes"),
		attribute.Int("start_at", startAt),
		attribute.Int("max_results", maxResults),
	)

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	endpoint := fmt.Sprintf("rest/api/%v/workflowscheme?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	page := new(model.WorkflowSchemePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalWorkflowSchemeImpl) Create(ctx context.Context, payload *model.WorkflowSchemePayloadScheme) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkflowSchemeImpl).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "create_workflow_scheme"),
	)

	endpoint := fmt.Sprintf("rest/api/%v/workflowscheme", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	workflowScheme := new(model.WorkflowSchemeScheme)
	response, err := i.c.Call(request, workflowScheme)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return workflowScheme, response, nil
}

func (i *internalWorkflowSchemeImpl) Get(ctx context.Context, schemeID int, returnDraftIfExists bool) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkflowSchemeImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "get_workflow_scheme"),
		attribute.Int("scheme_id", schemeID),
		attribute.Bool("return_draft_if_exists", returnDraftIfExists),
	)

	if schemeID == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoWorkflowSchemeID)
		recordError(span, err)
		return nil, nil, err
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/workflowscheme/%v", i.version, schemeID))

	if returnDraftIfExists {

		params := url.Values{}
		params.Add("returnDraftIfExists", "true")

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	workflowScheme := new(model.WorkflowSchemeScheme)
	response, err := i.c.Call(request, workflowScheme)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return workflowScheme, response, nil
}

func (i *internalWorkflowSchemeImpl) Update(ctx context.Context, schemeID int, payload *model.WorkflowSchemePayloadScheme) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkflowSchemeImpl).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "update_workflow_scheme"),
		attribute.Int("scheme_id", schemeID),
	)

	if schemeID == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoWorkflowSchemeID)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/workflowscheme/%v", i.version, schemeID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	workflowScheme := new(model.WorkflowSchemeScheme)
	response, err := i.c.Call(request, workflowScheme)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return workflowScheme, response, nil
}

func (i *internalWorkflowSchemeImpl) Delete(ctx context.Context, schemeID int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkflowSchemeImpl).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "delete_workflow_scheme"),
		attribute.Int("scheme_id", schemeID),
	)

	if schemeID == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoWorkflowSchemeID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/workflowscheme/%v", i.version, schemeID)

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

func (i *internalWorkflowSchemeImpl) Associations(ctx context.Context, projectIDs []int) (*model.WorkflowSchemeAssociationPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkflowSchemeImpl).Associations", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "get_workflow_scheme_associations"),
		attribute.Int("project_ids_count", len(projectIDs)),
	)

	if len(projectIDs) == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoProjects)
		recordError(span, err)
		return nil, nil, err
	}

	params := url.Values{}
	for _, projectID := range projectIDs {
		params.Add("projectId", strconv.Itoa(projectID))
	}

	endpoint := fmt.Sprintf("rest/api/%v/workflowscheme/project?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	mapping := new(model.WorkflowSchemeAssociationPageScheme)
	response, err := i.c.Call(request, mapping)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return mapping, response, nil
}

func (i *internalWorkflowSchemeImpl) Assign(ctx context.Context, schemeID, projectID string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkflowSchemeImpl).Assign", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "assign_workflow_scheme_to_project"),
		attribute.String("scheme_id", schemeID),
		attribute.String("project_id", projectID),
	)

	if schemeID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoWorkflowSchemeID)
		recordError(span, err)
		return nil, err
	}

	if projectID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoProjectIDOrKey)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/workflowscheme/project", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]interface{}{"workflowSchemeId": schemeID, "projectId": projectID})
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
