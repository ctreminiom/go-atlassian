package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewResolutionService creates a new instance of ResolutionService.
func NewResolutionService(client service.Connector, version string) (*ResolutionService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ResolutionService{
		internalClient: &internalResolutionImpl{c: client, version: version},
	}, nil
}

// ResolutionService provides methods to manage issue resolutions in Jira Service Management.
type ResolutionService struct {
	// internalClient is the connector interface for resolution operations.
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
		return nil, nil, model.ErrNoResolutionID
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
