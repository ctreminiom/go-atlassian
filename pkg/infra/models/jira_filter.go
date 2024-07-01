package models

// FilterPageScheme represents a page of filters in Jira.
type FilterPageScheme struct {
	Self       string          `json:"self,omitempty"`
	MaxResults int             `json:"maxResults,omitempty"`
	StartAt    int             `json:"startAt,omitempty"`
	Total      int             `json:"total,omitempty"`
	IsLast     bool            `json:"isLast,omitempty"`
	Values     []*FilterScheme `json:"values,omitempty"`
}

// FilterSearchPageScheme represents a page of filter details in Jira.
type FilterSearchPageScheme struct {
	Self       string                `json:"self,omitempty"`
	MaxResults int                   `json:"maxResults,omitempty"`
	StartAt    int                   `json:"startAt,omitempty"`
	Total      int                   `json:"total,omitempty"`
	IsLast     bool                  `json:"isLast,omitempty"`
	Values     []*FilterDetailScheme `json:"values,omitempty"`
}

// FilterDetailScheme represents the details of a filter in Jira.
type FilterDetailScheme struct {
	Self             string                      `json:"self,omitempty"`
	ID               string                      `json:"id,omitempty"`
	Name             string                      `json:"name,omitempty"`
	Owner            *UserScheme                 `json:"owner,omitempty"`
	JQL              string                      `json:"jql,omitempty"`
	ViewURL          string                      `json:"viewUrl,omitempty"`
	SearchURL        string                      `json:"searchUrl,omitempty"`
	Favourite        bool                        `json:"favourite,omitempty"`
	FavouritedCount  int                         `json:"favouritedCount,omitempty"`
	SharePermissions []*SharePermissionScheme    `json:"sharePermissions,omitempty"`
	Subscriptions    []*FilterSubscriptionScheme `json:"subscriptions,omitempty"`
}

// FilterScheme represents a filter in Jira.
type FilterScheme struct {
	Self             string                        `json:"self,omitempty"`
	ID               string                        `json:"id,omitempty"`
	Name             string                        `json:"name,omitempty"`
	Owner            *UserScheme                   `json:"owner,omitempty"`
	JQL              string                        `json:"jql,omitempty"`
	ViewURL          string                        `json:"viewUrl,omitempty"`
	SearchURL        string                        `json:"searchUrl,omitempty"`
	Favourite        bool                          `json:"favourite,omitempty"`
	FavouritedCount  int                           `json:"favouritedCount,omitempty"`
	SharePermissions []*SharePermissionScheme      `json:"sharePermissions,omitempty"`
	ShareUsers       *FilterUsersScheme            `json:"sharedUsers,omitempty"`
	Subscriptions    *FilterSubscriptionPageScheme `json:"subscriptions,omitempty"`
}

// FilterSubscriptionPageScheme represents a page of filter subscriptions in Jira.
type FilterSubscriptionPageScheme struct {
	Size       int                         `json:"size,omitempty"`
	Items      []*FilterSubscriptionScheme `json:"items,omitempty"`
	MaxResults int                         `json:"max-results,omitempty"`
	StartIndex int                         `json:"start-index,omitempty"`
	EndIndex   int                         `json:"end-index,omitempty"`
}

// FilterSubscriptionScheme represents a filter subscription in Jira.
type FilterSubscriptionScheme struct {
	ID    int          `json:"id,omitempty"`
	User  *UserScheme  `json:"user,omitempty"`
	Group *GroupScheme `json:"group,omitempty"`
}

// FilterUsersScheme represents the users of a filter in Jira.
type FilterUsersScheme struct {
	Size       int           `json:"size,omitempty"`
	Items      []*UserScheme `json:"items,omitempty"`
	MaxResults int           `json:"max-results,omitempty"`
	StartIndex int           `json:"start-index,omitempty"`
	EndIndex   int           `json:"end-index,omitempty"`
}

// FilterPayloadScheme represents the payload for a filter in Jira.
type FilterPayloadScheme struct {
	Name             string                   `json:"name,omitempty"`
	Description      string                   `json:"description,omitempty"`
	JQL              string                   `json:"jql,omitempty"`
	Favorite         bool                     `json:"favourite,omitempty"`
	SharePermissions []*SharePermissionScheme `json:"sharePermissions,omitempty"`
	EditPermissions  []*SharePermissionScheme `json:"editPermissions,omitempty"`
}

// FilterSearchOptionScheme represents the search options for a filter in Jira.
type FilterSearchOptionScheme struct {
	Name      string
	AccountID string
	Group     string
	OrderBy   string
	ProjectID int
	IDs       []int
	Expand    []string
}

// ShareFilterScopeScheme represents the scope of a shared filter in Jira.
type ShareFilterScopeScheme struct {
	Scope string `json:"scope"`
}

// PermissionFilterPayloadScheme represents the payload for a permission filter in Jira.
type PermissionFilterPayloadScheme struct {
	Type          string `json:"type,omitempty"`
	ProjectID     string `json:"projectId,omitempty"`
	GroupName     string `json:"groupname,omitempty"`
	ProjectRoleID string `json:"projectRoleId,omitempty"`
}
