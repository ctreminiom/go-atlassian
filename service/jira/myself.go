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
}
