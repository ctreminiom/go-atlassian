package models

type FieldConfigurationSchemePageScheme struct {
	MaxResults int                               `json:"maxResults,omitempty"`
	StartAt    int                               `json:"startAt,omitempty"`
	Total      int                               `json:"total,omitempty"`
	IsLast     bool                              `json:"isLast,omitempty"`
	Values     []*FieldConfigurationSchemeScheme `json:"values,omitempty"`
}

type FieldConfigurationSchemeScheme struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type FieldConfigurationIssueTypeItemPageScheme struct {
	MaxResults int                                      `json:"maxResults,omitempty"`
	StartAt    int                                      `json:"startAt,omitempty"`
	Total      int                                      `json:"total,omitempty"`
	IsLast     bool                                     `json:"isLast,omitempty"`
	Values     []*FieldConfigurationIssueTypeItemScheme `json:"values,omitempty"`
}

type FieldConfigurationIssueTypeItemScheme struct {
	FieldConfigurationSchemeID string `json:"fieldConfigurationSchemeId,omitempty"`
	IssueTypeID                string `json:"issueTypeId,omitempty"`
	FieldConfigurationID       string `json:"fieldConfigurationId,omitempty"`
}

type FieldConfigurationSchemeProjectPageScheme struct {
	MaxResults int                                      `json:"maxResults,omitempty"`
	StartAt    int                                      `json:"startAt,omitempty"`
	Total      int                                      `json:"total,omitempty"`
	IsLast     bool                                     `json:"isLast,omitempty"`
	Values     []*FieldConfigurationSchemeProjectScheme `json:"values,omitempty"`
}

type FieldConfigurationSchemeProjectScheme struct {
	ProjectIds               []string                        `json:"projectIds,omitempty"`
	FieldConfigurationScheme *FieldConfigurationSchemeScheme `json:"fieldConfigurationScheme,omitempty"`
}

type FieldConfigurationSchemeAssignPayload struct {
	FieldConfigurationSchemeID string `json:"fieldConfigurationSchemeId"`
	ProjectID                  string `json:"projectId"`
}

type FieldConfigurationToIssueTypeMappingPayloadScheme struct {
	Mappings []*FieldConfigurationToIssueTypeMappingScheme `json:"mappings,omitempty"`
}

type FieldConfigurationToIssueTypeMappingScheme struct {
	IssueTypeID          string `json:"issueTypeId,omitempty"`
	FieldConfigurationID string `json:"fieldConfigurationId,omitempty"`
}
