package v3

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type PermissionGrantSchemeService struct{ client *Client }

type PermissionSchemeGrantsScheme struct {
	Permissions []*PermissionGrantScheme `json:"permissions,omitempty"`
	Expand      string                   `json:"expand,omitempty"`
}

type PermissionGrantScheme struct {
	ID         int                          `json:"id,omitempty"`
	Self       string                       `json:"self,omitempty"`
	Holder     *PermissionGrantHolderScheme `json:"holder,omitempty"`
	Permission string                       `json:"permission,omitempty"`
}

type PermissionGrantHolderScheme struct {
	Type      string `json:"type,omitempty"`
	Parameter string `json:"parameter,omitempty"`
	Expand    string `json:"expand,omitempty"`
}

type PermissionGrantPayloadScheme struct {
	Holder     *PermissionGrantHolderScheme `json:"holder,omitempty"`
	Permission string                       `json:"permission,omitempty"`
}

// Create creates a permission grant in a permission scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#create-permission-grant
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-permission-schemes/#api-rest-api-3-permissionscheme-schemeid-permission-post
func (p *PermissionGrantSchemeService) Create(ctx context.Context, permissionSchemeID int, payload *PermissionGrantPayloadScheme) (
	result *PermissionGrantScheme, response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/permissionscheme/%v/permission", permissionSchemeID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

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

// Gets returns all permission grants for a permission scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#get-permission-scheme-grants
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-permission-schemes/#api-rest-api-3-permissionscheme-schemeid-permission-get
func (p *PermissionGrantSchemeService) Gets(ctx context.Context, permissionSchemeID int, expand []string) (
	result *PermissionSchemeGrantsScheme, response *ResponseScheme, err error) {

	if permissionSchemeID == 0 {
		return nil, nil, notPermissionSchemeIDError
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/3/permissionscheme/%v/permission", permissionSchemeID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns a permission grant.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#get-permission-scheme-grant
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-permission-schemes/#api-rest-api-3-permissionscheme-schemeid-permission-permissionid-get
func (p *PermissionGrantSchemeService) Get(ctx context.Context, permissionSchemeID, permissionGrantID int, expand []string) (
	result *PermissionGrantScheme, response *ResponseScheme, err error) {

	if permissionSchemeID == 0 {
		return nil, nil, notPermissionSchemeIDError
	}

	if permissionGrantID == 0 {
		return nil, nil, notPermissionGrantIDError
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/3/permissionscheme/%v/permission/%v", permissionSchemeID, permissionGrantID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes a permission grant from a permission scheme. See About permission schemes and grants for more details.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme/grant#delete-permission-scheme-grant
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-permission-schemes/#api-rest-api-3-permissionscheme-schemeid-permission-permissionid-delete
func (p *PermissionGrantSchemeService) Delete(ctx context.Context, permissionSchemeID, permissionGrantID int) (
	response *ResponseScheme, err error) {

	if permissionSchemeID == 0 {
		return nil, notPermissionSchemeIDError
	}

	if permissionGrantID == 0 {
		return nil, notPermissionGrantIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/permissionscheme/%v/permission/%v", permissionSchemeID, permissionGrantID)

	request, err := p.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

var (
	notPermissionSchemeIDError = fmt.Errorf("error, please provide a permissionSchemeID value")
	notPermissionGrantIDError  = fmt.Errorf("error, please provide a permissionGrantID value")
)
