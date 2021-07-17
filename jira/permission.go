package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PermissionService struct {
	client *Client
	Scheme *PermissionSchemeService
}

type PermissionScheme struct {
	Key         string `json:"key,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

// Gets Returns all permissions
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-permissions/#api-rest-api-3-permissions-get
func (p *PermissionService) Gets(ctx context.Context) (result []*PermissionScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/permissions"

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
			result = append(result, &PermissionScheme{
				Key:         key,
				Name:        data["name"].(string),
				Type:        data["type"].(string),
				Description: data["description"].(string),
			})
		}

	}

	return
}

type PermissionCheckPayload struct {
	GlobalPermissions  []string                        `json:"globalPermissions,omitempty"`
	AccountID          string                          `json:"accountId,omitempty"`
	ProjectPermissions []*BulkProjectPermissionsScheme `json:"projectPermissions,omitempty"`
}

type BulkProjectPermissionsScheme struct {
	Issues      []int    `json:"issues"`
	Projects    []int    `json:"projects"`
	Permissions []string `json:"permissions"`
}

type PermissionGrantsScheme struct {
	ProjectPermissions []struct {
		Permission string `json:"permission,omitempty"`
		Issues     []int  `json:"issues,omitempty"`
		Projects   []int  `json:"projects,omitempty"`
	} `json:"projectPermissions,omitempty"`
	GlobalPermissions []string `json:"globalPermissions,omitempty"`
}

// Check search the permissions linked to an accountID, then check if the user permissions.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-permissions/#api-rest-api-3-permissions-check-post
// Docs: N/A
func (p *PermissionService) Check(ctx context.Context, payload *PermissionCheckPayload) (result *PermissionGrantsScheme,
	response *ResponseScheme, err error) {

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	if len(payload.ProjectPermissions) == 0 {
		return nil, nil, fmt.Errorf("error!, the ProjectPermissions values is required by the Atlassian")
	}

	var endpoint = "/rest/api/3/permissions/check"

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
