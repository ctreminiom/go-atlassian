package models

import (
	"encoding/json"

	"dario.cat/mergo"
)

// IssueScheme represents an issue in Jira.
type IssueScheme struct {
	ID             string                   `json:"id,omitempty"`
	Key            string                   `json:"key,omitempty"`
	Self           string                   `json:"self,omitempty"`
	Transitions    []*IssueTransitionScheme `json:"transitions,omitempty"`
	Changelog      *IssueChangelogScheme    `json:"changelog,omitempty"`
	Fields         *IssueFieldsScheme       `json:"fields,omitempty"`
	RenderedFields map[string]interface{}   `json:"renderedFields,omitempty"`
}

// MergeCustomFields merges custom fields into the issue scheme.
// It returns a map representation of the issue scheme with the merged fields.
// If the provided fields are nil or empty, it returns an error.
func (i *IssueScheme) MergeCustomFields(fields *CustomFields) (map[string]interface{}, error) {

	if fields == nil || len(fields.Fields) == 0 {
		return map[string]interface{}{}, nil
	}

	//Convert the IssueScheme struct to map[string]interface{}
	issueSchemeAsBytes, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	issueSchemeAsMap := make(map[string]interface{})
	if err := json.Unmarshal(issueSchemeAsBytes, &issueSchemeAsMap); err != nil {
		return nil, err
	}

	//For each customField created, merge it into the eAsMap
	for _, customField := range fields.Fields {
		if err := mergo.Merge(&issueSchemeAsMap, customField, mergo.WithOverride); err != nil {
			return nil, err
		}
	}

	return issueSchemeAsMap, nil
}

// MergeOperations merges operations into the issue scheme.
// It returns a map representation of the issue scheme with the merged operations.
// If the provided operations are nil or empty, it returns an error.
//
// Parameters:
// - operations: A pointer to UpdateOperations containing the operations to be merged.
//
// Returns:
// - A map[string]interface{} representing the issue scheme with the merged operations.
// - An error if the operations are nil, empty, or if there is an issue during the merging process.
func (i *IssueScheme) MergeOperations(operations *UpdateOperations) (map[string]interface{}, error) {

	if operations == nil || len(operations.Fields) == 0 {
		return map[string]interface{}{}, nil
	}

	//Convert the IssueScheme struct to map[string]interface{}
	issueSchemeAsBytes, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	issueSchemeAsMap := make(map[string]interface{})
	if err := json.Unmarshal(issueSchemeAsBytes, &issueSchemeAsMap); err != nil {
		return nil, err
	}

	//For each customField created, merge it into the eAsMap
	for _, customField := range operations.Fields {
		if err := mergo.Merge(&issueSchemeAsMap, customField, mergo.WithOverride); err != nil {
			return nil, err
		}
	}

	return issueSchemeAsMap, nil
}

// ToMap converts the issue to a map[string]interface{}.
// It returns a map[string]interface{} representing the issue and an error if any occurred.
func (i *IssueScheme) ToMap() (map[string]interface{}, error) {

	//Convert the IssueScheme struct to map[string]interface{}
	issueSchemeAsBytes, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	issueSchemeAsMap := make(map[string]interface{})
	if err := json.Unmarshal(issueSchemeAsBytes, &issueSchemeAsMap); err != nil {
		return nil, err
	}

	return issueSchemeAsMap, nil
}

