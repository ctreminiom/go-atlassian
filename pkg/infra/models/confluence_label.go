package models

// LabelDetailsScheme represents the details of a label in Confluence.
type LabelDetailsScheme struct {
	Label              *ContentLabelScheme               `json:"label"`              // The label.
	AssociatedContents *LabelAssociatedContentPageScheme `json:"associatedContents"` // The associated contents of the label.
}

// LabelAssociatedContentPageScheme represents a page of associated contents of a label in Confluence.
type LabelAssociatedContentPageScheme struct {
	Results []*LabelAssociatedContentScheme `json:"results,omitempty"` // The associated contents in the page.
	Start   int                             `json:"start,omitempty"`   // The start index of the associated contents in the page.
	Limit   int                             `json:"limit,omitempty"`   // The limit of the associated contents in the page.
	Size    int                             `json:"size,omitempty"`    // The size of the associated contents in the page.
}

// LabelAssociatedContentScheme represents an associated content of a label in Confluence.
type LabelAssociatedContentScheme struct {
	ContentType string `json:"contentType,omitempty"` // The content type of the associated content.
	ContentID   int    `json:"contentId,omitempty"`   // The ID of the associated content.
	Title       string `json:"title,omitempty"`       // The title of the associated content.
}
