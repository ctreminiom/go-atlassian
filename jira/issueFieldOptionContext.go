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
	FieldID     string
	ContextID   int
	OptionID    int
	OnlyOptions bool
}

type FieldContextOptionScheme struct {
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		ID       string `json:"id"`
		Value    string `json:"value"`
		Disabled bool   `json:"disabled,omitempty"`
		OptionID string `json:"optionId,omitempty"`
	} `json:"values"`
}

// Returns a paginated list of all custom field option for a context.
// Options are returned first then cascading options, in the order they display in Jira.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-custom-field-options/#api-rest-api-3-field-fieldid-context-contextid-option-get
func (f *FieldOptionContextService) Options(ctx context.Context, opts *FieldOptionContextParams, startAt, maxResults int) (result *FieldContextOptionScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(opts.FieldID) != 0 {
		params.Add("fieldId", opts.FieldID)
	}

	if opts.ContextID != 0 {
		params.Add("contextId", strconv.Itoa(opts.ContextID))
	}

	if opts.OptionID != 0 {
		params.Add("optionId", strconv.Itoa(opts.OptionID))
	}

	if opts.OnlyOptions {
		params.Add("onlyOptions", "true")
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/option?%v", opts.FieldID, opts.ContextID, params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldContextOptionScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
