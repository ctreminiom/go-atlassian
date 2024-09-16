package assets

import (
	"context"

	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// IconConnector represents the assets icons endpoints.
// Use it to search and get asset icons.
type IconConnector interface {

	// Get loads a single asset icon by id.
	//
	// GET /jsm/assets/workspace/{workspaceID}/v1/icon/{id}
	//
	// https://docs.go-atlassian.io/jira-assets/icons#get-icon
	Get(ctx context.Context, workspaceID, iconID string) (*models.IconScheme, *models.ResponseScheme, error)

	// Global returns all global icons i.e. icons not associated with a particular object schema.
	//
	// GET /jsm/assets/workspace/{workspaceID}/v1/icon/global
	//
	// https://docs.go-atlassian.io/jira-assets/icons#get-global-icons
	Global(ctx context.Context, workspaceID string) ([]*models.IconScheme, *models.ResponseScheme, error)
}
