package jira

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type FieldConfigurationService struct{ client *Client }

type FieldConfigurationPageScheme struct {
	MaxResults int                         `json:"maxResults,omitempty"`
	StartAt    int                         `json:"startAt,omitempty"`
	Total      int                         `json:"total,omitempty"`
	IsLast     bool                        `json:"isLast,omitempty"`
	Values     []*FieldConfigurationScheme `json:"values,omitempty"`
}

type FieldConfigurationScheme struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IsDefault   bool   `json:"isDefault,omitempty"`
}

// Gets Returns a paginated list of all field configurations.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-all-field-configurations
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-field-configurations/#api-rest-api-3-fieldconfiguration-get
func (f *FieldConfigurationService) Gets(ctx context.Context, ids []int, isDefault bool, startAt, maxResults int) (
	result *FieldConfigurationPageScheme, response *ResponseScheme, err error) {

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

type FieldConfigurationItemPageScheme struct {
	MaxResults int                             `json:"maxResults,omitempty"`
	StartAt    int                             `json:"startAt,omitempty"`
	Total      int                             `json:"total,omitempty"`
	IsLast     bool                            `json:"isLast,omitempty"`
	Values     []*FieldConfigurationItemScheme `json:"values,omitempty"`
}

type FieldConfigurationItemScheme struct {
	ID          string `json:"id,omitempty"`
	IsHidden    bool   `json:"isHidden,omitempty"`
	IsRequired  bool   `json:"isRequired,omitempty"`
	Description string `json:"description,omitempty"`
}

// Items Returns a paginated list of all fields for a configuration.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-field-configuration-items
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-field-configurations/#api-rest-api-3-fieldconfiguration-id-fields-get
func (f *FieldConfigurationService) Items(ctx context.Context, fieldConfigurationID, startAt, maxResults int) (
	result *FieldConfigurationItemPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("rest/api/3/fieldconfiguration/%v/fields?%v", fieldConfigurationID, params.Encode())

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

type FieldConfigurationSchemePageScheme struct {
	MaxResults int                               `json:"maxResults,omitempty"`
	StartAt    int                               `json:"startAt,omitempty"`
	Total      int                               `json:"total,omitempty"`
	IsLast     bool                              `json:"isLast,omitempty"`
	Values     []*FieldConfigurationSchemeScheme `json:"values,omitempty"`
}

type FieldConfigurationSchemeScheme struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// Schemes Returns a paginated list of field configuration schemes.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-all-field-configuration-schemes
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-field-configurations/#api-rest-api-3-fieldconfigurationscheme-get
func (f *FieldConfigurationService) Schemes(ctx context.Context, IDs []int, startAt, maxResults int) (
	result *FieldConfigurationSchemePageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range IDs {
		params.Add("id", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/3/fieldconfigurationscheme?%v", params.Encode())

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

type FieldConfigurationIssueTypeItemPageScheme struct {
	MaxResults int                                      `json:"maxResults,omitempty"`
	StartAt    int                                      `json:"startAt,omitempty"`
	Total      int                                      `json:"total,omitempty"`
	IsLast     bool                                     `json:"isLast,omitempty"`
	Values     []*FieldConfigurationIssueTypeItemScheme `json:"values,omitempty"`
}

type FieldConfigurationIssueTypeItemScheme struct {
	FieldConfigurationSchemeID string `json:"fieldConfigurationSchemeId,omitempty"`
	IssueTypeID                string `json:"issueTypeId,omitempty"`
	FieldConfigurationID       string `json:"fieldConfigurationId,omitempty"`
}

// IssueTypeItems Returns a paginated list of field configuration issue type items.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-field-configuration-issue-type-items
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-field-configurations/#api-rest-api-3-fieldconfigurationscheme-mapping-get
func (f *FieldConfigurationService) IssueTypeItems(ctx context.Context, fieldConfigIDs []int, startAt, maxResults int) (
	result *FieldConfigurationIssueTypeItemPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range fieldConfigIDs {
		params.Add("fieldConfigurationSchemeId", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/3/fieldconfigurationscheme/mapping?%v", params.Encode())

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

type FieldConfigurationSchemeProjectPageScheme struct {
	MaxResults int                                      `json:"maxResults,omitempty"`
	StartAt    int                                      `json:"startAt,omitempty"`
	Total      int                                      `json:"total,omitempty"`
	IsLast     bool                                     `json:"isLast,omitempty"`
	Values     []*FieldConfigurationSchemeProjectScheme `json:"values,omitempty"`
}

type FieldConfigurationSchemeProjectScheme struct {
	ProjectIds               []string                        `json:"projectIds,omitempty"`
	FieldConfigurationScheme *FieldConfigurationSchemeScheme `json:"fieldConfigurationScheme,omitempty"`
}

// SchemesByProject Returns a paginated list of field configuration schemes and, for each scheme, a list of the projects that use it.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/configuration#get-field-configuration-schemes-for-projects
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-field-configurations/#api-rest-api-3-fieldconfigurationscheme-project-get
func (f *FieldConfigurationService) SchemesByProject(ctx context.Context, projectIDs []int, startAt, maxResults int) (
	result *FieldConfigurationSchemeProjectPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, projectID := range projectIDs {
		params.Add("projectId", strconv.Itoa(projectID))
	}

	var endpoint = fmt.Sprintf("rest/api/3/fieldconfigurationscheme/project?%v", params.Encode())

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

// Assign Assigns a field configuration scheme to a project.
// If the field configuration scheme ID is null, the operation assigns the default field configuration scheme.
// Docs: TODO Issue 50
// Official Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-field-configurations/#api-rest-api-3-fieldconfigurationscheme-project-put
func (f *FieldConfigurationService) Assign(ctx context.Context, fieldConfigurationSchemeID, projectID string) (response *ResponseScheme, err error) {

	if len(projectID) == 0 {
		return nil, notProjectIDError
	}

	payload := struct {
		SchemeID  string `json:"fieldConfigurationSchemeId,omitempty"`
		ProjectID string `json:"projectId,omitempty"`
	}{
		SchemeID:  fieldConfigurationSchemeID,
		ProjectID: projectID,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	var endpoint = "rest/api/3/fieldconfigurationscheme/project"

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

var (
	notProjectIDError = fmt.Errorf("error!, please provide a vaild projectID value")
)
