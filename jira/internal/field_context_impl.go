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

// NewIssueFieldContextService creates a new instance of IssueFieldContextService.
// It takes a service.Connector, a version string, and an IssueFieldContextOptionService as input.
// Returns a pointer to IssueFieldContextService and an error if the version is not provided.
func NewIssueFieldContextService(client service.Connector, version string, option *IssueFieldContextOptionService) (*IssueFieldContextService, error) {

	if version == "" {
		return nil, fmt.Errorf("jira: %w", model.ErrNoVersionProvided)
	}

	return &IssueFieldContextService{
		internalClient: &internalIssueFieldContextServiceImpl{c: client, version: version},
		Option:         option,
	}, nil
}

// IssueFieldContextService provides methods to manage field contexts in Jira Service Management.
type IssueFieldContextService struct {
	// internalClient is the connector interface for field context operations.
	internalClient jira.FieldContextConnector
	// Option is the service for managing field context options.
	Option *IssueFieldContextOptionService
}

// Gets returns a paginated list of contexts for a custom field. Contexts can be returned as follows:
//
// 1. By defining id only, all contexts from the list of IDs.
//
// 2. By defining isAnyIssueType
//
// 3. By defining isGlobalContext
//
// GET /rest/api/{2-3}/field/{fieldID}/context
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#get-custom-field-contexts
func (i *IssueFieldContextService) Gets(ctx context.Context, fieldID string, options *model.FieldContextOptionsScheme, startAt, maxResults int) (*model.CustomFieldContextPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldContextService).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_field_contexts"),
		attribute.String("jira.field.id", fieldID),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
	)

	result, response, err := i.internalClient.Gets(ctx, fieldID, options, startAt, maxResults)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Create creates a custom field context.
//
// 1. If projectIDs is empty, a global context is created. A global context is one that applies to all project.
//
// 2. If issueTypeIDs is empty, the context applies to all issue types.
//
// POST /rest/api/{2-3}/field/{fieldID}/context
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#create-custom-field-context
func (i *IssueFieldContextService) Create(ctx context.Context, fieldID string, payload *model.FieldContextPayloadScheme) (*model.FieldContextScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldContextService).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "create_field_context"),
		attribute.String("jira.field.id", fieldID),
	)

	result, response, err := i.internalClient.Create(ctx, fieldID, payload)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// GetDefaultValues returns a paginated list of defaults for a custom field.
//
// The results can be filtered by contextID, otherwise all values are returned. If no defaults are set for a context, nothing is returned.
//
// GET /rest/api/{2-3}/field/{fieldID}/context/defaultValue
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#get-custom-field-contexts-default-values
func (i *IssueFieldContextService) GetDefaultValues(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (*model.CustomFieldDefaultValuePageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldContextService).GetDefaultValues", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_field_context_default_values"),
		attribute.String("jira.field.id", fieldID),
		attribute.Int("jira.context.ids_count", len(contextIDs)),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
	)

	result, response, err := i.internalClient.GetDefaultValues(ctx, fieldID, contextIDs, startAt, maxResults)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// SetDefaultValue sets default for contexts of a custom field.
//
// PUT /rest/api/{2-3}/field/{fieldID}/context/defaultValue
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#set-custom-field-contexts-default-values
func (i *IssueFieldContextService) SetDefaultValue(ctx context.Context, fieldID string, payload *model.FieldContextDefaultPayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldContextService).SetDefaultValue", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "set_field_context_default_value"),
		attribute.String("jira.field.id", fieldID),
	)

	response, err := i.internalClient.SetDefaultValue(ctx, fieldID, payload)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

