package v3

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type IssueSearchService struct{ client *Client }

type IssueSearchScheme struct {
	Expand          string                `json:"expand"`
	StartAt         int                   `json:"startAt"`
	MaxResults      int                   `json:"maxResults"`
	Total           int                   `json:"total"`
	Issues          []*models.IssueScheme `json:"issues"`
	WarningMessages []string              `json:"warningMessages"`
}

type IssueTransitionsScheme struct {
	Expand      string                   `json:"expand,omitempty"`
	Transitions []*IssueTransitionScheme `json:"transitions,omitempty"`
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

type IssueChangelogScheme struct {
	StartAt    int                            `json:"startAt"`
	MaxResults int                            `json:"maxResults"`
	Total      int                            `json:"total"`
	Histories  []*IssueChangelogHistoryScheme `json:"histories"`
}

type IssueChangelogHistoryScheme struct {
	ID string `json:"id"`

	Author struct {
		Self         string `json:"self"`
		AccountID    string `json:"accountId"`
		EmailAddress string `json:"emailAddress"`
		AvatarUrls   struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		DisplayName string `json:"displayName"`
		Active      bool   `json:"active"`
		TimeZone    string `json:"timeZone"`
		AccountType string `json:"accountType"`
	} `json:"author"`

	Created string `json:"created"`

	Items []*IssueChangelogHistoryItemScheme `json:"items"`
}

type IssueChangelogHistoryItemScheme struct {
	Field      string `json:"field"`
	Fieldtype  string `json:"fieldtype"`
	FieldID    string `json:"fieldId"`
	From       string `json:"from"`
	FromString string `json:"fromString"`
	To         string `json:"to"`
	ToString   string `json:"toString"`
}

// Get search issues using JQL query under the HTTP Method GET
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/search#search-for-issues-using-jql-get
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-search/#api-rest-api-3-search-get
func (s *IssueSearchService) Get(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int,
	validate string) (result *IssueSearchScheme, response *ResponseScheme, err error) {

	if len(jql) == 0 {
		return nil, nil, notJQLError
	}

	params := url.Values{}
	params.Add("jql", jql)
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(expands) != 0 {
		params.Add("expand", strings.Join(expands, ","))
	}

	if len(validate) != 0 {
		params.Add("validateQuery", validate)
	}

	if len(fields) != 0 {
		params.Add("fields", strings.Join(fields, ","))
	}

	var endpoint = fmt.Sprintf("rest/api/3/search?%v", params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Post search issues using JQL query under the HTTP Method POST
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/search#search-for-issues-using-jql-post
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-search/#api-rest-api-3-search-post
func (s *IssueSearchService) Post(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int,
	validate string) (result *IssueSearchScheme, response *ResponseScheme, err error) {

	payload := struct {
		Expand        []string `json:"expand,omitempty"`
		Jql           string   `json:"jql,omitempty"`
		MaxResults    int      `json:"maxResults,omitempty"`
		Fields        []string `json:"fields,omitempty"`
		StartAt       int      `json:"startAt,omitempty"`
		ValidateQuery string   `json:"validateQuery,omitempty"`
	}{
		Expand:        expands,
		Jql:           jql,
		MaxResults:    maxResults,
		Fields:        fields,
		StartAt:       startAt,
		ValidateQuery: validate,
	}

	var endpoint = "rest/api/3/search"

	payloadAsReader, _ := transformStructToReader(&payload)
	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

var (
	notJQLError = fmt.Errorf("error, please provide a valid JQL query")
)