package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/confluence"
	"net/http"
)

// NewPermissionService creates a new instance of PermissionService.
// It takes a service.Connector as input and returns a pointer to PermissionService.
func NewPermissionService(client service.Connector) *PermissionService {

	return &PermissionService{
		internalClient: &internalPermissionImpl{c: client},
	}
}

// PermissionService provides methods to interact with content permission operations in Confluence.
type PermissionService struct {
	internalClient confluence.ContentPermissionConnector
}

// Check if a user or a group can perform an operation to the specified content.
//
// The operation to check must be provided.
//
// The user’s account ID or the ID of the group can be provided in the subject to check permissions
// against a specified user or group.
//
// The following permission checks are done to make sure that the user or group has the proper access:
//
// 1. site permissions
//
// 2. space permissions
//
// 3. content restrictions
//
// POST /wiki/rest/api/content/{id}/permission/check
//
// https://docs.go-atlassian.io/confluence-cloud/content/permissions#check-content-permissions
func (p *PermissionService) Check(ctx context.Context, contentID string, payload *model.CheckPermissionScheme) (*model.PermissionCheckResponseScheme, *model.ResponseScheme, error) {
	return p.internalClient.Check(ctx, contentID, payload)
}

type internalPermissionImpl struct {
	c service.Connector
}

func (i *internalPermissionImpl) Check(ctx context.Context, contentID string, payload *model.CheckPermissionScheme) (*model.PermissionCheckResponseScheme, *model.ResponseScheme, error) {

	if contentID == "" {
		return nil, nil, model.ErrNoContentID
	}

	endpoint := fmt.Sprintf("wiki/rest/api/content/%v/permission/check", contentID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	checker := new(model.PermissionCheckResponseScheme)
	response, err := i.c.Call(request, checker)
	if err != nil {
		return nil, response, err
	}

	return checker, response, nil
}
