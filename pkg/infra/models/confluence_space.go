package models

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

type SpaceHistoryScheme struct {
	CreatedDate string             `json:"createdDate,omitempty"`
	CreatedBy   *ContentUserScheme `json:"createdBy,omitempty"`
}

type SpaceOperationScheme struct {
	Operation  string `json:"operation,omitempty"`
	TargetType string `json:"targetType,omitempty"`
}

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

type CreateSpaceScheme struct {
	Key              string                        `json:"key,omitempty"`
	Name             string                        `json:"name,omitempty"`
	Description      *CreateSpaceDescriptionScheme `json:"description,omitempty"`
	AnonymousAccess  bool                          `json:"anonymousAccess,omitempty"`
	UnlicensedAccess bool                          `json:"unlicensedAccess,omitempty"`
}

type CreateSpaceDescriptionScheme struct {
	Plain *CreateSpaceDescriptionPlainScheme `json:"plain"`
}

type CreateSpaceDescriptionPlainScheme struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}

type SpacePermissionScheme struct {
	Subject          *SubjectPermissionScheme   `json:"subjects,omitempty"`
	Operation        *OperationPermissionScheme `json:"operation,omitempty"`
	AnonymousAccess  bool                       `json:"anonymousAccess,omitempty"`
	UnlicensedAccess bool                       `json:"unlicensedAccess,omitempty"`
}

type OperationPermissionScheme struct {
	Operation  string `json:"operation,omitempty"`
	TargetType string `json:"targetType,omitempty"`
}

type SubjectPermissionScheme struct {
	User  *UserPermissionScheme  `json:"user,omitempty"`
	Group *GroupPermissionScheme `json:"group,omitempty"`
}

type UserPermissionScheme struct {
	Results []*ContentUserScheme `json:"results,omitempty"`
	Size    int                  `json:"size,omitempty"`
}

type GroupPermissionScheme struct {
	Results []*SpaceGroupScheme `json:"results,omitempty"`
	Size    int                 `json:"size,omitempty"`
}

type SpaceGroupScheme struct {
	Type  string      `json:"type,omitempty"`
	Name  string      `json:"name,omitempty"`
	ID    string      `json:"id,omitempty"`
	Links *LinkScheme `json:"_links,omitempty"`
}

type UpdateSpaceScheme struct {
	Name        string                        `json:"name,omitempty"`
	Description *CreateSpaceDescriptionScheme `json:"description,omitempty"`
	Homepage    *UpdateSpaceHomepageScheme    `json:"homepage,omitempty"`
}

type UpdateSpaceHomepageScheme struct {
	ID string `json:"id"`
}
