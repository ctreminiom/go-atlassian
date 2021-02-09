package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PermissionSchemeService struct {
	client *Client
	Grant  *PermissionGrantSchemeService
}

type PermissionSchemesScheme struct {
	PermissionSchemes []PermissionScheme `json:"permissionSchemes"`
}

// Returns all permission schemes.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-permission-schemes/#api-rest-api-3-permissionscheme-get
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
func (p *PermissionSchemeService) Get(ctx context.Context, permissionSchemeID int) (result *PermissionScheme, response *Response, err error) {

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

// Deletes a permission scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-permission-schemes/#api-rest-api-3-permissionscheme-schemeid-delete
func (p *PermissionSchemeService) Delete(ctx context.Context, permissionSchemeID int) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/permissionscheme/%v", permissionSchemeID)

	request, err := p.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Creates a new permission scheme.
// You can create a permission scheme with or without defining a set of permission grants.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-permission-schemes/#api-rest-api-3-permissionscheme-post
func (p *PermissionSchemeService) Create(ctx context.Context, name, description string, permissions []PermissionGrantPayloadScheme) (result *PermissionScheme, response *Response, err error) {

	payload := struct {
		Permissions []PermissionGrantPayloadScheme `json:"permissions,omitempty"`
		Name        string                         `json:"name,omitempty"`
		Description string                         `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
		Permissions: permissions,
	}

	var endpoint = "rest/api/3/permissionscheme"

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

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

// Updates a permission scheme.
// Below are some important things to note when using this resource:
// 1. If a permissions list is present in the request, then it is set in the permission scheme, overwriting all existing grants.
// 2. If you want to update only the name and description, then do not send a permissions list in the request.
// 3. Sending an empty list will remove all permission grants from the permission scheme.
func (p *PermissionSchemeService) Update(ctx context.Context, schemeID int, name, description string, permissions []PermissionGrantPayloadScheme) (result *PermissionScheme, response *Response, err error) {

	payload := struct {
		Permissions []PermissionGrantPayloadScheme `json:"permissions,omitempty"`
		Name        string                         `json:"name,omitempty"`
		Description string                         `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
		Permissions: permissions,
	}

	var endpoint = fmt.Sprintf("rest/api/3/permissionscheme/%v", schemeID)

	request, err := p.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

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
