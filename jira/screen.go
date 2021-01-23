package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type ScreenService struct {
	client *Client
	Tab    *ScreenTabService
	Scheme *ScreenSchemeService
}

type ScreenScheme struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ScreenFieldSearchScheme struct {
	MaxResults int            `json:"maxResults"`
	StartAt    int            `json:"startAt"`
	Total      int            `json:"total"`
	IsLast     bool           `json:"isLast"`
	Values     []ScreenScheme `json:"values"`
}

// Returns a paginated list of the screens a field is used in.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screens/#api-rest-api-3-field-fieldid-screens-get
func (s *ScreenService) Get(ctx context.Context, fieldID string, startAt, maxResults int) (result *ScreenFieldSearchScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/screens?%v", fieldID, params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ScreenFieldSearchScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ScreenSearchPageScheme struct {
	Self       string         `json:"self"`
	MaxResults int            `json:"maxResults"`
	StartAt    int            `json:"startAt"`
	Total      int            `json:"total"`
	IsLast     bool           `json:"isLast"`
	Values     []ScreenScheme `json:"values"`
}

// Returns a paginated list of all screens or those specified by one or more screen IDs.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screens/#api-rest-api-3-screens-get
func (s *ScreenService) Gets(ctx context.Context, screenIDs []int, startAt, maxResults int) (result *ScreenSearchPageScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, screenID := range screenIDs {
		params.Add("id", strconv.Itoa(screenID))
	}

	var endpoint = fmt.Sprintf("rest/api/3/screens?%v", params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ScreenSearchPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Creates a screen with a default field tab.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screens/#api-rest-api-3-screens-post
func (s *ScreenService) Create(ctx context.Context, name, description string) (result *ScreenScheme, response *Response, err error) {

	payload := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}{
		Name:        name,
		Description: description,
	}

	var endpoint = "rest/api/3/screens"

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ScreenScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

func (s *ScreenService) AddField(ctx context.Context, fieldID string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/screens/addToDefault/%v", fieldID)

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Updates a screen. Only screens used in classic projects can be updated.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screens/#api-rest-api-3-screens-screenid-put
func (s *ScreenService) Update(ctx context.Context, screenID, name, description string) (result *ScreenScheme, response *Response, err error) {

	payload := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}{
		Name:        name,
		Description: description,
	}

	var endpoint = fmt.Sprintf("rest/api/3/screens/%v", screenID)

	request, err := s.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ScreenScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes a screen.
// A screen cannot be deleted if it is used in a screen scheme,
// workflow, or workflow draft. Only screens used in classic projects can be deleted.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screens/#api-rest-api-3-screens-screenid-delete
func (s *ScreenService) Delete(ctx context.Context, screenID string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/screens/%v", screenID)

	request, err := s.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	return
}

type AvailableScreenFieldScheme struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (s *ScreenService) Available(ctx context.Context, screenID string) (result *[]AvailableScreenFieldScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/screens/%v/availableFields", screenID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new([]AvailableScreenFieldScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
