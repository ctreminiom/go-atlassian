package models

// ProjectTypeScheme represents a project type in Jira.
type ProjectTypeScheme struct {
	Key                string `json:"key"`                // The key of the project type.
	FormattedKey       string `json:"formattedKey"`       // The formatted key of the project type.
	DescriptionI18NKey string `json:"descriptionI18nKey"` // The internationalized description key of the project type.
	Icon               string `json:"icon"`               // The icon of the project type.
	Color              string `json:"color"`              // The color associated with the project type.
}
