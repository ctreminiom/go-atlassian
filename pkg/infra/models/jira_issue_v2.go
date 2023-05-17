package models

import (
	"encoding/json"
	"github.com/imdario/mergo"
)

type IssueSchemeV2 struct {
	ID          string                   `json:"id,omitempty"`
	Key         string                   `json:"key,omitempty"`
	Self        string                   `json:"self,omitempty"`
	Transitions []*IssueTransitionScheme `json:"transitions,omitempty"`
	Changelog   *IssueChangelogScheme    `json:"changelog,omitempty"`
	Fields      *IssueFieldsSchemeV2     `json:"fields,omitempty"`
}

func (i *IssueSchemeV2) MergeCustomFields(fields *CustomFields) (map[string]interface{}, error) {

	if fields == nil || len(fields.Fields) == 0 {
		return nil, ErrNoCustomFieldError
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

func (i *IssueSchemeV2) MergeOperations(operations *UpdateOperations) (map[string]interface{}, error) {

	if operations == nil || len(operations.Fields) == 0 {
		return nil, ErrNoOperatorError
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

func (i *IssueSchemeV2) ToMap() (map[string]interface{}, error) {

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
	Resolutiondate           string                          `json:"resolutiondate,omitempty"`
	Workratio                int                             `json:"workratio,omitempty"`
	StatusCategoryChangeDate string                          `json:"statuscategorychangedate,omitempty"`
	LastViewed               string                          `json:"lastViewed,omitempty"`
	Summary                  string                          `json:"summary,omitempty"`
	Created                  string                          `json:"created,omitempty"`
	Updated                  string                          `json:"updated,omitempty"`
	Labels                   []string                        `json:"labels,omitempty"`
	Status                   *StatusScheme                   `json:"status,omitempty"`
	Description              string                          `json:"description,omitempty"`
	Comment                  *IssueCommentPageSchemeV2       `json:"comment,omitempty"`
	Subtasks                 []*IssueScheme                  `json:"subtasks,omitempty"`
	Security                 *SecurityScheme                 `json:"security,omitempty"`
	Worklog                  *IssueWorklogRichTextPageScheme `json:"worklog,omitempty"`
}

type ParentScheme struct {
	ID     string              `json:"id,omitempty"`
	Key    string              `json:"key,omitempty"`
	Self   string              `json:"self,omitempty"`
	Fields *ParentFieldsScheme `json:"fields,omitempty"`
}

type ParentFieldsScheme struct {
	Summary string        `json:"summary,omitempty"`
	Status  *StatusScheme `json:"status,omitempty"`
}

type IssueResponseScheme struct {
	ID   string `json:"id,omitempty"`
	Key  string `json:"key,omitempty"`
	Self string `json:"self,omitempty"`
}

type IssueBulkSchemeV2 struct {
	Payload      *IssueSchemeV2
	CustomFields *CustomFields
}

type BulkIssueSchemeV2 struct {
	Issues []*IssueSchemeV2 `json:"issues,omitempty"`
}

type IssueBulkResponseScheme struct {
	Issues []struct {
		ID   string `json:"id,omitempty"`
		Key  string `json:"key,omitempty"`
		Self string `json:"self,omitempty"`
	} `json:"issues,omitempty"`
	Errors []*IssueBulkResponseErrorScheme `json:"errors,omitempty"`
}

type IssueBulkResponseErrorScheme struct {
	Status        int `json:"status"`
	ElementErrors struct {
		ErrorMessages []string `json:"errorMessages"`
		Status        int      `json:"status"`
	} `json:"elementErrors"`
	FailedElementNumber int `json:"failedElementNumber"`
}

type IssueMoveOptionsV2 struct {
	Fields       *IssueSchemeV2
	CustomFields *CustomFields
	Operations   *UpdateOperations
}
