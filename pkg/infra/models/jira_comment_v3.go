// Package models provides the data structures used in the admin package.
package models

// CommentNodeScheme represents a node in a comment.
type CommentNodeScheme struct {
	Version int                    `json:"version,omitempty"` // The version of the node.
	Type    string                 `json:"type,omitempty"`    // The type of the node.
	Content []*CommentNodeScheme   `json:"content,omitempty"` // The content of the node.
	Text    string                 `json:"text,omitempty"`    // The text of the node.
	Attrs   map[string]interface{} `json:"attrs,omitempty"`   // The attributes of the node.
	Marks   []*MarkScheme          `json:"marks,omitempty"`   // The marks of the node.
}

// AppendNode appends a node to the content of the CommentNodeScheme.
func (n *CommentNodeScheme) AppendNode(node *CommentNodeScheme) {
	n.Content = append(n.Content, node)
}

// MarkScheme represents a mark in a comment.
type MarkScheme struct {
	Type  string                 `json:"type,omitempty"`  // The type of the mark.
	Attrs map[string]interface{} `json:"attrs,omitempty"` // The attributes of the mark.
}

// CommentPayloadScheme represents the payload of a comment.
type CommentPayloadScheme struct {
	Visibility *CommentVisibilityScheme `json:"visibility,omitempty"` // The visibility of the comment.
	Body       *CommentNodeScheme       `json:"body,omitempty"`       // The body of the comment.
}

// IssueCommentPageScheme represents a page of issue comments.
type IssueCommentPageScheme struct {
	StartAt    int                   `json:"startAt,omitempty"`    // The start index of the comments.
	MaxResults int                   `json:"maxResults,omitempty"` // The maximum number of comments per page.
	Total      int                   `json:"total,omitempty"`      // The total number of comments.
	Comments   []*IssueCommentScheme `json:"comments,omitempty"`   // The comments on the page.
}

// IssueCommentScheme represents a comment on an issue.
type IssueCommentScheme struct {
	Self         string                   `json:"self,omitempty"`         // The self link of the comment.
	ID           string                   `json:"id,omitempty"`           // The ID of the comment.
	Author       *UserScheme              `json:"author,omitempty"`       // The author of the comment.
	RenderedBody string                   `json:"renderedBody,omitempty"` // The rendered body of the comment.
	Body         *CommentNodeScheme       `json:"body,omitempty"`         // The body of the comment.
	JSDPublic    bool                     `json:"jsdPublic,omitempty"`    // Whether the comment is public.
	UpdateAuthor *UserScheme              `json:"updateAuthor,omitempty"` // The author of the last update.
	Created      string                   `json:"created,omitempty"`      // The creation time of the comment.
	Updated      string                   `json:"updated,omitempty"`      // The last update time of the comment.
	Visibility   *CommentVisibilityScheme `json:"visibility,omitempty"`   // The visibility of the comment.
}

// CommentVisibilityScheme represents the visibility of a comment.
type CommentVisibilityScheme struct {
	Type  string `json:"type,omitempty"`  // The type of the visibility.
	Value string `json:"value,omitempty"` // The value of the visibility.
}
