package jira

import (
	"context"

	"github.com/tidwall/gjson"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type MetadataConnector interface {

	// Get edit issue metadata returns the edit screen fields for an issue that are visible to and editable by the user.
	//
	// Use the information to populate the requests in Edit issue.
	//
	// GET /rest/api/{2-3}/issue/{issueKeyOrID}/editmeta
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/metadata#get-edit-issue-metadata
	Get(ctx context.Context, issueKeyOrID string, overrideScreenSecurity, overrideEditableFlag bool) (gjson.Result, *model.ResponseScheme, error)

	// Create returns details of projects, issue types within projects, and, when requested,
	//
	// the create screen fields for each issue type for the user.
	//
	// GET /rest/api/{2-3}/issue/createmeta
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/metadata#get-create-issue-metadata
	Create(ctx context.Context, opts *model.IssueMetadataCreateOptions) (gjson.Result, *model.ResponseScheme, error)

	// FetchIssueMappings returns a page of issue type metadata for a specified project.
	//
	// Use the information to populate the requests in Create issue and Create issues.
	//
	// This operation can be accessed anonymously.
	//
	// GET /rest/api/{2-3}/issue/createmeta/{projectIdOrKey}/issuetypes
	//
	// Parameters:
	// - ctx: The context for the request.
	// - projectKeyOrID: The key or ID of the project.
	// - startAt: The starting index of the returned issues.
	// - maxResults: The maximum number of results to return.
	//
	// Returns:
	// - A gjson.Result containing the issue type metadata.
	// - A pointer to the response scheme.
	// - An error if the retrieval fails.
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/metadata#get-create-metadata-issue-types-for-a-project
	FetchIssueMappings(ctx context.Context, projectKeyOrID string, startAt, maxResults int) (gjson.Result, *model.ResponseScheme, error)

	// FetchFieldMappings returns a page of field metadata for a specified project and issue type.
	//
	// Use the information to populate the requests in Create issue and Create issues.
	//
	// This operation can be accessed anonymously.
	//
	// GET /rest/api/{2-3}/issue/createmeta/{projectIdOrKey}/fields/{issueTypeId}
	//
	// Parameters:
	// - ctx: The context for the request.
	// - projectKeyOrID: The key or ID of the project.
	// - issueTypeID: The ID of the issue type whose metadata is to be retrieved.
	// - startAt: The starting index of the returned fields.
	// - maxResults: The maximum number of results to return.
	//
	// Returns:
	// - A gjson.Result containing the field metadata.
	// - A pointer to the response scheme.
	// - An error if the retrieval fails.
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/metadata#get-create-field-metadata-for-a-project-and-issue-type-id
	FetchFieldMappings(ctx context.Context, projectKeyOrID, issueTypeID string, startAt, maxResults int) (gjson.Result, *model.ResponseScheme, error)
}
