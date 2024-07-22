package models

// CustomerPageScheme represents a page of customers in a system.
type CustomerPageScheme struct {
	Expands    []interface{}            `json:"_expands,omitempty"`   // Additional data related to the customers.
	Size       int                      `json:"size,omitempty"`       // The number of customers on the page.
	Start      int                      `json:"start,omitempty"`      // The index of the first customer on the page.
	Limit      int                      `json:"limit,omitempty"`      // The maximum number of customers that can be on the page.
	IsLastPage bool                     `json:"isLastPage,omitempty"` // Indicates if this is the last page of customers.
	Links      *CustomerPageLinksScheme `json:"_links,omitempty"`     // Links related to the page of customers.
	Values     []*CustomerScheme        `json:"values,omitempty"`     // The customers on the page.
}

// CustomerPageLinksScheme represents links related to a page of customers.
type CustomerPageLinksScheme struct {
	Base    string `json:"base,omitempty"`    // The base URL for the links.
	Context string `json:"context,omitempty"` // The context for the links.
	Next    string `json:"next,omitempty"`    // The URL for the next page of customers.
	Prev    string `json:"prev,omitempty"`    // The URL for the previous page of customers.
}

// CustomerScheme represents a customer in a system.
type CustomerScheme struct {
	AccountID    string              `json:"accountId,omitempty"`    // The account ID of the customer.
	Name         string              `json:"name,omitempty"`         // The name of the customer.
	Key          string              `json:"key,omitempty"`          // The key of the customer.
	EmailAddress string              `json:"emailAddress,omitempty"` // The email address of the customer.
	DisplayName  string              `json:"displayName,omitempty"`  // The display name of the customer.
	Active       bool                `json:"active,omitempty"`       // Indicates if the customer is active.
	TimeZone     string              `json:"timeZone,omitempty"`     // The time zone of the customer.
	Links        *CustomerLinkScheme `json:"_links,omitempty"`       // Links related to the customer.
}

// CustomerLinkScheme represents links related to a customer.
type CustomerLinkScheme struct {
	JiraREST   string           `json:"jiraRest"`   // The Jira REST API link for the customer.
	AvatarURLs *AvatarURLScheme `json:"avatarUrls"` // The URLs for the customer's avatars.
	Self       string           `json:"self"`       // The URL for the customer itself.
}
