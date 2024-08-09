package models

// IssueFieldScheme represents an issue field in Jira.
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

// IssueFieldSchemaScheme represents the schema of an issue field in Jira.
type IssueFieldSchemaScheme struct {
	Type     string `json:"type,omitempty"`
	Items    string `json:"items,omitempty"`
	System   string `json:"system,omitempty"`
	Custom   string `json:"custom,omitempty"`
	CustomID int    `json:"customId,omitempty"`
}

// IssueFieldLastUsedScheme represents the last used information of an issue field in Jira.
type IssueFieldLastUsedScheme struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// FieldSearchPageScheme represents a search page of fields in Jira.
type FieldSearchPageScheme struct {
	MaxResults int                 `json:"maxResults,omitempty"`
	StartAt    int                 `json:"startAt,omitempty"`
	Total      int                 `json:"total,omitempty"`
	IsLast     bool                `json:"isLast,omitempty"`
	Values     []*IssueFieldScheme `json:"values,omitempty"`
}

// FieldSearchOptionsScheme represents the search options for a field in Jira.
type FieldSearchOptionsScheme struct {
	Types   []string
	IDs     []string
	Query   string
	OrderBy string
	Expand  []string
}

// CustomFieldScheme represents a custom field in Jira.
type CustomFieldScheme struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	FieldType   string `json:"type,omitempty"`
	SearcherKey string `json:"searcherKey,omitempty"`
}

// CustomFieldAssetScheme represents a custom field asset in Jira.
type CustomFieldAssetScheme struct {
	WorkspaceID string `json:"workspaceId,omitempty"` // The ID of the workspace.
	ID          string `json:"id,omitempty"`          // The ID of the custom field asset.
	ObjectID    string `json:"objectId,omitempty"`    // The object ID of the custom field asset.
}

// CustomFieldRequestTypeScheme represents a custom field request type in Jira.
type CustomFieldRequestTypeScheme struct {
	Links         *CustomFieldRequestTypeLinkScheme   `json:"_links"`        // Links related to the custom field request type.
	RequestType   *CustomerRequestTypeScheme          `json:"requestType"`   // The request type of the custom field.
	CurrentStatus *CustomerRequestCurrentStatusScheme `json:"currentStatus"` // The current status of the request type.
}

// CustomFieldRequestTypeLinkScheme represents the links of a custom field request type in Jira.
type CustomFieldRequestTypeLinkScheme struct {
	Self     string `json:"self,omitempty"`     // The URL of the custom field request type itself.
	JiraRest string `json:"jiraRest,omitempty"` // The Jira REST API link for the custom field request type.
	Web      string `json:"web,omitempty"`      // The web link for the custom field request type.
	Agent    string `json:"agent,omitempty"`    // The agent link for the custom field request type.
}

// CustomFieldTempoAccountScheme represents a custom field tempo account in Jira.
type CustomFieldTempoAccountScheme struct {
	ID    int    `json:"id"`
	Value string `json:"value"`
}
