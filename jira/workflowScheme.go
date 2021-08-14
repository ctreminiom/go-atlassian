package jira

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type WorkflowSchemeService struct{ client *Client }

// Gets returns a paginated list of all workflow schemes, not including draft workflow schemes.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-schemes/#api-rest-api-3-workflowscheme-get
func (w *WorkflowSchemeService) Gets(ctx context.Context, startAt, maxResults int) (result *WorkflowSchemePageScheme,
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
func (w *WorkflowSchemeService) Create(ctx context.Context, payload *WorkflowSchemePayloadScheme) (result *WorkflowSchemeScheme,
	response *ResponseScheme, err error) {

	var endpoint = "/rest/api/3/workflowscheme"

	payloadAsReader, err := transformStructToReader(&payload)
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
func (w *WorkflowSchemeService) Get(ctx context.Context, workflowSchemeID int, isExits bool) (result *WorkflowSchemeScheme,
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
func (w *WorkflowSchemeService) Update(ctx context.Context, workflowSchemeID int, payload *WorkflowSchemePayloadScheme) (result *WorkflowSchemeScheme,
	response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("/rest/api/3/workflowscheme/%v", workflowSchemeID)

	payloadAsReader, err := transformStructToReader(&payload)
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

type WorkflowSchemePayloadScheme struct {
	DefaultWorkflow   string      `json:"defaultWorkflow,omitempty"`
	Name              string      `json:"name,omitempty"`
	Description       string      `json:"description,omitempty"`
	IssueTypeMappings interface{} `json:"issueTypeMappings,omitempty"`
}

type WorkflowSchemePageScheme struct {
	Self       string                  `json:"self,omitempty"`
	NextPage   string                  `json:"nextPage,omitempty"`
	MaxResults int                     `json:"maxResults,omitempty"`
	StartAt    int                     `json:"startAt,omitempty"`
	Total      int                     `json:"total,omitempty"`
	IsLast     bool                    `json:"isLast,omitempty"`
	Values     []*WorkflowSchemeScheme `json:"values,omitempty"`
}

type WorkflowSchemeScheme struct {
	ID                  int         `json:"id,omitempty"`
	Name                string      `json:"name,omitempty"`
	Description         string      `json:"description,omitempty"`
	DefaultWorkflow     string      `json:"defaultWorkflow,omitempty"`
	Draft               bool        `json:"draft,omitempty"`
	LastModifiedUser    *UserScheme `json:"lastModifiedUser,omitempty"`
	LastModified        string      `json:"lastModified,omitempty"`
	Self                string      `json:"self,omitempty"`
	UpdateDraftIfNeeded bool        `json:"updateDraftIfNeeded,omitempty"`
}
