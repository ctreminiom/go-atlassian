package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type FilterShareService struct{ client *Client }

type shareFilterScopeScheme struct {
	Scope string `json:"scope"`
}

// Returns the default sharing settings for new filters and dashboards for a user.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filter-sharing/#api-rest-api-3-filter-defaultsharescope-get
func (f *FilterShareService) Scope(ctx context.Context) (scope string, response *Response, err error) {

	var endpoint = "rest/api/3/filter/defaultShareScope"
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result := new(shareFilterScopeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return "", response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	scope = result.Scope
	return
}

// Sets the default sharing for new filters and dashboards for a user.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filter-sharing/#api-rest-api-3-filter-defaultsharescope-put
// Valid values: GLOBAL, AUTHENTICATED, PRIVATE
func (f *FilterShareService) SetScope(ctx context.Context, scope string) (response *Response, err error) {

	//Valid the share filter scope
	var (
		validScopeValuesAsList = []string{"GLOBAL", "AUTHENTICATED", "PRIVATE"}
		isValid                bool
	)

	for _, validScope := range validScopeValuesAsList {
		if validScope == scope {
			isValid = true
			break
		}
	}

	if !isValid {
		//Join the valid values and create the custom error
		var validScopeValuesAsString = strings.Join(validScopeValuesAsList, ",")
		return nil, fmt.Errorf("invalid scope, please provide one of the following: %v", validScopeValuesAsString)
	}

	var endpoint = "rest/api/3/filter/defaultShareScope"
	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, shareFilterScopeScheme{Scope: scope})
	if err != nil {
		return
	}

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	return
}

type ShareFilterPermissionScheme struct {
	ID      int    `json:"id,omitempty"`
	Type    string `json:"type,omitempty"`
	Project struct {
		Self       string `json:"self,omitempty"`
		ID         string `json:"id,omitempty"`
		Key        string `json:"key,omitempty"`
		Name       string `json:"name,omitempty"`
		AvatarUrls struct {
			Four8X48  string `json:"48x48,omitempty"`
			Two4X24   string `json:"24x24,omitempty"`
			One6X16   string `json:"16x16,omitempty"`
			Three2X32 string `json:"32x32,omitempty"`
		} `json:"avatarUrls,omitempty"`
		ProjectCategory struct {
			Self        string `json:"self,omitempty"`
			ID          string `json:"id,omitempty"`
			Name        string `json:"name,omitempty"`
			Description string `json:"description,omitempty"`
		} `json:"projectCategory,omitempty"`
		Simplified bool   `json:"simplified,omitempty"`
		Style      string `json:"style,omitempty"`
		Insight    struct {
			TotalIssueCount     int    `json:"totalIssueCount,omitempty"`
			LastIssueUpdateTime string `json:"lastIssueUpdateTime,omitempty"`
		} `json:"insight,omitempty"`
	} `json:"project,omitempty"`
	Role struct {
		Self        string `json:"self,omitempty"`
		Name        string `json:"name,omitempty"`
		ID          int    `json:"id,omitempty"`
		Description string `json:"description,omitempty"`
		Actors      []struct {
			ID          int    `json:"id,omitempty"`
			DisplayName string `json:"displayName,omitempty"`
			Type        string `json:"type,omitempty"`
			Name        string `json:"name,omitempty"`
			ActorGroup  struct {
				Name        string `json:"name,omitempty"`
				DisplayName string `json:"displayName,omitempty"`
			} `json:"actorGroup,omitempty"`
			ActorUser struct {
				AccountID string `json:"accountId,omitempty"`
			} `json:"actorUser,omitempty"`
		} `json:"actors,omitempty"`
		Scope struct {
			Type    string `json:"type,omitempty"`
			Project struct {
				ID   string `json:"id,omitempty"`
				Key  string `json:"key,omitempty"`
				Name string `json:"name,omitempty"`
			} `json:"project,omitempty"`
		} `json:"scope,omitempty"`
	} `json:"role,omitempty"`
	Group struct {
		Name string `json:"name,omitempty"`
		Self string `json:"self,omitempty"`
	} `json:"group,omitempty"`
}

// Returns the share permissions for a filter.
// A filter can be shared with groups, projects, all logged-in users, or the public.
// Sharing with all logged-in users or the public is known as a global share permission.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filter-sharing/#api-rest-api-3-filter-id-permission-get
func (f *FilterShareService) Gets(ctx context.Context, filterID int) (result *[]ShareFilterPermissionScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/filter/%v/permission", filterID)
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new([]ShareFilterPermissionScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

type PermissionFilterBodyScheme struct {
	Type          string `json:"type,omitempty"`
	ProjectID     string `json:"projectId,omitempty"`
	GroupName     string `json:"groupname,omitempty"`
	ProjectRoleID string `json:"projectRoleId,omitempty"`
}

// Add a share permissions to a filter.
// If you add a global share permission (one for all logged-in users or the public)
// it will overwrite all share permissions for the filter.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filter-sharing/#api-rest-api-3-filter-id-permission-post
func (f *FilterShareService) Add(ctx context.Context, filterID int, payload *PermissionFilterBodyScheme) (result *[]ShareFilterPermissionScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, payload value is nil, please provide a valid PermissionFilterBodyScheme pointer")
	}

	var endpoint = fmt.Sprintf("rest/api/3/filter/%v/permission", filterID)
	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new([]ShareFilterPermissionScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Returns a share permission for a filter.
// A filter can be shared with groups, projects, all logged-in users, or the public.
// Sharing with all logged-in users or the public is known as a global share permission.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filter-sharing/#api-rest-api-3-filter-id-permission-permissionid-get
func (f *FilterShareService) Get(ctx context.Context, filterID, permissionID int) (result *ShareFilterPermissionScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/filter/%v/permission/%v", filterID, permissionID)
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(ShareFilterPermissionScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Deletes a share permission from a filter.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-filter-sharing/#api-rest-api-3-filter-id-permission-permissionid-delete
func (f *FilterShareService) Delete(ctx context.Context, filterID, permissionID int) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/filter/%v/permission/%v", filterID, permissionID)
	request, err := f.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	return
}
