package v3

import (
	"context"
	"fmt"
	"net/http"
)

type TaskService struct{ client *Client }

type TaskScheme struct {
	Self           string `json:"self"`
	ID             string `json:"id"`
	Description    string `json:"description"`
	Status         string `json:"status"`
	Result         string `json:"result"`
	SubmittedBy    int    `json:"submittedBy"`
	Progress       int    `json:"progress"`
	ElapsedRuntime int    `json:"elapsedRuntime"`
	Submitted      int64  `json:"submitted"`
	Started        int64  `json:"started"`
	Finished       int64  `json:"finished"`
	LastUpdate     int64  `json:"lastUpdate"`
}

// Get returns the status of a long-running asynchronous task.
// When a task has finished, this operation returns the JSON blob applicable to the task.
// See the documentation of the operation that created the task for details.
// Task details are not permanently retained.
// As of September 2019, details are retained for 14 days although this period may change without notice.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/tasks#get-task
func (t *TaskService) Get(ctx context.Context, taskID string) (result *TaskScheme, response *ResponseScheme, err error) {

	if len(taskID) == 0 {
		return nil, nil, notTaskIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/task/%v", taskID)

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
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-tasks/#api-rest-api-3-task-taskid-cancel-post
func (t *TaskService) Cancel(ctx context.Context, taskID string) (response *ResponseScheme, err error) {

	if len(taskID) == 0 {
		return nil, notTaskIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/task/%v/cancel", taskID)

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
