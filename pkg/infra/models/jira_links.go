package models

// IssueLinkScheme represents a link of an issue in Jira.
type IssueLinkScheme struct {
	ID           string             `json:"id,omitempty"`           // The ID of the link.
	Type         *LinkTypeScheme    `json:"type,omitempty"`         // The type of the link.
	InwardIssue  *LinkedIssueScheme `json:"inwardIssue,omitempty"`  // The inward issue of the link.
	OutwardIssue *LinkedIssueScheme `json:"outwardIssue,omitempty"` // The outward issue of the link.
}

// LinkTypeScheme represents the type of a link in Jira.
type LinkTypeScheme struct {
	Self    string `json:"self,omitempty"`    // The URL of the link type.
	ID      string `json:"id,omitempty"`      // The ID of the link type.
	Name    string `json:"name,omitempty"`    // The name of the link type.
	Inward  string `json:"inward,omitempty"`  // The inward description of the link type.
	Outward string `json:"outward,omitempty"` // The outward description of the link type.
}

// LinkedIssueScheme represents a linked issue in Jira.
type LinkedIssueScheme struct {
	ID     string                 `json:"id,omitempty"`     // The ID of the linked issue.
	Key    string                 `json:"key,omitempty"`    // The key of the linked issue.
	Self   string                 `json:"self,omitempty"`   // The URL of the linked issue.
	Fields *IssueLinkFieldsScheme `json:"fields,omitempty"` // The fields of the linked issue.
}

// IssueLinkFieldsScheme represents the fields of a linked issue in Jira.
type IssueLinkFieldsScheme struct {
	IssueType                *IssueTypeScheme    `json:"issuetype,omitempty"`                // The type of the linked issue.
	IssueLinks               []*IssueLinkScheme  `json:"issuelinks,omitempty"`               // The links of the linked issue.
	Watcher                  *IssueWatcherScheme `json:"watches,omitempty"`                  // The watchers of the linked issue.
	Votes                    *IssueVoteScheme    `json:"votes,omitempty"`                    // The votes for the linked issue.
	Versions                 []*VersionScheme    `json:"versions,omitempty"`                 // The versions of the linked issue.
	Project                  *ProjectScheme      `json:"project,omitempty"`                  // The project of the linked issue.
	FixVersions              []*VersionScheme    `json:"fixVersions,omitempty"`              // The fix versions for the linked issue.
	Priority                 *PriorityScheme     `json:"priority,omitempty"`                 // The priority of the linked issue.
	Components               []*ComponentScheme  `json:"components,omitempty"`               // The components of the linked issue.
	Creator                  *UserScheme         `json:"creator,omitempty"`                  // The creator of the linked issue.
	Reporter                 *UserScheme         `json:"reporter,omitempty"`                 // The reporter of the linked issue.
	Assignee                 *UserScheme         `json:"assignee,omitempty"`                 // The assignee of the linked issue.
	Resolution               *ResolutionScheme   `json:"resolution,omitempty"`               // The resolution of the linked issue.
	ResolutionDate           string              `json:"resolutiondate,omitempty"`           // The date the linked issue was resolved.
	Workratio                int                 `json:"workratio,omitempty"`                // The work ratio of the linked issue.
	StatusCategoryChangeDate string              `json:"statuscategorychangedate,omitempty"` // The date the status category of the linked issue changed.
	LastViewed               string              `json:"lastViewed,omitempty"`               // The last time the linked issue was viewed.
	Summary                  string              `json:"summary,omitempty"`                  // The summary of the linked issue.
	Created                  string              `json:"created,omitempty"`                  // The date the linked issue was created.
	Updated                  string              `json:"updated,omitempty"`                  // The date the linked issue was last updated.
	Labels                   []string            `json:"labels,omitempty"`                   // The labels of the linked issue.
	Status                   *StatusScheme       `json:"status,omitempty"`                   // The status of the linked issue.
	Security                 *SecurityScheme     `json:"security,omitempty"`                 // The security level of the linked issue.
}

// LinkPayloadSchemeV3 represents the payload for a version 3 link in Jira.
type LinkPayloadSchemeV3 struct {
	Comment      *CommentPayloadScheme `json:"comment,omitempty"`      // The comment for the link.
	InwardIssue  *LinkedIssueScheme    `json:"inwardIssue,omitempty"`  // The inward issue for the link.
	OutwardIssue *LinkedIssueScheme    `json:"outwardIssue,omitempty"` // The outward issue for the link.
	Type         *LinkTypeScheme       `json:"type,omitempty"`         // The type of the link.
}

// LinkPayloadSchemeV2 represents the payload for a version 2 link in Jira.
type LinkPayloadSchemeV2 struct {
	Comment      *CommentPayloadSchemeV2 `json:"comment,omitempty"`      // The comment for the link.
	InwardIssue  *LinkedIssueScheme      `json:"inwardIssue,omitempty"`  // The inward issue for the link.
	OutwardIssue *LinkedIssueScheme      `json:"outwardIssue,omitempty"` // The outward issue for the link.
	Type         *LinkTypeScheme         `json:"type,omitempty"`         // The type of the link.
}

// IssueLinkPageScheme represents a page of links in Jira.
type IssueLinkPageScheme struct {
	Expand string                `json:"expand,omitempty"` // The expand option for the page.
	ID     string                `json:"id,omitempty"`     // The ID of the page.
	Self   string                `json:"self,omitempty"`   // The URL of the page.
	Key    string                `json:"key,omitempty"`    // The key of the page.
	Fields *IssueLinkFieldScheme `json:"fields,omitempty"` // The fields of the page.
}

// IssueLinkFieldScheme represents the fields of a page of links in Jira.
type IssueLinkFieldScheme struct {
	IssueLinks []*IssueLinkScheme `json:"issuelinks,omitempty"` // The links in the page.
}

// IssueLinkTypeSearchScheme represents a search for link types in Jira.
type IssueLinkTypeSearchScheme struct {
	IssueLinkTypes []*LinkTypeScheme `json:"issueLinkTypes,omitempty"` // The link types in the search.
}

// IssueLinkTypeScheme represents a link type in Jira.
type IssueLinkTypeScheme struct {
	ID      string `json:"id,omitempty"`      // The ID of the link type.
	Name    string `json:"name,omitempty"`    // The name of the link type.
	Inward  string `json:"inward,omitempty"`  // The inward description of the link type.
	Outward string `json:"outward,omitempty"` // The outward description of the link type.
	Self    string `json:"self,omitempty"`    // The URL of the link type.
}
