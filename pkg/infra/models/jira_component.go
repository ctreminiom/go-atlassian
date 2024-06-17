package models

// ComponentPayloadScheme represents the payload for a component in Jira.
type ComponentPayloadScheme struct {
	IsAssigneeTypeValid bool   `json:"isAssigneeTypeValid,omitempty"` // Indicates if the assignee type is valid.
	Name                string `json:"name,omitempty"`                // The name of the component.
	Description         string `json:"description,omitempty"`         // The description of the component.
	Project             string `json:"project,omitempty"`             // The project to which the component belongs.
	AssigneeType        string `json:"assigneeType,omitempty"`        // The type of the assignee.
	LeadAccountID       string `json:"leadAccountId,omitempty"`       // The account ID of the lead.
}

// ComponentScheme represents a component in Jira.
type ComponentScheme struct {
	Self                string      `json:"self,omitempty"`                // The URL of the component.
	ID                  string      `json:"id,omitempty"`                  // The ID of the component.
	Name                string      `json:"name,omitempty"`                // The name of the component.
	Description         string      `json:"description,omitempty"`         // The description of the component.
	Lead                *UserScheme `json:"lead,omitempty"`                // The lead of the component.
	LeadUserName        string      `json:"leadUserName,omitempty"`        // The username of the lead.
	AssigneeType        string      `json:"assigneeType,omitempty"`        // The type of the assignee.
	Assignee            *UserScheme `json:"assignee,omitempty"`            // The assignee of the component.
	RealAssigneeType    string      `json:"realAssigneeType,omitempty"`    // The real type of the assignee.
	RealAssignee        *UserScheme `json:"realAssignee,omitempty"`        // The real assignee of the component.
	IsAssigneeTypeValid bool        `json:"isAssigneeTypeValid,omitempty"` // Indicates if the assignee type is valid.
	Project             string      `json:"project,omitempty"`             // The project to which the component belongs.
	ProjectID           int         `json:"projectId,omitempty"`           // The ID of the project to which the component belongs.
}

// ComponentCountScheme represents the count of components in Jira.
type ComponentCountScheme struct {
	Self       string `json:"self,omitempty"`       // The URL of the component count.
	IssueCount int    `json:"issueCount,omitempty"` // The count of issues in the component.
}
