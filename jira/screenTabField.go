package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ScreenTabFieldService struct{ client *Client }

type ScreenTabFieldScheme struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Returns all fields for a screen tab.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/tabs/fields#get-all-screen-tab-fields
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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/tabs/fields#add-screen-tab-field
func (s *ScreenTabFieldService) Add(ctx context.Context, screenID, tabID int, fieldID string) (result *ScreenTabFieldScheme, response *Response, err error) {

	if len(fieldID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide valid fieldID value")
	}

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
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/tabs/fields#remove-screen-tab-field
func (s *ScreenTabFieldService) Remove(ctx context.Context, screenID, tabID int, fieldID string) (response *Response, err error) {

	if len(fieldID) == 0 {
		return nil, fmt.Errorf("error, please provide valid fieldID value")
	}

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
