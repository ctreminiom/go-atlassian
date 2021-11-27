package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strings"
)

type PermissionSchemeService struct {
	client *Client
	Grant  *PermissionGrantSchemeService
}

// Gets returns all permission schemes.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#get-all-permission-schemes
// Atlassian Docs; https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-permission-schemes/#api-rest-api-3-permissionscheme-get
func (p *PermissionSchemeService) Gets(ctx context.Context) (result *models.PermissionSchemePageScheme, response *ResponseScheme,
	err error) {

	var endpoint = "rest/api/3/permissionscheme"

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
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

// Get returns a permission scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#get-permission-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-permission-schemes/#api-rest-api-3-permissionscheme-schemeid-get
func (p *PermissionSchemeService) Get(ctx context.Context, permissionSchemeID int, expand []string) (
	result *models.PermissionSchemeScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/3/permissionscheme/%v", permissionSchemeID))

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

// Delete deletes a permission scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#delete-permission-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-permission-schemes/#api-rest-api-3-permissionscheme-schemeid-delete
func (p *PermissionSchemeService) Delete(ctx context.Context, permissionSchemeID int) (response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/permissionscheme/%v", permissionSchemeID)

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

// Create creates a new permission scheme.
// You can create a permission scheme with or without defining a set of permission grants.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#create-permission-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-permission-schemes/#api-rest-api-3-permissionscheme-post
func (p *PermissionSchemeService) Create(ctx context.Context, payload *models.PermissionSchemeScheme) (result *models.PermissionSchemeScheme,
	response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/permissionscheme"

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

// Update updates a permission scheme.
// Below are some important things to note when using this resource:
// 1. If a permissions list is present in the request, then it is set in the permission scheme, overwriting all existing grants.
// 2. If you want to update only the name and description, then do not send a permissions list in the request.
// 3. Sending an empty list will remove all permission grants from the permission scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#update-permission-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-permission-schemes/#api-rest-api-3-permissionscheme-schemeid-put
func (p *PermissionSchemeService) Update(ctx context.Context, schemeID int, payload *models.PermissionSchemeScheme) (
	result *models.PermissionSchemeScheme, response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/permissionscheme/%v", schemeID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := p.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
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
