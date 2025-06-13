package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type ResolutionConnector interface {

	// Gets returns a list of all issue resolution values.
	//
	// GET /rest/api/{2-3}/resolution
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/resolutions#get-resolutions
	Gets(ctx context.Context) ([]*model.ResolutionScheme, *model.ResponseScheme, error)

	// Get returns an issue resolution value.
	//
	//
	// GET /rest/api/{2-3}/resolution/{resolutionID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/resolutions#get-resolution
	Get(ctx context.Context, resolutionID string) (*model.ResolutionScheme, *model.ResponseScheme, error)
}
