package models

type CustomFieldAssetScheme struct {
	WorkspaceId string `json:"workspaceId,omitempty"`
	Id          string `json:"id,omitempty"`
	ObjectId    string `json:"objectId,omitempty"`
}

type CustomFieldContextOptionPageScheme struct {
	Self       string                            `json:"self,omitempty"`
	NextPage   string                            `json:"nextPage,omitempty"`
	MaxResults int                               `json:"maxResults,omitempty"`
	StartAt    int                               `json:"startAt,omitempty"`
	Total      int                               `json:"total,omitempty"`
	IsLast     bool                              `json:"isLast,omitempty"`
	Values     []*CustomFieldContextOptionScheme `json:"values,omitempty"`
}

type CustomFieldContextOptionScheme struct {
	ID       string `json:"id,omitempty"`
	Value    string `json:"value,omitempty"`
	Disabled bool   `json:"disabled"`
	OptionID string `json:"optionId,omitempty"`
}

type FieldContextOptionListScheme struct {
	Options []*CustomFieldContextOptionScheme `json:"options,omitempty"`
}

type FieldOptionContextParams struct {
	OptionID    int
	OnlyOptions bool
}

type OrderFieldOptionPayloadScheme struct {
	After                string   `json:"after,omitempty"`
	Position             string   `json:"position,omitempty"`
	CustomFieldOptionIds []string `json:"customFieldOptionIds,omitempty"`
}

type CascadingSelectScheme struct {
	Self  string                      `json:"self,omitempty"`
	Value string                      `json:"value,omitempty"`
	Id    string                      `json:"id,omitempty"`
	Child *CascadingSelectChildScheme `json:"child,omitempty"`
}

type CascadingSelectChildScheme struct {
	Self  string `json:"self,omitempty"`
	Value string `json:"value,omitempty"`
	Id    string `json:"id,omitempty"`
}
