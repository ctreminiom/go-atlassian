package confluence

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
)

type ContentPermissionService struct {
	client *Client
}

// Check if a user or a group can perform an operation to the specified content.
// The operation to check must be provided.
// The user’s account ID or the ID of the group can be provided in the subject to check permissions
//  against a specified user or group.
// The following permission checks are done to make sure that the user or group has the proper access:
// 1. site permissions
// 2. space permissions
// 3. content restrictions
func (c *ContentPermissionService) Check(ctx context.Context, contentID string,
	payload *model.CheckPermissionScheme) (result *model.PermissionCheckResponseScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, model.ErrNoContentIDError
	}

	var endpoint = fmt.Sprintf("/wiki/rest/api/content/%v/permission/check", contentID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := c.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}
