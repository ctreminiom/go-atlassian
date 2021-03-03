package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type PermissionGrantSchemeService struct{ client *Client }

type PermissionSchemeGrantsScheme struct {
	Permissions []PermissionGrantScheme `json:"permissions"`
	Expand      string                  `json:"expand"`
}

type PermissionGrantScheme struct {
	ID     int    `json:"id"`
	Self   string `json:"self"`
	Holder struct {
		Type        string `json:"type"`
		Parameter   string `json:"parameter"`
		ProjectRole struct {
			Self        string `json:"self"`
			Name        string `json:"name"`
			ID          int    `json:"id"`
			Description string `json:"description"`
		} `json:"projectRole"`
		Expand string `json:"expand"`
	} `json:"holder"`
	Permission string `json:"permission"`
}

type PermissionGrantPayloadScheme struct {
	Holder     *PermissionGrantHolderPayloadScheme `json:"holder,omitempty"`
	Permission string                              `json:"permission,omitempty"`
}

type PermissionGrantHolderPayloadScheme struct {
	Parameter string `json:"parameter,omitempty"`
	Type      string `json:"type,omitempty"`
}

// Creates a permission grant in a permission scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#create-permission-grant
func (p *PermissionGrantSchemeService) Create(ctx context.Context, schemeID int, payload *PermissionGrantPayloadScheme) (result *PermissionGrantScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, please provide a PermissionGrantPayloadScheme pointer")
	}

	var endpoint = fmt.Sprintf("rest/api/3/permissionscheme/%v/permission", schemeID)

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

	result = new(PermissionGrantScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns all permission grants for a permission scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#get-permission-scheme-grants
func (p *PermissionGrantSchemeService) Gets(ctx context.Context, permissionSchemeID int, expands []string) (result *PermissionSchemeGrantsScheme, response *Response, err error) {

	params := url.Values{}

	var expand string
	for index, value := range expands {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	var endpoint string
	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("rest/api/3/permissionscheme/%v/permission?%v", permissionSchemeID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/api/3/permissionscheme/%v/permission", permissionSchemeID)
	}

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(PermissionSchemeGrantsScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns a permission grant.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#get-permission-scheme-grant
func (p *PermissionGrantSchemeService) Get(ctx context.Context, schemeID, permissionID int, expands []string) (result *PermissionGrantScheme, response *Response, err error) {

	params := url.Values{}

	var expand string
	for index, value := range expands {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	var endpoint string

	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("rest/api/3/permissionscheme/%v/permission/%v?%v", schemeID, permissionID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/api/3/permissionscheme/%v/permission/%v", schemeID, permissionID)
	}

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.Do(request)
	if err != nil {
		return
	}

	result = new(PermissionGrantScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes a permission grant from a permission scheme. See About permission schemes and grants for more details.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#delete-permission-scheme-grant
func (p *PermissionGrantSchemeService) Delete(ctx context.Context, schemeID, permissionID int) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/permissionscheme/%v/permission/%v", schemeID, permissionID)

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
