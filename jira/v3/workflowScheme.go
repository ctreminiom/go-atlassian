package v3

import (
	"context"
	"fmt"
	models2 "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type WorkflowSchemeService struct{ client *Client }

// Gets returns a paginated list of all workflow schemes, not including draft workflow schemes.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-schemes/#api-rest-api-3-workflowscheme-get
func (w *WorkflowSchemeService) Gets(ctx context.Context, startAt, maxResults int) (result *models2.WorkflowSchemePageScheme,
	response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("/rest/api/3/workflowscheme?%v", params.Encode())

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

// Create creates a workflow scheme.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-schemes/#api-rest-api-3-workflowscheme-post
func (w *WorkflowSchemeService) Create(ctx context.Context, payload *models2.WorkflowSchemePayloadScheme) (result *models2.WorkflowSchemeScheme,
	response *ResponseScheme, err error) {

	var endpoint = "/rest/api/3/workflowscheme"

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

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

// Get returns a workflow scheme.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-schemes/#api-rest-api-3-workflowscheme-id-get
func (w *WorkflowSchemeService) Get(ctx context.Context, workflowSchemeID int, isExits bool) (result *models2.WorkflowSchemeScheme,
	response *ResponseScheme, err error) {

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/rest/api/3/workflowscheme/%v", workflowSchemeID))

	query := url.Values{}
	if isExits {
		query.Add("returnDraftIfExists", "true")
	}

	if query.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", query.Encode()))
	}

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
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

// Update updates a workflow scheme, including the name, default workflow, issue type to project mappings,
// and more. If the workflow scheme is active (that is, being used by at least one project),
// then a draft workflow scheme is created or updated instead, provided that updateDraftIfNeeded is set to true.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-schemes/#api-rest-api-3-workflowscheme-id-put
func (w *WorkflowSchemeService) Update(ctx context.Context, workflowSchemeID int, payload *models2.WorkflowSchemePayloadScheme) (result *models2.WorkflowSchemeScheme,
	response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("/rest/api/3/workflowscheme/%v", workflowSchemeID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := w.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
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

// Delete deletes a workflow scheme.
// Note that a workflow scheme cannot be deleted if it is active (that is, being used by at least one project).
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-schemes/#api-rest-api-3-workflowscheme-id-delete
func (w *WorkflowSchemeService) Delete(ctx context.Context, workflowSchemeID int) (response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("/rest/api/3/workflowscheme/%v", workflowSchemeID)

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

// Associations returns a list of the workflow schemes associated with a list of projects.
// Each returned workflow scheme includes a list of the requested projects associated with it.
// Any team-managed or non-existent projects in the request are ignored and no errors are returned.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-scheme-project-associations/#api-rest-api-3-workflowscheme-project-get
func (w *WorkflowSchemeService) Associations(ctx context.Context, projectIDs []int) (result *models2.WorkflowSchemeAssociationPageScheme,
	response *ResponseScheme, err error) {

	if len(projectIDs) == 0 {
		return nil, nil, models2.ErrNoProjectsError
	}

	params := url.Values{}
	for _, projectID := range projectIDs {
		params.Add("projectId", strconv.Itoa(projectID))
	}

	endpoint := fmt.Sprintf("rest/api/3/workflowscheme/project?%v", params.Encode())

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

// Assign assigns a workflow scheme to a project.
// This operation is performed only when there are no issues in the project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-scheme-project-associations/#api-rest-api-3-workflowscheme-project-put
func (w *WorkflowSchemeService) Assign(ctx context.Context, workflowSchemeID, projectID string) (response *ResponseScheme, err error) {

	if len(projectID) == 0 {
		return nil, models2.ErrNoProjectIDError
	}

	if len(workflowSchemeID) == 0 {
		return nil, models2.ErrNoWorkflowSchemeIDError
	}

	payload := struct {
		WorkflowSchemeID string `json:"workflowSchemeId"`
		ProjectID        string `json:"projectId"`
	}{
		WorkflowSchemeID: workflowSchemeID,
		ProjectID:        projectID,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	endpoint := "rest/api/3/workflowscheme/project"

	request, err := w.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = w.client.call(request, nil)
	if err != nil {
		return
	}

	return
}