// IssueTypesContext returns a paginated list of context to issue type mappings for a custom field.
//
// 1. Mappings are returned for all contexts or a list of contexts.
//
// 2. Mappings are ordered first by context ID and then by issue type ID.
//
// GET /rest/api/{2-3}/field/{fieldID}/context/issuetypemapping
//
// Docs: TODO: The documentation needs to be created, raise a ticket here: https://github.com/ctreminiom/go-atlassian/issues
func (i *IssueFieldContextService) IssueTypesContext(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (*model.IssueTypeToContextMappingPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldContextService).IssueTypesContext", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_field_context_issue_types"),
		attribute.String("jira.field.id", fieldID),
		attribute.Int("jira.context.ids_count", len(contextIDs)),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
	)

	result, response, err := i.internalClient.IssueTypesContext(ctx, fieldID, contextIDs, startAt, maxResults)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// ProjectsContext returns a paginated list of context to project mappings for a custom field.
//
// 1. The result can be filtered by contextID, or otherwise all mappings are returned.
//
// 2. Invalid IDs are ignored.
//
// GET /rest/api/{2-3}/field/{fieldID}/context/projectmapping
//
// Docs: TODO: The documentation needs to be created, raise a ticket here: https://github.com/ctreminiom/go-atlassian/issues
func (i *IssueFieldContextService) ProjectsContext(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (*model.CustomFieldContextProjectMappingPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldContextService).ProjectsContext", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "get_field_context_projects"),
		attribute.String("jira.field.id", fieldID),
		attribute.Int("jira.context.ids_count", len(contextIDs)),
		attribute.Int("jira.pagination.start_at", startAt),
		attribute.Int("jira.pagination.max_results", maxResults),
	)

	result, response, err := i.internalClient.ProjectsContext(ctx, fieldID, contextIDs, startAt, maxResults)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return result, response, nil
}

// Update updates a custom field context
//
// PUT /rest/api/{2-3}/field/{fieldID}/context/{contextID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#update-custom-field-context
func (i *IssueFieldContextService) Update(ctx context.Context, fieldID string, contextID int, name, description string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldContextService).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "update_field_context"),
		attribute.String("jira.field.id", fieldID),
		attribute.Int("jira.context.id", contextID),
		attribute.String("jira.context.name", name),
	)

	response, err := i.internalClient.Update(ctx, fieldID, contextID, name, description)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

// Delete deletes a custom field context.
//
// DELETE /rest/api/{2-3}/field/{fieldID}/context/{contextID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#delete-custom-field-context
func (i *IssueFieldContextService) Delete(ctx context.Context, fieldID string, contextID int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldContextService).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "delete_field_context"),
		attribute.String("jira.field.id", fieldID),
		attribute.Int("jira.context.id", contextID),
	)

	response, err := i.internalClient.Delete(ctx, fieldID, contextID)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

// AddIssueTypes adds issue types to a custom field context, appending the issue types to the issue types list.
//
// PUT /rest/api/{2-3}/field/{fieldID}/context/{contextID}/issuetype
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#add-issue-types-to-context
func (i *IssueFieldContextService) AddIssueTypes(ctx context.Context, fieldID string, contextID int, issueTypesIDs []string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldContextService).AddIssueTypes", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "add_issue_types_to_context"),
		attribute.String("jira.field.id", fieldID),
		attribute.Int("jira.context.id", contextID),
		attribute.Int("jira.issue_types.count", len(issueTypesIDs)),
	)

	response, err := i.internalClient.AddIssueTypes(ctx, fieldID, contextID, issueTypesIDs)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

// RemoveIssueTypes removes issue types from a custom field context. A custom field context without any issue types applies to all issue types.
//
// POST /rest/api/{2-3}/field/{fieldID}/context/{contextID}/issuetype/remove
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#remove-issue-types-from-context
func (i *IssueFieldContextService) RemoveIssueTypes(ctx context.Context, fieldID string, contextID int, issueTypesIDs []string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldContextService).RemoveIssueTypes", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "remove_issue_types_from_context"),
		attribute.String("jira.field.id", fieldID),
		attribute.Int("jira.context.id", contextID),
		attribute.Int("jira.issue_types.count", len(issueTypesIDs)),
	)

	response, err := i.internalClient.RemoveIssueTypes(ctx, fieldID, contextID, issueTypesIDs)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

