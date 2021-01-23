package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type WorkflowSchemeService struct{ client *Client }

type WorkflowSchemePageScheme struct {
	MaxResults int                    `json:"maxResults"`
	StartAt    int                    `json:"startAt"`
	Total      int                    `json:"total"`
	IsLast     bool                   `json:"isLast"`
	Values     []WorkflowSchemeScheme `json:"values"`
}

type WorkflowSchemeScheme struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	DefaultWorkflow   string `json:"defaultWorkflow"`
	IssueTypeMappings struct {
		Num10000 string `json:"10000"`
		Num10001 string `json:"10001"`
	} `json:"issueTypeMappings"`
	Draft bool   `json:"draft"`
	Self  string `json:"self"`
}

// Returns a paginated list of all workflow schemes, not including draft workflow schemes.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-schemes/#api-rest-api-3-workflowscheme-get
func (w *WorkflowSchemeService) Gets(ctx context.Context, startAt, maxResults int) (result *WorkflowSchemePageScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("rest/api/3/workflowscheme?%v", params.Encode())

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = w.client.Do(request)
	if err != nil {
		return
	}

	result = new(WorkflowSchemePageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns a workflow scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-schemes/#api-rest-api-3-workflowscheme-id-get
func (w *WorkflowSchemeService) Get(ctx context.Context, workflowSchemeID int) (result *WorkflowSchemeScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/workflowscheme/%v", workflowSchemeID)

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = w.client.Do(request)
	if err != nil {
		return
	}

	result = new(WorkflowSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
