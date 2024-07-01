package models

// UserGroupsScheme represents a collection of user groups in Jira.
type UserGroupsScheme struct {
	Size       int                `json:"size,omitempty"`        // The size of the collection.
	Items      []*UserGroupScheme `json:"items,omitempty"`       // The user groups.
	MaxResults int                `json:"max-results,omitempty"` // The maximum number of results in the collection.
}

// UserGroupScheme represents a user group in Jira.
type UserGroupScheme struct {
	Name string `json:"name,omitempty"` // The name of the user group.
	Self string `json:"self,omitempty"` // The URL of the user group.
}

// GroupScheme represents a group in Jira.
type GroupScheme struct {
	Name   string               `json:"name,omitempty"`   // The name of the group.
	Self   string               `json:"self,omitempty"`   // The URL of the group.
	Users  *GroupUserPageScheme `json:"users,omitempty"`  // The users in the group.
	Expand string               `json:"expand,omitempty"` // The fields to be expanded in the group.
}

// GroupUserPageScheme represents a page of users in a group in Jira.
type GroupUserPageScheme struct {
	Size       int           `json:"size,omitempty"`        // The size of the page.
	Items      []*UserScheme `json:"items,omitempty"`       // The users in the page.
	MaxResults int           `json:"max-results,omitempty"` // The maximum number of results in the page.
	StartIndex int           `json:"start-index,omitempty"` // The starting index of the page.
	EndIndex   int           `json:"end-index,omitempty"`   // The ending index of the page.
}

// BulkGroupScheme represents a bulk of groups in Jira.
type BulkGroupScheme struct {
	MaxResults int                  `json:"maxResults,omitempty"` // The maximum number of results in the bulk.
	StartAt    int                  `json:"startAt,omitempty"`    // The starting index of the bulk.
	Total      int                  `json:"total,omitempty"`      // The total number of groups in the bulk.
	IsLast     bool                 `json:"isLast,omitempty"`     // Indicates if the bulk is the last one.
	Values     []*GroupDetailScheme `json:"values,omitempty"`     // The groups in the bulk.
}

// GroupDetailScheme represents the details of a group in Jira.
type GroupDetailScheme struct {
	Self    string `json:"self,omitempty"`    // The URL of the group.
	Name    string `json:"name,omitempty"`    // The name of the group.
	GroupID string `json:"groupId,omitempty"` // The ID of the group.
}

// GroupMemberPageScheme represents a page of group members in Jira.
type GroupMemberPageScheme struct {
	Self       string                   `json:"self,omitempty"`       // The URL of the page.
	NextPage   string                   `json:"nextPage,omitempty"`   // The URL of the next page.
	MaxResults int                      `json:"maxResults,omitempty"` // The maximum number of results in the page.
	StartAt    int                      `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                      `json:"total,omitempty"`      // The total number of group members.
	IsLast     bool                     `json:"isLast,omitempty"`     // Indicates if the page is the last one.
	Values     []*GroupUserDetailScheme `json:"values,omitempty"`     // The group members in the page.
}

// GroupUserDetailScheme represents the details of a group user in Jira.
type GroupUserDetailScheme struct {
	Self         string `json:"self"`         // The URL of the group user.
	Name         string `json:"name"`         // The name of the group user.
	Key          string `json:"key"`          // The key of the group user.
	AccountID    string `json:"accountId"`    // The account ID of the group user.
	EmailAddress string `json:"emailAddress"` // The email address of the group user.
	DisplayName  string `json:"displayName"`  // The display name of the group user.
	Active       bool   `json:"active"`       // Indicates if the group user is active.
	TimeZone     string `json:"timeZone"`     // The time zone of the group user.
	AccountType  string `json:"accountType"`  // The account type of the group user.
}

// GroupBulkOptionsScheme represents the bulk options for a group in Jira.
type GroupBulkOptionsScheme struct {
	GroupIDs   []string // The IDs of the groups.
	GroupNames []string // The names of the groups.
}
