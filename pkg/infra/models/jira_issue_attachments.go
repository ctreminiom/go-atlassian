package models

type AttachmentSettingScheme struct {
	Enabled     bool `json:"enabled,omitempty"`
	UploadLimit int  `json:"uploadLimit,omitempty"`
}

type IssueAttachmentScheme struct {
	Self      string      `json:"self,omitempty"`
	ID        string      `json:"id,omitempty"`
	Filename  string      `json:"filename,omitempty"`
	Author    *UserScheme `json:"author,omitempty"`
	Created   string      `json:"created,omitempty"`
	Size      int         `json:"size,omitempty"`
	MimeType  string      `json:"mimeType,omitempty"`
	Content   string      `json:"content,omitempty"`
	Thumbnail string      `json:"thumbnail,omitempty"`
}

type IssueAttachmentMetadataScheme struct {
	ID        int         `json:"id,omitempty"`
	Self      string      `json:"self,omitempty"`
	Filename  string      `json:"filename,omitempty"`
	Author    *UserScheme `json:"author,omitempty"`
	Created   string      `json:"created,omitempty"`
	Size      int         `json:"size,omitempty"`
	MimeType  string      `json:"mimeType,omitempty"`
	Content   string      `json:"content,omitempty"`
	Thumbnail string      `json:"thumbnail,omitempty"`
}

type IssueAttachmentHumanMetadataScheme struct {
	ID              int                                        `json:"id,omitempty"`
	Name            string                                     `json:"name,omitempty"`
	Entries         []*IssueAttachmentHumanMetadataEntryScheme `json:"entries,omitempty"`
	TotalEntryCount int                                        `json:"totalEntryCount,omitempty"`
	MediaType       string                                     `json:"mediaType,omitempty"`
}

type IssueAttachmentHumanMetadataEntryScheme struct {
	Path      string `json:"path,omitempty"`
	Index     int    `json:"index,omitempty"`
	Size      string `json:"size,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
	Label     string `json:"label,omitempty"`
}
