package v2

import (
	"context"
	"fmt"
	models2 "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type ScreenService struct {
	client *Client
	Tab    *ScreenTabService
	Scheme *ScreenSchemeService
}

// Fields returns a paginated list of the screens a field is used in.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens#get-screens-for-a-field
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screens/#api-rest-api-2-field-fieldid-screens-get
func (s *ScreenService) Fields(ctx context.Context, fieldID string, startAt, maxResults int) (result *models2.ScreenFieldPageScheme,
	response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, nil, models2.ErrNoFieldIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("rest/api/2/field/%v/screens?%v", fieldID, params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Gets returns a paginated list of all screens or those specified by one or more screen IDs.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens#get-screens
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screens/#api-rest-api-2-screens-get
func (s *ScreenService) Gets(ctx context.Context, screenIDs []int, startAt, maxResults int) (result *models2.ScreenSearchPageScheme,
	response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, screenID := range screenIDs {
		params.Add("id", strconv.Itoa(screenID))
	}

	var endpoint = fmt.Sprintf("rest/api/2/screens?%v", params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Create creates a screen with a default field tab.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens#create-screen
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screens/#api-rest-api-2-screens-post
func (s *ScreenService) Create(ctx context.Context, name, description string) (result *models2.ScreenScheme,
	response *ResponseScheme, err error) {

	if len(name) == 0 {
		return nil, nil, models2.ErrNoScreenNameError
	}

	payload := struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	var endpoint = "rest/api/2/screens"

	payloadAsReader, _ := transformStructToReader(&payload)
	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// AddToDefault adds a field to the default tab of the default screen.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens#add-field-to-default-screen
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screens/#api-rest-api-2-screens-addtodefault-fieldid-post
func (s *ScreenService) AddToDefault(ctx context.Context, fieldID string) (response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, models2.ErrNoFieldIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/screens/addToDefault/%v", fieldID)

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Update updates a screen. Only screens used in classic projects can be updated.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens#update-screen
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screens/#api-rest-api-2-screens-screenid-put
func (s *ScreenService) Update(ctx context.Context, screenID int, name, description string) (result *models2.ScreenScheme,
	response *ResponseScheme, err error) {

	payload := struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	var endpoint = fmt.Sprintf("rest/api/2/screens/%v", screenID)

	payloadAsReader, _ := transformStructToReader(&payload)
	request, err := s.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes a screen.
// A screen cannot be deleted if it is used in a screen scheme,
// workflow, or workflow draft. Only screens used in classic projects can be deleted.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens#delete-screen
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screens/#api-rest-api-2-screens-screenid-delete
func (s *ScreenService) Delete(ctx context.Context, screenID int) (response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/api/2/screens/%v", screenID)

	request, err := s.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = s.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Available returns the fields that can be added to a tab on a screen.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens#get-available-screen-fields
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screens/#api-rest-api-2-screens-screenid-availablefields-get
func (s *ScreenService) Available(ctx context.Context, screenID int) (result []*models2.AvailableScreenFieldScheme,
	response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/api/2/screens/%v/availableFields", screenID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
