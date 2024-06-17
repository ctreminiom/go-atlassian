package models

// IssueTypeScheme represents an issue type in Jira.
type IssueTypeScheme struct {
	Self           string                `json:"self,omitempty"`           // The URL of the issue type.
	ID             string                `json:"id,omitempty"`             // The ID of the issue type.
	Description    string                `json:"description,omitempty"`    // The description of the issue type.
	IconURL        string                `json:"iconUrl,omitempty"`        // The URL of the icon of the issue type.
	Name           string                `json:"name,omitempty"`           // The name of the issue type.
	Subtask        bool                  `json:"subtask,omitempty"`        // Indicates if the issue type is a subtask.
	AvatarID       int                   `json:"avatarId,omitempty"`       // The ID of the avatar of the issue type.
	EntityID       string                `json:"entityId,omitempty"`       // The entity ID of the issue type.
	HierarchyLevel int                   `json:"hierarchyLevel,omitempty"` // The hierarchy level of the issue type.
	Scope          *IssueTypeScopeScheme `json:"scope,omitempty"`          // The scope of the issue type.
}

// IssueTypeScopeScheme represents the scope of an issue type in Jira.
type IssueTypeScopeScheme struct {
	Type    string         `json:"type,omitempty"`    // The type of the scope.
	Project *ProjectScheme `json:"project,omitempty"` // The project of the scope.
}

// IssueTypePayloadScheme represents the payload for an issue type in Jira.
type IssueTypePayloadScheme struct {
	Name           string `json:"name,omitempty"`           // The name of the issue type.
	Description    string `json:"description,omitempty"`    // The description of the issue type.
	Type           string `json:"type,omitempty"`           // The type of the issue type.
	HierarchyLevel int    `json:"hierarchyLevel,omitempty"` // The hierarchy level of the issue type.
	AvatarID       int    `json:"avatarId,omitempty"`       // The ID of the avatar of the issue type.
}
