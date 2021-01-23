package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PermissionSchemeService struct{ client *Client }

type PermissionSchemesScheme struct {
	PermissionSchemes []struct {
		Expand      string `json:"expand"`
		ID          int    `json:"id"`
		Self        string `json:"self"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Permissions []struct {
			ID     int    `json:"id"`
			Self   string `json:"self"`
			Holder struct {
				Type      string `json:"type"`
				Parameter string `json:"parameter"`
				Expand    string `json:"expand"`
			} `json:"holder"`
			Permission string `json:"permission"`
		} `json:"permissions"`
	} `json:"permissionSchemes"`
}

func (p *PermissionSchemeService) Gets(ctx context.Context) (result *PermissionSchemesScheme, response *Response, err error) {

	var endpoint = "rest/api/3/permissionscheme"

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(PermissionSchemesScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type PermissionScheme struct {
	ID          int    `json:"id"`
	Self        string `json:"self"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Permissions []struct {
		ID     int    `json:"id"`
		Self   string `json:"self"`
		Holder struct {
			Type      string `json:"type"`
			Parameter string `json:"parameter"`
			Expand    string `json:"expand"`
		} `json:"holder"`
		Permission string `json:"permission"`
	} `json:"permissions"`
}

// Returns a permission scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-permission-schemes/#api-rest-api-3-permissionscheme-schemeid-get
func (p *PermissionSchemeService) Get(ctx context.Context, permissionSchemeID string) (result *PermissionScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/permissionscheme/%v", permissionSchemeID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(PermissionScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
