package models

type RequestParticipantPageScheme struct {
	Size       int                               `json:"size,omitempty"`
	Start      int                               `json:"start,omitempty"`
	Limit      int                               `json:"limit,omitempty"`
	IsLastPage bool                              `json:"isLastPage,omitempty"`
	Values     []*RequestParticipantScheme       `json:"values,omitempty"`
	Expands    []string                          `json:"_expands,omitempty"`
	Links      *RequestParticipantPageLinkScheme `json:"_links,omitempty"`
}

type RequestParticipantPageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type RequestParticipantScheme struct {
	AccountID    string                        `json:"accountId,omitempty"`
	Name         string                        `json:"name,omitempty"`
	Key          string                        `json:"key,omitempty"`
	EmailAddress string                        `json:"emailAddress,omitempty"`
	DisplayName  string                        `json:"displayName,omitempty"`
	Active       bool                          `json:"active,omitempty"`
	TimeZone     string                        `json:"timeZone,omitempty"`
	Links        *RequestParticipantLinkScheme `json:"_links,omitempty"`
}

type RequestParticipantLinkScheme struct {
	Self     string `json:"self,omitempty"`
	JiraRest string `json:"jiraRest,omitempty"`
}
