package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type ProjectConnector interface {

	// Create creates a project based on a project type template
	//
	// POST /rest/api/{2-3}/project
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#create-project
	Create(ctx context.Context, payload *model.ProjectPayloadScheme) (*model.NewProjectCreatedScheme, *model.ResponseScheme, error)

	// Search returns a paginated list of projects visible to the user.
	//
	// GET /rest/api/{2-3}/project/search
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#get-projects-paginated
	Search(ctx context.Context, options *model.ProjectSearchOptionsScheme, startAt, maxResults int) (*model.ProjectSearchScheme, *model.ResponseScheme, error)

	// Get returns the project details for a project.
	//
	// GET /rest/api/{2-3}project/{projectIdOrKey}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#get-project
	Get(ctx context.Context, projectKeyOrId string, expand []string) (*model.ProjectScheme, *model.ResponseScheme, error)

	// Update updates the project details of a project.
	//
	// PUT /rest/api/{2-3}/project/{projectIdOrKey}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#update-project
	Update(ctx context.Context, projectKeyOrId string, payload *model.ProjectUpdateScheme) (*model.ProjectScheme, *model.ResponseScheme, error)

	// Delete deletes a project.
	//
	// You can't delete a project if it's archived. To delete an archived project, restore the project and then delete it.
	//
	// To restore a project, use the Jira UI.
	//
	// DELETE /rest/api/{2-3}/project/{projectIdOrKey}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#delete-project
	Delete(ctx context.Context, projectKeyOrId string, enableUndo bool) (*model.ResponseScheme, error)

	// DeleteAsynchronously deletes a project asynchronously.
	//
	// 1. transactional, that is, if part of to delete fails the project is not deleted.
	//
	// 2. asynchronous. Follow the location link in the response to determine the status of the task and use Get task to obtain subsequent updates.
	//
	// POST /rest/api/{2-3}/project/{projectIdOrKey}/delete
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#delete-project-asynchronously
	DeleteAsynchronously(ctx context.Context, projectKeyOrId string) (*model.TaskScheme, *model.ResponseScheme, error)

	// Archive archives a project. Archived projects cannot be deleted.
	//
	// To delete an archived project, restore the project and then delete it.
	//
	// To restore a project, use the Jira UI.
	//
	// POST /rest/api/{2-3}/project/{projectIdOrKey}/archive
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#archive-project
	Archive(ctx context.Context, projectKeyOrId string) (*model.ResponseScheme, error)

	// Restore restores a project from the Jira recycle bin.
	//
	// POST /rest/api/3/project/{projectIdOrKey}/restore
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#restore-deleted-project
	Restore(ctx context.Context, projectKeyOrId string) (*model.ProjectScheme, *model.ResponseScheme, error)

	// Statuses returns the valid statuses for a project.
	//
	// The statuses are grouped by issue type, as each project has a set of valid issue types and each issue type has a set of valid statuses.
	//
	// GET /rest/api/{2-3}/project/{projectIdOrKey}/statuses
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#get-all-statuses-for-project
	Statuses(ctx context.Context, projectKeyOrId string) ([]*model.ProjectStatusPageScheme, *model.ResponseScheme, error)

	// NotificationScheme gets the notification scheme associated with the project.
	//
	// GET /rest/api/{2-3}/project/{projectKeyOrId}/notificationscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#get-project-notification-scheme
	NotificationScheme(ctx context.Context, projectKeyOrId string, expand []string) (*model.NotificationSchemeScheme, *model.ResponseScheme, error)
}

type ProjectCategoryConnector interface {

	// Gets returns all project categories.
	//
	// GET /rest/api/{2-3}/projectCategory
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/categories#get-all-project-categories
	Gets(ctx context.Context) ([]*model.ProjectCategoryScheme, *model.ResponseScheme, error)

	// Get returns a project category.
	//
	// GET /rest/api/{2-3}/projectCategory/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/categories#get-project-category-by-id
	Get(ctx context.Context, categoryId int) (*model.ProjectCategoryScheme, *model.ResponseScheme, error)

	// Create creates a project category.
	//
	// POST /rest/api/{2-3}/projectCategory
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/categories#create-project-category
	Create(ctx context.Context, payload *model.ProjectCategoryPayloadScheme) (*model.ProjectCategoryScheme, *model.ResponseScheme, error)

	// Update updates a project category.
	//
	// PUT /rest/api/{2-3}/projectCategory/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/categories#update-project-category
	Update(ctx context.Context, categoryId int, payload *model.ProjectCategoryPayloadScheme) (*model.ProjectCategoryScheme, *model.ResponseScheme, error)

	// Delete deletes a project category.
	//
	// DELETE /rest/api/{2-3}/projectCategory/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/categories#delete-project-category
	Delete(ctx context.Context, categoryId int) (*model.ResponseScheme, error)
}
