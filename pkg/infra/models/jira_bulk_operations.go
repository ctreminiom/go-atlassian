package models

type IssueBulkEditPayloadScheme struct {
	Fields           *JiraIssueFieldsScheme `json:"editedFieldsInput"`
	Actions          []string               `json:"selectedActions"`
	Issues           []string               `json:"selectedIssueIdsOrKeys"`
	SendNotification bool                   `json:"sendNotification,omitempty"`
}

// JiraIssueFieldsScheme represents the fields of a Jira issue for bulk operations.
type JiraIssueFieldsScheme struct {

	/*
		Add or clear a cascading select field:
			- To add, specify optionId for both parent and child.
			- To clear the child, set its optionId to null.
			- To clear both, set the parent's optionId to null.
	*/
	CascadingSelectFields []*JiraCascadingSelectFieldScheme `json:"cascadingSelectFields,omitempty"`

	/*
		Add or clear a number field:
			- To add, specify a numeric value.
			- To clear, set value to null.
	*/
	ClearableNumberFields []*JiraNumberFieldScheme `json:"clearableNumberFields,omitempty"`

	/*
		Add or clear a color field:
			- To add, specify the color name. Available colors are: purple, blue, green, teal, yellow, orange, grey, dark purple, dark blue, dark green, dark teal, dark yellow, dark orange, dark grey.
			- To clear, set the color name to an empty string.
	*/
	ColorFields []*JiraColorFieldScheme `json:"colorFields,omitempty"`

	/*
		Add or clear a date picker field:
			- To add, specify the date in d/mmm/yy format or ISO format dd-mm-yyyy.
			- To clear, set formattedDate to an empty string.
	*/
	DatePickerFields []*JiraDateFieldScheme `json:"datePickerFields,omitempty"`

	/*
		Add or clear the planned start date and time:
			- To add, specify the date and time in ISO format for formattedDateTime.
			- To clear, provide an empty string for formattedDateTime.
	*/
	DateTimePickerFields []*JiraDateTimeFieldScheme `json:"dateTimePickerFields,omitempty"`

	/*
		Set the issue type field by providing an issueTypeId.
	*/
	IssueType *JiraIssueTypeFieldScheme `json:"issueType,omitempty"`

	/*
		Edit a labels field:
			- Options include ADD, REPLACE, REMOVE, or REMOVE_ALL for bulk edits.
			- To clear labels, use the REMOVE_ALL option with an empty labels array.
	*/
	LabelsFields []*JiraLabelFieldScheme `json:"labelsFields,omitempty"`

	/*
		Add or clear a multi-group picker field:
			- To add groups, provide an array of groups with groupNames.
			- To clear all groups, use an empty groups array.
	*/
	MultipleGroupPickerFields []*JiraMultipleGroupPickerFieldScheme `json:"multipleGroupPickerFields,omitempty"`

	/*
		Assign or unassign multiple users to/from a field:
			- To assign, provide an array of user accountIds.
			- To clear, set users to null.
	*/
	MultipleSelectClearableUserPickerFields []*JiraMultipleSelectClearableUserPickerFieldScheme `json:"multipleSelectClearableUserPickerFields,omitempty"`

	/*
		Add or clear a multi-select field:
			- To add, provide an array of options with optionIds.
			- To clear, use an empty options array.
	*/
	MultipleSelectFields []*JiraMultipleSelectFieldScheme `json:"multipleSelectFields,omitempty"`

	/*
		Edit a multi-version picker field like Fix Versions/Affects Versions:
			- Options include ADD, REPLACE, REMOVE, or REMOVE_ALL for bulk edits.
			- To clear the field, use the REMOVE_ALL option with an empty versions array.
	*/
	MultipleVersionPickerFields []*JiraMultipleVersionPickerFieldScheme `json:"multipleVersionPickerFields,omitempty"`

	/*
		Edit a multi select components field:
			- Options include ADD, REPLACE, REMOVE, or REMOVE_ALL for bulk edits.
			- To clear, use the REMOVE_ALL option with an empty components array.
	*/
	MultiselectComponents *JiraMultiSelectComponentFieldScheme `json:"multiselectComponents,omitempty"`

	// Edit the original estimate field.
	OriginalEstimateField *JiraDurationFieldScheme `json:"originalEstimateField,omitempty"`

	// Set the priority of an issue by specifying a priorityId.
	Priority *JiraPriorityFieldScheme `json:"priority,omitempty"`

	/*
		Add or clear a rich text field:
			- To add, provide adfValue. Note that rich text fields only support ADF values.
			- To clear, use an empty richText object.
	*/
	RichTextFields []*JiraRichTextFieldScheme `json:"richTextFields,omitempty"`

	/*
		Add or clear a single group picker field:
			- To add, specify the group with groupName
			- To clear, set groupName to an empty string.
	*/
	SingleGroupPickerFields []*JiraSingleGroupPickerFieldScheme `json:"singleGroupPickerFields,omitempty"`

	/*
		Add or clear a single line text field:
			- To add, provide the text value.
			- To clear, set text to an empty string.
	*/
	SingleLineTextFields []*JiraSingleLineTextFieldScheme `json:"singleLineTextFields,omitempty"`

	/*
		Edit assignment for single select user picker fields like Assignee/Reporter:
			- To assign an issue, specify the user's accountId.
			- To unassign an issue, set user to null.
			- For automatic assignment, set accountId to -1.
	*/
	SingleSelectClearableUserPickerFields []*JiraSingleSelectUserPickerFieldScheme `json:"singleSelectClearableUserPickerFields,omitempty"`

	/*
		Add or clear a single select field:
			- To add, specify the option with an optionId.
			- To clear, pass an option with optionId as -1.
	*/
	SingleSelectFields []*JiraSingleSelectFieldScheme `json:"singleSelectFields,omitempty"`

	/*
		Add or clear a single version picker field:
			- To add, specify the version with a versionId.
			- To clear, set versionId to -1.
	*/
	SingleVersionPickerFields []*JiraSingleVersionPickerFieldScheme `json:"singleVersionPickerFields,omitempty"`

	// Edit the time tracking field.
	TimeTrackingField *JiraTimeTrackingFieldScheme `json:"timeTrackingField,omitempty"`

	/*
		Add or clear a URL field:
			- To add, provide the url with the desired URL value.
			- To clear, set url to an empty string.
	*/
	UrlFields []*JiraUrlFieldScheme `json:"urlFields,omitempty"`
}

