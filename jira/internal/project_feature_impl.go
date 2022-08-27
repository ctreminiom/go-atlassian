package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

func NewProjectFeatureService(client service.Client, version string) (*ProjectFeatureService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ProjectFeatureService{
		internalClient: &internalProjectFeatureImpl{c: client, version: version},
	}, nil
}

type ProjectFeatureService struct {
	internalClient jira.ProjectFeatureConnector
}

// Gets returns the list of features for a project.
//
// GET /rest/api/{2-3}/project/{projectIdOrKey}/features
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/features#get-project-features
func (p *ProjectFeatureService) Gets(ctx context.Context, projectKeyOrId string) (*model.ProjectFeaturesScheme, *model.ResponseScheme, error) {
	return p.internalClient.Gets(ctx, projectKeyOrId)
}

// Set sets the state of a project feature.
//
// PUT /rest/api/{2-3}/project/{projectIdOrKey}/features/{featureKey}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects/features#set-project-feature-state
func (p *ProjectFeatureService) Set(ctx context.Context, projectKeyOrId, featureKey, state string) (*model.ProjectFeaturesScheme, *model.ResponseScheme, error) {
	return p.internalClient.Set(ctx, projectKeyOrId, featureKey, state)
}

type internalProjectFeatureImpl struct {
	c       service.Client
	version string
}

func (i *internalProjectFeatureImpl) Gets(ctx context.Context, projectKeyOrId string) (*model.ProjectFeaturesScheme, *model.ResponseScheme, error) {

	if projectKeyOrId == "" {
		return nil, nil, model.ErrNoProjectIDOrKeyError
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/features", i.version, projectKeyOrId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

func (i *internalProjectFeatureImpl) Set(ctx context.Context, projectKeyOrId, featureKey, state string) (*model.ProjectFeaturesScheme, *model.ResponseScheme, error) {

	if projectKeyOrId == "" {
		return nil, nil, model.ErrNoProjectIDOrKeyError
	}

	if featureKey == "" {
		return nil, nil, model.ErrNoProjectFeatureKeyError
	}

	payload := struct {
		State string `json:"state,omitempty"`
	}{
		State: state,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/features/%v", i.version, projectKeyOrId, featureKey)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
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
