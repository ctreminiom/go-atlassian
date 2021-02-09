package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type IssueSearchService struct{ client *Client }

type IssueSearchScheme struct {
	Expand          string        `json:"expand"`
	StartAt         int           `json:"startAt"`
	MaxResults      int           `json:"maxResults"`
	Total           int           `json:"total"`
	Issues          []IssueScheme `json:"issues"`
	WarningMessages []string      `json:"warningMessages"`
}

type IssueSearchTransitionScheme struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	To   struct {
		Self           string `json:"self"`
		Description    string `json:"description"`
		IconURL        string `json:"iconUrl"`
		Name           string `json:"name"`
		ID             string `json:"id"`
		StatusCategory struct {
			Self      string `json:"self"`
			ID        int    `json:"id"`
			Key       string `json:"key"`
			ColorName string `json:"colorName"`
			Name      string `json:"name"`
		} `json:"statusCategory"`
	} `json:"to"`
	HasScreen     bool `json:"hasScreen"`
	IsGlobal      bool `json:"isGlobal"`
	IsInitial     bool `json:"isInitial"`
	IsAvailable   bool `json:"isAvailable"`
	IsConditional bool `json:"isConditional"`
	IsLooped      bool `json:"isLooped"`
}

type IssueChangelogScheme struct {
	StartAt    int                           `json:"startAt"`
	MaxResults int                           `json:"maxResults"`
	Total      int                           `json:"total"`
	Histories  []IssueChangelogHistoryScheme `json:"histories"`
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

	Items []IssueChangelogHistoryItemScheme `json:"items"`
}

type IssueChangelogHistoryItemScheme struct {
	Field      string `json:"field"`
	Fieldtype  string `json:"fieldtype"`
	From       string `json:"from"`
	FromString string `json:"fromString"`
	To         string `json:"to"`
}

// Search issues using JQL query under the HTTP Method GET
// If the JQL query expression is too large to be encoded as a query parameter, use the POST version of this resource.
// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-search/#api-rest-api-3-search-get
func (s *IssueSearchService) Get(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int) (result *IssueSearchScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("jql", jql)
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var expand string
	for index, value := range expands {

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
	for index, value := range fields {

		if index == 0 {
			fieldFormatted = value
			continue
		}
		fieldFormatted += "," + value
	}

	if len(fieldFormatted) != 0 {
		params.Add("fields", fieldFormatted)
	}

	var endpoint = fmt.Sprintf("rest/api/3/search?%v", params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueSearchScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Search issues using JQL query under the HTTP Method POST
// There is a GET version of this resource that can be used for smaller JQL query expressions.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-search/#api-rest-api-3-search-post
func (s *IssueSearchService) Post(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int) (result *IssueSearchScheme, response *Response, err error) {

	payload := struct {
		Expand     []string `json:"expand"`
		Jql        string   `json:"jql"`
		MaxResults int      `json:"maxResults"`
		Fields     []string `json:"fields"`
		StartAt    int      `json:"startAt"`
	}{
		Expand:     expands,
		Jql:        jql,
		MaxResults: maxResults,
		Fields:     fields,
		StartAt:    startAt,
	}

	var endpoint = "rest/api/3/search"

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueSearchScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
