package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type FilterSharingConnector interface {

	// Scope returns the default sharing settings for new filters and dashboards for a user.
	//
	// GET /rest/api/{2-3}/filter/defaultShareScope
	//
	// https://docs.go-atlassian.io/jira-software-cloud/filters/sharing#get-default-share-scope
	Scope(ctx context.Context) (*model.ShareFilterScopeScheme, *model.ResponseScheme, error)

	// SetScope sets the default sharing for new filters and dashboards for a user.
	//
	// PUT /rest/api/{2-3}/filter/defaultShareScope
	//
	// https://docs.go-atlassian.io/jira-software-cloud/filters/sharing#set-default-share-scope
	SetScope(ctx context.Context, scope string) (*model.ResponseScheme, error)

	// Gets returns the share permissions for a filter.
	//
	// 1.A filter can be shared with groups, projects, all logged-in users, or the public.
	//
	// 2.Sharing with all logged-in users or the public is known as a global share permission.
	//
	// GET /rest/api/{2-3}/filter/{filterID}/permission
	//
	// https://docs.go-atlassian.io/jira-software-cloud/filters/sharing#get-share-permissions
	Gets(ctx context.Context, filterID int) ([]*model.SharePermissionScheme, *model.ResponseScheme, error)

	// Add a share permissions to a filter.
	//
	// If you add a global share permission (one for all logged-in users or the public)
	//
	// it will overwrite all share permissions for the filter.
	//
	// POST /rest/api/{2-3}/filter/{filterID}/permission
	//
	// https://docs.go-atlassian.io/jira-software-cloud/filters/sharing#add-share-permission
	Add(ctx context.Context, filterID int, payload *model.PermissionFilterPayloadScheme) ([]*model.SharePermissionScheme, *model.ResponseScheme, error)

	// Get returns a share permission for a filter.
	//
	// A filter can be shared with groups, projects, all logged-in users, or the public.
	//
	// Sharing with all logged-in users or the public is known as a global share permission.
	//
	// GET /rest/api/{2-3}/filter/{filterID}/permission/{permissionID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/filters/sharing#get-share-permission
	Get(ctx context.Context, filterID, permissionID int) (*model.SharePermissionScheme, *model.ResponseScheme, error)

	// Delete deletes a share permission from a filter.
	//
	// DELETE /rest/api/{2-3}/filter/{filterID}/permission/{permissionID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/filters/sharing#delete-share-permission
	Delete(ctx context.Context, filterID, permissionID int) (*model.ResponseScheme, error)
}
