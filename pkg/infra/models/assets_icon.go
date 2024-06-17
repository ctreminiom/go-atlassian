// Package models provides the data structures used in the asset icon management.
package models

// IconScheme represents an asset icon.
// ID is the unique identifier of the icon.
// Name is the name of the icon.
// URL16 is the URL for the 16x16 version of the icon.
// URL48 is the URL for the 48x48 version of the icon.
type IconScheme struct {
	ID    string `json:"id,omitempty"`    // The ID of the icon.
	Name  string `json:"name,omitempty"`  // The name of the icon.
	URL16 string `json:"url16,omitempty"` // The URL for the 16x16 version of the icon.
	URL48 string `json:"url48,omitempty"` // The URL for the 48x48 version of the icon.
}
