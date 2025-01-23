package jira

import (
	"context"
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

type JQLConnector interface {

	// Parse parses and validates JQL queries.
	//
	// Validation is performed in context of the current user.
	//
	// POST /rest/api/{2-3}/jql/parse
	//
	// https://docs.go-atlassian.io/jira-software-cloud/jql#parse-jql-query
	Parse(ctx context.Context, validationType string, JqlQueries []string) (*models.ParsedQueryPageScheme, *models.ResponseScheme, error)
}
