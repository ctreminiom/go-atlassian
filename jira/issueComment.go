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
	Self   string      `json:"self"`
	ID     string      `json:"id"`
	Author *UserScheme `json:"author"`
	Body   struct {
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
	Created    string `json:"created"`
	Updated    string `json:"updated"`
	Visibility struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"visibility"`
}

// Returns all comments for an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comments
func (c *CommentService) Gets(ctx context.Context, issueKeyOrID, orderBy string, expands []string, startAt, maxResults int) (result *IssueCommentScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

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

// Returns a comment.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comment
func (c *CommentService) Get(ctx context.Context, issueKeyOrID, commentID string) (result *CommentScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	if len(commentID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid commentID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/comment/%v", issueKeyOrID, commentID)

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#delete-comment
func (c *CommentService) Delete(ctx context.Context, issueKeyOrID, commentID string) (response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	if len(commentID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid commentID value")
	}

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
