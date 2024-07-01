package models

// ScreenSchemeParamsScheme represents the parameters for a screen scheme in Jira.
type ScreenSchemeParamsScheme struct {
	IDs         []int    // The IDs of the screen schemes.
	QueryString string   // The query string for the screen scheme search.
	OrderBy     string   // The order by field for the screen scheme search.
	Expand      []string // The fields to be expanded in the screen scheme.
}

// IssueTypeScreenSchemePageScheme represents a page of issue type screen schemes in Jira.
type IssueTypeScreenSchemePageScheme struct {
	Self       string                         `json:"self,omitempty"`       // The URL of the page.
	NextPage   string                         `json:"nextPage,omitempty"`   // The URL of the next page.
	MaxResults int                            `json:"maxResults,omitempty"` // The maximum results per page.
	StartAt    int                            `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                            `json:"total,omitempty"`      // The total number of issue type screen schemes.
	IsLast     bool                           `json:"isLast,omitempty"`     // Indicates if this is the last page.
	Values     []*IssueTypeScreenSchemeScheme `json:"values,omitempty"`     // The issue type screen schemes in the page.
}

// IssueTypeScreenSchemePayloadScheme represents the payload for an issue type screen scheme in Jira.
type IssueTypeScreenSchemePayloadScheme struct {
	Name              string                                       `json:"name,omitempty"`              // The name of the issue type screen scheme.
	Description       string                                       `json:"description,omitempty"`       // The description of the issue type screen scheme.
	IssueTypeMappings []*IssueTypeScreenSchemeMappingPayloadScheme `json:"issueTypeMappings,omitempty"` // The issue type mappings for the screen scheme.
}

// IssueTypeScreenSchemeMappingPayloadScheme represents a mapping payload for an issue type screen scheme in Jira.
type IssueTypeScreenSchemeMappingPayloadScheme struct {
	IssueTypeID    string `json:"issueTypeId,omitempty"`    // The ID of the issue type.
	ScreenSchemeID string `json:"screenSchemeId,omitempty"` // The ID of the screen scheme.
}

// IssueTypeScreenSchemeScheme represents an issue type screen scheme in Jira.
type IssueTypeScreenSchemeScheme struct {
	ID          string               `json:"id,omitempty"`          // The ID of the issue type screen scheme.
	Name        string               `json:"name,omitempty"`        // The name of the issue type screen scheme.
	Description string               `json:"description,omitempty"` // The description of the issue type screen scheme.
	Projects    *ProjectSearchScheme `json:"projects,omitempty"`    // The projects associated with the screen scheme.
}

// IssueTypeScreenScreenCreatedScheme represents a newly created issue type screen scheme in Jira.
type IssueTypeScreenScreenCreatedScheme struct {
	ID string `json:"id"` // The ID of the newly created issue type screen scheme.
}

// IssueTypeProjectScreenSchemePageScheme represents a page of project issue type screen schemes in Jira.
type IssueTypeProjectScreenSchemePageScheme struct {
	Self       string                                 `json:"self,omitempty"`       // The URL of the page.
	NextPage   string                                 `json:"nextPage,omitempty"`   // The URL of the next page.
	MaxResults int                                    `json:"maxResults,omitempty"` // The maximum results per page.
	StartAt    int                                    `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                                    `json:"total,omitempty"`      // The total number of project issue type screen schemes.
	IsLast     bool                                   `json:"isLast,omitempty"`     // Indicates if this is the last page.
	Values     []*IssueTypeScreenSchemesProjectScheme `json:"values,omitempty"`     // The project issue type screen schemes in the page.
}

// IssueTypeScreenSchemesProjectScheme represents the project issue type screen schemes in Jira.
type IssueTypeScreenSchemesProjectScheme struct {
	IssueTypeScreenScheme *IssueTypeScreenSchemeScheme `json:"issueTypeScreenScheme,omitempty"` // The issue type screen scheme.
	ProjectIDs            []string                     `json:"projectIds,omitempty"`            // The IDs of the projects.
}

// IssueTypeScreenSchemeMappingScheme represents a mapping of an issue type screen scheme in Jira.
type IssueTypeScreenSchemeMappingScheme struct {
	Self       string                             `json:"self,omitempty"`       // The URL of the mapping.
	NextPage   string                             `json:"nextPage,omitempty"`   // The URL of the next page.
	MaxResults int                                `json:"maxResults,omitempty"` // The maximum results per page.
	StartAt    int                                `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                                `json:"total,omitempty"`      // The total number of issue type screen scheme mappings.
	IsLast     bool                               `json:"isLast,omitempty"`     // Indicates if this is the last page.
	Values     []*IssueTypeScreenSchemeItemScheme `json:"values,omitempty"`     // The issue type screen scheme items in the mapping.
}

// IssueTypeScreenSchemeItemScheme represents an item of an issue type screen scheme in Jira.
type IssueTypeScreenSchemeItemScheme struct {
	IssueTypeScreenSchemeID string `json:"issueTypeScreenSchemeId,omitempty"` // The ID of the issue type screen scheme.
	IssueTypeID             string `json:"issueTypeId,omitempty"`             // The ID of the issue type.
	ScreenSchemeID          string `json:"screenSchemeId,omitempty"`          // The ID of the screen scheme.
}

// IssueTypeScreenSchemeByProjectPageScheme represents a page of issue type screen schemes by project in Jira.
type IssueTypeScreenSchemeByProjectPageScheme struct {
	MaxResults int                    `json:"maxResults,omitempty"` // The maximum results per page.
	StartAt    int                    `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                    `json:"total,omitempty"`      // The total number of issue type screen schemes by project.
	IsLast     bool                   `json:"isLast,omitempty"`     // Indicates if this is the last page.
	Values     []*ProjectDetailScheme `json:"values,omitempty"`     // The project details in the page.
}

// ProjectDetailScheme represents the details of a project in Jira.
type ProjectDetailScheme struct {
	Self            string                 `json:"self,omitempty"`            // The URL of the project.
	ID              string                 `json:"id,omitempty"`              // The ID of the project.
	Key             string                 `json:"key,omitempty"`             // The key of the project.
	Name            string                 `json:"name,omitempty"`            // The name of the project.
	ProjectTypeKey  string                 `json:"projectTypeKey,omitempty"`  // The type key of the project.
	Simplified      bool                   `json:"simplified,omitempty"`      // Indicates if the project is simplified.
	ProjectCategory *ProjectCategoryScheme `json:"projectCategory,omitempty"` // The category of the project.
}
