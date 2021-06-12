package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type MySelfService struct{ client *Client }

// Details returns details for the current user.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-myself/#api-rest-api-3-myself-get
func (m *MySelfService) Details(ctx context.Context, expands []string) (result *UserScheme, response *Response, err error) {

	params := url.Values{}

	var expand string
	for index, value := range expands {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	var endpoint string
	if params.Encode() != "" {
		endpoint = fmt.Sprintf("rest/api/3/myself?%v", params.Encode())
	} else {
		endpoint = "rest/api/3/myself"
	}

	request, err := m.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = m.client.Do(request)
	if err != nil {
		return
	}

	result = new(UserScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
