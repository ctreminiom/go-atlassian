package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SearchRichTextService provides methods to manage rich text searches in Jira Service Management.
type SearchRichTextService struct {
	// internalClient is the connector interface for rich text search operations.
	internalClient jira.SearchRichTextConnector
}

// Checks checks whether one or more issues would be returned by one or more JQL queries.
//
// POST /rest/api/{2-3}/jql/match
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/search#check-issues-against-jql
func (s *SearchRichTextService) Checks(ctx context.Context, payload *model.IssueSearchCheckPayloadScheme) (*model.IssueMatchesPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Checks(ctx, payload)
}

// Get search issues using JQL query under the HTTP Method GET
//
// GET /rest/api/2/search
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/search#search-for-issues-using-jql-get
func (s *SearchRichTextService) Get(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchSchemeV2, *model.ResponseScheme, error) {
	return s.internalClient.Get(ctx, jql, fields, expands, startAt, maxResults, validate)
}

// Post search issues using JQL query under the HTTP Method POST
//
// POST /rest/api/2/search
//
// https://docs.go-atlassian.io/jira-software-cloud/issues/search#search-for-issues-using-jql-get
func (s *SearchRichTextService) Post(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchSchemeV2, *model.ResponseScheme, error) {
	return s.internalClient.Post(ctx, jql, fields, expands, startAt, maxResults, validate)
}

type internalSearchRichTextImpl struct {
	c       service.Connector
	version string
}

func (i *internalSearchRichTextImpl) Checks(ctx context.Context, payload *model.IssueSearchCheckPayloadScheme) (*model.IssueMatchesPageScheme, *model.ResponseScheme, error) {

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

func (i *internalSearchRichTextImpl) Get(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchSchemeV2, *model.ResponseScheme, error) {

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

	issues := new(model.IssueSearchSchemeV2)
	response, err := i.c.Call(request, issues)
	if err != nil {
		return nil, response, err
	}

	return issues, response, nil
}

func (i *internalSearchRichTextImpl) Post(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchSchemeV2, *model.ResponseScheme, error) {

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

	issues := new(model.IssueSearchSchemeV2)
	response, err := i.c.Call(request, issues)
	if err != nil {
		return nil, response, err
	}

	return issues, response, nil
}
