package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewAuditRecordService creates a new instance of AuditRecordService.
// It takes a service.Connector and a version string as input.
// Returns a pointer to AuditRecordService and an error if the version is not provided.
func NewAuditRecordService(client service.Connector, version string) (*AuditRecordService, error) {

	if version == "" {
		return nil, fmt.Errorf("jira: %w", model.ErrNoVersionProvided)
	}

	return &AuditRecordService{
		internalClient: &internalAuditRecordImpl{c: client, version: version},
	}, nil
}

// AuditRecordService provides methods to interact with audit record operations in Jira Service Management.
type AuditRecordService struct {
	// internalClient is the connector interface for audit record operations.
	internalClient jira.AuditRecordConnector
}

// Get allows you to retrieve the audit records for specific activities that have occurred within Jira.
//
// GET /rest/api/{2-3}/auditing/record
//
// https://docs.go-atlassian.io/jira-software-cloud/audit-records#get-audit-records
func (a *AuditRecordService) Get(ctx context.Context, options *model.AuditRecordGetOptions, offSet, limit int) (*model.AuditRecordPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*AuditRecordService).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_audit_records"),
		attribute.Int("jira.audit.offset", offSet),
		attribute.Int("jira.audit.limit", limit),
	)

	result, response, err := a.internalClient.Get(ctx, options, offSet, limit)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

type internalAuditRecordImpl struct {
	c       service.Connector
	version string
}

func (i *internalAuditRecordImpl) Get(ctx context.Context, options *model.AuditRecordGetOptions, offSet, limit int) (*model.AuditRecordPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalAuditRecordImpl).Get", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_audit_records"),
		attribute.Int("jira.audit.offset", offSet),
		attribute.Int("jira.audit.limit", limit),
		attribute.String("api.version", i.version),
	)

	params := url.Values{}
	params.Add("offset", strconv.Itoa(offSet))
	params.Add("limit", strconv.Itoa(limit))

	if options != nil {

		if options.Filter != "" {
			params.Add("", options.Filter)
			addAttributes(span, attribute.String("jira.audit.filter", options.Filter))
		}

		if !options.To.IsZero() {
			params.Add("to", options.To.Format("2006-01-02"))
			addAttributes(span, attribute.String("jira.audit.to", options.To.Format("2006-01-02")))
		}

		if !options.From.IsZero() {
			params.Add("from", options.From.Format("2006-01-02"))
			addAttributes(span, attribute.String("jira.audit.from", options.From.Format("2006-01-02")))
		}

	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/auditing/record", i.version))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	records := new(model.AuditRecordPageScheme)
	response, err := i.c.Call(request, records)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	setOK(span)
	return records, response, nil
}
