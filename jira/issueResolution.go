package jira

import (
	"context"
	"encoding/json"
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

// Returns a list of all issue resolution values.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/resolutions#get-resolutions
func (r *ResolutionService) Gets(ctx context.Context) (result *[]IssueResolutionScheme, response *Response, err error) {

	var endpoint = "rest/api/3/resolution"
	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new([]IssueResolutionScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Returns an issue resolution value.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/resolutions#get-resolution
func (r *ResolutionService) Get(ctx context.Context, resolutionID string) (result *IssueResolutionScheme, response *Response, err error) {

	if len(resolutionID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid resolutionID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/resolution/%v", resolutionID)

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueResolutionScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}
