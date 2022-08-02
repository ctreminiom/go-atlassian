package jira

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type TypeConnector interface {

	// Gets returns all issue types.
	//
	// GET /rest/api/3/issuetype
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/type#get-all-issue-types-for-user
	Gets(ctx context.Context) ([]*model.IssueTypeScheme, *model.ResponseScheme, error)

	// Create creates an issue type and adds it to the default issue type scheme.
	//
	// POST /rest/api/3/issuetype
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/type#create-issue-type
	Create(ctx context.Context, payload *model.IssueTypePayloadScheme) (*model.IssueTypeScheme, *model.ResponseScheme, error)

	// Get returns an issue type.
	//
	// GET /rest/api/3/issuetype/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/type#get-issue-type
	Get(ctx context.Context, issueTypeId string) (*model.IssueTypeScheme, *model.ResponseScheme, error)

	// Update updates the issue type.
	//
	// PUT /rest/api/3/issuetype/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/type#update-issue-type
	Update(ctx context.Context, issueTypeId string, payload *model.IssueTypePayloadScheme) (*model.IssueTypeScheme, *model.ResponseScheme, error)

	// Delete deletes the issue type.
	//
	// If the issue type is in use, all uses are updated with the alternative issue type (alternativeIssueTypeId).
	// A list of alternative issue types are obtained from the Get alternative issue types resource.
	//
	// DELETE /rest/api/3/issuetype/{id}
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/type#delete-issue-type
	Delete(ctx context.Context, issueTypeId string) (*model.ResponseScheme, error)

	// Alternatives returns a list of issue types that can be used to replace the issue type.
	//
	// The alternative issue types are those assigned to the same workflow scheme, field configuration scheme, and screen scheme.
	//
	// GET /rest/api/3/issuetype/{id}/alternatives
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/type#get-alternative-issue-types
	Alternatives(ctx context.Context, issueTypeId string) ([]*model.IssueTypeScheme, *model.ResponseScheme, error)
}
