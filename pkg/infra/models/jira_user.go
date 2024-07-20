package models

// UserScheme represents a user in Jira.
type UserScheme struct {
	Self             string                      `json:"self,omitempty"`             // The URL of the user.
	Key              string                      `json:"key,omitempty"`              // The key of the user.
	AccountID        string                      `json:"accountId,omitempty"`        // The account ID of the user.
	AccountType      string                      `json:"accountType,omitempty"`      // The account type of the user.
	Name             string                      `json:"name,omitempty"`             // The name of the user.
	EmailAddress     string                      `json:"emailAddress,omitempty"`     // The email address of the user.
	AvatarURLs       *AvatarURLScheme            `json:"avatarUrls,omitempty"`       // The avatar URLs of the user.
	DisplayName      string                      `json:"displayName,omitempty"`      // The display name of the user.
	Active           bool                        `json:"active,omitempty"`           // Indicates if the user is active.
	TimeZone         string                      `json:"timeZone,omitempty"`         // The time zone of the user.
	Locale           string                      `json:"locale,omitempty"`           // The locale of the user.
	Groups           *UserGroupsScheme           `json:"groups,omitempty"`           // The groups of the user.
	ApplicationRoles *UserApplicationRolesScheme `json:"applicationRoles,omitempty"` // The application roles of the user.
	Expand           string                      `json:"expand,omitempty"`           // The fields that are expanded in the results.
}

// UserApplicationRolesScheme represents the application roles of a user in Jira.
type UserApplicationRolesScheme struct {
	Size       int                               `json:"size,omitempty"`        // The size of the application roles.
	Items      []*UserApplicationRoleItemsScheme `json:"items,omitempty"`       // The items of the application roles.
	MaxResults int                               `json:"max-results,omitempty"` // The maximum number of results returned.
}

// UserApplicationRoleItemsScheme represents an item of the application roles of a user in Jira.
type UserApplicationRoleItemsScheme struct {
	Key                  string   `json:"key,omitempty"`                  // The key of the item.
	Groups               []string `json:"groups,omitempty"`               // The groups of the item.
	Name                 string   `json:"name,omitempty"`                 // The name of the item.
	DefaultGroups        []string `json:"defaultGroups,omitempty"`        // The default groups of the item.
	SelectedByDefault    bool     `json:"selectedByDefault,omitempty"`    // Indicates if the item is selected by default.
	Defined              bool     `json:"defined,omitempty"`              // Indicates if the item is defined.
	NumberOfSeats        int      `json:"numberOfSeats,omitempty"`        // The number of seats of the item.
	RemainingSeats       int      `json:"remainingSeats,omitempty"`       // The remaining seats of the item.
	UserCount            int      `json:"userCount,omitempty"`            // The user count of the item.
	UserCountDescription string   `json:"userCountDescription,omitempty"` // The user count description of the item.
	HasUnlimitedSeats    bool     `json:"hasUnlimitedSeats,omitempty"`    // Indicates if the item has unlimited seats.
	Platform             bool     `json:"platform,omitempty"`             // Indicates if the item is a platform.
}

// UserPayloadScheme represents the payload for a user in Jira.
type UserPayloadScheme struct {
	Password     string `json:"password,omitempty"`     // The password of the user.
	EmailAddress string `json:"emailAddress,omitempty"` // The email address of the user.
	DisplayName  string `json:"displayName,omitempty"`  // The display name of the user.
	Notification bool   `json:"notification,omitempty"` // Indicates if the user receives notifications.
}

// UserSearchPageScheme represents a page of users in Jira.
type UserSearchPageScheme struct {
	MaxResults int           `json:"maxResults,omitempty"` // The maximum number of results returned.
	StartAt    int           `json:"startAt,omitempty"`    // The index of the first result returned.
	Total      int           `json:"total,omitempty"`      // The total number of results available.
	IsLast     bool          `json:"isLast,omitempty"`     // Indicates if this is the last page of results.
	Values     []*UserScheme `json:"values,omitempty"`     // The users on the page.
}

// UserPermissionCheckParamsScheme represents the parameters for checking a user's permissions in Jira.
type UserPermissionCheckParamsScheme struct {
	Query      string // The query for the check.
	AccountID  string // The account ID for the check.
	IssueKey   string // The issue key for the check.
	ProjectKey string // The project key for the check.
}
