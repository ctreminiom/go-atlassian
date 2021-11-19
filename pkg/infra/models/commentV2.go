package models

type IssueCommentPageSchemeV2 struct {
	StartAt    int                     `json:"startAt,omitempty"`
	MaxResults int                     `json:"maxResults,omitempty"`
	Total      int                     `json:"total,omitempty"`
	Comments   []*IssueCommentSchemeV2 `json:"comments,omitempty"`
}

type IssueCommentSchemeV2 struct {
	Self         string                   `json:"self,omitempty"`
	ID           string                   `json:"id,omitempty"`
	Body         string                   `json:"body,omitempty"`
	RenderedBody string                   `json:"renderedBody,omitempty"`
	Author       *UserScheme              `json:"author,omitempty"`
	JSDPublic    bool                     `json:"jsdPublic,omitempty"`
	UpdateAuthor *UserScheme              `json:"updateAuthor,omitempty"`
	Created      string                   `json:"created,omitempty"`
	Updated      string                   `json:"updated,omitempty"`
	Visibility   *CommentVisibilityScheme `json:"visibility,omitempty"`
}

type CommentPayloadSchemeV2 struct {
	Visibility *CommentVisibilityScheme `json:"visibility,omitempty"`
	Body       string                   `json:"body,omitempty"`
}
