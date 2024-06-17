package models

// SecurityScheme represents a security scheme in Jira.
type SecurityScheme struct {
	Self        string `json:"self,omitempty"`        // The URL of the security scheme.
	ID          string `json:"id,omitempty"`          // The ID of the security scheme.
	Name        string `json:"name,omitempty"`        // The name of the security scheme.
	Description string `json:"description,omitempty"` // The description of the security scheme.
}
