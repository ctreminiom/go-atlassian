package models

// PermissionSchemePageScheme represents a page of permission schemes in Jira.
type PermissionSchemePageScheme struct {
	PermissionSchemes []*PermissionSchemeScheme `json:"permissionSchemes,omitempty"` // The permission schemes in the page.
}

// PermissionSchemeScheme represents a permission scheme in Jira.
type PermissionSchemeScheme struct {
	Expand      string                         `json:"expand,omitempty"`      // The expand option for the permission scheme.
	ID          int                            `json:"id,omitempty"`          // The ID of the permission scheme.
	Self        string                         `json:"self,omitempty"`        // The URL of the permission scheme.
	Name        string                         `json:"name,omitempty"`        // The name of the permission scheme.
	Description string                         `json:"description,omitempty"` // The description of the permission scheme.
	Permissions []*PermissionGrantScheme       `json:"permissions,omitempty"` // The permissions of the permission scheme.
	Scope       *TeamManagedProjectScopeScheme `json:"scope,omitempty"`       // The scope of the permission scheme.
}

// PermissionScopeItemScheme represents an item in the scope of a permission in Jira.
type PermissionScopeItemScheme struct {
	Self            string                 `json:"self,omitempty"`            // The URL of the item.
	ID              string                 `json:"id,omitempty"`              // The ID of the item.
	Key             string                 `json:"key,omitempty"`             // The key of the item.
	Name            string                 `json:"name,omitempty"`            // The name of the item.
	ProjectTypeKey  string                 `json:"projectTypeKey,omitempty"`  // The project type key of the item.
	Simplified      bool                   `json:"simplified,omitempty"`      // Indicates if the item is simplified.
	ProjectCategory *ProjectCategoryScheme `json:"projectCategory,omitempty"` // The project category of the item.
}
