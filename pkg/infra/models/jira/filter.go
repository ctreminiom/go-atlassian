package jira

type FilterPageScheme struct {
	Self       string          `json:"self,omitempty"`
	MaxResults int             `json:"maxResults,omitempty"`
	StartAt    int             `json:"startAt,omitempty"`
	Total      int             `json:"total,omitempty"`
	IsLast     bool            `json:"isLast,omitempty"`
	Values     []*FilterScheme `json:"values,omitempty"`
}

type FilterSearchPageScheme struct {
	Self       string                `json:"self,omitempty"`
	MaxResults int                   `json:"maxResults,omitempty"`
	StartAt    int                   `json:"startAt,omitempty"`
	Total      int                   `json:"total,omitempty"`
	IsLast     bool                  `json:"isLast,omitempty"`
	Values     []*FilterDetailScheme `json:"values,omitempty"`
}

type FilterDetailScheme struct {
	Self             string                      `json:"self,omitempty"`
	ID               string                      `json:"id,omitempty"`
	Name             string                      `json:"name,omitempty"`
	Owner            *UserScheme                 `json:"owner,omitempty"`
	Jql              string                      `json:"jql,omitempty"`
	ViewURL          string                      `json:"viewUrl,omitempty"`
	SearchURL        string                      `json:"searchUrl,omitempty"`
	Favourite        bool                        `json:"favourite,omitempty"`
	FavouritedCount  int                         `json:"favouritedCount,omitempty"`
	SharePermissions []*SharePermissionScheme    `json:"sharePermissions,omitempty"`
	Subscriptions    []*FilterSubscriptionScheme `json:"subscriptions,omitempty"`
}

type FilterScheme struct {
	Self             string                        `json:"self,omitempty"`
	ID               string                        `json:"id,omitempty"`
	Name             string                        `json:"name,omitempty"`
	Owner            *UserScheme                   `json:"owner,omitempty"`
	Jql              string                        `json:"jql,omitempty"`
	ViewURL          string                        `json:"viewUrl,omitempty"`
	SearchURL        string                        `json:"searchUrl,omitempty"`
	Favourite        bool                          `json:"favourite,omitempty"`
	FavouritedCount  int                           `json:"favouritedCount,omitempty"`
	SharePermissions []*SharePermissionScheme      `json:"sharePermissions,omitempty"`
	ShareUsers       *FilterUsersScheme            `json:"sharedUsers,omitempty"`
	Subscriptions    *FilterSubscriptionPageScheme `json:"subscriptions,omitempty"`
}

type FilterSubscriptionPageScheme struct {
	Size       int                         `json:"size,omitempty"`
	Items      []*FilterSubscriptionScheme `json:"items,omitempty"`
	MaxResults int                         `json:"max-results,omitempty"`
	StartIndex int                         `json:"start-index,omitempty"`
	EndIndex   int                         `json:"end-index,omitempty"`
}

type FilterSubscriptionScheme struct {
	ID    int          `json:"id,omitempty"`
	User  *UserScheme  `json:"user,omitempty"`
	Group *GroupScheme `json:"group,omitempty"`
}

type FilterUsersScheme struct {
	Size       int           `json:"size,omitempty"`
	Items      []*UserScheme `json:"items,omitempty"`
	MaxResults int           `json:"max-results,omitempty"`
	StartIndex int           `json:"start-index,omitempty"`
	EndIndex   int           `json:"end-index,omitempty"`
}
