package models

// SpaceScheme represents a space in Confluence.
type SpaceScheme struct {
	ID          int                      `json:"id,omitempty"`
	Key         string                   `json:"key,omitempty"`
	Name        string                   `json:"name,omitempty"`
	Type        string                   `json:"type,omitempty"`
	Status      string                   `json:"status,omitempty"`
	HomePage    *ContentScheme           `json:"homepage,omitempty"`
	Operations  *[]SpaceOperationScheme  `json:"operations,omitempty"`
	Permissions []*SpacePermissionScheme `json:"permissions,omitempty"`
	History     *SpaceHistoryScheme      `json:"history,omitempty"`
	Expandable  *ExpandableScheme        `json:"_expandable,omitempty"`
	Links       *LinkScheme              `json:"_links,omitempty"`
}

// SpaceHistoryScheme represents the history of a space in Confluence.
type SpaceHistoryScheme struct {
	CreatedDate string             `json:"createdDate,omitempty"`
	CreatedBy   *ContentUserScheme `json:"createdBy,omitempty"`
}

// SpaceOperationScheme represents an operation in a space in Confluence.
type SpaceOperationScheme struct {
	Operation  string `json:"operation,omitempty"`
	TargetType string `json:"targetType,omitempty"`
}

// SpacePageScheme represents a page of spaces in Confluence.
type SpacePageScheme struct {
	Results []*SpaceScheme `json:"results,omitempty"`
	Start   int            `json:"start"`
	Limit   int            `json:"limit"`
	Size    int            `json:"size"`
	Links   struct {
		Base    string `json:"base"`
		Context string `json:"context"`
		Self    string `json:"self"`
	} `json:"_links"`
}

// GetSpacesOptionScheme represents the options for getting spaces in Confluence.
type GetSpacesOptionScheme struct {
	SpaceKeys       []string
	SpaceIDs        []int
	SpaceType       string
	Status          string
	Labels          []string
	Favorite        bool
	FavoriteUserKey string
	Expand          []string
}

// CreateSpaceScheme represents the scheme for creating a space in Confluence.
type CreateSpaceScheme struct {
	Key              string                        `json:"key,omitempty"`
	Name             string                        `json:"name,omitempty"`
	Description      *CreateSpaceDescriptionScheme `json:"description,omitempty"`
	AnonymousAccess  bool                          `json:"anonymousAccess,omitempty"`
	UnlicensedAccess bool                          `json:"unlicensedAccess,omitempty"`
}

// CreateSpaceDescriptionScheme represents the description of a space in Confluence.
type CreateSpaceDescriptionScheme struct {
	Plain *CreateSpaceDescriptionPlainScheme `json:"plain"`
}

// CreateSpaceDescriptionPlainScheme represents the plain text description of a space in Confluence.
type CreateSpaceDescriptionPlainScheme struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}

// SpacePermissionScheme represents a permission in a space in Confluence.
type SpacePermissionScheme struct {
	Subject          *SubjectPermissionScheme   `json:"subjects,omitempty"`
	Operation        *OperationPermissionScheme `json:"operation,omitempty"`
	AnonymousAccess  bool                       `json:"anonymousAccess,omitempty"`
	UnlicensedAccess bool                       `json:"unlicensedAccess,omitempty"`
}

// OperationPermissionScheme represents an operation in a permission in Confluence.
type OperationPermissionScheme struct {
	Operation  string `json:"operation,omitempty"`
	TargetType string `json:"targetType,omitempty"`
}

// SubjectPermissionScheme represents a subject in a permission in Confluence.
type SubjectPermissionScheme struct {
	User  *UserPermissionScheme  `json:"user,omitempty"`
	Group *GroupPermissionScheme `json:"group,omitempty"`
}

// UserPermissionScheme represents a user in a subject in a permission in Confluence.
type UserPermissionScheme struct {
	Results []*ContentUserScheme `json:"results,omitempty"`
	Size    int                  `json:"size,omitempty"`
}

// GroupPermissionScheme represents a group in a subject in a permission in Confluence.
type GroupPermissionScheme struct {
	Results []*SpaceGroupScheme `json:"results,omitempty"`
	Size    int                 `json:"size,omitempty"`
}

// SpaceGroupScheme represents a group in a subject in a permission in Confluence.
type SpaceGroupScheme struct {
	Type  string      `json:"type,omitempty"`
	Name  string      `json:"name,omitempty"`
	ID    string      `json:"id,omitempty"`
	Links *LinkScheme `json:"_links,omitempty"`
}

// UpdateSpaceScheme represents the scheme for updating a space in Confluence.
type UpdateSpaceScheme struct {
	Name        string                        `json:"name,omitempty"`
	Description *CreateSpaceDescriptionScheme `json:"description,omitempty"`
	Homepage    *UpdateSpaceHomepageScheme    `json:"homepage,omitempty"`
}

// UpdateSpaceHomepageScheme represents the home page of a space in Confluence.
type UpdateSpaceHomepageScheme struct {
	ID string `json:"id"`
}
