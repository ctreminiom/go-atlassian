package jira

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type FieldService struct {
	client        *Client
	Configuration *FieldConfigurationService
	Context       *FieldContextService
}

type IssueFieldScheme struct {
	ID            string                         `json:"id,omitempty"`
	Key           string                         `json:"key,omitempty"`
	Name          string                         `json:"name,omitempty"`
	Custom        bool                           `json:"custom,omitempty"`
	Orderable     bool                           `json:"orderable,omitempty"`
	Navigable     bool                           `json:"navigable,omitempty"`
	Searchable    bool                           `json:"searchable,omitempty"`
	ClauseNames   []string                       `json:"clauseNames,omitempty"`
	Scope         *TeamManagedProjectScopeScheme `json:"scope,omitempty"`
	Schema        *IssueFieldSchemaScheme        `json:"schema,omitempty"`
	Description   string                         `json:"description,omitempty"`
	IsLocked      bool                           `json:"isLocked,omitempty"`
	SearcherKey   string                         `json:"searcherKey,omitempty"`
	ScreensCount  int                            `json:"screensCount,omitempty"`
	ContextsCount int                            `json:"contextsCount,omitempty"`
	LastUsed      *IssueFieldLastUsedScheme      `json:"lastUsed,omitempty"`
}

type IssueFieldLastUsedScheme struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type TeamManagedProjectScopeScheme struct {
	Type    string         `json:"type,omitempty"`
	Project *ProjectScheme `json:"project,omitempty"`
}

type IssueFieldSchemaScheme struct {
	Type     string `json:"type,omitempty"`
	Items    string `json:"items,omitempty"`
	System   string `json:"system,omitempty"`
	Custom   string `json:"custom,omitempty"`
	CustomID int    `json:"customId,omitempty"`
}

// Gets returns system and custom issue fields according to the following rules:
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields#get-fields
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-fields/#api-rest-api-3-field-get
func (f *FieldService) Gets(ctx context.Context) (result []*IssueFieldScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/field"

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type CustomFieldScheme struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	FieldType   string `json:"type,omitempty"`
	SearcherKey string `json:"searcherKey,omitempty"`
}

// Create creates a custom field.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields#create-custom-field
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-fields/#api-rest-api-3-field-post
func (f *FieldService) Create(ctx context.Context, payload *CustomFieldScheme) (result *IssueFieldScheme,
	response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/field"

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type FieldSearchOptionsScheme struct {
	Types   []string
	IDs     []string
	Query   string
	OrderBy string
	Expand  []string
}

type FieldSearchPageScheme struct {
	MaxResults int                 `json:"maxResults,omitempty"`
	StartAt    int                 `json:"startAt,omitempty"`
	Total      int                 `json:"total,omitempty"`
	IsLast     bool                `json:"isLast,omitempty"`
	Values     []*IssueFieldScheme `json:"values,omitempty"`
}

// Search returns a paginated list of fields for Classic Jira projects.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields#get-fields-paginated
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-fields/#api-rest-api-3-field-search-get
// NOTE: Experimental Endpoint
func (f *FieldService) Search(ctx context.Context, options *FieldSearchOptionsScheme, startAt, maxResults int) (
	result *FieldSearchPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}

		if len(options.Types) != 0 {
			params.Add("type", strings.Join(options.Types, ","))
		}

		if len(options.IDs) != 0 {
			params.Add("id", strings.Join(options.IDs, ","))
		}

		if len(options.OrderBy) != 0 {
			params.Add("orderBy", options.OrderBy)
		}

		if len(options.Query) != 0 {
			params.Add("query", options.Query)
		}
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/search?%v", params.Encode())

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
