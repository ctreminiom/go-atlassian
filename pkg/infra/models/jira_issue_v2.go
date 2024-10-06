package models

import (
	"encoding/json"

	"dario.cat/mergo"
)

// IssueSchemeV2 represents the scheme of an issue in Jira version 2.
type IssueSchemeV2 struct {
	ID             string                   `json:"id,omitempty"`          // The ID of the issue.
	Key            string                   `json:"key,omitempty"`         // The key of the issue.
	Self           string                   `json:"self,omitempty"`        // The URL of the issue.
	Transitions    []*IssueTransitionScheme `json:"transitions,omitempty"` // The transitions of the issue.
	Changelog      *IssueChangelogScheme    `json:"changelog,omitempty"`   // The changelog of the issue.
	Fields         *IssueFieldsSchemeV2     `json:"fields,omitempty"`      // The fields of the issue.
	RenderedFields map[string]interface{}   `json:"renderedFields,omitempty"`
}

// MergeCustomFields merges custom fields into the issue scheme.
// It returns a map representation of the issue scheme with the merged fields.
// If the provided fields are nil or empty, it returns an error.
func (i *IssueSchemeV2) MergeCustomFields(fields *CustomFields) (map[string]interface{}, error) {

	if fields == nil || len(fields.Fields) == 0 {
		return map[string]interface{}{}, nil
	}

	// Convert the IssueScheme struct to map[string]interface{}
	issueSchemeAsBytes, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	issueSchemeAsMap := make(map[string]interface{})
	if err := json.Unmarshal(issueSchemeAsBytes, &issueSchemeAsMap); err != nil {
		return nil, err
	}

	// For each customField created, merge it into the eAsMap
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
func (i *IssueSchemeV2) MergeOperations(operations *UpdateOperations) (map[string]interface{}, error) {

	if operations == nil || len(operations.Fields) == 0 {
		return map[string]interface{}{}, nil
	}

	// Convert the IssueScheme struct to map[string]interface{}
	issueSchemeAsBytes, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}

	issueSchemeAsMap := make(map[string]interface{})
	if err := json.Unmarshal(issueSchemeAsBytes, &issueSchemeAsMap); err != nil {
		return nil, err
	}

	// For each customField created, merge it into the eAsMap
	for _, customField := range operations.Fields {
		if err := mergo.Merge(&issueSchemeAsMap, customField, mergo.WithOverride); err != nil {
			return nil, err
		}
	}

	return issueSchemeAsMap, nil
}

