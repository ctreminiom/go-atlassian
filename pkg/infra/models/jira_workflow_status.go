package models

type WorkflowStatusDetailPageScheme struct {
	StartAt    int                           `json:"startAt,omitempty"`
	Total      int                           `json:"total,omitempty"`
	IsLast     bool                          `json:"isLast,omitempty"`
	MaxResults int                           `json:"maxResults,omitempty"`
	Values     []*WorkflowStatusDetailScheme `json:"values,omitempty"`
	Self       string                        `json:"self,omitempty"`
	NextPage   string                        `json:"nextPage,omitempty"`
}

type WorkflowStatusDetailScheme struct {
	ID              string                     `json:"id,omitempty"`
	Name            string                     `json:"name,omitempty"`
	StatusCategory  string                     `json:"statusCategory,omitempty"`
	StatusReference string                     `json:"statusReference,omitempty"`
	Scope           *WorkflowStatusScopeScheme `json:"scope,omitempty"`
	Description     string                     `json:"description,omitempty"`
	Usages          []*ProjectIssueTypesScheme `json:"usages,omitempty"`
}

type WorkflowStatusScopeScheme struct {
	Type    string                       `json:"type,omitempty"`
	Project *WorkflowStatusProjectScheme `json:"project,omitempty"`
}

type WorkflowStatusProjectScheme struct {
	ID string `json:"id,omitempty"`
}

type ProjectIssueTypesScheme struct {
	Project    *ProjectScheme `json:"project,omitempty"`
	IssueTypes []string       `json:"issueTypes,omitempty"`
}

type WorkflowStatusPayloadScheme struct {
	Statuses []*WorkflowStatusNodeScheme `json:"statuses,omitempty"`
	Scope    *WorkflowStatusScopeScheme  `json:"scope,omitempty"`
}

type WorkflowStatusNodeScheme struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty"`
	StatusCategory string `json:"statusCategory,omitempty"`
	Description    string `json:"description,omitempty"`
}

type WorkflowStatusSearchParams struct {
	ProjectID      string
	SearchString   string
	StatusCategory string
	Expand         []string
}

type StatusDetailScheme struct {
	Self           string                     `json:"self,omitempty"`
	Description    string                     `json:"description,omitempty"`
	IconURL        string                     `json:"iconUrl,omitempty"`
	Name           string                     `json:"name,omitempty"`
	ID             string                     `json:"id,omitempty"`
	StatusCategory *StatusCategoryScheme      `json:"statusCategory,omitempty"`
	Scope          *WorkflowStatusScopeScheme `json:"scope,omitempty"`
}
