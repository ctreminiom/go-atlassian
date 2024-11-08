package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewTaskService creates a new instance of TaskService.
func NewTaskService(client service.Connector, version string) (*TaskService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &TaskService{
		internalClient: &internalTaskServiceImpl{c: client, version: version},
	}, nil
}

// TaskService provides methods to manage tasks in Jira Service Management.
type TaskService struct {
	// internalClient is the connector interface for task operations.
	internalClient jira.TaskConnector
}

// Get returns the status of a long-running asynchronous task.
//
// When a task has finished, this operation returns the JSON blob applicable to the task.
//
// See the documentation of the operation that created the task for details.
//
// Task details are not permanently retained.
//
// GET /rest/api/{2-3}/task/{taskID}
//
// https://docs.go-atlassian.io/jira-software-cloud/tasks#get-task
func (t *TaskService) Get(ctx context.Context, taskID string) (*model.TaskScheme, *model.ResponseScheme, error) {
	return t.internalClient.Get(ctx, taskID)
}

// Cancel cancels a task.
//
// POST /rest/api/{2-3}/task/{taskID}/cancel
//
// https://docs.go-atlassian.io/jira-software-cloud/tasks#cancel-task
func (t *TaskService) Cancel(ctx context.Context, taskID string) (*model.ResponseScheme, error) {
	return t.internalClient.Cancel(ctx, taskID)
}

type internalTaskServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalTaskServiceImpl) Get(ctx context.Context, taskID string) (*model.TaskScheme, *model.ResponseScheme, error) {

	if taskID == "" {
		return nil, nil, model.ErrNoTaskID
	}

	endpoint := fmt.Sprintf("rest/api/%v/task/%v", i.version, taskID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	task := new(model.TaskScheme)
	response, err := i.c.Call(request, task)
	if err != nil {
		return nil, response, err
	}

	return task, response, nil
}

func (i *internalTaskServiceImpl) Cancel(ctx context.Context, taskID string) (*model.ResponseScheme, error) {

	if taskID == "" {
		return nil, model.ErrNoTaskID
	}

	endpoint := fmt.Sprintf("rest/api/%v/task/%v/cancel", i.version, taskID)
	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