// JiraCascadingSelectFieldScheme represents a cascading select field in Jira.
type JiraCascadingSelectFieldScheme struct {
	ChildOptionValue  *JiraSelectedOptionFieldSchemes `json:"childOptionValue,omitempty"`
	FieldId           string                          `json:"fieldId,omitempty"`
	ParentOptionValue *JiraSelectedOptionFieldSchemes `json:"parentOptionValue,omitempty"`
}

// JiraSelectedOptionFieldSchemes represents the selected option for a field in Jira.
type JiraSelectedOptionFieldSchemes struct {
	OptionID int `json:"optionId,omitempty"`
}

// JiraNumberFieldScheme represents a number field in Jira.
type JiraNumberFieldScheme struct {
	FieldID string `json:"fieldId,omitempty"`
	Value   int    `json:"value,omitempty"`
}

// JiraColorFieldScheme represents a color field in Jira.
type JiraColorFieldScheme struct {
	Color   *JiraColorInputScheme `json:"color,omitempty"`
	FieldID string                `json:"fieldId,omitempty"`
}

// JiraColorInputScheme represents the input scheme for a color field in Jira.
type JiraColorInputScheme struct {
	Name string `json:"name,omitempty"`
}

// JiraDateFieldScheme represents a date field in Jira.
type JiraDateFieldScheme struct {
	Date    *JiraDateInputScheme `json:"date,omitempty"`
	FieldId string               `json:"fieldId,omitempty"`
}

// JiraDateInputScheme represents the input scheme for a date field in Jira.
type JiraDateInputScheme struct {
	FormattedDate string `json:"formattedDate,omitempty"`
}

// JiraDateTimeFieldScheme represents a date-time field in Jira.
type JiraDateTimeFieldScheme struct {
	DateTime *JiraDateTimeInputScheme `json:"dateTime,omitempty"`
	FieldID  string                   `json:"fieldId,omitempty"`
}

// JiraDateTimeInputScheme represents the input scheme for a date-time field in Jira.
type JiraDateTimeInputScheme struct {
	FormattedDateTime string `json:"formattedDateTime,omitempty"`
}

// JiraIssueTypeFieldScheme represents an issue type field in Jira.
type JiraIssueTypeFieldScheme struct {
	IssueTypeID string `json:"issueTypeId,omitempty"`
}

