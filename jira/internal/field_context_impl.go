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
)

func NewIssueFieldContextService(client service.Client, version string, option *IssueFieldContextOptionService) (*IssueFieldContextService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &IssueFieldContextService{
		internalClient: &internalIssueFieldContextServiceImpl{c: client, version: version},
		Option:         option,
	}, nil
}

type IssueFieldContextService struct {
	internalClient jira.FieldContextConnector
	Option         *IssueFieldContextOptionService
}

// Gets returns a paginated list of contexts for a custom field. Contexts can be returned as follows:
//
// 1. By defining id only, all contexts from the list of IDs.
//
// 2. By defining isAnyIssueType
//
// 3. By defining isGlobalContext
//
// GET /rest/api/{2-3}/field/{fieldId}/context
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#get-custom-field-contexts
func (i *IssueFieldContextService) Gets(ctx context.Context, fieldId string, options *model.FieldContextOptionsScheme, startAt, maxResults int) (*model.CustomFieldContextPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.Gets(ctx, fieldId, options, startAt, maxResults)
}

// Create creates a custom field context.
//
// 1. If projectIds is empty, a global context is created. A global context is one that applies to all project.
//
// 2. If issueTypeIds is empty, the context applies to all issue types.
//
// POST /rest/api/{2-3}/field/{fieldId}/context
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#create-custom-field-context
func (i *IssueFieldContextService) Create(ctx context.Context, fieldId string, payload *model.FieldContextPayloadScheme) (*model.FieldContextScheme, *model.ResponseScheme, error) {
	return i.internalClient.Create(ctx, fieldId, payload)
}

// GetDefaultValues returns a paginated list of defaults for a custom field.
//
// The results can be filtered by contextId, otherwise all values are returned. If no defaults are set for a context, nothing is returned.
//
// GET /rest/api/{2-3}/field/{fieldId}/context/defaultValue
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#get-custom-field-contexts-default-values
func (i *IssueFieldContextService) GetDefaultValues(ctx context.Context, fieldId string, contextIds []int, startAt, maxResults int) (*model.CustomFieldDefaultValuePageScheme, *model.ResponseScheme, error) {
	return i.internalClient.GetDefaultValues(ctx, fieldId, contextIds, startAt, maxResults)
}

// SetDefaultValue sets default for contexts of a custom field.
//
// PUT /rest/api/{2-3}/field/{fieldId}/context/defaultValue
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#set-custom-field-contexts-default-values
func (i *IssueFieldContextService) SetDefaultValue(ctx context.Context, fieldId string, payload *model.FieldContextDefaultPayloadScheme) (*model.ResponseScheme, error) {
	return i.internalClient.SetDefaultValue(ctx, fieldId, payload)
}

// IssueTypesContext returns a paginated list of context to issue type mappings for a custom field.
//
// 1. Mappings are returned for all contexts or a list of contexts.
//
// 2. Mappings are ordered first by context ID and then by issue type ID.
//
// GET /rest/api/{2-3}/field/{fieldId}/context/issuetypemapping
//
// Docs: TODO: The documentation needs to be created, raise a ticket here: https://github.com/ctreminiom/go-atlassian/issues
func (i *IssueFieldContextService) IssueTypesContext(ctx context.Context, fieldId string, contextIds []int, startAt, maxResults int) (*model.IssueTypeToContextMappingPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.IssueTypesContext(ctx, fieldId, contextIds, startAt, maxResults)
}

// ProjectsContext returns a paginated list of context to project mappings for a custom field.
//
// 1. The result can be filtered by contextId, or otherwise all mappings are returned.
//
// 2. Invalid IDs are ignored.
//
// GET /rest/api/{2-3}/field/{fieldId}/context/projectmapping
//
// Docs: TODO: The documentation needs to be created, raise a ticket here: https://github.com/ctreminiom/go-atlassian/issues
func (i *IssueFieldContextService) ProjectsContext(ctx context.Context, fieldId string, contextIds []int, startAt, maxResults int) (*model.CustomFieldContextProjectMappingPageScheme, *model.ResponseScheme, error) {
	return i.internalClient.ProjectsContext(ctx, fieldId, contextIds, startAt, maxResults)
}

// Update updates a custom field context
//
// PUT /rest/api/{2-3}/field/{fieldId}/context/{contextId}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#update-custom-field-context
func (i *IssueFieldContextService) Update(ctx context.Context, fieldId string, contextId int, name, description string) (*model.ResponseScheme, error) {
	return i.internalClient.Update(ctx, fieldId, contextId, name, description)
}

// Delete deletes a custom field context.
//
// DELETE /rest/api/{2-3}/field/{fieldId}/context/{contextId}
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#delete-custom-field-context
func (i *IssueFieldContextService) Delete(ctx context.Context, fieldId string, contextId int) (*model.ResponseScheme, error) {
	return i.internalClient.Delete(ctx, fieldId, contextId)
}

// AddIssueTypes adds issue types to a custom field context, appending the issue types to the issue types list.
//
// PUT /rest/api/{2-3}/field/{fieldId}/context/{contextId}/issuetype
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#add-issue-types-to-context
func (i *IssueFieldContextService) AddIssueTypes(ctx context.Context, fieldId string, contextId int, issueTypesIds []string) (*model.ResponseScheme, error) {
	return i.internalClient.AddIssueTypes(ctx, fieldId, contextId, issueTypesIds)
}

// RemoveIssueTypes removes issue types from a custom field context. A custom field context without any issue types applies to all issue types.
//
// POST /rest/api/{2-3}/field/{fieldId}/context/{contextId}/issuetype/remove
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#remove-issue-types-from-context
func (i *IssueFieldContextService) RemoveIssueTypes(ctx context.Context, fieldId string, contextId int, issueTypesIds []string) (*model.ResponseScheme, error) {
	return i.internalClient.RemoveIssueTypes(ctx, fieldId, contextId, issueTypesIds)
}

