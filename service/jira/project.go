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

type ProjectComponentConnector interface {

	// Create creates a component. Use components to provide containers for issues within a project.
	//
	// POST /rest/api/{2-3}/component
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/components#create-component
	Create(ctx context.Context, payload *model.ComponentPayloadScheme) (*model.ComponentScheme, *model.ResponseScheme, error)

	// Gets returns all components in a project.
	//
	// GET /rest/api/{2-3}/project/{projectIdOrKey}/components
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/components#get-project-components
	Gets(ctx context.Context, projectIdOrKey string) ([]*model.ComponentScheme, *model.ResponseScheme, error)

	// Count returns the counts of issues assigned to the component.
	//
	// GET /rest/api/{2-3}/component/{id}/relatedIssueCounts
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/components#get-component-issues-count
	Count(ctx context.Context, componentId string) (*model.ComponentCountScheme, *model.ResponseScheme, error)

	// Delete deletes a component.
	//
	// DELETE /rest/api/{2-3}/component/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/components#delete-component
	Delete(ctx context.Context, componentId string) (*model.ResponseScheme, error)

	// Update updates a component.
	//
	// Any fields included in the request are overwritten
	//
	// PUT /rest/api/{2-3}/component/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/components#update-component
	Update(ctx context.Context, componentId string, payload *model.ComponentPayloadScheme) (*model.ComponentScheme, *model.ResponseScheme, error)

	// Get returns a component.
	//
	// GET /rest/api/{2-3}/component/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/components#get-component
	Get(ctx context.Context, componentId string) (*model.ComponentScheme, *model.ResponseScheme, error)
}

type ProjectFeatureConnector interface {

	// Gets returns the list of features for a project.
	//
	// GET /rest/api/{2-3}/project/{projectIdOrKey}/features
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/features#get-project-features
	Gets(ctx context.Context, projectKeyOrId string) (*model.ProjectFeaturesScheme, *model.ResponseScheme, error)

	// Set sets the state of a project feature.
	//
	// PUT /rest/api/{2-3}/project/{projectIdOrKey}/features/{featureKey}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/features#set-project-feature-state
	Set(ctx context.Context, projectKeyOrId, featureKey, state string) (*model.ProjectFeaturesScheme, *model.ResponseScheme, error)
}

type ProjectPermissionSchemeConnector interface {

	// Get search the permission scheme associated with the project.
	//
	// GET /rest/api/{2-3}/project/{projectKeyOrId}/permissionscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/permission-schemes#get-assigned-permission-scheme
	Get(ctx context.Context, projectKeyOrId string, expand []string) (*model.PermissionSchemeScheme, *model.ResponseScheme, error)

	// Assign assigns a permission scheme with a project.
	//
	// See Managing project permissions for more information about permission schemes.
	//
	// PUT /rest/api/{2-3}/project/{projectKeyOrId}/permissionscheme
	Assign(ctx context.Context, projectKeyOrId string, permissionSchemeId int) (*model.PermissionSchemeScheme, *model.ResponseScheme, error)

	// SecurityLevels returns all issue security levels for the project that the user has access to.
	//
	// GET /rest/api/{2-3}/project/{projectKeyOrId}/securitylevel
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/permission-schemes#get-project-issue-security-levels
	SecurityLevels(ctx context.Context, projectKeyOrId string) (*model.IssueSecurityLevelsScheme, *model.ResponseScheme, error)
}

type ProjectPropertyConnector interface {

	// Gets returns all project property keys for the project.
	//
	// GET /rest/api/{2-3}/project/{projectIdOrKey}/properties
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/properties#get-project-properties-keys
	Gets(ctx context.Context, projectKeyOrId string) (*model.ProjectPropertyPageScheme, *model.ResponseScheme, error)

	// Get returns the value of a project property.
	//
	// GET /rest/api/{2-3}/project/{projectIdOrKey}/properties/{propertyKey}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/properties#get-project-property
	Get(ctx context.Context, projectKeyOrId, propertyKey string) (*model.EntityPropertyScheme, *model.ResponseScheme, error)

	// Set sets the value of the project property.
	//
	// You can use project properties to store custom data against the project.
	//
	// The value of the request body must be a valid, non-empty JSON blob.
	//
	// The maximum length is 32768 characters.
	//
	// PUT /rest/api/{2-3}/project/{projectIdOrKey}/properties/{propertyKey}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/properties#set-project-property
	Set(ctx context.Context, projectKeyOrId, propertyKey string, payload interface{}) (*model.ResponseScheme, error)

	// Delete deletes the property from a project.
	//
	// DELETE /rest/api/{2-3}/project/{projectIdOrKey}/properties/{propertyKey}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/properties#delete-project-property
	Delete(ctx context.Context, projectKeyOrId, propertyKey string) (*model.ResponseScheme, error)
}

type ProjectRoleConnector interface {

	// Gets returns a list of project roles for the project returning the name and self URL for each role.
	//
	// GET /rest/api/{2-3}/project/{projectIdOrKey}/role
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-project-roles-for-project
	Gets(ctx context.Context, projectKeyOrId string) (*map[string]int, *model.ResponseScheme, error)

	// Get returns a project role's details and actors associated with the project.
	//
	// GET /rest/api/{2-3}/project/{projectIdOrKey}/role/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-project-role-for-project
	Get(ctx context.Context, projectKeyOrId string, roleId int) (*model.ProjectRoleScheme, *model.ResponseScheme, error)

	// Details returns all project roles and the details for each role.
	//
	// GET /rest/api/{2-3}/project/{projectIdOrKey}/roledetails
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-project-role-details
	Details(ctx context.Context, projectKeyOrId string) ([]*model.ProjectRoleDetailScheme, *model.ResponseScheme, error)

	// Global gets a list of all project roles, complete with project role details and default actors.
	//
	// GET /rest/api/{2-3}/role
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-all-project-roles
	Global(ctx context.Context) ([]*model.ProjectRoleScheme, *model.ResponseScheme, error)

	// Create creates a new project role with no default actors.
	//
	// POST /rest/api/{2-3}/role
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/roles#create-project-role
	Create(ctx context.Context, payload *model.ProjectRolePayloadScheme) (*model.ProjectRoleScheme, *model.ResponseScheme, error)
}

type ProjectRoleActorConnector interface {

	// Add adds actors to a project role for the project.
	//
	// POST /rest/api/{2-3}/project/{projectIdOrKey}/role/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/roles/actors#add-actors-to-project-role
	Add(ctx context.Context, projectKeyOrId string, roleId int, accountIds, groups []string) (*model.ProjectRoleScheme, *model.ResponseScheme, error)

	// Delete deletes actors from a project role for the project.
	//
	// DELETE /rest/api/{2-3}/project/{projectIdOrKey}/role/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/roles/actors#delete-actors-from-project-role
	Delete(ctx context.Context, projectKeyOrId string, roleId int, accountId, group string) (*model.ResponseScheme, error)
}