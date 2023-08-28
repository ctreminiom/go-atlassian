package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func NewAuditRecordService(client service.Connector, version string) (*AuditRecordService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &AuditRecordService{
		internalClient: &internalAuditRecordImpl{c: client, version: version},
	}, nil
}

type AuditRecordService struct {
	internalClient jira.AuditRecordConnector
}

// Get allows you to retrieve the audit records for specific activities that have occurred within Jira.
//
// GET /rest/api/{2-3}/auditing/record
//
// https://docs.go-atlassian.io/jira-software-cloud/audit-records#get-audit-records
func (a *AuditRecordService) Get(ctx context.Context, options *model.AuditRecordGetOptions, offSet, limit int) (*model.AuditRecordPageScheme, *model.ResponseScheme, error) {
	return a.internalClient.Get(ctx, options, offSet, limit)
}

type internalAuditRecordImpl struct {
	c       service.Connector
	version string
}

func (i *internalAuditRecordImpl) Get(ctx context.Context, options *model.AuditRecordGetOptions, offSet, limit int) (*model.AuditRecordPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offSet))
	params.Add("limit", strconv.Itoa(limit))

	if options != nil {

		if options.Filter != "" {
			params.Add("", options.Filter)
		}

		if !options.To.IsZero() {
			params.Add("to", options.To.Format("2006-01-02"))
		}

		if !options.From.IsZero() {
			params.Add("from", options.From.Format("2006-01-02"))
		}

	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/auditing/record", i.version))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	records := new(model.AuditRecordPageScheme)
	response, err := i.c.Call(request, records)
	if err != nil {
		return nil, nil, err
	}

	return records, response, nil
}
