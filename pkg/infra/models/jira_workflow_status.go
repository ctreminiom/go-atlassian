package models

// WorkflowStatusDetailPageScheme represents a page of workflow status details in Jira.
type WorkflowStatusDetailPageScheme struct {
	StartAt    int                           `json:"startAt,omitempty"`    // The index of the first result returned.
	Total      int                           `json:"total,omitempty"`      // The total number of results available.
	IsLast     bool                          `json:"isLast,omitempty"`     // Indicates if this is the last page of results.
	MaxResults int                           `json:"maxResults,omitempty"` // The maximum number of results returned.
	Values     []*WorkflowStatusDetailScheme `json:"values,omitempty"`     // The workflow status details on the page.
	Self       string                        `json:"self,omitempty"`       // The URL of the page.
	NextPage   string                        `json:"nextPage,omitempty"`   // The URL of the next page.
}

// WorkflowStatusDetailScheme represents a workflow status detail in Jira.
type WorkflowStatusDetailScheme struct {
	ID             string                     `json:"id,omitempty"`             // The ID of the workflow status.
	Name           string                     `json:"name,omitempty"`           // The name of the workflow status.
	StatusCategory string                     `json:"statusCategory,omitempty"` // The status category of the workflow status.
	Scope          *WorkflowStatusScopeScheme `json:"scope,omitempty"`          // The scope of the workflow status.
	Description    string                     `json:"description,omitempty"`    // The description of the workflow status.
	Usages         []*ProjectIssueTypesScheme `json:"usages,omitempty"`         // The usages of the workflow status.
}

// WorkflowStatusScopeScheme represents the scope of a workflow status in Jira.
type WorkflowStatusScopeScheme struct {
	Type    string                       `json:"type,omitempty"`    // The type of the scope.
	Project *WorkflowStatusProjectScheme `json:"project,omitempty"` // The project associated with the scope.
}

// WorkflowStatusProjectScheme represents a project associated with a workflow status scope in Jira.
type WorkflowStatusProjectScheme struct {
	ID string `json:"id,omitempty"` // The ID of the project.
}

// ProjectIssueTypesScheme represents the issue types associated with a project in Jira.
type ProjectIssueTypesScheme struct {
	Project    *ProjectScheme `json:"project,omitempty"`    // The project associated with the issue types.
	IssueTypes []string       `json:"issueTypes,omitempty"` // The issue types associated with the project.
}

// WorkflowStatusPayloadScheme represents the payload for a workflow status in Jira.
type WorkflowStatusPayloadScheme struct {
	Statuses []*WorkflowStatusNodeScheme `json:"statuses,omitempty"` // The statuses of the workflow.
	Scope    *WorkflowStatusScopeScheme  `json:"scope,omitempty"`    // The scope of the workflow status.
}

// WorkflowStatusNodeScheme represents a node of a workflow status in Jira.
type WorkflowStatusNodeScheme struct {
	ID             string `json:"id,omitempty"`             // The ID of the workflow status node.
	Name           string `json:"name,omitempty"`           // The name of the workflow status node.
	StatusCategory string `json:"statusCategory,omitempty"` // The status category of the workflow status node.
	Description    string `json:"description,omitempty"`    // The description of the workflow status node.
}

// WorkflowStatusSearchParams represents the search parameters for a workflow status in Jira.
type WorkflowStatusSearchParams struct {
	ProjectID      string   // The ID of the project.
	SearchString   string   // The search string.
	StatusCategory string   // The status category.
	Expand         []string // The fields to expand in the response.
}

// StatusDetailScheme represents a status detail in Jira.
type StatusDetailScheme struct {
	Self           string                     `json:"self,omitempty"`           // The URL of the status detail.
	Description    string                     `json:"description,omitempty"`    // The description of the status detail.
	IconURL        string                     `json:"iconUrl,omitempty"`        // The URL of the icon for the status detail.
	Name           string                     `json:"name,omitempty"`           // The name of the status detail.
	ID             string                     `json:"id,omitempty"`             // The ID of the status detail.
	StatusCategory *StatusCategoryScheme      `json:"statusCategory,omitempty"` // The status category of the status detail.
	Scope          *WorkflowStatusScopeScheme `json:"scope,omitempty"`          // The scope of the status detail.
}
