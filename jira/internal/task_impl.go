package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

func NewTaskService(client service.Connector, version string) (*TaskService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &TaskService{
		internalClient: &internalTaskServiceImpl{c: client, version: version},
	}, nil
}

type TaskService struct {
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
// GET /rest/api/{2-3}/task/{taskId}
//
// https://docs.go-atlassian.io/jira-software-cloud/tasks#get-task
func (t *TaskService) Get(ctx context.Context, taskId string) (*model.TaskScheme, *model.ResponseScheme, error) {
	return t.internalClient.Get(ctx, taskId)
}

// Cancel cancels a task.
//
// POST /rest/api/{2-3}/task/{taskId}/cancel
//
// https://docs.go-atlassian.io/jira-software-cloud/tasks#cancel-task
func (t *TaskService) Cancel(ctx context.Context, taskId string) (*model.ResponseScheme, error) {
	return t.internalClient.Cancel(ctx, taskId)
}

type internalTaskServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalTaskServiceImpl) Get(ctx context.Context, taskId string) (*model.TaskScheme, *model.ResponseScheme, error) {

	if taskId == "" {
		return nil, nil, model.ErrNoTaskIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/task/%v", i.version, taskId)

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

func (i *internalTaskServiceImpl) Cancel(ctx context.Context, taskId string) (*model.ResponseScheme, error) {

	if taskId == "" {
		return nil, model.ErrNoTaskIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/task/%v/cancel", i.version, taskId)
	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
