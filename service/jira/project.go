package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
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
	// GET /rest/api/{2-3}project/{projectKeyOrID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#get-project
	Get(ctx context.Context, projectKeyOrID string, expand []string) (*model.ProjectScheme, *model.ResponseScheme, error)

	// Update updates the project details of a project.
	//
	// PUT /rest/api/{2-3}/project/{projectKeyOrID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#update-project
	Update(ctx context.Context, projectKeyOrID string, payload *model.ProjectUpdateScheme) (*model.ProjectScheme, *model.ResponseScheme, error)

	// Delete deletes a project.
	//
	// You can't delete a project if it's archived. To delete an archived project, restore the project and then delete it.
	//
	// To restore a project, use the Jira UI.
	//
	// DELETE /rest/api/{2-3}/project/{projectKeyOrID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#delete-project
	Delete(ctx context.Context, projectKeyOrID string, enableUndo bool) (*model.ResponseScheme, error)

	// DeleteAsynchronously deletes a project asynchronously.
	//
	// 1. transactional, that is, if part of to delete fails the project is not deleted.
	//
	// 2. asynchronous. Follow the location link in the response to determine the status of the task and use Get task to obtain subsequent updates.
	//
	// POST /rest/api/{2-3}/project/{projectKeyOrID}/delete
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#delete-project-asynchronously
	DeleteAsynchronously(ctx context.Context, projectKeyOrID string) (*model.TaskScheme, *model.ResponseScheme, error)

	// Archive archives a project. Archived projects cannot be deleted.
	//
	// To delete an archived project, restore the project and then delete it.
	//
	// To restore a project, use the Jira UI.
	//
	// POST /rest/api/{2-3}/project/{projectKeyOrID}/archive
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#archive-project
	Archive(ctx context.Context, projectKeyOrID string) (*model.ResponseScheme, error)

	// Restore restores a project from the Jira recycle bin.
	//
	// POST /rest/api/3/project/{projectKeyOrID}/restore
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#restore-deleted-project
	Restore(ctx context.Context, projectKeyOrID string) (*model.ProjectScheme, *model.ResponseScheme, error)

	// Statuses returns the valid statuses for a project.
	//
	// The statuses are grouped by issue type, as each project has a set of valid issue types and each issue type has a set of valid statuses.
	//
	// GET /rest/api/{2-3}/project/{projectKeyOrID}/statuses
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#get-all-statuses-for-project
	Statuses(ctx context.Context, projectKeyOrID string) ([]*model.ProjectStatusPageScheme, *model.ResponseScheme, error)

	// NotificationScheme gets the notification scheme associated with the project.
	//
	// GET /rest/api/{2-3}/project/{projectKeyOrID}/notificationscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects#get-project-notification-scheme
	NotificationScheme(ctx context.Context, projectKeyOrID string, expand []string) (*model.NotificationSchemeScheme, *model.ResponseScheme, error)
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
	Get(ctx context.Context, categoryID int) (*model.ProjectCategoryScheme, *model.ResponseScheme, error)

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
	Update(ctx context.Context, categoryID int, payload *model.ProjectCategoryPayloadScheme) (*model.ProjectCategoryScheme, *model.ResponseScheme, error)

	// Delete deletes a project category.
	//
	// DELETE /rest/api/{2-3}/projectCategory/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/categories#delete-project-category
	Delete(ctx context.Context, categoryID int) (*model.ResponseScheme, error)
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
	// GET /rest/api/{2-3}/project/{projectKeyOrID}/components
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/components#get-project-components
	Gets(ctx context.Context, projectKeyOrID string) ([]*model.ComponentScheme, *model.ResponseScheme, error)

	// Count returns the counts of issues assigned to the component.
	//
	// GET /rest/api/{2-3}/component/{id}/relatedIssueCounts
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/components#get-component-issues-count
	Count(ctx context.Context, componentID string) (*model.ComponentCountScheme, *model.ResponseScheme, error)

	// Delete deletes a component.
	//
	// DELETE /rest/api/{2-3}/component/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/components#delete-component
	Delete(ctx context.Context, componentID string) (*model.ResponseScheme, error)

	// Update updates a component.
	//
	// Any fields included in the request are overwritten
	//
	// PUT /rest/api/{2-3}/component/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/components#update-component
	Update(ctx context.Context, componentID string, payload *model.ComponentPayloadScheme) (*model.ComponentScheme, *model.ResponseScheme, error)

	// Get returns a component.
	//
	// GET /rest/api/{2-3}/component/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/components#get-component
	Get(ctx context.Context, componentID string) (*model.ComponentScheme, *model.ResponseScheme, error)
}

