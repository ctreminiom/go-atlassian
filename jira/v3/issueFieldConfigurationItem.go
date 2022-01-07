package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type FieldConfigurationItemService struct{ client *Client }

// Gets Returns a paginated list of all fields for a configuration.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/items#get-field-configuration-items
func (f *FieldConfigurationItemService) Gets(ctx context.Context, fieldConfigurationID, startAt, maxResults int) (
	result *models.FieldConfigurationItemPageScheme, response *ResponseScheme, err error) {

	if fieldConfigurationID == 0 {
		return nil, nil, models.ErrNoFieldConfigurationIDError
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("rest/api/2/fieldconfiguration/%v/fields?%v", fieldConfigurationID, params.Encode())

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

// Update updates fields in a field configuration. The properties of the field configuration fields provided
// override the existing values.
// This operation can only update field configurations used in company-managed (classic) projects.
// EXPERIMENTAL
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration/items#update-field-configuration-items
func (f *FieldConfigurationItemService) Update(ctx context.Context, fieldConfigurationID int, payload *models.UpdateFieldConfigurationItemPayloadScheme) (
	response *ResponseScheme, err error) {

	if fieldConfigurationID == 0 {
		return nil, models.ErrNoFieldConfigurationIDError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/2/fieldconfiguration/%v/fields", fieldConfigurationID)

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
