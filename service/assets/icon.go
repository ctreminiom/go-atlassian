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
	// GET /jsm/assets/workspace/{workspaceId}/v1/icon/{id}
	Get(ctx context.Context, workspaceID, iconID string) (*models.IconScheme, *models.ResponseScheme, error)

	// Global returns all global icons i.e. icons not associated with a particular object schema.
	//
	// GET /jsm/assets/workspace/{workspaceId}/v1/icon/global
	Global(ctx context.Context, workspaceID string) ([]*models.IconScheme, *models.ResponseScheme, error)
}
