package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type MySelfConnector interface {

	// Details returns details for the current user.
	//
	// GET /rest/api/{2-3}/myself
	//
	// https://docs.go-atlassian.io/jira-software-cloud/myself#get-current-user
	Details(ctx context.Context, expand []string) (*model.UserScheme, *model.ResponseScheme, error)

	// Get returns the values of the user's preferences.
	//
	// GET /rest/api/{2-3}/mypreferences
	//
	// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-myself/#api-rest-api-3-mypreferences-get
	Get(ctx context.Context, key string) (map[string]interface{}, *model.ResponseScheme, error)

	// Set sets the value of the user's preference.
	//
	// PUT /rest/api/{2-3}/mypreferences
	//
	// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-myself/#api-rest-api-3-mypreferences-put
	Set(ctx context.Context, key string, value string) (map[string]interface{}, *model.ResponseScheme, error)

	// Delete deletes the user's preference.
	//
	// DELETE /rest/api/{2-3}/mypreferences
	//
	// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-myself/#api-rest-api-3-mypreferences-delete
	Delete(ctx context.Context, key string) (*model.ResponseScheme, error)
}
