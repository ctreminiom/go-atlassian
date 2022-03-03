package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ProjectService struct {
	client     *Client
	Category   *ProjectCategoryService
	Component  *ProjectComponentService
	Valid      *ProjectValidationService
	Permission *ProjectPermissionSchemeService
	Role       *ProjectRoleService
	Type       *ProjectTypeService
	Version    *ProjectVersionService
	Feature    *ProjectFeatureService
	Property   *ProjectPropertyService
}

// Create creates a project based on a project type template, as shown in the following table:
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-post
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-post
func (p *ProjectService) Create(ctx context.Context, payload *models.ProjectPayloadScheme) (result *models.NewProjectCreatedScheme,
	response *ResponseScheme, err error) {

	var endpoint = "rest/api/2/project"

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Search returns a paginated list of projects visible to the user.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-search-get
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-search-get
func (p *ProjectService) Search(ctx context.Context, options *models.ProjectSearchOptionsScheme, startAt, maxResults int) (
	result *models.ProjectSearchScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}

		if len(options.OrderBy) != 0 {
			params.Add("orderBy", options.OrderBy)
		}

		if len(options.Query) != 0 {
			params.Add("query", options.Query)
		}

		if len(options.ProjectKeyType) != 0 {
			params.Add("typeKey", options.ProjectKeyType)
		}

		if options.CategoryID != 0 {
			params.Add("categoryId", strconv.Itoa(options.CategoryID))
		}

		if len(options.Action) != 0 {
			params.Add("action", options.Action)
		}
	}

	var endpoint = fmt.Sprintf("rest/api/2/project/search?%v", params.Encode())

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Gets returns all projects visible to the user.
// Deprecated, use Get projects paginated that supports search and pagination.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-get
// Note: available on the Data Center versions
func (p *ProjectService) Gets(ctx context.Context, expand []string) (result []*models.ProjectScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString("rest/api/2/project")

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns the project details for a project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-projectidorkey-get
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-projectidorkey-get
func (p *ProjectService) Get(ctx context.Context, projectKeyOrID string, expand []string) (result *models.ProjectScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, models.ErrNoProjectIDError
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/2/project/%v", projectKeyOrID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update updates the project details of a project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-projectidorkey-put
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-projectidorkey-put
func (p *ProjectService) Update(ctx context.Context, projectKeyOrID string, payload *models.ProjectUpdateScheme) (result *models.ProjectScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, models.ErrNoProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/project/%v", projectKeyOrID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := p.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes a project.
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-projectidorkey-delete
func (p *ProjectService) Delete(ctx context.Context, projectKeyOrID string, enableUndo bool) (response *ResponseScheme,
	err error) {

	if len(projectKeyOrID) == 0 {
		return nil, models.ErrNoProjectIDError
	}

	params := url.Values{}
	if enableUndo {
		params.Add("enableUndo", "true")
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/2/project/%v", projectKeyOrID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := p.client.newRequest(ctx, http.MethodDelete, endpoint.String(), nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// DeleteAsynchronously deletes a project asynchronously.
// 1. transactional, that is, if part of to delete fails the project is not deleted.
// 2. asynchronous. Follow the location link in the response to determine the status of the task and use Get task to obtain subsequent updates.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-projectidorkey-delete-post
func (p *ProjectService) DeleteAsynchronously(ctx context.Context, projectKeyOrID string) (result *models.TaskScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, models.ErrNoProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/project/%v/delete", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Archive archives a project. Archived projects cannot be deleted.
// To delete an archived project, restore the project and then delete it.
// To restore a project, use the Jira UI.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-projectidorkey-archive-post
func (p *ProjectService) Archive(ctx context.Context, projectKeyOrID string) (response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, models.ErrNoProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/project/%v/archive", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Restore restores a project from the Jira recycle bin.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-projectidorkey-restore-post
func (p *ProjectService) Restore(ctx context.Context, projectKeyOrID string) (result *models.ProjectScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, models.ErrNoProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/project/%v/restore", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Statuses returns the valid statuses for a project.
// The statuses are grouped by issue type, as each project has a set of valid issue types and each issue type has a set of valid statuses.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-projectidorkey-statuses-get
func (p *ProjectService) Statuses(ctx context.Context, projectKeyOrID string) (result []*models.ProjectStatusPageScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, models.ErrNoProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/project/%v/statuses", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Hierarchy get the issue type hierarchy for a next-gen project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-projectid-hierarchy-get
func (p *ProjectService) Hierarchy(ctx context.Context, projectKeyOrID string) (result *models.ProjectHierarchyScheme,
	response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, models.ErrNoProjectIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/project/%v/hierarchy", projectKeyOrID)

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// NotificationScheme search a notification scheme associated with the project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-projects/#api-rest-api-2-project-projectkeyorid-notificationscheme-get
func (p *ProjectService) NotificationScheme(ctx context.Context, projectKeyOrID string, expand []string) (
	result *models.NotificationSchemeScheme, response *ResponseScheme, err error) {

	if len(projectKeyOrID) == 0 {
		return nil, nil, models.ErrNoProjectIDError
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/2/project/%v/notificationscheme", projectKeyOrID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := p.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	response, err = p.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
