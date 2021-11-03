package jira

type ProjectRoleDetailScheme struct {
	Self             string                         `json:"self,omitempty"`
	Name             string                         `json:"name,omitempty"`
	ID               int                            `json:"id,omitempty"`
	Description      string                         `json:"description,omitempty"`
	Admin            bool                           `json:"admin,omitempty"`
	Scope            *TeamManagedProjectScopeScheme `json:"scope,omitempty"`
	RoleConfigurable bool                           `json:"roleConfigurable,omitempty"`
	TranslatedName   string                         `json:"translatedName,omitempty"`
	Default          bool                           `json:"default,omitempty"`
}

type ProjectRolePayloadScheme struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
