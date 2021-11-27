package models

type IssueTypeSchemePayloadScheme struct {
	DefaultIssueTypeID string   `json:"defaultIssueTypeId,omitempty"`
	IssueTypeIds       []string `json:"issueTypeIds,omitempty"`
	Name               string   `json:"name,omitempty"`
	Description        string   `json:"description,omitempty"`
}

type NewIssueTypeSchemeScheme struct {
	IssueTypeSchemeID string `json:"issueTypeSchemeId"`
}

type IssueTypeSchemeItemPageScheme struct {
	MaxResults int                             `json:"maxResults,omitempty"`
	StartAt    int                             `json:"startAt,omitempty"`
	Total      int                             `json:"total,omitempty"`
	IsLast     bool                            `json:"isLast,omitempty"`
	Values     []*IssueTypeSchemeMappingScheme `json:"values,omitempty"`
}

type IssueTypeSchemeMappingScheme struct {
	IssueTypeSchemeID string `json:"issueTypeSchemeId,omitempty"`
	IssueTypeID       string `json:"issueTypeId,omitempty"`
}

type ProjectIssueTypeSchemePageScheme struct {
	MaxResults int                              `json:"maxResults"`
	StartAt    int                              `json:"startAt"`
	Total      int                              `json:"total"`
	IsLast     bool                             `json:"isLast"`
	Values     []*IssueTypeSchemeProjectsScheme `json:"values"`
}

type IssueTypeSchemeProjectsScheme struct {
	IssueTypeScheme *IssueTypeSchemeScheme `json:"issueTypeScheme,omitempty"`
	ProjectIds      []string               `json:"projectIds,omitempty"`
}
