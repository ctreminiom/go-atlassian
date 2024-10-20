package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// VoteConnector is an interface that defines the methods available from VoteConnector API.
// Use it to get details of votes on an issue as well as cast and withdrawal votes.
type VoteConnector interface {

	// Gets returns details about the votes on an issue.
	//
	// This operation requires allowing users to vote on issues option to be ON
	//
	// GET /rest/api/{2-3}/issue/{issueKeyOrID}/votes
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/vote#get-votes
	Gets(ctx context.Context, issueKeyOrID string) (*model.IssueVoteScheme, *model.ResponseScheme, error)

	// Add adds the user's vote to an issue. This is the equivalent of the user clicking Vote on an issue in Jira.
	//
	// This operation requires the Allow users to vote on issues option to be ON.
	//
	// POST /rest/api/{2-3}/issue/{issueKeyOrID}/votes
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/vote#add-vote
	Add(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error)

	// Delete deletes a user's vote from an issue. This is the equivalent of the user clicking Unvote on an issue in Jira.
	//
	// This operation requires the Allow users to vote on issues option to be ON.
	//
	// DELETE /rest/api/{2-3}/issue/{issueKeyOrID}/votes
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/vote#delete-vote
	Delete(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error)
}
