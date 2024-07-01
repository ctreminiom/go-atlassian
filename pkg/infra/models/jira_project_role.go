package models

// ProjectRoleDetailScheme represents the details of a project role in Jira.
type ProjectRoleDetailScheme struct {
	Self             string                         `json:"self,omitempty"`             // The URL of the project role.
	Name             string                         `json:"name,omitempty"`             // The name of the project role.
	ID               int                            `json:"id,omitempty"`               // The ID of the project role.
	Description      string                         `json:"description,omitempty"`      // The description of the project role.
	Admin            bool                           `json:"admin,omitempty"`            // Indicates if the project role has admin privileges.
	Scope            *TeamManagedProjectScopeScheme `json:"scope,omitempty"`            // The scope of the project role.
	RoleConfigurable bool                           `json:"roleConfigurable,omitempty"` // Indicates if the project role is configurable.
	TranslatedName   string                         `json:"translatedName,omitempty"`   // The translated name of the project role.
	Default          bool                           `json:"default,omitempty"`          // Indicates if the project role is the default role.
}

// ProjectRolePayloadScheme represents the payload for a project role in Jira.
type ProjectRolePayloadScheme struct {
	Name        string `json:"name,omitempty"`        // The name of the project role.
	Description string `json:"description,omitempty"` // The description of the project role.
}
