package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type WorkflowStatusService struct{ client *Client }

type WorkflowStatusScheme struct {
	Self           string `json:"self"`
	Description    string `json:"description"`
	IconURL        string `json:"iconUrl"`
	Name           string `json:"name"`
	ID             string `json:"id"`
	StatusCategory struct {
		Self      string `json:"self"`
		ID        int    `json:"id"`
		Key       string `json:"key"`
		ColorName string `json:"colorName"`
		Name      string `json:"name"`
	} `json:"statusCategory"`
}

// Returns a list of all statuses associated with workflows.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-statuses/#api-rest-api-3-status-get
func (w *WorkflowStatusService) Gets(ctx context.Context) (result *[]WorkflowStatusScheme, response *Response, err error) {

	var endpoint = "rest/api/3/status"

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = w.client.Do(request)
	if err != nil {
		return
	}

	result = new([]WorkflowStatusScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns a status. The status must be associated with a workflow to be returned.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-statuses/#api-rest-api-3-status-idorname-get
func (w *WorkflowStatusService) Get(ctx context.Context, statusNameOrID string) (result *WorkflowStatusScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/status/%v", statusNameOrID)

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = w.client.Do(request)
	if err != nil {
		return
	}

	result = new(WorkflowStatusScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
