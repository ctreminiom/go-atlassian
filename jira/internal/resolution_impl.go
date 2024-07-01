package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
)

func NewResolutionService(client service.Connector, version string) (*ResolutionService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ResolutionService{
		internalClient: &internalResolutionImpl{c: client, version: version},
	}, nil
}

type ResolutionService struct {
	internalClient jira.ResolutionConnector
}

// Gets returns a list of all issue resolution values.
//
// GET /rest/api/{2-3}/resolution
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/resolutions#get-resolutions
func (r *ResolutionService) Gets(ctx context.Context) ([]*model.ResolutionScheme, *model.ResponseScheme, error) {
	return r.internalClient.Gets(ctx)
}

// Get returns an issue resolution value.
//
// GET /rest/api/{2-3}/resolution/{resolutionID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/resolutions#get-resolution
func (r *ResolutionService) Get(ctx context.Context, resolutionID string) (*model.ResolutionScheme, *model.ResponseScheme, error) {
	return r.internalClient.Get(ctx, resolutionID)
}

type internalResolutionImpl struct {
	c       service.Connector
	version string
}

func (i *internalResolutionImpl) Gets(ctx context.Context) ([]*model.ResolutionScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/resolution", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var resolutions []*model.ResolutionScheme
	response, err := i.c.Call(request, &resolutions)
	if err != nil {
		return nil, response, err
	}

	return resolutions, response, nil
}

func (i *internalResolutionImpl) Get(ctx context.Context, resolutionID string) (*model.ResolutionScheme, *model.ResponseScheme, error) {

	if resolutionID == "" {
		return nil, nil, model.ErrNoResolutionIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/resolution/%v", i.version, resolutionID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	resolution := new(model.ResolutionScheme)
	response, err := i.c.Call(request, resolution)
	if err != nil {
		return nil, response, err
	}

	return resolution, response, nil
}
