package models

// CustomFieldContextOptionPageScheme represents a page of custom field context options in Jira.
type CustomFieldContextOptionPageScheme struct {
	Self       string                            `json:"self,omitempty"`       // The URL of the page.
	NextPage   string                            `json:"nextPage,omitempty"`   // The URL of the next page.
	MaxResults int                               `json:"maxResults,omitempty"` // The maximum number of results in the page.
	StartAt    int                               `json:"startAt,omitempty"`    // The starting index of the page.
	Total      int                               `json:"total,omitempty"`      // The total number of custom field context options.
	IsLast     bool                              `json:"isLast,omitempty"`     // Indicates if the page is the last one.
	Values     []*CustomFieldContextOptionScheme `json:"values,omitempty"`     // The custom field context options in the page.
}

// CustomFieldContextOptionScheme represents a custom field context option in Jira.
type CustomFieldContextOptionScheme struct {
	ID       string `json:"id,omitempty"`       // The ID of the custom field context option.
	Value    string `json:"value,omitempty"`    // The value of the custom field context option.
	Disabled bool   `json:"disabled"`           // Indicates if the custom field context option is disabled.
	OptionID string `json:"optionId,omitempty"` // The ID of the option.
}

// FieldContextOptionListScheme represents a list of field context options in Jira.
type FieldContextOptionListScheme struct {
	Options []*CustomFieldContextOptionScheme `json:"options,omitempty"` // The field context options.
}

// FieldOptionContextParams represents the parameters for a field option context in Jira.
type FieldOptionContextParams struct {
	OptionID    int  // The ID of the option.
	OnlyOptions bool // Indicates if only options are applicable.
}

// OrderFieldOptionPayloadScheme represents the payload for ordering a field option in Jira.
type OrderFieldOptionPayloadScheme struct {
	After                string   `json:"after,omitempty"`                // The ID of the option after which the current option should be placed.
	Position             string   `json:"position,omitempty"`             // The position of the option.
	CustomFieldOptionIDs []string `json:"customFieldOptionIds,omitempty"` // The IDs of the custom field options.
}

// CascadingSelectScheme represents a cascading select in Jira.
type CascadingSelectScheme struct {
	Self  string                      `json:"self,omitempty"`  // The URL of the cascading select.
	Value string                      `json:"value,omitempty"` // The value of the cascading select.
	ID    string                      `json:"id,omitempty"`    // The ID of the cascading select.
	Child *CascadingSelectChildScheme `json:"child,omitempty"` // The child of the cascading select.
}

// CascadingSelectChildScheme represents a child of a cascading select in Jira.
type CascadingSelectChildScheme struct {
	Self  string `json:"self,omitempty"`  // The URL of the child.
	Value string `json:"value,omitempty"` // The value of the child.
	ID    string `json:"id,omitempty"`    // The ID of the child.
}
