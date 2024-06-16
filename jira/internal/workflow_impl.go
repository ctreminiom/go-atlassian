package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func NewWorkflowService(client service.Connector,
	version string,
	scheme *WorkflowSchemeService,
	status *WorkflowStatusService,
	validator *WorkflowValidatorService,
) (*WorkflowService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &WorkflowService{
		internalClient: &internalWorkflowImpl{c: client, version: version},
		Scheme:         scheme,
		Status:         status,
		Validator:      validator,
	}, nil
}

type WorkflowService struct {
	internalClient jira.WorkflowConnector
	Scheme         *WorkflowSchemeService
	Status         *WorkflowStatusService
	Validator      *WorkflowValidatorService
}

// Create creates a workflow.
//
// You can define transition rules using the shapes detailed in the following sections.
//
// If no transitional rules are specified the default system transition rules are used.
//
// POST /rest/api/{2-3}/workflow
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow#create-workflow
func (w *WorkflowService) Create(ctx context.Context, payload *model.WorkflowPayloadScheme) (*model.WorkflowCreatedResponseScheme, *model.ResponseScheme, error) {
	return w.internalClient.Create(ctx, payload)
}

// Gets returns a paginated list of published classic workflows.
//
// When workflow names are specified, details of those workflows are returned.
//
// Otherwise, all published classic workflows are returned.
//
// GET /rest/api/{2-3}/workflow/search
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow#search-workflows
func (w *WorkflowService) Gets(ctx context.Context, options *model.WorkflowSearchOptions, startAt, maxResults int) (*model.WorkflowPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Gets(ctx, options, startAt, maxResults)
}

// Delete deletes a workflow.
//
// The workflow cannot be deleted if it is:
//
// 1. an active workflow.
// 2. a system workflow.
// 3. associated with any workflow scheme.
// 4. associated with any draft workflow scheme.
//
// DELETE /rest/api/{2-3}/workflow/{entityId}
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow#search-workflows
func (w *WorkflowService) Delete(ctx context.Context, workflowId string) (*model.ResponseScheme, error) {
	return w.internalClient.Delete(ctx, workflowId)
}

// Bulk returns a list of workflows and related statuses by providing workflow names,
//
// workflow IDs, or project and issue types.
//
// POST /rest/api/{2-3}/workflows
func (w *WorkflowService) Bulk(ctx context.Context, options *model.WorkflowBulkOptionsScheme, expand []string) (*model.WorkflowReadResponseScheme, *model.ResponseScheme, error) {
	return w.internalClient.Bulk(ctx, options, expand)
}

// Capabilities get the list of workflow capabilities for a specific workflow using either the workflow ID, or the project and issue type ID pair.
//
// The response includes the scope of the workflow, defined as global/project-based,
//
// and a list of project types that the workflow is scoped to.
//
// It also includes all rules organised into their broad categories (conditions, validators, actions, triggers, screens)
//
// as well as the source location (Atlassian-provided, Connect, Forge).
//
// GET /rest/api/{2-3}/workflows/capabilities
func (w *WorkflowService) Capabilities(ctx context.Context, workflowID, projectID, issueTypeID string) (*model.WorkflowCapabilitiesScheme, *model.ResponseScheme, error) {
	return w.internalClient.Capabilities(ctx, workflowID, projectID, issueTypeID)
}

// Creates creates workflows and related statuses.
//
// POST /rest/api/{2-3}/workflows/create
func (w *WorkflowService) Creates(ctx context.Context, payload *model.WorkflowCreatesPayloadScheme) (*model.WorkflowCreateResponseScheme, *model.ResponseScheme, error) {
	return w.internalClient.Creates(ctx, payload)
}

// Updates updates workflows and related statuses.
//
// POST /rest/api/{2-3}/workflows/update
func (w *WorkflowService) Updates(ctx context.Context, payload *model.WorkflowUpdatesPayloadScheme, expand []string) (*model.WorkflowUpdateResponseScheme, *model.ResponseScheme, error) {
	return w.internalClient.Updates(ctx, payload, expand)
}

type internalWorkflowImpl struct {
	c       service.Connector
	version string
}

func (i *internalWorkflowImpl) Bulk(ctx context.Context, options *model.WorkflowBulkOptionsScheme, expand []string) (*model.WorkflowReadResponseScheme, *model.ResponseScheme, error) {

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/workflows", i.version))

	if len(expand) != 0 {

		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", options)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.WorkflowReadResponseScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalWorkflowImpl) Capabilities(ctx context.Context, workflowID, projectID, issueTypeID string) (*model.WorkflowCapabilitiesScheme, *model.ResponseScheme, error) {

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/workflows/capabilities", i.version))

	params := url.Values{}

	if workflowID != "" {
		params.Add("workflowId", workflowID)
	}

	if projectID != "" {
		params.Add("projectId", projectID)
	}

	if issueTypeID != "" {
		params.Add("issueTypeId", issueTypeID)
	}

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	capabilities := new(model.WorkflowCapabilitiesScheme)
	response, err := i.c.Call(request, capabilities)
	if err != nil {
		return nil, response, err
	}

	return capabilities, response, nil
}

func (i *internalWorkflowImpl) Creates(ctx context.Context, payload *model.WorkflowCreatesPayloadScheme) (*model.WorkflowCreateResponseScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/workflows/create", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	workflows := new(model.WorkflowCreateResponseScheme)
	response, err := i.c.Call(request, workflows)
	if err != nil {
		return nil, response, err
	}

	return workflows, response, nil
}

func (i *internalWorkflowImpl) Updates(ctx context.Context, payload *model.WorkflowUpdatesPayloadScheme, expand []string) (*model.WorkflowUpdateResponseScheme, *model.ResponseScheme, error) {

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/workflows/update", i.version))

	if len(expand) != 0 {

		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	workflows := new(model.WorkflowUpdateResponseScheme)
	response, err := i.c.Call(request, workflows)
	if err != nil {
		return nil, response, err
	}

	return workflows, response, nil
}

func (i *internalWorkflowImpl) Create(ctx context.Context, payload *model.WorkflowPayloadScheme) (*model.WorkflowCreatedResponseScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/workflow", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	workflow := new(model.WorkflowCreatedResponseScheme)
	response, err := i.c.Call(request, workflow)
	if err != nil {
		return nil, response, err
	}

	return workflow, response, nil
}

func (i *internalWorkflowImpl) Gets(ctx context.Context, options *model.WorkflowSearchOptions, startAt, maxResults int) (*model.WorkflowPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {
		params.Add("isActive", fmt.Sprintf("%v", options.IsActive))

		for _, name := range options.WorkflowName {
			params.Add("workflowName", name)
		}

		if options.QueryString != "" {
			params.Add("queryString", options.QueryString)
		}

		if options.OrderBy != "" {
			params.Add("orderBy", options.OrderBy)
		}

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/workflow/search?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.WorkflowPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalWorkflowImpl) Delete(ctx context.Context, workflowId string) (*model.ResponseScheme, error) {

	if workflowId == "" {
		return nil, model.ErrNoWorkflowIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/workflow/%v", i.version, workflowId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
