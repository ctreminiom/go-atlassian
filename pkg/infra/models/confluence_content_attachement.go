package models

// GetContentAttachmentsOptionsScheme represents the options for getting content attachments.
// Expand is a slice of strings representing the fields to expand in the content attachments.
// FileName is a string representing the file name of the content attachments.
// MediaType is a string representing the media type of the content attachments.
type GetContentAttachmentsOptionsScheme struct {
	Expand    []string // The fields to expand in the content attachments.
	FileName  string   // The file name of the content attachments.
	MediaType string   // The media type of the content attachments.
}
