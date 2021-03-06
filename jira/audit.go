package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type AuditService struct{ client *Client }

type AuditRecordScheme struct {
	Offset  int `json:"offset,omitempty"`
	Limit   int `json:"limit,omitempty"`
	Total   int `json:"total,omitempty"`
	Records []struct {
		ID            int    `json:"id,omitempty"`
		Summary       string `json:"summary,omitempty"`
		RemoteAddress string `json:"remoteAddress,omitempty"`
		AuthorKey     string `json:"authorKey,omitempty"`
		Created       string `json:"created,omitempty"`
		Category      string `json:"category,omitempty"`
		EventSource   string `json:"eventSource,omitempty"`
		Description   string `json:"description,omitempty"`
		ObjectItem    struct {
			ID         string `json:"id,omitempty"`
			Name       string `json:"name,omitempty"`
			TypeName   string `json:"typeName,omitempty"`
			ParentID   string `json:"parentId,omitempty"`
			ParentName string `json:"parentName,omitempty"`
		} `json:"objectItem,omitempty"`
		ChangedValues []struct {
			FieldName   string `json:"fieldName,omitempty"`
			ChangedFrom string `json:"changedFrom,omitempty"`
			ChangedTo   string `json:"changedTo,omitempty"`
		} `json:"changedValues,omitempty"`
		AssociatedItems []struct {
			ID         string `json:"id,omitempty"`
			Name       string `json:"name,omitempty"`
			TypeName   string `json:"typeName,omitempty"`
			ParentID   string `json:"parentId,omitempty"`
			ParentName string `json:"parentName,omitempty"`
		} `json:"associatedItems,omitempty"`
	} `json:"records,omitempty"`
}

type AuditRecordGetOptions struct {
	Filter string
	From   string
	To     string
}

// Returns a list of audit records. The list can be filtered to include items:
// Docs: https://docs.go-atlassian.io/jira-software-cloud/audit-records#get-audit-records
func (a *AuditService) Get(ctx context.Context, options *AuditRecordGetOptions, offset, limit int) (result *AuditRecordScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	if options != nil {

		if len(options.Filter) != 0 {
			params.Add("filter", options.Filter)
		}

		if len(options.From) != 0 {
			params.Add("from", options.From)
		}

		if len(options.To) != 0 {
			params.Add("to", options.To)
		}
	}

	var endpoint = fmt.Sprintf("rest/api/3/auditing/record?%s", params.Encode())
	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = a.client.Do(request)
	if err != nil {
		return
	}

	result = new(AuditRecordScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
