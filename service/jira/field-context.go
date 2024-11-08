package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

//FieldContextOptionConnector

// FieldContextConnector is the interface that wraps the Jira field context
//
// It contains the methods required to manipulate the field context associated with a Jira field, you can use to:
//  1. get, create, update, and delete custom field contexts.
//  2. get context to issue types and projects mappings.
//  3. get custom field contexts for projects and issue types.
//  4. assign custom field contexts to projects.
//  5. remove custom field contexts from projects.
//  6. add issue types to custom field contexts.
type FieldContextConnector interface {

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
	Gets(ctx context.Context, fieldID string, options *model.FieldContextOptionsScheme, startAt, maxResults int) (
		*model.CustomFieldContextPageScheme, *model.ResponseScheme, error)

	// Create creates a custom field context.
	//
	// 1. If projectIDs is empty, a global context is created. A global context is one that applies to all project.
	//
	// 2. If issueTypeIDs is empty, the context applies to all issue types.
	//
	// POST /rest/api/{2-3}/field/{fieldID}/context
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#create-custom-field-context
	Create(ctx context.Context, fieldID string, payload *model.FieldContextPayloadScheme) (*model.FieldContextScheme,
		*model.ResponseScheme, error)

	// GetDefaultValues returns a paginated list of defaults for a custom field.
	//
	// The results can be filtered by contextID, otherwise all values are returned. If no defaults are set for a context, nothing is returned.
	//
	// GET /rest/api/{2-3}/field/{fieldID}/context/defaultValue
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#get-custom-field-contexts-default-values
	GetDefaultValues(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (*model.CustomFieldDefaultValuePageScheme,
		*model.ResponseScheme, error)

	// SetDefaultValue sets default for contexts of a custom field.
	//
	// PUT /rest/api/{2-3}/field/{fieldID}/context/defaultValue
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#set-custom-field-contexts-default-values
	SetDefaultValue(ctx context.Context, fieldID string, payload *model.FieldContextDefaultPayloadScheme) (*model.ResponseScheme, error)

	// IssueTypesContext returns a paginated list of context to issue type mappings for a custom field.
	//
	// 1. Mappings are returned for all contexts or a list of contexts.
	//
	// 2. Mappings are ordered first by context ID and then by issue type ID.
	//
	// GET /rest/api/{2-3}/field/{fieldID}/context/issuetypemapping
	//
	// Docs: TODO: The documentation needs to be created, raise a ticket here: https://github.com/ctreminiom/go-atlassian/issues
	IssueTypesContext(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (*model.IssueTypeToContextMappingPageScheme,
		*model.ResponseScheme, error)

	// ProjectsContext returns a paginated list of context to project mappings for a custom field.
	//
	// 1. The result can be filtered by contextID, or otherwise all mappings are returned.
	//
	// 2. Invalid IDs are ignored.
	//
	// GET /rest/api/{2-3}/field/{fieldID}/context/projectmapping
	//
	// Docs: TODO: The documentation needs to be created, raise a ticket here: https://github.com/ctreminiom/go-atlassian/issues
	ProjectsContext(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (*model.CustomFieldContextProjectMappingPageScheme,
		*model.ResponseScheme, error)

	// Update updates a custom field context
	//
	// PUT /rest/api/{2-3}/field/{fieldID}/context/{contextID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#update-custom-field-context
	Update(ctx context.Context, fieldID string, contextID int, name, description string) (*model.ResponseScheme, error)

	// Delete deletes a custom field context.
	//
	// DELETE /rest/api/{2-3}/field/{fieldID}/context/{contextID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#delete-custom-field-context
	Delete(ctx context.Context, fieldID string, contextID int) (*model.ResponseScheme, error)

	// AddIssueTypes adds issue types to a custom field context, appending the issue types to the issue types list.
	//
	// PUT /rest/api/{2-3}/field/{fieldID}/context/{contextID}/issuetype
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#add-issue-types-to-context
	AddIssueTypes(ctx context.Context, fieldID string, contextID int, issueTypesIDs []string) (*model.ResponseScheme, error)

	// RemoveIssueTypes removes issue types from a custom field context. A custom field context without any issue types applies to all issue types.
	//
	// POST /rest/api/{2-3}/field/{fieldID}/context/{contextID}/issuetype/remove
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#remove-issue-types-from-context
	RemoveIssueTypes(ctx context.Context, fieldID string, contextID int, issueTypesIDs []string) (*model.ResponseScheme, error)

	// Link assigns a custom field context to projects. If any project in the request is assigned to any context of the custom field, the operation fails.
	//
	// PUT /rest/api/{2-3}/field/{fieldID}/context/{contextID}/project
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#assign-custom-field-context-to-projects
	Link(ctx context.Context, fieldID string, contextID int, projectIDs []string) (*model.ResponseScheme, error)

	// UnLink removes a custom field context from projects.
	//
	// 1. A custom field context without any projects applies to all projects.
	//
	// 2. Removing all projects from a custom field context would result in it applying to all projects.
	//
	// POST /rest/api/{2-3}/field/{fieldID}/context/{contextID}/project/remove
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#remove-custom-field-context-from-projects
	UnLink(ctx context.Context, fieldID string, contextID int, projectIDs []string) (*model.ResponseScheme, error)
}

// FieldContextOptionConnector is the interface that wraps the Jira field context options
//
// It contains the methods required to manipulate the field options associated with a field context
// and represents custom issue field select list options created in Jira or using the REST API.
// Use it to retrieve, create, update, order, and delete custom field options.
type FieldContextOptionConnector interface {

	// Gets returns a paginated list of all custom field option for a context.
	//
	// Options are returned first then cascading options, in the order they display in Jira.
	//
	// GET /rest/api/{2-3}/field/{fieldID}/context/{contextID}/option
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#get-custom-field-options
	Gets(ctx context.Context, fieldID string, contextID int, options *model.FieldOptionContextParams, startAt, maxResults int) (*model.CustomFieldContextOptionPageScheme, *model.ResponseScheme, error)

	// Create creates options and, where the custom select field is of the type Select List (cascading), cascading options for a custom select field.
	//
	// 1. The options are added to a context of the field.
	//
	// 2. The maximum number of options that can be created per request is 1000 and each field can have a maximum of 10000 options.
	//
	// POST /rest/api/{2-3}/field/{fieldID}/context/{contextID}/option
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#create-custom-field-options
	Create(ctx context.Context, fieldID string, contextID int, payload *model.FieldContextOptionListScheme) (*model.FieldContextOptionListScheme, *model.ResponseScheme, error)

	// Update updates the options of a custom field.
	//
	// 1. If any of the options are not found, no options are updated.
	//
	// 2. Options where the values in the request match the current values aren't updated and aren't reported in the response.
	//
	// PUT /rest/api/{2-3}/field/{fieldID}/context/{contextID}/option
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#update-custom-field-options
	Update(ctx context.Context, fieldID string, contextID int, payload *model.FieldContextOptionListScheme) (*model.FieldContextOptionListScheme, *model.ResponseScheme, error)

	// Delete deletes a custom field option.
	//
	// 1. Options with cascading options cannot be deleted without deleting the cascading options first.
	//
	// DELETE /rest/api/{2-3}/field/{fieldID}/context/{contextID}/option/{optionID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#delete-custom-field-options
	Delete(ctx context.Context, fieldID string, contextID, optionID int) (*model.ResponseScheme, error)

	// Order changes the order of custom field options or cascading options in a context.
	//
	// PUT /rest/api/{2-3}/field/{fieldID}/context/{contextID}/option/move
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#reorder-custom-field-options
	Order(ctx context.Context, fieldID string, contextID int, payload *model.OrderFieldOptionPayloadScheme) (*model.ResponseScheme, error)
}
