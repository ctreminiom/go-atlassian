package models

type PermissionSchemePageScheme struct {
	PermissionSchemes []*PermissionSchemeScheme `json:"permissionSchemes,omitempty"`
}

type PermissionSchemeScheme struct {
	Expand      string                         `json:"expand,omitempty"`
	ID          int                            `json:"id,omitempty"`
	Self        string                         `json:"self,omitempty"`
	Name        string                         `json:"name,omitempty"`
	Description string                         `json:"description,omitempty"`
	Permissions []*PermissionGrantScheme       `json:"permissions,omitempty"`
	Scope       *TeamManagedProjectScopeScheme `json:"scope,omitempty"`
}

type PermissionScopeItemScheme struct {
	Self            string                 `json:"self,omitempty"`
	ID              string                 `json:"id,omitempty"`
	Key             string                 `json:"key,omitempty"`
	Name            string                 `json:"name,omitempty"`
	ProjectTypeKey  string                 `json:"projectTypeKey,omitempty"`
	Simplified      bool                   `json:"simplified,omitempty"`
	ProjectCategory *ProjectCategoryScheme `json:"projectCategory,omitempty"`
}