// JiraLabelFieldScheme represents a label field in Jira.
type JiraLabelFieldScheme struct {
	BulkEditMultiSelectFieldOption string                  `json:"bulkEditMultiSelectFieldOption,omitempty"`
	FieldID                        string                  `json:"fieldId,omitempty"`
	Labels                         []*JiraLabelInputScheme `json:"labels,omitempty"`
}

// JiraLabelInputScheme represents the input scheme for a label in Jira.
type JiraLabelInputScheme struct {
	Name string `json:"name,omitempty"`
}

// JiraMultipleGroupPickerFieldScheme represents a multiple group picker field in Jira.
type JiraMultipleGroupPickerFieldScheme struct {
	FieldID string                  `json:"fieldId,omitempty"`
	Groups  []*JiraGroupInputScheme `json:"groups,omitempty"`
}

// JiraGroupInputScheme represents the input scheme for a group in Jira.
type JiraGroupInputScheme struct {
	GroupName string `json:"groupName,omitempty"`
}

// JiraMultipleSelectClearableUserPickerFieldScheme represents a multiple select clearable user picker field in Jira.
type JiraMultipleSelectClearableUserPickerFieldScheme struct {
	FieldID string                 `json:"fieldId,omitempty"`
	Users   []*JiraUserFieldScheme `json:"users,omitempty"`
}

// JiraUserFieldScheme represents the input scheme for a user in Jira.
type JiraUserFieldScheme struct {
	AccountID string `json:"accountId,omitempty"`
}

// JiraMultipleSelectFieldScheme represents a multiple select field in Jira.
type JiraMultipleSelectFieldScheme struct {
	FieldID string                     `json:"fieldId,omitempty"`
	Options []*JiraSelectedOptionField `json:"options,omitempty"`
}

// JiraSelectedOptionField represents a selected option field in Jira.
type JiraSelectedOptionField struct {
	OptionID int `json:"optionId,omitempty"`
}

// JiraMultipleVersionPickerFieldScheme represents a multiple version picker field in Jira.
type JiraMultipleVersionPickerFieldScheme struct {
	BulkEditMultiSelectFieldOption string                    `json:"bulkEditMultiSelectFieldOption,omitempty"`
	FieldId                        string                    `json:"fieldId,omitempty"`
	Versions                       []*JiraVersionFieldScheme `json:"versions,omitempty"`
}

// JiraVersionFieldScheme represents a version field in Jira.
type JiraVersionFieldScheme struct {
	VersionID string `json:"versionId,omitempty"`
}

// JiraMultiSelectComponentFieldScheme represents a multi-select component field in Jira.
type JiraMultiSelectComponentFieldScheme struct {
	BulkEditMultiSelectFieldOption string                      `json:"bulkEditMultiSelectFieldOption,omitempty"`
	Components                     []*JiraComponentFieldScheme `json:"components,omitempty"`
	FieldID                        string                      `json:"fieldId,omitempty"`
}

// JiraComponentFieldScheme represents a component field in Jira.
type JiraComponentFieldScheme struct {
	ComponentID int `json:"componentId,omitempty"`
}

// JiraDurationFieldScheme represents a duration field in Jira.
type JiraDurationFieldScheme struct {
	OriginalEstimateField string `json:"originalEstimateField,omitempty"`
}

// JiraPriorityFieldScheme represents a priority field in Jira.
type JiraPriorityFieldScheme struct {
	PriorityID string `json:"priorityId,omitempty"`
}

// JiraRichTextFieldScheme represents a rich text field in Jira.
type JiraRichTextFieldScheme struct {
	FieldID  string                   `json:"fieldId,omitempty"`
	RichText *JiraRichTextInputScheme `json:"richText,omitempty"`
}

// JiraRichTextInputScheme represents the input scheme for a rich text field in Jira.
type JiraRichTextInputScheme struct {
	AdfValue *CommentNodeScheme `json:"adfValue,omitempty"`
}

// JiraSingleGroupPickerFieldScheme represents a single group picker field in Jira.
type JiraSingleGroupPickerFieldScheme struct {
	FieldID string                `json:"fieldId,omitempty"`
	Group   *JiraGroupInputScheme `json:"group,omitempty"`
}

