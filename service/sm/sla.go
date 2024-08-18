package sm

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type ServiceLevelAgreementConnector interface {

	// Gets  returns all the SLA records on a customer request.
	//
	// A customer request can have zero or more SLAs. Each SLA can have recordings for zero or more "completed cycles" and zero or 1 "ongoing cycle".
	//
	// GET /rest/servicedeskapi/request/{issueKeyOrID}/sla
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/sla#get-sla-information
	Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.RequestSLAPageScheme, *model.ResponseScheme, error)

	// Get returns the details for an SLA on a customer request.
	//
	// GET /rest/servicedeskapi/request/{issueKeyOrID}/sla/{slaMetricId}
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/sla#get-sla-information-by-id
	Get(ctx context.Context, issueKeyOrID string, metricID int) (*model.RequestSLAScheme, *model.ResponseScheme, error)
}
