package jira

import (
	"context"
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

// Gets returns the list of all issue priorities.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/priorities#get-priorities
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-priorities/#api-rest-api-3-priority-get
func (p *PriorityService) Gets(ctx context.Context) (result []*PriorityScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/priority"
	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns an issue priority.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/priorities#get-priority
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-priorities/#api-rest-api-3-priority-id-get
func (p *PriorityService) Get(ctx context.Context, priorityID string) (result *PriorityScheme, response *ResponseScheme,
	err error) {

	if len(priorityID) == 0 {
		return nil, nil, notPriorityIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/priority/%v", priorityID)
	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

var (
	notPriorityIDError = fmt.Errorf("error, please provide a valid priorityID value")
)
