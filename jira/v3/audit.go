package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type AuditService struct{ client *Client }

type AuditRecordGetOptions struct {
	Filter   string
	From, To time.Time
}

// Get returns a list of audit records. The list can be filtered to include items:
// Docs: https://docs.go-atlassian.io/jira-software-cloud/audit-records#get-audit-records
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-audit-records/#api-rest-api-3-auditing-record-get
func (a *AuditService) Get(ctx context.Context, options *AuditRecordGetOptions, offset, limit int) (result *models.AuditRecordPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	if options != nil {

		if len(options.Filter) != 0 {
			params.Add("filter", options.Filter)
		}

		if !options.From.IsZero() {
			params.Add("from", options.From.Format(DateFormatJira))
		}

		if !options.To.IsZero() {
			params.Add("to", options.To.Format(DateFormatJira))
		}
	}

	var endpoint = fmt.Sprintf("rest/api/3/auditing/record?%s", params.Encode())
	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = a.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
