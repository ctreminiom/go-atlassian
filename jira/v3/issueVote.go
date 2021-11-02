package v3

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"net/http"
)

type VoteService struct{ client *Client }

// Gets returns details about the votes on an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/vote#get-votes
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-votes/#api-rest-api-3-issue-issueidorkey-votes-get
func (v *VoteService) Gets(ctx context.Context, issueKeyOrID string) (result *models.IssueVoteScheme, response *ResponseScheme,
	err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, models.ErrNoIssueKeyOrIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/votes", issueKeyOrID)

	request, err := v.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = v.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Add adds the user's vote to an issue. This is the equivalent of the user clicking Vote on an issue in Jira.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/vote#add-vote
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-votes/#api-rest-api-3-issue-issueidorkey-votes-post
func (v *VoteService) Add(ctx context.Context, issueKeyOrID string) (response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, models.ErrNoIssueKeyOrIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/votes", issueKeyOrID)

	request, err := v.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = v.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Delete deletes a user's vote from an issue. This is the equivalent of the user clicking Unvote on an issue in Jira.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/vote#delete-vote
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-votes/#api-rest-api-3-issue-issueidorkey-votes-delete
func (v *VoteService) Delete(ctx context.Context, issueKeyOrID string) (response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, models.ErrNoIssueKeyOrIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/votes", issueKeyOrID)

	request, err := v.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = v.client.call(request, nil)
	if err != nil {
		return
	}

	return
}
