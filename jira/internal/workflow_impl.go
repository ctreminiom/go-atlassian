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

// NewWorkflowService creates a new instance of WorkflowService.
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

// WorkflowService provides methods to manage workflows in Jira Service Management.
type WorkflowService struct {
	// internalClient is the connector interface for workflow operations.
	internalClient jira.WorkflowConnector
	// Scheme is the service for managing workflow schemes.
	Scheme *WorkflowSchemeService
	// Status is the service for managing workflow statuses.
	Status *WorkflowStatusService
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

// Search searches for workflows based on specified criteria.
//
// This method returns a paginated list of workflows that match the search criteria provided in the `options` parameter.
//
// The search can be expanded to include additional details, such as transition links, by specifying the `expand` and `transitionLinks` parameters.
//
// When search criteria are provided in the `options` parameter, only workflows matching those criteria are returned.
// If no criteria are specified, all workflows are returned.
//
// GET /rest/api/{2-3}/workflow/search
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow#bulk-get-workflows
func (w *WorkflowService) Search(ctx context.Context, options *model.WorkflowSearchCriteria, expand []string, transitionLinks bool) (*model.WorkflowReadResponseScheme, *model.ResponseScheme, error) {
	return w.internalClient.Search(ctx, options, expand, transitionLinks)
}

// Capabilities Get the list of workflow capabilities for a specific workflow using either the workflow ID, or the project and issue type ID pair.
//
// The response includes the scope of the workflow, defined as global/project-based, and a list of project types that the workflow is scoped to.
//
// It also includes all rules organised into their broad categories (conditions, validators, actions, triggers, screens) as well as the source location (Atlassian-provided, Connect, Forge).
//
// GET /rest/api/{2-3}/workflows/capabilities
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow#get-workflow-capabilities
func (w *WorkflowService) Capabilities(ctx context.Context, workflowID, projectID, issueTypeID string) (*model.WorkflowCapabilitiesScheme, *model.ResponseScheme, error) {
	return w.internalClient.Capabilities(ctx, workflowID, projectID, issueTypeID)
}

// Creates create workflows and related statuses.
//
// This method allows you to create a workflow by defining transition rules
// using the shapes detailed in the Atlassian REST API documentation.
// If no transition rules are specified, the default system transition rules
// will be used.
//
// POST /rest/api/{2-3}/workflows
//
// For more details, refer to:
// https://docs.go-atlassian.io/jira-software-cloud/workflow#bulk-create-workflows
//
// Example:
//
//	payload := &models.WorkflowCreatesPayload{
//		Scope: &models.WorkflowScopeScheme{Type: "GLOBAL"},
//		// The workflows to create, with transition rules.
//	}
//
//	response, respScheme, err := client.Workflow.Creates(ctx, payload)
//	if err != nil {
//		log.Fatalf("Failed to create workflow: %v", err)
//	}
//
//	fmt.Printf("Workflow created with ID: %s", response.ID)
func (w *WorkflowService) Creates(ctx context.Context, payload *model.WorkflowCreatesPayload) (*model.WorkflowCreateResponseScheme, *model.ResponseScheme, error) {
	return w.internalClient.Creates(ctx, payload)
}

// ValidateCreateWorkflows validates workflows before creating them.
//
// This method allows you to validate the configuration of one or more workflows
// before they are created in Jira. It helps ensure that the workflows adhere
// to the defined rules and constraints.
//
// The validation checks will include all aspects of the workflows, such as transitions,
// statuses, and any related conditions or validators.
//
// POST /rest/api/{2-3}/workflows/create/validation
//
// For more details, refer to:
// https://docs.go-atlassian.io/jira-software-cloud/workflow#validate-create-workflows
//
// Example:
//
//	 ctx := context.Background()
//	 payload := &model.ValidationOptionsForCreateScheme{
//	     Payload: payload,
//	    Options: &models.ValidationOptionsLevelScheme{
//			  Levels: []string{"ERROR", "WARNING"},
//		  }
//
//	 validationResult, response, err := client.Workflow.ValidateCreateWorkflows(ctx, payload)
//	 if err != nil {
//	     log.Fatalf("Failed to validate workflow: %v", err)
//	 }
//
//	 if validationResult != nil {
//	     log.Printf("Workflow validation errors: %v", validationResult.Errors)
//	 }
//
//	 log.Printf("Validation response: %v", response)
func (w *WorkflowService) ValidateCreateWorkflows(ctx context.Context, payload *model.ValidationOptionsForCreateScheme) (*model.WorkflowValidationErrorListScheme, *model.ResponseScheme, error) {
	return w.internalClient.ValidateCreateWorkflows(ctx, payload)
}

// Updates updates workflows.
//
// This method allows you to update workflows by providing a payload containing the details
// of the updates. You can expand specific fields using the 'expand' parameter.
//
// The update follows the API detailed in the Atlassian documentation.
//
// POST /rest/api/{2-3}/workflows/update
//
// For more details, refer to:
// https://docs.go-atlassian.io/jira-software-cloud/workflow#bulk-update-workflows
func (w *WorkflowService) Updates(ctx context.Context, payload *model.WorkflowUpdatesPayloadScheme, expand []string) (*model.WorkflowUpdateResponseScheme, *model.ResponseScheme, error) {
	return w.internalClient.Updates(ctx, payload, expand)
}

// ValidateUpdateWorkflows validates the update of one or more workflows.
//
// This method allows you to validate changes to workflows before they are applied.
// The validation checks for potential issues that could prevent the workflows from being updated successfully.
//
// The validation process will check for conditions such as:
//
//	1.Whether the transitions are valid.
//	2. Whether the status and transition names are unique within the workflow.
//	3. If the validation fails, a list of validation errors is returned, which should be resolved before applying the changes.
//
// POST /rest/api/{2-3}/workflows/update/validation
//
// For more details, refer to:
// https://docs.go-atlassian.io/jira-software-cloud/workflow#validate-update-workflows
//
// Example:
//
//	options := &model.ValidationOptionsForUpdateScheme{
//	    // populate the validation options here
//	}
//
//	result, response, err := client.Workflow.ValidateUpdateWorkflows(ctx, options)
//	if err != nil {
//	    log.Fatalf("Validation failed: %v", err)
//	}
//
//	if len(result.Errors) > 0 {
//	    log.Printf("Validation errors: %v", result.Errors)
//	} else {
//	    log.Println("Validation passed, you can proceed with the update.")
//	}
func (w *WorkflowService) ValidateUpdateWorkflows(ctx context.Context, payload *model.ValidationOptionsForUpdateScheme) (*model.WorkflowValidationErrorListScheme, *model.ResponseScheme, error) {
	return w.internalClient.ValidateUpdateWorkflows(ctx, payload)
}

type internalWorkflowImpl struct {
	c       service.Connector
	version string
}

func (i *internalWorkflowImpl) Search(ctx context.Context, options *model.WorkflowSearchCriteria, expand []string, transitionLinks bool) (*model.WorkflowReadResponseScheme, *model.ResponseScheme, error) {
	params := url.Values{}

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	if transitionLinks {
		params.Add("useTransitionLinksFormat", "true")
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/workflows", i.version))

	if params.Encode() != "" {
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

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/workflows/capabilities", i.version))

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

func (i *internalWorkflowImpl) Creates(ctx context.Context, payload *model.WorkflowCreatesPayload) (*model.WorkflowCreateResponseScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/workflows/create", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	result := new(model.WorkflowCreateResponseScheme)
	response, err := i.c.Call(request, result)
	if err != nil {
		return nil, response, err
	}

	return result, response, nil
}

func (i *internalWorkflowImpl) ValidateCreateWorkflows(ctx context.Context, payload *model.ValidationOptionsForCreateScheme) (*model.WorkflowValidationErrorListScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/workflows/create/validation", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	errorList := new(model.WorkflowValidationErrorListScheme)
	response, err := i.c.Call(request, errorList)
	if err != nil {
		return nil, response, err
	}

	return errorList, response, nil
}

func (i *internalWorkflowImpl) Updates(ctx context.Context, payload *model.WorkflowUpdatesPayloadScheme, expand []string) (*model.WorkflowUpdateResponseScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/workflows/update", i.version))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", payload)
	if err != nil {
		return nil, nil, err
	}

	result := new(model.WorkflowUpdateResponseScheme)
	response, err := i.c.Call(request, result)
	if err != nil {
		return nil, response, err
	}

	return result, response, nil
}

func (i *internalWorkflowImpl) ValidateUpdateWorkflows(ctx context.Context, payload *model.ValidationOptionsForUpdateScheme) (*model.WorkflowValidationErrorListScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/workflows/update/validation", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	errorList := new(model.WorkflowValidationErrorListScheme)
	response, err := i.c.Call(request, errorList)
	if err != nil {
		return nil, response, err
	}

	return errorList, response, nil
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
		return nil, model.ErrNoWorkflowID
	}

	endpoint := fmt.Sprintf("rest/api/%v/workflow/%v", i.version, workflowID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
