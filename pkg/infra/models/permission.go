package models

type PermissionScheme struct {
	Key         string `json:"key,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

type PermissionCheckPayload struct {
	GlobalPermissions  []string                        `json:"globalPermissions,omitempty"`
	AccountID          string                          `json:"accountId,omitempty"`
	ProjectPermissions []*BulkProjectPermissionsScheme `json:"projectPermissions,omitempty"`
}

type BulkProjectPermissionsScheme struct {
	Issues      []int    `json:"issues"`
	Projects    []int    `json:"projects"`
	Permissions []string `json:"permissions"`
}

type PermissionGrantsScheme struct {
	ProjectPermissions []struct {
		Permission string `json:"permission,omitempty"`
		Issues     []int  `json:"issues,omitempty"`
		Projects   []int  `json:"projects,omitempty"`
	} `json:"projectPermissions,omitempty"`
	GlobalPermissions []string `json:"globalPermissions,omitempty"`
}

type PermissionSchemeGrantsScheme struct {
	Permissions []*PermissionGrantScheme `json:"permissions,omitempty"`
	Expand      string                   `json:"expand,omitempty"`
}

type PermissionGrantScheme struct {
	ID         int                          `json:"id,omitempty"`
	Self       string                       `json:"self,omitempty"`
	Holder     *PermissionGrantHolderScheme `json:"holder,omitempty"`
	Permission string                       `json:"permission,omitempty"`
}

type PermissionGrantHolderScheme struct {
	Type      string `json:"type,omitempty"`
	Parameter string `json:"parameter,omitempty"`
	Expand    string `json:"expand,omitempty"`
}

type PermissionGrantPayloadScheme struct {
	Holder     *PermissionGrantHolderScheme `json:"holder,omitempty"`
	Permission string                       `json:"permission,omitempty"`
}
