package jira

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type MySelfService struct{ client *Client }

// Details returns details for the current user.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-myself/#api-rest-api-3-myself-get
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-myself/#api-rest-api-3-myself-get
func (m *MySelfService) Details(ctx context.Context, expand []string) (result *UserScheme, response *ResponseScheme,
	err error) {

	params := url.Values{}
	if len(expand) != 0 {
		params.Add("expand", strings.Join(expand, ","))
	}

	var endpoint strings.Builder
	endpoint.WriteString("rest/api/3/myself")

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := m.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = m.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
