package models

type LabelDetailsScheme struct {
	Label              *ContentLabelScheme               `json:"label"`
	AssociatedContents *LabelAssociatedContentPageScheme `json:"associatedContents"`
}

type LabelAssociatedContentPageScheme struct {
	Results []*LabelAssociatedContentScheme `json:"results,omitempty"`
	Start   int                             `json:"start,omitempty"`
	Limit   int                             `json:"limit,omitempty"`
	Size    int                             `json:"size,omitempty"`
}

type LabelAssociatedContentScheme struct {
	ContentType string `json:"contentType,omitempty"`
	ContentID   int    `json:"contentId,omitempty"`
	Title       string `json:"title,omitempty"`
}
