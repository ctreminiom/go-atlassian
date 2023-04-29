package models

type WorklogOptionsScheme struct {
	Notify               bool
	AdjustEstimate       string
	NewEstimate          string
	ReduceBy             string
	OverrideEditableFlag bool
	Expand               []string
}

type WorklogADFPayloadScheme struct {
	Comment          *CommentNodeScheme            `json:"comment,omitempty"`
	Visibility       *IssueWorklogVisibilityScheme `json:"visibility,omitempty"`
	Started          string                        `json:"started,omitempty"`
	TimeSpent        string                        `json:"timeSpent,omitempty"`
	TimeSpentSeconds int                           `json:"timeSpentSeconds,omitempty"`
}

type WorklogRichTextPayloadScheme struct {
	Comment          *CommentPayloadSchemeV2       `json:"comment,omitempty"`
	Visibility       *IssueWorklogVisibilityScheme `json:"visibility,omitempty"`
	Started          string                        `json:"started,omitempty"`
	TimeSpent        string                        `json:"timeSpent,omitempty"`
	TimeSpentSeconds int                           `json:"timeSpentSeconds,omitempty"`
}

type ChangedWorklogPageScheme struct {
	Since    int                     `json:"since,omitempty"`
	Until    int                     `json:"until,omitempty"`
	Self     string                  `json:"self,omitempty"`
	NextPage string                  `json:"nextPage,omitempty"`
	LastPage bool                    `json:"lastPage,omitempty"`
	Values   []*ChangedWorklogScheme `json:"values,omitempty"`
}

type ChangedWorklogScheme struct {
	WorklogID   int                             `json:"worklogId,omitempty"`
	UpdatedTime int                             `json:"updatedTime,omitempty"`
	Properties  []*ChangedWorklogPropertyScheme `json:"properties,omitempty"`
}

type ChangedWorklogPropertyScheme struct {
	Key string `json:"key,omitempty"`
}

type IssueWorklogRichTextPageScheme struct {
	StartAt    int                           `json:"startAt,omitempty"`
	MaxResults int                           `json:"maxResults,omitempty"`
	Total      int                           `json:"total,omitempty"`
	Worklogs   []*IssueWorklogRichTextScheme `json:"worklogs,omitempty"`
}

type IssueWorklogADFPageScheme struct {
	StartAt    int                      `json:"startAt,omitempty"`
	MaxResults int                      `json:"maxResults,omitempty"`
	Total      int                      `json:"total,omitempty"`
	Worklogs   []*IssueWorklogADFScheme `json:"worklogs,omitempty"`
}

type IssueWorklogVisibilityScheme struct {
	Type       string `json:"type,omitempty"`
	Value      string `json:"value,omitempty"`
	Identifier string `json:"identifier,omitempty"`
}

type IssueWorklogRichTextScheme struct {
	Self             string                        `json:"self,omitempty"`
	Author           *UserDetailScheme             `json:"author,omitempty"`
	UpdateAuthor     *UserDetailScheme             `json:"updateAuthor,omitempty"`
	Comment          string                        `json:"comment,omitempty"`
	Updated          string                        `json:"updated,omitempty"`
	Visibility       *IssueWorklogVisibilityScheme `json:"visibility,omitempty"`
	Started          string                        `json:"started,omitempty"`
	TimeSpent        string                        `json:"timeSpent,omitempty"`
	TimeSpentSeconds int                           `json:"timeSpentSeconds,omitempty"`
	ID               string                        `json:"id,omitempty"`
	IssueID          string                        `json:"issueId,omitempty"`
}

type IssueWorklogADFScheme struct {
	Self             string                        `json:"self,omitempty"`
	Author           *UserDetailScheme             `json:"author,omitempty"`
	UpdateAuthor     *UserDetailScheme             `json:"updateAuthor,omitempty"`
	Comment          *CommentNodeScheme            `json:"comment,omitempty"`
	Updated          string                        `json:"updated,omitempty"`
	Visibility       *IssueWorklogVisibilityScheme `json:"visibility,omitempty"`
	Started          string                        `json:"started,omitempty"`
	TimeSpent        string                        `json:"timeSpent,omitempty"`
	TimeSpentSeconds int                           `json:"timeSpentSeconds,omitempty"`
	ID               string                        `json:"id,omitempty"`
	IssueID          string                        `json:"issueId,omitempty"`
}
