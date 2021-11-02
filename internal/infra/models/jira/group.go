package jira

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
