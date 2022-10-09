package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/confluence"
	"net/http"
	"net/url"
	"strconv"
)

func NewTaskService(client service.Client) *TaskService {

	return &TaskService{
		internalClient: &internalTaskImpl{c: client},
	}
}

type TaskService struct {
	internalClient confluence.TaskConnector
}

// Gets returns information about all active long-running tasks (e.g. space export),
//
// such as how long each task has been running and the percentage of each task that has completed.
//
// GET /wiki/rest/api/longtask
//
// https://docs.go-atlassian.io/confluence-cloud/long-task#get-long-running-tasks
func (t *TaskService) Gets(ctx context.Context, start, limit int) (*model.LongTaskPageScheme, *model.ResponseScheme, error) {
	return t.internalClient.Gets(ctx, start, limit)
}

// Get returns information about an active long-running task (e.g. space export), such as how long it has been running
//
// and the percentage of the task that has completed.
//
// GET /wiki/rest/api/longtask/{id}
//
// https://docs.go-atlassian.io/confluence-cloud/long-task#get-long-running-task
func (t *TaskService) Get(ctx context.Context, taskID string) (*model.LongTaskScheme, *model.ResponseScheme, error) {
	return t.internalClient.Get(ctx, taskID)
}

type internalTaskImpl struct {
	c service.Client
}

func (i *internalTaskImpl) Gets(ctx context.Context, start, limit int) (*model.LongTaskPageScheme, *model.ResponseScheme, error) {

	query := url.Values{}
	query.Add("start", strconv.Itoa(start))
	query.Add("limit", strconv.Itoa(limit))

	endpoint := fmt.Sprintf("wiki/rest/api/longtask?%v", query.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.LongTaskPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalTaskImpl) Get(ctx context.Context, taskID string) (*model.LongTaskScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("wiki/rest/api/longtask/%v", taskID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	task := new(model.LongTaskScheme)
	response, err := i.c.Call(request, task)
	if err != nil {
		return nil, response, err
	}

	return task, response, nil
}
