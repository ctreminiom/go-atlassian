package models

// WorkflowSchemePayloadScheme represents the payload for a workflow scheme in Jira.
type WorkflowSchemePayloadScheme struct {
	DefaultWorkflow     string      `json:"defaultWorkflow,omitempty"`     // The default workflow of the scheme.
	Name                string      `json:"name,omitempty"`                // The name of the scheme.
	Description         string      `json:"description,omitempty"`         // The description of the scheme.
	IssueTypeMappings   interface{} `json:"issueTypeMappings,omitempty"`   // The issue type mappings of the scheme.
	UpdateDraftIfNeeded bool        `json:"updateDraftIfNeeded,omitempty"` // Indicates if the draft should be updated if needed.
}

// WorkflowSchemePageScheme represents a page of workflow schemes in Jira.
type WorkflowSchemePageScheme struct {
	Self       string                  `json:"self,omitempty"`       // The URL of the page.
	NextPage   string                  `json:"nextPage,omitempty"`   // The URL of the next page.
	MaxResults int                     `json:"maxResults,omitempty"` // The maximum number of results returned.
	StartAt    int                     `json:"startAt,omitempty"`    // The index of the first result returned.
	Total      int                     `json:"total,omitempty"`      // The total number of results available.
	IsLast     bool                    `json:"isLast,omitempty"`     // Indicates if this is the last page of results.
	Values     []*WorkflowSchemeScheme `json:"values,omitempty"`     // The workflow schemes on the page.
}

// WorkflowSchemeScheme represents a workflow scheme in Jira.
type WorkflowSchemeScheme struct {
	ID                  int         `json:"id,omitempty"`                  // The ID of the scheme.
	Name                string      `json:"name,omitempty"`                // The name of the scheme.
	Description         string      `json:"description,omitempty"`         // The description of the scheme.
	DefaultWorkflow     string      `json:"defaultWorkflow,omitempty"`     // The default workflow of the scheme.
	Draft               bool        `json:"draft,omitempty"`               // Indicates if the scheme is a draft.
	LastModifiedUser    *UserScheme `json:"lastModifiedUser,omitempty"`    // The user who last modified the scheme.
	LastModified        string      `json:"lastModified,omitempty"`        // The date and time when the scheme was last modified.
	Self                string      `json:"self,omitempty"`                // The URL of the scheme.
	UpdateDraftIfNeeded bool        `json:"updateDraftIfNeeded,omitempty"` // Indicates if the draft should be updated if needed.
}

// WorkflowSchemeAssociationPageScheme represents a page of workflow scheme associations in Jira.
type WorkflowSchemeAssociationPageScheme struct {
	Values []*WorkflowSchemeAssociationsScheme `json:"values,omitempty"` // The workflow scheme associations on the page.
}

// WorkflowSchemeAssociationsScheme represents a workflow scheme association in Jira.
type WorkflowSchemeAssociationsScheme struct {
	ProjectIDs     []string              `json:"projectIds,omitempty"`     // The IDs of the projects associated with the scheme.
	WorkflowScheme *WorkflowSchemeScheme `json:"workflowScheme,omitempty"` // The workflow scheme associated with the projects.
}
