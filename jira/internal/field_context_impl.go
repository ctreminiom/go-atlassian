package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewIssueFieldContextService creates a new instance of IssueFieldContextService.
// It takes a service.Connector, a version string, and an IssueFieldContextOptionService as input.
// Returns a pointer to IssueFieldContextService and an error if the version is not provided.
func NewIssueFieldContextService(client service.Connector, version string, option *IssueFieldContextOptionService) (*IssueFieldContextService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
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
	return i.internalClient.Gets(ctx, fieldID, options, startAt, maxResults)
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
	return i.internalClient.Create(ctx, fieldID, payload)
}

// GetDefaultValues returns a paginated list of defaults for a custom field.
//
// The results can be filtered by contextID, otherwise all values are returned. If no defaults are set for a context, nothing is returned.
//
// GET /rest/api/{2-3}/field/{fieldID}/context/defaultValue
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#get-custom-field-contexts-default-values
func (i *IssueFieldContextService) GetDefaultValues(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (*model.CustomFieldDefaultValuePageScheme, *model.ResponseScheme, error) {
	return i.internalClient.GetDefaultValues(ctx, fieldID, contextIDs, startAt, maxResults)
}

// SetDefaultValue sets default for contexts of a custom field.
//
// PUT /rest/api/{2-3}/field/{fieldID}/context/defaultValue
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#set-custom-field-contexts-default-values
func (i *IssueFieldContextService) SetDefaultValue(ctx context.Context, fieldID string, payload *model.FieldContextDefaultPayloadScheme) (*model.ResponseScheme, error) {
	return i.internalClient.SetDefaultValue(ctx, fieldID, payload)
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
	return i.internalClient.IssueTypesContext(ctx, fieldID, contextIDs, startAt, maxResults)
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
	return i.internalClient.ProjectsContext(ctx, fieldID, contextIDs, startAt, maxResults)
}

// Update updates a custom field context
//
// PUT /rest/api/{2-3}/field/{fieldID}/context/{contextID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#update-custom-field-context
func (i *IssueFieldContextService) Update(ctx context.Context, fieldID string, contextID int, name, description string) (*model.ResponseScheme, error) {
	return i.internalClient.Update(ctx, fieldID, contextID, name, description)
}

// Delete deletes a custom field context.
//
// DELETE /rest/api/{2-3}/field/{fieldID}/context/{contextID}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#delete-custom-field-context
func (i *IssueFieldContextService) Delete(ctx context.Context, fieldID string, contextID int) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, fieldID, contextID)
}

// AddIssueTypes adds issue types to a custom field context, appending the issue types to the issue types list.
//
// PUT /rest/api/{2-3}/field/{fieldID}/context/{contextID}/issuetype
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#add-issue-types-to-context
func (i *IssueFieldContextService) AddIssueTypes(ctx context.Context, fieldID string, contextID int, issueTypesIDs []string) (*model.ResponseScheme, error) {
	return i.internalClient.AddIssueTypes(ctx, fieldID, contextID, issueTypesIDs)
}

// RemoveIssueTypes removes issue types from a custom field context. A custom field context without any issue types applies to all issue types.
//
// POST /rest/api/{2-3}/field/{fieldID}/context/{contextID}/issuetype/remove
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#remove-issue-types-from-context
func (i *IssueFieldContextService) RemoveIssueTypes(ctx context.Context, fieldID string, contextID int, issueTypesIDs []string) (*model.ResponseScheme, error) {
	return i.internalClient.RemoveIssueTypes(ctx, fieldID, contextID, issueTypesIDs)
}

// Link assigns a custom field context to projects. If any project in the request is assigned to any context of the custom field, the operation fails.
//
// PUT /rest/api/{2-3}/field/{fieldID}/context/{contextID}/project
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#assign-custom-field-context-to-projects
func (i *IssueFieldContextService) Link(ctx context.Context, fieldID string, contextID int, projectIDs []string) (*model.ResponseScheme, error) {
	return i.internalClient.Link(ctx, fieldID, contextID, projectIDs)
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
	return i.internalClient.UnLink(ctx, fieldID, contextID, projectIDs)
}

type internalIssueFieldContextServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalIssueFieldContextServiceImpl) Gets(ctx context.Context, fieldID string, options *model.FieldContextOptionsScheme, startAt, maxResults int) (*model.CustomFieldContextPageScheme, *model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, nil, model.ErrNoFieldID
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
		return nil, nil, err
	}

	contexts := new(model.CustomFieldContextPageScheme)
	response, err := i.c.Call(request, contexts)
	if err != nil {
		return nil, response, err
	}

	return contexts, response, nil
}

