package sm

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type RequestSLAService struct{ client *Client }

// Gets  returns all the SLA records on a customer request.
// A customer request can have zero or more SLAs. Each SLA can have recordings for zero or more "completed cycles" and zero or 1 "ongoing cycle".
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/sla#get-sla-information
func (r *RequestSLAService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (result *model.RequestSLAPageScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/sla?%v", issueKeyOrID, params.Encode())

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns the details for an SLA on a customer request.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/sla#get-sla-information-by-id
func (r *RequestSLAService) Get(ctx context.Context, issueKeyOrID string, slaMetricID int) (result *model.RequestSLAScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/request/%v/sla/%v", issueKeyOrID, slaMetricID)

	request, err := r.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = r.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}
