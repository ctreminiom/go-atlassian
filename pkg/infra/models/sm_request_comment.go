package models

// RequestCommentOptionsScheme represents the options for filtering and paginating request comments.
type RequestCommentOptionsScheme struct {
	Public   *bool    `url:"public,omitempty"`
	Internal *bool    `url:"internal,omitempty"`
	Expand   []string `url:"expand,omitempty"`
	Start    int      `url:"start,omitempty"`
	Limit    int      `url:"limit,omitempty"`
}

// RequestCommentPageScheme represents a page of request comments in a system.
type RequestCommentPageScheme struct {
	Size       int                           `json:"size"`       // The number of request comments on the page.
	Start      int                           `json:"start"`      // The index of the first request comment on the page.
	Limit      int                           `json:"limit"`      // The maximum number of request comments that can be on the page.
	IsLastPage bool                          `json:"isLastPage"` // Indicates if this is the last page of request comments.
	Values     []*RequestCommentScheme       `json:"values"`     // The request comments on the page.
	Expands    []string                      `json:"_expands"`   // Additional data related to the request comments.
	Links      *RequestCommentPageLinkScheme `json:"_links"`     // Links related to the page of request comments.
}

// RequestCommentPageLinkScheme represents links related to a page of request comments.
type RequestCommentPageLinkScheme struct {
	Self    string `json:"self"`    // The URL of the page itself.
	Base    string `json:"base"`    // The base URL for the links.
	Context string `json:"context"` // The context for the links.
	Next    string `json:"next"`    // The URL for the next page of request comments.
	Prev    string `json:"prev"`    // The URL for the previous page of request comments.
}

// RequestCommentScheme represents a request comment in a system.
type RequestCommentScheme struct {
	ID           string                       `json:"id,omitempty"`           // The ID of the request comment.
	Body         string                       `json:"body,omitempty"`         // The body of the request comment.
	RenderedBody *RequestCommentRenderScheme  `json:"renderedBody,omitempty"` // The rendered body of the request comment.
	Author       *RequestAuthorScheme         `json:"author,omitempty"`       // The author of the request comment.
	Created      *CustomerRequestDateScheme   `json:"created,omitempty"`      // The created date of the request comment.
	Attachments  *RequestAttachmentPageScheme `json:"attachments,omitempty"`  // The attachments of the request comment.
	Expands      []string                     `json:"_expands,omitempty"`     // The fields to expand in the request comment.
	Public       bool                         `json:"public,omitempty"`       // Indicates if the request comment is public.
	Links        *RequestCommentLinkScheme    `json:"_links,omitempty"`       // Links related to the request comment.
}

// RequestCommentLinkScheme represents links related to a request comment.
type RequestCommentLinkScheme struct {
	Self string `json:"self"` // The URL of the request comment itself.
}

// RequestCommentRenderScheme represents the rendered body of a request comment.
type RequestCommentRenderScheme struct {
	HTML string `json:"html"` // The HTML of the rendered body.
}
