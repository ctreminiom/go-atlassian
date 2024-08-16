package models

// AttachmentScheme represents an attachment in Confluence.
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
type AttachmentVersionScheme struct {
	CreatedAt  string                       `json:"createdAt,omitempty"`
	Message    string                       `json:"message,omitempty"`
	Number     int                          `json:"number,omitempty"`
	MinorEdit  bool                         `json:"minorEdit,omitempty"`
	AuthorID   string                       `json:"authorId,omitempty"`
	Attachment *AttachmentVersionBodyScheme `json:"attachment,omitempty"`
}

// AttachmentVersionBodyScheme represents the body of a version of an attachment in Confluence.
type AttachmentVersionBodyScheme struct {
	Title string `json:"title,omitempty"`
	ID    string `json:"id,omitempty"`
}

// AttachmentParamsScheme represents the parameters for an attachment in Confluence.
type AttachmentParamsScheme struct {
	Sort         string
	MediaType    string
	FileName     string
	SerializeIDs bool
}

// AttachmentPageScheme represents a paginated list of attachments in Confluence.
type AttachmentPageScheme struct {
	Results []*AttachmentScheme `json:"results,omitempty"`
	Links   *PageLinkScheme     `json:"_links,omitempty"`
}

// PageLinkScheme represents a link in a page in Confluence.
type PageLinkScheme struct {
	Next string `json:"next,omitempty"`
}

// AttachmentVersionPageScheme represents a paginated list of versions of an attachment in Confluence.
type AttachmentVersionPageScheme struct {
	Results []*AttachmentVersionScheme `json:"results,omitempty"`
	Links   *PageLinkScheme            `json:"_links,omitempty"`
}

// DetailedVersionScheme represents a detailed version in Confluence.
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
