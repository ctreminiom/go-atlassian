package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/tidwall/gjson"
)

type MetadataConnector interface {

	// Get edit issue metadata returns the edit screen fields for an issue that are visible to and editable by the user.
	//
	// Use the information to populate the requests in Edit issue.
	//
	// TODO: the documentation needs to be created
	Get(ctx context.Context, issueKeyOrId string, overrideScreenSecurity, overrideEditableFlag bool) (gjson.Result, *model.ResponseScheme, error)

	// Create returns details of projects, issue types within projects, and, when requested,
	//
	// the create screen fields for each issue type for the user.
	// TODO: the documentation needs to be created
	Create(ctx context.Context, opts *model.IssueMetadataCreateOptions) (gjson.Result, *model.ResponseScheme, error)
}
