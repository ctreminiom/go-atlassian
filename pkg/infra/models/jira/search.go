package jira

type IssueSearchScheme struct {
	Expand          string         `json:"expand,omitempty"`
	StartAt         int            `json:"startAt,omitempty"`
	MaxResults      int            `json:"maxResults,omitempty"`
	Total           int            `json:"total,omitempty"`
	Issues          []*IssueScheme `json:"issues,omitempty"`
	WarningMessages []string       `json:"warningMessages,omitempty"`
}

type IssueTransitionsScheme struct {
	Expand      string                   `json:"expand,omitempty"`
	Transitions []*IssueTransitionScheme `json:"transitions,omitempty"`
}
