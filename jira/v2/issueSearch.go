package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type IssueSearchService struct{ client *Client }

// Get search issues using JQL query under the HTTP Method GET
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/search#search-for-issues-using-jql-get
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-search/#api-rest-api-2-search-get
func (s *IssueSearchService) Get(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int,
	validate string) (result *models.IssueSearchSchemeV2, response *ResponseScheme, err error) {

	if len(jql) == 0 {
		return nil, nil, models.ErrNoJQLError
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

	var endpoint = fmt.Sprintf("rest/api/2/search?%v", params.Encode())

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
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-search/#api-rest-api-2-search-post
func (s *IssueSearchService) Post(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int,
	validate string) (result *models.IssueSearchSchemeV2, response *ResponseScheme, err error) {

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

	var endpoint = "rest/api/2/search"

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

// Checks checks whether one or more issues would be returned by one or more JQL queries.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/search#check-issues-against-jql
func (s *IssueSearchService) Checks(ctx context.Context, payload *models.IssueSearchCheckPayloadScheme) (result *models.IssueMatchesPageScheme,
	response *ResponseScheme, err error) {

	var endpoint = "rest/api/2/jql/match"

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

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
