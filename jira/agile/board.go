package agile

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type BoardService struct{ client *Client }

// Get returns the board for the given board ID.
// This board will only be returned if the user has permission to view it.
// Admins without the view permission will see the board as a private one,
// so will see only a subset of the board's data (board location for instance).
func (b *BoardService) Get(ctx context.Context, boardID int) (result *BoardScheme, response *ResponseScheme, err error) {

	if boardID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid boardID value")
	}

	var endpoint = fmt.Sprintf("/rest/agile/1.0/board/%v", boardID)

	request, err := b.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = b.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

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

// Create creates a new board. Board name, type and filter ID is required.
// Docs: N/A
func (b *BoardService) Create(ctx context.Context, payload *BoardPayloadScheme) (result *BoardScheme, response *ResponseScheme, err error) {

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint = "/rest/agile/1.0/board"

	request, err := b.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = b.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

type BoardPageScheme struct {
	MaxResults int            `json:"maxResults"`
	StartAt    int            `json:"startAt"`
	Total      int            `json:"total"`
	IsLast     bool           `json:"isLast"`
	Values     []*BoardScheme `json:"values"`
}

// Filter returns any boards which use the provided filter id.
// This method can be executed by users without a valid software license in order
// to find which boards are using a particular filter.
// Docs: N/A
func (b *BoardService) Filter(ctx context.Context, filterID, startAt, maxResults int) (result *BoardPageScheme, response *ResponseScheme, err error) {

	if filterID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid filterID value")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("/rest/agile/1.0/board/filter/%v?%v", filterID, params.Encode())

	request, err := b.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = b.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
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

// Backlog returns all issues from the board's backlog, for the given board ID.
// This only includes issues that the user has permission to view.
// The backlog contains incomplete issues that are not assigned to any future or active sprint.
// Note, if the user does not have permission to view the board, no issues will be returned at all.
// Issues returned from this resource include Agile fields, like sprint, closedSprints, flagged, and epic.
// By default, the returned issues are ordered by rank.
// Docs: N/A
func (b *BoardService) Backlog(ctx context.Context, boardID, startAt, maxResults int, opts *IssueOptionScheme) (result *BoardIssuePageScheme, response *ResponseScheme, err error) {

	if boardID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid boardID value")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if !opts.ValidateQuery {
			params.Add("validateQuery", "false")
		} else {
			params.Add("validateQuery", "true")
		}

		if len(opts.JQL) != 0 {
			params.Add("jql", opts.JQL)
		}

		if len(opts.Expand) != 0 {
			params.Add("expand", strings.Join(opts.Expand, ","))
		}

		if len(opts.Fields) != 0 {
			params.Add("fields", strings.Join(opts.Fields, ","))
		}

	}

	var endpoint = fmt.Sprintf("/rest/agile/1.0/board/%v/backlog?%v", boardID, params.Encode())

	request, err := b.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = b.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

type BoardIssuePageScheme struct {
	Expand     string `json:"expand"`
	StartAt    int    `json:"startAt"`
	MaxResults int    `json:"maxResults"`
	Total      int    `json:"total"`
	Issues     []struct {
		Expand string `json:"expand"`
		ID     string `json:"id"`
		Self   string `json:"self"`
		Key    string `json:"key"`
	} `json:"issues"`
}

// Configuration get the board configuration.
// Docs: N/A
func (b *BoardService) Configuration(ctx context.Context, boardID int) (result *BoardConfigurationScheme, response *ResponseScheme, err error) {

	if boardID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid boardID value")
	}

	var endpoint = fmt.Sprintf("/rest/agile/1.0/board/%v/configuration", boardID)

	request, err := b.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = b.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
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

// Epics returns all epics from the board, for the given board ID.
// This only includes epics that the user has permission to view.
// Note, if the user does not have permission to view the board, no epics will be returned at all.
func (b *BoardService) Epics(ctx context.Context, boardID, startAt, maxResults int, done bool) (result *BoardEpicPageScheme, response *ResponseScheme, err error) {

	if boardID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid boardID value")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if done {
		params.Add("done", "true")
	} else {
		params.Add("done", "false")
	}

	var endpoint = fmt.Sprintf("/rest/agile/1.0/board/%v/epic?%v", boardID, params.Encode())

	request, err := b.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = b.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
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

// IssuesWithoutEpic returns all issues that do not belong to any epic on a board, for a given board ID.
// This only includes issues that the user has permission to view.
// Issues returned from this resource include Agile fields, like sprint, closedSprints, flagged, and epic.
// By default, the returned issues are ordered by rank.
// Docs: N/A
func (b *BoardService) IssuesWithoutEpic(ctx context.Context, boardID, startAt, maxResults int, opts *IssueOptionScheme) (
	result *BoardIssuePageScheme, response *ResponseScheme, err error) {

	if boardID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid boardID value")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if !opts.ValidateQuery {
			params.Add("validateQuery ", "false")
		}

		if len(opts.JQL) != 0 {
			params.Add("jql", opts.JQL)
		}

		var expand string
		for index, value := range opts.Expand {

			if index == 0 {
				expand = value
				continue
			}

			expand += "," + value
		}

		if len(expand) != 0 {
			params.Add("expand", expand)
		}

		var fieldFormatted string
		for index, value := range opts.Fields {

			if index == 0 {
				fieldFormatted = value
				continue
			}
			fieldFormatted += "," + value
		}

		if len(fieldFormatted) != 0 {
			params.Add("fields", fieldFormatted)
		}

	}

	var endpoint = fmt.Sprintf("/rest/agile/1.0/board/%v/epic/none/issue?%v", boardID, params.Encode())

	request, err := b.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = b.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// IssuesByEpic returns all issues that belong to an epic on the board, for the given epic ID and the board ID.
