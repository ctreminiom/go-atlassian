package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type CommentService struct{ client *Client }

// Gets returns all comments for an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comments
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-comments/#api-rest-api-2-issue-issueidorkey-comment-get
func (c *CommentService) Gets(ctx context.Context, issueKeyOrID, orderBy string, expand []string, startAt,
	maxResults int) (result *models.IssueCommentPageSchemeV2, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, models.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	if orderBy != "" {
		params.Add("orderBy", orderBy)
	}

	var endpoint = fmt.Sprintf("rest/api/2/issue/%v/comment?%v", issueKeyOrID, params.Encode())
	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = c.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns a comment.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comment
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-comments/#api-rest-api-2-comment-list-post
func (c *CommentService) Get(ctx context.Context, issueKeyOrID, commentID string) (result *models.IssueCommentSchemeV2,
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, models.ErrNoIssueKeyOrIDError
	}

	if len(commentID) == 0 {
		return nil, nil, models.ErrNoCommentIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/issue/%v/comment/%v", issueKeyOrID, commentID)

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = c.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes a comment.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#delete-comment
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-comments/#api-rest-api-2-issue-issueidorkey-comment-id-delete
func (c *CommentService) Delete(ctx context.Context, issueKeyOrID, commentID string) (response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, models.ErrNoIssueKeyOrIDError
	}

	if len(commentID) == 0 {
		return nil, models.ErrNoCommentIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/issue/%v/comment/%v", issueKeyOrID, commentID)

	request, err := c.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = c.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Add adds a comment to an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#add-comment
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-comments/#api-rest-api-2-issue-issueidorkey-comment-post
func (c *CommentService) Add(ctx context.Context, issueKeyOrID string, payload *models.CommentPayloadSchemeV2, expand []string) (
	result *models.IssueCommentSchemeV2, response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, models.ErrNoIssueKeyOrIDError
	}

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/2/issue/%v/comment", issueKeyOrID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := c.client.newRequest(ctx, http.MethodPost, endpoint.String(), payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
