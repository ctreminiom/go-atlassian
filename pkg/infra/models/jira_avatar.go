package models

// AvatarURLScheme represents the URLs for different sizes of an avatar in Jira.
type AvatarURLScheme struct {
	Four8X48  string `json:"48x48,omitempty"` // The URL for the 48x48 size of the avatar.
	Two4X24   string `json:"24x24,omitempty"` // The URL for the 24x24 size of the avatar.
	One6X16   string `json:"16x16,omitempty"` // The URL for the 16x16 size of the avatar.
	Three2X32 string `json:"32x32,omitempty"` // The URL for the 32x32 size of the avatar.
}
