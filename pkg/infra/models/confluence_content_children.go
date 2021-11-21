package models

type ContentChildrenScheme struct {
	Attachment *ContentPageScheme `json:"attachment,omitempty"`
	Comments   *ContentPageScheme `json:"comment,omitempty"`
	Page       *ContentPageScheme `json:"page,omitempty"`
	Links      *LinkScheme        `json:"_links,omitempty"`
}

type CopyOptionsScheme struct {
	CopyAttachments    bool                       `json:"copyAttachments,omitempty"`
	CopyPermissions    bool                       `json:"copyPermissions,omitempty"`
	CopyProperties     bool                       `json:"copyProperties,omitempty"`
	CopyLabels         bool                       `json:"copyLabels,omitempty"`
	CopyCustomContents bool                       `json:"copyCustomContents,omitempty"`
	DestinationPageID  string                     `json:"destinationPageId,omitempty"`
	TitleOptions       *CopyTitleOptionScheme     `json:"titleOptions,omitempty"`
	Destination        *CopyPageDestinationScheme `json:"destination,omitempty"`
	PageTitle          string                     `json:"pageTitle,omitempty"`
	Body               *CopyPageBodyScheme        `json:"body,omitempty"`
}

type CopyTitleOptionScheme struct {
	Prefix  string `json:"prefix,omitempty"`
	Replace string `json:"replace,omitempty"`
	Search  string `json:"search,omitempty"`
}

type CopyPageDestinationScheme struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type CopyPageBodyScheme struct {
	Storage *BodyNodeScheme `json:"storage"`
	Editor2 *BodyNodeScheme `json:"editor2"`
}

type ContentTaskScheme struct {
	ID    string          `json:"id,omitempty"`
	Links *TaskLinkScheme `json:"links,omitempty"`
}

type TaskLinkScheme struct {
	Status string `json:"status"`
}
