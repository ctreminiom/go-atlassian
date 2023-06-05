package internal

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/sm"
	"net/http"
)

func NewWorkSpaceService(client service.Client, version string) (*WorkSpaceService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &WorkSpaceService{
		internalClient: &internalWorkSpaceImpl{c: client, version: version},
	}, nil
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
func (w *WorkSpaceService) Gets(ctx context.Context) (*model.WorkSpacePageScheme, *model.ResponseScheme, error) {
	return w.internalClient.Gets(ctx)
}

type internalWorkSpaceImpl struct {
	c       service.Client
	version string
}

func (i *internalWorkSpaceImpl) Gets(ctx context.Context) (*model.WorkSpacePageScheme, *model.ResponseScheme, error) {

	endpoint := "/rest/servicedeskapi/assets/workspace"

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.WorkSpacePageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil

}
