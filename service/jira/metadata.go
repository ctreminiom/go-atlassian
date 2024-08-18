package jira

import (
	"context"

	"github.com/tidwall/gjson"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
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
}
