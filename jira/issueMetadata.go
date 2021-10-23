package jira

import (
	"context"
	"fmt"
	"github.com/tidwall/gjson"
	"net/http"
	"net/url"
	"strings"
)

type IssueMetadataService struct{ client *Client }

// Get edit issue metadata returns the edit screen fields for an issue that are visible to and editable by the user.
// Use the information to populate the requests in Edit issue.
// Atlassian URL: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issues/#api-rest-api-3-issue-issueidorkey-editmeta-get
// Docs: N/A
func (i *IssueMetadataService) Get(ctx context.Context, issueKeyOrID string, overrideScreenSecurity, overrideEditableFlag bool) (result gjson.Result,
	response *ResponseScheme, err error) {

	if len(issueKeyOrID) == 0 {
		return gjson.Result{}, nil, notIssueKeyOrIDError
	}

	params := url.Values{}

	if overrideEditableFlag {
		params.Add("overrideEditableFlag", "true")
	}

	if overrideScreenSecurity {
		params.Add("overrideScreenSecurity", "true")
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/3/issue/%v/editmeta", issueKeyOrID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return
	}

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return gjson.ParseBytes(response.Bytes.Bytes()), response, nil
}
