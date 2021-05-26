package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type FieldOptionContextService struct{ client *Client }

type FieldOptionContextParams struct {
	OptionID    int
	OnlyOptions bool
}

type CustomFieldContextOptionPageScheme struct {
	Self       string                            `json:"self,omitempty"`
	NextPage   string                            `json:"nextPage,omitempty"`
	MaxResults int                               `json:"maxResults,omitempty"`
	StartAt    int                               `json:"startAt,omitempty"`
	Total      int                               `json:"total,omitempty"`
	IsLast     bool                              `json:"isLast,omitempty"`
	Values     []*CustomFieldContextOptionScheme `json:"values,omitempty"`
}

type CustomFieldContextOptionScheme struct {
	ID       string `json:"id,omitempty"`
	Value    string `json:"value,omitempty"`
	Disabled bool   `json:"disabled"`
	OptionID string `json:"optionId,omitempty"`
}

// Returns a paginated list of all custom field option for a context.
// Options are returned first then cascading options, in the order they display in Jira.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#get-custom-field-options
func (f *FieldOptionContextService) Gets(ctx context.Context, fieldID string, contextID int, opts *FieldOptionContextParams, startAt, maxResults int) (result *CustomFieldContextOptionPageScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if opts.OptionID != 0 {
			params.Add("optionId", strconv.Itoa(opts.OptionID))
		}

		if opts.OnlyOptions {
			params.Add("onlyOptions", "true")
		}
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/option?%v", fieldID, contextID, params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(CustomFieldContextOptionPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

type FieldContextOptionListScheme struct {
	Options []*CustomFieldContextOptionScheme `json:"options,omitempty"`
}

// Creates options and, where the custom select field is of the type Select List (cascading),
// cascading options for a custom select field. The options are added to a context of the field.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#create-custom-field-options
func (f *FieldOptionContextService) Create(ctx context.Context, fieldID string, contextID int, payload *FieldContextOptionListScheme) (result *FieldContextOptionListScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, payload value is nil, please provide a valid CustomFieldOptionPayloadScheme pointer")
	}

	if fieldID == "" {
		return nil, nil, fmt.Errorf("error, fieldID value is nil, please provide a valid fieldID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/option", fieldID, contextID)

	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldContextOptionListScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Updates the options of a custom field.
// If any of the options are not found, no options are updated.
// Options where the values in the request match the current values aren't updated and aren't reported in the response.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#update-custom-field-options
func (f *FieldOptionContextService) Update(ctx context.Context, fieldID string, contextID int, payload *FieldContextOptionListScheme) (result *FieldContextOptionListScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, payload value is nil, please provide a valid CustomFieldOptionPayloadScheme pointer")
	}

	if fieldID == "" {
		return nil, nil, fmt.Errorf("error, fieldID value is nil, please provide a valid fieldID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/option", fieldID, contextID)

	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, payload)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldContextOptionListScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Deletes a custom field option.
// Options with cascading options cannot be deleted without deleting the cascading options first.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#delete-custom-field-options
func (f *FieldOptionContextService) Delete(ctx context.Context, fieldID string, contextID, optionID int) (response *Response, err error) {

	if fieldID == "" {
		return nil, fmt.Errorf("error, fieldID value is nil, please provide a valid fieldID value")
	}

	if contextID == 0 {
		return nil, fmt.Errorf("error, fieldID value is nil, please provide a valid contextID value")
	}

	if optionID == 0 {
		return nil, fmt.Errorf("error, fieldID value is nil, please provide a valid optionID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/option/%v", fieldID, contextID, optionID)

	request, err := f.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	return
}