func (i *internalIssueFieldContextServiceImpl) Create(ctx context.Context, fieldID string, payload *model.FieldContextPayloadScheme) (*model.FieldContextScheme, *model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, nil, model.ErrNoFieldID
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context", i.version, fieldID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	newContext := new(model.FieldContextScheme)
	response, err := i.c.Call(request, newContext)
	if err != nil {
		return nil, response, err
	}

	return newContext, response, nil
}

func (i *internalIssueFieldContextServiceImpl) GetDefaultValues(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (*model.CustomFieldDefaultValuePageScheme, *model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, nil, model.ErrNoFieldID
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
		return nil, nil, err
	}

	values := new(model.CustomFieldDefaultValuePageScheme)
	response, err := i.c.Call(request, values)
	if err != nil {
		return nil, response, err
	}

	return values, response, nil
}

func (i *internalIssueFieldContextServiceImpl) SetDefaultValue(ctx context.Context, fieldID string, payload *model.FieldContextDefaultPayloadScheme) (*model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, model.ErrNoFieldID
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/defaultValue", i.version, fieldID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldContextServiceImpl) IssueTypesContext(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (*model.IssueTypeToContextMappingPageScheme, *model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, nil, model.ErrNoFieldID
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
		return nil, nil, err
	}

	mapping := new(model.IssueTypeToContextMappingPageScheme)
	response, err := i.c.Call(request, mapping)
	if err != nil {
		return nil, response, err
	}

	return mapping, response, nil
}

func (i *internalIssueFieldContextServiceImpl) ProjectsContext(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (*model.CustomFieldContextProjectMappingPageScheme, *model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, nil, model.ErrNoFieldID
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
		return nil, nil, err
	}

	mapping := new(model.CustomFieldContextProjectMappingPageScheme)
	response, err := i.c.Call(request, mapping)
	if err != nil {
		return nil, response, err
	}

	return mapping, response, nil
}

func (i *internalIssueFieldContextServiceImpl) Update(ctx context.Context, fieldID string, contextID int, name, description string) (*model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, model.ErrNoFieldID
	}

	if contextID == 0 {
		return nil, model.ErrNoFieldContextID
	}

	payload := map[string]interface{}{"name": name}

	if description != "" {
		payload["description"] = description
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v", i.version, fieldID, contextID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldContextServiceImpl) Delete(ctx context.Context, fieldID string, contextID int) (*model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, model.ErrNoFieldID
	}

	if contextID == 0 {
		return nil, model.ErrNoFieldContextID
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v", i.version, fieldID, contextID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldContextServiceImpl) AddIssueTypes(ctx context.Context, fieldID string, contextID int, issueTypesIDs []string) (*model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, model.ErrNoFieldID
	}

	if len(issueTypesIDs) == 0 {
		return nil, model.ErrNoIssueTypes
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/issuetype", i.version, fieldID, contextID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]interface{}{"issueTypeIds": issueTypesIDs})
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldContextServiceImpl) RemoveIssueTypes(ctx context.Context, fieldID string, contextID int, issueTypesIDs []string) (*model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, model.ErrNoFieldID
	}

	if len(issueTypesIDs) == 0 {
		return nil, model.ErrNoIssueTypes
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/issuetype/remove", i.version, fieldID, contextID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"issueTypeIds": issueTypesIDs})
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldContextServiceImpl) Link(ctx context.Context, fieldID string, contextID int, projectIDs []string) (*model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, model.ErrNoFieldID
	}

	if contextID == 0 {
		return nil, model.ErrNoFieldContextID
	}

	if len(projectIDs) == 0 {
		return nil, model.ErrNoProjectIDs
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/project", i.version, fieldID, contextID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", map[string]interface{}{"projectIds": projectIDs})
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldContextServiceImpl) UnLink(ctx context.Context, fieldID string, contextID int, projectIDs []string) (*model.ResponseScheme, error) {

	if fieldID == "" {
		return nil, model.ErrNoFieldID
	}

	if contextID == 0 {
		return nil, model.ErrNoFieldContextID
	}

	if len(projectIDs) == 0 {
		return nil, model.ErrNoProjectIDs
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/project/remove", i.version, fieldID, contextID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", map[string]interface{}{"projectIds": projectIDs})
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
