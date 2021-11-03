package v2

import (
	"context"
	"fmt"
	models "github.com/ctreminiom/go-atlassian/pkg/infra/models/jira"
	"net/http"
)

type ScreenTabFieldService struct{ client *Client }

// Gets returns all fields for a screen tab.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/tabs/fields#get-all-screen-tab-fields
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screen-tab-fields/#api-rest-api-2-screens-screenid-tabs-tabid-fields-get
func (s *ScreenTabFieldService) Gets(ctx context.Context, screenID, tabID int) (result []*models.ScreenTabFieldScheme,
	response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/api/2/screens/%v/tabs/%v/fields", screenID, tabID)

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

// Add adds a field to a screen tab.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/tabs/fields#add-screen-tab-field
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screen-tab-fields/#api-rest-api-2-screens-screenid-tabs-tabid-fields-post
func (s *ScreenTabFieldService) Add(ctx context.Context, screenID, tabID int, fieldID string) (result *models.ScreenTabFieldScheme,
	response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, nil, models.ErrNoFieldIDError
	}

	payload := struct {
		FieldID string `json:"fieldId"`
	}{
		FieldID: fieldID,
	}

	var endpoint = fmt.Sprintf("rest/api/2/screens/%v/tabs/%v/fields", screenID, tabID)

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

// Remove removes a field from a screen tab.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/tabs/fields#remove-screen-tab-field
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screen-tab-fields/#api-rest-api-2-screens-screenid-tabs-tabid-fields-id-delete
func (s *ScreenTabFieldService) Remove(ctx context.Context, screenID, tabID int, fieldID string) (
	response *ResponseScheme, err error) {

	if len(fieldID) == 0 {
		return nil, models.ErrNoFieldIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/screens/%v/tabs/%v/fields/%v", screenID, tabID, fieldID)

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
