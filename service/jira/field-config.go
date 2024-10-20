package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// FieldConfigConnector interface holds the methods available for the FieldConfig resource.
type FieldConfigConnector interface {

	// Gets Returns a paginated list of all field configurations.
	//
	// GET /rest/api/{2-3}/fieldconfiguration
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-all-field-configurations
	Gets(ctx context.Context, ids []int, isDefault bool, startAt, maxResults int) (*model.FieldConfigurationPageScheme,
		*model.ResponseScheme, error)

	// Create creates a field configuration. The field configuration is created with the same field properties as the
	// default configuration, with all the fields being optional.
	//
	// This operation can only create configurations for use in company-managed (classic) projects.
	//
	// POST /rest/api/{2-3}/fieldconfiguration
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#create-field-configuration
	Create(ctx context.Context, name, description string) (*model.FieldConfigurationScheme, *model.ResponseScheme, error)

	// Update updates a field configuration. The name and the description provided in the request override the existing values.
	//
	// This operation can only update configurations used in company-managed (classic) projects.
	//
	// PUT /rest/api/{2-3}/fieldconfiguration/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#update-field-configuration
	Update(ctx context.Context, id int, name, description string) (*model.ResponseScheme, error)

	// Delete deletes a field configuration.
	//
	// This operation can only delete configurations used in company-managed (classic) projects.
	//
	// DELETE /rest/api/{2-3}/fieldconfiguration/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#delete-field-configuration
	Delete(ctx context.Context, id int) (*model.ResponseScheme, error)
}

// FieldConfigItemConnector interface holds the methods available for the FieldConfigItem resource.
type FieldConfigItemConnector interface {

	// Gets Returns a paginated list of all fields for a configuration.
	//
	// GET /rest/api/{2-3}/fieldconfiguration/{id}/fields
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/items#get-field-configuration-items
	Gets(ctx context.Context, id, startAt, maxResults int) (*model.FieldConfigurationItemPageScheme, *model.ResponseScheme, error)

	// Update updates fields in a field configuration. The properties of the field configuration fields provided
	// override the existing values.
	//
	// 1. This operation can only update field configurations used in company-managed (classic) projects.
	//
	// PUT /rest/api/{2-3}/fieldconfiguration/{id}/fields
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/items#update-field-configuration-items
	Update(ctx context.Context, id int, payload *model.UpdateFieldConfigurationItemPayloadScheme) (*model.ResponseScheme, error)
}

// FieldConfigSchemeConnector interface holds the methods available for the FieldConfigScheme resource.
type FieldConfigSchemeConnector interface {

	// Gets returns a paginated list of field configuration schemes.
	//
	// Only field configuration schemes used in classic projects are returned.
	//
	// GET /rest/api/{2-3}/fieldconfigurationscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#get-field-configuration-schemes
	Gets(ctx context.Context, ids []int, startAt, maxResults int) (*model.FieldConfigurationSchemePageScheme, *model.ResponseScheme, error)

	// Create creates a field configuration scheme.
	//
	// This operation can only create field configuration schemes used in company-managed (classic) projects.
	//
	// POST /rest/api/{2-3}/fieldconfigurationscheme
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#create-field-configuration-scheme
	Create(ctx context.Context, name, description string) (*model.FieldConfigurationSchemeScheme, *model.ResponseScheme, error)

	// Mapping returns a paginated list of field configuration issue type items.
	//
	// Only items used in classic projects are returned.
	//
	// GET /rest/api/{2-3}/fieldconfigurationscheme/mapping
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#get-field-configuration-scheme-mapping
	Mapping(ctx context.Context, fieldConfigIDs []int, startAt, maxResults int) (*model.FieldConfigurationIssueTypeItemPageScheme,
		*model.ResponseScheme, error)

	// Project returns a paginated list of field configuration schemes and, for each scheme, a list of the projects that use it.
	//
	// 1. The list is sorted by field configuration scheme ID. The first item contains the list of project IDs assigned to the default field configuration scheme.
	//
	// 2. Only field configuration schemes used in classic projects are returned.\
	//
	// GET /rest/api/{2-3}/fieldconfigurationscheme/project
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#get-field-configuration-schemes-by-project
	Project(ctx context.Context, projectIDs []int, startAt, maxResults int) (*model.FieldConfigurationSchemeProjectPageScheme,
		*model.ResponseScheme, error)

	// Assign assigns a field configuration scheme to a project. If the field configuration scheme ID is null,
	//
	// the operation assigns the default field configuration scheme.
	//
	// Field configuration schemes can only be assigned to classic projects.
	//
	// PUT /rest/api/{2-3}/fieldconfigurationscheme/project
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#assign-field-configuration-scheme
	Assign(ctx context.Context, payload *model.FieldConfigurationSchemeAssignPayload) (*model.ResponseScheme, error)

	// Update updates a field configuration scheme.
	//
	// This operation can only update field configuration schemes used in company-managed (classic) projects.
	//
	// PUT /rest/api/{2-3}/fieldconfigurationscheme/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#update-field-configuration-scheme
	Update(ctx context.Context, schemeID int, name, description string) (*model.ResponseScheme, error)

	// Delete deletes a field configuration scheme.
	//
	// This operation can only delete field configuration schemes used in company-managed (classic) projects.
	//
	// DELETE /rest/api/{2-3}/fieldconfigurationscheme/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#delete-field-configuration-scheme
	Delete(ctx context.Context, schemeID int) (*model.ResponseScheme, error)

	// Link assigns issue types to field configurations on field configuration scheme.
	//
	// This operation can only modify field configuration schemes used in company-managed (classic) projects.
	//
	// PUT /rest/api/{2-3}/fieldconfigurationscheme/{id}/mapping
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#assign-issue-types-to-field-configuration
	Link(ctx context.Context, schemeID int, payload *model.FieldConfigurationToIssueTypeMappingPayloadScheme) (
		*model.ResponseScheme, error)

	// Unlink removes issue types from the field configuration scheme.
	//
	// This operation can only modify field configuration schemes used in company-managed (classic) projects.
	//
	// POST /rest/api/{2-3}/fieldconfigurationscheme/{id}/mapping/delete
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/schemes#remove-issue-types-to-field-configuration
	Unlink(ctx context.Context, schemeID int, issueTypeIDs []string) (*model.ResponseScheme, error)
}
