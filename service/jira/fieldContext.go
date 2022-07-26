package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type FieldContext interface {

	/*
		Gets returns a paginated list of contexts for a custom field. Contexts can be returned as follows:
		With no other parameters set, all contexts.
		 1. By defining id only, all contexts from the list of IDs.
		 2. By defining isAnyIssueType
		 3. By defining isGlobalContext
		GET /rest/api/3/field/{fieldId}/context
		Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#get-custom-field-contexts
	*/
	Gets(ctx context.Context, fieldId string, options *model.FieldContextOptionsScheme, startAt, maxResults int) (
		*model.CustomFieldContextPageScheme, *model.ResponseScheme, error)

	/*
		Create creates a custom field context.
		1. If projectIds is empty, a global context is created. A global context is one that applies to all project.
		2. If issueTypeIds is empty, the context applies to all issue types.
		POST /rest/api/3/field/{fieldId}/context
		Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#create-custom-field-context
	*/
	Create(ctx context.Context, fieldId string, payload *model.FieldContextPayloadScheme) (*model.FieldContextScheme,
		*model.ResponseScheme, error)

	/*
		GetDefaultValues returns a paginated list of defaults for a custom field.
		The results can be filtered by contextId, otherwise all values are returned.
		If no defaults are set for a context, nothing is returned.
		GET /rest/api/3/field/{fieldId}/context/defaultValue
		Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#get-custom-field-contexts-default-values
	*/
	GetDefaultValues(ctx context.Context, fieldId string, contextIds []int, startAt, maxResults int) (*model.CustomFieldDefaultValuePageScheme,
		*model.ResponseScheme, error)

	/*
		SetDefaultValue sets default for contexts of a custom field.
		Default are defined using these objects:
		PUT /rest/api/3/field/{fieldId}/context/defaultValue
		Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#set-custom-field-contexts-default-values
	*/
	SetDefaultValue(ctx context.Context, fieldId string, payload *model.FieldContextDefaultPayloadScheme) (*model.ResponseScheme, error)

	/*
		IssueTypesContext returns a paginated list of context to issue type mappings for a custom field.
		Mappings are returned for all contexts or a list of contexts.
		Mappings are ordered first by context ID and then by issue type ID.
		GET /rest/api/3/field/{fieldId}/context/issuetypemapping
		// Docs: TODO: The documentation needs to be created, raise a ticket here: https://github.com/ctreminiom/go-atlassian/issues
	*/
	IssueTypesContext(ctx context.Context, fieldId string, contextIds []int, startAt, maxResults int) (*model.IssueTypeToContextMappingPageScheme,
		*model.ResponseScheme, error)

	/*
		ProjectsContext returns a paginated list of context to project mappings for a custom field.
		The result can be filtered by contextId, or otherwise all mappings are returned.
		Invalid IDs are ignored.
		GET /rest/api/3/field/{fieldId}/context/projectmapping
		// Docs: TODO: The documentation needs to be created, raise a ticket here: https://github.com/ctreminiom/go-atlassian/issues
	*/
	ProjectsContext(ctx context.Context, fieldId string, contextIds []int, startAt, maxResults int) (*model.CustomFieldContextProjectMappingPageScheme,
		*model.ResponseScheme, error)

	/*
		Update updates a custom field context
		PUT /rest/api/3/field/{fieldId}/context/{contextId}
		Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#update-custom-field-context
	*/
	Update(ctx context.Context, fieldId string, contextId int, name, description string) (*model.ResponseScheme, error)

	/*
		Delete deletes a custom field context.
		DELETE /rest/api/3/field/{fieldId}/context/{contextId}
		Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#delete-custom-field-context
	*/
	Delete(ctx context.Context, fieldId string, contextId int) (*model.ResponseScheme, error)

	/*
		AddIssueTypes adds issue types to a custom field context, appending the issue types to the issue types list.
		PUT /rest/api/3/field/{fieldId}/context/{contextId}/issuetype
		Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#add-issue-types-to-context
	*/
	AddIssueTypes(ctx context.Context, fieldId string, contextId int, issueTypesIds []string) (*model.ResponseScheme, error)

	/*
		RemoveIssueTypes removes issue types from a custom field context.
		A custom field context without any issue types applies to all issue types.
		POST /rest/api/3/field/{fieldId}/context/{contextId}/issuetype/remove
		Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#remove-issue-types-from-context
	*/
	RemoveIssueTypes(ctx context.Context, fieldId string, contextId int, issueTypesIds []string) (*model.ResponseScheme, error)

	/*
		Link assigns a custom field context to projects.
		If any project in the request is assigned to any context of the custom field, the operation fails.
		PUT /rest/api/3/field/{fieldId}/context/{contextId}/project
		Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#assign-custom-field-context-to-projects
	*/
	Link(ctx context.Context, fieldId string, contextId int, projectIds []string) (*model.ResponseScheme, error)

	/*
		UnLink removes a custom field context from projects.
		A custom field context without any projects applies to all projects.
		Removing all projects from a custom field context would result in it applying to all projects.
		POST /rest/api/3/field/{fieldId}/context/{contextId}/project/remove
		Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#remove-custom-field-context-from-projects
	*/
	UnLink(ctx context.Context, fieldId string, contextId int, projectIds []string) (*model.ResponseScheme, error)
}
