package models

// AttachmentScheme represents an attachment in Confluence.
// ID is the unique identifier of the attachment.
// BlogPostID is the ID of the blog post to which the attachment belongs.
// CustomContentID is the custom content ID of the attachment.
// Comment is the comment for the attachment.
// MediaTypeDescription is the description of the media type of the attachment.
// WebuiLink is the web UI link of the attachment.
// DownloadLink is the download link of the attachment.
// Title is the title of the attachment.
// Status is the status of the attachment.
// FileSize is the size of the file of the attachment.
// MediaType is the media type of the attachment.
// PageID is the ID of the page to which the attachment belongs.
// FileID is the ID of the file of the attachment.
// Version is the version of the attachment.
type AttachmentScheme struct {
	ID                   string                   `json:"id,omitempty"`
	BlogPostID           string                   `json:"blogPostId,omitempty"`
	CustomContentID      string                   `json:"customContentId,omitempty"`
	Comment              string                   `json:"comment,omitempty"`
	MediaTypeDescription string                   `json:"mediaTypeDescription,omitempty"`
	WebuiLink            string                   `json:"webuiLink,omitempty"`
	DownloadLink         string                   `json:"downloadLink,omitempty"`
	Title                string                   `json:"title,omitempty"`
	Status               string                   `json:"status,omitempty"`
	FileSize             int                      `json:"fileSize,omitempty"`
	MediaType            string                   `json:"mediaType,omitempty"`
	PageID               string                   `json:"pageId,omitempty"`
	FileID               string                   `json:"fileId,omitempty"`
	Version              *AttachmentVersionScheme `json:"version,omitempty"`
}

// AttachmentVersionScheme represents a version of an attachment in Confluence.
// CreatedAt is the creation time of the version.
// Message is the message for the version.
// Number is the number of the version.
// MinorEdit indicates if the version is a minor edit.
// AuthorID is the ID of the author of the version.
// Attachment is the attachment of the version.
type AttachmentVersionScheme struct {
	CreatedAt  string                       `json:"createdAt,omitempty"`
	Message    string                       `json:"message,omitempty"`
	Number     int                          `json:"number,omitempty"`
	MinorEdit  bool                         `json:"minorEdit,omitempty"`
	AuthorID   string                       `json:"authorId,omitempty"`
	Attachment *AttachmentVersionBodyScheme `json:"attachment,omitempty"`
}

// AttachmentVersionBodyScheme represents the body of a version of an attachment in Confluence.
// Title is the title of the body.
// ID is the ID of the body.
type AttachmentVersionBodyScheme struct {
	Title string `json:"title,omitempty"`
	ID    string `json:"id,omitempty"`
}

// AttachmentParamsScheme represents the parameters for an attachment in Confluence.
// Sort is used to sort the result by a particular field.
// MediaType filters the mediaType of attachments.
// FileName filters on the file-name of attachments.
// SerializeIDs indicates if IDs should be serialized.
type AttachmentParamsScheme struct {
	Sort         string
	MediaType    string
	FileName     string
	SerializeIDs bool
}

// AttachmentPageScheme represents a paginated list of attachments in Confluence.
// Results is a slice of the attachments in the current page.
// Links is a collection of links related to the page.
type AttachmentPageScheme struct {
	Results []*AttachmentScheme `json:"results,omitempty"`
	Links   *PageLinkScheme     `json:"_links,omitempty"`
}

// PageLinkScheme represents a link in a page in Confluence.
// Next is the URL to the next page.
type PageLinkScheme struct {
	Next string `json:"next,omitempty"`
}

// AttachmentVersionPageScheme represents a paginated list of versions of an attachment in Confluence.
// Results is a slice of the versions in the current page.
// Links is a collection of links related to the page.
type AttachmentVersionPageScheme struct {
	Results []*AttachmentVersionScheme `json:"results,omitempty"`
	Links   *PageLinkScheme            `json:"_links,omitempty"`
}

// DetailedVersionScheme represents a detailed version in Confluence.
// Number is the number of the version.
// AuthorID is the ID of the author of the version.
// Message is the message for the version.
// CreatedAt is the creation time of the version.
// MinorEdit indicates if the version is a minor edit.
// ContentTypeModified indicates if the content type was modified in the version.
// Collaborators is a slice of the collaborators of the version.
// PrevVersion is the number of the previous version.
// NextVersion is the number of the next version.
type DetailedVersionScheme struct {
	Number              int      `json:"number,omitempty"`
	AuthorID            string   `json:"authorId,omitempty"`
	Message             string   `json:"message,omitempty"`
	CreatedAt           string   `json:"createdAt,omitempty"`
	MinorEdit           bool     `json:"minorEdit,omitempty"`
	ContentTypeModified bool     `json:"contentTypeModified,omitempty"`
	Collaborators       []string `json:"collaborators,omitempty"`
	PrevVersion         int      `json:"prevVersion,omitempty"`
	NextVersion         int      `json:"nextVersion,omitempty"`
}
