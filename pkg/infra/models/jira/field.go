package jira

type IssueFieldScheme struct {
	ID            string                         `json:"id,omitempty"`
	Key           string                         `json:"key,omitempty"`
	Name          string                         `json:"name,omitempty"`
	Custom        bool                           `json:"custom,omitempty"`
	Orderable     bool                           `json:"orderable,omitempty"`
	Navigable     bool                           `json:"navigable,omitempty"`
	Searchable    bool                           `json:"searchable,omitempty"`
	ClauseNames   []string                       `json:"clauseNames,omitempty"`
	Scope         *TeamManagedProjectScopeScheme `json:"scope,omitempty"`
	Schema        *IssueFieldSchemaScheme        `json:"schema,omitempty"`
	Description   string                         `json:"description,omitempty"`
	IsLocked      bool                           `json:"isLocked,omitempty"`
	SearcherKey   string                         `json:"searcherKey,omitempty"`
	ScreensCount  int                            `json:"screensCount,omitempty"`
	ContextsCount int                            `json:"contextsCount,omitempty"`
	LastUsed      *IssueFieldLastUsedScheme      `json:"lastUsed,omitempty"`
}

type IssueFieldSchemaScheme struct {
	Type     string `json:"type,omitempty"`
	Items    string `json:"items,omitempty"`
	System   string `json:"system,omitempty"`
	Custom   string `json:"custom,omitempty"`
	CustomID int    `json:"customId,omitempty"`
}

type IssueFieldLastUsedScheme struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type FieldSearchPageScheme struct {
	MaxResults int                 `json:"maxResults,omitempty"`
	StartAt    int                 `json:"startAt,omitempty"`
	Total      int                 `json:"total,omitempty"`
	IsLast     bool                `json:"isLast,omitempty"`
	Values     []*IssueFieldScheme `json:"values,omitempty"`
}

type FieldSearchOptionsScheme struct {
	Types   []string
	IDs     []string
	Query   string
	OrderBy string
	Expand  []string
}

type CustomFieldScheme struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	FieldType   string `json:"type,omitempty"`
	SearcherKey string `json:"searcherKey,omitempty"`
}
