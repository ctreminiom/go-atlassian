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

type IssueTypeSchemePageScheme struct {
	Self       string                   `json:"self,omitempty"`
	NextPage   string                   `json:"nextPage,omitempty"`
	MaxResults int                      `json:"maxResults,omitempty"`
	StartAt    int                      `json:"startAt,omitempty"`
	Total      int                      `json:"total,omitempty"`
	IsLast     bool                     `json:"isLast,omitempty"`
	Values     []*IssueTypeSchemeScheme `json:"values,omitempty"`
}

type IssueTypeSchemeScheme struct {
	ID                 string `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	Description        string `json:"description,omitempty"`
	DefaultIssueTypeID string `json:"defaultIssueTypeId,omitempty"`
	IsDefault          bool   `json:"isDefault,omitempty"`
}

// Returns a paginated list of issue type schemes.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#get-all-issue-type-schemes
func (i *IssueTypeSchemeService) Gets(ctx context.Context, issueTypeSchemeIDs []int, startAt, maxResults int) (result *IssueTypeSchemePageScheme, response *Response, err error) {

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

	result = new(IssueTypeSchemePageScheme)
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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#create-issue-type-scheme
func (i *IssueTypeSchemeService) Create(ctx context.Context, payload *IssueTypeSchemePayloadScheme) (issueTypeSchemeID string, response *Response, err error) {

	if payload == nil {
		return "", nil, fmt.Errorf("error, payload value is nil, please provide a valid IssueTypeSchemePayloadScheme pointer")
	}

	/*
		Validation considerations for the Atlassian Documentation.

		-------------------------
		value: defaultIssueTypeId
		validation: This ID must be included in issueTypeIds.
		-------------------------

		-------------------------
		value: issueTypeIds
		validation: At least one standard issue type ID is required
	*/

	var containsTheIssueType bool
	for _, issueType := range payload.IssueTypeIds {

		// The DefaultIssueTypeID value is not required
		if payload.DefaultIssueTypeID == "" {
			containsTheIssueType = true
			break
		}

		if issueType == payload.DefaultIssueTypeID {
			containsTheIssueType = true
			break
		}
	}

	if !containsTheIssueType {
		return "", nil, fmt.Errorf("error, please add the DefaultIssueTypeID value on the IssueTypeIds value")
	}

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

type IssueTypeSchemeItemPageScheme struct {
	MaxResults int                             `json:"maxResults,omitempty"`
	StartAt    int                             `json:"startAt,omitempty"`
	Total      int                             `json:"total,omitempty"`
	IsLast     bool                            `json:"isLast,omitempty"`
	Values     []*IssueTypeSchemeMappingScheme `json:"values,omitempty"`
}

type IssueTypeSchemeMappingScheme struct {
	IssueTypeSchemeID string `json:"issueTypeSchemeId,omitempty"`
	IssueTypeID       string `json:"issueTypeId,omitempty"`
}

// Returns a paginated list of issue type scheme items.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#get-issue-type-scheme-items
func (i *IssueTypeSchemeService) Items(ctx context.Context, issueTypeSchemeIDs []int, startAt, maxResults int) (result *IssueTypeSchemeItemPageScheme, response *Response, err error) {

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

	result = new(IssueTypeSchemeItemPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ProjectIssueTypeSchemePageScheme struct {
	MaxResults int                              `json:"maxResults"`
	StartAt    int                              `json:"startAt"`
	Total      int                              `json:"total"`
	IsLast     bool                             `json:"isLast"`
	Values     []*IssueTypeSchemeProjectsScheme `json:"values"`
}

type IssueTypeSchemeProjectsScheme struct {
	IssueTypeScheme *IssueTypeSchemeScheme `json:"issueTypeScheme,omitempty"`
	ProjectIds      []string               `json:"projectIds,omitempty"`
}

// Returns a paginated list of issue type schemes and, for each issue type scheme, a list of the projects that use it.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#get-issue-type-schemes-for-projects
func (i *IssueTypeSchemeService) Projects(ctx context.Context, projectIDs []int, startAt, maxResults int) (result *ProjectIssueTypeSchemePageScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(projectIDs) == 0 {
		return nil, nil, fmt.Errorf("error, please provide values on the projectIDs param")
	}

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

	result = new(ProjectIssueTypeSchemePageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Assigns an issue type scheme to a project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#assign-issue-type-scheme-to-project
func (i *IssueTypeSchemeService) Assign(ctx context.Context, issueTypeSchemeID, projectID string) (response *Response, err error) {

	if len(issueTypeSchemeID) == 0 {
		return nil, fmt.Errorf("error!, please provide a valid issueTypeSchemeID value")
	}

	if len(projectID) == 0 {
		return nil, fmt.Errorf("error!, please provide a valid projectID value")
	}

	payload := struct {
		IssueTypeSchemeID string `json:"issueTypeSchemeId"`
		ProjectID         string `json:"projectId"`
	}{
		IssueTypeSchemeID: issueTypeSchemeID,
		ProjectID:         projectID,
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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#update-issue-type-scheme
func (i *IssueTypeSchemeService) Update(ctx context.Context, issueTypeSchemeID int, payload *IssueTypeSchemePayloadScheme) (response *Response, err error) {

	if issueTypeSchemeID == 0 {
		return nil, fmt.Errorf("error!, please provide a valid issueTypeSchemeID value")
	}

	if payload == nil {
		return nil, fmt.Errorf("error, payload value is nil, please provide a valid IssueTypeSchemePayloadScheme pointer")
	}

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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#delete-issue-type-scheme
func (i *IssueTypeSchemeService) Delete(ctx context.Context, issueTypeSchemeID int) (response *Response, err error) {

	if issueTypeSchemeID == 0 {
		return nil, fmt.Errorf("error!, please provide a valid issueTypeSchemeID value")
	}

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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#add-issue-types-to-issue-type-scheme
func (i *IssueTypeSchemeService) AddIssueTypes(ctx context.Context, issueTypeSchemeID int, issueTypeIDs []int) (response *Response, err error) {

	if len(issueTypeIDs) == 0 {
		return nil, fmt.Errorf("error, please provide a issue types ID under the issueTypeIDs param")
	}

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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#remove-issue-type-from-issue-type-scheme
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
