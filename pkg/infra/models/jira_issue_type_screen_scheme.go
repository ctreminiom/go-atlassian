package models

type ScreenSchemeParamsScheme struct {
	IDs         []int
	QueryString string
	OrderBy     string
	Expand      []string
}

type IssueTypeScreenSchemePageScheme struct {
	Self       string                         `json:"self,omitempty"`
	NextPage   string                         `json:"nextPage,omitempty"`
	MaxResults int                            `json:"maxResults,omitempty"`
	StartAt    int                            `json:"startAt,omitempty"`
	Total      int                            `json:"total,omitempty"`
	IsLast     bool                           `json:"isLast,omitempty"`
	Values     []*IssueTypeScreenSchemeScheme `json:"values,omitempty"`
}

type IssueTypeScreenSchemePayloadScheme struct {
	Name              string                                       `json:"name,omitempty"`
	IssueTypeMappings []*IssueTypeScreenSchemeMappingPayloadScheme `json:"issueTypeMappings,omitempty"`
}

type IssueTypeScreenSchemeMappingPayloadScheme struct {
	IssueTypeID    string `json:"issueTypeId,omitempty"`
	ScreenSchemeID string `json:"screenSchemeId,omitempty"`
}

type IssueTypeScreenSchemeScheme struct {
	ID          string               `json:"id,omitempty"`
	Name        string               `json:"name,omitempty"`
	Description string               `json:"description,omitempty"`
	Projects    *ProjectSearchScheme `json:"projects,omitempty"`
}

type IssueTypeScreenScreenCreatedScheme struct {
	ID string `json:"id"`
}

type IssueTypeProjectScreenSchemePageScheme struct {
	Self       string                                 `json:"self,omitempty"`
	NextPage   string                                 `json:"nextPage,omitempty"`
	MaxResults int                                    `json:"maxResults,omitempty"`
	StartAt    int                                    `json:"startAt,omitempty"`
	Total      int                                    `json:"total,omitempty"`
	IsLast     bool                                   `json:"isLast,omitempty"`
	Values     []*IssueTypeScreenSchemesProjectScheme `json:"values,omitempty"`
}

type IssueTypeScreenSchemesProjectScheme struct {
	IssueTypeScreenScheme *IssueTypeScreenSchemeScheme `json:"issueTypeScreenScheme,omitempty"`
	ProjectIds            []string                     `json:"projectIds,omitempty"`
}

type IssueTypeScreenSchemeMappingScheme struct {
	Self       string                             `json:"self,omitempty"`
	NextPage   string                             `json:"nextPage,omitempty"`
	MaxResults int                                `json:"maxResults,omitempty"`
	StartAt    int                                `json:"startAt,omitempty"`
	Total      int                                `json:"total,omitempty"`
	IsLast     bool                               `json:"isLast,omitempty"`
	Values     []*IssueTypeScreenSchemeItemScheme `json:"values,omitempty"`
}

type IssueTypeScreenSchemeItemScheme struct {
	IssueTypeScreenSchemeID string `json:"issueTypeScreenSchemeId,omitempty"`
	IssueTypeID             string `json:"issueTypeId,omitempty"`
	ScreenSchemeID          string `json:"screenSchemeId,omitempty"`
}

type IssueTypeScreenSchemeByProjectPageScheme struct {
	MaxResults int                    `json:"maxResults,omitempty"`
	StartAt    int                    `json:"startAt,omitempty"`
	Total      int                    `json:"total,omitempty"`
	IsLast     bool                   `json:"isLast,omitempty"`
	Values     []*ProjectDetailScheme `json:"values,omitempty"`
}

type ProjectDetailScheme struct {
	Self            string                 `json:"self,omitempty"`
	ID              string                 `json:"id,omitempty"`
	Key             string                 `json:"key,omitempty"`
	Name            string                 `json:"name,omitempty"`
	ProjectTypeKey  string                 `json:"projectTypeKey,omitempty"`
	Simplified      bool                   `json:"simplified,omitempty"`
	ProjectCategory *ProjectCategoryScheme `json:"projectCategory,omitempty"`
}
