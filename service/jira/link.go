package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type LinkSharedConnector interface {

	// Get returns an issue link.
	//
	// GET /rest/api/{2-3}/issueLink/{linkID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link#get-issue-link
	Get(ctx context.Context, linkID string) (*model.IssueLinkScheme, *model.ResponseScheme, error)

	// Gets get the issue links ID's associated with a Jira Issue
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link#get-issue-links
	Gets(ctx context.Context, issueKeyOrID string) (*model.IssueLinkPageScheme, *model.ResponseScheme, error)

	// Delete deletes an issue link.
	//
	// DELETE /rest/api/{2-3}/issueLink/{linkID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link#delete-issue-link
	Delete(ctx context.Context, linkID string) (*model.ResponseScheme, error)
}

type LinkRichTextConnector interface {
	LinkSharedConnector

	// Create creates a link between two issues. Use this operation to indicate a relationship between two issues
	//
	// and optionally add a comment to the from (outward) issue.
	//
	// To use this resource the site must have Issue Linking enabled.
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link#create-issue-link
	Create(ctx context.Context, payload *model.LinkPayloadSchemeV2) (*model.ResponseScheme, error)
}

type LinkAdfIssueConnector interface {
	LinkSharedConnector

	// Create creates a link between two issues. Use this operation to indicate a relationship between two issues
	//
	// and optionally add a comment to the from (outward) issue.
	//
	// To use this resource the site must have Issue Linking enabled.
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link#create-issue-link
	Create(ctx context.Context, payload *model.LinkPayloadSchemeV3) (*model.ResponseScheme, error)
}

// LinkTypeConnector is an interface that defines the methods available from Issue Link Type  API.
// Use it to get, create, update, and delete link issue types as well as get lists of all link issue types.
type LinkTypeConnector interface {

	// Gets returns a list of all issue link types.
	//
	// GET /rest/api/{2-3}/issueLinkType
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#get-issue-link-types
	Gets(ctx context.Context) (*model.IssueLinkTypeSearchScheme, *model.ResponseScheme, error)

	// Get returns an issue link type.
	//
	//
	// GET /rest/api/{2-3}/issueLinkType/{issueLinkTypeID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#get-issue-link-type
	Get(ctx context.Context, issueLinkTypeID string) (*model.LinkTypeScheme, *model.ResponseScheme, error)

	// Create creates an issue link type.
	//
	// Use this operation to create descriptions of the reasons why issues are linked.
	//
	// The issue link type consists of a name and descriptions for a link's inward and outward relationships.
	//
	// POST /rest/api/{2-3}/issueLinkType
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#create-issue-link-type
	Create(ctx context.Context, payload *model.LinkTypeScheme) (*model.LinkTypeScheme, *model.ResponseScheme, error)

	// Update updates an issue link type.
	//
	// PUT /rest/api/{2-3}/issueLinkType/{issueLinkTypeID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#update-issue-link-type
	Update(ctx context.Context, issueLinkTypeID string, payload *model.LinkTypeScheme) (*model.LinkTypeScheme, *model.ResponseScheme, error)

	// Delete deletes an issue link type.
	//
	// DELETE /rest/api/{2-3}/issueLinkType/{issueLinkTypeID}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link/types#delete-issue-link-type
	Delete(ctx context.Context, issueLinkTypeID string) (*model.ResponseScheme, error)
}
