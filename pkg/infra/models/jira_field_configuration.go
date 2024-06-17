package models

// FieldConfigurationPageScheme represents a page of field configurations in Jira.
type FieldConfigurationPageScheme struct {
	MaxResults int                         `json:"maxResults,omitempty"` // The maximum number of results in the page.
	StartAt    int                         `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                         `json:"total,omitempty"`      // The total number of field configurations.
	IsLast     bool                        `json:"isLast,omitempty"`     // Indicates if the page is the last one.
	Values     []*FieldConfigurationScheme `json:"values,omitempty"`     // The field configurations in the page.
}

// FieldConfigurationScheme represents a field configuration in Jira.
type FieldConfigurationScheme struct {
	ID          int    `json:"id,omitempty"`          // The ID of the field configuration.
	Name        string `json:"name,omitempty"`        // The name of the field configuration.
	Description string `json:"description,omitempty"` // The description of the field configuration.
	IsDefault   bool   `json:"isDefault,omitempty"`   // Indicates if the field configuration is the default one.
}

// FieldConfigurationItemPageScheme represents a page of field configuration items in Jira.
type FieldConfigurationItemPageScheme struct {
	MaxResults int                             `json:"maxResults,omitempty"` // The maximum number of results in the page.
	StartAt    int                             `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                             `json:"total,omitempty"`      // The total number of field configuration items.
	IsLast     bool                            `json:"isLast,omitempty"`     // Indicates if the page is the last one.
	Values     []*FieldConfigurationItemScheme `json:"values,omitempty"`     // The field configuration items in the page.
}

// UpdateFieldConfigurationItemPayloadScheme represents the payload for updating a field configuration item in Jira.
type UpdateFieldConfigurationItemPayloadScheme struct {
	FieldConfigurationItems []*FieldConfigurationItemScheme `json:"fieldConfigurationItems"` // The field configuration items to be updated.
}

// FieldConfigurationItemScheme represents a field configuration item in Jira.
type FieldConfigurationItemScheme struct {
	ID          string `json:"id,omitempty"`          // The ID of the field configuration item.
	IsHidden    bool   `json:"isHidden,omitempty"`    // Indicates if the field configuration item is hidden.
	IsRequired  bool   `json:"isRequired,omitempty"`  // Indicates if the field configuration item is required.
	Description string `json:"description,omitempty"` // The description of the field configuration item.
	Renderer    string `json:"renderer,omitempty"`    // The renderer of the field configuration item.
}
