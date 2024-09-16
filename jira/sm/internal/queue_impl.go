package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/sm"
	"net/http"
	"net/url"
	"strconv"
)

// NewQueueService creates a new instance of QueueService.
// It takes a service.Connector and a version string as input and returns a pointer to QueueService.
func NewQueueService(client service.Connector, version string) *QueueService {

	return &QueueService{
		internalClient: &internalQueueServiceImpl{c: client, version: version},
	}
}

// QueueService provides methods to interact with queue operations in Jira Service Management.
type QueueService struct {
	// internalClient is the connector interface for queue operations.
	internalClient sm.QueueConnector
}

// Gets returns the queues in a service desk
//
// GET /rest/servicedeskapi/servicedesk/{serviceDeskId}/queue
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk/queue#get-queues
func (q *QueueService) Gets(ctx context.Context, serviceDeskID int, includeCount bool, start, limit int) (*model.ServiceDeskQueuePageScheme, *model.ResponseScheme, error) {
	return q.internalClient.Gets(ctx, serviceDeskID, includeCount, start, limit)
}

// Get returns a specific queues in a service desk.
//
// GET /rest/servicedeskapi/servicedesk/{serviceDeskId}/queue/{queueId}
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk/queue#get-queue
func (q *QueueService) Get(ctx context.Context, serviceDeskID, queueID int, includeCount bool) (*model.ServiceDeskQueueScheme, *model.ResponseScheme, error) {
	return q.internalClient.Get(ctx, serviceDeskID, queueID, includeCount)
}

// Issues returns the customer requests in a queue
//
// GET /rest/servicedeskapi/servicedesk/{serviceDeskId}/queue/{queueId}/issue
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk/queue#get-issues-in-queue
func (q *QueueService) Issues(ctx context.Context, serviceDeskID, queueID, start, limit int) (*model.ServiceDeskIssueQueueScheme, *model.ResponseScheme, error) {
	return q.internalClient.Issues(ctx, serviceDeskID, queueID, start, limit)
}

type internalQueueServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalQueueServiceImpl) Gets(ctx context.Context, serviceDeskID int, includeCount bool, start, limit int) (*model.ServiceDeskQueuePageScheme, *model.ResponseScheme, error) {

	if serviceDeskID == 0 {
		return nil, nil, model.ErrNoServiceDeskID
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))
	params.Add("includeCount", fmt.Sprintf("%v", includeCount))

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/queue?%v", serviceDeskID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ServiceDeskQueuePageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalQueueServiceImpl) Get(ctx context.Context, serviceDeskID, queueID int, includeCount bool) (*model.ServiceDeskQueueScheme, *model.ResponseScheme, error) {

	if serviceDeskID == 0 {
		return nil, nil, model.ErrNoServiceDeskID
	}

	if queueID == 0 {
		return nil, nil, model.ErrNoQueueID
	}

	params := url.Values{}
	params.Add("includeCount", fmt.Sprintf("%v", includeCount))

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/queue/%v?%v", serviceDeskID, queueID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	queue := new(model.ServiceDeskQueueScheme)
	res, err := i.c.Call(req, queue)
	if err != nil {
		return nil, res, err
	}

	return queue, res, nil
}

func (i *internalQueueServiceImpl) Issues(ctx context.Context, serviceDeskID, queueID, start, limit int) (*model.ServiceDeskIssueQueueScheme, *model.ResponseScheme, error) {

	if serviceDeskID == 0 {
		return nil, nil, model.ErrNoServiceDeskID
	}

	if queueID == 0 {
		return nil, nil, model.ErrNoQueueID
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/queue/%v/issue?%v", serviceDeskID, queueID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	issues := new(model.ServiceDeskIssueQueueScheme)
	res, err := i.c.Call(req, issues)
	if err != nil {
		return nil, res, err
	}

	return issues, res, nil
}
