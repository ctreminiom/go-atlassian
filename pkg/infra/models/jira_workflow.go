package models

import "fmt"

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
// Note that a transition can have either the deprecated to/from fields or the toStatusReference/links fields,
// but never both nor a combination.
type WorkflowTransitionScheme struct {
	Actions            []*WorkflowRuleConfigurationScheme `json:"actions,omitempty"`
	Conditions         *ConditionGroupConfigurationScheme `json:"conditions,omitempty"`
	CustomIssueEventID string                             `json:"customIssueEventId,omitempty"`
	Description        string                             `json:"description,omitempty"`
	From               []*WorkflowStatusAndPortScheme     `json:"from,omitempty"` // This field is deprecated - use toStatusReference/links instead.
	ID                 string                             `json:"id,omitempty"`
	Links              []*WorkflowTransitionLinkScheme    `json:"links,omitempty"`
	Name               string                             `json:"name,omitempty"` // The name of the transition.
	To                 *WorkflowStatusAndPortScheme       `json:"to,omitempty"`   // The status to which this transition goes.
	ToStatusReference  string                             `json:"toStatusReference,omitempty"`
	TransitionScreen   *WorkflowRuleConfigurationScheme   `json:"transitionScreen,omitempty"`
	Triggers           []*WorkflowTriggerScheme           `json:"triggers,omitempty"`
	Type               string                             `json:"type,omitempty"` // The type of the transition.
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

type WorkflowTransitionLinkScheme struct {
	FromPort            int    `json:"fromPort,omitempty"`
	FromStatusReference string `json:"fromStatusReference,omitempty"`
	ToPort              int    `json:"toPort,omitempty"`
}

type WorkflowTriggerScheme struct {
	ID      string `json:"id,omitempty"`
	RuleKey string `json:"ruleKey,omitempty"`
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

// WorkflowSearchCriteria represents the criteria for searching workflows in Jira.
type WorkflowSearchCriteria struct {
	ProjectAndIssueTypes []*WorkflowSearchProjectIssueTypeMapping `json:"projectAndIssueTypes,omitempty"` // ProjectAndIssueTypes is a list of project and issue type mappings to filter the search.
	WorkflowIDs          []string                                 `json:"workflowIds,omitempty"`          // WorkflowIDs is a list of workflow IDs to filter the search.
	WorkflowNames        []string                                 `json:"workflowNames,omitempty"`        // WorkflowNames is a list of workflow names to filter the search.
}

// WorkflowSearchProjectIssueTypeMapping represents a mapping of project and issue type for workflow search.
type WorkflowSearchProjectIssueTypeMapping struct {
	IssueTypeID string `json:"issueTypeId,omitempty"` // IssueTypeID is the ID of the issue type.
	ProjectID   string `json:"projectId,omitempty"`   // ProjectID is the ID of the project.
}

// WorkflowReadResponseScheme represents the response scheme for reading workflows in Jira.
type WorkflowReadResponseScheme struct {
	Statuses  []*WorkflowStatusDetailScheme `json:"statuses,omitempty"`  // Statuses is a list of workflow status details.
	Workflows []*JiraWorkflowScheme         `json:"workflows,omitempty"` // Workflows is a list of Jira workflows.
}

// JiraWorkflowScheme represents a workflow in Jira.
type JiraWorkflowScheme struct {
	Description      string                           `json:"description,omitempty"`      // Description is the description of the workflow.
	ID               string                           `json:"id,omitempty"`               // ID is the ID of the workflow.
	IsEditable       bool                             `json:"isEditable,omitempty"`       // IsEditable indicates if the workflow is editable.
	Name             string                           `json:"name,omitempty"`             // Name is the name of the workflow.
	Scope            *WorkflowStatusScopeScheme       `json:"scope,omitempty"`            // Scope is the scope of the workflow.
	StartPointLayout *WorkflowLayoutScheme            `json:"startPointLayout,omitempty"` // StartPointLayout is the layout of the start point of the workflow.
	Statuses         []*WorkflowReferenceStatusScheme `json:"statuses,omitempty"`         // Statuses is a list of reference statuses in the workflow.
	TaskID           string                           `json:"taskId,omitempty"`           // TaskID is the ID of the task associated with the workflow.
	Transitions      []*WorkflowTransitionScheme      `json:"transitions,omitempty"`      // Transitions is a list of transitions in the workflow.
	Usages           []*ProjectIssueTypesScheme       `json:"usages,omitempty"`           // Usages is a list of project issue types that use the workflow.
	Version          *WorkflowDocumentVersionScheme   `json:"version,omitempty"`          // Version is the version of the workflow document.
}

// WorkflowLayoutScheme represents the layout of a workflow element in Jira.
type WorkflowLayoutScheme struct {
	X float64 `json:"x,omitempty"` // X is the X coordinate of the layout.
	Y float64 `json:"y,omitempty"` // Y is the Y coordinate of the layout.
}

// WorkflowReferenceStatusScheme represents a reference status in a workflow in Jira.
type WorkflowReferenceStatusScheme struct {
	Deprecated      bool                  `json:"deprecated,omitempty"`      // Deprecated indicates if the status is deprecated.
	Layout          *WorkflowLayoutScheme `json:"layout,omitempty"`          // Layout is the layout of the status.
	StatusReference string                `json:"statusReference,omitempty"` // StatusReference is the reference of the status.
}

// WorkflowDocumentVersionScheme represents the version of a workflow document in Jira.
type WorkflowDocumentVersionScheme struct {
	ID            string `json:"id,omitempty"`            // ID is the ID of the document version.
	VersionNumber int    `json:"versionNumber,omitempty"` // VersionNumber is the version number of the document.
}

// WorkflowCapabilitiesScheme represents the capabilities of a workflow in Jira.
type WorkflowCapabilitiesScheme struct {
	ConnectRules []*AvailableWorkflowConnectRuleScheme `json:"connectRules,omitempty"` // ConnectRules is a list of available workflow connect rules.
	EditorScope  string                                `json:"editorScope,omitempty"`  // EditorScope is the scope of the editor.
	ForgeRules   []*AvailableWorkflowForgeRuleScheme   `json:"forgeRules,omitempty"`   // ForgeRules is a list of available workflow forge rules.
	SystemRules  []*AvailableWorkflowSystemRuleScheme  `json:"systemRules,omitempty"`  // SystemRules is a list of available workflow system rules.
	TriggerRules []*AvailableWorkflowTriggers          `json:"triggerRules,omitempty"` // TriggerRules is a list of available workflow trigger rules.
}

// AvailableWorkflowConnectRuleScheme represents a connect rule in a workflow.
type AvailableWorkflowConnectRuleScheme struct {
	AddonKey    string `json:"addonKey,omitempty"`    // AddonKey is the key of the addon.
	CreateURL   string `json:"createUrl,omitempty"`   // CreateURL is the URL to create the rule.
	Description string `json:"description,omitempty"` // Description is the description of the rule.
	EditURL     string `json:"editUrl,omitempty"`     // EditURL is the URL to edit the rule.
	ModuleKey   string `json:"moduleKey,omitempty"`   // ModuleKey is the key of the module.
	Name        string `json:"name,omitempty"`        // Name is the name of the rule.
	RuleKey     string `json:"ruleKey,omitempty"`     // RuleKey is the key of the rule.
	RuleType    string `json:"ruleType,omitempty"`    // RuleType is the type of the rule.
	ViewURL     string `json:"viewUrl,omitempty"`     // ViewURL is the URL to view the rule.
}

// AvailableWorkflowForgeRuleScheme represents a forge rule in a workflow.
type AvailableWorkflowForgeRuleScheme struct {
	Description string `json:"description,omitempty"` // Description is the description of the rule.
	ID          string `json:"id,omitempty"`          // ID is the ID of the rule.
	Name        string `json:"name,omitempty"`        // Name is the name of the rule.
	RuleKey     string `json:"ruleKey,omitempty"`     // RuleKey is the key of the rule.
	RuleType    string `json:"ruleType,omitempty"`    // RuleType is the type of the rule.
}

// AvailableWorkflowSystemRuleScheme represents a system rule in a workflow.
type AvailableWorkflowSystemRuleScheme struct {
	Description                     string   `json:"description,omitempty"`                     // Description is the description of the rule.
	IncompatibleRuleKeys            []string `json:"incompatibleRuleKeys,omitempty"`            // IncompatibleRuleKeys is a list of keys of incompatible rules.
	IsAvailableForInitialTransition bool     `json:"isAvailableForInitialTransition,omitempty"` // IsAvailableForInitialTransition indicates if the rule is available for the initial transition.
	IsVisible                       bool     `json:"isVisible,omitempty"`                       // IsVisible indicates if the rule is visible.
	Name                            string   `json:"name,omitempty"`                            // Name is the name of the rule.
	RuleKey                         string   `json:"ruleKey,omitempty"`                         // RuleKey is the key of the rule.
	RuleType                        string   `json:"ruleType,omitempty"`                        // RuleType is the type of the rule.
}

// AvailableWorkflowTriggers represents the triggers available in a workflow.
type AvailableWorkflowTriggers struct {
	AvailableTypes []*AvailableWorkflowTriggerTypeScheme `json:"availableTypes,omitempty"` // AvailableTypes is a list of available trigger types.
	RuleKey        string                                `json:"ruleKey,omitempty"`        // RuleKey is the key of the rule.
}

// AvailableWorkflowTriggerTypeScheme represents a trigger type in a workflow.
type AvailableWorkflowTriggerTypeScheme struct {
	Description string `json:"description,omitempty"` // Description is the description of the trigger type.
	Name        string `json:"name,omitempty"`        // Name is the name of the trigger type.
	Type        string `json:"type,omitempty"`        // Type is the type of the trigger.
}

// WorkflowCreatesPayload represents the payload for creating workflows in Jira.
type WorkflowCreatesPayload struct {
	Scope     *WorkflowScopeScheme          `json:"scope,omitempty"`     // Scope is the scope of the workflow.
	Statuses  []*WorkflowStatusUpdateScheme `json:"statuses,omitempty"`  // Statuses is a list of statuses in the workflow.
	Workflows []*WorkflowCreateScheme       `json:"workflows,omitempty"` // Workflows is a list of workflows to be created.
}

// AddWorkflow adds a new workflow and its statuses to the payload, ensuring no duplicate statuses.
func (w *WorkflowCreatesPayload) AddWorkflow(workflow *WorkflowCreateScheme) error {

	// Check if the status references are present in the payload
	for _, referenceID := range workflow.Statuses {
		found := false
		for _, status := range w.Statuses {
			if referenceID.StatusReference == status.StatusReference {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("status reference %s not found", referenceID.StatusReference)
		}
	}

	w.Workflows = append(w.Workflows, workflow)
	return nil
}

// AddStatus adds a new status to the WorkflowCreatesPayload.
func (w *WorkflowCreatesPayload) AddStatus(status *WorkflowStatusUpdateScheme) {
	w.Statuses = append(w.Statuses, status)
}

// WorkflowScopeScheme represents the scope of a workflow in Jira.
type WorkflowScopeScheme struct {
	Project *WorkflowScopeProjectScheme `json:"project,omitempty"` // Project is the project associated with the workflow.
	Type    string                      `json:"type,omitempty"`    // Valid values: PROJECT, GLOBAL.
}

// WorkflowScopeProjectScheme represents a project in the workflow scope.
type WorkflowScopeProjectScheme struct {
	ID string `json:"id,omitempty"` // ID is the ID of the project.
}

// WorkflowStatusUpdateScheme represents an update to a workflow status in Jira.
type WorkflowStatusUpdateScheme struct {
	Description     string `json:"description,omitempty"`     // Description is the description of the status.
	ID              string `json:"id,omitempty"`              // ID is the ID of the status.
	Name            string `json:"name,omitempty"`            // Name is the name of the status.
	StatusCategory  string `json:"statusCategory,omitempty"`  // StatusCategory is the category of the status.
	StatusReference string `json:"statusReference,omitempty"` // StatusReference is the reference of the status.
}

// WorkflowCreateScheme represents the creation of a workflow in Jira.
type WorkflowCreateScheme struct {
	Description      string                       `json:"description,omitempty"`      // Description is the description of the workflow.
	Name             string                       `json:"name,omitempty"`             // Name is the name of the workflow.
	StartPointLayout *WorkflowLayoutScheme        `json:"startPointLayout,omitempty"` // StartPointLayout is the layout of the start point of the workflow.
	Statuses         []*StatusLayoutUpdateScheme  `json:"statuses,omitempty"`         // Statuses is a list of statuses in the workflow.
	Transitions      []*TransitionUpdateDTOScheme `json:"transitions,omitempty"`      // Transitions is a list of transitions in the workflow.
}

// AddStatus adds a new status to the WorkflowCreateScheme.
func (w *WorkflowCreateScheme) AddStatus(status *StatusLayoutUpdateScheme) {
	w.Statuses = append(w.Statuses, status)
}

// AddTransition adds a new transition to the WorkflowCreateScheme.
func (w *WorkflowCreateScheme) AddTransition(transition *TransitionUpdateDTOScheme) error {
	if !w.isStatusReferenceAdded(transition.To.StatusReference) {
		return fmt.Errorf("status reference %s not found", transition.To.StatusReference)
	}

	w.Transitions = append(w.Transitions, transition)
	return nil
}

// isStatusReferenceAdded checks if a status reference is already added to the workflow.
// It returns true if the status reference is found, otherwise false.
func (w *WorkflowCreateScheme) isStatusReferenceAdded(statusReference string) bool {
	for _, status := range w.Statuses {
		if status.StatusReference == statusReference {
			return true
		}
	}
	return false
}

// StatusLayoutUpdateScheme represents an update to the layout of a status in a workflow.
type StatusLayoutUpdateScheme struct {
	// Layout is the layout of the status.
	Layout *WorkflowLayoutScheme `json:"layout,omitempty"`
	// StatusReference is the reference of the status.
	StatusReference string `json:"statusReference"`
}

// TransitionUpdateDTOScheme represents an update to a transition in a workflow.
type TransitionUpdateDTOScheme struct {
	Actions            []*WorkflowRuleConfigurationScheme `json:"actions,omitempty"`            // Actions is a list of actions associated with the transition.
	Conditions         *ConditionGroupUpdateScheme        `json:"conditions,omitempty"`         // Conditions is a list of conditions associated with the transition.
	CustomIssueEventID string                             `json:"customIssueEventId,omitempty"` // CustomIssueEventID is the custom issue event ID associated with the transition.
	Description        string                             `json:"description,omitempty"`        // Description is the description of the transition.
	From               []*StatusReferenceAndPortScheme    `json:"from,omitempty"`               // From is a list of statuses from which this transition can be executed.
	ID                 string                             `json:"id,omitempty"`                 // ID is the ID of the transition.
	Links              []*WorkflowTransitionLinkScheme    `json:"links,omitempty"`              // Links is a list of links associated with the transition.
	Name               string                             `json:"name,omitempty"`               // Name is the name of the transition.
	To                 *StatusReferenceAndPortScheme      `json:"to,omitempty"`                 // To is the status to which this transition goes.
	ToStatusReference  string                             `json:"toStatusReference,omitempty"`  // ToStatusReference is the reference of the status to which this transition goes.
	TransitionScreen   *WorkflowRuleConfigurationScheme   `json:"transitionScreen,omitempty"`   // TransitionScreen is the screen associated with the transition.
	Triggers           []*WorkflowTriggerScheme           `json:"triggers,omitempty"`           // Triggers is a list of triggers associated with the transition.
	Type               string                             `json:"type,omitempty"`               // Type is the type of the transition.
	Validators         []*WorkflowRuleConfigurationScheme `json:"validators,omitempty"`         // Validators is a list of validators associated with the transition.
}

// ConditionGroupUpdateScheme represents an update to a condition group in a workflow.
type ConditionGroupUpdateScheme struct {
	ConditionGroups []*ConditionGroupUpdateScheme      `json:"conditionGroups,omitempty"` // ConditionGroups is a list of nested condition groups.
	Conditions      []*WorkflowRuleConfigurationScheme `json:"conditions,omitempty"`      // Conditions is a list of conditions.
	Operation       string                             `json:"operation,omitempty"`       // Operation is the operation applied to the conditions.
}

// StatusReferenceAndPortScheme represents a status reference and port in a workflow.
type StatusReferenceAndPortScheme struct {
	Port            int    `json:"port,omitempty"`            // Port is the port associated with the status.
	StatusReference string `json:"statusReference,omitempty"` // StatusReference is the reference of the status.
}

// WorkflowCreateResponseScheme represents the response after creating a workflow in Jira.
type WorkflowCreateResponseScheme struct {
	Statuses  []*JiraWorkflowStatusScheme `json:"statuses,omitempty"`  // Statuses is a list of statuses in the workflow.
	Workflows []*JiraWorkflowScheme       `json:"workflows,omitempty"` // Workflows is a list of Jira workflows.
}

// JiraWorkflowStatusScheme represents a status in a Jira workflow.
type JiraWorkflowStatusScheme struct {
	Description     string                     `json:"description,omitempty"`     // Description is the description of the status.
	ID              string                     `json:"id,omitempty"`              // ID is the ID of the status.
	Name            string                     `json:"name,omitempty"`            // Name is the name of the status.
	Scope           *WorkflowScopeScheme       `json:"scope,omitempty"`           // Scope is the scope of the status.
	StatusCategory  string                     `json:"statusCategory,omitempty"`  // StatusCategory is the category of the status.
	StatusReference string                     `json:"statusReference,omitempty"` // StatusReference is the reference of the status.
	Usages          []*ProjectIssueTypesScheme `json:"usages,omitempty"`          // Usages is a list of project issue types that use the status.
}

// ValidationOptionsForCreateScheme represents the validation options for creating a workflow.
type ValidationOptionsForCreateScheme struct {
	Payload *WorkflowCreatesPayload       `json:"payload,omitempty"` // Payload is the payload for creating workflows.
	Options *ValidationOptionsLevelScheme `json:"validationOptions"` // Options are the validation options.
}

// ValidationOptionsLevelScheme represents the levels of validation options.
type ValidationOptionsLevelScheme struct {
	Levels []string `json:"levels,omitempty"` // Valid values: WARNING, ERROR.
}

// WorkflowValidationErrorListScheme represents a list of workflow validation errors.
type WorkflowValidationErrorListScheme struct {
	Errors []*WorkflowValidationErrorScheme `json:"errors,omitempty"` // Errors is a list of workflow validation errors.
}

// WorkflowValidationErrorScheme represents a workflow validation error.
type WorkflowValidationErrorScheme struct {
	Code             string                          `json:"code,omitempty"`             // Valid values: INVALID_TRANSITION, INVALID_STATUS, INVALID_TRANSITION_TO_STATUS, INVALID_TRANSITION_FROM_STATUS, INVALID_TRANSITION_TO_STATUS
	ElementReference *WorkflowElementReferenceScheme `json:"elementReference,omitempty"` // ElementReference is the reference to the element that caused the error.
	Level            string                          `json:"level,omitempty"`            // Valid values: WARNING, ERROR.
	Message          string                          `json:"message,omitempty"`          // Message is the error message.
	Type             string                          `json:"type,omitempty"`             // Valid values: TRANSITION, STATUS
}

// WorkflowElementReferenceScheme represents a reference to an element in a workflow.
type WorkflowElementReferenceScheme struct {
	PropertyKey            string                         `json:"propertyKey,omitempty"`            // PropertyKey is the key of the property.
	RuleID                 string                         `json:"ruleId,omitempty"`                 // RuleID is the ID of the rule.
	StatusMappingReference *ProjectAndIssueTypePairScheme `json:"statusMappingReference,omitempty"` // StatusMappingReference is the reference to the status mapping.
	StatusReference        string                         `json:"statusReference,omitempty"`        // StatusReference is the reference of the status.
	TransitionID           string                         `json:"transitionId,omitempty"`           // TransitionID is the ID of the transition.
}

// ProjectAndIssueTypePairScheme represents a pair of project and issue type.
type ProjectAndIssueTypePairScheme struct {
	IssueTypeID string `json:"issueTypeId,omitempty"` // IssueTypeID is the ID of the issue type.
	ProjectID   string `json:"projectId,omitempty"`   // ProjectID is the ID of the project.
}

// WorkflowUpdatesPayloadScheme represents the payload for updating workflows in Jira.
type WorkflowUpdatesPayloadScheme struct {
	Statuses  []*WorkflowStatusUpdateScheme `json:"statuses,omitempty"`  // Statuses is a list of statuses in the workflow.
	Workflows []*WorkflowUpdateScheme       `json:"workflows,omitempty"` // Workflows is a list of workflows to be updated.
}

// InjectWorkflow adds a new workflow to the payload for updating workflows.
// It takes a JiraWorkflowScheme as input and appends it to the Workflows slice.
func (w *WorkflowUpdatesPayloadScheme) InjectWorkflow(workflow *JiraWorkflowScheme) {
	w.Workflows = append(w.Workflows, &WorkflowUpdateScheme{
		Description:      workflow.Description,
		ID:               workflow.ID,
		StartPointLayout: workflow.StartPointLayout,
		Version:          workflow.Version,
	})
}

// WorkflowUpdateScheme represents the update scheme for a workflow in Jira.
type WorkflowUpdateScheme struct {
	DefaultStatusMappings []*StatusMigrationScheme       `json:"defaultStatusMappings,omitempty"` // DefaultStatusMappings is a list of default status mappings.
	Description           string                         `json:"description,omitempty"`           // Description is the description of the workflow.
	ID                    string                         `json:"id,omitempty"`                    // ID is the ID of the workflow.
	StartPointLayout      *WorkflowLayoutScheme          `json:"startPointLayout,omitempty"`      // StartPointLayout is the layout of the start point of the workflow.
	StatusMappings        []*StatusMappingDTOScheme      `json:"statusMappings,omitempty"`        // StatusMappings is a list of status mappings in the workflow.
	Statuses              []*StatusLayoutUpdateScheme    `json:"statuses,omitempty"`              // Statuses is a list of statuses in the workflow.
	Transitions           []*TransitionUpdateDTOScheme   `json:"transitions,omitempty"`           // Transitions is a list of transitions in the workflow.
	Version               *WorkflowDocumentVersionScheme `json:"version,omitempty"`               // Version is the version of the workflow document.
}

// StatusMigrationScheme represents a status migration in a workflow in Jira.
type StatusMigrationScheme struct {
	NewStatusReference string `json:"newStatusReference,omitempty"` // NewStatusReference is the reference of the new status.
	OldStatusReference string `json:"oldStatusReference,omitempty"` // OldStatusReference is the reference of the old status.
}

// StatusMappingDTOScheme represents a status mapping DTO in a workflow in Jira.
type StatusMappingDTOScheme struct {
	IssueTypeID      string                       `json:"issueTypeId,omitempty"`      // IssueTypeID is the ID of the issue type.
	ProjectID        string                       `json:"projectId,omitempty"`        // ProjectID is the ID of the project.
	StatusMigrations []*StatusMigrationScheme     `json:"statusMigrations,omitempty"` // StatusMigrations is a list of status migrations.
	Statuses         []*StatusLayoutUpdateScheme  `json:"statuses,omitempty"`         // Statuses is a list of statuses in the workflow.
	Transitions      []*TransitionUpdateDTOScheme `json:"transitions,omitempty"`      // Transitions is a list of transitions in the workflow.
}

// WorkflowUpdateResponseScheme represents the response after updating a workflow in Jira.
type WorkflowUpdateResponseScheme struct {
	Statuses  []*JiraWorkflowStatusScheme `json:"statuses,omitempty"`  // Statuses is a list of statuses in the workflow.
	TaskID    string                      `json:"taskId,omitempty"`    // TaskID is the ID of the task associated with the workflow update.
	Workflows []*JiraWorkflowScheme       `json:"workflows,omitempty"` // Workflows is a list of Jira workflows.
}

// ValidationOptionsForUpdateScheme represents the validation options for updating a workflow.
type ValidationOptionsForUpdateScheme struct {
	Payload *WorkflowUpdatesPayloadScheme `json:"payload,omitempty"`           // Payload is the payload for updating workflows.
	Options *ValidationOptionsLevelScheme `json:"validationOptions,omitempty"` // Options are the validation options.
}
