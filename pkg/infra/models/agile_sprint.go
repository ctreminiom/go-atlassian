// Package models provides the data structures used in the agile sprint management.
package models

import "time"

// SprintScheme represents an agile sprint.
// ID is the unique identifier of the sprint.
// Self is the self URL of the sprint.
// State is the state of the sprint.
// Name is the name of the sprint.
// StartDate is the start date of the sprint.
// EndDate is the end date of the sprint.
// CompleteDate is the completion date of the sprint.
// OriginBoardID is the ID of the board where the sprint originated.
// Goal is the goal of the sprint.
type SprintScheme struct {
	ID            int       `json:"id,omitempty"`
	Self          string    `json:"self,omitempty"`
	State         string    `json:"state,omitempty"`
	Name          string    `json:"name,omitempty"`
	StartDate     time.Time `json:"startDate,omitempty"`
	EndDate       time.Time `json:"endDate,omitempty"`
	CompleteDate  time.Time `json:"completeDate,omitempty"`
	OriginBoardID int       `json:"originBoardId,omitempty"`
	Goal          string    `json:"goal,omitempty"`
}

// SprintPayloadScheme represents the payload for creating or updating a sprint.
// Name is the name of the sprint.
// StartDate is the start date of the sprint.
// EndDate is the end date of the sprint.
// OriginBoardID is the ID of the board where the sprint originated.
// Goal is the goal of the sprint.
// State is the state of the sprint.
type SprintPayloadScheme struct {
	Name          string `json:"name,omitempty"`
	StartDate     string `json:"startDate,omitempty"`
	EndDate       string `json:"endDate,omitempty"`
	OriginBoardID int    `json:"originBoardId,omitempty"`
	Goal          string `json:"goal,omitempty"`
	State         string `json:"state,omitempty"`
}

// SprintIssuePageScheme represents a page of issues in a sprint.
// Expand is a string that contains the instructions for expanding the issues in the page.
// StartAt is the starting index of the page.
// MaxResults is the maximum number of results per page.
// Total is the total number of issues.
// Issues is a slice of the issues in the page.
type SprintIssuePageScheme struct {
	Expand     string               `json:"expand,omitempty"`
	StartAt    int                  `json:"startAt,omitempty"`
	MaxResults int                  `json:"maxResults,omitempty"`
	Total      int                  `json:"total,omitempty"`
	Issues     []*SprintIssueScheme `json:"issues,omitempty"`
}

// SprintIssueScheme represents an issue in a sprint.
// Expand is a string that contains the instructions for expanding the issue.
// ID is the unique identifier of the issue.
// Self is the self URL of the issue.
// Key is the key of the issue.
type SprintIssueScheme struct {
	Expand string `json:"expand,omitempty"`
	ID     string `json:"id,omitempty"`
	Self   string `json:"self,omitempty"`
	Key    string `json:"key,omitempty"`
}

// SprintMovePayloadScheme represents the payload for moving an issue in a sprint.
// Issues is a slice of the issues to be moved.
// RankBeforeIssue is the rank of the issue before the move.
// RankAfterIssue is the rank of the issue after the move.
// RankCustomFieldID is the ID of the custom field used for ranking.
type SprintMovePayloadScheme struct {
	Issues            []string `json:"issues,omitempty"`
	RankBeforeIssue   string   `json:"rankBeforeIssue,omitempty"`
	RankAfterIssue    string   `json:"rankAfterIssue,omitempty"`
	RankCustomFieldID int      `json:"rankCustomFieldId,omitempty"`
}

// SprintDetailScheme represents the details of a sprint.
// ID is the unique identifier of the sprint.
// State is the state of the sprint.
// Name is the name of the sprint.
// StartDate is the start date of the sprint.
// EndDate is the end date of the sprint.
// CompleteDate is the completion date of the sprint.
// OriginBoardID is the ID of the board where the sprint originated.
// Goal is the goal of the sprint.
// BoardID is the ID of the board where the sprint is located.
type SprintDetailScheme struct {
	ID            int    `json:"id,omitempty"`
	State         string `json:"state,omitempty"`
	Name          string `json:"name,omitempty"`
	StartDate     string `json:"startDate,omitempty"`
	EndDate       string `json:"endDate,omitempty"`
	CompleteDate  string `json:"completeDate,omitempty"`
	OriginBoardID int    `json:"originBoardId,omitempty"`
	Goal          string `json:"goal,omitempty"`
	BoardID       int    `json:"boardId,omitempty"`
}
