package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type AppRoleConnector interface {

	// Gets returns all application roles.
	//
	// In Jira, application roles are managed using the Application access configuration page.
	//
	// GET /rest/api/{2-3}/applicationrole
	//
	// https://docs.go-atlassian.io/jira-software-cloud/application-roles#get-all-application-roles
	Gets(ctx context.Context) ([]*model.ApplicationRoleScheme, *model.ResponseScheme, error)

	// Get returns an application role.
	//
	// GET /rest/api/{2-3}/applicationrole/{key}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/application-roles#get-application-role
	Get(ctx context.Context, key string) (*model.ApplicationRoleScheme, *model.ResponseScheme, error)
}
