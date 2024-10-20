package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
	"net/http"
)

// NewSpacePermissionService creates a new instance of SpacePermissionService.
// It takes a service.Connector as input and returns a pointer to SpacePermissionService.
func NewSpacePermissionService(client service.Connector) *SpacePermissionService {

	return &SpacePermissionService{
		internalClient: &internalSpacePermissionImpl{c: client},
	}
}

// SpacePermissionService provides methods to interact with space permission operations in Confluence.
type SpacePermissionService struct {
	// internalClient is the connector interface for space permission operations.
	internalClient confluence.SpacePermissionConnector
}

// Add adds new permission to space.
//
// If the permission to be added is a group permission, the group can be identified by its group name or group id.
//
// Note: Apps cannot access this REST resource - including when utilizing user impersonation.
//
// POST /wiki/rest/api/space/{spaceKey}/permission
//
// https://docs.go-atlassian.io/confluence-cloud/space/permissions#add-new-permission-to-space
func (s *SpacePermissionService) Add(ctx context.Context, spaceKey string, payload *model.SpacePermissionPayloadScheme) (*model.SpacePermissionV2Scheme, *model.ResponseScheme, error) {
	return s.internalClient.Add(ctx, spaceKey, payload)
}

// Bulk adds new custom content permission to space.
//
// If the permission to be added is a group permission, the group can be identified by its group name or group id.
//
// Note: Only apps can access this REST resource and only make changes to the respective app permissions.
//
// POST /wiki/rest/api/space/{spaceKey}/permission/custom-content
//
// https://docs.go-atlassian.io/confluence-cloud/space/permissions#add-new-custom-content-permission-to-space
func (s *SpacePermissionService) Bulk(ctx context.Context, spaceKey string, payload *model.SpacePermissionArrayPayloadScheme) (*model.ResponseScheme, error) {
	return s.internalClient.Bulk(ctx, spaceKey, payload)
}

// Remove removes a space permission.
//
// Note that removing Read Space permission for a user or group will remove all the space permissions for that user or group.
//
// DELETE /wiki/rest/api/space/{spaceKey}/permission/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/space/permissions#remove-a-space-permission
func (s *SpacePermissionService) Remove(ctx context.Context, spaceKey string, permissionID int) (*model.ResponseScheme, error) {
	return s.internalClient.Remove(ctx, spaceKey, permissionID)
}

type internalSpacePermissionImpl struct {
	c service.Connector
}

func (i *internalSpacePermissionImpl) Add(ctx context.Context, spaceKey string, payload *model.SpacePermissionPayloadScheme) (*model.SpacePermissionV2Scheme, *model.ResponseScheme, error) {

	if spaceKey == "" {
		return nil, nil, model.ErrNoSpaceKey
	}

	endpoint := fmt.Sprintf("wiki/rest/api/space/%v/permission", spaceKey)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	permission := new(model.SpacePermissionV2Scheme)
	response, err := i.c.Call(request, permission)
	if err != nil {
		return nil, response, err
	}

	return permission, response, nil
}

func (i *internalSpacePermissionImpl) Bulk(ctx context.Context, spaceKey string, payload *model.SpacePermissionArrayPayloadScheme) (*model.ResponseScheme, error) {

	if spaceKey == "" {
		return nil, model.ErrNoSpaceKey
	}

	endpoint := fmt.Sprintf("wiki/rest/api/space/%v/permission/custom-content", spaceKey)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalSpacePermissionImpl) Remove(ctx context.Context, spaceKey string, permissionID int) (*model.ResponseScheme, error) {

	if spaceKey == "" {
		return nil, model.ErrNoSpaceKey
	}

	endpoint := fmt.Sprintf("wiki/rest/api/space/%v/permission/%v", spaceKey, permissionID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
