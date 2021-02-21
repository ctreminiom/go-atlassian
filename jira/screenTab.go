package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type ScreenTabService struct {
	client *Client
	Field  *ScreenTabFieldService
}

type ScreenTabScheme struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Returns the list of tabs for a screen.
// https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-tabs/#api-rest-api-3-screens-screenid-tabs-get
func (s *ScreenTabService) Gets(ctx context.Context, screenID int, projectKey string) (result *[]ScreenTabScheme, response *Response, err error) {

	params := url.Values{}

	if len(projectKey) != 0 {
		params.Add("projectKey", projectKey)
	}

	var endpoint string
	if len(params.Encode()) != 0 {
		endpoint = fmt.Sprintf("rest/api/3/screens/%v/tabs?%v", screenID, params.Encode())
	} else {
		endpoint = fmt.Sprintf("rest/api/3/screens/%v/tabs", screenID)
	}

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new([]ScreenTabScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Creates a tab for a screen.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-tabs/#api-rest-api-3-screens-screenid-tabs-post
func (s *ScreenTabService) Create(ctx context.Context, screenID int, tabName string) (result *ScreenTabScheme, response *Response, err error) {

	if len(tabName) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid tabName value")
	}

	payload := struct {
		Name string `json:"name"`
	}{Name: tabName}

	var endpoint = fmt.Sprintf("rest/api/3/screens/%v/tabs", screenID)

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

	result = new(ScreenTabScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Updates the name of a screen tab.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-tabs/#api-rest-api-3-screens-screenid-tabs-tabid-put
func (s *ScreenTabService) Update(ctx context.Context, screenID, tabID int, newTabName string) (result *ScreenTabScheme, response *Response, err error) {

	if len(newTabName) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid newTabName value")
	}

	payload := struct {
		Name string `json:"name"`
	}{Name: newTabName}

	var endpoint = fmt.Sprintf("rest/api/3/screens/%v/tabs/%v", screenID, tabID)

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

	result = new(ScreenTabScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes a screen tab.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-tabs/#api-rest-api-3-screens-screenid-tabs-tabid-delete
func (s *ScreenTabService) Delete(ctx context.Context, screenID, tabID int) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/screens/%v/tabs/%v", screenID, tabID)

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

// Moves a screen tab.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-tabs/#api-rest-api-3-screens-screenid-tabs-tabid-move-pos-post
func (s *ScreenTabService) Move(ctx context.Context, screenID, tabID, tabPosition int) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/screens/%v/tabs/%v/move/%v", screenID, tabID, tabPosition)

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
