package models

import "time"

type BoardPayloadScheme struct {
	Name     string                      `json:"name,omitempty"`
	Type     string                      `json:"type,omitempty"`
	FilterID int                         `json:"filterId,omitempty"`
	Location *BoardPayloadLocationScheme `json:"location,omitempty"`
}

type BoardPayloadLocationScheme struct {
	Type           string `json:"type,omitempty"`
	ProjectKeyOrID string `json:"projectKeyOrId,omitempty"`
}

type BoardPageScheme struct {
	MaxResults int            `json:"maxResults"`
	StartAt    int            `json:"startAt"`
	Total      int            `json:"total"`
	IsLast     bool           `json:"isLast"`
	Values     []*BoardScheme `json:"values"`
}

type BoardScheme struct {
	ID       int                  `json:"id,omitempty"`
	Self     string               `json:"self,omitempty"`
	Name     string               `json:"name,omitempty"`
	Type     string               `json:"type,omitempty"`
	Location *BoardLocationScheme `json:"location,omitempty"`
}

type BoardLocationScheme struct {
	ProjectID      int    `json:"projectId,omitempty"`
	DisplayName    string `json:"displayName,omitempty"`
	ProjectName    string `json:"projectName,omitempty"`
	ProjectKey     string `json:"projectKey,omitempty"`
	ProjectTypeKey string `json:"projectTypeKey,omitempty"`
	AvatarURI      string `json:"avatarURI,omitempty"`
	Name           string `json:"name,omitempty"`
}

type BoardIssuePageScheme struct {
	Expand     string           `json:"expand,omitempty"`
	StartAt    int              `json:"startAt,omitempty"`
	MaxResults int              `json:"maxResults,omitempty"`
	Total      int              `json:"total,omitempty"`
	Issues     []*IssueSchemeV2 `json:"issues,omitempty"`
}

type BoardConfigurationScheme struct {
	ID           int                             `json:"id,omitempty"`
	Name         string                          `json:"name,omitempty"`
	Type         string                          `json:"type,omitempty"`
	Self         string                          `json:"self,omitempty"`
	Location     *BoardLocationScheme            `json:"location,omitempty"`
	Filter       *BoardFilterScheme              `json:"filter,omitempty"`
	ColumnConfig *BoardColumnConfigurationScheme `json:"columnConfig,omitempty"`
	Estimation   *BoardEstimationScheme          `json:"estimation,omitempty"`
	Ranking      *BoardRankingScheme             `json:"ranking,omitempty"`
}

type BoardFilterScheme struct {
	ID   string `json:"id,omitempty"`
	Self string `json:"self,omitempty"`
}

type BoardRankingScheme struct {
	RankCustomFieldID int `json:"rankCustomFieldId,omitempty"`
}

type BoardEstimationScheme struct {
	Type  string                      `json:"type,omitempty"`
	Field *BoardEstimationFieldScheme `json:"field,omitempty"`
}

type BoardEstimationFieldScheme struct {
	FieldID     string `json:"fieldId,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type BoardColumnConfigurationScheme struct {
	Columns        []*BoardColumnScheme `json:"columns,omitempty"`
	ConstraintType string               `json:"constraintType,omitempty"`
}

type BoardColumnScheme struct {
	Name     string                     `json:"name,omitempty"`
	Statuses []*BoardColumnStatusScheme `json:"statuses,omitempty"`
}

type BoardColumnStatusScheme struct {
	ID   string `json:"id,omitempty"`
	Self string `json:"self,omitempty"`
}

type BoardEpicPageScheme struct {
	MaxResults int                `json:"maxResults,omitempty"`
	StartAt    int                `json:"startAt,omitempty"`
	IsLast     bool               `json:"isLast,omitempty"`
	Values     []*BoardEpicScheme `json:"values,omitempty"`
}

type BoardEpicScheme struct {
	ID      int    `json:"id,omitempty"`
	Key     string `json:"key,omitempty"`
	Self    string `json:"self,omitempty"`
	Name    string `json:"name,omitempty"`
	Summary string `json:"summary,omitempty"`
	Color   struct {
		Key string `json:"key,omitempty"`
	} `json:"color,omitempty"`
	Done bool `json:"done,omitempty"`
}

type BoardMovementPayloadScheme struct {
	Issues            []string `json:"issues,omitempty"`
	RankBeforeIssue   string   `json:"rankBeforeIssue,omitempty"`
	RankAfterIssue    string   `json:"rankAfterIssue,omitempty"`
	RankCustomFieldID int      `json:"rankCustomFieldId,omitempty"`
}

type BoardProjectPageScheme struct {
	MaxResults int                   `json:"maxResults,omitempty"`
	StartAt    int                   `json:"startAt,omitempty"`
	Total      int                   `json:"total,omitempty"`
	IsLast     bool                  `json:"isLast,omitempty"`
	Values     []*BoardProjectScheme `json:"values,omitempty"`
}

type BoardProjectScheme struct {
	Self            string                      `json:"self,omitempty"`
	ID              string                      `json:"id,omitempty"`
	Key             string                      `json:"key,omitempty"`
	Name            string                      `json:"name,omitempty"`
	ProjectCategory *BoardProjectCategoryScheme `json:"projectCategory,omitempty"`
	Simplified      bool                        `json:"simplified,omitempty"`
	Style           string                      `json:"style,omitempty"`
	Insight         *BoardProjectInsightScheme  `json:"insight,omitempty"`
}

type BoardProjectCategoryScheme struct {
	Self        string `json:"self,omitempty"`
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type BoardProjectInsightScheme struct {
	TotalIssueCount     int    `json:"totalIssueCount,omitempty"`
	LastIssueUpdateTime string `json:"lastIssueUpdateTime,omitempty"`
}

type BoardSprintPageScheme struct {
	MaxResults int                  `json:"maxResults,omitempty"`
	StartAt    int                  `json:"startAt,omitempty"`
	IsLast     bool                 `json:"isLast,omitempty"`
	Total      int                  `json:"total,omitempty"`
	Values     []*BoardSprintScheme `json:"values,omitempty"`
}

type BoardSprintScheme struct {
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

type BoardVersionPageScheme struct {
	MaxResults int                   `json:"maxResults,omitempty"`
	StartAt    int                   `json:"startAt,omitempty"`
	IsLast     bool                  `json:"isLast,omitempty"`
	Values     []*BoardVersionScheme `json:"values,omitempty"`
}

type BoardVersionScheme struct {
	Self        string    `json:"self,omitempty"`
	ID          int       `json:"id,omitempty"`
	ProjectID   int       `json:"projectId,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	Archived    bool      `json:"archived,omitempty"`
	Released    bool      `json:"released,omitempty"`
	ReleaseDate time.Time `json:"releaseDate,omitempty"`
}

type GetBoardsOptions struct {
	BoardType               string
	BoardName               string
	ProjectKeyOrID          string
	AccountIDLocation       string
	ProjectIDLocation       string
	IncludePrivate          bool
	NegateLocationFiltering bool
	OrderBy                 string
	Expand                  string
	FilterID                int
}

type IssueOptionScheme struct {
	JQL           string
	ValidateQuery bool
	Fields        []string
	Expand        []string
}

type BoardBacklogPayloadScheme struct {
	Issues            []string `json:"issues,omitempty"`
	RankBeforeIssue   string   `json:"rankBeforeIssue,omitempty"`
	RankAfterIssue    string   `json:"rankAfterIssue,omitempty"`
	RankCustomFieldID int      `json:"rankCustomFieldId,omitempty"`
}
