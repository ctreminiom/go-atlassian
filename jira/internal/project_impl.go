package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// ProjectChildServices holds various child services related to project management in Jira Service Management.
type ProjectChildServices struct {
	// Category is the service for managing project categories.
	Category *ProjectCategoryService
	// Component is the service for managing project components.
	Component *ProjectComponentService
	// Feature is the service for managing project features.
	Feature *ProjectFeatureService
	// Permission is the service for managing project permission schemes.
	Permission *ProjectPermissionSchemeService
	// Property is the service for managing project properties.
	Property *ProjectPropertyService
	// Role is the service for managing project roles.
	Role *ProjectRoleService
	// Type is the service for managing project types.
	Type *ProjectTypeService
	// Validator is the service for managing project validators.
	Validator *ProjectValidatorService
	// Version is the service for managing project versions.
	Version *ProjectVersionService
}

// NewProjectService creates a new instance of ProjectService.
func NewProjectService(client service.Connector, version string, subServices *ProjectChildServices) (*ProjectService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ProjectService{
		internalClient: &internalProjectImpl{c: client, version: version},
		Category:       subServices.Category,
		Component:      subServices.Component,
		Feature:        subServices.Feature,
		Permission:     subServices.Permission,
		Property:       subServices.Property,
		Role:           subServices.Role,
		Type:           subServices.Type,
		Validator:      subServices.Validator,
		Version:        subServices.Version,
	}, nil
}

// ProjectService provides methods to manage projects in Jira Service Management.
type ProjectService struct {
	// internalClient is the connector interface for project operations.
	internalClient jira.ProjectConnector
	// Category is the service for managing project categories.
	Category *ProjectCategoryService
	// Component is the service for managing project components.
	Component *ProjectComponentService
	// Feature is the service for managing project features.
	Feature *ProjectFeatureService
	// Permission is the service for managing project permission schemes.
	Permission *ProjectPermissionSchemeService
	// Property is the service for managing project properties.
	Property *ProjectPropertyService
	// Role is the service for managing project roles.
	Role *ProjectRoleService
	// Type is the service for managing project types.
	Type *ProjectTypeService
	// Validator is the service for managing project validators.
	Validator *ProjectValidatorService
	// Version is the service for managing project versions.
	Version *ProjectVersionService
}

// Create creates a project based on a project type template
//
// POST /rest/api/{2-3}/project
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#create-project
func (p *ProjectService) Create(ctx context.Context, payload *model.ProjectPayloadScheme) (*model.NewProjectCreatedScheme, *model.ResponseScheme, error) {
	return p.internalClient.Create(ctx, payload)
}

// Search returns a paginated list of projects visible to the user.
//
// GET /rest/api/{2-3}/project/search
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#get-projects-paginated
func (p *ProjectService) Search(ctx context.Context, options *model.ProjectSearchOptionsScheme, startAt, maxResults int) (*model.ProjectSearchScheme, *model.ResponseScheme, error) {
	return p.internalClient.Search(ctx, options, startAt, maxResults)
}

// Get returns the project details for a project.
//
// GET /rest/api/{2-3}project/{projectKeyOrID}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#get-project
func (p *ProjectService) Get(ctx context.Context, projectKeyOrID string, expand []string) (*model.ProjectScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, projectKeyOrID, expand)
}

// Update updates the project details of a project.
//
// PUT /rest/api/{2-3}/project/{projectKeyOrID}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#update-project
func (p *ProjectService) Update(ctx context.Context, projectKeyOrID string, payload *model.ProjectUpdateScheme) (*model.ProjectScheme, *model.ResponseScheme, error) {
	return p.internalClient.Update(ctx, projectKeyOrID, payload)
}

// Delete deletes a project.
//
// You can't delete a project if it's archived. To delete an archived project, restore the project and then delete it.
//
// To restore a project, use the Jira UI.
//
// DELETE /rest/api/{2-3}/project/{projectKeyOrID}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#delete-project
func (p *ProjectService) Delete(ctx context.Context, projectKeyOrID string, enableUndo bool) (*model.ResponseScheme, error) {
	return p.internalClient.Delete(ctx, projectKeyOrID, enableUndo)
}

// DeleteAsynchronously deletes a project asynchronously.
//
// 1. transactional, that is, if part of to delete fails the project is not deleted.
//
// 2. asynchronous. Follow the location link in the response to determine the status of the task and use Get task to obtain subsequent updates.
//
// POST /rest/api/{2-3}/project/{projectKeyOrID}/delete
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#delete-project-asynchronously
func (p *ProjectService) DeleteAsynchronously(ctx context.Context, projectKeyOrID string) (*model.TaskScheme, *model.ResponseScheme, error) {
	return p.internalClient.DeleteAsynchronously(ctx, projectKeyOrID)
}

// Archive archives a project. Archived projects cannot be deleted.
//
// To delete an archived project, restore the project and then delete it.
//
// To restore a project, use the Jira UI.
//
// POST /rest/api/{2-3}/project/{projectKeyOrID}/archive
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#archive-project
func (p *ProjectService) Archive(ctx context.Context, projectKeyOrID string) (*model.ResponseScheme, error) {
	return p.internalClient.Archive(ctx, projectKeyOrID)
}

// Restore restores a project from the Jira recycle bin.
//
// POST /rest/api/3/project/{projectKeyOrID}/restore
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#restore-deleted-project
func (p *ProjectService) Restore(ctx context.Context, projectKeyOrID string) (*model.ProjectScheme, *model.ResponseScheme, error) {
	return p.internalClient.Restore(ctx, projectKeyOrID)
}

