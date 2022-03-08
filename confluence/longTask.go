package confluence

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type LongTaskService struct{ client *Client }

// Gets returns information about all active long-running tasks (e.g. space export),
// such as how long each task has been running and the percentage of each task that has completed.
func (l *LongTaskService) Gets(ctx context.Context, start, limit int) (result *models.LongTaskPageScheme,
	response *ResponseScheme, err error) {

	query := url.Values{}
	query.Add("start", strconv.Itoa(start))
	query.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("/wiki/rest/api/longtask?%v", query.Encode())

	request, err := l.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = l.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Get returns information about an active long-running task (e.g. space export), such as how long it has been running
//and the percentage of the task that has completed.
func (l *LongTaskService) Get(ctx context.Context, taskID string) (result *models.LongTaskScheme, response *ResponseScheme,
	err error) {

	var endpoint = fmt.Sprintf("/wiki/rest/api/longtask/%v", taskID)

	request, err := l.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	request.Header.Set("Accept", "application/json")

	response, err = l.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}
