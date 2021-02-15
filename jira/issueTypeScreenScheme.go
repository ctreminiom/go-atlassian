package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"net/url"
	"strconv"
)

type IssueTypeScreenSchemeService struct{ client *Client }

type IssueTypeScreenSchemesScheme struct {
	MaxResults int  `json:"maxResults"`
	StartAt    int  `json:"startAt"`
	Total      int  `json:"total"`
	IsLast     bool `json:"isLast"`
	Values     []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"values"`
}

// Returns a paginated list of issue type screen schemes.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-get
func (i *IssueTypeScreenSchemeService) Gets(ctx context.Context, ids []int, startAt, maxResults int) (result *IssueTypeScreenSchemesScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range ids {
		params.Add("id", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescreenscheme?%v", params.Encode())

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueTypeScreenSchemesScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type IssueTypeScreenSchemePayloadScheme struct {
	Name              string                                      `json:"name" validate:"required"`
	IssueTypeMappings []IssueTypeScreenSchemeMappingPayloadScheme `json:"issueTypeMappings" validate:"required"`
}

type IssueTypeScreenSchemeMappingPayloadScheme struct {
	IssueTypeID    string `json:"issueTypeId"`
	ScreenSchemeID string `json:"screenSchemeId"`
}

type issueTypeScreenScreenCreatedScheme struct {
	ID string `json:"id"`
}

// Creates an issue type screen scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-post
func (i *IssueTypeScreenSchemeService) Create(ctx context.Context, payload *IssueTypeScreenSchemePayloadScheme) (issueTypeScreenSchemeID int, response *Response, err error) {

	if payload == nil {
		return 0, nil, fmt.Errorf("error, payload value is nil, please provide a valid IssueTypeScreenSchemePayloadScheme pointer")
	}

	validate := validator.New()
	if err = validate.Struct(payload); err != nil {
		return
	}

	var endpoint = "rest/api/3/issuetypescreenscheme"

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

	result := new(issueTypeScreenScreenCreatedScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	newIssueTypeScreenSchemeIDAsInt, err := strconv.Atoi(result.ID)
	if err != nil {
		return
	}

	issueTypeScreenSchemeID = newIssueTypeScreenSchemeIDAsInt
	return
}

// Assigns an issue type screen scheme to a project.
// Issue type screen schemes can only be assigned to classic projects.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-project-put
func (i *IssueTypeScreenSchemeService) Assign(ctx context.Context, issueTypeScreenSchemeID, projectID string) (response *Response, err error) {

	if len(issueTypeScreenSchemeID) == 0 {
		return nil, fmt.Errorf("error, please provide a issueTypeScreenSchemeID value")
	}

	if len(projectID) == 0 {
		return nil, fmt.Errorf("error, please provide a projectID value")
	}

	payload := struct {
		IssueTypeScreenSchemeID string `json:"issueTypeScreenSchemeId"`
		ProjectID               string `json:"projectId"`
	}{
		IssueTypeScreenSchemeID: issueTypeScreenSchemeID,
		ProjectID:               projectID,
	}

	var endpoint = "rest/api/3/issuetypescreenscheme/project"

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

// Updates an issue type screen scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-issuetypescreenschemeid-put
func (i *IssueTypeScreenSchemeService) Update(ctx context.Context, issueTypeScreenSchemeID, name, description string) (response *Response, err error) {

	if len(issueTypeScreenSchemeID) == 0 {
		return nil, fmt.Errorf("error, please provide a issueTypeScreenSchemeID value")
	}

	payload := struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescreenscheme/%v", issueTypeScreenSchemeID)

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

// Deletes an issue type screen scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-issuetypescreenschemeid-delete
func (i *IssueTypeScreenSchemeService) Delete(ctx context.Context, issueTypeScreenSchemeID string) (response *Response, err error) {

	if len(issueTypeScreenSchemeID) == 0 {
		return nil, fmt.Errorf("error, please provide a issueTypeScreenSchemeID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescreenscheme/%v", issueTypeScreenSchemeID)

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

// Appends issue type to screen scheme mappings to an issue type screen scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-issuetypescreenschemeid-mapping-put
func (i *IssueTypeScreenSchemeService) Append(ctx context.Context, issueTypeScreenSchemeID string, mappings *[]IssueTypeScreenSchemeMappingPayloadScheme) (response *Response, err error) {

	if len(issueTypeScreenSchemeID) == 0 {
		return nil, fmt.Errorf("error, please provide a issueTypeScreenSchemeID value")
	}

	if mappings == nil {
		return nil, fmt.Errorf("error, payload value is nil, please provide a valid IssueTypeScreenSchemeMappingPayloadScheme slice pointer")
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescreenscheme/%v/mapping", issueTypeScreenSchemeID)

	payload := struct {
		IssueTypeMappings []IssueTypeScreenSchemeMappingPayloadScheme `json:"issueTypeMappings"`
	}{
		IssueTypeMappings: *mappings,
	}

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
