package models

type UserGroupsScheme struct {
	Size       int                `json:"size,omitempty"`
	Items      []*UserGroupScheme `json:"items,omitempty"`
	MaxResults int                `json:"max-results,omitempty"`
}

type UserGroupScheme struct {
	Name string `json:"name,omitempty"`
	Self string `json:"self,omitempty"`
}

type GroupScheme struct {
	Name   string               `json:"name,omitempty"`
	Self   string               `json:"self,omitempty"`
	Users  *GroupUserPageScheme `json:"users,omitempty"`
	Expand string               `json:"expand,omitempty"`
}

type GroupUserPageScheme struct {
	Size       int           `json:"size,omitempty"`
	Items      []*UserScheme `json:"items,omitempty"`
	MaxResults int           `json:"max-results,omitempty"`
	StartIndex int           `json:"start-index,omitempty"`
	EndIndex   int           `json:"end-index,omitempty"`
}

type BulkGroupScheme struct {
	MaxResults int  `json:"maxResults,omitempty"`
	StartAt    int  `json:"startAt,omitempty"`
	Total      int  `json:"total,omitempty"`
	IsLast     bool `json:"isLast,omitempty"`
	Values     []struct {
		Name    string `json:"name,omitempty"`
		GroupID string `json:"groupId,omitempty"`
	} `json:"values,omitempty"`
}

type GroupMemberPageScheme struct {
	Self       string                   `json:"self,omitempty"`
	NextPage   string                   `json:"nextPage,omitempty"`
	MaxResults int                      `json:"maxResults,omitempty"`
	StartAt    int                      `json:"startAt,omitempty"`
	Total      int                      `json:"total,omitempty"`
	IsLast     bool                     `json:"isLast,omitempty"`
	Values     []*GroupUserDetailScheme `json:"values,omitempty"`
}

type GroupUserDetailScheme struct {
	Self         string `json:"self"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	AccountID    string `json:"accountId"`
	EmailAddress string `json:"emailAddress"`
	DisplayName  string `json:"displayName"`
	Active       bool   `json:"active"`
	TimeZone     string `json:"timeZone"`
	AccountType  string `json:"accountType"`
}

type GroupBulkOptionsScheme struct {
	GroupIDs   []string
	GroupNames []string
}
