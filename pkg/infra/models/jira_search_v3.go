package models

// IssueSearchScheme represents the results of an issue search in Jira.
type IssueSearchScheme struct {
	Expand          string         `json:"expand,omitempty"`          // The fields that are expanded in the results.
	StartAt         int            `json:"startAt,omitempty"`         // The index of the first result returned.
	MaxResults      int            `json:"maxResults,omitempty"`      // The maximum number of results returned.
	Total           int            `json:"total,omitempty"`           // The total number of results available.
	Issues          []*IssueScheme `json:"issues,omitempty"`          // The issues returned in the results.
	WarningMessages []string       `json:"warningMessages,omitempty"` // Any warning messages generated during the search.
}

// IssueTransitionsScheme represents the transitions of an issue in Jira.
type IssueTransitionsScheme struct {
	Expand      string                   `json:"expand,omitempty"`      // The fields that are expanded in the results.
	Transitions []*IssueTransitionScheme `json:"transitions,omitempty"` // The transitions of the issue.
}

// IssueSearchJQLScheme represents the response from the new JQL search endpoint for ADF (v3 API)
type IssueSearchJQLScheme struct {
	StartAt       int               `json:"startAt,omitempty"`
	MaxResults    int               `json:"maxResults,omitempty"`
	Total         int               `json:"total,omitempty"`
	Issues        []*IssueScheme    `json:"issues,omitempty"`
	Names         map[string]string `json:"names,omitempty"`
	Schema        map[string]string `json:"schema,omitempty"`
	NextPageToken string            `json:"nextPageToken,omitempty"`
}

// IssueBulkFetchScheme represents the response from the bulk fetch endpoint for ADF (v3 API)
type IssueBulkFetchScheme struct {
	Issues []*IssueScheme `json:"issues,omitempty"`
}
