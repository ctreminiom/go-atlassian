package models

import "time"

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

type SprintPayloadScheme struct {
	Name          string `json:"name,omitempty"`
	StartDate     string `json:"startDate,omitempty"`
	EndDate       string `json:"endDate,omitempty"`
	OriginBoardID int    `json:"originBoardId,omitempty"`
	Goal          string `json:"goal,omitempty"`
	State         string `json:"state,omitempty"`
}

type SprintIssuePageScheme struct {
	Expand     string               `json:"expand,omitempty"`
	StartAt    int                  `json:"startAt,omitempty"`
	MaxResults int                  `json:"maxResults,omitempty"`
	Total      int                  `json:"total,omitempty"`
	Issues     []*SprintIssueScheme `json:"issues,omitempty"`
}

type SprintIssueScheme struct {
	Expand string `json:"expand,omitempty"`
	ID     string `json:"id,omitempty"`
	Self   string `json:"self,omitempty"`
	Key    string `json:"key,omitempty"`
}

type SprintMovePayloadScheme struct {
	Issues            []string `json:"issues,omitempty"`
	RankBeforeIssue   string   `json:"rankBeforeIssue,omitempty"`
	RankAfterIssue    string   `json:"rankAfterIssue,omitempty"`
	RankCustomFieldId int      `json:"rankCustomFieldId,omitempty"`
}

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
