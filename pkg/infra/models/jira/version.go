package jira

type VersionScheme struct {
	Self                      string                                  `json:"self,omitempty"`
	ID                        string                                  `json:"id,omitempty"`
	Description               string                                  `json:"description,omitempty"`
	Name                      string                                  `json:"name,omitempty"`
	Archived                  bool                                    `json:"archived,omitempty"`
	Released                  bool                                    `json:"released,omitempty"`
	ReleaseDate               string                                  `json:"releaseDate,omitempty"`
	Overdue                   bool                                    `json:"overdue,omitempty"`
	UserReleaseDate           string                                  `json:"userReleaseDate,omitempty"`
	ProjectID                 int                                     `json:"projectId,omitempty"`
	Operations                []*VersionOperation                     `json:"operations,omitempty"`
	IssuesStatusForFixVersion *VersionIssuesStatusForFixVersionScheme `json:"issuesStatusForFixVersion,omitempty"`
}

type VersionOperation struct {
	ID         string `json:"id,omitempty"`
	StyleClass string `json:"styleClass,omitempty"`
	Label      string `json:"label,omitempty"`
	Href       string `json:"href,omitempty"`
	Weight     int    `json:"weight,omitempty"`
}

type VersionIssuesStatusForFixVersionScheme struct {
	Unmapped   int `json:"unmapped,omitempty"`
	ToDo       int `json:"toDo,omitempty"`
	InProgress int `json:"inProgress,omitempty"`
	Done       int `json:"done,omitempty"`
}

type VersionPageScheme struct {
	Self       string           `json:"self,omitempty"`
	NextPage   string           `json:"nextPage,omitempty"`
	MaxResults int              `json:"maxResults,omitempty"`
	StartAt    int              `json:"startAt,omitempty"`
	Total      int              `json:"total,omitempty"`
	IsLast     bool             `json:"isLast,omitempty"`
	Values     []*VersionScheme `json:"values,omitempty"`
}

type VersionGetsOptions struct {
	OrderBy string
	Query   string
	Status  string
	Expand  []string
}

type VersionPayloadScheme struct {
	Archived    bool   `json:"archived,omitempty"`
	ReleaseDate string `json:"releaseDate,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ProjectID   int    `json:"projectId,omitempty"`
	Released    bool   `json:"released,omitempty"`
	StartDate   string `json:"startDate,omitempty"`
}

type VersionIssueCountsScheme struct {
	Self                                     string                                     `json:"self,omitempty"`
	IssuesFixedCount                         int                                        `json:"issuesFixedCount,omitempty"`
	IssuesAffectedCount                      int                                        `json:"issuesAffectedCount,omitempty"`
	IssueCountWithCustomFieldsShowingVersion int                                        `json:"issueCountWithCustomFieldsShowingVersion,omitempty"`
	CustomFieldUsage                         []*VersionIssueCountCustomFieldUsageScheme `json:"customFieldUsage,omitempty"`
}

type VersionIssueCountCustomFieldUsageScheme struct {
	FieldName                          string `json:"fieldName,omitempty"`
	CustomFieldID                      int    `json:"customFieldId,omitempty"`
	IssueCountWithVersionInCustomField int    `json:"issueCountWithVersionInCustomField,omitempty"`
}

type VersionUnresolvedIssuesCountScheme struct {
	Self                  string `json:"self"`
	IssuesUnresolvedCount int    `json:"issuesUnresolvedCount"`
	IssuesCount           int    `json:"issuesCount"`
}