// Statuses returns the valid statuses for a project.
//
// The statuses are grouped by issue type, as each project has a set of valid issue types and each issue type has a set of valid statuses.
//
// GET /rest/api/{2-3}/project/{projectKeyOrID}/statuses
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#get-all-statuses-for-project
func (p *ProjectService) Statuses(ctx context.Context, projectKeyOrID string) ([]*model.ProjectStatusPageScheme, *model.ResponseScheme, error) {
	return p.internalClient.Statuses(ctx, projectKeyOrID)
}

// NotificationScheme gets the notification scheme associated with the project.
//
// GET /rest/api/{2-3}/project/{projectKeyOrID}/notificationscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#get-project-notification-scheme
func (p *ProjectService) NotificationScheme(ctx context.Context, projectKeyOrID string, expand []string) (*model.NotificationSchemeScheme, *model.ResponseScheme, error) {
	return p.internalClient.NotificationScheme(ctx, projectKeyOrID, expand)
}

type internalProjectImpl struct {
	c       service.Connector
	version string
}

func (i *internalProjectImpl) Create(ctx context.Context, payload *model.ProjectPayloadScheme) (*model.NewProjectCreatedScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/project", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	project := new(model.NewProjectCreatedScheme)
	response, err := i.c.Call(request, project)
	if err != nil {
		return nil, response, err
	}

	return project, response, nil
}

func (i *internalProjectImpl) Search(ctx context.Context, options *model.ProjectSearchOptionsScheme, startAt, maxResults int) (*model.ProjectSearchScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}

		if len(options.IDs) != 0 {
			for _, id := range options.IDs {
				params.Add("id", strconv.Itoa(id))
			}
		}

		if len(options.Keys) != 0 {
			for _, key := range options.Keys {
				params.Add("keys", key)
			}
		}

		if len(options.OrderBy) != 0 {
			params.Add("orderBy", options.OrderBy)
		}

		if len(options.Query) != 0 {
			params.Add("query", options.Query)
		}

		if len(options.TypeKeys) != 0 {
			params.Add("typeKey", strings.Join(options.TypeKeys, ","))
		}

		if options.CategoryID != 0 {
			params.Add("categoryId", strconv.Itoa(options.CategoryID))
		}

		if len(options.Action) != 0 {
			params.Add("action", options.Action)
		}

		if len(options.Status) != 0 {
			params.Add("status", strings.Join(options.Status, ","))
		}

		if len(options.Properties) != 0 {
			params.Add("properties", strings.Join(options.Properties, ","))
		}
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/search?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ProjectSearchScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalProjectImpl) Get(ctx context.Context, projectKeyOrID string, expand []string) (*model.ProjectScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/project/%v", i.version, projectKeyOrID))

	if expand != nil {

		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	project := new(model.ProjectScheme)
	response, err := i.c.Call(request, project)
	if err != nil {
		return nil, response, err
	}

	return project, response, nil
}

func (i *internalProjectImpl) Update(ctx context.Context, projectKeyOrID string, payload *model.ProjectUpdateScheme) (*model.ProjectScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v", i.version, projectKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	project := new(model.ProjectScheme)
	response, err := i.c.Call(request, project)
	if err != nil {
		return nil, response, err
	}

	return project, response, nil
}

func (i *internalProjectImpl) Delete(ctx context.Context, projectKeyOrID string, enableUndo bool) (*model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, model.ErrNoProjectIDOrKey
	}

	params := url.Values{}
	params.Add("enableUndo", fmt.Sprintf("%v", enableUndo))

	endpoint := fmt.Sprintf("rest/api/%v/project/%v?%v", i.version, projectKeyOrID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalProjectImpl) DeleteAsynchronously(ctx context.Context, projectKeyOrID string) (*model.TaskScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/delete", i.version, projectKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	task := new(model.TaskScheme)
	response, err := i.c.Call(request, task)
	if err != nil {
		return nil, response, err
	}

	return task, response, nil
}

func (i *internalProjectImpl) Archive(ctx context.Context, projectKeyOrID string) (*model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, model.ErrNoProjectIDOrKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/archive", i.version, projectKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalProjectImpl) Restore(ctx context.Context, projectKeyOrID string) (*model.ProjectScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/restore", i.version, projectKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	project := new(model.ProjectScheme)
	response, err := i.c.Call(request, project)
	if err != nil {
		return nil, response, err
	}

	return project, response, nil
}

func (i *internalProjectImpl) Statuses(ctx context.Context, projectKeyOrID string) ([]*model.ProjectStatusPageScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/statuses", i.version, projectKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	var statuses []*model.ProjectStatusPageScheme
	response, err := i.c.Call(request, &statuses)
	if err != nil {
		return nil, response, err
	}

	return statuses, response, nil
}

func (i *internalProjectImpl) NotificationScheme(ctx context.Context, projectKeyOrID string, expand []string) (*model.NotificationSchemeScheme, *model.ResponseScheme, error) {

	if projectKeyOrID == "" {
		return nil, nil, model.ErrNoProjectIDOrKey
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/project/%v/notificationscheme", i.version, projectKeyOrID))

	if expand != nil {

		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), "", nil)
	if err != nil {
		return nil, nil, err
	}

	notificationScheme := new(model.NotificationSchemeScheme)
	response, err := i.c.Call(request, notificationScheme)
	if err != nil {
		return nil, response, err
	}

	return notificationScheme, response, nil
}
