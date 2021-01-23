package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type IssueSecuritySchemeService struct{ client *Client }

type IssueSecuritySchemesScheme struct {
	IssueSecuritySchemes []IssueSecurityScheme `json:"issueSecuritySchemes"`
}

// Returns all issue security schemes.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-security-schemes/#api-rest-api-3-issuesecurityschemes-get
func (i *IssueSecuritySchemeService) Gets(ctx context.Context) (result *IssueSecuritySchemesScheme, response *Response, err error) {

	var endpoint = "rest/api/3/issuesecurityschemes"

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueSecuritySchemesScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type IssueSecurityScheme struct {
	Self                   string `json:"self"`
	ID                     int    `json:"id"`
	Name                   string `json:"name"`
	Description            string `json:"description"`
	DefaultSecurityLevelID int    `json:"defaultSecurityLevelId"`
	Levels                 []struct {
		Self        string `json:"self"`
		ID          string `json:"id"`
		Description string `json:"description"`
		Name        string `json:"name"`
	} `json:"levels"`
}

// Returns an issue security scheme along with its security levels.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-security-schemes/#api-rest-api-3-issuesecurityschemes-id-get
func (i *IssueSecuritySchemeService) Get(ctx context.Context, issueSecuritySchemeID int) (result *IssueSecurityScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issuesecurityschemes/%v", issueSecuritySchemeID)

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueSecurityScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
