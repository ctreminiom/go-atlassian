package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
)

type PermissionService struct {
	client *Client
	Scheme *PermissionSchemeService
}

// Gets Returns all permissions
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-permissions/#api-rest-api-2-permissions-get
func (p *PermissionService) Gets(ctx context.Context) (result []*models.PermissionScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/2/permissions"

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, nil)
	if err != nil {
		return
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(response.Bytes.Bytes(), &jsonMap)
	if err != nil {
		return
	}

	for key, value := range jsonMap["permissions"].(map[string]interface{}) {
		data, ok := value.(map[string]interface{})

		if ok {
			result = append(result, &models.PermissionScheme{
				Key:         key,
				Name:        data["name"].(string),
				Type:        data["type"].(string),
				Description: data["description"].(string),
			})
		}
	}

	return
}

// Check search the permissions linked to an accountID, then check if the user permissions.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-permissions/#api-rest-api-2-permissions-check-post
// Docs: N/A
func (p *PermissionService) Check(ctx context.Context, payload *models.PermissionCheckPayload) (result *models.PermissionGrantsScheme,
	response *ResponseScheme, err error) {

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	if len(payload.ProjectPermissions) == 0 {
		return nil, nil, fmt.Errorf("error!, the ProjectPermissions values is required by the Atlassian")
	}

	var endpoint = "/rest/api/2/permissions/check"

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