type ProjectFeatureConnector interface {

	// Gets returns the list of features for a project.
	//
	// GET /rest/api/{2-3}/project/{projectKeyOrID}/features
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/features#get-project-features
	Gets(ctx context.Context, projectKeyOrID string) (*model.ProjectFeaturesScheme, *model.ResponseScheme, error)

	// Set sets the state of a project feature.
	//
	// PUT /rest/api/{2-3}/project/{projectKeyOrID}/features/{featureKey}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/features#set-project-feature-state
	Set(ctx context.Context, projectKeyOrID, featureKey, state string) (*model.ProjectFeaturesScheme, *model.ResponseScheme, error)
}

type ProjectPermissionSchemeConnector interface {

	// Get search the permission scheme associated with the project.
	//
	// GET /rest/api/{2-3}/project/{projectKeyOrID}/permissionscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/permission-schemes#get-assigned-permission-scheme
	Get(ctx context.Context, projectKeyOrID string, expand []string) (*model.PermissionSchemeScheme, *model.ResponseScheme, error)

	// Assign assigns a permission scheme with a project.
	//
	// See Managing project permissions for more information about permission schemes.
	//
	// PUT /rest/api/{2-3}/project/{projectKeyOrID}/permissionscheme
	Assign(ctx context.Context, projectKeyOrID string, permissionSchemeID int) (*model.PermissionSchemeScheme, *model.ResponseScheme, error)

	// SecurityLevels returns all issue security levels for the project that the user has access to.
	//
	// GET /rest/api/{2-3}/project/{projectKeyOrID}/securitylevel
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/permission-schemes#get-project-issue-security-levels
	SecurityLevels(ctx context.Context, projectKeyOrID string) (*model.IssueSecurityLevelsScheme, *model.ResponseScheme, error)
}

type ProjectPropertyConnector interface {

	// Gets returns all project property keys for the project.
	//
	// GET /rest/api/{2-3}/project/{projectKeyOrID}/properties
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/properties#get-project-properties-keys
	Gets(ctx context.Context, projectKeyOrID string) (*model.PropertyPageScheme, *model.ResponseScheme, error)

	// Get returns the value of a project property.
	//
	// GET /rest/api/{2-3}/project/{projectKeyOrID}/properties/{propertyKey}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/properties#get-project-property
	Get(ctx context.Context, projectKeyOrID, propertyKey string) (*model.EntityPropertyScheme, *model.ResponseScheme, error)

	// Set sets the value of the project property.
	//
	// You can use project properties to store custom data against the project.
	//
	// The value of the request body must be a valid, non-empty JSON blob.
	//
	// The maximum length is 32768 characters.
	//
	// PUT /rest/api/{2-3}/project/{projectKeyOrID}/properties/{propertyKey}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/properties#set-project-property
	Set(ctx context.Context, projectKeyOrID, propertyKey string, payload interface{}) (*model.ResponseScheme, error)

	// Delete deletes the property from a project.
	//
	// DELETE /rest/api/{2-3}/project/{projectKeyOrID}/properties/{propertyKey}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/properties#delete-project-property
	Delete(ctx context.Context, projectKeyOrID, propertyKey string) (*model.ResponseScheme, error)
}

type ProjectRoleConnector interface {

	// Gets returns a list of project roles for the project returning the name and self URL for each role.
	//
	// GET /rest/api/{2-3}/project/{projectKeyOrID}/role
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-project-roles-for-project
	Gets(ctx context.Context, projectKeyOrID string) (*map[string]int, *model.ResponseScheme, error)

	// Get returns a project role's details and actors associated with the project.
	//
	// GET /rest/api/{2-3}/project/{projectKeyOrID}/role/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-project-role-for-project
	Get(ctx context.Context, projectKeyOrID string, roleID int) (*model.ProjectRoleScheme, *model.ResponseScheme, error)

	// Details returns all project roles and the details for each role.
	//
	// GET /rest/api/{2-3}/project/{projectKeyOrID}/roledetails
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/roles#get-project-role-details
	Details(ctx context.Context, projectKeyOrID string) ([]*model.ProjectRoleDetailScheme, *model.ResponseScheme, error)

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
	// POST /rest/api/{2-3}/project/{projectKeyOrID}/role/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/roles/actors#add-actors-to-project-role
	Add(ctx context.Context, projectKeyOrID string, roleID int, accountIDs, groups []string) (*model.ProjectRoleScheme, *model.ResponseScheme, error)

	// Delete deletes actors from a project role for the project.
	//
	// DELETE /rest/api/{2-3}/project/{projectKeyOrID}/role/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/roles/actors#delete-actors-from-project-role
	Delete(ctx context.Context, projectKeyOrID string, roleID int, accountID, group string) (*model.ResponseScheme, error)
}