// IssueFieldsScheme represents the fields of an issue in Jira.
type IssueFieldsScheme struct {
	Parent                   *ParentScheme              `json:"parent,omitempty"`                   // The parent of the issue.
	IssueType                *IssueTypeScheme           `json:"issuetype,omitempty"`                // The type of the issue.
	IssueLinks               []*IssueLinkScheme         `json:"issuelinks,omitempty"`               // The links associated with the issue.
	Watcher                  *IssueWatcherScheme        `json:"watches,omitempty"`                  // The watchers of the issue.
	Votes                    *IssueVoteScheme           `json:"votes,omitempty"`                    // The votes for the issue.
	Versions                 []*VersionScheme           `json:"versions,omitempty"`                 // The versions associated with the issue.
	Project                  *ProjectScheme             `json:"project,omitempty"`                  // The project the issue belongs to.
	FixVersions              []*VersionScheme           `json:"fixVersions,omitempty"`              // The fix versions for the issue.
	Priority                 *PriorityScheme            `json:"priority,omitempty"`                 // The priority of the issue.
	Components               []*ComponentScheme         `json:"components,omitempty"`               // The components associated with the issue.
	Creator                  *UserScheme                `json:"creator,omitempty"`                  // The user who created the issue.
	Reporter                 *UserScheme                `json:"reporter,omitempty"`                 // The user who reported the issue.
	Assignee                 *UserScheme                `json:"assignee,omitempty"`                 // The user assigned to the issue.
	Resolution               *ResolutionScheme          `json:"resolution,omitempty"`               // The resolution of the issue.
	Resolutiondate           *DateTimeScheme            `json:"resolutiondate,omitempty"`           // The date the issue was resolved.
	Workratio                int                        `json:"workratio,omitempty"`                // The work ratio of the issue.
	StatusCategoryChangeDate *DateTimeScheme            `json:"statuscategorychangedate,omitempty"` // The date the status category changed.
	LastViewed               string                     `json:"lastViewed,omitempty"`               // The last time the issue was viewed.
	Summary                  string                     `json:"summary,omitempty"`                  // The summary of the issue.
	Created                  *DateTimeScheme            `json:"created,omitempty"`                  // The date the issue was created.
	Updated                  *DateTimeScheme            `json:"updated,omitempty"`                  // The date the issue was last updated.
	Labels                   []string                   `json:"labels,omitempty"`                   // The labels associated with the issue.
	Status                   *StatusScheme              `json:"status,omitempty"`                   // The status of the issue.
	Description              *CommentNodeScheme         `json:"description,omitempty"`              // The description of the issue.
	Comment                  *IssueCommentPageScheme    `json:"comment,omitempty"`                  // The comments on the issue.
	Subtasks                 []*IssueScheme             `json:"subtasks,omitempty"`                 // The subtasks of the issue.
	Security                 *SecurityScheme            `json:"security,omitempty"`                 // The security level of the issue.
	Attachment               []*AttachmentScheme        `json:"attachment,omitempty"`               // The attachments of the issue.
	Worklog                  *IssueWorklogADFPageScheme `json:"worklog,omitempty"`                  // The worklog of the issue.
	DueDate                  *DateScheme                `json:"duedate,omitempty"`                  // The due date of the issue.
}

// IssueTransitionScheme represents a transition of an issue in Jira.
type IssueTransitionScheme struct {
	ID            string        `json:"id,omitempty"`            // The ID of the transition.
	Name          string        `json:"name,omitempty"`          // The name of the transition.
	To            *StatusScheme `json:"to,omitempty"`            // The status the issue transitions to.
	HasScreen     bool          `json:"hasScreen,omitempty"`     // Indicates if the transition has a screen.
	IsGlobal      bool          `json:"isGlobal,omitempty"`      // Indicates if the transition is global.
	IsInitial     bool          `json:"isInitial,omitempty"`     // Indicates if the transition is initial.
	IsAvailable   bool          `json:"isAvailable,omitempty"`   // Indicates if the transition is available.
	IsConditional bool          `json:"isConditional,omitempty"` // Indicates if the transition is conditional.
	IsLooped      bool          `json:"isLooped,omitempty"`      // Indicates if the transition is looped.
}

// StatusScheme represents the status of an issue in Jira.
type StatusScheme struct {
	Self           string                `json:"self,omitempty"`           // The URL of the status.
	Description    string                `json:"description,omitempty"`    // The description of the status.
	IconURL        string                `json:"iconUrl,omitempty"`        // The icon URL of the status.
	Name           string                `json:"name,omitempty"`           // The name of the status.
	ID             string                `json:"id,omitempty"`             // The ID of the status.
	StatusCategory *StatusCategoryScheme `json:"statusCategory,omitempty"` // The category of the status.
}

