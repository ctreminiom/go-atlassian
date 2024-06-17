package models

// RequestSLAPageScheme represents a page of request SLAs.
type RequestSLAPageScheme struct {
	Size       int                       `json:"size,omitempty"`       // The size of the page.
	Start      int                       `json:"start,omitempty"`      // The start index of the page.
	Limit      int                       `json:"limit,omitempty"`      // The limit of the page.
	IsLastPage bool                      `json:"isLastPage,omitempty"` // Indicates if this is the last page.
	Values     []*RequestSLAScheme       `json:"values,omitempty"`     // The request SLAs in the page.
	Expands    []string                  `json:"_expands,omitempty"`   // The fields to expand in the page.
	Links      *RequestSLAPageLinkScheme `json:"_links,omitempty"`     // The links related to the page.
}

// RequestSLAPageLinkScheme represents the links related to a page of request SLAs.
type RequestSLAPageLinkScheme struct {
	Self    string `json:"self,omitempty"`    // The self link of the page.
	Base    string `json:"base,omitempty"`    // The base link of the page.
	Context string `json:"context,omitempty"` // The context link of the page.
	Next    string `json:"next,omitempty"`    // The next link of the page.
	Prev    string `json:"prev,omitempty"`    // The previous link of the page.
}

// RequestSLAScheme represents a request SLA.
type RequestSLAScheme struct {
	ID           string                        `json:"id,omitempty"`           // The ID of the SLA.
	Name         string                        `json:"name,omitempty"`         // The name of the SLA.
	OngoingCycle *RequestSLAOngoingCycleScheme `json:"ongoingCycle,omitempty"` // The ongoing cycle of the SLA.
	Links        *RequestSLALinkScheme         `json:"_links,omitempty"`       // The links related to the SLA.
}

// RequestSLAOngoingCycleScheme represents the ongoing cycle of a request SLA.
type RequestSLAOngoingCycleScheme struct {
	Breached            bool `json:"breached,omitempty"`            // Indicates if the SLA is breached.
	Paused              bool `json:"paused,omitempty"`              // Indicates if the SLA is paused.
	WithinCalendarHours bool `json:"withinCalendarHours,omitempty"` // Indicates if the SLA is within calendar hours.
}

// RequestSLALinkScheme represents the links related to a request SLA.
type RequestSLALinkScheme struct {
	Self string `json:"self,omitempty"` // The self link of the SLA.
}
