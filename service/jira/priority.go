package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type PriorityConnector interface {

	// Gets returns the list of all issue priorities.
	//
	// GET /rest/api/{2-3}/priority
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/priorities#get-priorities
	// Deprecated: This endpoint is deprecated in the Jira API spec.
	// TODO: Cannot change without breaking API compatibility. Consider removing in next major version.
	Gets(ctx context.Context) ([]*model.PriorityScheme, *model.ResponseScheme, error)

	// Get returns an issue priority.
	//
	// GET /rest/api/{2-3}/priority/{priorityID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/priorities#get-priority
	// Deprecated: This endpoint is deprecated in the Jira API spec.
	// TODO Cannot change without breaking API compatibility. Consider removing in next major version.
	Get(ctx context.Context, priorityID string) (*model.PriorityScheme, *model.ResponseScheme, error)
}
