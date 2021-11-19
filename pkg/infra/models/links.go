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
	ID     string             `json:"id,omitempty"`
	Key    string             `json:"key,omitempty"`
	Self   string             `json:"self,omitempty"`
	Fields *IssueFieldsScheme `json:"fields,omitempty"`
}

type LinkPayloadScheme struct {
	Comment      *CommentPayloadScheme `json:"comment,omitempty"`
	InwardIssue  *LinkedIssueScheme    `json:"inwardIssue,omitempty"`
	OutwardIssue *LinkedIssueScheme    `json:"outwardIssue,omitempty"`
	Type         *LinkTypeScheme       `json:"type,omitempty"`
}

type IssueLinkPageScheme struct {
	Expand string `json:"expand,omitempty"`
	ID     string `json:"id,omitempty"`
	Self   string `json:"self,omitempty"`
	Key    string `json:"key,omitempty"`
	Fields struct {
		IssueLinks []*IssueLinkScheme `json:"issuelinks,omitempty"`
	} `json:"fields,omitempty"`
}

type IssueLinkTypeSearchScheme struct {
	IssueLinkTypes []*LinkTypeScheme `json:"issueLinkTypes,omitempty"`
}

type IssueLinkTypeScheme struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Inward  string `json:"inward"`
	Outward string `json:"outward"`
	Self    string `json:"self"`
}
