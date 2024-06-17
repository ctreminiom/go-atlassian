package models

// VersionScheme represents a version in Jira.
type VersionScheme struct {
	Self                      string                                  `json:"self,omitempty"`                      // The URL of the version.
	ID                        string                                  `json:"id,omitempty"`                        // The ID of the version.
	Description               string                                  `json:"description,omitempty"`               // The description of the version.
	Name                      string                                  `json:"name,omitempty"`                      // The name of the version.
	Archived                  bool                                    `json:"archived,omitempty"`                  // Indicates if the version is archived.
	Released                  bool                                    `json:"released,omitempty"`                  // Indicates if the version is released.
	ReleaseDate               string                                  `json:"releaseDate,omitempty"`               // The release date of the version.
	Overdue                   bool                                    `json:"overdue,omitempty"`                   // Indicates if the version is overdue.
	UserReleaseDate           string                                  `json:"userReleaseDate,omitempty"`           // The user release date of the version.
	ProjectID                 int                                     `json:"projectId,omitempty"`                 // The project ID of the version.
	Operations                []*VersionOperation                     `json:"operations,omitempty"`                // The operations of the version.
	IssuesStatusForFixVersion *VersionIssuesStatusForFixVersionScheme `json:"issuesStatusForFixVersion,omitempty"` // The issues status for fix version of the version.
}

// VersionOperation represents an operation in a version in Jira.
type VersionOperation struct {
	ID         string `json:"id,omitempty"`         // The ID of the operation.
	StyleClass string `json:"styleClass,omitempty"` // The style class of the operation.
	Label      string `json:"label,omitempty"`      // The label of the operation.
	Href       string `json:"href,omitempty"`       // The href of the operation.
	Weight     int    `json:"weight,omitempty"`     // The weight of the operation.
}

// VersionIssuesStatusForFixVersionScheme represents the issues status for fix version in a version in Jira.
type VersionIssuesStatusForFixVersionScheme struct {
	Unmapped   int `json:"unmapped,omitempty"`   // The unmapped status.
	ToDo       int `json:"toDo,omitempty"`       // The to do status.
	InProgress int `json:"inProgress,omitempty"` // The in progress status.
	Done       int `json:"done,omitempty"`       // The done status.
}

// VersionPageScheme represents a page of versions in Jira.
type VersionPageScheme struct {
	Self       string           `json:"self,omitempty"`       // The URL of the page.
	NextPage   string           `json:"nextPage,omitempty"`   // The URL of the next page.
	MaxResults int              `json:"maxResults,omitempty"` // The maximum number of results returned.
	StartAt    int              `json:"startAt,omitempty"`    // The index of the first result returned.
	Total      int              `json:"total,omitempty"`      // The total number of results available.
	IsLast     bool             `json:"isLast,omitempty"`     // Indicates if this is the last page of results.
	Values     []*VersionScheme `json:"values,omitempty"`     // The versions on the page.
}

// VersionGetsOptions represents the options for getting versions in Jira.
type VersionGetsOptions struct {
	OrderBy string   // The order by option.
	Query   string   // The query option.
	Status  string   // The status option.
	Expand  []string // The expand option.
}

// VersionPayloadScheme represents the payload for a version in Jira.
type VersionPayloadScheme struct {
	Archived    bool   `json:"archived,omitempty"`    // Indicates if the version is archived.
	ReleaseDate string `json:"releaseDate,omitempty"` // The release date of the version.
	Name        string `json:"name,omitempty"`        // The name of the version.
	Description string `json:"description,omitempty"` // The description of the version.
	ProjectID   int    `json:"projectId,omitempty"`   // The project ID of the version.
	Released    bool   `json:"released,omitempty"`    // Indicates if the version is released.
	StartDate   string `json:"startDate,omitempty"`   // The start date of the version.
}

// VersionIssueCountsScheme represents the issue counts for a version in Jira.
type VersionIssueCountsScheme struct {
	Self                                     string                                     `json:"self,omitempty"`                                     // The URL of the issue counts.
	IssuesFixedCount                         int                                        `json:"issuesFixedCount,omitempty"`                         // The count of issues fixed.
	IssuesAffectedCount                      int                                        `json:"issuesAffectedCount,omitempty"`                      // The count of issues affected.
	IssueCountWithCustomFieldsShowingVersion int                                        `json:"issueCountWithCustomFieldsShowingVersion,omitempty"` // The count of issues with custom fields showing version.
	CustomFieldUsage                         []*VersionIssueCountCustomFieldUsageScheme `json:"customFieldUsage,omitempty"`                         // The custom field usage of the version.
}

// VersionIssueCountCustomFieldUsageScheme represents the custom field usage for a version in Jira.
type VersionIssueCountCustomFieldUsageScheme struct {
	FieldName                          string `json:"fieldName,omitempty"`                          // The field name of the custom field usage.
	CustomFieldID                      int    `json:"customFieldId,omitempty"`                      // The custom field ID of the custom field usage.
	IssueCountWithVersionInCustomField int    `json:"issueCountWithVersionInCustomField,omitempty"` // The issue count with version in custom field of the custom field usage.
}

// VersionUnresolvedIssuesCountScheme represents the unresolved issues count for a version in Jira.
type VersionUnresolvedIssuesCountScheme struct {
	Self                  string `json:"self"`                  // The URL of the unresolved issues count.
	IssuesUnresolvedCount int    `json:"issuesUnresolvedCount"` // The count of unresolved issues.
	IssuesCount           int    `json:"issuesCount"`           // The count of issues.
}

// VersionDetailScheme represents the detail of a version in Jira.
type VersionDetailScheme struct {
	Self        string `json:"self,omitempty"`        // The URL of the detail.
	ID          string `json:"id,omitempty"`          // The ID of the detail.
	Description string `json:"description,omitempty"` // The description of the detail.
	Name        string `json:"name,omitempty"`        // The name of the detail.
	Archived    bool   `json:"archived,omitempty"`    // Indicates if the detail is archived.
	Released    bool   `json:"released,omitempty"`    // Indicates if the detail is released.
	ReleaseDate string `json:"releaseDate,omitempty"` // The release date of the detail.
}
