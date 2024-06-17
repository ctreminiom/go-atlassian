package models

// DashboardPageScheme represents a page of dashboards in Jira.
type DashboardPageScheme struct {
	StartAt    int                `json:"startAt,omitempty"`    // The starting index of the page.
	MaxResults int                `json:"maxResults,omitempty"` // The maximum number of results in the page.
	Total      int                `json:"total,omitempty"`      // The total number of dashboards.
	Dashboards []*DashboardScheme `json:"dashboards,omitempty"` // The dashboards in the page.
}

// DashboardScheme represents a dashboard in Jira.
type DashboardScheme struct {
	ID               string                   `json:"id,omitempty"`               // The ID of the dashboard.
	IsFavourite      bool                     `json:"isFavourite,omitempty"`      // Indicates if the dashboard is a favourite.
	Name             string                   `json:"name,omitempty"`             // The name of the dashboard.
	Owner            *UserScheme              `json:"owner,omitempty"`            // The owner of the dashboard.
	Popularity       int                      `json:"popularity,omitempty"`       // The popularity of the dashboard.
	Rank             int                      `json:"rank,omitempty"`             // The rank of the dashboard.
	Self             string                   `json:"self,omitempty"`             // The URL of the dashboard.
	SharePermissions []*SharePermissionScheme `json:"sharePermissions,omitempty"` // The share permissions of the dashboard.
	EditPermission   []*SharePermissionScheme `json:"editPermissions,omitempty"`  // The edit permissions of the dashboard.
	View             string                   `json:"view,omitempty"`             // The view of the dashboard.
}

// SharePermissionScheme represents a share permission in Jira.
type SharePermissionScheme struct {
	ID      int                `json:"id,omitempty"`      // The ID of the share permission.
	Type    string             `json:"type,omitempty"`    // The type of the share permission.
	Project *ProjectScheme     `json:"project,omitempty"` // The project of the share permission.
	Role    *ProjectRoleScheme `json:"role,omitempty"`    // The role of the share permission.
	Group   *GroupScheme       `json:"group,omitempty"`   // The group of the share permission.
	User    *UserDetailScheme  `json:"user,omitempty"`    // The user of the share permission.
}

// DashboardSearchPageScheme represents a search page of dashboards in Jira.
type DashboardSearchPageScheme struct {
	Self       string             `json:"self,omitempty"`       // The URL of the search page.
	MaxResults int                `json:"maxResults,omitempty"` // The maximum number of results in the search page.
	StartAt    int                `json:"startAt,omitempty"`    // The starting index of the search page.
	Total      int                `json:"total,omitempty"`      // The total number of dashboards in the search page.
	IsLast     bool               `json:"isLast,omitempty"`     // Indicates if the search page is the last one.
	Values     []*DashboardScheme `json:"values,omitempty"`     // The dashboards in the search page.
}

// DashboardPayloadScheme represents the payload for a dashboard in Jira.
type DashboardPayloadScheme struct {
	Name             string                   `json:"name,omitempty"`             // The name of the dashboard.
	Description      string                   `json:"description,omitempty"`      // The description of the dashboard.
	SharePermissions []*SharePermissionScheme `json:"sharePermissions,omitempty"` // The share permissions of the dashboard.
	EditPermissions  []*SharePermissionScheme `json:"editPermissions,omitempty"`  // The edit permissions of the dashboard.
}

// DashboardSearchOptionsScheme represents the search options for a dashboard in Jira.
type DashboardSearchOptionsScheme struct {
	DashboardName       string   // The name of the dashboard.
	OwnerAccountID      string   // The account ID of the owner of the dashboard.
	GroupPermissionName string   // The name of the group permission of the dashboard.
	OrderBy             string   // The order by criteria of the dashboard.
	Expand              []string // The fields to be expanded in the dashboard.
}
