package assets

import (
	"context"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// AQLAssetConnector represents the assets AQL search endpoints.
// Use it to execute AQL searches.
type AQLAssetConnector interface {

	// Filter search objects based on Assets Query Language (AQL)
	//
	// POST /jsm/assets/workspace/{workspaceID}/v1/aql/objects
	//
	// Deprecated. Please use Object.Filter() instead.
	//
	// https://docs.go-atlassian.io/jira-assets/aql#filter-objects
	Filter(ctx context.Context, workspaceID string, parameters *models.AQLSearchParamsScheme) (*models.ObjectListScheme, *models.ResponseScheme, error)
}
