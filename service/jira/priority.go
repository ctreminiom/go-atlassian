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
	Gets(ctx context.Context) ([]*model.PriorityScheme, *model.ResponseScheme, error)

	// Get returns an issue priority.
	//
	// GET /rest/api/{2-3}/priority/{priorityID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/priorities#get-priority
	Get(ctx context.Context, priorityID string) (*model.PriorityScheme, *model.ResponseScheme, error)
}
