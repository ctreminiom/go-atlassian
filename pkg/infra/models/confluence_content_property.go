package models

type ContentPropertyPayloadScheme struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ContentPropertyPageScheme struct {
	Results []*ContentPropertyScheme `json:"results,omitempty"`
	Start   int                      `json:"start,omitempty"`
	Limit   int                      `json:"limit,omitempty"`
	Size    int                      `json:"size,omitempty"`
}

type ContentPropertyScheme struct {
	ID         string                        `json:"id,omitempty"`
	Key        string                        `json:"key,omitempty"`
	Value      interface{}                   `json:"value,omitempty"`
	Version    *ContentPropertyVersionScheme `json:"version,omitempty"`
	Expandable struct {
		Content              string `json:"content,omitempty"`
		AdditionalProperties string `json:"additionalProperties,omitempty"`
	} `json:"_expandable,omitempty"`
}

type ContentPropertyVersionScheme struct {
	When                string `json:"when,omitempty"`
	Message             string `json:"message,omitempty"`
	Number              int    `json:"number,omitempty"`
	MinorEdit           bool   `json:"minorEdit,omitempty"`
	ContentTypeModified bool   `json:"contentTypeModified,omitempty"`
}
