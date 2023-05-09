package models

type ScreenParamsScheme struct {

	// The list of screen IDs
	IDs []int

	// String used to perform a case-insensitive partial match with screen name.
	QueryString string

	// The scope filter string. To filter by multiple scope,
	Scope []string

	// Order the results by a field:
	// 1. id Sorts by screen ID.
	// 2. name Sorts by screen name.
	OrderBy string
}

type ScreenScheme struct {
	ID          int                            `json:"id,omitempty"`
	Name        string                         `json:"name,omitempty"`
	Description string                         `json:"description,omitempty"`
	Scope       *TeamManagedProjectScopeScheme `json:"scope,omitempty"`
}

type ScreenFieldPageScheme struct {
	Self       string                 `json:"self,omitempty"`
	NextPage   string                 `json:"nextPage,omitempty"`
	MaxResults int                    `json:"maxResults,omitempty"`
	StartAt    int                    `json:"startAt,omitempty"`
	Total      int                    `json:"total,omitempty"`
	IsLast     bool                   `json:"isLast,omitempty"`
	Values     []*ScreenWithTabScheme `json:"values,omitempty"`
}

type ScreenWithTabScheme struct {
	ID          int                            `json:"id,omitempty"`
	Name        string                         `json:"name,omitempty"`
	Description string                         `json:"description,omitempty"`
	Scope       *TeamManagedProjectScopeScheme `json:"scope,omitempty"`
	Tab         *ScreenTabScheme               `json:"tab,omitempty"`
}

type ScreenSearchPageScheme struct {
	Self       string          `json:"self,omitempty"`
	MaxResults int             `json:"maxResults,omitempty"`
	StartAt    int             `json:"startAt,omitempty"`
	Total      int             `json:"total,omitempty"`
	IsLast     bool            `json:"isLast,omitempty"`
	Values     []*ScreenScheme `json:"values,omitempty"`
}

type AvailableScreenFieldScheme struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ScreenTabScheme struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ScreenTabFieldScheme struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
