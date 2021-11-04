package confluence

import (
	"context"
	"fmt"
	"net/http"
)

type ContentPermissionService struct {
	client *Client
}

type CheckPermissionScheme struct {
	Subject   *PermissionSubjectScheme `json:"subject,omitempty"`
	Operation string                   `json:"operation,omitempty"`
}

type PermissionSubjectScheme struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
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
	payload *CheckPermissionScheme) (result *PermissionCheckResponseScheme, response *ResponseScheme, err error) {

	if len(contentID) == 0 {
		return nil, nil, notContentIDError
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

type PermissionCheckResponseScheme struct {
	HasPermission bool                            `json:"hasPermission"`
	Errors        []*PermissionCheckMessageScheme `json:"errors,omitempty"`
}

type PermissionCheckMessageScheme struct {
	Translation string `json:"translation"`
	Args        []struct {
	} `json:"args"`
}