// ToMap converts the issue scheme to a map representation.
// It returns a map[string]interface{} where the keys are the field names and the values are the field values.
func (i *IssueSchemeV2) ToMap() (map[string]interface{}, error) {

	// Convert the IssueScheme struct to map[string]interface{}
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

// IssueFieldsSchemeV2 represents the fields of an issue in Jira version 2.
type IssueFieldsSchemeV2 struct {
	Parent                   *ParentScheme                   `json:"parent,omitempty"`
	IssueType                *IssueTypeScheme                `json:"issuetype,omitempty"`
	IssueLinks               []*IssueLinkScheme              `json:"issuelinks,omitempty"`
	Watcher                  *IssueWatcherScheme             `json:"watches,omitempty"`
	Votes                    *IssueVoteScheme                `json:"votes,omitempty"`
	Versions                 []*VersionScheme                `json:"versions,omitempty"`
	Project                  *ProjectScheme                  `json:"project,omitempty"`
	FixVersions              []*VersionScheme                `json:"fixVersions,omitempty"`
	Priority                 *PriorityScheme                 `json:"priority,omitempty"`
	Components               []*ComponentScheme              `json:"components,omitempty"`
	Creator                  *UserScheme                     `json:"creator,omitempty"`
	Reporter                 *UserScheme                     `json:"reporter,omitempty"`
	Assignee                 *UserScheme                     `json:"assignee,omitempty"`
	Resolution               *ResolutionScheme               `json:"resolution,omitempty"`
	ResolutionDate           *DateTimeScheme                 `json:"resolutiondate,omitempty"`
	Workratio                int                             `json:"workratio,omitempty"`
	StatusCategoryChangeDate *DateTimeScheme                 `json:"statuscategorychangedate,omitempty"`
	LastViewed               string                          `json:"lastViewed,omitempty"`
	Summary                  string                          `json:"summary,omitempty"`
	Created                  *DateTimeScheme                 `json:"created,omitempty"`
	Updated                  *DateTimeScheme                 `json:"updated,omitempty"`
	Labels                   []string                        `json:"labels,omitempty"`
	Status                   *StatusScheme                   `json:"status,omitempty"`
	Description              string                          `json:"description,omitempty"`
	Comment                  *IssueCommentPageSchemeV2       `json:"comment,omitempty"`
	Subtasks                 []*IssueScheme                  `json:"subtasks,omitempty"`
	Security                 *SecurityScheme                 `json:"security,omitempty"`
	Worklog                  *IssueWorklogRichTextPageScheme `json:"worklog,omitempty"`
	DueDate                  *DateScheme                     `json:"duedate,omitempty"`
}

// ParentScheme represents the parent of an issue in Jira.
type ParentScheme struct {
	ID     string              `json:"id,omitempty"`     // The ID of the parent issue.
	Key    string              `json:"key,omitempty"`    // The key of the parent issue.
	Self   string              `json:"self,omitempty"`   // The URL of the parent issue.
	Fields *ParentFieldsScheme `json:"fields,omitempty"` // The fields of the parent issue.
}

// ParentFieldsScheme represents the fields of a parent issue in Jira.
type ParentFieldsScheme struct {
	Summary string        `json:"summary,omitempty"` // The summary of the parent issue.
	Status  *StatusScheme `json:"status,omitempty"`  // The status of the parent issue.
}

// IssueResponseScheme represents the response of an issue operation in Jira.
type IssueResponseScheme struct {
	ID   string `json:"id,omitempty"`   // The ID of the issue.
	Key  string `json:"key,omitempty"`  // The key of the issue.
	Self string `json:"self,omitempty"` // The URL of the issue.
}

// IssueBulkSchemeV2 represents the bulk operation scheme for issues in Jira.
type IssueBulkSchemeV2 struct {
	Payload      *IssueSchemeV2 // The payload of the bulk operation.
	CustomFields *CustomFields  // The custom fields of the bulk operation.
}

// BulkIssueSchemeV2 represents the bulk issue scheme in Jira.
type BulkIssueSchemeV2 struct {
	Issues []*IssueSchemeV2 `json:"issues,omitempty"` // The issues in the bulk operation.
}

// IssueBulkResponseScheme represents the response of a bulk issue operation in Jira.
type IssueBulkResponseScheme struct {
	Issues []struct {
		ID   string `json:"id,omitempty"`   // The ID of the issue.
		Key  string `json:"key,omitempty"`  // The key of the issue.
		Self string `json:"self,omitempty"` // The URL of the issue.
	} `json:"issues,omitempty"` // The issues in the response.
	Errors []*IssueBulkResponseErrorScheme `json:"errors,omitempty"` // The errors in the response.
}

// IssueBulkResponseErrorScheme represents the error scheme of a bulk issue operation in Jira.
type IssueBulkResponseErrorScheme struct {
	Status        int `json:"status"` // The status of the error.
	ElementErrors struct {
		ErrorMessages []string `json:"errorMessages"` // The error messages.
		Status        int      `json:"status"`        // The status of the error messages.
	} `json:"elementErrors"` // The element errors in the response.
	FailedElementNumber int `json:"failedElementNumber"` // The number of the failed element.
}

// IssueMoveOptionsV2 represents the move options for an issue in Jira.
type IssueMoveOptionsV2 struct {
	Fields       *IssueSchemeV2    // The fields of the issue.
	CustomFields *CustomFields     // The custom fields of the issue.
	Operations   *UpdateOperations // The operations for the issue.
}
