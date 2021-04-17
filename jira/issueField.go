package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
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

// Returns system and custom issue fields according to the following rules:
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields#get-fields
func (f *FieldService) Gets(ctx context.Context) (result *[]IssueFieldScheme, response *Response, err error) {

	var endpoint = "rest/api/3/field"
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new([]IssueFieldScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

type CustomFieldScheme struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	FieldType   string `json:"type,omitempty"`
	SearcherKey string `json:"searcherKey,omitempty"`
}

func (c *CustomFieldScheme) Format() {
	//Append the atlassian plugin name convention
	c.FieldType = fmt.Sprintf("com.atlassian.jira.plugin.system.customfieldtypes:%v", c.FieldType)
	c.SearcherKey = fmt.Sprintf("com.atlassian.jira.plugin.system.customfieldtypes:%v", c.SearcherKey)
}

// Creates a custom field.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields#create-custom-field
func (f *FieldService) Create(ctx context.Context, payload *CustomFieldScheme) (result *IssueFieldScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, payload value is nil, please provide a valid CustomFieldScheme pointer")
	}

	payload.Format()

	var endpoint = "rest/api/3/field"
	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueFieldScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
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

type FieldSearchScheme struct {
	MaxResults int                 `json:"maxResults,omitempty"`
	StartAt    int                 `json:"startAt,omitempty"`
	Total      int                 `json:"total,omitempty"`
	IsLast     bool                `json:"isLast,omitempty"`
	Values     []*IssueFieldScheme `json:"values,omitempty"`
}

// Returns a paginated list of fields for Classic Jira projects.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields#get-fields-paginated
func (f *FieldService) Search(ctx context.Context, opts *FieldSearchOptionsScheme, startAt, maxResults int) (result *FieldSearchScheme, response *Response, err error) {

	if opts == nil {
		return nil, nil, fmt.Errorf("error, opts value is nil, please provide a valid FieldSearchOptionsScheme pointer")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var expand string
	for index, value := range opts.Expand {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	var fieldTypes string
	for index, value := range opts.Types {

		if index == 0 {
			fieldTypes = value
			continue
		}

		fieldTypes += "," + value
	}

	if len(fieldTypes) != 0 {
		params.Add("type", fieldTypes)
	}

	if opts.Query != "" {
		params.Add("orderBy", opts.Query)
	}

	if opts.OrderBy != "" {
		params.Add("orderBy", opts.OrderBy)
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/search?%v", params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldSearchScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}
