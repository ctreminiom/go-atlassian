package models

// GetContentAttachmentsOptionsScheme represents the options for getting content attachments.
type GetContentAttachmentsOptionsScheme struct {
	Expand    []string // The fields to expand in the content attachments.
	FileName  string   // The file name of the content attachments.
	MediaType string   // The media type of the content attachments.
}
