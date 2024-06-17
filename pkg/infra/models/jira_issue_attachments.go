package models

// AttachmentSettingScheme represents the attachment settings in Jira.
type AttachmentSettingScheme struct {
	Enabled     bool `json:"enabled,omitempty"`     // Indicates if attachments are enabled.
	UploadLimit int  `json:"uploadLimit,omitempty"` // The upload limit for attachments.
}

// IssueAttachmentScheme represents an attachment of an issue in Jira.
type IssueAttachmentScheme struct {
	Self      string      `json:"self,omitempty"`      // The URL of the attachment.
	ID        string      `json:"id,omitempty"`        // The ID of the attachment.
	Filename  string      `json:"filename,omitempty"`  // The filename of the attachment.
	Author    *UserScheme `json:"author,omitempty"`    // The author of the attachment.
	Created   string      `json:"created,omitempty"`   // The creation time of the attachment.
	Size      int         `json:"size,omitempty"`      // The size of the attachment.
	MimeType  string      `json:"mimeType,omitempty"`  // The MIME type of the attachment.
	Content   string      `json:"content,omitempty"`   // The content of the attachment.
	Thumbnail string      `json:"thumbnail,omitempty"` // The thumbnail of the attachment.
}

// IssueAttachmentMetadataScheme represents the metadata of an attachment of an issue in Jira.
type IssueAttachmentMetadataScheme struct {
	ID        int         `json:"id,omitempty"`        // The ID of the attachment.
	Self      string      `json:"self,omitempty"`      // The URL of the attachment.
	Filename  string      `json:"filename,omitempty"`  // The filename of the attachment.
	Author    *UserScheme `json:"author,omitempty"`    // The author of the attachment.
	Created   string      `json:"created,omitempty"`   // The creation time of the attachment.
	Size      int         `json:"size,omitempty"`      // The size of the attachment.
	MimeType  string      `json:"mimeType,omitempty"`  // The MIME type of the attachment.
	Content   string      `json:"content,omitempty"`   // The content of the attachment.
	Thumbnail string      `json:"thumbnail,omitempty"` // The thumbnail of the attachment.
}

// IssueAttachmentHumanMetadataScheme represents the human-readable metadata of an attachment of an issue in Jira.
type IssueAttachmentHumanMetadataScheme struct {
	ID              int                                        `json:"id,omitempty"`              // The ID of the attachment.
	Name            string                                     `json:"name,omitempty"`            // The name of the attachment.
	Entries         []*IssueAttachmentHumanMetadataEntryScheme `json:"entries,omitempty"`         // The entries of the attachment.
	TotalEntryCount int                                        `json:"totalEntryCount,omitempty"` // The total count of entries of the attachment.
	MediaType       string                                     `json:"mediaType,omitempty"`       // The media type of the attachment.
}

// IssueAttachmentHumanMetadataEntryScheme represents an entry of the human-readable metadata of an attachment of an issue in Jira.
type IssueAttachmentHumanMetadataEntryScheme struct {
	Path      string `json:"path,omitempty"`      // The path of the entry.
	Index     int    `json:"index,omitempty"`     // The index of the entry.
	Size      string `json:"size,omitempty"`      // The size of the entry.
	MediaType string `json:"mediaType,omitempty"` // The media type of the entry.
	Label     string `json:"label,omitempty"`     // The label of the entry.
}
