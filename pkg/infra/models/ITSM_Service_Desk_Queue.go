package models

type ServiceDeskQueuePageScheme struct {
	Size       int                             `json:"size,omitempty"`
	Start      int                             `json:"start,omitempty"`
	Limit      int                             `json:"limit,omitempty"`
	IsLastPage bool                            `json:"isLastPage,omitempty"`
	Values     []*ServiceDeskQueueScheme       `json:"values,omitempty"`
	Expands    []string                        `json:"_expands,omitempty"`
	Links      *ServiceDeskQueuePageLinkScheme `json:"_links,omitempty"`
}

type ServiceDeskQueuePageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type ServiceDeskQueueScheme struct {
	ID         string   `json:"id,omitempty"`
	Name       string   `json:"name,omitempty"`
	Jql        string   `json:"jql,omitempty"`
	Fields     []string `json:"fields,omitempty"`
	IssueCount int      `json:"issueCount,omitempty"`
}

type ServiceDeskIssueQueueScheme struct {
	Size       int                             `json:"size,omitempty"`
	Start      int                             `json:"start,omitempty"`
	Limit      int                             `json:"limit,omitempty"`
	IsLastPage bool                            `json:"isLastPage,omitempty"`
	Values     []*IssueSchemeV2                `json:"values,omitempty"`
	Expands    []string                        `json:"_expands,omitempty"`
	Links      *ServiceDeskQueuePageLinkScheme `json:"_links,omitempty"`
}
