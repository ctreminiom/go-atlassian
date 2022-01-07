package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type FieldConfigurationService struct {
	client *Client
	Item   *FieldConfigurationItemService
	Scheme *FieldConfigurationSchemeService
}

// Gets Returns a paginated list of all field configurations.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-all-field-configurations
func (f *FieldConfigurationService) Gets(ctx context.Context, ids []int, isDefault bool, startAt, maxResults int) (
	result *models.FieldConfigurationPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if isDefault {
		params.Add("isDefault", "true")
	}

	for _, id := range ids {
		params.Add("id", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/3/fieldconfiguration?%v", params.Encode())

	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Create creates a field configuration. The field configuration is created with the same field properties as the
// default configuration, with all the fields being optional.
// This operation can only create configurations for use in company-managed (classic) projects.
// EXPERIMENTAL
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#create-field-configuration
func (f *FieldConfigurationService) Create(ctx context.Context, name, description string) (result *models.FieldConfigurationScheme,
	response *ResponseScheme, err error) {

	if name == "" {
		return nil, nil, models.ErrNoFieldConfigurationNameError
	}

	payload := struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}
	payloadAsReader, _ := transformStructToReader(&payload)
	endpoint := "rest/api/3/fieldconfiguration"

	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update updates a field configuration. The name and the description provided in the request override the existing values.
// This operation can only update configurations used in company-managed (classic) projects.
// EXPERIMENTAL
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#update-field-configuration
func (f *FieldConfigurationService) Update(ctx context.Context, fieldConfigurationID int, name, description string) (
	response *ResponseScheme, err error) {

	if fieldConfigurationID == 0 {
		return nil, models.ErrNoFieldConfigurationIDError
	}

	if name == "" {
		return nil, models.ErrNoFieldConfigurationNameError
	}

	payload := struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	endpoint := fmt.Sprintf("rest/api/3/fieldconfiguration/%v", fieldConfigurationID)

	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Delete deletes a field configuration.
// This operation can only delete configurations used in company-managed (classic) projects.
// EXPERIMENTAL
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#delete-field-configuration
func (f *FieldConfigurationService) Delete(ctx context.Context, fieldConfigurationID int) (response *ResponseScheme, err error) {

	if fieldConfigurationID == 0 {
		return nil, models.ErrNoFieldConfigurationIDError
	}

	endpoint := fmt.Sprintf("rest/api/3/fieldconfiguration/%v", fieldConfigurationID)

	request, err := f.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.call(request, nil)
	if err != nil {
		return
	}

	return
}
