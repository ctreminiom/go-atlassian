package jira

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type FieldContextService struct {
	client *Client
	Option *FieldOptionContextService
}

type FieldContextOptionsScheme struct {
	IsAnyIssueType  bool
	IsGlobalContext bool
	ContextID       []int
}

type CustomFieldContextPageScheme struct {
	MaxResults int                   `json:"maxResults,omitempty"`
	StartAt    int                   `json:"startAt,omitempty"`
	Total      int                   `json:"total,omitempty"`
	IsLast     bool                  `json:"isLast,omitempty"`
	Values     []*FieldContextScheme `json:"values,omitempty"`
}

type FieldContextScheme struct {
	ID              string   `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Description     string   `json:"description,omitempty"`
	IsGlobalContext bool     `json:"isGlobalContext,omitempty"`
	IsAnyIssueType  bool     `json:"isAnyIssueType,omitempty"`
	ProjectIds      []string `json:"projectIds,omitempty"`
	IssueTypeIds    []string `json:"issueTypeIds,omitempty"`
}

// Gets returns a paginated list of contexts for a custom field. Contexts can be returned as follows:
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#get-custom-field-contexts
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-get
func (f *FieldContextService) Gets(ctx context.Context, fieldID string, opts *FieldContextOptionsScheme, startAt, maxResults int) (
	result *CustomFieldContextPageScheme, response *ResponseScheme, err error) {

	if fieldID == "" {
		return nil, nil, notFieldIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if opts.IsAnyIssueType {
			params.Add("isAnyIssueType", "true")
		}

		if opts.IsGlobalContext {
			params.Add("isGlobalContext", "true")
		}

		for _, contextID := range opts.ContextID {
			params.Add("contextId", strconv.Itoa(contextID))
		}

	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context?%v", fieldID, params.Encode())

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type FieldContextPayloadScheme struct {
	IssueTypeIDs []int  `json:"issueTypeIds,omitempty"`
	ProjectIDs   []int  `json:"projectIds,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
}

