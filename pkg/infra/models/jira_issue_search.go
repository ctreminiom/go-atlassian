package models

type IssueSearchCheckPayloadScheme struct {
	IssueIds []int    `json:"issueIds,omitempty"`
	JQLs     []string `json:"jqls,omitempty"`
}

type IssueMatchesPageScheme struct {
	Matches []*IssueMatchesScheme `json:"matches,omitempty"`
}

type IssueMatchesScheme struct {
	MatchedIssues []int    `json:"matchedIssues,omitempty"`
	Errors        []string `json:"errors,omitempty"`
}
