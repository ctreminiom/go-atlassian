package models

type AnnouncementBannerScheme struct {
	HashId        string `json:"hashId,omitempty"`
	IsDismissible bool   `json:"isDismissible,omitempty"`
	IsEnabled     bool   `json:"isEnabled,omitempty"`
	Message       string `json:"message,omitempty"`
	Visibility    string `json:"visibility,omitempty"`
}

type AnnouncementBannerPayloadScheme struct {
	IsDismissible bool   `json:"isDismissible,omitempty"`
	IsEnabled     bool   `json:"isEnabled,omitempty"`
	Message       string `json:"message,omitempty"`
	Visibility    string `json:"visibility,omitempty"`
}
