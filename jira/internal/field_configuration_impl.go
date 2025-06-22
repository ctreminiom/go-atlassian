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

// NewIssueFieldConfigurationService creates a new instance of IssueFieldConfigService.
// It takes a service.Connector, a version string, an IssueFieldConfigItemService, and an IssueFieldConfigSchemeService as input.
// Returns a pointer to IssueFieldConfigService and an error if the version is not provided.
func NewIssueFieldConfigurationService(client service.Connector, version string, item *IssueFieldConfigItemService,
	scheme *IssueFieldConfigSchemeService) (*IssueFieldConfigService, error) {

	if version == "" {
		return nil, fmt.Errorf("jira: %w", model.ErrNoVersionProvided)
	}

	return &IssueFieldConfigService{
		internalClient: &internalIssueFieldConfigServiceImpl{c: client, version: version},
		Item:           item,
		Scheme:         scheme,
	}, nil
}

// IssueFieldConfigService provides methods to manage field configurations in Jira Service Management.
type IssueFieldConfigService struct {
	// internalClient is the connector interface for field configuration operations.
	internalClient jira.FieldConfigConnector
	// Item is the service for managing field configuration items.
	Item *IssueFieldConfigItemService
	// Scheme is the service for managing field configuration schemes.
	Scheme *IssueFieldConfigSchemeService
}

// Gets Returns a paginated list of all field configurations.
//
// GET /rest/api/{2-3}/fieldconfiguration
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-all-field-configurations
func (i *IssueFieldConfigService) Gets(ctx context.Context, ids []int, isDefault bool, startAt, maxResults int) (*model.FieldConfigurationPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldConfigService).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_field_configurations"),
		attribute.Bool("jira.field_config.is_default", isDefault),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
		attribute.Int("jira.field_config.ids_count", len(ids)),
	)

	result, response, err := i.internalClient.Gets(ctx, ids, isDefault, startAt, maxResults)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Create creates a field configuration. The field configuration is created with the same field properties as the
// default configuration, with all the fields being optional.
//
// This operation can only create configurations for use in company-managed (classic) projects.
//
// POST /rest/api/{2-3}/fieldconfiguration
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#create-field-configuration
func (i *IssueFieldConfigService) Create(ctx context.Context, name, description string) (*model.FieldConfigurationScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldConfigService).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "create_field_configuration"),
		attribute.String("jira.field_config.name", name),
	)

	result, response, err := i.internalClient.Create(ctx, name, description)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Update updates a field configuration. The name and the description provided in the request override the existing values.
//
// This operation can only update configurations used in company-managed (classic) projects.
//
// PUT /rest/api/{2-3}/fieldconfiguration/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#update-field-configuration
func (i *IssueFieldConfigService) Update(ctx context.Context, id int, name, description string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldConfigService).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update_field_configuration"),
		attribute.Int("jira.field_config.id", id),
		attribute.String("jira.field_config.name", name),
	)

	response, err := i.internalClient.Update(ctx, id, name, description)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

// Delete deletes a field configuration.
//
// This operation can only delete configurations used in company-managed (classic) projects.
//
// DELETE /rest/api/{2-3}/fieldconfiguration/{id}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#delete-field-configuration
func (i *IssueFieldConfigService) Delete(ctx context.Context, id int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldConfigService).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete_field_configuration"),
		attribute.Int("jira.field_config.id", id),
	)

	response, err := i.internalClient.Delete(ctx, id)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

type internalIssueFieldConfigServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalIssueFieldConfigServiceImpl) Gets(ctx context.Context, ids []int, isDefault bool, startAt, maxResults int) (*model.FieldConfigurationPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldConfigServiceImpl).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("isDefault", fmt.Sprintf("%v", isDefault))

	for _, id := range ids {
		params.Add("id", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	page := new(model.FieldConfigurationPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return page, response, nil
}

func (i *internalIssueFieldConfigServiceImpl) Create(ctx context.Context, name, description string) (*model.FieldConfigurationScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldConfigServiceImpl).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if name == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldConfigurationName)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration", i.version)

	payload := map[string]interface{}{"name": name}

	if description != "" {
		payload["description"] = description
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	issueConfig := new(model.FieldConfigurationScheme)
	response, err := i.c.Call(request, issueConfig)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return issueConfig, response, nil
}

func (i *internalIssueFieldConfigServiceImpl) Update(ctx context.Context, id int, name, description string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldConfigServiceImpl).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if id == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldConfigurationID)
		recordError(span, err)
		return nil, err
	}

	if name == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldConfigurationName)
		recordError(span, err)
		return nil, err
	}

	payload := map[string]interface{}{"name": name}

	if description != "" {
		payload["description"] = description
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration/%v", i.version, id)

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

func (i *internalIssueFieldConfigServiceImpl) Delete(ctx context.Context, id int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldConfigServiceImpl).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if id == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldConfigurationID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/fieldconfiguration/%v", i.version, id)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
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
