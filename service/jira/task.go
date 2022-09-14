package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type TaskConnector interface {

	// Get returns the status of a long-running asynchronous task.
	//
	// When a task has finished, this operation returns the JSON blob applicable to the task.
	//
	// See the documentation of the operation that created the task for details.
	//
	// Task details are not permanently retained.
	//
	// GET /rest/api/{2-3}/task/{taskId}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/tasks#get-task
	Get(ctx context.Context, taskId string) (*model.TaskScheme, *model.ResponseScheme, error)

	// Cancel cancels a task.
	//
	// POST /rest/api/{2-3}/task/{taskId}/cancel
	//
	// https://docs.go-atlassian.io/jira-software-cloud/tasks#cancel-task
	Cancel(ctx context.Context, taskId string) (*model.ResponseScheme, error)
}
