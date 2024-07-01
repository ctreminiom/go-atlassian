package models

// ScreenSchemePageScheme represents a page of screen schemes in Jira.
type ScreenSchemePageScheme struct {
	Self       string                `json:"self,omitempty"`       // The URL of the screen scheme page.
	NextPage   string                `json:"nextPage,omitempty"`   // The URL of the next page.
	MaxResults int                   `json:"maxResults,omitempty"` // The maximum number of results per page.
	StartAt    int                   `json:"startAt,omitempty"`    // The index of the first item returned in the page.
	Total      int                   `json:"total,omitempty"`      // The total number of items available.
	IsLast     bool                  `json:"isLast,omitempty"`     // Indicates if this is the last page.
	Values     []*ScreenSchemeScheme `json:"values,omitempty"`     // The screen schemes in the page.
}

// ScreenSchemeScheme represents a screen scheme in Jira.
type ScreenSchemeScheme struct {
	ID                     int                        `json:"id,omitempty"`                     // The ID of the screen scheme.
	Name                   string                     `json:"name,omitempty"`                   // The name of the screen scheme.
	Description            string                     `json:"description,omitempty"`            // The description of the screen scheme.
	Screens                *ScreenTypesScheme         `json:"screens,omitempty"`                // The types of screens in the screen scheme.
	IssueTypeScreenSchemes *IssueTypeSchemePageScheme `json:"issueTypeScreenSchemes,omitempty"` // The issue type screen schemes in the screen scheme.
}

// ScreenTypesScheme represents the types of screens in a screen scheme in Jira.
type ScreenTypesScheme struct {
	Create  int `json:"create,omitempty"`  // The ID of the create screen.
	Default int `json:"default,omitempty"` // The ID of the default screen.
	View    int `json:"view,omitempty"`    // The ID of the view screen.
	Edit    int `json:"edit,omitempty"`    // The ID of the edit screen.
}

// IssueTypeSchemePageScheme represents a page of issue type schemes in Jira.
type IssueTypeSchemePageScheme struct {
	Self       string                   `json:"self,omitempty"`       // The URL of the issue type scheme page.
	NextPage   string                   `json:"nextPage,omitempty"`   // The URL of the next page.
	MaxResults int                      `json:"maxResults,omitempty"` // The maximum number of results per page.
	StartAt    int                      `json:"startAt,omitempty"`    // The index of the first item returned in the page.
	Total      int                      `json:"total,omitempty"`      // The total number of items available.
	IsLast     bool                     `json:"isLast,omitempty"`     // Indicates if this is the last page.
	Values     []*IssueTypeSchemeScheme `json:"values,omitempty"`     // The issue type schemes in the page.
}

// IssueTypeSchemeScheme represents an issue type scheme in Jira.
type IssueTypeSchemeScheme struct {
	ID                 string `json:"id,omitempty"`                 // The ID of the issue type scheme.
	Name               string `json:"name,omitempty"`               // The name of the issue type scheme.
	Description        string `json:"description,omitempty"`        // The description of the issue type scheme.
	DefaultIssueTypeID string `json:"defaultIssueTypeId,omitempty"` // The ID of the default issue type in the scheme.
	IsDefault          bool   `json:"isDefault,omitempty"`          // Indicates if this is the default issue type scheme.
}

// ScreenSchemePayloadScheme represents the payload for a screen scheme in Jira.
type ScreenSchemePayloadScheme struct {
	Screens     *ScreenTypesScheme `json:"screens,omitempty"`     // The types of screens in the screen scheme.
	Name        string             `json:"name,omitempty"`        // The name of the screen scheme.
	Description string             `json:"description,omitempty"` // The description of the screen scheme.
}
