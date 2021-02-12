package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type FieldContextService struct {
	client *Client
	Option *FieldOptionContextService
}

type FieldContextOptionsScheme struct {
	IsAnyIssueType  bool
	IsGlobalContext bool
	ContextID       []int
}

type FieldContextSearchScheme struct {
	MaxResults int                  `json:"maxResults"`
	StartAt    int                  `json:"startAt"`
	Total      int                  `json:"total"`
	IsLast     bool                 `json:"isLast"`
	Values     []FieldContextScheme `json:"values"`
}

type FieldContextScheme struct {
	ID              string        `json:"id"`
	Name            string        `json:"name"`
	Description     string        `json:"description"`
	IsGlobalContext bool          `json:"isGlobalContext"`
	IsAnyIssueType  bool          `json:"isAnyIssueType"`
	ProjectIds      []interface{} `json:"projectIds,omitempty"`
	IssueTypeIds    []string      `json:"issueTypeIds,omitempty"`
}

// Returns a paginated list of contexts for a custom field. Contexts can be returned as follows:
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-get
func (f *FieldContextService) Gets(ctx context.Context, fieldID string, opts *FieldContextOptionsScheme, startAt, maxResults int) (result *FieldContextSearchScheme, response *Response, err error) {

	if opts == nil {
		return nil, nil, fmt.Errorf("error, payload value is nil, please provide a valid FieldContextOptionsScheme pointer")
	}

	if fieldID == "" {
		return nil, nil, fmt.Errorf("error, fieldID value is nil, please provide a valid fieldID value")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts.IsAnyIssueType {
		params.Add("isAnyIssueType", "true")
	}

	if opts.IsGlobalContext {
		params.Add("isGlobalContext", "true")
	}

	for _, contextID := range opts.ContextID {
		params.Add("contextId", strconv.Itoa(contextID))
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context?%v", fieldID, params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldContextSearchScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

type FieldContextPayloadScheme struct {
	IssueTypeIDs []int  `json:"issueTypeIds,omitempty"`
	ProjectIDs   []int  `json:"projectIds,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
}

// Creates a custom field context.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-contexts/#api-rest-api-3-field-fieldid-context-post
func (f *FieldContextService) Create(ctx context.Context, fieldID string, payload *FieldContextPayloadScheme) (result *FieldContextScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, payload value is nil, please provide a valid FieldContextPayloadScheme pointer")
	}

	if fieldID == "" {
		return nil, nil, fmt.Errorf("error, fieldID value is nil, please provide a valid fieldID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context", fieldID)
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

	result = new(FieldContextScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}
