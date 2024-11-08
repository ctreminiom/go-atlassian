package internal

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/sm"
	"net/http"
)

// NewWorkSpaceService creates a new instance of WorkSpaceService.
// It takes a service.Connector and a version string as input and returns a pointer to WorkSpaceService.
func NewWorkSpaceService(client service.Connector, version string) *WorkSpaceService {

	return &WorkSpaceService{
		internalClient: &internalWorkSpaceImpl{c: client, version: version},
	}
}

// WorkSpaceService provides methods to interact with workspace operations in Jira Service Management.
type WorkSpaceService struct {
	// internalClient is the connector interface for workspace operations.
	internalClient sm.WorkSpaceConnector
}

// Gets retrieves workspace assets
//
// This endpoint is used to fetch the assets associated with a workspace.
//
// These assets may include knowledge base articles, request types, request fields, customer portals, queues, etc.
//
// GET /rest/servicedeskapi/assets/workspace
//
// https://docs.go-atlassian.io/jira-service-management/workspaces#get-workspaces
func (w *WorkSpaceService) Gets(ctx context.Context) (*model.WorkSpacePageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Gets(ctx)
}

type internalWorkSpaceImpl struct {
	c       service.Connector
	version string
}

func (i *internalWorkSpaceImpl) Gets(ctx context.Context) (*model.WorkSpacePageScheme, *model.ResponseScheme, error) {

	endpoint := "/rest/servicedeskapi/assets/workspace"

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.WorkSpacePageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil

}
