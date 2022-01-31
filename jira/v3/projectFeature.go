package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
)

type ProjectFeatureService struct{ client *Client }

// Gets returns the list of features for a project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/features#get-project-features
func (p *ProjectFeatureService) Gets(ctx context.Context, projectKeyOrID string) (result *models.ProjectFeaturesScheme, response *ResponseScheme, err error) {

	if projectKeyOrID == "" {
		return nil, nil, models.ErrNoProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/features", projectKeyOrID)

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

// Set sets the state of a project feature.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/projects/features#set-project-feature-state
func (p *ProjectFeatureService) Set(ctx context.Context, projectKeyOrID, featureKey, state string) (result *models.ProjectFeaturesScheme, response *ResponseScheme, err error) {

	if projectKeyOrID == "" {
		return nil, nil, models.ErrNoProjectIDError
	}

	if featureKey == "" {
		return nil, nil, models.ErrNoProjectFeatureKeyError
	}

	payload := struct {
		State string `json:"state,omitempty"`
	}{
		State: state,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = fmt.Sprintf("rest/api/3/project/%v/features/%v", projectKeyOrID, featureKey)

	request, err := p.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
