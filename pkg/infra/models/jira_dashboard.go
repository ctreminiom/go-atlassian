package models

type DashboardPageScheme struct {
	StartAt    int                `json:"startAt,omitempty"`
	MaxResults int                `json:"maxResults,omitempty"`
	Total      int                `json:"total,omitempty"`
	Dashboards []*DashboardScheme `json:"dashboards,omitempty"`
}

type DashboardScheme struct {
	ID               string                   `json:"id,omitempty"`
	IsFavourite      bool                     `json:"isFavourite,omitempty"`
	Name             string                   `json:"name,omitempty"`
	Owner            *UserScheme              `json:"owner,omitempty"`
	Popularity       int                      `json:"popularity,omitempty"`
	Rank             int                      `json:"rank,omitempty"`
	Self             string                   `json:"self,omitempty"`
	SharePermissions []*SharePermissionScheme `json:"sharePermissions,omitempty"`
	EditPermission   []*SharePermissionScheme `json:"editPermissions,omitempty"`
	View             string                   `json:"view,omitempty"`
}

type SharePermissionScheme struct {
	ID      int                `json:"id,omitempty"`
	Type    string             `json:"type,omitempty"`
	Project *ProjectScheme     `json:"project,omitempty"`
	Role    *ProjectRoleScheme `json:"role,omitempty"`
	Group   *GroupScheme       `json:"group,omitempty"`
}

type DashboardSearchPageScheme struct {
	Self       string             `json:"self,omitempty"`
	MaxResults int                `json:"maxResults,omitempty"`
	StartAt    int                `json:"startAt,omitempty"`
	Total      int                `json:"total,omitempty"`
	IsLast     bool               `json:"isLast,omitempty"`
	Values     []*DashboardScheme `json:"values,omitempty"`
}

type DashboardPayloadScheme struct {
	Name             string                   `json:"name,omitempty"`
	Description      string                   `json:"description,omitempty"`
	SharePermissions []*SharePermissionScheme `json:"sharePermissions,omitempty"`
	EditPermissions  []*SharePermissionScheme `json:"editPermissions,omitempty"`
}

type DashboardSearchOptionsScheme struct {
	DashboardName       string
	OwnerAccountID      string
	GroupPermissionName string
	OrderBy             string
	Expand              []string
}
