package models

// IssueSearchApproximateCountScheme represents the response from the approximate count endpoint
type IssueSearchApproximateCountScheme struct {
	Count int `json:"count,omitempty"`
}
