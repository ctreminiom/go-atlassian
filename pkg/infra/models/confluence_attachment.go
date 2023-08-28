package models

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

type AttachmentVersionScheme struct {
	CreatedAt  string                       `json:"createdAt,omitempty"`
	Message    string                       `json:"message,omitempty"`
	Number     int                          `json:"number,omitempty"`
	MinorEdit  bool                         `json:"minorEdit,omitempty"`
	AuthorID   string                       `json:"authorId,omitempty"`
	Attachment *AttachmentVersionBodyScheme `json:"attachment,omitempty"`
}

type AttachmentVersionBodyScheme struct {
	Title string `json:"title,omitempty"`
	ID    string `json:"id,omitempty"`
}

type AttachmentParamsScheme struct {

	// Sort it's used to sort the result by a particular field.
	// Valid values:
	// created-date
	// -created-date
	// modified-date
	// -modified-date
	Sort string

	// MediaType filters the mediaType of attachments. Only one may be specified.
	MediaType string

	// FileName filters on the file-name of attachments. Only one may be specified.
	FileName string

	SerializeIDs bool
}

type AttachmentPageScheme struct {
	Results []*AttachmentScheme `json:"results,omitempty"`
	Links   *PageLinkScheme     `json:"_links,omitempty"`
}

type PageLinkScheme struct {
	Next string `json:"next,omitempty"`
}

type AttachmentVersionPageScheme struct {
	Results []*AttachmentVersionScheme `json:"results,omitempty"`
	Links   *PageLinkScheme            `json:"_links,omitempty"`
}

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
