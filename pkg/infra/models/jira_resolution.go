package models

// ResolutionScheme represents a resolution in Jira.
type ResolutionScheme struct {
	Self        string `json:"self"`        // The URL of the resolution.
	ID          string `json:"id"`          // The ID of the resolution.
	Description string `json:"description"` // The description of the resolution.
	Name        string `json:"name"`        // The name of the resolution.
}
