package jira

import (
	"context"
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

// Gets returns a paginated list of all custom field option for a context.
// Options are returned first then cascading options, in the order they display in Jira.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#get-custom-field-options
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-options/#api-rest-api-3-field-fieldid-context-contextid-option-get
func (f *FieldOptionContextService) Gets(ctx context.Context, fieldID string, contextID int, opts *FieldOptionContextParams,
	startAt, maxResults int) (result *CustomFieldContextOptionPageScheme, response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, nil, notFieldIDError
	}

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

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

type FieldContextOptionListScheme struct {
	Options []*CustomFieldContextOptionScheme `json:"options,omitempty"`
}

// Create creates options and, where the custom select field is of the type Select List (cascading),
// cascading options for a custom select field. The options are added to a context of the field.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#create-custom-field-options
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-options/#api-rest-api-3-field-fieldid-context-contextid-option-post
func (f *FieldOptionContextService) Create(ctx context.Context, fieldID string, contextID int, payload *FieldContextOptionListScheme) (
	result *FieldContextOptionListScheme, response *ResponseScheme, err error) {

	if fieldID == "" {
		return nil, nil, notFieldIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/option", fieldID, contextID)

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

// Update updates the options of a custom field.
// If any of the options are not found, no options are updated.
// Options where the values in the request match the current values aren't updated and aren't reported in the response.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#update-custom-field-options
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-options/#api-rest-api-3-field-fieldid-context-contextid-option-put
func (f *FieldOptionContextService) Update(ctx context.Context, fieldID string, contextID int,
	payload *FieldContextOptionListScheme) (result *FieldContextOptionListScheme, response *ResponseScheme, err error) {

	if fieldID == "" {
		return nil, nil, notFieldIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/option", fieldID, contextID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
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

// Delete deletes a custom field option.
// Options with cascading options cannot be deleted without deleting the cascading options first.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context/option#delete-custom-field-options
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-options/#api-rest-api-3-field-fieldid-context-contextid-option-optionid-delete
func (f *FieldOptionContextService) Delete(ctx context.Context, fieldID string, contextID, optionID int) (
	response *ResponseScheme, err error) {

	if fieldID == "" {
		return nil, notFieldIDError
	}

	if contextID == 0 {
		return nil, notContextIDError
	}

	if optionID == 0 {
		return nil, notOptionIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/option/%v", fieldID, contextID, optionID)

	request, err := f.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = f.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

type OrderFieldOptionPayloadScheme struct {
	After                string   `json:"after,omitempty"`
	Position             string   `json:"position,omitempty"`
	CustomFieldOptionIds []string `json:"customFieldOptionIds,omitempty"`
}

// Order changes the order of custom field options or cascading options in a context.
// Docs: TODO Issue 53
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-options/#api-rest-api-3-field-fieldid-context-contextid-option-move-put
func (f *FieldOptionContextService) Order(ctx context.Context, fieldID string, contextID int, payload *OrderFieldOptionPayloadScheme) (
	response *ResponseScheme, err error) {

	if fieldID == "" {
		return nil, notFieldIDError
	}

	if contextID == 0 {
		return nil, notContextIDError
	}

	var endpoint = fmt.Sprintf("/rest/api/3/field/%v/context/%v/option/move", fieldID, contextID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

var (
	notOptionIDError = fmt.Errorf("error, fieldID value is nil, please provide a valid optionID value")
)
