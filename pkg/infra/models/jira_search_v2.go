package models

// IssueSearchSchemeV2 represents the results of an issue search in Jira.
type IssueSearchSchemeV2 struct {
	Expand          string           `json:"expand,omitempty"`          // The fields that are expanded in the results.
	StartAt         int              `json:"startAt,omitempty"`         // The index of the first result returned.
	MaxResults      int              `json:"maxResults,omitempty"`      // The maximum number of results returned.
	Total           int              `json:"total,omitempty"`           // The total number of results available.
	Issues          []*IssueSchemeV2 `json:"issues,omitempty"`          // The issues returned in the results.
	WarningMessages []string         `json:"warningMessages,omitempty"` // Any warning messages generated during the search.
}

// IssueSearchJQLSchemeV2 represents the response from the new JQL search endpoint for richtext (v2 API)
type IssueSearchJQLSchemeV2 struct {
	StartAt       int               `json:"startAt,omitempty"`
	MaxResults    int               `json:"maxResults,omitempty"`
	Total         int               `json:"total,omitempty"`
	Issues        []*IssueSchemeV2  `json:"issues,omitempty"`
	Names         map[string]string `json:"names,omitempty"`
	Schema        map[string]string `json:"schema,omitempty"`
	NextPageToken string            `json:"nextPageToken,omitempty"`
}

// IssueBulkFetchSchemeV2 represents the response from the bulk fetch endpoint for richtext (v2 API)
type IssueBulkFetchSchemeV2 struct {
	Issues []*IssueSchemeV2 `json:"issues,omitempty"`
}
