package models

// WorkflowSearchOptions represents the search options for a workflow in Jira.
type WorkflowSearchOptions struct {
	WorkflowName []string // The names of the workflows to search for.
	Expand       []string // The fields to expand in the response.
	QueryString  string   // The query string for the search.
	OrderBy      string   // The field to order the results by.
	IsActive     bool     // Indicates if only active workflows should be returned.
}

// WorkflowPageScheme represents a page of workflows in Jira.
type WorkflowPageScheme struct {
	Self       string            `json:"self,omitempty"`       // The URL of the page.
	NextPage   string            `json:"nextPage,omitempty"`   // The URL of the next page.
	MaxResults int               `json:"maxResults,omitempty"` // The maximum number of results returned.
	StartAt    int               `json:"startAt,omitempty"`    // The index of the first result returned.
	Total      int               `json:"total,omitempty"`      // The total number of results available.
	IsLast     bool              `json:"isLast,omitempty"`     // Indicates if this is the last page of results.
	Values     []*WorkflowScheme `json:"values,omitempty"`     // The workflows on the page.
}

// WorkflowScheme represents a workflow in Jira.
type WorkflowScheme struct {
	ID          *WorkflowPublishedIDScheme  `json:"id,omitempty"`          // The ID of the workflow.
	Transitions []*WorkflowTransitionScheme `json:"transitions,omitempty"` // The transitions of the workflow.
	Statuses    []*WorkflowStatusScheme     `json:"statuses,omitempty"`    // The statuses of the workflow.
	Description string                      `json:"description,omitempty"` // The description of the workflow.
	IsDefault   bool                        `json:"isDefault,omitempty"`   // Indicates if the workflow is the default workflow.
}

// WorkflowPublishedIDScheme represents the published ID of a workflow in Jira.
type WorkflowPublishedIDScheme struct {
	Name     string `json:"name,omitempty"`     // The name of the workflow.
	EntityID string `json:"entityId,omitempty"` // The entity ID of the workflow.
}

// WorkflowTransitionScheme represents a transition in a workflow in Jira.
type WorkflowTransitionScheme struct {
	ID          string                          `json:"id,omitempty"`          // The ID of the transition.
	Name        string                          `json:"name,omitempty"`        // The name of the transition.
	Description string                          `json:"description,omitempty"` // The description of the transition.
	From        []string                        `json:"from,omitempty"`        // The statuses from which this transition can be executed.
	To          string                          `json:"to,omitempty"`          // The status to which this transition goes.
	Type        string                          `json:"type,omitempty"`        // The type of the transition.
	Screen      *WorkflowTransitionScreenScheme `json:"screen,omitempty"`      // The screen associated with the transition.
	Rules       *WorkflowTransitionRulesScheme  `json:"rules,omitempty"`       // The rules of the transition.
}

// WorkflowTransitionScreenScheme represents a screen associated with a transition in a workflow in Jira.
type WorkflowTransitionScreenScheme struct {
	ID         string      `json:"id,omitempty"`         // The ID of the screen.
	Properties interface{} `json:"properties,omitempty"` // The properties of the screen.
}

// WorkflowTransitionRulesScheme represents the rules of a transition in a workflow in Jira.
type WorkflowTransitionRulesScheme struct {
	Conditions    []*WorkflowTransitionRuleScheme `json:"conditions,omitempty"`    // The conditions of the transition.
	Validators    []*WorkflowTransitionRuleScheme `json:"validators,omitempty"`    // The validators of the transition.
	PostFunctions []*WorkflowTransitionRuleScheme `json:"postFunctions,omitempty"` // The post functions of the transition.
}

// WorkflowTransitionRuleScheme represents a rule of a transition in a workflow in Jira.
type WorkflowTransitionRuleScheme struct {
	Type          string      `json:"type,omitempty"`          // The type of the rule.
	Configuration interface{} `json:"configuration,omitempty"` // The configuration of the rule.
}

// WorkflowStatusScheme represents a status in a workflow in Jira.
type WorkflowStatusScheme struct {
	ID         string                          `json:"id,omitempty"`         // The ID of the status.
	Name       string                          `json:"name,omitempty"`       // The name of the status.
	Properties *WorkflowStatusPropertiesScheme `json:"properties,omitempty"` // The properties of the status.
}

// WorkflowStatusPropertiesScheme represents the properties of a status in a workflow in Jira.
type WorkflowStatusPropertiesScheme struct {
	IssueEditable bool `json:"issueEditable,omitempty"` // Indicates if the issue is editable.
}

// WorkflowCreatedResponseScheme represents the response after a workflow is created in Jira.
type WorkflowCreatedResponseScheme struct {
	Name     string `json:"name,omitempty"`     // The name of the created workflow.
	EntityID string `json:"entityId,omitempty"` // The entity ID of the created workflow.
}

// WorkflowPayloadScheme represents the payload for creating a workflow in Jira.
type WorkflowPayloadScheme struct {
	Name        string                             `json:"name,omitempty"`        // The name of the workflow.
	Description string                             `json:"description,omitempty"` // The description of the workflow.
	Statuses    []*WorkflowTransitionScreenScheme  `json:"statuses,omitempty"`    // The statuses of the workflow.
	Transitions []*WorkflowTransitionPayloadScheme `json:"transitions,omitempty"` // The transitions of the workflow.
}

// WorkflowTransitionPayloadScheme represents the payload for a transition in a workflow in Jira.
type WorkflowTransitionPayloadScheme struct {
	Name        string                                 `json:"name,omitempty"`        // The name of the transition.
	Description string                                 `json:"description,omitempty"` // The description of the transition.
	From        []string                               `json:"from,omitempty"`        // The statuses from which this transition can be executed.
	To          string                                 `json:"to,omitempty"`          // The status to which this transition goes.
	Type        string                                 `json:"type,omitempty"`        // The type of the transition.
	Rules       *WorkflowTransitionRulePayloadScheme   `json:"rules,omitempty"`       // The rules of the transition.
	Screen      *WorkflowTransitionScreenPayloadScheme `json:"screen,omitempty"`      // The screen associated with the transition.
	Properties  interface{}                            `json:"properties,omitempty"`  // The properties of the transition.
}

// WorkflowTransitionScreenPayloadScheme represents the payload for a screen associated with a transition in a workflow in Jira.
type WorkflowTransitionScreenPayloadScheme struct {
	ID string `json:"id"` // The ID of the screen.
}

// WorkflowTransitionRulePayloadScheme represents the payload for the rules of a transition in a workflow in Jira.
type WorkflowTransitionRulePayloadScheme struct {
	Conditions    *WorkflowConditionScheme        `json:"conditions,omitempty"`    // The conditions of the transition.
	PostFunctions []*WorkflowTransitionRuleScheme `json:"postFunctions,omitempty"` // The post functions of the transition.
	Validators    []*WorkflowTransitionRuleScheme `json:"validators,omitempty"`    // The validators of the transition.
}

// WorkflowConditionScheme represents a condition in a workflow in Jira.
type WorkflowConditionScheme struct {
	Conditions    []*WorkflowConditionScheme `json:"conditions,omitempty"`    // The conditions of the workflow.
	Configuration interface{}                `json:"configuration,omitempty"` // The configuration of the condition.
	Operator      string                     `json:"operator,omitempty"`      // The operator of the condition.
	Type          string                     `json:"type,omitempty"`          // The type of the condition.
}
