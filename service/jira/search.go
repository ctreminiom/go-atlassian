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
	Get(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchSchemeV2, *model.ResponseScheme, error)

	// Post search issues using JQL query under the HTTP Method POST
	//
	// POST /rest/api/2/search
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/search#search-for-issues-using-jql-get
	Post(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchSchemeV2, *model.ResponseScheme, error)
}

type SearchADFConnector interface {
	SearchSharedConnector

	// Get search issues using JQL query under the HTTP Method GET
	//
	// GET /rest/api/3/search
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/search#search-for-issues-using-jql-get
	Get(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchScheme, *model.ResponseScheme, error)

	// Post search issues using JQL query under the HTTP Method POST
	//
	// POST /rest/api/3/search
	//
	// https://docs.go-atlassian.io/jira-software-cloud/issues/search#search-for-issues-using-jql-get
	Post(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchScheme, *model.ResponseScheme, error)
}
