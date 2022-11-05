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

func NewWorkflowStatusService(client service.Client, version string) (*WorkflowStatusService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &WorkflowStatusService{
		internalClient: &internalWorkflowStatusImpl{c: client, version: version},
	}, nil
}

type WorkflowStatusService struct {
	internalClient jira.WorkflowStatusConnector
}

// Gets returns a list of the statuses specified by one or more status IDs.
//
// GET /rest/api/{2-3}/statuses
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#gets-workflow-statuses
func (w *WorkflowStatusService) Gets(ctx context.Context, ids, expand []string) ([]*model.WorkflowStatusDetailScheme, *model.ResponseScheme, error) {
	return w.internalClient.Gets(ctx, ids, expand)
}

// Update updates statuses by ID.
//
// PUT /rest/api/{2-3}/statuses
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#update-workflow-statuses
func (w *WorkflowStatusService) Update(ctx context.Context, payload *model.WorkflowStatusPayloadScheme) (*model.ResponseScheme, error) {
	return w.internalClient.Update(ctx, payload)
}

// Create creates statuses for a global or project scope.
//
// POST /rest/api/{2-3}/statuses
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#create-workflow-statuses
func (w *WorkflowStatusService) Create(ctx context.Context, payload *model.WorkflowStatusPayloadScheme) ([]*model.WorkflowStatusDetailScheme, *model.ResponseScheme, error) {
	return w.internalClient.Create(ctx, payload)
}

// Delete deletes statuses by ID.
//
// DELETE /rest/api/{2-3}/statuses
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#delete-workflow-statuses
func (w *WorkflowStatusService) Delete(ctx context.Context, ids []string) (*model.ResponseScheme, error) {
	return w.internalClient.Delete(ctx, ids)
}

// Search returns a paginated list of statuses that match a search on name or project.
//
// GET /rest/api/{2-3}/statuses/search
//
// https://docs.go-atlassian.io/jira-software-cloud/workflow/status#search-workflow-statuses
func (w *WorkflowStatusService) Search(ctx context.Context, options *model.WorkflowStatusSearchParams, startAt, maxResults int) (*model.WorkflowStatusDetailPageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Search(ctx, options, startAt, maxResults)
}

func (w *WorkflowStatusService) Bulk(ctx context.Context) ([]*model.StatusDetailScheme, *model.ResponseScheme, error) {
	return w.internalClient.Bulk(ctx)
}

type internalWorkflowStatusImpl struct {
	c       service.Client
	version string
}

func (i *internalWorkflowStatusImpl) Bulk(ctx context.Context) ([]*model.StatusDetailScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("/rest/api/%v/status", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var statuses []*model.StatusDetailScheme
	response, err := i.c.Call(request, &statuses)
	if err != nil {
		return nil, response, err
	}

	return statuses, response, nil
}

func (i *internalWorkflowStatusImpl) Gets(ctx context.Context, ids, expand []string) ([]*model.WorkflowStatusDetailScheme, *model.ResponseScheme, error) {

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

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	var statuses []*model.WorkflowStatusDetailScheme
	response, err := i.c.Call(request, &statuses)
	if err != nil {
		return nil, response, err
	}

	return statuses, response, nil
}

func (i *internalWorkflowStatusImpl) Update(ctx context.Context, payload *model.WorkflowStatusPayloadScheme) (*model.ResponseScheme, error) {

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/statuses", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalWorkflowStatusImpl) Create(ctx context.Context, payload *model.WorkflowStatusPayloadScheme) ([]*model.WorkflowStatusDetailScheme, *model.ResponseScheme, error) {

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	if len(payload.Statuses) == 0 {
		return nil, nil, model.ErrNoWorkflowStatusesError
	}

	if payload.Scope == nil {
		return nil, nil, model.ErrNoWorkflowScopeError
	}

	endpoint := fmt.Sprintf("rest/api/%v/statuses", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	var workflowStatuses []*model.WorkflowStatusDetailScheme
	response, err := i.c.Call(request, &workflowStatuses)
	if err != nil {
		return nil, response, err
	}

	return workflowStatuses, response, nil
}

func (i *internalWorkflowStatusImpl) Delete(ctx context.Context, ids []string) (*model.ResponseScheme, error) {

	if len(ids) == 0 {
		return nil, model.ErrNoWorkflowStatusesError
	}

	params := url.Values{}
	for _, id := range ids {
		params.Add("id", id)
	}

	endpoint := fmt.Sprintf("rest/api/%v/statuses?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalWorkflowStatusImpl) Search(ctx context.Context, options *model.WorkflowStatusSearchParams, startAt, maxResults int) (*model.WorkflowStatusDetailPageScheme, *model.ResponseScheme, error) {

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

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.WorkflowStatusDetailPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}
