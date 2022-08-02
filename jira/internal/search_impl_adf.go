package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/jira"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type SearchADFService struct {
	internalClient jira.SearchADFConnector
}

func (s *SearchADFService) Checks(ctx context.Context, payload *model.IssueSearchCheckPayloadScheme) (*model.IssueMatchesPageScheme, *model.ResponseScheme, error) {
	return s.internalClient.Checks(ctx, payload)
}

func (s *SearchADFService) Get(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchScheme, *model.ResponseScheme, error) {
	return s.internalClient.Get(ctx, jql, fields, expands, startAt, maxResults, validate)
}

func (s *SearchADFService) Post(ctx context.Context, jql string, fields, expands []string, startAt, maxResults int, validate string) (*model.IssueSearchScheme, *model.ResponseScheme, error) {
	return s.internalClient.Post(ctx, jql, fields, expands, startAt, maxResults, validate)
}

type internalSearchADFImpl struct {
	c       service.Client
	version string
}

func (i *internalSearchADFImpl) Checks(ctx context.Context, payload *model.IssueSearchCheckPayloadScheme) (*model.IssueMatchesPageScheme, *model.ResponseScheme, error) {

	reader, err := i.c.TransformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/jql/match", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
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
		return nil, nil, model.ErrNoJQLError
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

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
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

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := fmt.Sprintf("rest/api/%v/search", i.version)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
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
