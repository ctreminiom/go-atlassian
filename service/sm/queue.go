package sm

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type QueueConnector interface {

	// Gets returns the queues in a service desk
	//
	// GET /rest/servicedeskapi/servicedesk/{serviceDeskID}/queue
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk/queue#get-queues
	Gets(ctx context.Context, serviceDeskID int, includeCount bool, start, limit int) (*model.ServiceDeskQueuePageScheme, *model.ResponseScheme, error)

	// Get returns a specific queues in a service desk.
	//
	// GET /rest/servicedeskapi/servicedesk/{serviceDeskID}/queue/{queueID}
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk/queue#get-queue
	Get(ctx context.Context, serviceDeskID, queueID int, includeCount bool) (*model.ServiceDeskQueueScheme, *model.ResponseScheme, error)

	// Issues returns the customer requests in a queue
	//
	// GET /rest/servicedeskapi/servicedesk/{serviceDeskID}/queue/{queueID}/issue
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/service-desk/queue#get-issues-in-queue
	Issues(ctx context.Context, serviceDeskID, queueID, start, limit int) (*model.ServiceDeskIssueQueueScheme, *model.ResponseScheme, error)
}
