package models

// ContentPropertyPayloadScheme represents the payload for a content property in Confluence.
type ContentPropertyPayloadScheme struct {
	Key   string `json:"key"`   // The key of the content property.
	Value string `json:"value"` // The value of the content property.
}

// ContentPropertyPageScheme represents a page of content properties in Confluence.
type ContentPropertyPageScheme struct {
	Results []*ContentPropertyScheme `json:"results,omitempty"` // The content properties in the page.
	Start   int                      `json:"start,omitempty"`   // The start index of the content properties in the page.
	Limit   int                      `json:"limit,omitempty"`   // The limit of the content properties in the page.
	Size    int                      `json:"size,omitempty"`    // The size of the content properties in the page.
}

// ContentPropertyScheme represents a content property in Confluence.
type ContentPropertyScheme struct {
	ID         string                        `json:"id,omitempty"`      // The ID of the content property.
	Key        string                        `json:"key,omitempty"`     // The key of the content property.
	Value      interface{}                   `json:"value,omitempty"`   // The value of the content property.
	Version    *ContentPropertyVersionScheme `json:"version,omitempty"` // The version of the content property.
	Expandable struct {
		Content              string `json:"content,omitempty"`              // The content of the content property.
		AdditionalProperties string `json:"additionalProperties,omitempty"` // The additional properties of the content property.
	} `json:"_expandable,omitempty"` // The expandable fields of the content property.
}

// ContentPropertyVersionScheme represents the version of a content property in Confluence.
type ContentPropertyVersionScheme struct {
	When                string `json:"when,omitempty"`                // The timestamp of the version.
	Message             string `json:"message,omitempty"`             // The message of the version.
	Number              int    `json:"number,omitempty"`              // The number of the version.
	MinorEdit           bool   `json:"minorEdit,omitempty"`           // Indicates if the version is a minor edit.
	ContentTypeModified bool   `json:"contentTypeModified,omitempty"` // Indicates if the content type is modified in the version.
}
