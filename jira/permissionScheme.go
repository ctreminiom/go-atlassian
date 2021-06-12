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

type PermissionSchemePageScheme struct {
	PermissionSchemes []*PermissionSchemeScheme `json:"permissionSchemes,omitempty"`
}

// Returns all permission schemes.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#get-all-permission-schemes
func (p *PermissionSchemeService) Gets(ctx context.Context) (result *PermissionSchemePageScheme, response *Response, err error) {

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

	result = new(PermissionSchemePageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type PermissionSchemeScheme struct {
	Expand      string                         `json:"expand,omitempty"`
	ID          int                            `json:"id,omitempty"`
	Self        string                         `json:"self,omitempty"`
	Name        string                         `json:"name,omitempty"`
	Description string                         `json:"description,omitempty"`
	Permissions []*PermissionGrantScheme       `json:"permissions,omitempty"`
	Scope       *TeamManagedProjectScopeScheme `json:"scope,omitempty"`
}

type PermissionScopeItemScheme struct {
	Self            string                 `json:"self,omitempty"`
	ID              string                 `json:"id,omitempty"`
	Key             string                 `json:"key,omitempty"`
	Name            string                 `json:"name,omitempty"`
	ProjectTypeKey  string                 `json:"projectTypeKey,omitempty"`
	Simplified      bool                   `json:"simplified,omitempty"`
	ProjectCategory *ProjectCategoryScheme `json:"projectCategory,omitempty"`
}

// Returns a permission scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#get-permission-scheme
func (p *PermissionSchemeService) Get(ctx context.Context, permissionSchemeID int) (result *PermissionSchemeScheme, response *Response, err error) {

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

	result = new(PermissionSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes a permission scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#delete-permission-scheme
func (p *PermissionSchemeService) Delete(ctx context.Context, permissionSchemeID int) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/permissionscheme/%v", permissionSchemeID)

	request, err := p.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Creates a new permission scheme.
// You can create a permission scheme with or without defining a set of permission grants.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#create-permission-scheme
func (p *PermissionSchemeService) Create(ctx context.Context, payload *PermissionSchemeScheme) (result *PermissionSchemeScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, please provide a PermissionSchemeScheme pointer")
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

	result = new(PermissionSchemeScheme)
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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#update-permission-scheme
func (p *PermissionSchemeService) Update(ctx context.Context, schemeID int, payload *PermissionSchemeScheme) (result *PermissionSchemeScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, please provide a payload pointer")
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

	result = new(PermissionSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
