package models

type RequestSLAPageScheme struct {
	Size       int                       `json:"size,omitempty"`
	Start      int                       `json:"start,omitempty"`
	Limit      int                       `json:"limit,omitempty"`
	IsLastPage bool                      `json:"isLastPage,omitempty"`
	Values     []*RequestSLAScheme       `json:"values,omitempty"`
	Expands    []string                  `json:"_expands,omitempty"`
	Links      *RequestSLAPageLinkScheme `json:"_links,omitempty"`
}

type RequestSLAPageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type RequestSLAScheme struct {
	ID           string                        `json:"id,omitempty"`
	Name         string                        `json:"name,omitempty"`
	OngoingCycle *RequestSLAOngoingCycleScheme `json:"ongoingCycle,omitempty"`
	Links        *RequestSLALinkScheme         `json:"_links,omitempty"`
}

type RequestSLAOngoingCycleScheme struct {
	Breached            bool `json:"breached,omitempty"`
	Paused              bool `json:"paused,omitempty"`
	WithinCalendarHours bool `json:"withinCalendarHours,omitempty"`
}

type RequestSLALinkScheme struct {
	Self string `json:"self,omitempty"`
}
