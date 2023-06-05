package sm

import (
	"context"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type WorkSpaceConnector interface {

	// Gets retrieves workspace assets
	//
	// This endpoint is used to fetch the assets associated with a workspace.
	//
	// These assets may include knowledge base articles, request types, request fields, customer portals, queues, etc.
	//
	// GET /rest/servicedeskapi/assets/workspace
	Gets(ctx context.Context) (*models.WorkSpacePageScheme, *models.ResponseScheme, error)
}
