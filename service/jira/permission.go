package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type PermissionConnector interface {

	// Gets returns all permissions, including: global permissions, project permissions and global permissions added by plugins.
	//
	// GET /rest/api/{2-3}/permissions
	//
	// TODO: Add/Create documentation
	Gets(ctx context.Context) ([]*model.PermissionScheme, *model.ResponseScheme, error)

	// Check search the permissions linked to an accountID, then check if the user permissions.
	//
	// POST /rest/api/{2-3}/permissions/check
	//
	// https://docs.go-atlassian.io/jira-software-cloud/permissions#check-permissions
	Check(ctx context.Context, payload *model.PermissionCheckPayload) (*model.PermissionGrantsScheme, *model.ResponseScheme, error)

	// Projects returns all the projects where the user is granted a list of project permissions.
	//
	// POST /rest/api/{2-3}/permissions/project
	//
	// TODO: Add/Create documentation
	Projects(ctx context.Context, permissions []string) (*model.PermittedProjectsScheme, *model.ResponseScheme, error)
}

type PermissionSchemeConnector interface {

	// Gets returns all permission schemes.
	//
	// GET /rest/api/{2-3}/permissionscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#get-all-permission-schemes
	Gets(ctx context.Context) (*model.PermissionSchemePageScheme, *model.ResponseScheme, error)

	// Get returns a permission scheme.
	//
	// GET /rest/api/{2-3}/permissionscheme/{schemeId}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#get-permission-scheme
	Get(ctx context.Context, permissionSchemeId int, expand []string) (*model.PermissionSchemeScheme, *model.ResponseScheme, error)

	// Delete deletes a permission scheme.
	//
	// DELETE /rest/api/{2-3}/permissionscheme/{schemeId}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#delete-permission-scheme
	Delete(ctx context.Context, permissionSchemeId int) (*model.ResponseScheme, error)

	// Create creates a new permission scheme.
	//
	// You can create a permission scheme with or without defining a set of permission grants.
	//
	// POST /rest/api/{2-3}/permissionscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#create-permission-scheme
	Create(ctx context.Context, payload *model.PermissionSchemeScheme) (*model.PermissionSchemeScheme, *model.ResponseScheme, error)

	// Update updates a permission scheme.
	// Below are some important things to note when using this resource:
	//
	// 1. If a permissions list is present in the request, then it is set in the permission scheme, overwriting all existing grants.
	//
	// 2. If you want to update only the name and description, then do not send a permissions list in the request.
	//
	// 3. Sending an empty list will remove all permission grants from the permission scheme.
	//
	// PUT /rest/api/{2-3}/permissionscheme/{schemeId}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/permissions/scheme#update-permission-scheme
	Update(ctx context.Context, permissionSchemeId int, payload *model.PermissionSchemeScheme) (*model.PermissionSchemeScheme, *model.ResponseScheme, error)
}
