package models

// IssueTypeSchemePayloadScheme represents the payload for an issue type scheme in Jira.
type IssueTypeSchemePayloadScheme struct {
	DefaultIssueTypeID string   `json:"defaultIssueTypeId,omitempty"` // The default issue type ID.
	IssueTypeIDs       []string `json:"issueTypeIds,omitempty"`       // The issue type IDs.
	Name               string   `json:"name,omitempty"`               // The name of the issue type scheme.
	Description        string   `json:"description,omitempty"`        // The description of the issue type scheme.
}

// NewIssueTypeSchemeScheme represents a new issue type scheme in Jira.
type NewIssueTypeSchemeScheme struct {
	IssueTypeSchemeID string `json:"issueTypeSchemeId"` // The ID of the issue type scheme.
}

// IssueTypeSchemeItemPageScheme represents a page of issue type scheme items in Jira.
type IssueTypeSchemeItemPageScheme struct {
	MaxResults int                             `json:"maxResults,omitempty"` // The maximum results per page.
	StartAt    int                             `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                             `json:"total,omitempty"`      // The total number of issue type scheme items.
	IsLast     bool                            `json:"isLast,omitempty"`     // Indicates if this is the last page.
	Values     []*IssueTypeSchemeMappingScheme `json:"values,omitempty"`     // The issue type scheme items in the page.
}

// IssueTypeSchemeMappingScheme represents a mapping of an issue type scheme in Jira.
type IssueTypeSchemeMappingScheme struct {
	IssueTypeSchemeID string `json:"issueTypeSchemeId,omitempty"` // The ID of the issue type scheme.
	IssueTypeID       string `json:"issueTypeId,omitempty"`       // The ID of the issue type.
}

// ProjectIssueTypeSchemePageScheme represents a page of project issue type schemes in Jira.
type ProjectIssueTypeSchemePageScheme struct {
	MaxResults int                              `json:"maxResults"` // The maximum results per page.
	StartAt    int                              `json:"startAt"`    // The starting index of the page.
	Total      int                              `json:"total"`      // The total number of project issue type schemes.
	IsLast     bool                             `json:"isLast"`     // Indicates if this is the last page.
	Values     []*IssueTypeSchemeProjectsScheme `json:"values"`     // The project issue type schemes in the page.
}

// IssueTypeSchemeProjectsScheme represents the projects of an issue type scheme in Jira.
type IssueTypeSchemeProjectsScheme struct {
	IssueTypeScheme *IssueTypeSchemeScheme `json:"issueTypeScheme,omitempty"` // The issue type scheme.
	ProjectIDs      []string               `json:"projectIds,omitempty"`      // The IDs of the projects.
}
