package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
)

func NewWorkflowService(client service.Connector, version string, scheme *WorkflowSchemeService, status *WorkflowStatusService) (*WorkflowService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &WorkflowService{
		internalClient: &internalWorkflowImpl{c: client, version: version},
		Scheme:         scheme,
		Status:         status,
	}, nil
}

type WorkflowService struct {
	internalClient jira.WorkflowConnector
	Scheme         *WorkflowSchemeService
	Status         *WorkflowStatusService
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
// DELETE /rest/api/{2-3}/workflow/{workflowID}
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow#search-workflows
func (w *WorkflowService) Delete(ctx context.Context, workflowID string) (*model.ResponseScheme, error) {
	return w.internalClient.Delete(ctx, workflowID)
}

type internalWorkflowImpl struct {
	c       service.Connector
	version string
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

func (i *internalWorkflowImpl) Delete(ctx context.Context, workflowID string) (*model.ResponseScheme, error) {

	if workflowID == "" {
		return nil, model.ErrNoWorkflowIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/workflow/%v", i.version, workflowID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
