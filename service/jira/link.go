package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type LinkSharedConnector interface {

	// Get returns an issue link.
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link#get-issue-link
	Get(ctx context.Context, linkId string) (*model.IssueLinkScheme, *model.ResponseScheme, error)

	// Gets get the issue links ID's associated with a Jira Issue
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link#get-issue-links
	Gets(ctx context.Context, issueKeyOrId string) (*model.IssueLinkPageScheme, *model.ResponseScheme, error)

	// Delete deletes an issue link.
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/link#delete-issue-link
	Delete(ctx context.Context, linkId string) (*model.ResponseScheme, error)
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
