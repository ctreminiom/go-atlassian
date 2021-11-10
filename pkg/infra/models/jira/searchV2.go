package jira

type IssueSearchSchemeV2 struct {
	Expand          string           `json:"expand,omitempty"`
	StartAt         int              `json:"startAt,omitempty"`
	MaxResults      int              `json:"maxResults,omitempty"`
	Total           int              `json:"total,omitempty"`
	Issues          []*IssueSchemeV2 `json:"issues,omitempty"`
	WarningMessages []string         `json:"warningMessages,omitempty"`
}
