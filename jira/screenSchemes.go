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

type ScreenSchemeService struct{ client *Client }

type ScreenSchemePageScheme struct {
	Self       string               `json:"self"`
	MaxResults int                  `json:"maxResults"`
	StartAt    int                  `json:"startAt"`
	Total      int                  `json:"total"`
	IsLast     bool                 `json:"isLast"`
	Values     []ScreenSchemeScheme `json:"values"`
}

type ScreenSchemeScheme struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty" validate:"required"`
	Description string `json:"description,omitempty"`
	Screens     struct {
		Default int `json:"default,omitempty" validate:"required"`
		View    int `json:"view,omitempty" validate:"required"`
		Edit    int `json:"edit,omitempty" validate:"required"`
		Create  int `json:"create,omitempty"`
	} `json:"screens,omitempty"`
}

// Returns a paginated list of screen schemes.
// Only screen schemes used in classic projects are returned.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-schemes/#api-rest-api-3-screenscheme-get
func (s *ScreenSchemeService) Gets(ctx context.Context, screenSchemeIDs []int, startAt, maxResults int) (result *ScreenSchemePageScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, schemeScheme := range screenSchemeIDs {
		params.Add("id", strconv.Itoa(schemeScheme))
	}

	var endpoint = fmt.Sprintf("rest/api/3/screenscheme?%v", params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ScreenSchemePageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Creates a screen scheme.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-schemes/#api-rest-api-3-screenscheme-post
func (s *ScreenSchemeService) Create(ctx context.Context, payload *ScreenSchemeScheme) (result *ScreenSchemeScheme, response *Response, err error) {

	validate := validator.New()
	if err = validate.Struct(payload); err != nil {
		return
	}

	var endpoint = "rest/api/3/screenscheme"

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

	result = new(ScreenSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Updates a screen scheme. Only screen schemes used in classic projects can be updated.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-schemes/#api-rest-api-3-screenscheme-screenschemeid-put
func (s *ScreenSchemeService) Update(ctx context.Context, screenSchemeID string, payload *ScreenSchemeScheme) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/screenscheme/%v", screenSchemeID)

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

	return
}

// Deletes a screen scheme.
// A screen scheme cannot be deleted if it is used in an issue type screen scheme.
// Only screens schemes used in classic projects can be deleted.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-screen-schemes/#api-rest-api-3-screenscheme-screenschemeid-delete
func (s *ScreenSchemeService) Delete(ctx context.Context, screenSchemeID string) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/screenscheme/%v", screenSchemeID)

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
