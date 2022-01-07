package models

type FieldConfigurationPageScheme struct {
	MaxResults int                         `json:"maxResults,omitempty"`
	StartAt    int                         `json:"startAt,omitempty"`
	Total      int                         `json:"total,omitempty"`
	IsLast     bool                        `json:"isLast,omitempty"`
	Values     []*FieldConfigurationScheme `json:"values,omitempty"`
}

type FieldConfigurationScheme struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsDefault   bool   `json:"isDefault,omitempty"`
}

type FieldConfigurationItemPageScheme struct {
	MaxResults int                             `json:"maxResults,omitempty"`
	StartAt    int                             `json:"startAt,omitempty"`
	Total      int                             `json:"total,omitempty"`
	IsLast     bool                            `json:"isLast,omitempty"`
	Values     []*FieldConfigurationItemScheme `json:"values,omitempty"`
}

type UpdateFieldConfigurationItemPayloadScheme struct {
	FieldConfigurationItems []*FieldConfigurationItemScheme `json:"fieldConfigurationItems"`
}

type FieldConfigurationItemScheme struct {
	ID          string `json:"id,omitempty"`
	IsHidden    bool   `json:"isHidden,omitempty"`
	IsRequired  bool   `json:"isRequired,omitempty"`
	Description string `json:"description,omitempty"`
	Renderer    string `json:"renderer,omitempty"`
}
