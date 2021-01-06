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

func (a *AuditService) Get(ctx context.Context, offset, limit int, filter, from, to string) (records *AuditRecordScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offset))
	params.Add("limit", strconv.Itoa(limit))

	validateURLParam(&params, "filter", filter)
	validateURLParam(&params, "from", from)
	validateURLParam(&params, "to", to)

	var endpoint = fmt.Sprintf("rest/api/3/auditing/record?%s", params.Encode())
	request, err := a.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = a.client.Do(request)
	if err != nil {
		return
	}

	result := new(AuditRecordScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
