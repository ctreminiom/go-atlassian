package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
)

// SearchADFService provides methods to manage advanced document format (ADF) searches in Jira Service Management.
type SearchADFService struct {
	// internalClient is the connector interface for ADF search operations.
	internalClient jira.SearchADFConnector
}

// Checks checks whether one or more issues would be returned by one or more JQL queries.
//
// POST /rest/api/{2-3}/jql/match
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/search#check-issues-against-jql
func (s *SearchADFService) Checks(ctx context.Context, payload *model.IssueSearchCheckPayloadScheme) (*model.IssueMatchesPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Checks(ctx, payload)
}

// Get search issues using JQL query under the HTTP Method GET
//
// GET /rest/api/3/search
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/search#search-for-issues-using-jql-get
//
// Deprecated: This endpoint will be removed after May 1, 2025. Use SearchJQL, BulkFetch and ApproximateCount instead.
func (s *SearchADFService) Get(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchScheme, *model.ResponseScheme, error) {
	return s.internalClient.Get(ctx, jql, fields, expands, startAt, maxResults, validate)
}

// Post search issues using JQL query under the HTTP Method POST
//
// POST /rest/api/3/search
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/search#search-for-issues-using-jql-get
//
// Deprecated: This endpoint will be removed after May 1, 2025. Use SearchJQL, BulkFetch and ApproximateCount instead.
func (s *SearchADFService) Post(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchScheme, *model.ResponseScheme, error) {
	return s.internalClient.Post(ctx, jql, fields, expands, startAt, maxResults, validate)
}

// SearchJQL searches issues using the new JQL search endpoint
//
// POST /rest/api/3/search/jql
func (s *SearchADFService) SearchJQL(ctx context.Context, jql string, fields, expands []string, maxResults int, nextPageToken string) (*model.IssueSearchJQLScheme, *model.ResponseScheme, error) {
	return s.internalClient.SearchJQL(ctx, jql, fields, expands, maxResults, nextPageToken)
}

// ApproximateCount gets an approximate count of issues matching a JQL query
//
// POST /rest/api/3/search/approximate-count
func (s *SearchADFService) ApproximateCount(ctx context.Context, jql string) (*model.IssueSearchApproximateCountScheme, *model.ResponseScheme, error) {
	return s.internalClient.ApproximateCount(ctx, jql)
}

// BulkFetch fetches multiple issues by their IDs or keys
//
// POST /rest/api/3/issue/bulkfetch
func (s *SearchADFService) BulkFetch(ctx context.Context, issueIDsOrKeys []string, fields []string) (*model.IssueBulkFetchScheme, *model.ResponseScheme, error) {
	return s.internalClient.BulkFetch(ctx, issueIDsOrKeys, fields)
}

type internalSearchADFImpl struct {
	c       service.Connector
	version string
}

func (i *internalSearchADFImpl) Checks(ctx context.Context, payload *model.IssueSearchCheckPayloadScheme) (*model.IssueMatchesPageScheme, *model.ResponseScheme, error) {

	endpoint := fmt.Sprintf("rest/api/%v/jql/match", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	issues := new(model.IssueMatchesPageScheme)
	response, err := i.c.Call(request, issues)
	if err != nil {
		return nil, response, err
	}

	return issues, response, nil
}

func (i *internalSearchADFImpl) Get(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchScheme, *model.ResponseScheme, error) {

	if jql == "" {
		return nil, nil, model.ErrNoJQL
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

	endpoint := fmt.Sprintf("rest/api/%v/search?%v", i.version, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	issues := new(model.IssueSearchScheme)
	response, err := i.c.Call(request, issues)
	if err != nil {
		return nil, response, err
	}

	return issues, response, nil
}

func (i *internalSearchADFImpl) Post(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchScheme, *model.ResponseScheme, error) {

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

	endpoint := fmt.Sprintf("rest/api/%v/search", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	issues := new(model.IssueSearchScheme)
	response, err := i.c.Call(request, issues)
	if err != nil {
		return nil, response, err
	}

	return issues, response, nil
}

// SearchJQL searches issues using the new JQL search endpoint
//
// POST /rest/api/3/search/jql
func (i *internalSearchADFImpl) SearchJQL(ctx context.Context, jql string, fields, expands []string, maxResults int, nextPageToken string) (*model.IssueSearchJQLScheme, *model.ResponseScheme, error) {

	payload := struct {
		Jql           string   `json:"jql,omitempty"`
		MaxResults    int      `json:"maxResults,omitempty"`
		Fields        []string `json:"fields,omitempty"`
		Expand        []string `json:"expand,omitempty"`
		NextPageToken string   `json:"nextPageToken,omitempty"`
	}{
		Jql:           jql,
		MaxResults:    maxResults,
		Fields:        fields,
		Expand:        expands,
		NextPageToken: nextPageToken,
	}

	endpoint := fmt.Sprintf("rest/api/%v/search/jql", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	issues := new(model.IssueSearchJQLScheme)
	response, err := i.c.Call(request, issues)
	if err != nil {
		return nil, response, err
	}

	return issues, response, nil
}

// ApproximateCount gets an approximate count of issues matching a JQL query
//
// POST /rest/api/3/search/approximate-count
func (i *internalSearchADFImpl) ApproximateCount(ctx context.Context, jql string) (*model.IssueSearchApproximateCountScheme, *model.ResponseScheme, error) {

	payload := struct {
		Jql string `json:"jql,omitempty"`
	}{
		Jql: jql,
	}

	endpoint := fmt.Sprintf("rest/api/%v/search/approximate-count", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	count := new(model.IssueSearchApproximateCountScheme)
	response, err := i.c.Call(request, count)
	if err != nil {
		return nil, response, err
	}

	return count, response, nil
}

// BulkFetch fetches multiple issues by their IDs or keys
//
// POST /rest/api/3/issue/bulkfetch
func (i *internalSearchADFImpl) BulkFetch(ctx context.Context, issueIDsOrKeys []string, fields []string) (*model.IssueBulkFetchScheme, *model.ResponseScheme, error) {

	payload := struct {
		IssueIDsOrKeys []string `json:"issueIdsOrKeys,omitempty"`
		Fields         []string `json:"fields,omitempty"`
	}{
		IssueIDsOrKeys: issueIDsOrKeys,
		Fields:         fields,
	}

	endpoint := fmt.Sprintf("rest/api/%v/issue/bulkfetch", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	issues := new(model.IssueBulkFetchScheme)
	response, err := i.c.Call(request, issues)
	if err != nil {
		return nil, response, err
	}

	return issues, response, nil
}
