package models

// IssueLabelsScheme represents the labels of an issue in Jira.
type IssueLabelsScheme struct {
	MaxResults int      `json:"maxResults"` // The maximum number of results.
	StartAt    int      `json:"startAt"`    // The starting index of the results.
	Total      int      `json:"total"`      // The total number of results.
	IsLast     bool     `json:"isLast"`     // Indicates if this is the last page of results.
	Values     []string `json:"values"`     // The labels of the issue.
}
