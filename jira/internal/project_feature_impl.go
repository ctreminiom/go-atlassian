package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewProjectFeatureService creates a new instance of ProjectFeatureService.
func NewProjectFeatureService(client service.Connector, version string) (*ProjectFeatureService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ProjectFeatureService{
		internalClient: &internalProjectFeatureImpl{c: client, version: version},
	}, nil
}

// ProjectFeatureService provides methods to manage project features in Jira Service Management.
type ProjectFeatureService struct {
	// internalClient is the connector interface for project feature operations.
	internalClient jira.ProjectFeatureConnector
}

// Gets returns the list of features for a project.
//
// GET /rest/api/{2-3}/project/{projectKeyOrID}/features
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/features#get-project-features
func (p *ProjectFeatureService) Gets(ctx context.Context, projectKeyOrID string) (*model.ProjectFeaturesScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx, projectKeyOrID)
}

// Set sets the state of a project feature.
//
// PUT /rest/api/{2-3}/project/{projectKeyOrID}/features/{featureKey}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/features#set-project-feature-state
func (p *ProjectFeatureService) Set(ctx context.Context, projectKeyOrID, featureKey, state string) (*model.ProjectFeaturesScheme, *model.ResponseScheme, error) {
	return p.internalClient.Set(ctx, projectKeyOrID, featureKey, state)
}

type internalProjectFeatureImpl struct {
	c       service.Connector
	version string
}

func (i *internalProjectFeatureImpl) Gets(ctx context.Context, projectKeyOrID string) (*model.ProjectFeaturesScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/features", i.version, projectKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	features := new(model.ProjectFeaturesScheme)
	response, err := i.c.Call(request, features)
	if err != nil {
		return nil, response, err
	}

	return features, response, nil
}

func (i *internalProjectFeatureImpl) Set(ctx context.Context, projectKeyOrID, featureKey, state string) (*model.ProjectFeaturesScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	if featureKey == "" {
		return nil, nil, model.ErrNoProjectFeatureKey
	}

	if state == "" {
		return nil, nil, model.ErrNoProjectFeatureState
	}

	payload := map[string]interface{}{"state": state}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/features/%v", i.version, projectKeyOrID, featureKey)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	features := new(model.ProjectFeaturesScheme)
	response, err := i.c.Call(request, features)
	if err != nil {
		return nil, response, err
	}

	return features, response, nil
}
