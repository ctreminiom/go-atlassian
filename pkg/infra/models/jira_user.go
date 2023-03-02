package models

type UserScheme struct {
	Self             string                      `json:"self,omitempty"`
	Key              string                      `json:"key,omitempty"`
	AccountID        string                      `json:"accountId,omitempty"`
	AccountType      string                      `json:"accountType,omitempty"`
	Name             string                      `json:"name,omitempty"`
	EmailAddress     string                      `json:"emailAddress,omitempty"`
	AvatarUrls       *AvatarURLScheme            `json:"avatarUrls,omitempty"`
	DisplayName      string                      `json:"displayName,omitempty"`
	Active           bool                        `json:"active,omitempty"`
	TimeZone         string                      `json:"timeZone,omitempty"`
	Locale           string                      `json:"locale,omitempty"`
	Groups           *UserGroupsScheme           `json:"groups,omitempty"`
	ApplicationRoles *UserApplicationRolesScheme `json:"applicationRoles,omitempty"`
	Expand           string                      `json:"expand,omitempty"`
}

type UserApplicationRolesScheme struct {
	Size       int                               `json:"size,omitempty"`
	Items      []*UserApplicationRoleItemsScheme `json:"items,omitempty"`
	MaxResults int                               `json:"max-results,omitempty"`
}

type UserApplicationRoleItemsScheme struct {
	Key                  string   `json:"key,omitempty"`
	Groups               []string `json:"groups,omitempty"`
	Name                 string   `json:"name,omitempty"`
	DefaultGroups        []string `json:"defaultGroups,omitempty"`
	SelectedByDefault    bool     `json:"selectedByDefault,omitempty"`
	Defined              bool     `json:"defined,omitempty"`
	NumberOfSeats        int      `json:"numberOfSeats,omitempty"`
	RemainingSeats       int      `json:"remainingSeats,omitempty"`
	UserCount            int      `json:"userCount,omitempty"`
	UserCountDescription string   `json:"userCountDescription,omitempty"`
	HasUnlimitedSeats    bool     `json:"hasUnlimitedSeats,omitempty"`
	Platform             bool     `json:"platform,omitempty"`
}

type UserPayloadScheme struct {
	Password     string `json:"password,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Notification bool   `json:"notification,omitempty"`
}

type UserSearchPageScheme struct {
	MaxResults int           `json:"maxResults,omitempty"`
	StartAt    int           `json:"startAt,omitempty"`
	Total      int           `json:"total,omitempty"`
	IsLast     bool          `json:"isLast,omitempty"`
	Values     []*UserScheme `json:"values,omitempty"`
}

type UserPermissionCheckParamsScheme struct {
	Query      string
	AccountID  string
	IssueKey   string
	ProjectKey string
}
