package models

type WorkflowSchemePayloadScheme struct {
	DefaultWorkflow   string      `json:"defaultWorkflow,omitempty"`
	Name              string      `json:"name,omitempty"`
	Description       string      `json:"description,omitempty"`
	IssueTypeMappings interface{} `json:"issueTypeMappings,omitempty"`
}

type WorkflowSchemePageScheme struct {
	Self       string                  `json:"self,omitempty"`
	NextPage   string                  `json:"nextPage,omitempty"`
	MaxResults int                     `json:"maxResults,omitempty"`
	StartAt    int                     `json:"startAt,omitempty"`
	Total      int                     `json:"total,omitempty"`
	IsLast     bool                    `json:"isLast,omitempty"`
	Values     []*WorkflowSchemeScheme `json:"values,omitempty"`
}

type WorkflowSchemeScheme struct {
	ID                  int         `json:"id,omitempty"`
	Name                string      `json:"name,omitempty"`
	Description         string      `json:"description,omitempty"`
	DefaultWorkflow     string      `json:"defaultWorkflow,omitempty"`
	Draft               bool        `json:"draft,omitempty"`
	LastModifiedUser    *UserScheme `json:"lastModifiedUser,omitempty"`
	LastModified        string      `json:"lastModified,omitempty"`
	Self                string      `json:"self,omitempty"`
	UpdateDraftIfNeeded bool        `json:"updateDraftIfNeeded,omitempty"`
}

type WorkflowSchemeAssociationPageScheme struct {
	Values []*WorkflowSchemeAssociationsScheme `json:"values,omitempty"`
}

type WorkflowSchemeAssociationsScheme struct {
	ProjectIds     []string              `json:"projectIds,omitempty"`
	WorkflowScheme *WorkflowSchemeScheme `json:"workflowScheme,omitempty"`
}
