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

type WorkflowBulkOptionsScheme struct {
	ProjectAndIssueTypes []*ProjectAndIssueTypePairScheme `json:"projectAndIssueTypes,omitempty"`
	WorkflowIds          []string                         `json:"workflowIds,omitempty"`
	WorkflowNames        []string                         `json:"workflowNames,omitempty"`
}

type ProjectAndIssueTypePairScheme struct {
	IssueTypeID string `json:"issueTypeId,omitempty"`
	ProjectID   string `json:"projectId,omitempty"`
}

type WorkflowReadResponseScheme struct {
	Statuses  []*WorkflowStatusDetailScheme `json:"statuses,omitempty"`
	Workflows []*JiraWorkflowScheme         `json:"workflows,omitempty"`
}

type JiraWorkflowScheme struct {
	Description      string                           `json:"description,omitempty"`
	ID               string                           `json:"id,omitempty"`
	IsEditable       bool                             `json:"isEditable,omitempty"`
	Name             string                           `json:"name,omitempty"`
	Scope            *WorkflowStatusScopeScheme       `json:"scope,omitempty"`
	StartPointLayout *StartPointLayoutScheme          `json:"startPointLayout,omitempty"`
	Statuses         []*WorkflowReferenceStatusScheme `json:"statuses,omitempty"`
	TaskID           string                           `json:"taskId,omitempty"`
	Transitions      []*WorkflowTransitionsScheme     `json:"transitions,omitempty"`
	Usages           []*ProjectIssueTypesScheme       `json:"usages,omitempty"`
	Version          *WorkflowDocumentVersionScheme   `json:"version,omitempty"`
}

type StartPointLayoutScheme struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
}

type WorkflowReferenceStatusScheme struct {
	Deprecated      bool                    `json:"deprecated,omitempty"`
	Layout          *StartPointLayoutScheme `json:"layout,omitempty"`
	StatusReference string                  `json:"statusReference,omitempty"`
}

type WorkflowTransitionsScheme struct {
	Actions            []*WorkflowRuleConfigurationScheme `json:"actions,omitempty"`
	Conditions         *ConditionGroupConfigurationScheme `json:"conditions,omitempty"`
	CustomIssueEventID string                             `json:"customIssueEventId,omitempty"`
	Description        string                             `json:"description,omitempty"`
	From               []*WorkflowStatusAndPortScheme     `json:"from,omitempty"`
	ID                 string                             `json:"id,omitempty"`
	Name               string                             `json:"name,omitempty"`
	To                 *WorkflowStatusAndPortScheme       `json:"to,omitempty"`
	TransitionScreen   *WorkflowRuleConfigurationScheme   `json:"transitionScreen,omitempty"`
	Triggers           []*WorkflowTriggerScheme           `json:"triggers,omitempty"`
	Type               string                             `json:"type,omitempty"`
	Validators         []*WorkflowRuleConfigurationScheme `json:"validators,omitempty"`
}

type WorkflowRuleConfigurationScheme struct {
	ID      string `json:"id,omitempty"`
	RuleKey string `json:"ruleKey,omitempty"`
}

type ConditionGroupConfigurationScheme struct {
	ConditionGroups []*ConditionGroupConfigurationScheme `json:"conditionGroups,omitempty"`
	Conditions      []*WorkflowRuleConfigurationScheme   `json:"conditions,omitempty"`
	Operation       string                               `json:"operation,omitempty"`
}

type WorkflowStatusAndPortScheme struct {
	Port            int    `json:"port,omitempty"`
	StatusReference string `json:"statusReference,omitempty"`
}

type WorkflowTriggerScheme struct {
	ID      string `json:"id,omitempty"`
	RuleKey string `json:"ruleKey,omitempty"`
}

type WorkflowDocumentVersionScheme struct {
	ID            string `json:"id,omitempty"`
	VersionNumber int    `json:"versionNumber,omitempty"`
}

