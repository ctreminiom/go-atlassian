package jira

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type SearchSharedConnector interface {

	// Checks checks whether one or more issues would be returned by one or more JQL queries.
	//
	// POST /rest/api/{2-3}/jql/match
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/search#check-issues-against-jql
	Checks(ctx context.Context, payload *model.IssueSearchCheckPayloadScheme) (*model.IssueMatchesPageScheme, *model.ResponseScheme, error)
}

type SearchRichTextConnector interface {
	SearchSharedConnector

	// Get search issues using JQL query under the HTTP Method GET
	//
	// GET /rest/api/2/search
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/search#search-for-issues-using-jql-get
	//
	// Deprecated: This endpoint will be removed after May 1, 2025. Use SearchJQL, BulkFetch and ApproximateCount instead.
	Get(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchSchemeV2, *model.ResponseScheme, error)

	// Post search issues using JQL query under the HTTP Method POST
	//
	// POST /rest/api/2/search
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/search#search-for-issues-using-jql-get
	//
	// Deprecated: This endpoint will be removed after May 1, 2025. Use SearchJQL, BulkFetch and ApproximateCount instead.
	Post(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchSchemeV2, *model.ResponseScheme, error)

	// SearchJQL search issues using the new JQL search endpoint
	//
	// POST /rest/api/2/search/jql
	//
	SearchJQL(ctx context.Context, jql string, fields, expands []string, maxResults int, nextPageToken string) (*model.IssueSearchJQLSchemeV2, *model.ResponseScheme, error)

	// ApproximateCount gets an approximate count of issues matching a JQL query
	//
	// POST /rest/api/2/search/approximate-count
	//
	ApproximateCount(ctx context.Context, jql string) (*model.IssueSearchApproximateCountScheme, *model.ResponseScheme, error)

	// BulkFetch fetches multiple issues by their IDs or keys
	//
	// POST /rest/api/2/issue/bulkfetch
	//
	BulkFetch(ctx context.Context, issueIDsOrKeys []string, fields []string) (*model.IssueBulkFetchSchemeV2, *model.ResponseScheme, error)
}

type SearchADFConnector interface {
	SearchSharedConnector

	// Get search issues using JQL query under the HTTP Method GET
	//
	// GET /rest/api/3/search
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/search#search-for-issues-using-jql-get
	//
	// Deprecated: This endpoint will be removed after May 1, 2025. Use SearchJQL, BulkFetch and ApproximateCount instead.
	Get(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchScheme, *model.ResponseScheme, error)

	// Post search issues using JQL query under the HTTP Method POST
	//
	// POST /rest/api/3/search
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/search#search-for-issues-using-jql-get
	//
	// Deprecated: This endpoint will be removed after May 1, 2025. Use SearchJQL, BulkFetch and ApproximateCount instead.
	Post(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchScheme, *model.ResponseScheme, error)

	// SearchJQL searches issues using the new JQL search endpoint
	//
	// POST /rest/api/3/search/jql
	//
	SearchJQL(ctx context.Context, jql string, fields, expands []string, maxResults int, nextPageToken string) (*model.IssueSearchJQLScheme, *model.ResponseScheme, error)

	// ApproximateCount gets an approximate count of issues matching a JQL query
	//
	// POST /rest/api/3/search/approximate-count
	//
	ApproximateCount(ctx context.Context, jql string) (*model.IssueSearchApproximateCountScheme, *model.ResponseScheme, error)

	// BulkFetch fetches multiple issues by their IDs or keys
	//
	// POST /rest/api/3/issue/bulkfetch
	//
	BulkFetch(ctx context.Context, issueIDsOrKeys []string, fields []string) (*model.IssueBulkFetchScheme, *model.ResponseScheme, error)
}
