package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type ScreenSchemeService struct{ client *Client }

// Gets returns a paginated list of screen schemes.
// Only screen schemes used in classic projects are returned.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#get-screen-schemes
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screen-schemes/#api-rest-api-2-screenscheme-get
func (s *ScreenSchemeService) Gets(ctx context.Context, screenSchemeIDs []int, startAt, maxResults int) (
	result *models.ScreenSchemePageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, schemeScheme := range screenSchemeIDs {
		params.Add("id", strconv.Itoa(schemeScheme))
	}

	var endpoint = fmt.Sprintf("rest/api/2/screenscheme?%v", params.Encode())

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

// Create creates a screen scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#create-screen-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screen-schemes/#api-rest-api-2-screenscheme-post
func (s *ScreenSchemeService) Create(ctx context.Context, payload *models.ScreenSchemePayloadScheme) (result *models.ScreenSchemeScheme,
	response *ResponseScheme, err error) {

	var endpoint = "rest/api/2/screenscheme"

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

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

// Update updates a screen scheme. Only screen schemes used in classic projects can be updated.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#update-screen-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screen-schemes/#api-rest-api-2-screenscheme-screenschemeid-put
func (s *ScreenSchemeService) Update(ctx context.Context, screenSchemeID string, payload *models.ScreenSchemePayloadScheme) (
	response *ResponseScheme, err error) {

	if len(screenSchemeID) == 0 {
		return nil, models.ErrNoScreenSchemeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/screenscheme/%v", screenSchemeID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	request, err := s.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Delete deletes a screen scheme.
// A screen scheme cannot be deleted if it is used in an issue type screen scheme.
// Only screens schemes used in classic projects can be deleted.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#delete-screen-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-screen-schemes/#api-rest-api-2-screenscheme-screenschemeid-delete
func (s *ScreenSchemeService) Delete(ctx context.Context, screenSchemeID string) (response *ResponseScheme, err error) {

	if len(screenSchemeID) == 0 {
		return nil, models.ErrNoScreenSchemeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/screenscheme/%v", screenSchemeID)

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