// StatusCategoryScheme represents the category of a status in Jira.
type StatusCategoryScheme struct {
	Self      string `json:"self,omitempty"`      // The URL of the status category.
	ID        int    `json:"id,omitempty"`        // The ID of the status category.
	Key       string `json:"key,omitempty"`       // The key of the status category.
	ColorName string `json:"colorName,omitempty"` // The color name of the status category.
	Name      string `json:"name,omitempty"`      // The name of the status category.
}

// IssueNotifyOptionsScheme represents the options for notifying about an issue in Jira.
type IssueNotifyOptionsScheme struct {
	HTMLBody string                     `json:"htmlBody,omitempty"` // The HTML body of the notification.
	Subject  string                     `json:"subject,omitempty"`  // The subject of the notification.
	TextBody string                     `json:"textBody,omitempty"` // The text body of the notification.
	To       *IssueNotifyToScheme       `json:"to,omitempty"`       // The recipients of the notification.
	Restrict *IssueNotifyRestrictScheme `json:"restrict,omitempty"` // The restrictions for the notification.
}

// IssueNotifyRestrictScheme represents the restrictions for notifying about an issue in Jira.
type IssueNotifyRestrictScheme struct {
	Groups      []*IssueNotifyGroupScheme      `json:"groups,omitempty"`      // The groups to restrict the notification to.
	Permissions []*IssueNotifyPermissionScheme `json:"permissions,omitempty"` // The permissions to restrict the notification to.
}

// IssueNotifyToScheme represents the recipients for notifying about an issue in Jira.
type IssueNotifyToScheme struct {
	Reporter bool                      `json:"reporter,omitempty"` // Indicates if the reporter should be notified.
	Assignee bool                      `json:"assignee,omitempty"` // Indicates if the assignee should be notified.
	Watchers bool                      `json:"watchers,omitempty"` // Indicates if the watchers should be notified.
	Voters   bool                      `json:"voters,omitempty"`   // Indicates if the voters should be notified.
	Users    []*IssueNotifyUserScheme  `json:"users,omitempty"`    // The users to notify.
	Groups   []*IssueNotifyGroupScheme `json:"groups,omitempty"`   // The groups to notify.
}

// IssueNotifyPermissionScheme represents a permission for notifying about an issue in Jira.
type IssueNotifyPermissionScheme struct {
	ID  string `json:"id,omitempty"`  // The ID of the permission.
	Key string `json:"key,omitempty"` // The key of the permission.
}

// IssueNotifyUserScheme represents a user for notifying about an issue in Jira.
type IssueNotifyUserScheme struct {
	AccountID string `json:"accountId,omitempty"` // The account ID of the user.
}

// IssueNotifyGroupScheme represents a group for notifying about an issue in Jira.
type IssueNotifyGroupScheme struct {
	Name string `json:"name,omitempty"` // The name of the group.
}

// IssueBulkSchemeV3 represents a bulk operation on version 3 issues in Jira.
type IssueBulkSchemeV3 struct {
	Payload      *IssueScheme  // The payload for the bulk operation.
	CustomFields *CustomFields // The custom fields for the bulk operation.
}

// BulkIssueSchemeV3 represents a bulk of version 3 issues in Jira.
type BulkIssueSchemeV3 struct {
	Issues []*IssueScheme `json:"issues,omitempty"` // The issues in the bulk.
}

// IssueMoveOptionsV3 represents the options for moving a version 3 issue in Jira.
type IssueMoveOptionsV3 struct {
	Fields       *IssueScheme      // The fields for the move operation.
	CustomFields *CustomFields     // The custom fields for the move operation.
	Operations   *UpdateOperations // The operations for the move operation.
}