// JiraSingleLineTextFieldScheme represents a single line text field in Jira.
type JiraSingleLineTextFieldScheme struct {
	FieldID string `json:"fieldId,omitempty"`
	Text    string `json:"text,omitempty"`
}

// JiraSingleSelectUserPickerFieldScheme represents a single select user picker field in Jira.
type JiraSingleSelectUserPickerFieldScheme struct {
	FieldID string               `json:"fieldId,omitempty"`
	User    *JiraUserFieldScheme `json:"user,omitempty"`
}

// JiraSingleSelectFieldScheme represents a single select field in Jira.
type JiraSingleSelectFieldScheme struct {
	FieldID string                         `json:"fieldId,omitempty"`
	Option  *JiraSelectedOptionFieldScheme `json:"option,omitempty"`
}

// JiraSelectedOptionFieldScheme represents the selected option for a field in Jira.
type JiraSelectedOptionFieldScheme struct {
	OptionID int `json:"optionId,omitempty"`
}

// JiraSingleVersionPickerFieldScheme represents a single version picker field in Jira.
type JiraSingleVersionPickerFieldScheme struct {
	FieldID string                  `json:"fieldId,omitempty"`
	Version *JiraVersionFieldScheme `json:"version,omitempty"`
}

// JiraTimeTrackingFieldScheme represents a time tracking field in Jira.
type JiraTimeTrackingFieldScheme struct {
	TimeRemaining string `json:"timeRemaining,omitempty"`
}

// JiraUrlFieldScheme represents a URL field in Jira.
type JiraUrlFieldScheme struct {
	FieldID string `json:"fieldId,omitempty"`
	Url     string `json:"url,omitempty"`
}

type BulkTransitionSubmitInputScheme struct {
	SelectedIssueIdsOrKeys []string `json:"selectedIssueIdsOrKeys"`
	TransitionID           string   `json:"transitionId"`
}

type BulkEditGetFieldsScheme struct {
	EndingBefore  string                      `json:"endingBefore,omitempty"`
	Fields        []*IssueBulkEditFieldScheme `json:"fields,omitempty"`
	StartingAfter string                      `json:"startingAfter,omitempty"`
}

type IssueBulkEditFieldScheme struct {
	Description             string   `json:"description,omitempty"`
	ID                      string   `json:"id,omitempty"`
	IsRequired              bool     `json:"isRequired,omitempty"`
	MultiSelectFieldOptions []string `json:"multiSelectFieldOptions,omitempty"`
	Name                    string   `json:"name,omitempty"`
	SearchUrl               string   `json:"searchUrl,omitempty"`
	Type                    string   `json:"type,omitempty"`
	UnavailableMessage      string   `json:"unavailableMessage,omitempty"`
}

type SubmittedBulkOperationScheme struct {
	TaskID string `json:"taskId"`
}

type BulkTransitionGetAvailableTransitionsScheme struct {
	AvailableTransitions []*IssueBulkTransitionForWorkflowScheme `json:"availableTransitions,omitempty"`
	EndingBefore         string                                  `json:"endingBefore,omitempty"`
	StartingAfter        string                                  `json:"startingAfter,omitempty"`
}

type IssueBulkTransitionForWorkflowScheme struct {
	IsTransitionsFiltered bool     `json:"isTransitionsFiltered,omitempty"`
	Issues                []string `json:"issues,omitempty"`
}

type SimplifiedIssueTransitionScheme struct {
	To             *IssueTransitionStatusScheme `json:"to,omitempty"`
	TransitionID   int                          `json:"transitionId,omitempty"`
	TransitionName string                       `json:"transitionName,omitempty"`
}

type IssueTransitionStatusScheme struct {
	StatusID   int    `json:"statusId,omitempty"`
	StatusName string `json:"statusName,omitempty"`
}

type BulkOperationProgressScheme struct {
	Created                         int64  `json:"created"`
	InvalidOrInaccessibleIssueCount int    `json:"invalidOrInaccessibleIssueCount"`
	ProcessedAccessibleIssues       []int  `json:"processedAccessibleIssues"`
	ProgressPercent                 int    `json:"progressPercent"`
	Started                         int64  `json:"started"`
	Status                          string `json:"status"`
	TaskID                          string `json:"taskId"`
	TotalIssueCount                 int    `json:"totalIssueCount"`
	Updated                         int64  `json:"updated"`
	SubmittedBy                     struct {
		AccountId string `json:"accountId"`
	} `json:"submittedBy"`
}
