package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/sm"
)

// NewServiceLevelAgreementService creates a new instance of ServiceLevelAgreementService.
// It takes a service.Connector and a version string as input and returns a pointer to ServiceLevelAgreementService.
func NewServiceLevelAgreementService(client service.Connector, version string) *ServiceLevelAgreementService {
	return &ServiceLevelAgreementService{
		internalClient: &internalServiceLevelAgreementImpl{c: client, version: version},
	}
}

// ServiceLevelAgreementService provides methods to interact with SLA operations in Jira Service Management.
type ServiceLevelAgreementService struct {
	// internalClient is the connector interface for SLA operations.
	internalClient sm.ServiceLevelAgreementConnector
}

// Gets  returns all the SLA records on a customer request.
//
// A customer request can have zero or more SLAs. Each SLA can have recordings for zero or more "completed cycles" and zero or 1 "ongoing cycle".
//
// GET /rest/servicedeskapi/request/{issueKeyOrID}/sla
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/sla#get-sla-information
func (s *ServiceLevelAgreementService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.RequestSLAPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ServiceLevelAgreementService).Gets")
	defer span.End()

	return s.internalClient.Gets(ctx, issueKeyOrID, start, limit)
}

// Get returns the details for an SLA on a customer request.
//
// GET /rest/servicedeskapi/request/{issueKeyOrID}/sla/{slaMetricId}
//
// https://docs.go-atlassian.io/jira-service-management-cloud/request/sla#get-sla-information-by-id
func (s *ServiceLevelAgreementService) Get(ctx context.Context, issueKeyOrID string, metricID int) (*model.RequestSLAScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*ServiceLevelAgreementService).Get")
	defer span.End()

	return s.internalClient.Get(ctx, issueKeyOrID, metricID)
}

type internalServiceLevelAgreementImpl struct {
	c       service.Connector
	version string
}

func (i *internalServiceLevelAgreementImpl) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (*model.RequestSLAPageScheme, *model.ResponseScheme, error) {

	ctx, span := tracer().Start(ctx, "(*internalServiceLevelAgreementImpl).Gets")
	defer span.End()

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/sla?%v", issueKeyOrID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.RequestSLAPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalServiceLevelAgreementImpl) Get(ctx context.Context, issueKeyOrID string, metricID int) (*model.RequestSLAScheme, *model.ResponseScheme, error) {

	ctx, span := tracer().Start(ctx, "(*internalServiceLevelAgreementImpl).Get")
	defer span.End()

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	if metricID == 0 {
		return nil, nil, model.ErrNoSLAMetricID
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/request/%v/sla/%v", issueKeyOrID, metricID)

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	sla := new(model.RequestSLAScheme)
	res, err := i.c.Call(req, sla)
	if err != nil {
		return nil, res, err
	}

	return sla, res, nil
}
