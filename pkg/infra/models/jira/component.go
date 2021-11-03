package jira

type ComponentPayloadScheme struct {
	IsAssigneeTypeValid bool   `json:"isAssigneeTypeValid,omitempty"`
	Name                string `json:"name,omitempty"`
	Description         string `json:"description,omitempty"`
	Project             string `json:"project,omitempty"`
	AssigneeType        string `json:"assigneeType,omitempty"`
	LeadAccountID       string `json:"leadAccountId,omitempty"`
}

type ComponentScheme struct {
	Self                string      `json:"self,omitempty"`
	ID                  string      `json:"id,omitempty"`
	Name                string      `json:"name,omitempty"`
	Description         string      `json:"description,omitempty"`
	Lead                *UserScheme `json:"lead,omitempty"`
	LeadUserName        string      `json:"leadUserName,omitempty"`
	AssigneeType        string      `json:"assigneeType,omitempty"`
	Assignee            *UserScheme `json:"assignee,omitempty"`
	RealAssigneeType    string      `json:"realAssigneeType,omitempty"`
	RealAssignee        *UserScheme `json:"realAssignee,omitempty"`
	IsAssigneeTypeValid bool        `json:"isAssigneeTypeValid,omitempty"`
	Project             string      `json:"project,omitempty"`
	ProjectID           int         `json:"projectId,omitempty"`
}

type ComponentCountScheme struct {
	Self       string `json:"self,omitempty"`
	IssueCount int    `json:"issueCount,omitempty"`
}
