package models

type RequestCommentPageScheme struct {
	Size       int                           `json:"size"`
	Start      int                           `json:"start"`
	Limit      int                           `json:"limit"`
	IsLastPage bool                          `json:"isLastPage"`
	Values     []*RequestCommentScheme       `json:"values"`
	Expands    []string                      `json:"_expands"`
	Links      *RequestCommentPageLinkScheme `json:"_links"`
}

type RequestCommentPageLinkScheme struct {
	Self    string `json:"self"`
	Base    string `json:"base"`
	Context string `json:"context"`
	Next    string `json:"next"`
	Prev    string `json:"prev"`
}

type RequestCommentScheme struct {
	ID           string                       `json:"id,omitempty"`
	Body         string                       `json:"body,omitempty"`
	RenderedBody *RequestCommentRenderScheme  `json:"renderedBody,omitempty"`
	Author       *RequestAuthorScheme         `json:"author,omitempty"`
	Created      *CustomerRequestDateScheme   `json:"created,omitempty"`
	Attachments  *RequestAttachmentPageScheme `json:"attachments,omitempty"`
	Expands      []string                     `json:"_expands,omitempty"`
	Public       bool                         `json:"public,omitempty"`
	Links        *RequestCommentLinkScheme    `json:"_links,omitempty"`
}

type RequestCommentLinkScheme struct {
	Self string `json:"self"`
}

type RequestCommentRenderScheme struct {
	HTML string `json:"html"`
}
