package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
)

func NewVoteService(client service.Client, version string) (*VoteService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &VoteService{
		internalClient: &internalVoteImpl{c: client, version: version},
	}, nil
}

type VoteService struct {
	internalClient jira.VoteConnector
}

// Gets returns details about the votes on an issue.
//
// This operation requires allowing users to vote on issues option to be ON
//
// GET /rest/api/{2-3}/issue/{issueIdOrKey}/votes
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/vote#get-votes
func (v *VoteService) Gets(ctx context.Context, issueKeyOrId string) (*model.IssueVoteScheme, *model.ResponseScheme, error) {
	return v.internalClient.Gets(ctx, issueKeyOrId)
}

// Add adds the user's vote to an issue. This is the equivalent of the user clicking Vote on an issue in Jira.
//
// This operation requires the Allow users to vote on issues option to be ON.
//
// POST /rest/api/{2-3}/issue/{issueIdOrKey}/votes
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/vote#add-vote
func (v *VoteService) Add(ctx context.Context, issueKeyOrId string) (*model.ResponseScheme, error) {
	return v.internalClient.Add(ctx, issueKeyOrId)
}

// Delete deletes a user's vote from an issue. This is the equivalent of the user clicking Unvote on an issue in Jira.
//
// This operation requires the Allow users to vote on issues option to be ON.
//
// DELETE /rest/api/{2-3}/issue/{issueIdOrKey}/votes
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/vote#delete-vote
func (v *VoteService) Delete(ctx context.Context, issueKeyOrId string) (*model.ResponseScheme, error) {
	return v.internalClient.Delete(ctx, issueKeyOrId)
}

type internalVoteImpl struct {
	c       service.Client
	version string
}

func (i *internalVoteImpl) Gets(ctx context.Context, issueKeyOrId string) (*model.IssueVoteScheme, *model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, nil, model.ErrNoIssueKeyOrIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/votes", i.version, issueKeyOrId)

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

func (i *internalVoteImpl) Add(ctx context.Context, issueKeyOrId string) (*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/votes", i.version, issueKeyOrId)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalVoteImpl) Delete(ctx context.Context, issueKeyOrId string) (*model.ResponseScheme, error) {

	if issueKeyOrId == "" {
		return nil, model.ErrNoIssueKeyOrIDError
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/%v/votes", i.version, issueKeyOrId)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
