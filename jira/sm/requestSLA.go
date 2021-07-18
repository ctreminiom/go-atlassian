package sm

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type RequestSLAService struct{ client *Client }

// Gets  returns all the SLA records on a customer request.
// A customer request can have zero or more SLAs. Each SLA can have recordings for zero or more "completed cycles" and zero or 1 "ongoing cycle".
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/request/sla#get-sla-information
func (r *RequestSLAService) Gets(ctx context.Context, issueKeyOrID string, start, limit int) (result *RequestSLAPageScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, notIssueError
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
func (r *RequestSLAService) Get(ctx context.Context, issueKeyOrID string, slaMetricID int) (result *RequestSLAScheme, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, notIssueError
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

type RequestSLAPageScheme struct {
	Size       int                       `json:"size,omitempty"`
	Start      int                       `json:"start,omitempty"`
	Limit      int                       `json:"limit,omitempty"`
	IsLastPage bool                      `json:"isLastPage,omitempty"`
	Values     []*RequestSLAScheme       `json:"values,omitempty"`
	Expands    []string                  `json:"_expands,omitempty"`
	Links      *RequestSLAPageLinkScheme `json:"_links,omitempty"`
}

type RequestSLAPageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type RequestSLAScheme struct {
	ID           string                        `json:"id,omitempty"`
	Name         string                        `json:"name,omitempty"`
	OngoingCycle *RequestSLAOngoingCycleScheme `json:"ongoingCycle,omitempty"`
	Links        *RequestSLALinkScheme         `json:"_links,omitempty"`
}

type RequestSLAOngoingCycleScheme struct {
	Breached            bool `json:"breached,omitempty"`
	Paused              bool `json:"paused,omitempty"`
	WithinCalendarHours bool `json:"withinCalendarHours,omitempty"`
}

type RequestSLALinkScheme struct {
	Self string `json:"self,omitempty"`
}
