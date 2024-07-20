package models

// RequestParticipantPageScheme represents a page of request participants.
type RequestParticipantPageScheme struct {
	Size       int                               `json:"size,omitempty"`       // The size of the page.
	Start      int                               `json:"start,omitempty"`      // The start index of the page.
	Limit      int                               `json:"limit,omitempty"`      // The limit of the page.
	IsLastPage bool                              `json:"isLastPage,omitempty"` // Indicates if this is the last page.
	Values     []*RequestParticipantScheme       `json:"values,omitempty"`     // The request participants in the page.
	Expands    []string                          `json:"_expands,omitempty"`   // The fields to expand in the page.
	Links      *RequestParticipantPageLinkScheme `json:"_links,omitempty"`     // The links related to the page.
}

// RequestParticipantPageLinkScheme represents the links related to a page of request participants.
type RequestParticipantPageLinkScheme struct {
	Self    string `json:"self,omitempty"`    // The self link of the page.
	Base    string `json:"base,omitempty"`    // The base link of the page.
	Context string `json:"context,omitempty"` // The context link of the page.
	Next    string `json:"next,omitempty"`    // The next link of the page.
	Prev    string `json:"prev,omitempty"`    // The previous link of the page.
}

// RequestParticipantScheme represents a request participant.
type RequestParticipantScheme struct {
	AccountID    string                        `json:"accountId,omitempty"`    // The account ID of the participant.
	Name         string                        `json:"name,omitempty"`         // The name of the participant.
	Key          string                        `json:"key,omitempty"`          // The key of the participant.
	EmailAddress string                        `json:"emailAddress,omitempty"` // The email address of the participant.
	DisplayName  string                        `json:"displayName,omitempty"`  // The display name of the participant.
	Active       bool                          `json:"active,omitempty"`       // Indicates if the participant is active.
	TimeZone     string                        `json:"timeZone,omitempty"`     // The time zone of the participant.
	Links        *RequestParticipantLinkScheme `json:"_links,omitempty"`       // The links related to the participant.
}

// RequestParticipantLinkScheme represents the links related to a request participant.
type RequestParticipantLinkScheme struct {
	Self     string `json:"self,omitempty"`     // The self link of the participant.
	JiraREST string `json:"jiraRest,omitempty"` // The Jira REST link of the participant.
}
