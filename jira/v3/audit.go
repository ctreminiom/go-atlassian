package v3

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type AuditService struct{ client *Client }

type AuditRecordPageScheme struct {
	Offset  int                  `json:"offset,omitempty"`
	Limit   int                  `json:"limit,omitempty"`
	Total   int                  `json:"total,omitempty"`
	Records []*AuditRecordScheme `json:"records,omitempty"`
}

type AuditRecordScheme struct {
	ID              int                                `json:"id,omitempty"`
	Summary         string                             `json:"summary,omitempty"`
	RemoteAddress   string                             `json:"remoteAddress,omitempty"`
	AuthorKey       string                             `json:"authorKey,omitempty"`
	Created         string                             `json:"created,omitempty"`
	Category        string                             `json:"category,omitempty"`
	EventSource     string                             `json:"eventSource,omitempty"`
	Description     string                             `json:"description,omitempty"`
	ObjectItem      *AuditRecordObjectItemScheme       `json:"objectItem,omitempty"`
	ChangedValues   []*AuditRecordChangedValueScheme   `json:"changedValues,omitempty"`
	AssociatedItems []*AuditRecordAssociatedItemScheme `json:"associatedItems,omitempty"`
}

type AuditRecordObjectItemScheme struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	TypeName   string `json:"typeName,omitempty"`
	ParentID   string `json:"parentId,omitempty"`
	ParentName string `json:"parentName,omitempty"`
}

type AuditRecordChangedValueScheme struct {
	FieldName   string `json:"fieldName,omitempty"`
	ChangedFrom string `json:"changedFrom,omitempty"`
	ChangedTo   string `json:"changedTo,omitempty"`
}

type AuditRecordAssociatedItemScheme struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	TypeName   string `json:"typeName,omitempty"`
	ParentID   string `json:"parentId,omitempty"`
	ParentName string `json:"parentName,omitempty"`
}

type AuditRecordGetOptions struct {
	Filter   string
	From, To time.Time
}

// Get returns a list of audit records. The list can be filtered to include items:
// Docs: https://docs.go-atlassian.io/jira-software-cloud/audit-records#get-audit-records
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-audit-records/#api-rest-api-3-auditing-record-get
func (a *AuditService) Get(ctx context.Context, options *AuditRecordGetOptions, offset, limit int) (result *AuditRecordPageScheme, response *ResponseScheme, err error) {

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