// This only includes issues that the user has permission to view.
// Issues returned from this resource include Agile fields, like sprint, closedSprints,
// flagged, and epic. By default, the returned issues are ordered by rank.
func (b *BoardService) IssuesByEpic(ctx context.Context, boardID, epicID, startAt, maxResults int, opts *IssueOptionScheme) (
	result *BoardIssuePageScheme, response *ResponseScheme, err error) {

	if boardID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid boardID value")
	}

	if epicID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid epicID value")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if !opts.ValidateQuery {
			params.Add("validateQuery ", "false")
		}

		if len(opts.JQL) != 0 {
			params.Add("jql", opts.JQL)
		}

		if len(opts.Expand) != 0 {
			params.Add("expand", strings.Join(opts.Expand, ","))
		}

		if len(opts.Fields) != 0 {
			params.Add("fields", strings.Join(opts.Fields, ","))
		}

	}

	var endpoint = fmt.Sprintf("/rest/agile/1.0/board/%v/epic/%v/issue?%v", boardID, epicID, params.Encode())

	request, err := b.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = b.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Issues returns all issues from a board, for a given board ID.
// This only includes issues that the user has permission to view.
// An issue belongs to the board if its status is mapped to the board's column.
// Epic issues do not belongs to the scrum boards. Note, if the user does not have permission to view the board,
// no issues will be returned at all.
// Issues returned from this resource include Agile fields, like sprint, closedSprints, flagged, and epic.
// By default, the returned issues are ordered by rank.
// Docs: N/A
func (b *BoardService) Issues(ctx context.Context, boardID, startAt, maxResults int, opts *IssueOptionScheme) (
	result *BoardIssuePageScheme, response *ResponseScheme, err error) {

	if boardID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid boardID value")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if !opts.ValidateQuery {
			params.Add("validateQuery ", "false")
		}

		if len(opts.JQL) != 0 {
			params.Add("jql", opts.JQL)
		}

		if len(opts.Expand) != 0 {
			params.Add("expand", strings.Join(opts.Expand, ","))
		}

		if len(opts.Fields) != 0 {
			params.Add("fields", strings.Join(opts.Fields, ","))
		}

	}

	var endpoint = fmt.Sprintf("/rest/agile/1.0/board/%v/issue?%v", boardID, params.Encode())

	request, err := b.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = b.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

type BoardMovementPayloadScheme struct {
	Issues            []string `json:"issues,omitempty"`
	RankBeforeIssue   string   `json:"rankBeforeIssue,omitempty"`
	RankAfterIssue    string   `json:"rankAfterIssue,omitempty"`
	RankCustomFieldID int      `json:"rankCustomFieldId,omitempty"`
}

