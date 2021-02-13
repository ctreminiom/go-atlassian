package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type IssueLinkTypeService struct{ client *Client }

type IssueLinkTypeSearchScheme struct {
	IssueLinkTypes []IssueLinkTypeScheme `json:"issueLinkTypes"`
}

type IssueLinkTypeScheme struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Inward  string `json:"inward"`
	Outward string `json:"outward"`
	Self    string `json:"self"`
}

// Returns a list of all issue link types.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-link-types/#api-rest-api-3-issuelinktype-get
func (i *IssueLinkTypeService) Gets(ctx context.Context) (result *IssueLinkTypeSearchScheme, response *Response, err error) {

	var endpoint = "rest/api/3/issueLinkType"

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueLinkTypeSearchScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Returns an issue link type.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-link-types/#api-rest-api-3-issuelinktype-issuelinktypeid-get
func (i *IssueLinkTypeService) Get(ctx context.Context, issueLinkTypeID string) (result *IssueLinkTypeScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issueLinkType/%v", issueLinkTypeID)

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueLinkTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

type IssueLinkTypePayloadScheme struct {
	Inward  string `json:"inward,omitempty" validate:"required"`
	Name    string `json:"name,omitempty" validate:"required"`
	Outward string `json:"outward,omitempty" validate:"required"`
}

// Creates an issue link type.
// Use this operation to create descriptions of the reasons why issues are linked.
// The issue link type consists of a name and descriptions for a link's inward and outward relationships.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-link-types/#api-rest-api-3-issuelinktype-post
func (i *IssueLinkTypeService) Create(ctx context.Context, payload *IssueLinkTypePayloadScheme) (result *IssueLinkTypeScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, payload value is nil, please provide a valid IssueLinkTypePayloadScheme pointer")
	}

	validate := validator.New()
	if err = validate.Struct(payload); err != nil {
		return
	}

	var endpoint = "rest/api/3/issueLinkType"

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

	result = new(IssueLinkTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Updates an issue link type.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-link-types/#api-rest-api-3-issuelinktype-issuelinktypeid-put
func (i *IssueLinkTypeService) Update(ctx context.Context, issueLinkTypeID string, payload *IssueLinkTypePayloadScheme) (result *IssueLinkTypeScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, payload value is nil, please provide a valid IssueLinkTypePayloadScheme pointer")
	}

	var endpoint = fmt.Sprintf("rest/api/3/issueLinkType/%v", issueLinkTypeID)

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

	result = new(IssueLinkTypeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

// Deletes an issue link type.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-link-types/#api-rest-api-3-issuelinktype-issuelinktypeid-delete
func (i *IssueLinkTypeService) Delete(ctx context.Context, issueLinkTypeID string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issueLinkType/%v", issueLinkTypeID)

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
