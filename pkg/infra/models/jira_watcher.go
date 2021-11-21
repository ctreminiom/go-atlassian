package models

type IssueWatcherScheme struct {
	Self       string              `json:"self,omitempty"`
	IsWatching bool                `json:"isWatching,omitempty"`
	WatchCount int                 `json:"watchCount,omitempty"`
	Watchers   []*UserDetailScheme `json:"watchers,omitempty"`
}

type UserDetailScheme struct {
	Self         string `json:"self,omitempty"`
	Name         string `json:"name,omitempty"`
	Key          string `json:"key,omitempty"`
	AccountID    string `json:"accountId,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Active       bool   `json:"active,omitempty"`
	TimeZone     string `json:"timeZone,omitempty"`
	AccountType  string `json:"accountType,omitempty"`
}