// Create creates a custom field context.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#create-custom-field-context
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-post
func (f *FieldContextService) Create(ctx context.Context, fieldID string, payload *FieldContextPayloadScheme) (
	result *FieldContextScheme, response *ResponseScheme, err error) {

	if fieldID == "" {
		return nil, nil, notFieldIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context", fieldID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type CustomFieldDefaultValuePageScheme struct {
	MaxResults int                              `json:"maxResults,omitempty"`
	StartAt    int                              `json:"startAt,omitempty"`
	Total      int                              `json:"total,omitempty"`
	IsLast     bool                             `json:"isLast,omitempty"`
	Values     []*CustomFieldDefaultValueScheme `json:"values,omitempty"`
}

type CustomFieldDefaultValueScheme struct {
	ContextID         string   `json:"contextId,omitempty"`
	OptionID          string   `json:"optionId,omitempty"`
	CascadingOptionID string   `json:"cascadingOptionId,omitempty"`
	OptionIDs         []string `json:"optionIds,omitempty"`
	Type              string   `json:"type,omitempty"`
}

// GetDefaultValues returns a paginated list of defaults for a custom field.
// The results can be filtered by contextId, otherwise all values are returned.
// If no defaults are set for a context, nothing is returned.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#get-custom-field-contexts-default-values
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-defaultvalue-get
func (f *FieldContextService) GetDefaultValues(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (
	result *CustomFieldDefaultValuePageScheme, response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, nil, notFieldIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, contextID := range contextIDs {
		params.Add("contextId", strconv.Itoa(contextID))
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/defaultValue?%s", fieldID, params.Encode())

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type FieldContextDefaultPayloadScheme struct {
	DefaultValues []*CustomFieldDefaultValueScheme `json:"defaultValues,omitempty"`
}

// SetDefaultValue sets default for contexts of a custom field.
// Default are defined using these objects:
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#set-custom-field-contexts-default-values
// Official Docs; https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-defaultvalue-put
func (f *FieldContextService) SetDefaultValue(ctx context.Context, fieldID string, payload *FieldContextDefaultPayloadScheme) (
	response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, notFieldIDError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	if len(payload.DefaultValues) == 0 {
		return nil, notFieldContextDefaultError
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/defaultValue", fieldID)

	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

type IssueTypeToContextMappingPageScheme struct {
	MaxResults int                                     `json:"maxResults,omitempty"`
	StartAt    int                                     `json:"startAt,omitempty"`
	Total      int                                     `json:"total,omitempty"`
	IsLast     bool                                    `json:"isLast,omitempty"`
	Values     []*IssueTypeToContextMappingValueScheme `json:"values"`
}

type IssueTypeToContextMappingValueScheme struct {
	ContextID      string `json:"contextId"`
	IsAnyIssueType bool   `json:"isAnyIssueType,omitempty"`
	IssueTypeID    string `json:"issueTypeId,omitempty"`
}

// IssueTypesContext returns a paginated list of context to issue type mappings for a custom field.
// Mappings are returned for all contexts or a list of contexts.
// Mappings are ordered first by context ID and then by issue type ID.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-issuetypemapping-get
// Docs: TODO Issue 51
func (f *FieldContextService) IssueTypesContext(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (
	result *IssueTypeToContextMappingPageScheme, response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, nil, notFieldIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, contextID := range contextIDs {
		params.Add("contextId", strconv.Itoa(contextID))
	}

	var endpoint = fmt.Sprintf("/rest/api/3/field/%v/context/issuetypemapping?%v", fieldID, params.Encode())

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type CustomFieldContextProjectMappingPageScheme struct {
	Self       string                                         `json:"self,omitempty"`
	NextPage   string                                         `json:"nextPage,omitempty"`
	MaxResults int                                            `json:"maxResults,omitempty"`
	StartAt    int                                            `json:"startAt,omitempty"`
	Total      int                                            `json:"total,omitempty"`
	IsLast     bool                                           `json:"isLast,omitempty"`
	Values     []*CustomFieldContextProjectMappingValueScheme `json:"values,omitempty"`
}

type CustomFieldContextProjectMappingValueScheme struct {
	ContextID       string `json:"contextId,omitempty"`
	ProjectID       string `json:"projectId,omitempty"`
	IsGlobalContext bool   `json:"isGlobalContext,omitempty"`
}

// ProjectsContext returns a paginated list of context to project mappings for a custom field.
// The result can be filtered by contextId,
// or otherwise all mappings are returned.
// Invalid IDs are ignored.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-projectmapping-get
// Docs: TODO Issue 52
func (f *FieldContextService) ProjectsContext(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (
	result *CustomFieldContextProjectMappingPageScheme, response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, nil, notFieldIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, contextID := range contextIDs {
		params.Add("contextId", strconv.Itoa(contextID))
	}

	var endpoint = fmt.Sprintf("/rest/api/3/field/%v/context/projectmapping?%v", fieldID, params.Encode())

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update updates a custom field context
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#update-custom-field-context
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-contextid-put
func (f *FieldContextService) Update(ctx context.Context, fieldID string, contextID int, name, description string) (
	response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, notFieldIDError
	}

	if contextID == 0 {
		return nil, notContextIDError
	}

	payload := struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v", fieldID, contextID)

	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Delete deletes a custom field context.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#delete-custom-field-context
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-contextid-delete
func (f *FieldContextService) Delete(ctx context.Context, fieldID string, contextID int) (response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, notFieldIDError
	}

	if contextID == 0 {
		return nil, notContextIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v", fieldID, contextID)

	request, err := f.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// AddIssueTypes adds issue types to a custom field context, appending the issue types to the issue types list.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#add-issue-types-to-context
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-contextid-issuetype-put
func (f *FieldContextService) AddIssueTypes(ctx context.Context, fieldID string, contextID int, issueTypesIDs []string) (
	response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, notFieldIDError
	}

	if len(issueTypesIDs) == 0 {
		return nil, notIssueTypesError
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/issuetype", fieldID, contextID)

	payload := struct {
		IssueTypeIds []string `json:"issueTypeIds"`
	}{
		IssueTypeIds: issueTypesIDs,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// RemoveIssueTypes removes issue types from a custom field context.
// A custom field context without any issue types applies to all issue types.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#remove-issue-types-from-context
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-contextid-issuetype-remove-post
func (f *FieldContextService) RemoveIssueTypes(ctx context.Context, fieldID string, contextID int, issueTypesIDs []string) (
	response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, notFieldIDError
	}

	if len(issueTypesIDs) == 0 {
		return nil, notIssueTypesError
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/issuetype/remove", fieldID, contextID)

	payload := struct {
		IssueTypeIds []string `json:"issueTypeIds"`
	}{
		IssueTypeIds: issueTypesIDs,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Link assigns a custom field context to projects.
// If any project in the request is assigned to any context of the custom field, the operation fails.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#assign-custom-field-context-to-projects
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-contextid-project-put
func (f *FieldContextService) Link(ctx context.Context, fieldID string, contextID int, projectIDs []string) (
	response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, notFieldIDError
	}

	if len(projectIDs) == 0 {
		return nil, notProjectsError
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/project", fieldID, contextID)

	payload := struct {
		ProjectIds []string `json:"projectIds"`
	}{
		ProjectIds: projectIDs,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// UnLink removes a custom field context from projects.
// A custom field context without any projects applies to all projects.
// Removing all projects from a custom field context would result in it applying to all projects.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#remove-custom-field-context-from-projects
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-contextid-project-remove-post
func (f *FieldContextService) UnLink(ctx context.Context, fieldID string, contextID int, projectIDs []string) (response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, notFieldIDError
	}

	if len(projectIDs) == 0 {
		return nil, notProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/project/remove", fieldID, contextID)

	payload := struct {
		ProjectIds []string `json:"projectIds"`
	}{
		ProjectIds: projectIDs,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

var (
	notFieldIDError             = fmt.Errorf("error, fieldID value is nil, please provide a valid fieldID value")
	notFieldContextDefaultError = fmt.Errorf("error, please provide a valid Custom Field Context default value")
	notContextIDError           = fmt.Errorf("error, please provide a valid contextID value")
	notIssueTypesError          = fmt.Errorf("error, please provide a valid issueTypesIDs value")
	notProjectsError            = fmt.Errorf("error, please provide a valid projectIDs value")
)
