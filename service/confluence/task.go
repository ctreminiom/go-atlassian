package confluence

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type TaskConnector interface {

	// Gets returns information about all active long-running tasks (e.g. space export),
	//
	// such as how long each task has been running and the percentage of each task that has completed.
	//
	// GET /wiki/rest/api/longtask
	//
	// https://docs.go-atlassian.io/confluence-cloud/long-task#get-long-running-tasks
	Gets(ctx context.Context, start, limit int) (*model.LongTaskPageScheme, *model.ResponseScheme, error)

	// Get returns information about an active long-running task (e.g. space export), such as how long it has been running
	//
	// and the percentage of the task that has completed.
	//
	// GET /wiki/rest/api/longtask/{id}
	//
	// https://docs.go-atlassian.io/confluence-cloud/long-task#get-long-running-task
	Get(ctx context.Context, taskID string) (*model.LongTaskScheme, *model.ResponseScheme, error)
}
