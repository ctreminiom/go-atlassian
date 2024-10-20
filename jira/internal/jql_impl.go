package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/v2/service"
	"github.com/ctreminiom/go-atlassian/v2/service/jira"
	"net/http"
	"net/url"
	"strings"
)

// NewJQLService creates a new instance of JQLService.
func NewJQLService(client service.Connector, version string) (*JQLService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &JQLService{
		internalClient: &internalJQLServiceImpl{c: client, version: version},
	}, nil
}

// JQLService provides methods to manage JQL queries in Jira Service Management.
type JQLService struct {
	// internalClient is the connector interface for JQL operations.
	internalClient jira.JQLConnector
}

// Parse parses and validates JQL queries.
//
// Validation is performed in context of the current user.
//
// POST /rest/api/{2-3}/jql/parse
//
// https://docs.go-atlassian.io/jira-software-cloud/jql#parse-jql-query
func (j *JQLService) Parse(ctx context.Context, validationType string, JqlQueries []string) (*model.ParsedQueryPageScheme, *model.ResponseScheme, error) {
	return j.internalClient.Parse(ctx, validationType, JqlQueries)
}

type internalJQLServiceImpl struct {
	c       service.Connector
	version string
}

func (i *internalJQLServiceImpl) Parse(ctx context.Context, validationType string, JqlQueries []string) (*model.ParsedQueryPageScheme, *model.ResponseScheme, error) {

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("/rest/api/%v/jql/parse", i.version))

	if validationType != "" {
		params := url.Values{}
		params.Add("validation", validationType)

		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint.String(), "", map[string]interface{}{"queries": JqlQueries})
	if err != nil {
		return nil, nil, err
	}

	page := new(model.ParsedQueryPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}
