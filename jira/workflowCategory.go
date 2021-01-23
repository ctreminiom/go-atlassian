package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type WorkflowCategoryService struct{ client *Client }

type WorkflowCategoryScheme struct {
	Self      string `json:"self"`
	ID        int    `json:"id"`
	Key       string `json:"key"`
	ColorName string `json:"colorName"`
	Name      string `json:"name,omitempty"`
}

// Returns a list of all status categories.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-status-categories/#api-rest-api-3-statuscategory-get
func (w *WorkflowCategoryService) Gets(ctx context.Context) (result *[]WorkflowCategoryScheme, response *Response, err error) {

	var endpoint = "rest/api/3/statuscategory"

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = w.client.Do(request)
	if err != nil {
		return
	}

	result = new([]WorkflowCategoryScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns a status category. Status categories provided a mechanism for categorizing statuses.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-status-categories/#api-rest-api-3-statuscategory-idorkey-get
func (w *WorkflowCategoryService) Get(ctx context.Context, statusNameOrID string) (result *WorkflowCategoryScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/statuscategory/%v", statusNameOrID)

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = w.client.Do(request)
	if err != nil {
		return
	}

	result = new(WorkflowCategoryScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
