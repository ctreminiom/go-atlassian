package models

// IssueWatcherScheme represents the watcher information for an issue in Jira.
type IssueWatcherScheme struct {
	Self       string              `json:"self,omitempty"`       // The URL of the watcher information.
	IsWatching bool                `json:"isWatching,omitempty"` // Indicates if the current user is watching the issue.
	WatchCount int                 `json:"watchCount,omitempty"` // The number of watchers for the issue.
	Watchers   []*UserDetailScheme `json:"watchers,omitempty"`   // The users who are watching the issue.
}

// UserDetailScheme represents the detail of a user in Jira.
type UserDetailScheme struct {
	Self         string `json:"self,omitempty"`         // The URL of the user detail.
	Name         string `json:"name,omitempty"`         // The name of the user.
	Key          string `json:"key,omitempty"`          // The key of the user.
	AccountID    string `json:"accountId,omitempty"`    // The account ID of the user.
	EmailAddress string `json:"emailAddress,omitempty"` // The email address of the user.
	DisplayName  string `json:"displayName,omitempty"`  // The display name of the user.
	Active       bool   `json:"active,omitempty"`       // Indicates if the user is active.
	TimeZone     string `json:"timeZone,omitempty"`     // The time zone of the user.
	AccountType  string `json:"accountType,omitempty"`  // The account type of the user.
}
