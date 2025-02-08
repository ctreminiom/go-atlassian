// Package models provides the models for the issue archival service.
package models

// IssueArchivalSyncResponseScheme represents the response from the issue archival synchronization operation.
type IssueArchivalSyncResponseScheme struct {
	Errors                *IssueArchivalSyncErrorScheme `json:"errors"`
	NumberOfIssuesUpdated int                           `json:"numberOfIssuesUpdated"`
}

// IssueArchivalSyncErrorScheme represents the error details for the issue archival synchronization operation.
type IssueArchivalSyncErrorScheme struct {
	IssueIsSubtask             *IssueArchivalErrorScheme `json:"issueIsSubtask"`
	IssuesInArchivedProjects   *IssueArchivalErrorScheme `json:"issuesInArchivedProjects"`
	IssuesInUnlicensedProjects *IssueArchivalErrorScheme `json:"issuesInUnlicensedProjects"`
	IssuesNotFound             *IssueArchivalErrorScheme `json:"issuesNotFound"`
	UserDoesNotHavePermission  *IssueArchivalErrorScheme `json:"userDoesNotHavePermission"`
}

// IssueArchivalErrorScheme represents the error details for the issue archival operation.
type IssueArchivalErrorScheme struct {
	Count          int      `json:"count"`
	IssueIDsOrKeys []string `json:"issueIdsOrKeys"`
	Message        string   `json:"message"`
}

// IssueArchivalExportPayloadScheme represents the payload for the issue archival export operation.
type IssueArchivalExportPayloadScheme struct {
	ArchivedBy        []string                      `json:"archivedBy"`
	ArchivedDateRange *DateRangeFilterRequestScheme `json:"archivedDateRange,omitempty"`
	IssueTypes        []string                      `json:"issueTypes"`
	Projects          []string                      `json:"projects"`
	Reporters         []string                      `json:"reporters"`
}

// DateRangeFilterRequestScheme represents the date range filter for the issue archival export operation.
type DateRangeFilterRequestScheme struct {
	DateAfter  string `json:"dateAfter,omitempty"`
	DateBefore string `json:"dateBefore,omitempty"`
}

// IssueArchiveExportResultScheme represents the result of an issue archival export operation.
type IssueArchiveExportResultScheme struct {
	TaskID        string `json:"taskId,omitempty"`
	Payload       string `json:"payload,omitempty"`
	Progress      int    `json:"progress,omitempty"`
	SubmittedTime int64  `json:"submittedTime,omitempty"`
	Status        string `json:"status,omitempty"`
}
