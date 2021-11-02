package jira

type IssueVoteScheme struct {
	Self     string        `json:"self,omitempty"`
	Votes    int           `json:"votes,omitempty"`
	HasVoted bool          `json:"hasVoted,omitempty"`
	Voters   []*UserScheme `json:"voters,omitempty"`
}
