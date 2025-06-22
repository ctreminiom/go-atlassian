package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/assets"
)

// NewIconService creates a new instance of IconService.
// It takes a service.Connector as input and returns a pointer to IconService.
func NewIconService(client service.Connector) *IconService {
	return &IconService{
		internalClient: &internalIconImpl{c: client},
	}
}

// IconService provides methods to interact with asset icons in Jira.
type IconService struct {
	// internalClient is the connector interface for icon operations.
	internalClient assets.IconConnector
}

// Get loads a single asset icon by id.
//
// GET /jsm/assets/workspace/{workspaceId}/v1/icon/{id}
//
// https://docs.go-atlassian.io/jira-assets/icons#get-icon
func (i *IconService) Get(ctx context.Context, workspaceID, iconID string) (*model.IconScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IconService).Get")
	defer span.End()

	return i.internalClient.Get(ctx, workspaceID, iconID)
}

// Global returns all global icons i.e. icons not associated with a particular object schema.
//
// GET /jsm/assets/workspace/{workspaceId}/v1/icon/global
//
// https://docs.go-atlassian.io/jira-assets/icons#get-global-icons
func (i *IconService) Global(ctx context.Context, workspaceID string) ([]*model.IconScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IconService).Global")
	defer span.End()

	return i.internalClient.Global(ctx, workspaceID)
}

type internalIconImpl struct {
	c service.Connector
}

func (i *internalIconImpl) Get(ctx context.Context, workspaceID, iconID string) (*model.IconScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIconImpl).Get")
	defer span.End()

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	if iconID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoIconID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/icon/%v", workspaceID, iconID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	icon := new(model.IconScheme)
	res, err := i.c.Call(req, icon)
	if err != nil {
		return nil, res, err
	}

	return icon, res, nil
}

func (i *internalIconImpl) Global(ctx context.Context, workspaceID string) ([]*model.IconScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIconImpl).Global")
	defer span.End()

	if workspaceID == "" {
		return nil, nil, fmt.Errorf("assets: %w", model.ErrNoWorkspaceID)
	}

	endpoint := fmt.Sprintf("jsm/assets/workspace/%v/v1/icon/global", workspaceID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var icons []*model.IconScheme
	res, err := i.c.Call(req, &icons)
	if err != nil {
		return nil, res, err
	}

	return icons, res, nil
}
