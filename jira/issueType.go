package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type IssueTypeService struct {
	client *Client
	Scheme *IssueTypeSchemeService
}

type IssueTypeScheme struct {
	Self        string `json:"self"`
	ID          string `json:"id"`
	Description string `json:"description"`
	IconURL     string `json:"iconUrl"`
	Name        string `json:"name"`
	Subtask     bool   `json:"subtask"`
	AvatarID    int    `json:"avatarId"`
	EntityID    string `json:"entityId,omitempty"`
	Scope       struct {
		Type    string `json:"type"`
		Project struct {
			ID   string `json:"id"`
			Key  string `json:"key"`
			Name string `json:"name"`
		} `json:"project"`
	} `json:"scope,omitempty"`
}

// Returns all issue types.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-types/#api-rest-api-3-issuetype-get
func (i *IssueTypeService) Gets(ctx context.Context) (result *[]IssueTypeScheme, response *Response, err error) {

	var endpoint = "rest/api/3/issuetype"

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new([]IssueTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type IssueTypePayloadScheme struct {
	Name        string `json:"name,omitempty" validate:"required"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty" validate:"oneof= 'subtask' 'standard'"`
}

// Creates an issue type and adds it to the default issue type scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-types/#api-rest-api-3-issuetype-post
func (i *IssueTypeService) Create(ctx context.Context, payload *IssueTypePayloadScheme) (result *IssueTypeScheme, response *Response, err error) {

	validate := validator.New()
	if err = validate.Struct(payload); err != nil {
		return
	}

	var endpoint = "rest/api/3/issuetype"

	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (i *IssueTypeService) Get(ctx context.Context, issueTypeID string) (result *IssueTypeScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issuetype/%v", issueTypeID)

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Updates the issue type.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-types/#api-rest-api-3-issuetype-id-put
func (i *IssueTypeService) Update(ctx context.Context, issueTypeID string, payload *IssueTypePayloadScheme) (result *IssueTypeScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issuetype/%v", issueTypeID)

	request, err := i.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes the issue type.
// If the issue type is in use, all uses are updated with the alternative issue type (alternativeIssueTypeId).
// A list of alternative issue types are obtained from the Get alternative issue types resource.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-types/#api-rest-api-3-issuetype-id-delete
func (i *IssueTypeService) Delete(ctx context.Context, issueTypeID string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issuetype/%v", issueTypeID)

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

// Returns a list of issue types that can be used to replace the issue type.
// The alternative issue types are those assigned to the same workflow scheme, field configuration scheme, and screen scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-types/#api-rest-api-3-issuetype-id-alternatives-get
func (i *IssueTypeService) Alternatives(ctx context.Context, issueTypeID string) (result *[]IssueTypeScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issuetype/%v/alternatives", issueTypeID)

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new([]IssueTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
