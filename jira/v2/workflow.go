package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type WorkflowService struct {
	client *Client
	Scheme *WorkflowSchemeService
}

func (w *WorkflowService) Create(ctx context.Context, payload *models.WorkflowPayloadScheme) (result *models.WorkflowCreatedResponseScheme,
	response *ResponseScheme, err error) {

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := "/rest/api/2/workflow"
	request, err := w.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = w.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Gets returns a paginated list of published classic workflows.
// When workflow names are specified, details of those workflows are returned.
// Otherwise, all published classic workflows are returned.
func (w *WorkflowService) Gets(ctx context.Context, options *models.WorkflowSearchOptions, startAt, maxResults int) (result *models.WorkflowPageScheme,
	response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		for _, name := range options.WorkflowNames {
			params.Add("workflowName", name)
		}

		if options.QueryString != "" {
			params.Add("queryString", options.QueryString)
		}

		if options.OrderBy != "" {
			params.Add("orderBy", options.OrderBy)
		}

		params.Add("isActive", fmt.Sprintf("%T", options.IsActive))

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}

	}

	var endpoint = fmt.Sprintf("/rest/api/2/workflow/search?%v", params.Encode())

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = w.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes a workflow.
//
// The workflow cannot be deleted if it is:
//
// an active workflow.
// a system workflow.
// associated with any workflow scheme.
// associated with any draft workflow scheme.
func (w *WorkflowService) Delete(ctx context.Context, workflowID string) (response *ResponseScheme, err error) {

	if len(workflowID) == 0 {
		return nil, models.ErrNoWorkflowIDError
	}

	var endpoint = fmt.Sprintf("/rest/api/2/workflow/%v", workflowID)

	request, err := w.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = w.client.call(request, nil)
	if err != nil {
		return
	}

	return
}
