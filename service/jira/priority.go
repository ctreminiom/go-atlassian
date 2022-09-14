package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
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
	// GET /rest/api/{2-3}/priority/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/priorities#get-priority
	Get(ctx context.Context, priorityId string) (*model.PriorityScheme, *model.ResponseScheme, error)
}
