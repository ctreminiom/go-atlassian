package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// GroupUserPickerConnector is an interface that defines the methods available from GroupUserPickerConnector API.
type GroupUserPickerConnector interface {

	// Find returns a list of users and groups matching a string.
	//
	// GET /rest/api/{2-3}/groupuserpicker
	//
	// https://docs.go-atlassian.io/jira-software-cloud/groupuserpicker#find-users-and-groups
	Find(ctx context.Context, options *model.GroupUserPickerFindOptionScheme) (*model.GroupUserPickerFindScheme, *model.ResponseScheme, error)
}
