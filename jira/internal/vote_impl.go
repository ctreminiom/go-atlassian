package internal

import (
	"context"
	"fmt"
	"net/http"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// NewVoteService creates a new instance of VoteService.
func NewVoteService(client service.Connector, version string) (*VoteService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &VoteService{
		internalClient: &internalVoteImpl{c: client, version: version},
	}, nil
}

// VoteService provides methods to manage votes on issues in Jira Service Management.
type VoteService struct {
	// internalClient is the connector interface for vote operations.
	internalClient jira.VoteConnector
}

// Gets returns details about the votes on an issue.
//
// # This operation requires allowing users to vote on issues option to be ON
//
// GET /rest/api/{2-3}/issue/{issueKeyOrID}/votes
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/vote#get-votes
func (v *VoteService) Gets(ctx context.Context, issueKeyOrID string) (*model.IssueVoteScheme, *model.ResponseScheme, error) {
	return v.internalClient.Gets(ctx, issueKeyOrID)
}

// Add adds the user's vote to an issue. This is the equivalent of the user clicking Vote on an issue in Jira.
//
// This operation requires the Allow users to vote on issues option to be ON.
//
// POST /rest/api/{2-3}/issue/{issueKeyOrID}/votes
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/vote#add-vote
func (v *VoteService) Add(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error) {
	return v.internalClient.Add(ctx, issueKeyOrID)
}

// Delete deletes a user's vote from an issue. This is the equivalent of the user clicking Unvote on an issue in Jira.
//
// This operation requires the Allow users to vote on issues option to be ON.
//
// DELETE /rest/api/{2-3}/issue/{issueKeyOrID}/votes
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/vote#delete-vote
func (v *VoteService) Delete(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error) {
	return v.internalClient.Delete(ctx, issueKeyOrID)
}

type internalVoteImpl struct {
	c       service.Connector
	version string
}

func (i *internalVoteImpl) Gets(ctx context.Context, issueKeyOrID string) (*model.IssueVoteScheme, *model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, nil, model.ErrNoIssueKeyOrID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/votes", i.version, issueKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	votes := new(model.IssueVoteScheme)
	response, err := i.c.Call(request, votes)
	if err != nil {
		return nil, response, err
	}

	return votes, response, nil
}

func (i *internalVoteImpl) Add(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, model.ErrNoIssueKeyOrID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/votes", i.version, issueKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalVoteImpl) Delete(ctx context.Context, issueKeyOrID string) (*model.ResponseScheme, error) {

	if issueKeyOrID == "" {
		return nil, model.ErrNoIssueKeyOrID
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/votes", i.version, issueKeyOrID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
