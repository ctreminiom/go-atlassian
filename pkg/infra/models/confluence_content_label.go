package models

type ContentLabelPayloadScheme struct {
	Prefix string `json:"prefix,omitempty"`
	Name   string `json:"name,omitempty"`
}

type ContentLabelPageScheme struct {
	Results []*ContentLabelScheme `json:"results,omitempty"`
	Start   int                   `json:"start,omitempty"`
	Limit   int                   `json:"limit,omitempty"`
	Size    int                   `json:"size,omitempty"`
}

type ContentLabelScheme struct {
	Prefix string `json:"prefix,omitempty"`
	Name   string `json:"name,omitempty"`
	ID     string `json:"id,omitempty"`
	Label  string `json:"label,omitempty"`
}
