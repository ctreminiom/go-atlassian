package models

// IssueCommentPageSchemeV2 represents a page of issue comments in Jira.
type IssueCommentPageSchemeV2 struct {
	StartAt    int                     `json:"startAt,omitempty"`    // The starting index of the page.
	MaxResults int                     `json:"maxResults,omitempty"` // The maximum number of results in the page.
	Total      int                     `json:"total,omitempty"`      // The total number of comments.
	Comments   []*IssueCommentSchemeV2 `json:"comments,omitempty"`   // The comments in the page.
}

// IssueCommentSchemeV2 represents an issue comment in Jira.
type IssueCommentSchemeV2 struct {
	Self         string                   `json:"self,omitempty"`         // The URL of the comment.
	ID           string                   `json:"id,omitempty"`           // The ID of the comment.
	Body         string                   `json:"body,omitempty"`         // The body of the comment.
	RenderedBody string                   `json:"renderedBody,omitempty"` // The rendered body of the comment.
	Author       *UserScheme              `json:"author,omitempty"`       // The author of the comment.
	JSDPublic    bool                     `json:"jsdPublic,omitempty"`    // Indicates if the comment is public in Jira Service Desk.
	UpdateAuthor *UserScheme              `json:"updateAuthor,omitempty"` // The user who last updated the comment.
	Created      string                   `json:"created,omitempty"`      // The creation time of the comment.
	Updated      string                   `json:"updated,omitempty"`      // The last update time of the comment.
	Visibility   *CommentVisibilityScheme `json:"visibility,omitempty"`   // The visibility of the comment.
}

// CommentPayloadSchemeV2 represents the payload for an issue comment in Jira.
type CommentPayloadSchemeV2 struct {
	Visibility *CommentVisibilityScheme `json:"visibility,omitempty"` // The visibility of the comment.
	Body       string                   `json:"body,omitempty"`       // The body of the comment.
}
