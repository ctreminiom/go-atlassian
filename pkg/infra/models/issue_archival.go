package models

type IssueArchivalSyncResponseScheme struct {
	Errors                *IssueArchivalSyncErrorScheme `json:"errors"`
	NumberOfIssuesUpdated int                           `json:"numberOfIssuesUpdated"`
}

type IssueArchivalSyncErrorScheme struct {
	IssueIsSubtask             *IssueArchivalErrorScheme `json:"issueIsSubtask"`
	IssuesInArchivedProjects   *IssueArchivalErrorScheme `json:"issuesInArchivedProjects"`
	IssuesInUnlicensedProjects *IssueArchivalErrorScheme `json:"issuesInUnlicensedProjects"`
	IssuesNotFound             *IssueArchivalErrorScheme `json:"issuesNotFound"`
	UserDoesNotHavePermission  *IssueArchivalErrorScheme `json:"userDoesNotHavePermission"`
}

type IssueArchivalErrorScheme struct {
	Count          int      `json:"count"`
	IssueIdsOrKeys []string `json:"issueIdsOrKeys"`
	Message        string   `json:"message"`
}

type IssueArchivalExportPayloadScheme struct {
	ArchivedBy        []string                      `json:"archivedBy"`
	ArchivedDateRange *DateRangeFilterRequestScheme `json:"archivedDateRange,omitempty"`
	IssueTypes        []string                      `json:"issueTypes"`
	Projects          []string                      `json:"projects"`
	Reporters         []string                      `json:"reporters"`
}

type DateRangeFilterRequestScheme struct {
	DateAfter  string `json:"dateAfter,omitempty"`
	DateBefore string `json:"dateBefore,omitempty"`
}
