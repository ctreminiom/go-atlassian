package models

type FieldContextOptionsScheme struct {
	IsAnyIssueType  bool
	IsGlobalContext bool
	ContextID       []int
}

type CustomFieldContextPageScheme struct {
	MaxResults int                   `json:"maxResults,omitempty"`
	StartAt    int                   `json:"startAt,omitempty"`
	Total      int                   `json:"total,omitempty"`
	IsLast     bool                  `json:"isLast,omitempty"`
	Values     []*FieldContextScheme `json:"values,omitempty"`
}

type FieldContextScheme struct {
	ID              string   `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Description     string   `json:"description,omitempty"`
	IsGlobalContext bool     `json:"isGlobalContext,omitempty"`
	IsAnyIssueType  bool     `json:"isAnyIssueType,omitempty"`
	ProjectIds      []string `json:"projectIds,omitempty"`
	IssueTypeIds    []string `json:"issueTypeIds,omitempty"`
}

type CustomFieldDefaultValuePageScheme struct {
	MaxResults int                              `json:"maxResults,omitempty"`
	StartAt    int                              `json:"startAt,omitempty"`
	Total      int                              `json:"total,omitempty"`
	IsLast     bool                             `json:"isLast,omitempty"`
	Values     []*CustomFieldDefaultValueScheme `json:"values,omitempty"`
}

type CustomFieldDefaultValueScheme struct {
	ContextID         string   `json:"contextId,omitempty"`
	OptionID          string   `json:"optionId,omitempty"`
	CascadingOptionID string   `json:"cascadingOptionId,omitempty"`
	OptionIDs         []string `json:"optionIds,omitempty"`
	Type              string   `json:"type,omitempty"`
}

type FieldContextDefaultPayloadScheme struct {
	DefaultValues []*CustomFieldDefaultValueScheme `json:"defaultValues,omitempty"`
}

type IssueTypeToContextMappingPageScheme struct {
	MaxResults int                                     `json:"maxResults,omitempty"`
	StartAt    int                                     `json:"startAt,omitempty"`
	Total      int                                     `json:"total,omitempty"`
	IsLast     bool                                    `json:"isLast,omitempty"`
	Values     []*IssueTypeToContextMappingValueScheme `json:"values"`
}

type IssueTypeToContextMappingValueScheme struct {
	ContextID      string `json:"contextId"`
	IsAnyIssueType bool   `json:"isAnyIssueType,omitempty"`
	IssueTypeID    string `json:"issueTypeId,omitempty"`
}

type CustomFieldContextProjectMappingPageScheme struct {
	Self       string                                         `json:"self,omitempty"`
	NextPage   string                                         `json:"nextPage,omitempty"`
	MaxResults int                                            `json:"maxResults,omitempty"`
	StartAt    int                                            `json:"startAt,omitempty"`
	Total      int                                            `json:"total,omitempty"`
	IsLast     bool                                           `json:"isLast,omitempty"`
	Values     []*CustomFieldContextProjectMappingValueScheme `json:"values,omitempty"`
}

type CustomFieldContextProjectMappingValueScheme struct {
	ContextID       string `json:"contextId,omitempty"`
	ProjectID       string `json:"projectId,omitempty"`
	IsGlobalContext bool   `json:"isGlobalContext,omitempty"`
}

type FieldContextPayloadScheme struct {
	IssueTypeIDs []int  `json:"issueTypeIds,omitempty"`
	ProjectIDs   []int  `json:"projectIds,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
}
