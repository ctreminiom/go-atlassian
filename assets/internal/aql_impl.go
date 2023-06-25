package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/assets"
	"net/http"
)

func NewAQLService(client service.Client) *AQLService {

	return &AQLService{
		internalClient: &internalAQLImpl{c: client},
	}
}

type AQLService struct {
	internalClient assets.AQLAssetConnector
}

// Filter search objects based on Assets Query Language (AQL)
//
// POST /jsm/assets/workspace/{workspaceId}/v1/aql/objects
//
// Deprecated. Please use Object.Filter() instead.
//
// https://docs.go-atlassian.io/jira-assets/aql#filter-objects
func (a *AQLService) Filter(ctx context.Context, workspaceID string, payload *model.AQLSearchParamsScheme) (*model.ObjectListScheme, *model.ResponseScheme, error) {
	return a.internalClient.Filter(ctx, workspaceID, payload)
}

type internalAQLImpl struct {
	c service.Client
}

func (i *internalAQLImpl) Filter(ctx context.Context, workspaceID string, payload *model.AQLSearchParamsScheme) (*model.ObjectListScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/aql/objects", workspaceID)

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	object := new(model.ObjectListScheme)
	response, err := i.c.Call(request, object)
	if err != nil {
		return nil, response, err
	}

	return object, response, nil
}
