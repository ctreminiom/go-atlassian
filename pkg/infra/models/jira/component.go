package jira

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
