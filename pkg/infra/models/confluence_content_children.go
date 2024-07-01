package models

// ContentChildrenScheme represents the children of a content item in Confluence.
type ContentChildrenScheme struct {
	Attachment *ContentPageScheme `json:"attachment,omitempty"` // The attachment of the content item.
	Comments   *ContentPageScheme `json:"comment,omitempty"`    // The comments of the content item.
	Page       *ContentPageScheme `json:"page,omitempty"`       // The page of the content item.
	Links      *LinkScheme        `json:"_links,omitempty"`     // The links related to the content item.
}

// CopyOptionsScheme represents the options for copying a content item in Confluence.
type CopyOptionsScheme struct {
	CopyAttachments    bool                       `json:"copyAttachments,omitempty"`    // Indicates if the attachments should be copied.
	CopyPermissions    bool                       `json:"copyPermissions,omitempty"`    // Indicates if the permissions should be copied.
	CopyProperties     bool                       `json:"copyProperties,omitempty"`     // Indicates if the properties should be copied.
	CopyLabels         bool                       `json:"copyLabels,omitempty"`         // Indicates if the labels should be copied.
	CopyCustomContents bool                       `json:"copyCustomContents,omitempty"` // Indicates if the custom contents should be copied.
	DestinationPageID  string                     `json:"destinationPageId,omitempty"`  // The ID of the destination page.
	TitleOptions       *CopyTitleOptionScheme     `json:"titleOptions,omitempty"`       // The options for the title of the copied content.
	Destination        *CopyPageDestinationScheme `json:"destination,omitempty"`        // The destination of the copied content.
	PageTitle          string                     `json:"pageTitle,omitempty"`          // The title of the page.
	Body               *CopyPageBodyScheme        `json:"body,omitempty"`               // The body of the copied content.
}

// CopyTitleOptionScheme represents the options for the title of a copied content item in Confluence.
type CopyTitleOptionScheme struct {
	Prefix  string `json:"prefix,omitempty"`  // The prefix of the title.
	Replace string `json:"replace,omitempty"` // The replacement for the title.
	Search  string `json:"search,omitempty"`  // The search term for the title.
}

// CopyPageDestinationScheme represents the destination of a copied content item in Confluence.
type CopyPageDestinationScheme struct {
	Type  string `json:"type,omitempty"`  // The type of the destination.
	Value string `json:"value,omitempty"` // The value of the destination.
}

// CopyPageBodyScheme represents the body of a copied content item in Confluence.
type CopyPageBodyScheme struct {
	Storage *BodyNodeScheme `json:"storage"` // The storage of the body.
	Editor2 *BodyNodeScheme `json:"editor2"` // The editor2 of the body.
}

// ContentTaskScheme represents a task in a content item in Confluence.
type ContentTaskScheme struct {
	ID    string          `json:"id,omitempty"`    // The ID of the task.
	Links *TaskLinkScheme `json:"links,omitempty"` // The links related to the task.
}

// TaskLinkScheme represents the links related to a task in a content item in Confluence.
type TaskLinkScheme struct {
	Status string `json:"status"` // The status of the task.
}
