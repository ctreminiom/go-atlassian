package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type FieldConfiguration interface {

	// Gets Returns a paginated list of all field configurations.
	// GET /rest/api/{2-3}/fieldconfiguration
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-all-field-configurations
	Gets(ctx context.Context, ids []int, isDefault bool, startAt, maxResults int) (*model.FieldConfigurationPageScheme,
		*model.ResponseScheme, error)

	// Create creates a field configuration. The field configuration is created with the same field properties as the
	// default configuration, with all the fields being optional.
	// This operation can only create configurations for use in company-managed (classic) projects.
	// POST /rest/api/{2-3}/fieldconfiguration
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#create-field-configuration
	Create(ctx context.Context, name, description string) (*model.FieldConfigurationScheme, *model.ResponseScheme, error)

	// Update updates a field configuration. The name and the description provided in the request override the existing values.
	// This operation can only update configurations used in company-managed (classic) projects.
	// PUT /rest/api/{2-3}/fieldconfiguration/{id}
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#update-field-configuration
	Update(ctx context.Context, id int, name, description string) (*model.ResponseScheme, error)

	// Delete deletes a field configuration.
	// This operation can only delete configurations used in company-managed (classic) projects.
	// DELETE /rest/api/{2-3}/fieldconfiguration/{id}
	// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#delete-field-configuration
	Delete(ctx context.Context, id int) (*model.ResponseScheme, error)
}
