package models

type RequestAttachmentPageScheme struct {
	Size       int                              `json:"size,omitempty"`
	Start      int                              `json:"start,omitempty"`
	Limit      int                              `json:"limit,omitempty"`
	IsLastPage bool                             `json:"isLastPage,omitempty"`
	Values     []*RequestAttachmentScheme       `json:"values,omitempty"`
	Expands    []string                         `json:"_expands,omitempty"`
	Links      *RequestAttachmentPageLinkScheme `json:"_links,omitempty"`
}

type RequestAttachmentPageLinkScheme struct {
	Self    string `json:"self,omitempty"`
	Base    string `json:"base,omitempty"`
	Context string `json:"context,omitempty"`
	Next    string `json:"next,omitempty"`
	Prev    string `json:"prev,omitempty"`
}

type RequestAttachmentScheme struct {
	Filename string                       `json:"filename,omitempty"`
	Author   *RequestAuthorScheme         `json:"author,omitempty"`
	Created  *CustomerRequestDateScheme   `json:"created,omitempty"`
	Size     int                          `json:"size,omitempty"`
	MimeType string                       `json:"mimeType,omitempty"`
	Links    *RequestAttachmentLinkScheme `json:"_links,omitempty"`
}

type RequestAttachmentLinkScheme struct {
	Self      string `json:"self,omitempty"`
	JiraRest  string `json:"jiraRest,omitempty"`
	Content   string `json:"content,omitempty"`
	Thumbnail string `json:"thumbnail,omitempty"`
}

type RequestAuthorScheme struct {
	AccountID    string `json:"accountId,omitempty"`
	Name         string `json:"name,omitempty"`
	Key          string `json:"key,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	DisplayName  string `json:"displayName,omitempty"`
	Active       bool   `json:"active,omitempty"`
	TimeZone     string `json:"timeZone,omitempty"`
}

type RequestAttachmentCreationCommentScheme struct {
	Expands []string                   `json:"_expands,omitempty"`
	ID      string                     `json:"id,omitempty"`
	Body    string                     `json:"body,omitempty"`
	Public  bool                       `json:"public,omitempty"`
	Author  RequestAuthorScheme        `json:"author,omitempty"`
	Created *CustomerRequestDateScheme `json:"created,omitempty"`
	Links   struct {
		Self string `json:"self,omitempty"`
	} `json:"_links,omitempty"`
}

type RequestAttachmentCreationScheme struct {
	Comment     *RequestAttachmentCreationCommentScheme `json:"comment,omitempty"`
	Attachments *RequestAttachmentPageScheme            `json:"attachments,omitempty"`
}
