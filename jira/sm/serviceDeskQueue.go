package sm

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ServiceDeskQueueService struct{ client *Client }

// Gets returns the queues in a service desk
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk/queue#get-queues
func (s *ServiceDeskQueueService) Gets(ctx context.Context, serviceDeskID int, includeCount bool, start, limit int) (
	result *model.ServiceDeskQueuePageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if includeCount {
		params.Add("includeCount", "true")
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/queue?%v", serviceDeskID, params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns a specific queues in a service desk.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk/queue#get-queue
func (s *ServiceDeskQueueService) Get(ctx context.Context, serviceDeskID, queueID int, includeCount bool) (
	result *model.ServiceDeskQueueScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	if includeCount {
		params.Add("includeCount", "true")
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/queue/%v", serviceDeskID, queueID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Issues returns the customer requests in a queue
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk/queue#get-issues-in-queue
func (s *ServiceDeskQueueService) Issues(ctx context.Context, serviceDeskID, queueID, start, limit int) (
	result *model.ServiceDeskIssueQueueScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/queue/%v/issue?%v", serviceDeskID, queueID, params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}
