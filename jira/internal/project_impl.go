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
	"strings"
)

type ProjectChildServices struct {
	Category   *ProjectCategoryService
	Component  *ProjectComponentService
	Feature    *ProjectFeatureService
	Permission *ProjectPermissionSchemeService
}

func NewProjectService(client service.Client, version string, subServices *ProjectChildServices) (*ProjectService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &ProjectService{
		internalClient: &internalProjectImpl{c: client, version: version},
		Category:       subServices.Category,
		Component:      subServices.Component,
		Feature:        subServices.Feature,
		Permission:     subServices.Permission,
	}, nil
}

type ProjectService struct {
	internalClient jira.ProjectConnector
	Category       *ProjectCategoryService
	Component      *ProjectComponentService
	Feature        *ProjectFeatureService
	Permission     *ProjectPermissionSchemeService
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
// GET /rest/api/{2-3}project/{projectIdOrKey}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#get-project
func (p *ProjectService) Get(ctx context.Context, projectKeyOrId string, expand []string) (*model.ProjectScheme, *model.ResponseScheme, error) {
	return p.internalClient.Get(ctx, projectKeyOrId, expand)
}

// Update updates the project details of a project.
//
// PUT /rest/api/{2-3}/project/{projectIdOrKey}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#update-project
func (p *ProjectService) Update(ctx context.Context, projectKeyOrId string, payload *model.ProjectUpdateScheme) (*model.ProjectScheme, *model.ResponseScheme, error) {
	return p.internalClient.Update(ctx, projectKeyOrId, payload)
}

// Delete deletes a project.
//
// You can't delete a project if it's archived. To delete an archived project, restore the project and then delete it.
//
// To restore a project, use the Jira UI.
//
// DELETE /rest/api/{2-3}/project/{projectIdOrKey}
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#delete-project
func (p *ProjectService) Delete(ctx context.Context, projectKeyOrId string, enableUndo bool) (*model.ResponseScheme, error) {
	return p.internalClient.Delete(ctx, projectKeyOrId, enableUndo)
}

// DeleteAsynchronously deletes a project asynchronously.
//
// 1. transactional, that is, if part of to delete fails the project is not deleted.
//
// 2. asynchronous. Follow the location link in the response to determine the status of the task and use Get task to obtain subsequent updates.
//
// POST /rest/api/{2-3}/project/{projectIdOrKey}/delete
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#delete-project-asynchronously
func (p *ProjectService) DeleteAsynchronously(ctx context.Context, projectKeyOrId string) (*model.TaskScheme, *model.ResponseScheme, error) {
	return p.internalClient.DeleteAsynchronously(ctx, projectKeyOrId)
}

// Archive archives a project. Archived projects cannot be deleted.
//
// To delete an archived project, restore the project and then delete it.
//
// To restore a project, use the Jira UI.
//
// POST /rest/api/{2-3}/project/{projectIdOrKey}/archive
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#archive-project
func (p *ProjectService) Archive(ctx context.Context, projectKeyOrId string) (*model.ResponseScheme, error) {
	return p.internalClient.Archive(ctx, projectKeyOrId)
}

// Restore restores a project from the Jira recycle bin.
//
// POST /rest/api/3/project/{projectIdOrKey}/restore
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#restore-deleted-project
func (p *ProjectService) Restore(ctx context.Context, projectKeyOrId string) (*model.ProjectScheme, *model.ResponseScheme, error) {
	return p.internalClient.Restore(ctx, projectKeyOrId)
}

// Statuses returns the valid statuses for a project.
//
// The statuses are grouped by issue type, as each project has a set of valid issue types and each issue type has a set of valid statuses.
//
// GET /rest/api/{2-3}/project/{projectIdOrKey}/statuses
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#get-all-statuses-for-project
func (p *ProjectService) Statuses(ctx context.Context, projectKeyOrId string) ([]*model.ProjectStatusPageScheme, *model.ResponseScheme, error) {
	return p.internalClient.Statuses(ctx, projectKeyOrId)
}

// NotificationScheme gets the notification scheme associated with the project.
//
// GET /rest/api/{2-3}/project/{projectKeyOrId}/notificationscheme
//
// https://docs.go-atlassian.io/jira-software-cloud/projects#get-project-notification-scheme
func (p *ProjectService) NotificationScheme(ctx context.Context, projectKeyOrId string, expand []string) (*model.NotificationSchemeScheme, *model.ResponseScheme, error) {
	return p.internalClient.NotificationScheme(ctx, projectKeyOrId, expand)
}

type internalProjectImpl struct {
	c       service.Client
	version string
}

func (i *internalProjectImpl) Create(ctx context.Context, payload *model.ProjectPayloadScheme) (*model.NewProjectCreatedScheme, *model.ResponseScheme, error) {

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/project", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
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

	endpoint := fmt.Sprintf("rest/api/%v/project/search?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

func (i *internalProjectImpl) Get(ctx context.Context, projectKeyOrId string, expand []string) (*model.ProjectScheme, *model.ResponseScheme, error) {

	if projectKeyOrId == "" {
		return nil, nil, model.ErrNoProjectIDOrKeyError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/project/%v", i.version, projectKeyOrId))

	if expand != nil {

		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
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

func (i *internalProjectImpl) Update(ctx context.Context, projectKeyOrId string, payload *model.ProjectUpdateScheme) (*model.ProjectScheme, *model.ResponseScheme, error) {

	if projectKeyOrId == "" {
		return nil, nil, model.ErrNoProjectIDOrKeyError
	}

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v", i.version, projectKeyOrId)

	request, err := i.c.NewRequest(ctx, http.MethodPut, endpoint, reader)
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

func (i *internalProjectImpl) Delete(ctx context.Context, projectKeyOrId string, enableUndo bool) (*model.ResponseScheme, error) {

	if projectKeyOrId == "" {
		return nil, model.ErrNoProjectIDOrKeyError
	}

	params := url.Values{}
	params.Add("enableUndo", fmt.Sprintf("%v", enableUndo))

	endpoint := fmt.Sprintf("rest/api/%v/project/%v?%v", i.version, projectKeyOrId, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalProjectImpl) DeleteAsynchronously(ctx context.Context, projectKeyOrId string) (*model.TaskScheme, *model.ResponseScheme, error) {

	if projectKeyOrId == "" {
		return nil, nil, model.ErrNoProjectIDOrKeyError
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/delete", i.version, projectKeyOrId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, nil)
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

func (i *internalProjectImpl) Archive(ctx context.Context, projectKeyOrId string) (*model.ResponseScheme, error) {

	if projectKeyOrId == "" {
		return nil, model.ErrNoProjectIDOrKeyError
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/archive", i.version, projectKeyOrId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalProjectImpl) Restore(ctx context.Context, projectKeyOrId string) (*model.ProjectScheme, *model.ResponseScheme, error) {

	if projectKeyOrId == "" {
		return nil, nil, model.ErrNoProjectIDOrKeyError
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/restore", i.version, projectKeyOrId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, nil)
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

func (i *internalProjectImpl) Statuses(ctx context.Context, projectKeyOrId string) ([]*model.ProjectStatusPageScheme, *model.ResponseScheme, error) {

	if projectKeyOrId == "" {
		return nil, nil, model.ErrNoProjectIDOrKeyError
	}

	endpoint := fmt.Sprintf("rest/api/%v/project/%v/statuses", i.version, projectKeyOrId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

func (i *internalProjectImpl) NotificationScheme(ctx context.Context, projectKeyOrId string, expand []string) (*model.NotificationSchemeScheme, *model.ResponseScheme, error) {

	if projectKeyOrId == "" {
		return nil, nil, model.ErrNoProjectIDOrKeyError
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/%v/project/%v/notificationscheme", i.version, projectKeyOrId))

	if expand != nil {

		params := url.Values{}
		params.Add("expand", strings.Join(expand, ","))

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint.String(), nil)
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
