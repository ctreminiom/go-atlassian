package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type FieldConfigurationSchemeService struct{ client *Client }

// Gets returns a paginated list of field configuration schemes.
// Only field configuration schemes used in classic projects are returned.
func (f *FieldConfigurationSchemeService) Gets(ctx context.Context, ids []int, startAt, maxResults int) (
	result *models.FieldConfigurationSchemePageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range ids {
		params.Add("id", strconv.Itoa(id))
	}

	endpoint := fmt.Sprintf("rest/api/2/fieldconfigurationscheme?%v", params.Encode())

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

// Create creates a field configuration scheme.
// This operation can only create field configuration schemes used in company-managed (classic) projects.
// EXPERIMENTAL
func (f *FieldConfigurationSchemeService) Create(ctx context.Context, name, description string) (
	result *models.FieldConfigurationSchemeScheme, response *ResponseScheme, err error) {

	if name == "" {
		return nil, nil, models.ErrNoFieldConfigurationSchemeNameError
	}

	payload := struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}
	payloadAsReader, _ := transformStructToReader(&payload)
	endpoint := "rest/api/2/fieldconfigurationscheme"

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

// Mapping returns a paginated list of field configuration issue type items.
// Only items used in classic projects are returned.
func (f *FieldConfigurationSchemeService) Mapping(ctx context.Context, fieldConfigIDs []int, startAt, maxResults int) (
	result *models.FieldConfigurationIssueTypeItemPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range fieldConfigIDs {
		params.Add("fieldConfigurationSchemeId", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/2/fieldconfigurationscheme/mapping?%v", params.Encode())

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

// Project returns a paginated list of field configuration schemes and, for each scheme, a list of the projects that use it.
// The list is sorted by field configuration scheme ID. The first item contains the list of project IDs assigned to the default field configuration scheme.
// Only field configuration schemes used in classic projects are returned.
func (f *FieldConfigurationSchemeService) Project(ctx context.Context, projectIDs []int, startAt, maxResults int) (
	result *models.FieldConfigurationSchemeProjectPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, projectID := range projectIDs {
		params.Add("projectId", strconv.Itoa(projectID))
	}

	var endpoint = fmt.Sprintf("rest/api/2/fieldconfigurationscheme/project?%v", params.Encode())

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

// Assign assigns a field configuration scheme to a project. If the field configuration scheme ID is null,
// the operation assigns the default field configuration scheme.
// Field configuration schemes can only be assigned to classic projects.
func (f *FieldConfigurationSchemeService) Assign(ctx context.Context, payload *models.FieldConfigurationSchemeAssignPayload) (
	response *ResponseScheme, err error) {

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := "rest/api/2/fieldconfigurationscheme/project"

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

// Update updates a field configuration scheme.
// This operation can only update field configuration schemes used in company-managed (classic) projects.
// EXPERIMENTAL
func (f *FieldConfigurationSchemeService) Update(ctx context.Context, schemeID int, name, description string) (
	response *ResponseScheme, err error) {

	if schemeID == 0 {
		return nil, models.ErrNoFieldConfigurationSchemeIDError
	}

	if name == "" {
		return nil, models.ErrNoFieldConfigurationSchemeNameError
	}

	payload := struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}
	payloadAsReader, _ := transformStructToReader(&payload)
	endpoint := fmt.Sprintf("rest/api/2/fieldconfigurationscheme/%v", schemeID)

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

// Delete deletes a field configuration scheme.
// This operation can only delete field configuration schemes used in company-managed (classic) projects.
// EXPERIMENTAL
func (f *FieldConfigurationSchemeService) Delete(ctx context.Context, schemeID int) (response *ResponseScheme, err error) {

	if schemeID == 0 {
		return nil, models.ErrNoFieldConfigurationSchemeIDError
	}

	endpoint := fmt.Sprintf("rest/api/2/fieldconfigurationscheme/%v", schemeID)

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

// Link assigns issue types to field configurations on field configuration scheme.
// This operation can only modify field configuration schemes used in company-managed (classic) projects.
// EXPERIMENTAL
func (f *FieldConfigurationSchemeService) Link(ctx context.Context, schemeID int, payload *models.FieldConfigurationToIssueTypeMappingPayloadScheme) (
	response *ResponseScheme, err error) {

	if schemeID == 0 {
		return nil, models.ErrNoFieldConfigurationSchemeIDError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/api/2/fieldconfigurationscheme/%v/mapping", schemeID)

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

	return response, err
}

// Unlink removes issue types from the field configuration scheme.
// This operation can only modify field configuration schemes used in company-managed (classic) projects.
// EXPERIMENTAL
func (f *FieldConfigurationSchemeService) Unlink(ctx context.Context, schemeID int, issueTypeIDs []string) (response *ResponseScheme, err error) {

	if schemeID == 0 {
		return nil, models.ErrNoFieldConfigurationSchemeIDError
	}

	payload := struct {
		IssueTypeIds []string `json:"issueTypeIds"`
	}{
		IssueTypeIds: issueTypeIDs,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	endpoint := fmt.Sprintf("rest/api/2/fieldconfigurationscheme/%v/mapping/delete", schemeID)

	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
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
