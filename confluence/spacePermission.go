package confluence

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
)

type SpacePermissionService struct{ client *Client }

// Add adds new permission to space. If the permission to be added is a group permission, the group can be identified by its group name or group id.
// Docs: https://docs.go-atlassian.io/confluence-cloud/space/permissions#add-new-permission-to-space
func (s *SpacePermissionService) Add(ctx context.Context, spaceKey string, payload *models.SpacePermissionPayloadScheme) (
	result *models.SpacePermissionV2Scheme, response *ResponseScheme, err error) {

	if len(spaceKey) == 0 {
		return nil, nil, models.ErrNoSpaceKeyError
	}

	endpoint := fmt.Sprintf("/wiki/rest/api/space/%v/permission", spaceKey)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Bulk adds new custom content permission to space.
// If the permission to be added is a group permission, the group can be identified by its group name or group id.
// Docs: https://docs.go-atlassian.io/confluence-cloud/space/permissions#add-new-custom-content-permission-to-space
func (s *SpacePermissionService) Bulk(ctx context.Context, spaceKey string, payload *models.SpacePermissionArrayPayloadScheme) (response *ResponseScheme, err error) {

	if len(spaceKey) == 0 {
		return nil, models.ErrNoSpaceKeyError
	}

	endpoint := fmt.Sprintf("/wiki/rest/api/space/%v/permission/custom-content", spaceKey)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.Call(request, nil)
	if err != nil {
		return response, err
	}

	return
}

// Remove removes a space permission.
// Note that removing Read Space permission for a user or group will remove all the space permissions for that user or group.
// Docs: https://docs.go-atlassian.io/confluence-cloud/space/permissions#remove-a-space-permission
func (s *SpacePermissionService) Remove(ctx context.Context, spaceKey string, permissionId int) (response *ResponseScheme, err error) {

	if len(spaceKey) == 0 {
		return nil, models.ErrNoSpaceKeyError
	}

	endpoint := fmt.Sprintf("/wiki/rest/api/space/%v/permission/%v", spaceKey, permissionId)

	request, err := s.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	response, err = s.client.Call(request, nil)
	if err != nil {
		return response, err
	}

	return
}