// Link assigns a custom field context to projects. If any project in the request is assigned to any context of the custom field, the operation fails.
//
// PUT /rest/api/{2-3}/field/{fieldId}/context/{contextId}/project
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#assign-custom-field-context-to-projects
func (i *IssueFieldContextService) Link(ctx context.Context, fieldId string, contextId int, projectIds []string) (*model.ResponseScheme, error) {
	return i.internalClient.Link(ctx, fieldId, contextId, projectIds)
}

// UnLink removes a custom field context from projects.
//
// 1. A custom field context without any projects applies to all projects.
//
// 2. Removing all projects from a custom field context would result in it applying to all projects.
//
// POST /rest/api/{2-3}/field/{fieldId}/context/{contextId}/project/remove
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#remove-custom-field-context-from-projects
func (i *IssueFieldContextService) UnLink(ctx context.Context, fieldId string, contextId int, projectIds []string) (*model.ResponseScheme, error) {
	return i.internalClient.UnLink(ctx, fieldId, contextId, projectIds)
}

type internalIssueFieldContextServiceImpl struct {
	c       service.Client
	version string
}

func (i *internalIssueFieldContextServiceImpl) Gets(ctx context.Context, fieldId string, options *model.FieldContextOptionsScheme, startAt, maxResults int) (*model.CustomFieldContextPageScheme, *model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, nil, model.ErrNoFieldIDError
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

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context?%v", i.version, fieldId, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

func (i *internalIssueFieldContextServiceImpl) Create(ctx context.Context, fieldId string, payload *model.FieldContextPayloadScheme) (*model.FieldContextScheme, *model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, nil, model.ErrNoFieldIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context", i.version, fieldId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
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

func (i *internalIssueFieldContextServiceImpl) GetDefaultValues(ctx context.Context, fieldId string, contextIds []int, startAt, maxResults int) (*model.CustomFieldDefaultValuePageScheme, *model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, nil, model.ErrNoFieldIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range contextIds {
		params.Add("contextId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/defaultValue?%s", i.version, fieldId, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

func (i *internalIssueFieldContextServiceImpl) SetDefaultValue(ctx context.Context, fieldId string, payload *model.FieldContextDefaultPayloadScheme) (*model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, model.ErrNoFieldIDError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/defaultValue", i.version, fieldId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldContextServiceImpl) IssueTypesContext(ctx context.Context, fieldId string, contextIds []int, startAt, maxResults int) (*model.IssueTypeToContextMappingPageScheme, *model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, nil, model.ErrNoFieldIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range contextIds {
		params.Add("contextId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/issuetypemapping?%v", i.version, fieldId, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

func (i *internalIssueFieldContextServiceImpl) ProjectsContext(ctx context.Context, fieldId string, contextIds []int, startAt, maxResults int) (*model.CustomFieldContextProjectMappingPageScheme, *model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, nil, model.ErrNoFieldIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range contextIds {
		params.Add("contextId", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/projectmapping?%v", i.version, fieldId, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

func (i *internalIssueFieldContextServiceImpl) Update(ctx context.Context, fieldId string, contextId int, name, description string) (*model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, model.ErrNoFieldIDError
	}

	if contextId == 0 {
		return nil, model.ErrNoFieldContextIDError
	}

	payload := struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v", i.version, fieldId, contextId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldContextServiceImpl) Delete(ctx context.Context, fieldId string, contextId int) (*model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, model.ErrNoFieldIDError
	}

	if contextId == 0 {
		return nil, model.ErrNoFieldContextIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v", i.version, fieldId, contextId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldContextServiceImpl) AddIssueTypes(ctx context.Context, fieldId string, contextId int, issueTypesIds []string) (*model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, model.ErrNoFieldIDError
	}

	if len(issueTypesIds) == 0 {
		return nil, model.ErrNoIssueTypesError
	}

	payload := struct {
		IssueTypeIds []string `json:"issueTypeIds"`
	}{
		IssueTypeIds: issueTypesIds,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/issuetype", i.version, fieldId, contextId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldContextServiceImpl) RemoveIssueTypes(ctx context.Context, fieldId string, contextId int, issueTypesIds []string) (*model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, model.ErrNoFieldIDError
	}

	if len(issueTypesIds) == 0 {
		return nil, model.ErrNoIssueTypesError
	}

	payload := struct {
		IssueTypeIds []string `json:"issueTypeIds"`
	}{
		IssueTypeIds: issueTypesIds,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/issuetype/remove", i.version, fieldId, contextId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldContextServiceImpl) Link(ctx context.Context, fieldId string, contextId int, projectIds []string) (*model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, model.ErrNoFieldIDError
	}

	if len(projectIds) == 0 {
		return nil, model.ErrNoIssueTypesError
	}

	payload := struct {
		ProjectIds []string `json:"projectIds"`
	}{
		ProjectIds: projectIds,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/project", i.version, fieldId, contextId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalIssueFieldContextServiceImpl) UnLink(ctx context.Context, fieldId string, contextId int, projectIds []string) (*model.ResponseScheme, error) {

	if fieldId == "" {
		return nil, model.ErrNoFieldIDError
	}

	if len(projectIds) == 0 {
		return nil, model.ErrNoIssueTypesError
	}

	payload := struct {
		ProjectIds []string `json:"projectIds"`
	}{
		ProjectIds: projectIds,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/field/%v/context/%v/project/remove", i.version, fieldId, contextId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
