package jira

import (
	"encoding/json"
	"fmt"
	"github.com/imdario/mergo"
)

type IssueScheme struct {
	ID          string                   `json:"id,omitempty"`
	Key         string                   `json:"key,omitempty"`
	Self        string                   `json:"self,omitempty"`
	Transitions []*IssueTransitionScheme `json:"transitions,omitempty"`
	Changelog   *IssueChangelogScheme    `json:"changelog,omitempty"`
	Fields      *IssueFieldsScheme       `json:"fields,omitempty"`
}

func (i *IssueScheme) MergeCustomFields(fields *CustomFields) (result map[string]interface{}, err error) {

	if fields == nil {
		return nil, fmt.Errorf("error, please provide a value *CustomFields pointer")
	}

	if len(fields.Fields) == 0 {
		return nil, fmt.Errorf("error!, the Fields tag does not contains custom fields")
	}

	//Convert the IssueScheme struct to map[string]interface{}
	issueSchemeAsBytes, _ := json.Marshal(i)

	issueSchemeAsMap := make(map[string]interface{})
	_ = json.Unmarshal(issueSchemeAsBytes, &issueSchemeAsMap)

	//For each customField created, merge it into the eAsMap
	for _, customField := range fields.Fields {
		_ = mergo.Merge(&issueSchemeAsMap, customField, mergo.WithOverride)
	}

	return issueSchemeAsMap, nil
}

func (i *IssueScheme) MergeOperations(operations *UpdateOperations) (result map[string]interface{}, err error) {

	if operations == nil {
		return nil, fmt.Errorf("error, please provide a value *UpdateOperations pointer")
	}

	if len(operations.Fields) == 0 {
		return nil, fmt.Errorf("error!, the Fields tag does not contains custom fields")
	}

	//Convert the IssueScheme struct to map[string]interface{}
	issueSchemeAsBytes, _ := json.Marshal(i)

	issueSchemeAsMap := make(map[string]interface{})
	_ = json.Unmarshal(issueSchemeAsBytes, &issueSchemeAsMap)

	//For each customField created, merge it into the eAsMap
	for _, customField := range operations.Fields {
		_ = mergo.Merge(&issueSchemeAsMap, customField, mergo.WithOverride)
	}

	return issueSchemeAsMap, nil
}

func (i *IssueScheme) ToMap() (result map[string]interface{}, err error) {

	//Convert the IssueScheme struct to map[string]interface{}
	issueSchemeAsBytes, _ := json.Marshal(i)

	issueSchemeAsMap := make(map[string]interface{})
	_ = json.Unmarshal(issueSchemeAsBytes, &issueSchemeAsMap)

	return issueSchemeAsMap, err
}

type IssueFieldsScheme struct {
	IssueType                *IssueTypeScheme        `json:"issuetype,omitempty"`
	IssueLinks               []*IssueLinkScheme      `json:"issuelinks,omitempty"`
	Watcher                  *IssueWatcherScheme     `json:"watches,omitempty"`
	Votes                    *IssueVoteScheme        `json:"votes,omitempty"`
	Versions                 []*VersionScheme        `json:"versions,omitempty"`
	Project                  *ProjectScheme          `json:"project,omitempty"`
	FixVersions              []*VersionScheme        `json:"fixVersions,omitempty"`
	Priority                 *PriorityScheme         `json:"priority,omitempty"`
	Components               []*ComponentScheme      `json:"components,omitempty"`
	Creator                  *UserScheme             `json:"creator,omitempty"`
	Reporter                 *UserScheme             `json:"reporter,omitempty"`
	Resolution               *ResolutionScheme       `json:"resolution,omitempty"`
	Resolutiondate           string                  `json:"resolutiondate,omitempty"`
	Workratio                int                     `json:"workratio,omitempty"`
	StatusCategoryChangeDate string                  `json:"statuscategorychangedate,omitempty"`
	LastViewed               string                  `json:"lastViewed,omitempty"`
	Summary                  string                  `json:"summary,omitempty"`
	Created                  string                  `json:"created,omitempty"`
	Updated                  string                  `json:"updated,omitempty"`
	Labels                   []string                `json:"labels,omitempty"`
	Status                   *StatusScheme           `json:"status,omitempty"`
	Description              *CommentNodeScheme      `json:"description,omitempty"`
	Comment                  *IssueCommentPageScheme `json:"comment,omitempty"`
	Subtasks                 []*IssueScheme          `json:"subtasks,omitempty"`
}

type IssueTransitionScheme struct {
	ID            string        `json:"id,omitempty"`
	Name          string        `json:"name,omitempty"`
	To            *StatusScheme `json:"to,omitempty"`
	HasScreen     bool          `json:"hasScreen,omitempty"`
	IsGlobal      bool          `json:"isGlobal,omitempty"`
	IsInitial     bool          `json:"isInitial,omitempty"`
	IsAvailable   bool          `json:"isAvailable,omitempty"`
	IsConditional bool          `json:"isConditional,omitempty"`
	IsLooped      bool          `json:"isLooped,omitempty"`
}

type StatusScheme struct {
	Self           string                `json:"self,omitempty"`
	Description    string                `json:"description,omitempty"`
	IconURL        string                `json:"iconUrl,omitempty"`
	Name           string                `json:"name,omitempty"`
	ID             string                `json:"id,omitempty"`
	StatusCategory *StatusCategoryScheme `json:"statusCategory,omitempty"`
}

type StatusCategoryScheme struct {
	Self      string `json:"self,omitempty"`
	ID        int    `json:"id,omitempty"`
	Key       string `json:"key,omitempty"`
	ColorName string `json:"colorName,omitempty"`
	Name      string `json:"name,omitempty"`
}

type IssueNotifyOptionsScheme struct {
	HTMLBody string                     `json:"htmlBody,omitempty"`
	Subject  string                     `json:"subject,omitempty"`
	TextBody string                     `json:"textBody,omitempty"`
	To       *IssueNotifyToScheme       `json:"to,omitempty"`
	Restrict *IssueNotifyRestrictScheme `json:"restrict,omitempty"`
}

type IssueNotifyRestrictScheme struct {
	Groups      []*IssueNotifyGroupScheme      `json:"groups,omitempty"`
	Permissions []*IssueNotifyPermissionScheme `json:"permissions,omitempty"`
}

type IssueNotifyToScheme struct {
	Reporter bool                      `json:"reporter,omitempty"`
	Assignee bool                      `json:"assignee,omitempty"`
	Watchers bool                      `json:"watchers,omitempty"`
	Voters   bool                      `json:"voters,omitempty"`
	Users    []*IssueNotifyUserScheme  `json:"users,omitempty"`
	Groups   []*IssueNotifyGroupScheme `json:"groups,omitempty"`
}

type IssueNotifyPermissionScheme struct {
	ID  string `json:"id,omitempty"`
	Key string `json:"key,omitempty"`
}

type IssueNotifyUserScheme struct {
	AccountID string `json:"accountId,omitempty"`
}

type IssueNotifyGroupScheme struct {
	Name string `json:"name,omitempty"`
}
