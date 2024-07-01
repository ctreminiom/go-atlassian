package models

// OrganizationUsersPageScheme represents a page of organization users in a system.
type OrganizationUsersPageScheme struct {
	Size       int                              `json:"size,omitempty"`       // The number of organization users on the page.
	Start      int                              `json:"start,omitempty"`      // The index of the first organization user on the page.
	Limit      int                              `json:"limit,omitempty"`      // The maximum number of organization users that can be on the page.
	IsLastPage bool                             `json:"isLastPage,omitempty"` // Indicates if this is the last page of organization users.
	Values     []*OrganizationUserScheme        `json:"values,omitempty"`     // The organization users on the page.
	Expands    []string                         `json:"_expands,omitempty"`   // Additional data related to the organization users.
	Links      *OrganizationUsersPageLinkScheme `json:"_links,omitempty"`     // Links related to the page of organization users.
}

// OrganizationUsersPageLinkScheme represents links related to a page of organization users.
type OrganizationUsersPageLinkScheme struct {
	Self    string `json:"self,omitempty"`    // The URL of the page itself.
	Base    string `json:"base,omitempty"`    // The base URL for the links.
	Context string `json:"context,omitempty"` // The context for the links.
	Next    string `json:"next,omitempty"`    // The URL for the next page of organization users.
	Prev    string `json:"prev,omitempty"`    // The URL for the previous page of organization users.
}

// OrganizationUserScheme represents an organization user in a system.
type OrganizationUserScheme struct {
	AccountID    string                      `json:"accountId,omitempty"`    // The account ID of the organization user.
	Name         string                      `json:"name,omitempty"`         // The name of the organization user.
	Key          string                      `json:"key,omitempty"`          // The key of the organization user.
	EmailAddress string                      `json:"emailAddress,omitempty"` // The email address of the organization user.
	DisplayName  string                      `json:"displayName,omitempty"`  // The display name of the organization user.
	Active       bool                        `json:"active,omitempty"`       // Indicates if the organization user is active.
	TimeZone     string                      `json:"timeZone,omitempty"`     // The time zone of the organization user.
	Links        *OrganizationUserLinkScheme `json:"_links,omitempty"`       // Links related to the organization user.
}

// OrganizationUserLinkScheme represents links related to an organization user.
type OrganizationUserLinkScheme struct {
	Self     string `json:"self,omitempty"`     // The URL of the organization user itself.
	JiraRest string `json:"jiraRest,omitempty"` // The Jira REST API link for the organization user.
}

// OrganizationPageScheme represents a page of organizations in a system.
type OrganizationPageScheme struct {
	Size       int                         `json:"size,omitempty"`       // The number of organizations on the page.
	Start      int                         `json:"start,omitempty"`      // The index of the first organization on the page.
	Limit      int                         `json:"limit,omitempty"`      // The maximum number of organizations that can be on the page.
	IsLastPage bool                        `json:"isLastPage,omitempty"` // Indicates if this is the last page of organizations.
	Values     []*OrganizationScheme       `json:"values,omitempty"`     // The organizations on the page.
	Expands    []string                    `json:"_expands,omitempty"`   // Additional data related to the organizations.
	Links      *OrganizationPageLinkScheme `json:"_links,omitempty"`     // Links related to the page of organizations.
}

// OrganizationPageLinkScheme represents links related to a page of organizations.
type OrganizationPageLinkScheme struct {
	Self    string `json:"self,omitempty"`    // The URL of the page itself.
	Base    string `json:"base,omitempty"`    // The base URL for the links.
	Context string `json:"context,omitempty"` // The context for the links.
	Next    string `json:"next,omitempty"`    // The URL for the next page of organizations.
	Prev    string `json:"prev,omitempty"`    // The URL for the previous page of organizations.
}

// OrganizationScheme represents an organization in a system.
type OrganizationScheme struct {
	ID    string `json:"id,omitempty"`   // The ID of the organization.
	Name  string `json:"name,omitempty"` // The name of the organization.
	Links struct {
		Self string `json:"self,omitempty"` // The URL of the organization itself.
	} `json:"_links,omitempty"` // Links related to the organization.
}
