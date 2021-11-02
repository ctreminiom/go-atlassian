package v3

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"net/http"
)

type ResolutionService struct{ client *Client }

// Gets returns a list of all issue resolution values.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/resolutions#get-resolutions
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-resolutions/#api-rest-api-3-resolution-get
func (r *ResolutionService) Gets(ctx context.Context) (result []*models.ResolutionScheme, response *ResponseScheme, err error) {

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
func (r *ResolutionService) Get(ctx context.Context, resolutionID string) (result *models.ResolutionScheme,
	response *ResponseScheme, err error) {

	if len(resolutionID) == 0 {
		return nil, nil, models.ErrNoResolutionIDError
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
