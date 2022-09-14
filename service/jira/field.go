package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type FieldConnector interface {

	// Gets returns system and custom issue fields according to the following rules:
	//
	// 1. Fields that cannot be added to the issue navigator are always returned.
	//
	// 2. Fields that cannot be placed on an issue screen are always returned.
	//
	// 3. Fields that depend on global Jira settings are only returned if the setting is enabled.
	// That is, timetracking fields, subtasks, votes, and watches.
	//
	// 4. For all other fields, this operation only returns the fields that the user has permission to view
	//
	// GET /rest/api/{2-3}/field
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields#get-fields
	Gets(ctx context.Context) ([]*model.IssueFieldScheme, *model.ResponseScheme, error)

	// Create creates a custom field.
	//
	// POST /rest/api/{2-3}/field
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields#create-custom-field
	Create(ctx context.Context, payload *model.CustomFieldScheme) (*model.IssueFieldScheme, *model.ResponseScheme, error)

	// Search returns a paginated list of fields for Classic Jira projects.
	//
	// GET /rest/api/{2-3}/field/search
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/fields#get-fields-paginated
	Search(ctx context.Context, options *model.FieldSearchOptionsScheme, startAt, maxResults int) (*model.FieldSearchPageScheme, *model.ResponseScheme, error)

	// Delete deletes a custom field. The custom field is deleted whether it is in the trash or not.
	//
	// See Edit or delete a custom field for more information on trashing and deleting custom fields.
	//
	// DELETE /rest/api/{2-3}/field/{id}
	// TODO: The documentation for the method Field.Delete() needs to be created!
	Delete(ctx context.Context, fieldId string) (*model.TaskScheme, *model.ResponseScheme, error)
}
