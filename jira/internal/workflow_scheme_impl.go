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

func NewWorkflowSchemeService(client service.Client, version string) (*WorkflowSchemeService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &WorkflowSchemeService{
		internalClient: &internalWorkflowSchemeImpl{c: client, version: version},
	}, nil
}

type WorkflowSchemeService struct {
	internalClient jira.WorkflowSchemeConnector
}

// Gets returns a paginated list of all workflow schemes, not including draft workflow schemes.
//
// GET /rest/api/{2-3}/workflowscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#gets-workflows-schemes
func (w *WorkflowSchemeService) Gets(ctx context.Context, startAt, maxResults int) (*model.WorkflowSchemePageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Gets(ctx, startAt, maxResults)
}

// Create creates a workflow scheme.
//
// POST /rest/api/{2-3}/workflowscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#create-workflows-scheme
func (w *WorkflowSchemeService) Create(ctx context.Context, payload *model.WorkflowSchemePayloadScheme) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {
	return w.internalClient.Create(ctx, payload)
}

// Get returns a workflow scheme
//
// GET /rest/api/{2-3}/workflowscheme/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#get-workflow-scheme
func (w *WorkflowSchemeService) Get(ctx context.Context, schemeId int, returnDraftIfExists bool) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {
	return w.internalClient.Get(ctx, schemeId, returnDraftIfExists)
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
func (w *WorkflowSchemeService) Update(ctx context.Context, schemeId int, payload *model.WorkflowSchemePayloadScheme) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {
	return w.internalClient.Update(ctx, schemeId, payload)
}

// Delete deletes a workflow scheme.
//
// Note that a workflow scheme cannot be deleted if it is active (that is, being used by at least one project).
//
// DELETE /rest/api/{2-3}/workflowscheme/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/scheme#delete-workflow-scheme
func (w *WorkflowSchemeService) Delete(ctx context.Context, schemeId int) (*model.ResponseScheme, error) {
	return w.internalClient.Delete(ctx, schemeId)
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
func (w *WorkflowSchemeService) Associations(ctx context.Context, projectIds []int) (*model.WorkflowSchemeAssociationPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Associations(ctx, projectIds)
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
func (w *WorkflowSchemeService) Assign(ctx context.Context, schemeId, projectId string) (*model.ResponseScheme, error) {
	return w.internalClient.Assign(ctx, schemeId, projectId)
}

type internalWorkflowSchemeImpl struct {
	c       service.Client
	version string
}

func (i *internalWorkflowSchemeImpl) Gets(ctx context.Context, startAt, maxResults int) (*model.WorkflowSchemePageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	endpoint := fmt.Sprintf("rest/api/%v/workflowscheme?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.WorkflowSchemePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalWorkflowSchemeImpl) Create(ctx context.Context, payload *model.WorkflowSchemePayloadScheme) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/workflowscheme", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	workflowScheme := new(model.WorkflowSchemeScheme)
	response, err := i.c.Call(request, workflowScheme)
	if err != nil {
		return nil, response, err
	}

	return workflowScheme, response, nil
}

func (i *internalWorkflowSchemeImpl) Get(ctx context.Context, schemeId int, returnDraftIfExists bool) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {

	if schemeId == 0 {
		return nil, nil, model.ErrNoWorkflowSchemeIDError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/workflowscheme/%v", i.version, schemeId))

	if returnDraftIfExists {

		params := url.Values{}
		params.Add("returnDraftIfExists", "true")

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	workflowScheme := new(model.WorkflowSchemeScheme)
	response, err := i.c.Call(request, workflowScheme)
	if err != nil {
		return nil, response, err
	}

	return workflowScheme, response, nil
}

func (i *internalWorkflowSchemeImpl) Update(ctx context.Context, schemeId int, payload *model.WorkflowSchemePayloadScheme) (*model.WorkflowSchemeScheme, *model.ResponseScheme, error) {

	if schemeId == 0 {
		return nil, nil, model.ErrNoWorkflowSchemeIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/workflowscheme/%v", i.version, schemeId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	workflowScheme := new(model.WorkflowSchemeScheme)
	response, err := i.c.Call(request, workflowScheme)
	if err != nil {
		return nil, response, err
	}

	return workflowScheme, response, nil
}

func (i *internalWorkflowSchemeImpl) Delete(ctx context.Context, schemeId int) (*model.ResponseScheme, error) {

	if schemeId == 0 {
		return nil, model.ErrNoWorkflowSchemeIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/workflowscheme/%v", i.version, schemeId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalWorkflowSchemeImpl) Associations(ctx context.Context, projectIds []int) (*model.WorkflowSchemeAssociationPageScheme, *model.ResponseScheme, error) {

	if len(projectIds) == 0 {
		return nil, nil, model.ErrNoProjectsError
	}

	params := url.Values{}
	for _, projectID := range projectIds {
		params.Add("projectId", strconv.Itoa(projectID))
	}

	endpoint := fmt.Sprintf("rest/api/%v/workflowscheme/project?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	mapping := new(model.WorkflowSchemeAssociationPageScheme)
	response, err := i.c.Call(request, mapping)
	if err != nil {
		return nil, response, err
	}

	return mapping, response, nil
}

func (i *internalWorkflowSchemeImpl) Assign(ctx context.Context, schemeId, projectId string) (*model.ResponseScheme, error) {

	if schemeId == "" {
		return nil, model.ErrNoWorkflowSchemeIDError
	}

	if projectId == "" {
		return nil, model.ErrNoProjectIDOrKeyError
	}

	payload := struct {
		WorkflowSchemeID string `json:"workflowSchemeId"`
		ProjectID        string `json:"projectId"`
	}{
		WorkflowSchemeID: schemeId,
		ProjectID:        projectId,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/workflowscheme/project", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
