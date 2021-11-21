package models

type CustomerPageScheme struct {
	Expands    []interface{}            `json:"_expands,omitempty"`
	Size       int                      `json:"size,omitempty"`
	Start      int                      `json:"start,omitempty"`
	Limit      int                      `json:"limit,omitempty"`
	IsLastPage bool                     `json:"isLastPage,omitempty"`
	Links      *CustomerPageLinksScheme `json:"_links,omitempty"`
	Values     []*CustomerScheme        `json:"values,omitempty"`
}

type CustomerPageLinksScheme struct {
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type CustomerScheme struct {
	AccountID    string              `json:"accountId,omitempty"`
	Name         string              `json:"name,omitempty"`
	Key          string              `json:"key,omitempty"`
	EmailAddress string              `json:"emailAddress,omitempty"`
	DisplayName  string              `json:"displayName,omitempty"`
	Active       bool                `json:"active,omitempty"`
	TimeZone     string              `json:"timeZone,omitempty"`
	Links        *CustomerLinkScheme `json:"_links,omitempty"`
}

type CustomerLinkScheme struct {
	JiraRest   string           `json:"jiraRest"`
	AvatarUrls *AvatarURLScheme `json:"avatarUrls"`
	Self       string           `json:"self"`
}
