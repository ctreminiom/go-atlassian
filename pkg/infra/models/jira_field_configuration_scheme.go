package models

// FieldConfigurationSchemePageScheme represents a page of field configurations in Jira.
type FieldConfigurationSchemePageScheme struct {
	MaxResults int                               `json:"maxResults,omitempty"` // The maximum number of results in the page.
	StartAt    int                               `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                               `json:"total,omitempty"`      // The total number of field configurations.
	IsLast     bool                              `json:"isLast,omitempty"`     // Indicates if the page is the last one.
	Values     []*FieldConfigurationSchemeScheme `json:"values,omitempty"`     // The field configurations in the page.
}

// FieldConfigurationSchemeScheme represents a field configuration in Jira.
type FieldConfigurationSchemeScheme struct {
	ID          string `json:"id,omitempty"`          // The ID of the field configuration.
	Name        string `json:"name,omitempty"`        // The name of the field configuration.
	Description string `json:"description,omitempty"` // The description of the field configuration.
}

// FieldConfigurationIssueTypeItemPageScheme represents a page of field configuration items in Jira.
type FieldConfigurationIssueTypeItemPageScheme struct {
	MaxResults int                                      `json:"maxResults,omitempty"` // The maximum number of results in the page.
	StartAt    int                                      `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                                      `json:"total,omitempty"`      // The total number of field configuration items.
	IsLast     bool                                     `json:"isLast,omitempty"`     // Indicates if the page is the last one.
	Values     []*FieldConfigurationIssueTypeItemScheme `json:"values,omitempty"`     // The field configuration items in the page.
}

// FieldConfigurationIssueTypeItemScheme represents a field configuration item in Jira.
type FieldConfigurationIssueTypeItemScheme struct {
	FieldConfigurationSchemeID string `json:"fieldConfigurationSchemeId,omitempty"` // The ID of the field configuration scheme.
	IssueTypeID                string `json:"issueTypeId,omitempty"`                // The ID of the issue type.
	FieldConfigurationID       string `json:"fieldConfigurationId,omitempty"`       // The ID of the field configuration.
}

// FieldConfigurationSchemeProjectPageScheme represents a page of field configuration scheme projects in Jira.
type FieldConfigurationSchemeProjectPageScheme struct {
	MaxResults int                                      `json:"maxResults,omitempty"` // The maximum number of results in the page.
	StartAt    int                                      `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                                      `json:"total,omitempty"`      // The total number of field configuration scheme projects.
	IsLast     bool                                     `json:"isLast,omitempty"`     // Indicates if the page is the last one.
	Values     []*FieldConfigurationSchemeProjectScheme `json:"values,omitempty"`     // The field configuration scheme projects in the page.
}

// FieldConfigurationSchemeProjectScheme represents a field configuration scheme project in Jira.
type FieldConfigurationSchemeProjectScheme struct {
	ProjectIDs               []string                        `json:"projectIds,omitempty"`               // The IDs of the projects.
	FieldConfigurationScheme *FieldConfigurationSchemeScheme `json:"fieldConfigurationScheme,omitempty"` // The field configuration scheme.
}

// FieldConfigurationSchemeAssignPayload represents the payload for assigning a field configuration scheme in Jira.
type FieldConfigurationSchemeAssignPayload struct {
	FieldConfigurationSchemeID string `json:"fieldConfigurationSchemeId"` // The ID of the field configuration scheme.
	ProjectID                  string `json:"projectId"`                  // The ID of the project.
}

// FieldConfigurationToIssueTypeMappingPayloadScheme represents the payload for mapping a field configuration to an issue type in Jira.
type FieldConfigurationToIssueTypeMappingPayloadScheme struct {
	Mappings []*FieldConfigurationToIssueTypeMappingScheme `json:"mappings,omitempty"` // The mappings.
}

// FieldConfigurationToIssueTypeMappingScheme represents a mapping of a field configuration to an issue type in Jira.
type FieldConfigurationToIssueTypeMappingScheme struct {
	IssueTypeID          string `json:"issueTypeId,omitempty"`          // The ID of the issue type.
	FieldConfigurationID string `json:"fieldConfigurationId,omitempty"` // The ID of the field configuration.
}
