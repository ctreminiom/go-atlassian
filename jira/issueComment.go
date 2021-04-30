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

type IssueCommentPageScheme struct {
	StartAt    int                   `json:"startAt,omitempty"`
	MaxResults int                   `json:"maxResults,omitempty"`
	Total      int                   `json:"total,omitempty"`
	Comments   []*IssueCommentScheme `json:"comments,omitempty"`
}

type IssueCommentScheme struct {
	Self         string                   `json:"self,omitempty"`
	ID           string                   `json:"id,omitempty"`
	Author       *UserScheme              `json:"author,omitempty"`
	RenderedBody string                   `json:"renderedBody,omitempty"`
	Body         *CommentNodeScheme       `json:"body,omitempty"`
	JSDPublic    bool                     `json:"jsdPublic,omitempty"`
	UpdateAuthor *UserScheme              `json:"updateAuthor,omitempty"`
	Created      string                   `json:"created,omitempty"`
	Updated      string                   `json:"updated,omitempty"`
	Visibility   *CommentVisibilityScheme `json:"visibility,omitempty"`
}

// Returns all comments for an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comments
func (c *CommentService) Gets(ctx context.Context, issueKeyOrID, orderBy string, expands []string, startAt, maxResults int) (result *IssueCommentPageScheme, response *Response, err error) {

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

	result = new(IssueCommentPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns a comment.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/comments#get-comment
func (c *CommentService) Get(ctx context.Context, issueKeyOrID, commentID string) (result *IssueCommentScheme, response *Response, err error) {

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

	result = new(IssueCommentScheme)
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

func (c *CommentService) Add(ctx context.Context, issueKeyOrID string, payload *CommentPayloadScheme, expands []string) (result *IssueCommentScheme, response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid issueKeyOrID value")
	}

	if payload == nil {
		return nil, nil, fmt.Errorf("error, please provide a valid CommentNodeScheme pointer")
	}

	params := url.Values{}
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

	var endpoint string
	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("rest/api/3/issue/%v/comment?%v", issueKeyOrID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/api/3/issue/%v/comment", issueKeyOrID)
	}

	request, err := c.client.newRequest(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
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

type CommentNodeScheme struct {
	Version int                    `json:"version,omitempty"`
	Type    string                 `json:"type,omitempty"`
	Content []*CommentNodeScheme   `json:"content,omitempty"`
	Text    string                 `json:"text,omitempty"`
	Attrs   map[string]interface{} `json:"attrs,omitempty"`
	Marks   []*MarkScheme          `json:"marks,omitempty"`
}

func (n *CommentNodeScheme) AppendNode(node *CommentNodeScheme) {
	n.Content = append(n.Content, node)
}

type MarkScheme struct {
	Type  string                 `json:"type,omitempty"`
	Attrs map[string]interface{} `json:"attrs,omitempty"`
}
