package jira

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

type CommentVisibilityScheme struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
