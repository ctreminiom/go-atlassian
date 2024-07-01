package models

// PermissionScheme represents a permission scheme in Jira.
type PermissionScheme struct {
	Key         string `json:"key,omitempty"`         // The key of the permission scheme.
	Name        string `json:"name,omitempty"`        // The name of the permission scheme.
	Type        string `json:"type,omitempty"`        // The type of the permission scheme.
	Description string `json:"description,omitempty"` // The description of the permission scheme.
}

// PermissionCheckPayload represents the payload for a permission check in Jira.
type PermissionCheckPayload struct {
	GlobalPermissions  []string                        `json:"globalPermissions,omitempty"`  // The global permissions to check.
	AccountID          string                          `json:"accountId,omitempty"`          // The account ID to check the permissions for.
	ProjectPermissions []*BulkProjectPermissionsScheme `json:"projectPermissions,omitempty"` // The project permissions to check.
}

// BulkProjectPermissionsScheme represents a bulk of project permissions in Jira.
type BulkProjectPermissionsScheme struct {
	Issues      []int    `json:"issues"`      // The issues to check the permissions for.
	Projects    []int    `json:"projects"`    // The projects to check the permissions for.
	Permissions []string `json:"permissions"` // The permissions to check.
}

// PermissionGrantsScheme represents the grants of a permission in Jira.
type PermissionGrantsScheme struct {
	ProjectPermissions []*ProjectPermissionGrantsScheme `json:"projectPermissions,omitempty"` // The project permission grants.
	GlobalPermissions  []string                         `json:"globalPermissions,omitempty"`  // The global permission grants.
}

// ProjectPermissionGrantsScheme represents the grants of a project permission in Jira.
type ProjectPermissionGrantsScheme struct {
	Permission string `json:"permission,omitempty"` // The permission to grant.
	Issues     []int  `json:"issues,omitempty"`     // The issues to grant the permission for.
	Projects   []int  `json:"projects,omitempty"`   // The projects to grant the permission for.
}

// PermissionSchemeGrantsScheme represents the grants of a permission scheme in Jira.
type PermissionSchemeGrantsScheme struct {
	Permissions []*PermissionGrantScheme `json:"permissions,omitempty"` // The permission grants.
	Expand      string                   `json:"expand,omitempty"`      // The expand option for the permission scheme grants.
}

// PermissionGrantScheme represents a grant of a permission in Jira.
type PermissionGrantScheme struct {
	ID         int                          `json:"id,omitempty"`         // The ID of the permission grant.
	Self       string                       `json:"self,omitempty"`       // The URL of the permission grant.
	Holder     *PermissionGrantHolderScheme `json:"holder,omitempty"`     // The holder of the permission grant.
	Permission string                       `json:"permission,omitempty"` // The permission to grant.
}

// PermissionGrantHolderScheme represents a holder of a permission grant in Jira.
type PermissionGrantHolderScheme struct {
	Type      string `json:"type,omitempty"`      // The type of the holder.
	Parameter string `json:"parameter,omitempty"` // The parameter of the holder.
	Expand    string `json:"expand,omitempty"`    // The expand option for the holder.
}

// PermissionGrantPayloadScheme represents the payload for a permission grant in Jira.
type PermissionGrantPayloadScheme struct {
	Holder     *PermissionGrantHolderScheme `json:"holder,omitempty"`     // The holder of the permission grant.
	Permission string                       `json:"permission,omitempty"` // The permission to grant.
}

// PermittedProjectsScheme represents the permitted projects of a permission in Jira.
type PermittedProjectsScheme struct {
	Projects []*ProjectIdentifierScheme `json:"projects,omitempty"` // The permitted projects.
}
