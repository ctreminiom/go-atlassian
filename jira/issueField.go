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
	Option        *FieldOptionService
}

type FieldScheme struct {
	ID          string   `json:"id"`
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Custom      bool     `json:"custom"`
	Orderable   bool     `json:"orderable"`
	Navigable   bool     `json:"navigable"`
	Searchable  bool     `json:"searchable"`
	ClauseNames []string `json:"clauseNames"`
	Schema      struct {
		Type   string `json:"type"`
		System string `json:"system"`
	} `json:"schema,omitempty"`
	UntranslatedName string `json:"untranslatedName,omitempty"`
}

// Returns system and custom issue fields according to the following rules:
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-fields/
func (f *FieldService) Gets(ctx context.Context) (result *[]FieldScheme, response *Response, err error) {

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

	result = new([]FieldScheme)
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
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-fields/#api-rest-api-3-field-post
func (f *FieldService) Create(ctx context.Context, payload *CustomFieldScheme) (result *FieldScheme, response *Response, err error) {

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

	result = new(FieldScheme)
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
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Schema struct {
			Type   string `json:"type"`
			Items  string `json:"items"`
			System string `json:"system"`
		} `json:"schema,omitempty"`
		Description   string `json:"description,omitempty"`
		Key           string `json:"key,omitempty"`
		IsLocked      bool   `json:"isLocked,omitempty"`
		ScreensCount  int    `json:"screensCount,omitempty"`
		ContextsCount int    `json:"contextsCount,omitempty"`
		LastUsed      struct {
			Type string `json:"type,omitempty"`
		} `json:"lastUsed,omitempty"`
	} `json:"values,omitempty"`
}

// Returns a paginated list of fields for Classic Jira projects.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-fields/#api-rest-api-3-field-search-get
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
