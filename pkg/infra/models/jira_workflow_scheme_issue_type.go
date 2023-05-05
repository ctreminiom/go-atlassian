package models

type IssueTypeWorkflowMappingScheme struct {
	IssueType string `json:"issueType,omitempty"`
	Workflow  string `json:"workflow,omitempty"`
}

type IssueTypeWorkflowPayloadScheme struct {
	IssueType           string `json:"issueType,omitempty"`
	UpdateDraftIfNeeded bool   `json:"updateDraftIfNeeded,omitempty"`
	Workflow            string `json:"workflow,omitempty"`
}

type IssueTypesWorkflowMappingScheme struct {
	Workflow       string   `json:"workflow,omitempty"`
	IssueTypes     []string `json:"issueTypes,omitempty"`
	DefaultMapping bool     `json:"defaultMapping,omitempty"`
}
