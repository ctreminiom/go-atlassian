package models

// AnnouncementBannerScheme represents an announcement banner in Jira.
type AnnouncementBannerScheme struct {
	HashID        string `json:"hashId,omitempty"`        // The hash ID of the banner.
	IsDismissible bool   `json:"isDismissible,omitempty"` // Indicates if the banner is dismissible.
	IsEnabled     bool   `json:"isEnabled,omitempty"`     // Indicates if the banner is enabled.
	Message       string `json:"message,omitempty"`       // The message of the banner.
	Visibility    string `json:"visibility,omitempty"`    // The visibility of the banner.
}

// AnnouncementBannerPayloadScheme represents the payload for an announcement banner in Jira.
type AnnouncementBannerPayloadScheme struct {
	IsDismissible bool   `json:"isDismissible,omitempty"` // Indicates if the banner is dismissible.
	IsEnabled     bool   `json:"isEnabled,omitempty"`     // Indicates if the banner is enabled.
	Message       string `json:"message,omitempty"`       // The message of the banner.
	Visibility    string `json:"visibility,omitempty"`    // The visibility of the banner.
}
