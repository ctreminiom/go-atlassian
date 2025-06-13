package confluence

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type ContentPermissionConnector interface {

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
	Check(ctx context.Context, contentID string, payload *model.CheckPermissionScheme) (*model.PermissionCheckResponseScheme, *model.ResponseScheme, error)
}

type SpacePermissionConnector interface {

	// Add adds new permission to space.
	//
	// If the permission to be added is a group permission, the group can be identified by its group name or group id.
	//
	// Note: Apps cannot access this REST resource - including when utilizing user impersonation.
	//
	// POST /wiki/rest/api/space/{spaceKey}/permission
	//
	// https://docs.go-atlassian.io/confluence-cloud/space/permissions#add-new-permission-to-space
	Add(ctx context.Context, spaceKey string, payload *model.SpacePermissionPayloadScheme) (*model.SpacePermissionV2Scheme, *model.ResponseScheme, error)

	// Bulk adds new custom content permission to space.
	//
	// If the permission to be added is a group permission, the group can be identified by its group name or group id.
	//
	// Note: Only apps can access this REST resource and only make changes to the respective app permissions.
	//
	// POST /wiki/rest/api/space/{spaceKey}/permission/custom-content
	//
	// https://docs.go-atlassian.io/confluence-cloud/space/permissions#add-new-custom-content-permission-to-space
	Bulk(ctx context.Context, spaceKey string, payload *model.SpacePermissionArrayPayloadScheme) (*model.ResponseScheme, error)

	// Remove removes a space permission.
	//
	// Note that removing Read Space permission for a user or group will remove all the space permissions for that user or group.
	//
	// DELETE /wiki/rest/api/space/{spaceKey}/permission/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/space/permissions#remove-a-space-permission
	Remove(ctx context.Context, spaceKey string, permissionID int) (*model.ResponseScheme, error)
}
