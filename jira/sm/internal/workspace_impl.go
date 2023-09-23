package internal

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/sm"
	"net/http"
)

func NewWorkSpaceService(client service.Connector, version string) *WorkSpaceService {

	return &WorkSpaceService{
		internalClient: &internalWorkSpaceImpl{c: client, version: version},
	}
}

type WorkSpaceService struct {
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
