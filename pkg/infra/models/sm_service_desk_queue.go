package models

// ServiceDeskQueuePageScheme represents a page of service desk queues.
// It contains information about the page and a slice of service desk queues.
type ServiceDeskQueuePageScheme struct {
	Size       int                             `json:"size,omitempty"`       // The size of the page.
	Start      int                             `json:"start,omitempty"`      // The start index of the page.
	Limit      int                             `json:"limit,omitempty"`      // The limit of the page.
	IsLastPage bool                            `json:"isLastPage,omitempty"` // Indicates if this is the last page.
	Values     []*ServiceDeskQueueScheme       `json:"values,omitempty"`     // The service desk queues in the page.
	Expands    []string                        `json:"_expands,omitempty"`   // The fields to expand in the page.
	Links      *ServiceDeskQueuePageLinkScheme `json:"_links,omitempty"`     // The links related to the page.
}

// ServiceDeskQueuePageLinkScheme represents the links related to a page of service desk queues.
type ServiceDeskQueuePageLinkScheme struct {
	Self    string `json:"self,omitempty"`    // The self link of the page.
	Base    string `json:"base,omitempty"`    // The base link of the page.
	Context string `json:"context,omitempty"` // The context link of the page.
	Next    string `json:"next,omitempty"`    // The next link of the page.
	Prev    string `json:"prev,omitempty"`    // The previous link of the page.
}

// ServiceDeskQueueScheme represents a service desk queue.
// It contains information about the queue and a slice of fields.
type ServiceDeskQueueScheme struct {
	ID         string   `json:"id,omitempty"`         // The ID of the queue.
	Name       string   `json:"name,omitempty"`       // The name of the queue.
	JQL        string   `json:"jql,omitempty"`        // The JQL of the queue.
	Fields     []string `json:"fields,omitempty"`     // The fields of the queue.
	IssueCount int      `json:"issueCount,omitempty"` // The issue count of the queue.
}

// ServiceDeskIssueQueueScheme represents a service desk issue queue.
// It contains information about the page and a slice of issues.
type ServiceDeskIssueQueueScheme struct {
	Size       int                             `json:"size,omitempty"`       // The size of the page.
	Start      int                             `json:"start,omitempty"`      // The start index of the page.
	Limit      int                             `json:"limit,omitempty"`      // The limit of the page.
	IsLastPage bool                            `json:"isLastPage,omitempty"` // Indicates if this is the last page.
	Values     []*IssueSchemeV2                `json:"values,omitempty"`     // The issues in the page.
	Expands    []string                        `json:"_expands,omitempty"`   // The fields to expand in the page.
	Links      *ServiceDeskQueuePageLinkScheme `json:"_links,omitempty"`     // The links related to the page.
}
