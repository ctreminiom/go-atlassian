package jira

type ScreenSchemePageScheme struct {
	Self       string                `json:"self,omitempty"`
	NextPage   string                `json:"nextPage,omitempty"`
	MaxResults int                   `json:"maxResults,omitempty"`
	StartAt    int                   `json:"startAt,omitempty"`
	Total      int                   `json:"total,omitempty"`
	IsLast     bool                  `json:"isLast,omitempty"`
	Values     []*ScreenSchemeScheme `json:"values,omitempty"`
}

type ScreenSchemeScheme struct {
	ID                     int                        `json:"id,omitempty"`
	Name                   string                     `json:"name,omitempty"`
	Description            string                     `json:"description,omitempty"`
	Screens                *ScreenTypesScheme         `json:"screens,omitempty"`
	IssueTypeScreenSchemes *IssueTypeSchemePageScheme `json:"issueTypeScreenSchemes,omitempty"`
}

type ScreenTypesScheme struct {
	Create  int `json:"create,omitempty"`
	Default int `json:"default,omitempty"`
	View    int `json:"view,omitempty"`
	Edit    int `json:"edit,omitempty"`
}

type IssueTypeSchemePageScheme struct {
	Self       string                   `json:"self,omitempty"`
	NextPage   string                   `json:"nextPage,omitempty"`
	MaxResults int                      `json:"maxResults,omitempty"`
	StartAt    int                      `json:"startAt,omitempty"`
	Total      int                      `json:"total,omitempty"`
	IsLast     bool                     `json:"isLast,omitempty"`
	Values     []*IssueTypeSchemeScheme `json:"values,omitempty"`
}

type IssueTypeSchemeScheme struct {
	ID                 string `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	DefaultIssueTypeID string `json:"defaultIssueTypeId,omitempty"`
	IsDefault          bool   `json:"isDefault,omitempty"`
}

type ScreenSchemePayloadScheme struct {
	Screens     *ScreenTypesScheme `json:"screens,omitempty"`
	Name        string             `json:"name,omitempty"`
	Description string             `json:"description,omitempty"`
}
