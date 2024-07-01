package models

// ContentLabelPayloadScheme represents the payload for a content label in Confluence.
type ContentLabelPayloadScheme struct {
	Prefix string `json:"prefix,omitempty"` // The prefix of the content label.
	Name   string `json:"name,omitempty"`   // The name of the content label.
}

// ContentLabelPageScheme represents a page of content labels in Confluence.
type ContentLabelPageScheme struct {
	Results []*ContentLabelScheme `json:"results,omitempty"` // The content labels in the page.
	Start   int                   `json:"start,omitempty"`   // The start index of the content labels in the page.
	Limit   int                   `json:"limit,omitempty"`   // The limit of the content labels in the page.
	Size    int                   `json:"size,omitempty"`    // The size of the content labels in the page.
}

// ContentLabelScheme represents a content label in Confluence.
type ContentLabelScheme struct {
	Prefix string `json:"prefix,omitempty"` // The prefix of the content label.
	Name   string `json:"name,omitempty"`   // The name of the content label.
	ID     string `json:"id,omitempty"`     // The ID of the content label.
	Label  string `json:"label,omitempty"`  // The label of the content label.
}
