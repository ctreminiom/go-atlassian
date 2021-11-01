package v3

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type WorkflowService struct {
	client *Client
	Scheme *WorkflowSchemeService
}

// Gets returns a paginated list of published classic workflows.
// When workflow names are specified, details of those workflows are returned.
// Otherwise, all published classic workflows are returned.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflows/#api-rest-api-3-workflow-search-get
func (w *WorkflowService) Gets(ctx context.Context, workflowNames, expand []string, startAt, maxResults int) (result *WorkflowPageScheme,
	response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, workflowName := range workflowNames {
		params.Add("workflowName", workflowName)
	}

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint = fmt.Sprintf("/rest/api/3/workflow/search?%v", params.Encode())

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
//    an active workflow.
//    a system workflow.
//    associated with any workflow scheme.
//    associated with any draft workflow scheme.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflows/#api-rest-api-3-workflow-entityid-delete
// NOTE: Experimental Method
func (w *WorkflowService) Delete(ctx context.Context, workflowID string) (response *ResponseScheme, err error) {

	if len(workflowID) == 0 {
		return nil, notWorkflowIDError
	}

	var endpoint = fmt.Sprintf("/rest/api/3/workflow/%v", workflowID)

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

type WorkflowPageScheme struct {
	Self       string            `json:"self,omitempty"`
	NextPage   string            `json:"nextPage,omitempty"`
	MaxResults int               `json:"maxResults,omitempty"`
	StartAt    int               `json:"startAt,omitempty"`
	Total      int               `json:"total,omitempty"`
	IsLast     bool              `json:"isLast,omitempty"`
	Values     []*WorkflowScheme `json:"values,omitempty"`
}

type WorkflowScheme struct {
	ID          *WorkflowPublishedIDScheme  `json:"id,omitempty"`
	Transitions []*WorkflowTransitionScheme `json:"transitions,omitempty"`
	Statuses    []*WorkflowStatusScheme     `json:"statuses,omitempty"`
	Description string                      `json:"description,omitempty"`
	IsDefault   bool                        `json:"isDefault,omitempty"`
}

type WorkflowPublishedIDScheme struct {
	Name     string `json:"name,omitempty"`
	EntityID string `json:"entityId,omitempty"`
}

type WorkflowTransitionScheme struct {
	ID          string                          `json:"id,omitempty"`
	Name        string                          `json:"name,omitempty"`
	Description string                          `json:"description,omitempty"`
	From        []string                        `json:"from,omitempty"`
	To          string                          `json:"to,omitempty"`
	Type        string                          `json:"type,omitempty"`
	Screen      *WorkflowTransitionScreenScheme `json:"screen,omitempty"`
	Rules       *WorkflowTransitionRulesScheme  `json:"rules,omitempty"`
}

type WorkflowTransitionScreenScheme struct {
	ID string `json:"id,omitempty"`
}

type WorkflowTransitionRulesScheme struct {
	Conditions    []*WorkflowTransitionRuleScheme `json:"conditions,omitempty"`
	Validators    []*WorkflowTransitionRuleScheme `json:"validators,omitempty"`
	PostFunctions []*WorkflowTransitionRuleScheme `json:"postFunctions,omitempty"`
}

type WorkflowTransitionRuleScheme struct {
	Type string `json:"type"`
}

type WorkflowStatusScheme struct {
	ID         string                          `json:"id"`
	Name       string                          `json:"name"`
	Properties *WorkflowStatusPropertiesScheme `json:"properties"`
}

type WorkflowStatusPropertiesScheme struct {
	IssueEditable bool `json:"issueEditable"`
}

var (
	notWorkflowIDError = fmt.Errorf("error!, please provide a valid entity ID of the workflow")
)
