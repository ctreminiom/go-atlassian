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
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-field-configurations/#api-rest-api-3-fieldconfiguration-get
func (f *FieldConfigurationService) Gets(ctx context.Context, IDs []int, isDefault bool, startAt, maxResults int) (result *FieldConfigSearchScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if isDefault {
		params.Add("isDefault", "true")
	}

	var IDsAsString string
	for index, value := range IDs {

		if index == 0 {
			IDsAsString = strconv.Itoa(value)
			continue
		}

		IDsAsString += "," + strconv.Itoa(value)
	}

	if len(IDsAsString) != 0 {
		params.Add("id", IDsAsString)
	}

	var endpoint = fmt.Sprintf("rest/api/3/fieldconfiguration?%v", params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldConfigSearchScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
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
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-field-configurations/#api-rest-api-3-fieldconfiguration-id-fields-get
func (f *FieldConfigurationService) Items(ctx context.Context, fieldConfigID string, startAt, maxResults int) (result *FieldConfigItemSearchScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("rest/api/3/fieldconfiguration/%v/fields?%v", fieldConfigID, params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldConfigItemSearchScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
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
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-field-configurations/#api-rest-api-3-fieldconfigurationscheme-get
func (f *FieldConfigurationService) Schemes(ctx context.Context, IDs []int, startAt, maxResults int) (result *FieldConfigSchemeScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var IDsAsString string
	for index, value := range IDs {

		if index == 0 {
			IDsAsString = strconv.Itoa(value)
			continue
		}

		IDsAsString += "," + strconv.Itoa(value)
	}

	if len(IDsAsString) != 0 {
		params.Add("id", IDsAsString)
	}

	var endpoint = fmt.Sprintf("rest/api/3/fieldconfigurationscheme?%v", params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldConfigSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
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
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-field-configurations/#api-rest-api-3-fieldconfigurationscheme-mapping-get
func (f *FieldConfigurationService) SchemesItems(ctx context.Context, fieldConfigIDs []int, startAt, maxResults int) (result *FieldConfigSchemeItemsScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var IDsAsString string
	for index, value := range fieldConfigIDs {

		if index == 0 {
			IDsAsString = strconv.Itoa(value)
			continue
		}

		IDsAsString += "," + strconv.Itoa(value)
	}

	if len(IDsAsString) != 0 {
		params.Add("fieldConfigurationSchemeId", IDsAsString)
	}

	var endpoint = fmt.Sprintf("rest/api/3/fieldconfigurationscheme/mapping?%v", params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldConfigSchemeItemsScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
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
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-field-configurations/#api-rest-api-3-fieldconfigurationscheme-project-get
func (f *FieldConfigurationService) SchemesByProject(ctx context.Context, projectIDs []int, startAt, maxResults int) (result *FieldProjectSchemeScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var IDsAsString string
	for index, value := range projectIDs {

		if index == 0 {
			IDsAsString = strconv.Itoa(value)
			continue
		}

		IDsAsString += "," + strconv.Itoa(value)
	}

	if len(IDsAsString) != 0 {
		params.Add("projectId", IDsAsString)
	}

	var endpoint = fmt.Sprintf("rest/api/3/fieldconfigurationscheme/project?%v", params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldProjectSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
