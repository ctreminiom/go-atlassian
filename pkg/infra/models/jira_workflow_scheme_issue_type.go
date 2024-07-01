package models

// IssueTypeWorkflowMappingScheme represents the mapping between an issue type and a workflow in Jira.
type IssueTypeWorkflowMappingScheme struct {
	IssueType string `json:"issueType,omitempty"` // The type of the issue.
	Workflow  string `json:"workflow,omitempty"`  // The workflow associated with the issue type.
}

// IssueTypeWorkflowPayloadScheme represents the payload for an issue type and workflow mapping in Jira.
type IssueTypeWorkflowPayloadScheme struct {
	IssueType           string `json:"issueType,omitempty"`           // The type of the issue.
	UpdateDraftIfNeeded bool   `json:"updateDraftIfNeeded,omitempty"` // Indicates if the draft should be updated if needed.
	Workflow            string `json:"workflow,omitempty"`            // The workflow associated with the issue type.
}

// IssueTypesWorkflowMappingScheme represents the mapping between multiple issue types and a workflow in Jira.
type IssueTypesWorkflowMappingScheme struct {
	Workflow       string   `json:"workflow,omitempty"`       // The workflow associated with the issue types.
	IssueTypes     []string `json:"issueTypes,omitempty"`     // The types of the issues.
	DefaultMapping bool     `json:"defaultMapping,omitempty"` // Indicates if this is the default mapping.
}
