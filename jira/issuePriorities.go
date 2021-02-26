package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PriorityService struct{ client *Client }

type PriorityScheme struct {
	Self        string `json:"self,omitempty"`
	StatusColor string `json:"statusColor,omitempty"`
	Description string `json:"description,omitempty"`
	IconURL     string `json:"iconUrl,omitempty"`
	Name        string `json:"name,omitempty"`
	ID          string `json:"id,omitempty"`
}

// Returns the list of all issue priorities.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-priorities/#api-rest-api-3-priority-get
func (p *PriorityService) Gets(ctx context.Context) (result *[]PriorityScheme, response *Response, err error) {

	var endpoint = "rest/api/3/priority"
	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new([]PriorityScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

func (p *PriorityService) Get(ctx context.Context, priorityID string) (result *PriorityScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/priority/%v", priorityID)
	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(PriorityScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}
