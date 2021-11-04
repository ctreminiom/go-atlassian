package v3

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"net/http"
)

type PriorityService struct{ client *Client }

// Gets returns the list of all issue priorities.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/priorities#get-priorities
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-priorities/#api-rest-api-3-priority-get
func (p *PriorityService) Gets(ctx context.Context) (result []*models.PriorityScheme, response *ResponseScheme, err error) {

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
func (p *PriorityService) Get(ctx context.Context, priorityID string) (result *models.PriorityScheme, response *ResponseScheme,
	err error) {

	if len(priorityID) == 0 {
		return nil, nil, models.ErrNoPriorityIDError
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
