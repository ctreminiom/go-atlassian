package models

// ProjectCategoryPayloadScheme represents the payload for a project category in Jira.
type ProjectCategoryPayloadScheme struct {
	Name        string `json:"name,omitempty"`        // The name of the project category.
	Description string `json:"description,omitempty"` // The description of the project category.
}
