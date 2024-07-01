package models

// PriorityScheme represents a priority in Jira.
type PriorityScheme struct {
	Self        string `json:"self,omitempty"`        // The URL of the priority.
	StatusColor string `json:"statusColor,omitempty"` // The status color of the priority.
	Description string `json:"description,omitempty"` // The description of the priority.
	IconURL     string `json:"iconUrl,omitempty"`     // The URL of the icon for the priority.
	Name        string `json:"name,omitempty"`        // The name of the priority.
	ID          string `json:"id,omitempty"`          // The ID of the priority.
}
