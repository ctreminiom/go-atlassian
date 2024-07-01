package models

// IssueVoteScheme represents the voting information for an issue in Jira.
type IssueVoteScheme struct {
	Self     string        `json:"self,omitempty"`     // The URL of the voting information.
	Votes    int           `json:"votes,omitempty"`    // The number of votes for the issue.
	HasVoted bool          `json:"hasVoted,omitempty"` // Indicates if the current user has voted for the issue.
	Voters   []*UserScheme `json:"voters,omitempty"`   // The users who have voted for the issue.
}
