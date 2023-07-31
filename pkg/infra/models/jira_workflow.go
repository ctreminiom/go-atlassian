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
