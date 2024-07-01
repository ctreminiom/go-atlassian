package models

// IssueSecurityLevelsScheme represents the security levels of an issue in Jira.
type IssueSecurityLevelsScheme struct {
	Levels []*IssueSecurityLevelScheme `json:"levels,omitempty"` // The security levels of the issue.
}

// IssueSecurityLevelScheme represents a security level of an issue in Jira.
type IssueSecurityLevelScheme struct {
	Self        string `json:"self,omitempty"`        // The URL of the security level.
	ID          string `json:"id,omitempty"`          // The ID of the security level.
	Description string `json:"description,omitempty"` // The description of the security level.
	Name        string `json:"name,omitempty"`        // The name of the security level.
}