type ProjectTypeConnector interface {

	// Gets returns all project types, whether the instance has a valid license for each type.
	//
	// GET /rest/api/{2-3}/project/type
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/types#get-all-project-types
	Gets(ctx context.Context) ([]*model.ProjectTypeScheme, *model.ResponseScheme, error)

	// Licensed returns all project types with a valid license.
	//
	// GET /rest/api/{2-3}/project/type/accessible
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/types#get-licensed-project-types
	Licensed(ctx context.Context) ([]*model.ProjectTypeScheme, *model.ResponseScheme, error)

	// Get returns a project type
	//
	// GET /rest/api/{2-3}/project/type/{projectTypeKey}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/types#get-project-type-by-key
	Get(ctx context.Context, projectTypeKey string) (*model.ProjectTypeScheme, *model.ResponseScheme, error)

	// Accessible returns a project type if it is accessible to the user.
	//
	// GET /rest/api/{2-3}/project/type/{projectTypeKey}/accessible
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/types#get-accessible-project-type-by-key
	Accessible(ctx context.Context, projectTypeKey string) (*model.ProjectTypeScheme, *model.ResponseScheme, error)
}

type ProjectValidatorConnector interface {

	// Validate validates a project key by confirming the key is a valid string and not in use.
	//
	// GET /rest/api/{2-3}/projectvalidate/key
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/validation#validate-project-key
	Validate(ctx context.Context, key string) (*model.ProjectValidationMessageScheme, *model.ResponseScheme, error)

	// Key validates a project key and, if the key is invalid or in use,
	//
	// generates a valid random string for the project key.
	//
	// GET /rest/api/{2-3}/projectvalidate/validProjectKey
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/validation#get-valid-project-key
	Key(ctx context.Context, key string) (string, *model.ResponseScheme, error)

	// Name checks that a project name isn't in use.
	//
	// If the name isn't in use, the passed string is returned.
	//
	// If the name is in use, this operation attempts to generate a valid project name based on the one supplied,
	//
	// usually by adding a sequence number. If a valid project name cannot be generated, a 404 response is returned.
	//
	// GET /rest/api/{2-3}/projectvalidate/validProjectName
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/validation#get-valid-project-name
	Name(ctx context.Context, name string) (string, *model.ResponseScheme, error)
}

type ProjectVersionConnector interface {

	// Gets returns all versions in a project.
	//
	// The response is not paginated.
	//
	// Use Search() if you want to get the versions in a project with pagination.
	//
	// GET /rest/api/{2-3}/project/{projectKeyOrID}/versions
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-project-versions
	Gets(ctx context.Context, projectKeyOrID string) ([]*model.VersionScheme, *model.ResponseScheme, error)

	// Search returns a paginated list of all versions in a project.
	//
	// GET /rest/api/{2-3}/project/{projectKeyOrID}/version
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-project-versions-paginated
	Search(ctx context.Context, projectKeyOrID string, options *model.VersionGetsOptions, startAt, maxResults int) (*model.VersionPageScheme, *model.ResponseScheme, error)

	// Create creates a project version.
	//
	// POST /rest/api/{2-3}/version
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#create-version
	Create(ctx context.Context, payload *model.VersionPayloadScheme) (*model.VersionScheme, *model.ResponseScheme, error)

	// Get returns a project version.
	//
	// GET /rest/api/{2-3}/version/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-version
	Get(ctx context.Context, versionID string, expand []string) (*model.VersionScheme, *model.ResponseScheme, error)

	// Update updates a project version.
	//
	// PUT /rest/api/{2-3}/version/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#update-version
	Update(ctx context.Context, versionID string, payload *model.VersionPayloadScheme) (*model.VersionScheme, *model.ResponseScheme, error)

	// Merge merges two project versions.
	//
	// The merge is completed by deleting the version specified in id and replacing any occurrences of
	//
	// its ID in fixVersion with the version ID specified in moveIssuesTo.
	//
	// PUT /rest/api/{2-3}/version/{id}/mergeto/{moveIssuesTo}
	Merge(ctx context.Context, versionID, versionMoveIssuesTo string) (*model.ResponseScheme, error)

	// RelatedIssueCounts returns the following counts for a version:
	//
	// 1. Number of issues where the fixVersion is set to the version.
	//
	// 2. Number of issues where the affectedVersion is set to the version.
	//
	// 3. Number of issues where a version custom field is set to the version.
	//
	// GET /rest/api/{2-3}/version/{id}/relatedIssueCounts
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-versions-related-issues-count
	RelatedIssueCounts(ctx context.Context, versionID string) (*model.VersionIssueCountsScheme, *model.ResponseScheme, error)

	// UnresolvedIssueCount returns counts of the issues and unresolved issues for the project version.
	//
	// GET /rest/api/{2-3}/version/{id}/unresolvedIssueCount
	//
	// https://docs.go-atlassian.io/jira-software-cloud/projects/versions#get-versions-unresolved-issues-count
	UnresolvedIssueCount(ctx context.Context, versionID string) (*model.VersionUnresolvedIssuesCountScheme, *model.ResponseScheme, error)
}
