package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type FieldConfigurationService struct{ client *Client }

// Gets Returns a paginated list of all field configurations.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-all-field-configurations
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-field-configurations/#api-rest-api-2-fieldconfiguration-get
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

	var endpoint = fmt.Sprintf("rest/api/2/fieldconfiguration?%v", params.Encode())

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

// Items Returns a paginated list of all fields for a configuration.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-field-configuration-items
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-field-configurations/#api-rest-api-2-fieldconfiguration-id-fields-get
func (f *FieldConfigurationService) Items(ctx context.Context, fieldConfigurationID, startAt, maxResults int) (
	result *models.FieldConfigurationItemPageScheme, response *ResponseScheme, err error) {

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

// Schemes Returns a paginated list of field configuration schemes.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-all-field-configuration-schemes
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-field-configurations/#api-rest-api-2-fieldconfigurationscheme-get
func (f *FieldConfigurationService) Schemes(ctx context.Context, IDs []int, startAt, maxResults int) (
	result *models.FieldConfigurationSchemePageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range IDs {
		params.Add("id", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/2/fieldconfigurationscheme?%v", params.Encode())

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

// IssueTypeItems Returns a paginated list of field configuration issue type items.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-field-configuration-issue-type-items
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-field-configurations/#api-rest-api-2-fieldconfigurationscheme-mapping-get
func (f *FieldConfigurationService) IssueTypeItems(ctx context.Context, fieldConfigIDs []int, startAt, maxResults int) (
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

// SchemesByProject Returns a paginated list of field configuration schemes and, for each scheme, a list of the projects that use it.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-field-configuration-schemes-for-projects
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-field-configurations/#api-rest-api-2-fieldconfigurationscheme-project-get
func (f *FieldConfigurationService) SchemesByProject(ctx context.Context, projectIDs []int, startAt, maxResults int) (
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

// Assign assigns a field configuration scheme to a project.
// If the field configuration scheme ID is null, the operation assigns the default field configuration scheme.
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-field-configurations/#api-rest-api-2-fieldconfigurationscheme-project-put
func (f *FieldConfigurationService) Assign(ctx context.Context, fieldConfigurationSchemeID, projectID string) (response *ResponseScheme, err error) {

	if len(projectID) == 0 {
		return nil, models.ErrNoProjectIDError
	}

	payload := struct {
		SchemeID  string `json:"fieldConfigurationSchemeId,omitempty"`
		ProjectID string `json:"projectId,omitempty"`
	}{
		SchemeID:  fieldConfigurationSchemeID,
		ProjectID: projectID,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	var endpoint = "rest/api/2/fieldconfigurationscheme/project"

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