type WorkflowCapabilitiesScheme struct {
	ConnectRules []*AvailableWorkflowConnectRuleScheme `json:"connectRules,omitempty"`
	EditorScope  string                                `json:"editorScope,omitempty"`
	ForgeRules   []*AvailableWorkflowForgeRuleScheme   `json:"forgeRules,omitempty"`
	ProjectTypes []string                              `json:"projectTypes,omitempty"`
	SystemRules  []*AvailableWorkflowSystemRuleScheme  `json:"systemRules,omitempty"`
	TriggerRules []*AvailableWorkflowTriggersScheme    `json:"triggerRules,omitempty"`
}

type AvailableWorkflowConnectRuleScheme struct {
	AddonKey    string `json:"addonKey,omitempty"`
	CreateURL   string `json:"createUrl,omitempty"`
	Description string `json:"description,omitempty"`
	EditURL     string `json:"editUrl,omitempty"`
	ModuleKey   string `json:"moduleKey,omitempty"`
	Name        string `json:"name,omitempty"`
	RuleKey     string `json:"ruleKey,omitempty"`
	RuleType    string `json:"ruleType,omitempty"`
	ViewURL     string `json:"viewUrl,omitempty"`
}

type AvailableWorkflowForgeRuleScheme struct {
	Description string `json:"description,omitempty"`
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	RuleKey     string `json:"ruleKey,omitempty"`
	RuleType    string `json:"ruleType,omitempty"`
}

type AvailableWorkflowSystemRuleScheme struct {
	Description                     string   `json:"description,omitempty"`
	IncompatibleRuleKeys            []string `json:"incompatibleRuleKeys,omitempty"`
	IsAvailableForInitialTransition bool     `json:"isAvailableForInitialTransition,omitempty"`
	IsVisible                       bool     `json:"isVisible,omitempty"`
	Name                            string   `json:"name,omitempty"`
	RuleKey                         string   `json:"ruleKey,omitempty"`
	RuleType                        string   `json:"ruleType,omitempty"`
}

type AvailableWorkflowTriggersScheme struct {
	AvailableTypes []*AvailableWorkflowTriggerTypesScheme `json:"availableTypes,omitempty"`
	RuleKey        string                                 `json:"ruleKey,omitempty"`
}

type AvailableWorkflowTriggerTypesScheme struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
}

type WorkflowCreatesPayloadScheme struct {
	Scope     *WorkflowStatusScopeScheme     `json:"scope,omitempty"`
	Statuses  []*WorkflowStatusUpdateScheme  `json:"statuses,omitempty"`
	Workflows []*WorkflowCreatePayloadScheme `json:"workflows,omitempty"`
}

// WorkflowStatusUpdateScheme contains the details of the status being updated.
type WorkflowStatusUpdateScheme struct {
	Description     string `json:"description,omitempty"`
	ID              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	StatusCategory  string `json:"statusCategory,omitempty"`
	StatusReference string `json:"statusReference,omitempty"`
}

type WorkflowCreatePayloadScheme struct {
	Description      string                      `json:"description,omitempty"`
	Name             string                      `json:"name,omitempty"`
	StartPointLayout *StartPointLayoutScheme     `json:"startPointLayout,omitempty"`
	Statuses         []*StatusLayoutUpdateScheme `json:"statuses,omitempty"`
	Transitions      []*TransitionUpdateScheme   `json:"transitions"`
}

type StatusLayoutUpdateScheme struct {
	Layout          *StartPointLayoutScheme `json:"layout,omitempty"`
	StatusReference string                  `json:"statusReference,omitempty"`
}

type TransitionUpdateScheme struct {
	Actions            []*WorkflowRuleConfigurationScheme `json:"actions,omitempty"`
	Conditions         *ConditionGroupUpdateScheme        `json:"conditions,omitempty"`
	CustomIssueEventID string                             `json:"customIssueEventId,omitempty"`
	Description        string                             `json:"description,omitempty"`
	From               []*StatusReferenceAndPortScheme    `json:"from,omitempty"`
	ID                 string                             `json:"id,omitempty"`
	Name               string                             `json:"name,omitempty"`
	To                 *StatusReferenceAndPortScheme      `json:"to,omitempty"`
	TransitionScreen   *WorkflowRuleConfigurationScheme   `json:"transitionScreen,omitempty"`
	Triggers           []*WorkflowTriggerScheme           `json:"triggers,omitempty"`
	Type               string                             `json:"type,omitempty"`
	Validators         []*WorkflowRuleConfigurationScheme `json:"validators,omitempty"`
}