// Move issues from the backlog to the board (if they are already in the backlog of that board).
// This operation either moves an issue(s) onto a board from the backlog (by adding it to the issueList for the board)
// Or transitions the issue(s) to the first column for a kanban board with backlog.
// At most 50 issues may be moved at once.
// Docs: N/A
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/software/rest/api-group-board/#api-agile-1-0-board-boardid-issue-post
func (b *BoardService) Move(ctx context.Context, boardID int, payload *BoardMovementPayloadScheme) (response *ResponseScheme, err error) {

	if boardID == 0 {
		return nil, fmt.Errorf("error!, please provide a valid boardID value")
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	var endpoint = fmt.Sprintf("/rest/agile/1.0/board/%v/issue", boardID)

	request, err := b.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = b.client.Call(request, nil)
	if err != nil {
		return response, err
	}

	return
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

// Projects returns all projects that are associated with the board, for the given board ID.
// If the user does not have permission to view the board, no projects will be returned at all.
// Returned projects are ordered by the name.
// Docs: N/A
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/software/rest/api-group-board/#api-agile-1-0-board-boardid-project-get
func (b *BoardService) Projects(ctx context.Context, boardID, startAt, maxResults int) (
	result *BoardProjectPageScheme, response *ResponseScheme, err error) {

	if boardID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid boardID value")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("/rest/agile/1.0/board/%v/project?%v", boardID, params.Encode())

	request, err := b.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = b.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
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

// Sprints returns all sprints from a board, for a given board ID.
// This only includes sprints that the user has permission to view.
// Docs: N/A
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/software/rest/api-group-board/#api-agile-1-0-board-boardid-sprint-get
func (b *BoardService) Sprints(ctx context.Context, boardID, startAt, maxResults int, states []string) (
	result *BoardSprintPageScheme, response *ResponseScheme, err error) {

	if boardID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid boardID value")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))
	params.Add("state", strings.Join(states, ","))

	var endpoint = fmt.Sprintf("/rest/agile/1.0/board/%v/sprint?%v", boardID, params.Encode())

	request, err := b.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = b.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// IssuesBySprint get all issues you have access to that belong to the sprint from the board.
// Issue returned from this resource contains additional fields like: sprint, closedSprints, flagged and epic.
// Issues are returned ordered by rank. JQL order has higher priority than default rank.
func (b *BoardService) IssuesBySprint(ctx context.Context, boardID, sprintID, startAt, maxResults int,
	opts *IssueOptionScheme) (result *BoardIssuePageScheme, response *ResponseScheme, err error) {

	if boardID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid boardID value")
	}

	if sprintID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid sprintID value")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if !opts.ValidateQuery {
			params.Add("validateQuery ", "false")
		}

		if len(opts.JQL) != 0 {
			params.Add("jql", opts.JQL)
		}

		if len(opts.Expand) != 0 {
			params.Add("expand", strings.Join(opts.Expand, ","))
		}

		if len(opts.Fields) != 0 {
			params.Add("fields", strings.Join(opts.Fields, ","))
		}
	}

	var endpoint = fmt.Sprintf("/rest/agile/1.0/board/%v/sprint/%v/issue?%v", boardID, sprintID, params.Encode())

	request, err := b.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = b.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
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

// Versions returns all versions from a board, for a given board ID.
// This only includes versions that the user has permission to view.
// Note, if the user does not have permission to view the board, no versions will be returned at all.
// Returned versions are ordered by the name of the project from which they belong and then by sequence defined by user.
// Docs: N/A
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/software/rest/api-group-board/#api-agile-1-0-board-boardid-version-get
func (b *BoardService) Versions(ctx context.Context, boardID, startAt, maxResults int, released bool) (
	result *BoardVersionPageScheme, response *ResponseScheme, err error) {

	if boardID == 0 {
		return nil, nil, fmt.Errorf("error!, please provide a valid boardID value")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if released {
		params.Add("released", "true")
	} else {
		params.Add("released", "false")
	}

	var endpoint = fmt.Sprintf("/rest/agile/1.0/board/%v/version?%v", boardID, params.Encode())

	request, err := b.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = b.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}

// Delete deletes the board. Admin without the view permission can still remove the board.
// Docs: N/A
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/software/rest/api-group-board/#api-agile-1-0-board-boardid-delete
func (b *BoardService) Delete(ctx context.Context, boardID int) (response *ResponseScheme, err error) {

	if boardID == 0 {
		return nil, fmt.Errorf("error!, please provide a valid boardID value")
	}

	var endpoint = fmt.Sprintf("/rest/agile/1.0/board/%v", boardID)

	request, err := b.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = b.client.Call(request, nil)
	if err != nil {
		return
	}

	return
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

// Gets returns all boards. This only includes boards that the user has permission to view.
// Docs: https://developer.atlassian.com/cloud/jira/software/rest/api-group-other-operations/#api-agile-1-0-board-get
func (b *BoardService) Gets(ctx context.Context, opts *GetBoardsOptions, startAt, maxResults int) (result *BoardPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if opts.BoardType != "" {
			params.Add("type", opts.BoardType)
		}

		if opts.BoardName != "" {
			params.Add("name", opts.BoardName)
		}

		if opts.ProjectKeyOrID != "" {
			params.Add("projectKeyOrId", opts.ProjectKeyOrID)
		}

		if opts.AccountIDLocation != "" {
			params.Add("accountIdLocation", opts.AccountIDLocation)
		}

		if opts.ProjectIDLocation != "" {
			params.Add("projectLocation", opts.ProjectIDLocation)
		}

		if opts.IncludePrivate {
			params.Add("includePrivate", "true")
		}

		if opts.NegateLocationFiltering {
			params.Add("negateLocationFiltering", "true")
		}

		if opts.OrderBy != "" {
			params.Add("orderBy", opts.OrderBy)
		}

		if opts.Expand != "" {
			params.Add("expand", opts.Expand)
		}

		if opts.FilterID != 0 {
			params.Add("filterId", strconv.Itoa(opts.FilterID))
		}
	}

	var endpoint = fmt.Sprintf("/rest/agile/1.0/board?%v", params.Encode())

	request, err := b.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = b.client.Call(request, &result)
	if err != nil {
		return nil, response, err
	}

	return
}
