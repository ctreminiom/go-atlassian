package models

type WorkflowSearchOptions struct {
	WorkflowName []string
	Expand       []string
	QueryString  string
	OrderBy      string
	IsActive     bool
}

type WorkflowPageScheme struct {
	Self       string            `json:"self,omitempty"`
	NextPage   string            `json:"nextPage,omitempty"`
	MaxResults int               `json:"maxResults,omitempty"`
	StartAt    int               `json:"startAt,omitempty"`
	Total      int               `json:"total,omitempty"`
	IsLast     bool              `json:"isLast,omitempty"`
	Values     []*WorkflowScheme `json:"values,omitempty"`
}

type WorkflowScheme struct {
	ID          *WorkflowPublishedIDScheme  `json:"id,omitempty"`
	Transitions []*WorkflowTransitionScheme `json:"transitions,omitempty"`
	Statuses    []*WorkflowStatusScheme     `json:"statuses,omitempty"`
	Description string                      `json:"description,omitempty"`
	IsDefault   bool                        `json:"isDefault,omitempty"`
}

type WorkflowPublishedIDScheme struct {
	Name     string `json:"name,omitempty"`
	EntityID string `json:"entityId,omitempty"`
}

type WorkflowTransitionScheme struct {
	ID          string                          `json:"id,omitempty"`
	Name        string                          `json:"name,omitempty"`
	Description string                          `json:"description,omitempty"`
	From        []string                        `json:"from,omitempty"`
	To          string                          `json:"to,omitempty"`
	Type        string                          `json:"type,omitempty"`
	Screen      *WorkflowTransitionScreenScheme `json:"screen,omitempty"`
	Rules       *WorkflowTransitionRulesScheme  `json:"rules,omitempty"`
}

type WorkflowTransitionScreenScheme struct {
	ID         string      `json:"id,omitempty"`
	Properties interface{} `json:"properties,omitempty"`
}

type WorkflowTransitionRulesScheme struct {
	Conditions    []*WorkflowTransitionRuleScheme `json:"conditions,omitempty"`
	Validators    []*WorkflowTransitionRuleScheme `json:"validators,omitempty"`
	PostFunctions []*WorkflowTransitionRuleScheme `json:"postFunctions,omitempty"`
}

type WorkflowTransitionRuleScheme struct {
	Type          string      `json:"type,omitempty"`
	Configuration interface{} `json:"configuration,omitempty"`
}

type WorkflowStatusScheme struct {
	ID         string                          `json:"id,omitempty"`
	Name       string                          `json:"name,omitempty"`
	Properties *WorkflowStatusPropertiesScheme `json:"properties,omitempty"`
}

type WorkflowStatusPropertiesScheme struct {
	IssueEditable bool `json:"issueEditable,omitempty"`
}

type WorkflowCreatedResponseScheme struct {
	Name     string `json:"name,omitempty"`
	EntityID string `json:"entityId,omitempty"`
}

type WorkflowPayloadScheme struct {
	Name        string                             `json:"name,omitempty"`
	Description string                             `json:"description,omitempty"`
	Statuses    []*WorkflowTransitionScreenScheme  `json:"statuses,omitempty"`
	Transitions []*WorkflowTransitionPayloadScheme `json:"transitions,omitempty"`
}

type WorkflowTransitionPayloadScheme struct {
	Name        string                                 `json:"name,omitempty"`
	Description string                                 `json:"description,omitempty"`
	From        []string                               `json:"from,omitempty"`
	To          string                                 `json:"to,omitempty"`
	Type        string                                 `json:"type,omitempty"`
	Rules       *WorkflowTransitionRulePayloadScheme   `json:"rules,omitempty"`
	Screen      *WorkflowTransitionScreenPayloadScheme `json:"screen,omitempty"`
	Properties  interface{}                            `json:"properties,omitempty"`
}

type WorkflowTransitionScreenPayloadScheme struct {
	ID string `json:"id"`
}

type WorkflowTransitionRulePayloadScheme struct {
	Conditions    *WorkflowConditionScheme        `json:"conditions,omitempty"`
	PostFunctions []*WorkflowTransitionRuleScheme `json:"postFunctions,omitempty"`
	Validators    []*WorkflowTransitionRuleScheme `json:"validators,omitempty"`
}

type WorkflowConditionScheme struct {
	Conditions    []*WorkflowConditionScheme `json:"conditions,omitempty"`
	Configuration interface{}                `json:"configuration,omitempty"`
	Operator      string                     `json:"operator,omitempty"`
	Type          string                     `json:"type,omitempty"`
}
