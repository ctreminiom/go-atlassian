package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type FieldOptionService struct{ client *Client }

type FieldOptionSearchScheme struct {
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		ID               int      `json:"id"`
		Value            string   `json:"value"`
		CascadingOptions []string `json:"cascadingOptions"`
	} `json:"values"`
}

// Returns a paginated list of options and, where the custom select field is of the type Select List (cascading),
// cascading options for custom select fields. Cascading options are included in the item count when determining pagination.
// Only options from the global context are returned.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-options/#api-rest-api-3-customfield-fieldid-option-get
func (f *FieldOptionService) Gets(ctx context.Context, fieldID string, startAt, maxResults int) (result *FieldOptionSearchScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("rest/api/3/customField/%v/option?%v", fieldID, params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldOptionSearchScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Updates the options on a custom select field. Where an updated option is in use on an issue, the value on the issue is also updated.
// Options that are not found are ignored. A maximum of 1000 options, including sub-options of Select List (cascading) fields,
// can be updated per request.
// The options are updated on the global context of the field.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-options/#api-rest-api-3-customfield-fieldid-option-put
func (f *FieldOptionService) Update(ctx context.Context, fieldID string, payload interface{}) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/customField/%v/option", fieldID)
	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Creates options and, where the custom select field is of the type Select List (cascading),
// cascading options for a custom select field. The options are added to the global context of the field.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-options/#api-rest-api-3-customfield-fieldid-option-post
func (f *FieldOptionService) Create(ctx context.Context, fieldID string, payload interface{}) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/customField/%v/option", fieldID)
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

	return
}

type CustomFieldOptionScheme struct {
	Self  string `json:"self"`
	Value string `json:"value"`
}

// Returns a custom field option. For example, an option in a select list.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-options/#api-rest-api-3-customfieldoption-id-get
func (f *FieldOptionService) Option(ctx context.Context, customFieldOptionID string) (result *CustomFieldOptionScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/customFieldOption/%v", customFieldOptionID)
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(CustomFieldOptionScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
