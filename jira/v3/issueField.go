package v3

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type FieldService struct {
	client        *Client
	Configuration *FieldConfigurationService
	Context       *FieldContextService
}

// Gets returns system and custom issue fields according to the following rules:
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields#get-fields
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-fields/#api-rest-api-3-field-get
func (f *FieldService) Gets(ctx context.Context) (result []*models.IssueFieldScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/field"

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

// Create creates a custom field.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields#create-custom-field
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-fields/#api-rest-api-3-field-post
func (f *FieldService) Create(ctx context.Context, payload *models.CustomFieldScheme) (result *models.IssueFieldScheme,
	response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/field"

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

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

// Search returns a paginated list of fields for Classic Jira projects.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields#get-fields-paginated
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-fields/#api-rest-api-3-field-search-get
// NOTE: Experimental Endpoint
func (f *FieldService) Search(ctx context.Context, options *models.FieldSearchOptionsScheme, startAt, maxResults int) (
	result *models.FieldSearchPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if options != nil {

		if len(options.Expand) != 0 {
			params.Add("expand", strings.Join(options.Expand, ","))
		}

		if len(options.Types) != 0 {
			params.Add("type", strings.Join(options.Types, ","))
		}

		if len(options.IDs) != 0 {
			params.Add("id", strings.Join(options.IDs, ","))
		}

		if len(options.OrderBy) != 0 {
			params.Add("orderBy", options.OrderBy)
		}

		if len(options.Query) != 0 {
			params.Add("query", options.Query)
		}
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/search?%v", params.Encode())

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

// Delete deletes a custom field. The custom field is deleted whether it is in the trash or not.
// See Edit or delete a custom field for more information on trashing and deleting custom fields.
// Docs: TODO: The documentation needs to be created
func (f *FieldService) Delete(ctx context.Context, fieldId string) (result *models.TaskScheme, response *ResponseScheme,
	err error) {

	if fieldId == "" {
		return nil, nil, models.ErrNoFieldIDError
	}

	endpoint := fmt.Sprintf("rest/api/3/field/%v", fieldId)

	request, err := f.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = f.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Trash moves a custom field to trash
func (f *FieldService) Trash(ctx context.Context, fieldId string) (response *ResponseScheme, err error) {

	if fieldId == "" {
		return nil, models.ErrNoFieldIDError
	}

	endpoint := fmt.Sprintf("rest/api/3/field/%v/trash", fieldId)

	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, nil)
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

// Restore restores a custom field from trash.
func (f *FieldService) Restore(ctx context.Context, fieldId string) (response *ResponseScheme, err error) {

	if fieldId == "" {
		return nil, models.ErrNoFieldIDError
	}

	endpoint := fmt.Sprintf("rest/api/3/field/%v/restore", fieldId)

	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, nil)
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
