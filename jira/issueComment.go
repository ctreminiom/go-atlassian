package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type CommentService struct{ client *Client }

type IssueCommentScheme struct {
	StartAt    int             `json:"startAt"`
	MaxResults int             `json:"maxResults"`
	Total      int             `json:"total"`
	Comments   []CommentScheme `json:"comments"`
}

type CommentScheme struct {
	Self   string `json:"self"`
	ID     string `json:"id"`
	Author struct {
		Self         string `json:"self"`
		AccountID    string `json:"accountId"`
		EmailAddress string `json:"emailAddress"`
		AvatarUrls   struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
		TimeZone    string `json:"timeZone"`
		AccountType string `json:"accountType"`
	} `json:"author"`
	Body struct {
		Version int    `json:"version"`
		Type    string `json:"type"`
		Content []struct {
			Type    string `json:"type"`
			Content []struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"content"`
		} `json:"content"`
	} `json:"body"`
	JSDPublic    bool `json:"jsdPublic"`
	UpdateAuthor struct {
		Self        string `json:"self"`
		AccountID   string `json:"accountId"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
	} `json:"updateAuthor"`
	Created string `json:"created"`
	Updated string `json:"updated"`

	Visibility struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"visibility"`
}

// Returns all comments for an issue.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-comments/#api-rest-api-3-issue-issueidorkey-comment-get
func (c *CommentService) Comments(ctx context.Context, issueKeyOrID string, orderBy string, expands []string, startAt, maxResults int) (result *IssueCommentScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var expand string
	for index, value := range expands {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	if orderBy != "" {
		params.Add("orderBy", orderBy)
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/comment?%v", issueKeyOrID, params.Encode())
	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueCommentScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Adds a comment to an issue.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-comments/#api-rest-api-3-issue-issueidorkey-comment-post
func (c *CommentService) Add(ctx context.Context, issueKeyOrID string, payload interface{}) (result *CommentScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/comment", issueKeyOrID)
	request, err := c.client.newRequest(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Do(request)
	if err != nil {
		return
	}

	result = new(CommentScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns a comment.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-comments/#api-rest-api-3-issue-issueidorkey-comment-id-get
func (c *CommentService) Get(ctx context.Context, issueKeyOrID, commentID string) (result *CommentScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/comment/%v", issueKeyOrID, commentID)

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Do(request)
	if err != nil {
		return
	}

	result = new(CommentScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Updates a comment.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-comments/#api-rest-api-3-issue-issueidorkey-comment-id-put
func (c *CommentService) Update(ctx context.Context, issueKeyOrID, commentID string, payload interface{}) (result *CommentScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/comment/%v", issueKeyOrID, commentID)

	request, err := c.client.newRequest(ctx, http.MethodPut, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Do(request)
	if err != nil {
		return
	}

	result = new(CommentScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes a comment.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-comments/#api-rest-api-3-issue-issueidorkey-comment-id-delete
func (c *CommentService) Delete(ctx context.Context, issueKeyOrID, commentID string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/comment/%v", issueKeyOrID, commentID)

	request, err := c.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = c.client.Do(request)
	if err != nil {
		return
	}

	return
}
