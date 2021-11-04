package v3

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"net/http"
	"net/url"
	"strconv"
)

type FieldContextService struct {
	client *Client
	Option *FieldOptionContextService
}

// Gets returns a paginated list of contexts for a custom field. Contexts can be returned as follows:
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#get-custom-field-contexts
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-get
func (f *FieldContextService) Gets(ctx context.Context, fieldID string, opts *models.FieldContextOptionsScheme, startAt, maxResults int) (
	result *models.CustomFieldContextPageScheme, response *ResponseScheme, err error) {

	if fieldID == "" {
		return nil, nil, models.ErrNoFieldIDError
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

// Create creates a custom field context.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#create-custom-field-context
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-post
func (f *FieldContextService) Create(ctx context.Context, fieldID string, payload *models.FieldContextPayloadScheme) (
	result *models.FieldContextScheme, response *ResponseScheme, err error) {

	if fieldID == "" {
		return nil, nil, models.ErrNoFieldIDError
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

// GetDefaultValues returns a paginated list of defaults for a custom field.
// The results can be filtered by contextId, otherwise all values are returned.
// If no defaults are set for a context, nothing is returned.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#get-custom-field-contexts-default-values
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-defaultvalue-get
func (f *FieldContextService) GetDefaultValues(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (
	result *models.CustomFieldDefaultValuePageScheme, response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, nil, models.ErrNoFieldIDError
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

// SetDefaultValue sets default for contexts of a custom field.
// Default are defined using these objects:
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#set-custom-field-contexts-default-values
// Official Docs; https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-defaultvalue-put
func (f *FieldContextService) SetDefaultValue(ctx context.Context, fieldID string, payload *models.FieldContextDefaultPayloadScheme) (
	response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, models.ErrNoFieldIDError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	if len(payload.DefaultValues) == 0 {
		return nil, models.ErrNoFieldContextIDError
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

// IssueTypesContext returns a paginated list of context to issue type mappings for a custom field.
// Mappings are returned for all contexts or a list of contexts.
// Mappings are ordered first by context ID and then by issue type ID.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-issuetypemapping-get
// Docs: TODO Issue 51
func (f *FieldContextService) IssueTypesContext(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (
	result *models.IssueTypeToContextMappingPageScheme, response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, nil, models.ErrNoFieldIDError
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

// ProjectsContext returns a paginated list of context to project mappings for a custom field.
// The result can be filtered by contextId,
// or otherwise all mappings are returned.
// Invalid IDs are ignored.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-projectmapping-get
// Docs: TODO Issue 52
func (f *FieldContextService) ProjectsContext(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (
	result *models.CustomFieldContextProjectMappingPageScheme, response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, nil, models.ErrNoFieldIDError
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
		return nil, models.ErrNoFieldIDError
	}

	if contextID == 0 {
		return nil, models.ErrNoFieldContextIDError
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
		return nil, models.ErrNoFieldIDError
	}

	if contextID == 0 {
		return nil, models.ErrNoFieldContextIDError
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
		return nil, models.ErrNoFieldIDError
	}

	if len(issueTypesIDs) == 0 {
		return nil, models.ErrNoIssueTypesError
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
		return nil, models.ErrNoFieldIDError
	}

	if len(issueTypesIDs) == 0 {
		return nil, models.ErrNoIssueTypesError
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
		return nil, models.ErrNoFieldIDError
	}

	if len(projectIDs) == 0 {
		return nil, models.ErrNoIssueTypesError
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
func (f *FieldContextService) UnLink(ctx context.Context, fieldID string, contextID int, projectIDs []string) (response *ResponseScheme,
	err error) {

	if len(fieldID) == 0 {
		return nil, models.ErrNoFieldIDError
	}

	if len(projectIDs) == 0 {
		return nil, models.ErrNoProjectIDError
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
