package jira

import (
	"context"
	"fmt"
	"net/http"
)

type ResolutionService struct{ client *Client }

type IssueResolutionScheme struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Gets returns a list of all issue resolution values.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/resolutions#get-resolutions
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-resolutions/#api-rest-api-3-resolution-get
func (r *ResolutionService) Gets(ctx context.Context) (result []*IssueResolutionScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/resolution"
	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = r.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns an issue resolution value.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/resolutions#get-resolution
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-resolutions/#api-rest-api-3-resolution-id-get
func (r *ResolutionService) Get(ctx context.Context, resolutionID string) (result *IssueResolutionScheme,
	response *ResponseScheme, err error) {

	if len(resolutionID) == 0 {
		return nil, nil, notResolutionIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/resolution/%v", resolutionID)

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

var (
	notResolutionIDError = fmt.Errorf("error, please provide a valid resolutionID value")
)
