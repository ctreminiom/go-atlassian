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

// NewWorkflowStatusService creates a new instance of WorkflowStatusService.
func NewWorkflowStatusService(client service.Connector, version string) (*WorkflowStatusService, error) {

	if version == "" {
		return nil, fmt.Errorf("jira: %w", model.ErrNoVersionProvided)
	}

	return &WorkflowStatusService{
		internalClient: &internalWorkflowStatusImpl{c: client, version: version},
	}, nil
}

// WorkflowStatusService provides methods to manage workflow statuses in Jira Service Management.
type WorkflowStatusService struct {
	// internalClient is the connector interface for workflow status operations.
	internalClient jira.WorkflowStatusConnector
}

// Gets returns a list of the statuses specified by one or more status IDs.
//
// GET /rest/api/{2-3}/statuses
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#gets-workflow-statuses
func (w *WorkflowStatusService) Gets(ctx context.Context, ids, expand []string) ([]*model.WorkflowStatusDetailScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkflowStatusService).Gets")
	defer span.End()

	return w.internalClient.Gets(ctx, ids, expand)
}

// Update updates statuses by ID.
//
// PUT /rest/api/{2-3}/statuses
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#update-workflow-statuses
func (w *WorkflowStatusService) Update(ctx context.Context, payload *model.WorkflowStatusPayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkflowStatusService).Update")
	defer span.End()

	return w.internalClient.Update(ctx, payload)
}

// Create creates statuses for a global or project scope.
//
// POST /rest/api/{2-3}/statuses
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#create-workflow-statuses
func (w *WorkflowStatusService) Create(ctx context.Context, payload *model.WorkflowStatusPayloadScheme) ([]*model.WorkflowStatusDetailScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkflowStatusService).Create")
	defer span.End()

	return w.internalClient.Create(ctx, payload)
}

// Delete deletes statuses by ID.
//
// DELETE /rest/api/{2-3}/statuses
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#delete-workflow-statuses
func (w *WorkflowStatusService) Delete(ctx context.Context, ids []string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkflowStatusService).Delete")
	defer span.End()

	return w.internalClient.Delete(ctx, ids)
}

// Search returns a paginated list of statuses that match a search on name or project.
//
// GET /rest/api/{2-3}/statuses/search
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#search-workflow-statuses
func (w *WorkflowStatusService) Search(ctx context.Context, options *model.WorkflowStatusSearchParams, startAt, maxResults int) (*model.WorkflowStatusDetailPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkflowStatusService).Search")
	defer span.End()

	return w.internalClient.Search(ctx, options, startAt, maxResults)
}

// Bulk returns a list of all statuses associated with active workflows.
//
// GET /rest/api/{2-3}/status
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#bulk-workflow-statuses
func (w *WorkflowStatusService) Bulk(ctx context.Context) ([]*model.StatusDetailScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkflowStatusService).Bulk")
	defer span.End()

	return w.internalClient.Bulk(ctx)
}

// Get returns a status.
//
// The status must be associated with an active workflow to be returned.
//
// If a name is used on more than one status, only the status found first is returned.
//
// Therefore, identifying the status by its ID may be preferable.
//
// This operation can be accessed anonymously.
//
// GET /rest/api/{2-3}/status/{idOrName}
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#get-workflow-status
func (w *WorkflowStatusService) Get(ctx context.Context, idOrName string) (*model.StatusDetailScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*WorkflowStatusService).Get")
	defer span.End()

	return w.internalClient.Get(ctx, idOrName)
}

type internalWorkflowStatusImpl struct {
	c       service.Connector
	version string
}

func (i *internalWorkflowStatusImpl) Get(ctx context.Context, idOrName string) (*model.StatusDetailScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkflowStatusImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "get_workflow_status"),
		attribute.String("id_or_name", idOrName),
	)

	if idOrName == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoWorkflowStatusNameOrID)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("/rest/api/%v/status/%v", i.version, idOrName)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	workflowStatus := new(model.StatusDetailScheme)
	response, err := i.c.Call(request, workflowStatus)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return workflowStatus, response, nil
}

func (i *internalWorkflowStatusImpl) Bulk(ctx context.Context) ([]*model.StatusDetailScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkflowStatusImpl).Bulk", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "get_workflow_statuses_bulk"),
	)

	endpoint := fmt.Sprintf("/rest/api/%v/status", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	var statuses []*model.StatusDetailScheme
	response, err := i.c.Call(request, &statuses)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return statuses, response, nil
}

func (i *internalWorkflowStatusImpl) Gets(ctx context.Context, ids, expand []string) ([]*model.WorkflowStatusDetailScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkflowStatusImpl).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "get_workflow_statuses"),
		attribute.Int("ids_count", len(ids)),
		attribute.Int("expand_count", len(expand)),
	)

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/statuses", i.version))

	params := url.Values{}
	for _, id := range ids {
		params.Add("id", id)
	}

	if expand != nil {
		params.Add("expand", strings.Join(expand, ","))
	}

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	var statuses []*model.WorkflowStatusDetailScheme
	response, err := i.c.Call(request, &statuses)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return statuses, response, nil
}

func (i *internalWorkflowStatusImpl) Update(ctx context.Context, payload *model.WorkflowStatusPayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkflowStatusImpl).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "update_workflow_statuses"),
	)

	endpoint := fmt.Sprintf("rest/api/%v/statuses", i.version)

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

func (i *internalWorkflowStatusImpl) Create(ctx context.Context, payload *model.WorkflowStatusPayloadScheme) ([]*model.WorkflowStatusDetailScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkflowStatusImpl).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "create_workflow_statuses"),
	)

	if payload == nil || len(payload.Statuses) == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoWorkflowStatuses)
		recordError(span, err)
		return nil, nil, err
	}

	if payload.Scope == nil {
		err := fmt.Errorf("jira: %w", model.ErrNoWorkflowScope)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/statuses", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	var workflowStatuses []*model.WorkflowStatusDetailScheme
	response, err := i.c.Call(request, &workflowStatuses)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return workflowStatuses, response, nil
}

func (i *internalWorkflowStatusImpl) Delete(ctx context.Context, ids []string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkflowStatusImpl).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "delete_workflow_statuses"),
		attribute.Int("ids_count", len(ids)),
	)

	if len(ids) == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoWorkflowStatuses)
		recordError(span, err)
		return nil, err
	}

	params := url.Values{}
	for _, id := range ids {
		params.Add("id", id)
	}

	endpoint := fmt.Sprintf("rest/api/%v/statuses?%v", i.version, params.Encode())

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

func (i *internalWorkflowStatusImpl) Search(ctx context.Context, options *model.WorkflowStatusSearchParams, startAt, maxResults int) (*model.WorkflowStatusDetailPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalWorkflowStatusImpl).Search", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation", "search_workflow_statuses"),
		attribute.Int("start_at", startAt),
		attribute.Int("max_results", maxResults),
	)

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		if options.Expand != nil {
			params.Add("expand", strings.Join(options.Expand, ","))
		}

		if options.ProjectID != "" {
			params.Add("projectId", options.ProjectID)
		}

		if options.SearchString != "" {
			params.Add("searchString", options.SearchString)
		}

		if options.StatusCategory != "" {
			params.Add("statusCategory", options.StatusCategory)
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/statuses/search?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	page := new(model.WorkflowStatusDetailPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}
