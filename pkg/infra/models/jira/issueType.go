package jira

type IssueTypeScheme struct {
	Self           string                `json:"self,omitempty"`
	ID             string                `json:"id,omitempty"`
	Description    string                `json:"description,omitempty"`
	IconURL        string                `json:"iconUrl,omitempty"`
	Name           string                `json:"name,omitempty"`
	Subtask        bool                  `json:"subtask,omitempty"`
	AvatarID       int                   `json:"avatarId,omitempty"`
	EntityID       string                `json:"entityId,omitempty"`
	HierarchyLevel int                   `json:"hierarchyLevel,omitempty"`
	Scope          *IssueTypeScopeScheme `json:"scope,omitempty"`
}

type IssueTypeScopeScheme struct {
	Type    string         `json:"type,omitempty"`
	Project *ProjectScheme `json:"project,omitempty"`
}

type IssueTypePayloadScheme struct {
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	Type           string `json:"type,omitempty"`
	HierarchyLevel int    `json:"hierarchyLevel,omitempty"`
	AvatarID       int    `json:"avatarId,omitempty"`
}
