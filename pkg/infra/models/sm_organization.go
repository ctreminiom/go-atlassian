package models

type OrganizationUsersPageScheme struct {
	Size       int                              `json:"size,omitempty"`
	Start      int                              `json:"start,omitempty"`
	Limit      int                              `json:"limit,omitempty"`
	IsLastPage bool                             `json:"isLastPage,omitempty"`
	Values     []*OrganizationUserScheme        `json:"values,omitempty"`
	Expands    []string                         `json:"_expands,omitempty"`
	Links      *OrganizationUsersPageLinkScheme `json:"_links,omitempty"`
}

type OrganizationUsersPageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type OrganizationUserScheme struct {
	AccountID    string                      `json:"accountId,omitempty"`
	Name         string                      `json:"name,omitempty"`
	Key          string                      `json:"key,omitempty"`
	EmailAddress string                      `json:"emailAddress,omitempty"`
	DisplayName  string                      `json:"displayName,omitempty"`
	Active       bool                        `json:"active,omitempty"`
	TimeZone     string                      `json:"timeZone,omitempty"`
	Links        *OrganizationUserLinkScheme `json:"_links,omitempty"`
}

type OrganizationUserLinkScheme struct {
	Self     string `json:"self,omitempty"`
	JiraRest string `json:"jiraRest,omitempty"`
}

type OrganizationPageScheme struct {
	Size       int                         `json:"size,omitempty"`
	Start      int                         `json:"start,omitempty"`
	Limit      int                         `json:"limit,omitempty"`
	IsLastPage bool                        `json:"isLastPage,omitempty"`
	Values     []*OrganizationScheme       `json:"values,omitempty"`
	Expands    []string                    `json:"_expands,omitempty"`
	Links      *OrganizationPageLinkScheme `json:"_links,omitempty"`
}

type OrganizationPageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type OrganizationScheme struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}