// Link assigns a custom field context to projects. If any project in the request is assigned to any context of the custom field, the operation fails.
//
// PUT /rest/api/{2-3}/field/{fieldID}/context/{contextID}/project
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#assign-custom-field-context-to-projects
func (i *IssueFieldContextService) Link(ctx context.Context, fieldID string, contextID int, projectIDs []string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldContextService).Link", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "link_context_to_projects"),
		attribute.String("jira.field.id", fieldID),
		attribute.Int("jira.context.id", contextID),
		attribute.Int("jira.projects.count", len(projectIDs)),
	)

	response, err := i.internalClient.Link(ctx, fieldID, contextID, projectIDs)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

// UnLink removes a custom field context from projects.
//
// 1. A custom field context without any projects applies to all projects.
//
// 2. Removing all projects from a custom field context would result in it applying to all projects.
//
// POST /rest/api/{2-3}/field/{fieldID}/context/{contextID}/project/remove
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#remove-custom-field-context-from-projects
func (i *IssueFieldContextService) UnLink(ctx context.Context, fieldID string, contextID int, projectIDs []string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*IssueFieldContextService).UnLink", spanWithKind(trace.SpanKindClient))
	defer span.End()

	addAttributes(span,
		attribute.String("operation.name", "unlink_context_from_projects"),
		attribute.String("jira.field.id", fieldID),
		attribute.Int("jira.context.id", contextID),
		attribute.Int("jira.projects.count", len(projectIDs)),
	)

	response, err := i.internalClient.UnLink(ctx, fieldID, contextID, projectIDs)
	if err != nil {
		recordError(span, err)
		return response, err
	}

	setOK(span)
	return response, nil
}

type internalIssueFieldContextServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalIssueFieldContextServiceImpl) Gets(ctx context.Context, fieldID string, options *model.FieldContextOptionsScheme, startAt, maxResults int) (*model.CustomFieldContextPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldContextServiceImpl).Gets", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if fieldID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldID)
		recordError(span, err)
		return nil, nil, err
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {
		params.Add("isAnyIssueType", fmt.Sprintf("%v", options.IsAnyIssueType))
		params.Add("isGlobalContext", fmt.Sprintf("%v", options.IsGlobalContext))

		for _, id := range options.ContextID {
			params.Add("contextId", strconv.Itoa(id))
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context?%v", i.version, fieldID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	contexts := new(model.CustomFieldContextPageScheme)
	response, err := i.c.Call(request, contexts)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return contexts, response, nil
}

func (i *internalIssueFieldContextServiceImpl) Create(ctx context.Context, fieldID string, payload *model.FieldContextPayloadScheme) (*model.FieldContextScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldContextServiceImpl).Create", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if fieldID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldID)
		recordError(span, err)
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context", i.version, fieldID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	newContext := new(model.FieldContextScheme)
	response, err := i.c.Call(request, newContext)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return newContext, response, nil
}

func (i *internalIssueFieldContextServiceImpl) GetDefaultValues(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (*model.CustomFieldDefaultValuePageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldContextServiceImpl).GetDefaultValues", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if fieldID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldID)
		recordError(span, err)
		return nil, nil, err
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range contextIDs {
		params.Add("contextId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/defaultValue?%s", i.version, fieldID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	values := new(model.CustomFieldDefaultValuePageScheme)
	response, err := i.c.Call(request, values)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return values, response, nil
}

func (i *internalIssueFieldContextServiceImpl) SetDefaultValue(ctx context.Context, fieldID string, payload *model.FieldContextDefaultPayloadScheme) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldContextServiceImpl).SetDefaultValue", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if fieldID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/defaultValue", i.version, fieldID)

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

