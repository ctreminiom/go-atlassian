package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type ScreenTabFieldService struct{ client *Client }

type ScreenTabFieldScheme struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Returns all fields for a screen tab.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-tab-fields/#api-rest-api-3-screens-screenid-tabs-tabid-fields-get
func (s *ScreenTabFieldService) Gets(ctx context.Context, screenID, tabID int) (result *[]ScreenTabFieldScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/screens/%v/tabs/%v/fields", screenID, tabID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new([]ScreenTabFieldScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Adds a field to a screen tab.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-tab-fields/#api-rest-api-3-screens-screenid-tabs-tabid-fields-post
func (s *ScreenTabFieldService) Add(ctx context.Context, screenID, tabID int, fieldID string) (result *ScreenTabFieldScheme, response *Response, err error) {

	payload := struct {
		FieldID string `json:"fieldId"`
	}{FieldID: fieldID}

	var endpoint = fmt.Sprintf("rest/api/3/screens/%v/tabs/%v/fields", screenID, tabID)

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

	result = new(ScreenTabFieldScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Removes a field from a screen tab.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-tab-fields/#api-rest-api-3-screens-screenid-tabs-tabid-fields-id-delete
func (s *ScreenTabFieldService) Remove(ctx context.Context, screenID, tabID int, fieldID string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/screens/%v/tabs/%v/fields/%v", screenID, tabID, fieldID)

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

type ScreenTabFieldMoveOptions struct {
	ScreenID int    `validate:"required"`
	TabID    int    `validate:"required"`
	FieldID  string `validate:"required"`
	After    string
	Position string `validate:"oneof= 'Earlier' 'Later' 'First' 'Last'"`
}

// Moves a screen tab field.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-tab-fields/#api-rest-api-3-screens-screenid-tabs-tabid-fields-id-move-post
func (s *ScreenTabFieldService) Move(ctx context.Context, options *ScreenTabFieldMoveOptions) (response *Response, err error) {

	validate := validator.New()
	if err = validate.Struct(options); err != nil {
		return
	}

	var endpoint = fmt.Sprintf("rest/api/3/screens/%v/tabs/%v/fields/%v/move", options.ScreenID, options.TabID, options.FieldID)

	payload := struct {
		After    string `json:"after,omitempty"`
		Position string `json:"position,omitempty"`
	}{
		After:    options.After,
		Position: options.Position,
	}

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

	return
}
