package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewIssueFieldConfigurationItemService creates a new instance of IssueFieldConfigItemService.
// It takes a service.Connector and a version string as input.
// Returns a pointer to IssueFieldConfigItemService and an error if the version is not provided.
func NewIssueFieldConfigurationItemService(client service.Connector, version string) (*IssueFieldConfigItemService, error) {

	if version == "" {
		return nil, fmt.Errorf("jira: %w", model.ErrNoVersionProvided)
	}

	return &IssueFieldConfigItemService{
		internalClient: &internalIssueFieldConfigItemServiceImpl{c: client, version: version},
	}, nil
}

// IssueFieldConfigItemService provides methods to manage field configuration items in Jira Service Management.
type IssueFieldConfigItemService struct {
	// internalClient is the connector interface for field configuration item operations.
	internalClient jira.FieldConfigItemConnector
}

// Gets Returns a paginated list of all fields for a configuration.
//
// GET /rest/api/{2-3}/fieldconfiguration/{id}/fields
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/items#get-field-configuration-items
func (i *IssueFieldConfigItemService) Gets(ctx context.Context, id, startAt, maxResults int) (*model.FieldConfigurationItemPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldConfigItemService).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_field_configuration_items"),
		attribute.Int("jira.field_config.id", id),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
	)

	result, response, err := i.internalClient.Gets(ctx, id, startAt, maxResults)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Update updates fields in a field configuration. The properties of the field configuration fields provided
// override the existing values.
//
// 1. This operation can only update field configurations used in company-managed (classic) projects.
//
// PUT /rest/api/{2-3}/fieldconfiguration/{id}/fields
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/items#update-field-configuration-items
func (i *IssueFieldConfigItemService) Update(ctx context.Context, id int, payload *model.UpdateFieldConfigurationItemPayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldConfigItemService).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update_field_configuration_items"),
		attribute.Int("jira.field_config.id", id),
	)

	response, err := i.internalClient.Update(ctx, id, payload)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

type internalIssueFieldConfigItemServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalIssueFieldConfigItemServiceImpl) Gets(ctx context.Context, id, startAt, maxResults int) (*model.FieldConfigurationItemPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldConfigItemServiceImpl).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if id == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldConfigurationID)
		recordError(span, err)
		return nil, nil, err
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration/%v/fields?%v", i.version, id, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	page := new(model.FieldConfigurationItemPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalIssueFieldConfigItemServiceImpl) Update(ctx context.Context, id int, payload *model.UpdateFieldConfigurationItemPayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldConfigItemServiceImpl).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if id == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldConfigurationID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration/%v/fields", i.version, id)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, err
	}

	response, err := i.c.Call(request, nil)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}
