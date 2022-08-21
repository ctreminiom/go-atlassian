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