func (i *internalIssueFieldContextServiceImpl) IssueTypesContext(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (*model.IssueTypeToContextMappingPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldContextServiceImpl).IssueTypesContext", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if fieldID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldID)
		recordError(span, err)
		return nil, nil, err
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range contextIDs {
		params.Add("contextId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/issuetypemapping?%v", i.version, fieldID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	mapping := new(model.IssueTypeToContextMappingPageScheme)
	response, err := i.c.Call(request, mapping)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return mapping, response, nil
}

func (i *internalIssueFieldContextServiceImpl) ProjectsContext(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (*model.CustomFieldContextProjectMappingPageScheme, *model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldContextServiceImpl).ProjectsContext", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if fieldID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldID)
		recordError(span, err)
		return nil, nil, err
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range contextIDs {
		params.Add("contextId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/projectmapping?%v", i.version, fieldID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		recordError(span, err)
		return nil, nil, err
	}

	mapping := new(model.CustomFieldContextProjectMappingPageScheme)
	response, err := i.c.Call(request, mapping)
	if err != nil {
		recordError(span, err)
		return nil, response, err
	}

	setOK(span)
	return mapping, response, nil
}

func (i *internalIssueFieldContextServiceImpl) Update(ctx context.Context, fieldID string, contextID int, name, description string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldContextServiceImpl).Update", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if fieldID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldID)
		recordError(span, err)
		return nil, err
	}

	if contextID == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldContextID)
		recordError(span, err)
		return nil, err
	}

	payload := map[string]interface{}{"name": name}

	if description != "" {
		payload["description"] = description
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v", i.version, fieldID, contextID)

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

func (i *internalIssueFieldContextServiceImpl) Delete(ctx context.Context, fieldID string, contextID int) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldContextServiceImpl).Delete", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if fieldID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldID)
		recordError(span, err)
		return nil, err
	}

	if contextID == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldContextID)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v", i.version, fieldID, contextID)

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

func (i *internalIssueFieldContextServiceImpl) AddIssueTypes(ctx context.Context, fieldID string, contextID int, issueTypesIDs []string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldContextServiceImpl).AddIssueTypes", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if fieldID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldID)
		recordError(span, err)
		return nil, err
	}

	if len(issueTypesIDs) == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueTypes)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/issuetype", i.version, fieldID, contextID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]interface{}{"issueTypeIds": issueTypesIDs})
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

func (i *internalIssueFieldContextServiceImpl) RemoveIssueTypes(ctx context.Context, fieldID string, contextID int, issueTypesIDs []string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldContextServiceImpl).RemoveIssueTypes", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if fieldID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldID)
		recordError(span, err)
		return nil, err
	}

	if len(issueTypesIDs) == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoIssueTypes)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/issuetype/remove", i.version, fieldID, contextID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"issueTypeIds": issueTypesIDs})
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

func (i *internalIssueFieldContextServiceImpl) Link(ctx context.Context, fieldID string, contextID int, projectIDs []string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldContextServiceImpl).Link", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if fieldID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldID)
		recordError(span, err)
		return nil, err
	}

	if contextID == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldContextID)
		recordError(span, err)
		return nil, err
	}

	if len(projectIDs) == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoProjectIDs)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/project", i.version, fieldID, contextID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]interface{}{"projectIds": projectIDs})
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

func (i *internalIssueFieldContextServiceImpl) UnLink(ctx context.Context, fieldID string, contextID int, projectIDs []string) (*model.ResponseScheme, error) {
	ctx, span := tracer().Start(ctx, "(*internalIssueFieldContextServiceImpl).UnLink", spanWithKind(trace.SpanKindClient))
	defer span.End()

	if fieldID == "" {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldID)
		recordError(span, err)
		return nil, err
	}

	if contextID == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoFieldContextID)
		recordError(span, err)
		return nil, err
	}

	if len(projectIDs) == 0 {
		err := fmt.Errorf("jira: %w", model.ErrNoProjectIDs)
		recordError(span, err)
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/project/remove", i.version, fieldID, contextID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"projectIds": projectIDs})
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
