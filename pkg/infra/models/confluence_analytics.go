package models

// ContentViewScheme represents a content view in Confluence.
// ID is the unique identifier of the content view.
// Count is the number of times the content has been viewed.
type ContentViewScheme struct {
	ID    int `json:"id,omitempty"`    // The unique identifier of the content view.
	Count int `json:"count,omitempty"` // The number of times the content has been viewed.
}
