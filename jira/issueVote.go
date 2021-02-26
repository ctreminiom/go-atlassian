package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type VoteService struct{ client *Client }

type IssueVoteScheme struct {
	Self     string `json:"self,omitempty"`
	Votes    int    `json:"votes,omitempty"`
	HasVoted bool   `json:"hasVoted,omitempty"`
	Voters   []struct {
		Self       string `json:"self,omitempty"`
		Key        string `json:"key,omitempty"`
		AccountID  string `json:"accountId,omitempty"`
		Name       string `json:"name,omitempty"`
		AvatarUrls struct {
			Four8X48  string `json:"48x48,omitempty"`
			Two4X24   string `json:"24x24,omitempty"`
			One6X16   string `json:"16x16,omitempty"`
			Three2X32 string `json:"32x32,omitempty"`
		} `json:"avatarUrls,omitempty"`
		DisplayName string `json:"displayName,omitempty"`
		Active      bool   `json:"active,omitempty"`
	} `json:"voters,omitempty"`
}

// Returns details about the votes on an issue.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-votes/#api-rest-api-3-issue-issueidorkey-votes-get
func (v *VoteService) Gets(ctx context.Context, issueKeyOrID string) (result *IssueVoteScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/votes", issueKeyOrID)

	request, err := v.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = v.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueVoteScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Adds the user's vote to an issue. This is the equivalent of the user clicking Vote on an issue in Jira.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-votes/#api-rest-api-3-issue-issueidorkey-votes-post
func (v *VoteService) Add(ctx context.Context, issueKeyOrID string) (response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, fmt.Errorf("error, please provide a issueKeyOrID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/votes", issueKeyOrID)

	request, err := v.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = v.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Deletes a user's vote from an issue. This is the equivalent of the user clicking Unvote on an issue in Jira.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-votes/#api-rest-api-3-issue-issueidorkey-votes-delete
func (v *VoteService) Delete(ctx context.Context, issueKeyOrID string) (response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, fmt.Errorf("error, please provide a issueKeyOrID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/votes", issueKeyOrID)

	request, err := v.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = v.client.Do(request)
	if err != nil {
		return
	}

	return
}
