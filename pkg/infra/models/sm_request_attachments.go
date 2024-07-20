package models

// RequestAttachmentPageScheme represents a page of request attachments in a system.
type RequestAttachmentPageScheme struct {
	Size       int                              `json:"size,omitempty"`       // The number of request attachments on the page.
	Start      int                              `json:"start,omitempty"`      // The index of the first request attachment on the page.
	Limit      int                              `json:"limit,omitempty"`      // The maximum number of request attachments that can be on the page.
	IsLastPage bool                             `json:"isLastPage,omitempty"` // Indicates if this is the last page of request attachments.
	Values     []*RequestAttachmentScheme       `json:"values,omitempty"`     // The request attachments on the page.
	Expands    []string                         `json:"_expands,omitempty"`   // Additional data related to the request attachments.
	Links      *RequestAttachmentPageLinkScheme `json:"_links,omitempty"`     // Links related to the page of request attachments.
}

// RequestAttachmentPageLinkScheme represents links related to a page of request attachments.
type RequestAttachmentPageLinkScheme struct {
	Self    string `json:"self,omitempty"`    // The URL of the page itself.
	Base    string `json:"base,omitempty"`    // The base URL for the links.
	Context string `json:"context,omitempty"` // The context for the links.
	Next    string `json:"next,omitempty"`    // The URL for the next page of request attachments.
	Prev    string `json:"prev,omitempty"`    // The URL for the previous page of request attachments.
}

// RequestAttachmentScheme represents a request attachment in a system.
type RequestAttachmentScheme struct {
	Filename string                       `json:"filename,omitempty"` // The filename of the request attachment.
	Author   *RequestAuthorScheme         `json:"author,omitempty"`   // The author of the request attachment.
	Created  *CustomerRequestDateScheme   `json:"created,omitempty"`  // The created date of the request attachment.
	Size     int                          `json:"size,omitempty"`     // The size of the request attachment.
	MimeType string                       `json:"mimeType,omitempty"` // The MIME type of the request attachment.
	Links    *RequestAttachmentLinkScheme `json:"_links,omitempty"`   // Links related to the request attachment.
}

// RequestAttachmentLinkScheme represents links related to a request attachment.
type RequestAttachmentLinkScheme struct {
	Self      string `json:"self,omitempty"`      // The URL of the request attachment itself.
	JiraREST  string `json:"jiraRest,omitempty"`  // The Jira REST API link for the request attachment.
	Content   string `json:"content,omitempty"`   // The content link for the request attachment.
	Thumbnail string `json:"thumbnail,omitempty"` // The thumbnail link for the request attachment.
}

// RequestAuthorScheme represents an author in a system.
type RequestAuthorScheme struct {
	AccountID    string `json:"accountId,omitempty"`    // The account ID of the author.
	Name         string `json:"name,omitempty"`         // The name of the author.
	Key          string `json:"key,omitempty"`          // The key of the author.
	EmailAddress string `json:"emailAddress,omitempty"` // The email address of the author.
	DisplayName  string `json:"displayName,omitempty"`  // The display name of the author.
	Active       bool   `json:"active,omitempty"`       // Indicates if the author is active.
	TimeZone     string `json:"timeZone,omitempty"`     // The time zone of the author.
}

// RequestAttachmentCreationCommentScheme represents a comment during the creation of a request attachment.
type RequestAttachmentCreationCommentScheme struct {
	Expands []string                   `json:"_expands,omitempty"` // The fields to expand in the comment.
	ID      string                     `json:"id,omitempty"`       // The ID of the comment.
	Body    string                     `json:"body,omitempty"`     // The body of the comment.
	Public  bool                       `json:"public,omitempty"`   // Indicates if the comment is public.
	Author  RequestAuthorScheme        `json:"author,omitempty"`   // The author of the comment.
	Created *CustomerRequestDateScheme `json:"created,omitempty"`  // The created date of the comment.
	Links   struct {
		Self string `json:"self,omitempty"` // The URL of the comment itself.
	} `json:"_links,omitempty"` // Links related to the comment.
}

// RequestAttachmentCreationPayloadScheme represents the payload for creating a request attachment.
type RequestAttachmentCreationPayloadScheme struct {
	TemporaryAttachmentIDs []string                                                 `json:"temporaryAttachmentIds,omitempty"`
	Public                 bool                                                     `json:"public"`
	AdditionalComment      *RequestAttachmentCreationAdditionalCommentPayloadScheme `json:"additionalComment,omitempty"`
}

// RequestAttachmentCreationAdditionalCommentPayloadScheme represents the additional comment for creating a request attachment.
type RequestAttachmentCreationAdditionalCommentPayloadScheme struct {
	Body string `json:"body,omitempty"`
}

// RequestAttachmentCreationScheme represents the creation of a request attachment.
type RequestAttachmentCreationScheme struct {
	Comment     *RequestAttachmentCreationCommentScheme `json:"comment,omitempty"`     // The comment during the creation of the request attachment.
	Attachments *RequestAttachmentPageScheme            `json:"attachments,omitempty"` // The request attachment that was created.
}
