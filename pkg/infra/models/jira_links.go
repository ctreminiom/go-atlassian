package models

type IssueLinkScheme struct {
	ID           string             `json:"id,omitempty"`
	Type         *LinkTypeScheme    `json:"type,omitempty"`
	InwardIssue  *LinkedIssueScheme `json:"inwardIssue,omitempty"`
	OutwardIssue *LinkedIssueScheme `json:"outwardIssue,omitempty"`
}

type LinkTypeScheme struct {
	Self    string `json:"self,omitempty"`
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Inward  string `json:"inward,omitempty"`
	Outward string `json:"outward,omitempty"`
}

type LinkedIssueScheme struct {
	ID     string                 `json:"id,omitempty"`
	Key    string                 `json:"key,omitempty"`
	Self   string                 `json:"self,omitempty"`
	Fields *IssueLinkFieldsScheme `json:"fields,omitempty"`
}

type IssueLinkFieldsScheme struct {
	IssueType                *IssueTypeScheme    `json:"issuetype,omitempty"`
	IssueLinks               []*IssueLinkScheme  `json:"issuelinks,omitempty"`
	Watcher                  *IssueWatcherScheme `json:"watches,omitempty"`
	Votes                    *IssueVoteScheme    `json:"votes,omitempty"`
	Versions                 []*VersionScheme    `json:"versions,omitempty"`
	Project                  *ProjectScheme      `json:"project,omitempty"`
	FixVersions              []*VersionScheme    `json:"fixVersions,omitempty"`
	Priority                 *PriorityScheme     `json:"priority,omitempty"`
	Components               []*ComponentScheme  `json:"components,omitempty"`
	Creator                  *UserScheme         `json:"creator,omitempty"`
	Reporter                 *UserScheme         `json:"reporter,omitempty"`
	Assignee                 *UserScheme         `json:"assignee,omitempty"`
	Resolution               *ResolutionScheme   `json:"resolution,omitempty"`
	Resolutiondate           string              `json:"resolutiondate,omitempty"`
	Workratio                int                 `json:"workratio,omitempty"`
	StatusCategoryChangeDate string              `json:"statuscategorychangedate,omitempty"`
	LastViewed               string              `json:"lastViewed,omitempty"`
	Summary                  string              `json:"summary,omitempty"`
	Created                  string              `json:"created,omitempty"`
	Updated                  string              `json:"updated,omitempty"`
	Labels                   []string            `json:"labels,omitempty"`
	Status                   *StatusScheme       `json:"status,omitempty"`
	Security                 *SecurityScheme     `json:"security,omitempty"`
}

type LinkPayloadSchemeV3 struct {
	Comment      *CommentPayloadScheme `json:"comment,omitempty"`
	InwardIssue  *LinkedIssueScheme    `json:"inwardIssue,omitempty"`
	OutwardIssue *LinkedIssueScheme    `json:"outwardIssue,omitempty"`
	Type         *LinkTypeScheme       `json:"type,omitempty"`
}

type LinkPayloadSchemeV2 struct {
	Comment      *CommentPayloadSchemeV2 `json:"comment,omitempty"`
	InwardIssue  *LinkedIssueScheme      `json:"inwardIssue,omitempty"`
	OutwardIssue *LinkedIssueScheme      `json:"outwardIssue,omitempty"`
	Type         *LinkTypeScheme         `json:"type,omitempty"`
}

type IssueLinkPageScheme struct {
	Expand string                `json:"expand,omitempty"`
	ID     string                `json:"id,omitempty"`
	Self   string                `json:"self,omitempty"`
	Key    string                `json:"key,omitempty"`
	Fields *IssueLinkFieldScheme `json:"fields,omitempty"`
}

type IssueLinkFieldScheme struct {
	IssueLinks []*IssueLinkScheme `json:"issuelinks,omitempty"`
}

type IssueLinkTypeSearchScheme struct {
	IssueLinkTypes []*LinkTypeScheme `json:"issueLinkTypes,omitempty"`
}

type IssueLinkTypeScheme struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Inward  string `json:"inward,omitempty"`
	Outward string `json:"outward,omitempty"`
	Self    string `json:"self,omitempty"`
}
