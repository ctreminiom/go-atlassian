package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/assets"
	"net/http"
)

func NewIconService(client service.Client) *IconService {

	return &IconService{
		internalClient: &internalIconImpl{c: client},
	}
}

type IconService struct {
	internalClient assets.IconConnector
}

// Get loads a single asset icon by id.
//
// GET /jsm/assets/workspace/{workspaceId}/v1/icon/{id}
//
// https://docs.go-atlassian.io/jira-assets/icons#get-icon
func (i *IconService) Get(ctx context.Context, workspaceID, iconID string) (*model.IconScheme, *model.ResponseScheme, error) {
	return i.internalClient.Get(ctx, workspaceID, iconID)
}

// Global returns all global icons i.e. icons not associated with a particular object schema.
//
// GET /jsm/assets/workspace/{workspaceId}/v1/icon/global
//
// https://docs.go-atlassian.io/jira-assets/icons#get-global-icons
func (i *IconService) Global(ctx context.Context, workspaceID string) ([]*model.IconScheme, *model.ResponseScheme, error) {
	return i.internalClient.Global(ctx, workspaceID)
}

type internalIconImpl struct {
	c service.Client
}

func (i *internalIconImpl) Get(ctx context.Context, workspaceID, iconID string) (*model.IconScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	if iconID == "" {
		return nil, nil, model.ErrNoIconIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/icon/%v", workspaceID, iconID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	icon := new(model.IconScheme)
	response, err := i.c.Call(request, icon)
	if err != nil {
		return nil, response, err
	}

	return icon, response, nil
}

func (i *internalIconImpl) Global(ctx context.Context, workspaceID string) ([]*model.IconScheme, *model.ResponseScheme, error) {

	if workspaceID == "" {
		return nil, nil, model.ErrNoWorkspaceIDError
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/icon/global", workspaceID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	var icons []*model.IconScheme
	response, err := i.c.Call(request, &icons)
	if err != nil {
		return nil, response, err
	}

	return icons, response, nil
}
