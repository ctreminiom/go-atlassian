package v2

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"net/http"
)

type TaskService struct{ client *Client }

// Get returns the status of a long-running asynchronous task.
// When a task has finished, this operation returns the JSON blob applicable to the task.
// See the documentation of the operation that created the task for details.
// Task details are not permanently retained.
// As of September 2019, details are retained for 14 days although this period may change without notice.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/tasks#get-task
func (t *TaskService) Get(ctx context.Context, taskID string) (result *models.TaskScheme, response *ResponseScheme, err error) {

	if len(taskID) == 0 {
		return nil, nil, notTaskIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/task/%v", taskID)

	request, err := t.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = t.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Cancel cancels a task.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/tasks#cancel-task
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-tasks/#api-rest-api-2-task-taskid-cancel-post
func (t *TaskService) Cancel(ctx context.Context, taskID string) (response *ResponseScheme, err error) {

	if len(taskID) == 0 {
		return nil, notTaskIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/task/%v/cancel", taskID)

	request, err := t.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = t.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

var (
	notTaskIDError = fmt.Errorf("error, please provide a valid taskID value")
)
