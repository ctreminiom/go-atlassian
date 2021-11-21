package models

type IssueSecurityLevelsScheme struct {
	Levels []*IssueSecurityLevelScheme `json:"levels,omitempty"`
}

type IssueSecurityLevelScheme struct {
	Self        string `json:"self,omitempty"`
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}
