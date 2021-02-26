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
	Self        string                `json:"self,omitempty"`
	ID          string                `json:"id,omitempty"`
	Description string                `json:"description,omitempty"`
	IconURL     string                `json:"iconUrl,omitempty"`
	Name        string                `json:"name,omitempty"`
	Subtask     bool                  `json:"subtask,omitempty"`
	AvatarID    int                   `json:"avatarId,omitempty"`
	EntityID    string                `json:"entityId,omitempty"`
	Scope       *IssueTypeScopeScheme `json:"scope,omitempty"`
}

type IssueTypeScopeScheme struct {
	Type    string                       `json:"type,omitempty"`
	Project *IssueTypeScopeProjectScheme `json:"project,omitempty"`
}

type IssueTypeScopeProjectScheme struct {
	ID   string `json:"id,omitempty"`
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
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
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

type IssueTypePayloadScheme struct {
	Name        string `json:"name,omitempty" validate:"required"`
	Description string `json:"description,omitempty"`
	Type        string `json:"type,omitempty" validate:"required,oneof=subtask standard"`
}

// Creates an issue type and adds it to the default issue type scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-types/#api-rest-api-3-issuetype-post
func (i *IssueTypeService) Create(ctx context.Context, payload *IssueTypePayloadScheme) (result *IssueTypeScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, payload value is nil, please provide a valid IssueTypePayloadScheme pointer")
	}

	validate := validator.New()
	if err = validate.Struct(payload); err != nil {
		err = fmt.Errorf("error: issuetype type payload invalid: %v", err.Error())
		return
	}

	var endpoint = "rest/api/3/issuetype"

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

	result = new(IssueTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())

	}

	return
}

// Returns an issue type.
// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-types/#api-rest-api-3-issuetype-id-get
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
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())

	}
	return
}

// Updates the issue type.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-types/#api-rest-api-3-issuetype-id-put
func (i *IssueTypeService) Update(ctx context.Context, issueTypeID string, payload *IssueTypePayloadScheme) (result *IssueTypeScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, payload value is nil, please provide a valid IssueTypePayloadScheme pointer")
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetype/%v", issueTypeID)

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

	result = new(IssueTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())

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
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}
