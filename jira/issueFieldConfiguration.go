package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type FieldConfigurationService struct{ client *Client }

type FieldConfigSearchScheme struct {
	MaxResults int                 `json:"maxResults"`
	StartAt    int                 `json:"startAt"`
	Total      int                 `json:"total"`
	IsLast     bool                `json:"isLast"`
	Values     []FieldConfigScheme `json:"values"`
}

type FieldConfigScheme struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDefault   bool   `json:"isDefault,omitempty"`
}

// Returns a paginated list of all field configurations.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-all-field-configurations
func (f *FieldConfigurationService) Gets(ctx context.Context, IDs []int, isDefault bool, startAt, maxResults int) (result *FieldConfigSearchScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if isDefault {
		params.Add("isDefault", "true")
	}

	for _, id := range IDs {
		params.Add("id", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/3/fieldconfiguration?%v", params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldConfigSearchScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

type FieldConfigItemSearchScheme struct {
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		ID          string `json:"id"`
		IsHidden    bool   `json:"isHidden"`
		IsRequired  bool   `json:"isRequired"`
		Description string `json:"description,omitempty"`
	} `json:"values"`
}

// Returns a paginated list of all fields for a configuration.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-field-configuration-items
func (f *FieldConfigurationService) Items(ctx context.Context, fieldConfigID int, startAt, maxResults int) (result *FieldConfigItemSearchScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("rest/api/3/fieldconfiguration/%v/fields?%v", fieldConfigID, params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldConfigItemSearchScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

type FieldConfigSchemeScheme struct {
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"values"`
}

// Returns a paginated list of field configuration schemes.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-all-field-configuration-schemes
func (f *FieldConfigurationService) Schemes(ctx context.Context, IDs []int, startAt, maxResults int) (result *FieldConfigSchemeScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range IDs {
		params.Add("id", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/3/fieldconfigurationscheme?%v", params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldConfigSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

type FieldConfigSchemeItemsScheme struct {
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		FieldConfigurationSchemeID string `json:"fieldConfigurationSchemeId"`
		IssueTypeID                string `json:"issueTypeId"`
		FieldConfigurationID       string `json:"fieldConfigurationId"`
	} `json:"values"`
}

// Returns a paginated list of field configuration issue type items.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-field-configuration-issue-type-items
func (f *FieldConfigurationService) IssueTypeItems(ctx context.Context, fieldConfigIDs []int, startAt, maxResults int) (result *FieldConfigSchemeItemsScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range fieldConfigIDs {
		params.Add("fieldConfigurationSchemeId", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/3/fieldconfigurationscheme/mapping?%v", params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldConfigSchemeItemsScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

type FieldProjectSchemeScheme struct {
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		ProjectIds               []string `json:"projectIds"`
		FieldConfigurationScheme struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"fieldConfigurationScheme,omitempty"`
	} `json:"values"`
}

// Returns a paginated list of field configuration schemes and, for each scheme, a list of the projects that use it.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-field-configuration-schemes-for-projects
func (f *FieldConfigurationService) SchemesByProject(ctx context.Context, projectIDs []int, startAt, maxResults int) (result *FieldProjectSchemeScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, projectID := range projectIDs {
		params.Add("projectId", strconv.Itoa(projectID))
	}

	var endpoint = fmt.Sprintf("rest/api/3/fieldconfigurationscheme/project?%v", params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldProjectSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}
