package models

// ScreenParamsScheme represents the parameters for a screen in Jira.
type ScreenParamsScheme struct {

	// IDs is the list of screen IDs
	IDs []int

	// QueryString is used to perform a case-insensitive partial match with screen name.
	QueryString string

	// Scope is used to filter by multiple scopes.
	Scope []string

	// OrderBy is used to order the results by a field:
	// 1. id Sorts by screen ID.
	// 2. name Sorts by screen name.
	OrderBy string
}

// ScreenScheme represents a screen in Jira.
type ScreenScheme struct {
	ID          int                            `json:"id,omitempty"`          // The ID of the screen.
	Name        string                         `json:"name,omitempty"`        // The name of the screen.
	Description string                         `json:"description,omitempty"` // The description of the screen.
	Scope       *TeamManagedProjectScopeScheme `json:"scope,omitempty"`       // The scope of the screen.
}

// ScreenFieldPageScheme represents a page of screen fields in Jira.
type ScreenFieldPageScheme struct {
	Self       string                 `json:"self,omitempty"`       // The URL of the screen field page.
	NextPage   string                 `json:"nextPage,omitempty"`   // The URL of the next page.
	MaxResults int                    `json:"maxResults,omitempty"` // The maximum number of results per page.
	StartAt    int                    `json:"startAt,omitempty"`    // The index of the first item returned in the page.
	Total      int                    `json:"total,omitempty"`      // The total number of items available.
	IsLast     bool                   `json:"isLast,omitempty"`     // Indicates if this is the last page.
	Values     []*ScreenWithTabScheme `json:"values,omitempty"`     // The screen fields in the page.
}

// ScreenWithTabScheme represents a screen with a tab in Jira.
type ScreenWithTabScheme struct {
	ID          int                            `json:"id,omitempty"`          // The ID of the screen.
	Name        string                         `json:"name,omitempty"`        // The name of the screen.
	Description string                         `json:"description,omitempty"` // The description of the screen.
	Scope       *TeamManagedProjectScopeScheme `json:"scope,omitempty"`       // The scope of the screen.
	Tab         *ScreenTabScheme               `json:"tab,omitempty"`         // The tab of the screen.
}

// ScreenSearchPageScheme represents a page of screens in a search in Jira.
type ScreenSearchPageScheme struct {
	Self       string          `json:"self,omitempty"`       // The URL of the screen search page.
	MaxResults int             `json:"maxResults,omitempty"` // The maximum number of results per page.
	StartAt    int             `json:"startAt,omitempty"`    // The index of the first item returned in the page.
	Total      int             `json:"total,omitempty"`      // The total number of items available.
	IsLast     bool            `json:"isLast,omitempty"`     // Indicates if this is the last page.
	Values     []*ScreenScheme `json:"values,omitempty"`     // The screens in the page.
}

// AvailableScreenFieldScheme represents an available field in a screen in Jira.
type AvailableScreenFieldScheme struct {
	ID   string `json:"id"`   // The ID of the field.
	Name string `json:"name"` // The name of the field.
}

// ScreenTabScheme represents a tab in a screen in Jira.
type ScreenTabScheme struct {
	ID   int    `json:"id"`   // The ID of the tab.
	Name string `json:"name"` // The name of the tab.
}

// ScreenTabFieldScheme represents a field in a tab in a screen in Jira.
type ScreenTabFieldScheme struct {
	ID   string `json:"id,omitempty"`   // The ID of the field.
	Name string `json:"name,omitempty"` // The name of the field.
}
