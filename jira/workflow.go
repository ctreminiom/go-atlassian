package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type WorkflowService struct {
	client   *Client
	Scheme   *WorkflowSchemeService
	Status   *WorkflowStatusService
	Category *WorkflowCategoryService
}

type WorkflowSearchPageScheme struct {
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		ID struct {
			Name string `json:"name"`
		} `json:"id"`
		Description string `json:"description"`
		Transitions []struct {
			ID          string   `json:"id"`
			Name        string   `json:"name"`
			Description string   `json:"description"`
			From        []string `json:"from"`
			To          string   `json:"to"`
			Type        string   `json:"type"`
			Screen      struct {
				ID string `json:"id"`
			} `json:"screen"`
			Rules struct {
				Conditions []struct {
					Type          string `json:"type"`
					Configuration struct {
						PermissionKey string `json:"permissionKey"`
					} `json:"configuration"`
				} `json:"conditions"`
				Validators []struct {
					Type          string `json:"type"`
					Configuration struct {
						IgnoreContext bool     `json:"ignoreContext"`
						ErrorMessage  string   `json:"errorMessage"`
						Fields        []string `json:"fields"`
					} `json:"configuration"`
				} `json:"validators"`
				PostFunctions []struct {
					Type string `json:"type"`
				} `json:"postFunctions"`
			} `json:"rules"`
		} `json:"transitions"`
		Statuses []struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			Properties struct {
				IssueEditable bool `json:"issueEditable"`
			} `json:"properties"`
		} `json:"statuses"`
	} `json:"values"`
}

// Returns a paginated list of published classic workflows.
// When workflow names are specified, details of those workflows are returned.
// Otherwise, all published classic workflows are returned.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflows/#api-rest-api-3-workflow-search-get
func (w *WorkflowService) Search(ctx context.Context, expands []string, startAt, maxResults int) (result *WorkflowSearchPageScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var expand string
	for index, value := range expands {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	var endpoint = fmt.Sprintf("rest/api/3/workflow/search?%v", params.Encode())

	request, err := w.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = w.client.Do(request)
	if err != nil {
		return
	}

	result = new(WorkflowSearchPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
