package models

// FieldContextOptionsScheme represents the options for a field context in Jira.
type FieldContextOptionsScheme struct {
	IsAnyIssueType  bool  // Indicates if any issue type is applicable.
	IsGlobalContext bool  // Indicates if the context is global.
	ContextID       []int // The IDs of the contexts.
}

// CustomFieldContextPageScheme represents a page of custom field contexts in Jira.
type CustomFieldContextPageScheme struct {
	MaxResults int                   `json:"maxResults,omitempty"` // The maximum number of results in the page.
	StartAt    int                   `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                   `json:"total,omitempty"`      // The total number of custom field contexts.
	IsLast     bool                  `json:"isLast,omitempty"`     // Indicates if the page is the last one.
	Values     []*FieldContextScheme `json:"values,omitempty"`     // The custom field contexts in the page.
}

// FieldContextScheme represents a field context in Jira.
type FieldContextScheme struct {
	ID              string   `json:"id,omitempty"`              // The ID of the field context.
	Name            string   `json:"name,omitempty"`            // The name of the field context.
	Description     string   `json:"description,omitempty"`     // The description of the field context.
	IsGlobalContext bool     `json:"isGlobalContext,omitempty"` // Indicates if the context is global.
	IsAnyIssueType  bool     `json:"isAnyIssueType,omitempty"`  // Indicates if any issue type is applicable.
	ProjectIDs      []string `json:"projectIds,omitempty"`      // The IDs of the projects.
	IssueTypeIDs    []string `json:"issueTypeIds,omitempty"`    // The IDs of the issue types.
}

// CustomFieldDefaultValuePageScheme represents a page of default values for custom fields in Jira.
type CustomFieldDefaultValuePageScheme struct {
	MaxResults int                              `json:"maxResults,omitempty"` // The maximum number of results in the page.
	StartAt    int                              `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                              `json:"total,omitempty"`      // The total number of default values for custom fields.
	IsLast     bool                             `json:"isLast,omitempty"`     // Indicates if the page is the last one.
	Values     []*CustomFieldDefaultValueScheme `json:"values,omitempty"`     // The default values for custom fields in the page.
}

// CustomFieldDefaultValueScheme represents a default value for a custom field in Jira.
type CustomFieldDefaultValueScheme struct {
	ContextID         string   `json:"contextId,omitempty"`         // The ID of the context.
	OptionID          string   `json:"optionId,omitempty"`          // The ID of the option.
	CascadingOptionID string   `json:"cascadingOptionId,omitempty"` // The ID of the cascading option.
	OptionIDs         []string `json:"optionIds,omitempty"`         // The IDs of the options.
	Type              string   `json:"type,omitempty"`              // The type of the default value.
}

// FieldContextDefaultPayloadScheme represents the payload for a default field context in Jira.
type FieldContextDefaultPayloadScheme struct {
	DefaultValues []*CustomFieldDefaultValueScheme `json:"defaultValues,omitempty"` // The default values for the field context.
}

// IssueTypeToContextMappingPageScheme represents a page of issue type to context mappings in Jira.
type IssueTypeToContextMappingPageScheme struct {
	MaxResults int                                     `json:"maxResults,omitempty"` // The maximum number of results in the page.
	StartAt    int                                     `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                                     `json:"total,omitempty"`      // The total number of issue type to context mappings.
	IsLast     bool                                    `json:"isLast,omitempty"`     // Indicates if the page is the last one.
	Values     []*IssueTypeToContextMappingValueScheme `json:"values"`               // The issue type to context mappings in the page.
}

// IssueTypeToContextMappingValueScheme represents an issue type to context mapping in Jira.
type IssueTypeToContextMappingValueScheme struct {
	ContextID      string `json:"contextId"`                // The ID of the context.
	IsAnyIssueType bool   `json:"isAnyIssueType,omitempty"` // Indicates if any issue type is applicable.
	IssueTypeID    string `json:"issueTypeId,omitempty"`    // The ID of the issue type.
}

// CustomFieldContextProjectMappingPageScheme represents a page of project mappings for custom field contexts in Jira.
type CustomFieldContextProjectMappingPageScheme struct {
	Self       string                                         `json:"self,omitempty"`       // The URL of the page.
	NextPage   string                                         `json:"nextPage,omitempty"`   // The URL of the next page.
	MaxResults int                                            `json:"maxResults,omitempty"` // The maximum number of results in the page.
	StartAt    int                                            `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                                            `json:"total,omitempty"`      // The total number of project mappings for custom field contexts.
	IsLast     bool                                           `json:"isLast,omitempty"`     // Indicates if the page is the last one.
	Values     []*CustomFieldContextProjectMappingValueScheme `json:"values,omitempty"`     // The project mappings for custom field contexts in the page.
}

// CustomFieldContextProjectMappingValueScheme represents a project mapping for a custom field context in Jira.
type CustomFieldContextProjectMappingValueScheme struct {
	ContextID       string `json:"contextId,omitempty"`       // The ID of the context.
	ProjectID       string `json:"projectId,omitempty"`       // The ID of the project.
	IsGlobalContext bool   `json:"isGlobalContext,omitempty"` // Indicates if the context is global.
}

// FieldContextPayloadScheme represents the payload for a field context in Jira.
type FieldContextPayloadScheme struct {
	IssueTypeIDs []int  `json:"issueTypeIds,omitempty"` // The IDs of the issue types.
	ProjectIDs   []int  `json:"projectIds,omitempty"`   // The IDs of the projects.
	Name         string `json:"name,omitempty"`         // The name of the field context.
	Description  string `json:"description,omitempty"`  // The description of the field context.
}
