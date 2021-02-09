package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type IssueTypeSchemeService struct{ client *Client }

type IssueTypeSchemesScheme struct {
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		ID                 string `json:"id"`
		Name               string `json:"name"`
		Description        string `json:"description"`
		DefaultIssueTypeID string `json:"defaultIssueTypeId,omitempty"`
		IsDefault          bool   `json:"isDefault,omitempty"`
	} `json:"values"`
}

// Returns a paginated list of issue type schemes.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-schemes/#api-rest-api-3-issuetypescheme-get
func (i *IssueTypeSchemeService) Gets(ctx context.Context, issueTypeSchemeIDs []int, startAt, maxResults int) (result *IssueTypeSchemesScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range issueTypeSchemeIDs {
		params.Add("id", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescheme?%v", params.Encode())

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueTypeSchemesScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type IssueTypeSchemePayloadScheme struct {
	DefaultIssueTypeID string   `json:"defaultIssueTypeId,omitempty"`
	IssueTypeIds       []string `json:"issueTypeIds,omitempty"`
	Name               string   `json:"name,omitempty"`
	Description        string   `json:"description,omitempty"`
}

type newIssueTypeSchemeScheme struct {
	IssueTypeSchemeID string `json:"issueTypeSchemeId"`
}

// Creates an issue type scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-schemes/#api-rest-api-3-issuetypescheme-post
func (i *IssueTypeSchemeService) Create(ctx context.Context, payload *IssueTypeSchemePayloadScheme) (issueTypeSchemeID string, response *Response, err error) {

	var endpoint = "rest/api/3/issuetypescheme"

	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	responsePayload := new(newIssueTypeSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &responsePayload); err != nil {
		return
	}

	issueTypeSchemeID = responsePayload.IssueTypeSchemeID
	return
}

type IssueTypeSchemeItemsScheme struct {
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		IssueTypeSchemeID string `json:"issueTypeSchemeId"`
		IssueTypeID       string `json:"issueTypeId"`
	} `json:"values"`
}

// Returns a paginated list of issue type scheme items.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-schemes/#api-rest-api-3-issuetypescheme-mapping-get
func (i *IssueTypeSchemeService) Items(ctx context.Context, issueTypeSchemeIDs []int, startAt, maxResults int) (result *IssueTypeSchemeItemsScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range issueTypeSchemeIDs {
		params.Add("issueTypeSchemeId", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescheme/mapping?%v", params.Encode())

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueTypeSchemeItemsScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ProjectIssueTypeSchemeScheme struct {
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		IssueTypeScheme struct {
			ID                 string `json:"id"`
			Name               string `json:"name"`
			Description        string `json:"description"`
			DefaultIssueTypeID string `json:"defaultIssueTypeId"`
			IsDefault          bool   `json:"isDefault"`
		} `json:"issueTypeScheme"`
		ProjectIds []string `json:"projectIds"`
	} `json:"values"`
}

// Returns a paginated list of issue type schemes and, for each issue type scheme, a list of the projects that use it.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-schemes/#api-rest-api-3-issuetypescheme-project-get
func (i *IssueTypeSchemeService) Projects(ctx context.Context, projectIDs []int, startAt, maxResults int) (result *ProjectIssueTypeSchemeScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range projectIDs {
		params.Add("projectId", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescheme/project?%v", params.Encode())

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(ProjectIssueTypeSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Assigns an issue type scheme to a project.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-schemes/#api-rest-api-3-issuetypescheme-project-put
func (i *IssueTypeSchemeService) Assign(ctx context.Context, issueTypeSchemeID, projectID int) (response *Response, err error) {

	payload := struct {
		IssueTypeSchemeID string `json:"issueTypeSchemeId"`
		ProjectID         string `json:"projectId"`
	}{
		IssueTypeSchemeID: strconv.Itoa(issueTypeSchemeID),
		ProjectID:         strconv.Itoa(projectID),
	}

	var endpoint = "rest/api/3/issuetypescheme/project"

	request, err := i.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Updates an issue type scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-schemes/#api-rest-api-3-issuetypescheme-issuetypeschemeid-put
func (i *IssueTypeSchemeService) Update(ctx context.Context, issueTypeSchemeID int, payload *IssueTypeSchemePayloadScheme) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescheme/%v", issueTypeSchemeID)

	request, err := i.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Deletes an issue type scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-schemes/#api-rest-api-3-issuetypescheme-issuetypeschemeid-delete
func (i *IssueTypeSchemeService) Delete(ctx context.Context, issueTypeSchemeID int) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescheme/%v", issueTypeSchemeID)

	request, err := i.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Adds issue types to an issue type scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-schemes/#api-rest-api-3-issuetypescheme-issuetypeschemeid-issuetype-put
func (i *IssueTypeSchemeService) AddIssueTypes(ctx context.Context, issueTypeSchemeID int, issueTypeIDs []int) (response *Response, err error) {

	var issueTypesIDsAsStringSlice []string
	for _, issueTypeID := range issueTypeIDs {
		issueTypesIDsAsStringSlice = append(issueTypesIDsAsStringSlice, strconv.Itoa(issueTypeID))
	}

	payload := struct {
		IssueTypeIds []string `json:"issueTypeIds"`
	}{
		IssueTypeIds: issueTypesIDsAsStringSlice,
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescheme/%v/issuetype", issueTypeSchemeID)

	request, err := i.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Removes an issue type from an issue type scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-schemes/#api-rest-api-3-issuetypescheme-issuetypeschemeid-issuetype-issuetypeid-delete
func (i *IssueTypeSchemeService) RemoveIssueType(ctx context.Context, issueTypeSchemeID, issueTypeID int) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescheme/%v/issuetype/%v", issueTypeSchemeID, issueTypeID)

	request, err := i.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	return
}