type ConditionGroupUpdateScheme struct {
	ConditionGroups []*ConditionGroupUpdateScheme      `json:"conditionGroups,omitempty"`
	Conditions      []*WorkflowRuleConfigurationScheme `json:"conditions,omitempty"`
	Operation       string                             `json:"operation,omitempty"`
}

type StatusReferenceAndPortScheme struct {
	Port            int    `json:"port,omitempty"`
	StatusReference string `json:"statusReference,omitempty"`
}

type WorkflowCreateResponseScheme struct {
	Statuses  []*WorkflowStatusDetailScheme `json:"statuses,omitempty"`
	Workflows []*JiraWorkflowScheme         `json:"workflows,omitempty"`
}

type WorkflowUpdatesPayloadScheme struct {
	Statuses  []*WorkflowStatusUpdateScheme  `json:"statuses,omitempty"`
	Workflows []*WorkflowUpdatePayloadScheme `json:"workflows,omitempty"`
}

type WorkflowUpdatePayloadScheme struct {
	DefaultStatusMappings []*StatusMigrationScheme       `json:"defaultStatusMappings,omitempty"`
	Description           string                         `json:"description,omitempty"`
	ID                    string                         `json:"id,omitempty"`
	StartPointLayout      *StartPointLayoutScheme        `json:"startPointLayout,omitempty"`
	StatusMappings        []*StatusMappingScheme         `json:"statusMappings,omitempty"`
	Statuses              []*StatusLayoutUpdateScheme    `json:"statuses,omitempty"`
	Transitions           []*TransitionUpdateScheme      `json:"transitions,omitempty"`
	Version               *WorkflowDocumentVersionScheme `json:"version,omitempty"`
}

type StatusMigrationScheme struct {
	NewStatusReference string `json:"newStatusReference,omitempty"`
	OldStatusReference string `json:"oldStatusReference,omitempty"`
}

type StatusMappingScheme struct {
	IssueTypeID      string                   `json:"issueTypeId,omitempty"`
	ProjectID        string                   `json:"projectId,omitempty"`
	StatusMigrations []*StatusMigrationScheme `json:"statusMigrations,omitempty"`
}

type WorkflowUpdateResponseScheme struct {
	Statuses  []*WorkflowStatusDetailScheme `json:"statuses,omitempty"`
	TaskID    string                        `json:"taskId,omitempty"`
	Workflows []*JiraWorkflowScheme         `json:"workflows,omitempty"`
}

type WorkflowCreateValidatorPayloadScheme struct {
	Payload           *WorkflowCreatesPayloadScheme     `json:"payload,omitempty"`
	ValidationOptions *ValidationOptionsForCreateScheme `json:"validationOptions,omitempty"`
}

type WorkflowUpdateValidatorPayloadScheme struct {
	Payload           *WorkflowUpdatesPayloadScheme     `json:"payload,omitempty"`
	ValidationOptions *ValidationOptionsForCreateScheme `json:"validationOptions,omitempty"`
}

type ValidationOptionsForCreateScheme struct {
	Levels []string `json:"levels,omitempty"`
}

type WorkflowValidationErrorListScheme struct {
	Errors []*WorkflowValidationErrorScheme `json:"errors,omitempty"`
}

type WorkflowValidationErrorScheme struct {
	Code             string                          `json:"code,omitempty"`
	ElementReference *WorkflowElementReferenceScheme `json:"elementReference,omitempty"`
	Level            string                          `json:"level,omitempty"`
	Message          string                          `json:"message,omitempty"`
	Type             string                          `json:"type,omitempty"`
}

type WorkflowElementReferenceScheme struct {
	PropertyKey            string                         `json:"propertyKey,omitempty"`
	RuleID                 string                         `json:"ruleId,omitempty"`
	StatusMappingReference *ProjectAndIssueTypePairScheme `json:"statusMappingReference,omitempty"`
	StatusReference        string                         `json:"statusReference,omitempty"`
	TransitionID           string                         `json:"transitionId,omitempty"`
}
