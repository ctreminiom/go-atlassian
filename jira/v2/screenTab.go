package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strings"
)

type ScreenTabService struct {
	client *Client
	Field  *ScreenTabFieldService
}

// Gets returns the list of tabs for a screen.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/tabs#get-all-screen-tabs
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screen-tabs/#api-rest-api-2-screens-screenid-tabs-get
func (s *ScreenTabService) Gets(ctx context.Context, screenID int, projectKey string) (result []*models.ScreenTabScheme,
	response *ResponseScheme, err error) {

	params := url.Values{}
	if len(projectKey) != 0 {
		params.Add("projectKey", projectKey)
	}

	var endpoint strings.Builder
	endpoint.WriteString(fmt.Sprintf("rest/api/2/screens/%v/tabs", screenID))

	if params.Encode() != "" {
		endpoint.WriteString(fmt.Sprintf("?%v", params.Encode()))
	}

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint.String(), nil)
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

// Create creates a tab for a screen.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/tabs#create-screen-tab
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screen-tabs/#api-rest-api-2-screens-screenid-tabs-post
func (s *ScreenTabService) Create(ctx context.Context, screenID int, tabName string) (result *models.ScreenTabScheme,
	response *ResponseScheme, err error) {

	if len(tabName) == 0 {
		return nil, nil, models.ErrNoScreenTabNameError
	}

	payload := struct {
		Name string `json:"name"`
	}{
		Name: tabName,
	}

	var endpoint = fmt.Sprintf("rest/api/2/screens/%v/tabs", screenID)

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

// Update updates the name of a screen tab.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/tabs#update-screen-tab
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screen-tabs/#api-rest-api-2-screens-screenid-tabs-tabid-put
func (s *ScreenTabService) Update(ctx context.Context, screenID, tabID int, newTabName string) (result *models.ScreenTabScheme,
	response *ResponseScheme, err error) {

	if len(newTabName) == 0 {
		return nil, nil, models.ErrNoScreenTabNameError
	}

	payload := struct {
		Name string `json:"name"`
	}{
		Name: newTabName,
	}

	var endpoint = fmt.Sprintf("rest/api/2/screens/%v/tabs/%v", screenID, tabID)

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

// Delete deletes a screen tab.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/tabs#delete-screen-tab
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screen-tabs/#api-rest-api-2-screens-screenid-tabs-tabid-delete
func (s *ScreenTabService) Delete(ctx context.Context, screenID, tabID int) (response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/api/2/screens/%v/tabs/%v", screenID, tabID)

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

// Move moves a screen tab.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/tabs#move-screen-tab
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screen-tabs/#api-rest-api-2-screens-screenid-tabs-tabid-move-pos-post
func (s *ScreenTabService) Move(ctx context.Context, screenID, tabID, tabPosition int) (response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/api/2/screens/%v/tabs/%v/move/%v", screenID, tabID, tabPosition)

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
