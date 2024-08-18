package models

// IssueSearchCheckPayloadScheme represents the payload for checking issue search in Jira.
type IssueSearchCheckPayloadScheme struct {
	IssueIDs []int    `json:"issueIds,omitempty"` // The IDs of the issues.
	JQLs     []string `json:"jqls,omitempty"`     // The JQLs for the issue search.
}

// IssueMatchesPageScheme represents a page of issue matches in Jira.
type IssueMatchesPageScheme struct {
	Matches []*IssueMatchesScheme `json:"matches,omitempty"` // The issue matches in the page.
}

// IssueMatchesScheme represents the matches of an issue in Jira.
type IssueMatchesScheme struct {
	MatchedIssues []int    `json:"matchedIssues,omitempty"` // The matched issues.
	Errors        []string `json:"errors,omitempty"`        // The errors occurred during the matching process.
}
