package v3

import (
	"context"
	"fmt"
	models2 "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type IssueTypeScreenSchemeService struct{ client *Client }

// Gets returns a paginated list of issue type screen schemes.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#get-issue-type-screen-schemes
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-get
func (i *IssueTypeScreenSchemeService) Gets(ctx context.Context, ids []int, startAt, maxResults int) (
	result *models2.IssueTypeScreenSchemePageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range ids {
		params.Add("id", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescreenscheme?%v", params.Encode())

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Create creates an issue type screen scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#create-issue-type-screen-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-post
func (i *IssueTypeScreenSchemeService) Create(ctx context.Context, payload *models2.IssueTypeScreenSchemePayloadScheme) (
	result *models2.IssueTypeScreenScreenCreatedScheme, response *ResponseScheme, err error) {

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint = "rest/api/3/issuetypescreenscheme"

	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Assign assigns an issue type screen scheme to a project.
// Issue type screen schemes can only be assigned to classic projects.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#assign-issue-type-screen-scheme-to-project
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-project-put
func (i *IssueTypeScreenSchemeService) Assign(ctx context.Context, issueTypeScreenSchemeID, projectID string) (
	response *ResponseScheme, err error) {

	if len(issueTypeScreenSchemeID) == 0 {
		return nil, models2.ErrNoIssueTypeScreenSchemeIDError
	}

	if len(projectID) == 0 {
		return nil, models2.ErrNoProjectIDError
	}

	payload := struct {
		IssueTypeScreenSchemeID string `json:"issueTypeScreenSchemeId"`
		ProjectID               string `json:"projectId"`
	}{
		IssueTypeScreenSchemeID: issueTypeScreenSchemeID,
		ProjectID:               projectID,
	}

	var endpoint = "rest/api/3/issuetypescreenscheme/project"
	payloadAsReader, _ := transformStructToReader(&payload)

	request, err := i.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Projects returns a paginated list of issue type screen schemes and,
// for each issue type screen scheme, a list of the projects that use it.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#assign-issue-type-screen-scheme-to-project
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-project-get
func (i *IssueTypeScreenSchemeService) Projects(ctx context.Context, projectIDs []int, startAt, maxResults int) (
	result *models2.IssueTypeProjectScreenSchemePageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(projectIDs) == 0 {
		return nil, nil, models2.ErrNoProjectsError
	}

	for _, id := range projectIDs {
		params.Add("projectId", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescreenscheme/project?%v", params.Encode())

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Mapping returns a paginated list of issue type screen scheme items.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#get-issue-type-screen-scheme-items
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-mapping-get
func (i *IssueTypeScreenSchemeService) Mapping(ctx context.Context, issueTypeScreenSchemeIDs []int, startAt, maxResults int) (
	result *models2.IssueTypeScreenSchemeMappingScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range issueTypeScreenSchemeIDs {
		params.Add("issueTypeScreenSchemeId", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescreenscheme/mapping?%v", params.Encode())

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}

// Update updates an issue type screen scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#update-issue-type-screen-scheme
// Atlassian Docs; https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-issuetypescreenschemeid-put
func (i *IssueTypeScreenSchemeService) Update(ctx context.Context, issueTypeScreenSchemeID, name, description string) (
	response *ResponseScheme, err error) {

	if len(issueTypeScreenSchemeID) == 0 {
		return nil, notIssueTypeSchemeIDError
	}

	payload := struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescreenscheme/%v", issueTypeScreenSchemeID)

	request, err := i.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Delete deletes an issue type screen scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#delete-issue-type-screen-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-issuetypescreenschemeid-delete
func (i *IssueTypeScreenSchemeService) Delete(ctx context.Context, issueTypeScreenSchemeID string) (response *ResponseScheme,
	err error) {

	if len(issueTypeScreenSchemeID) == 0 {
		return nil, models2.ErrNoIssueTypeScreenSchemeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescreenscheme/%v", issueTypeScreenSchemeID)

	request, err := i.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Append appends issue type to screen scheme mappings to an issue type screen scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#append-mappings-to-issue-type-screen-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-issuetypescreenschemeid-mapping-put
func (i *IssueTypeScreenSchemeService) Append(ctx context.Context, issueTypeScreenSchemeID string,
	payload *models2.IssueTypeScreenSchemePayloadScheme) (response *ResponseScheme, err error) {

	if len(issueTypeScreenSchemeID) == 0 {
		return nil, models2.ErrNoIssueTypeScreenSchemeIDError
	}

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, err
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescreenscheme/%v/mapping", issueTypeScreenSchemeID)

	request, err := i.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// UpdateDefault updates the default screen scheme of an issue type screen scheme. The default screen scheme is used for all unmapped issue types.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#update-issue-type-screen-scheme-default-screen-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-issuetypescreenschemeid-mapping-default-put
func (i *IssueTypeScreenSchemeService) UpdateDefault(ctx context.Context, issueTypeScreenSchemeID, screenSchemeID string) (
	response *ResponseScheme, err error) {

	if len(issueTypeScreenSchemeID) == 0 {
		return nil, models2.ErrNoIssueTypeScreenSchemeIDError
	}

	if len(screenSchemeID) == 0 {
		return nil, models2.ErrNoScreenSchemeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescreenscheme/%v/mapping/default", issueTypeScreenSchemeID)

	payload := struct {
		ScreenSchemeID string `json:"screenSchemeId"`
	}{
		ScreenSchemeID: screenSchemeID,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	request, err := i.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Remove removes issue type to screen scheme mappings from an issue type screen scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#remove-mappings-from-issue-type-screen-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-issuetypescreenschemeid-mapping-remove-post
func (i *IssueTypeScreenSchemeService) Remove(ctx context.Context, issueTypeScreenSchemeID string, issueTypeIDs []string) (
	response *ResponseScheme, err error) {

	if len(issueTypeScreenSchemeID) == 0 {
		return nil, models2.ErrNoIssueTypeScreenSchemeIDError
	}

	if len(issueTypeIDs) == 0 {
		return nil, models2.ErrNoIssueTypesError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescreenscheme/%v/mapping/remove", issueTypeScreenSchemeID)

	payload := struct {
		IssueTypeIds []string `json:"issueTypeIds"`
	}{
		IssueTypeIds: issueTypeIDs,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	request, err := i.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// SchemesByProject returns a paginated list of projects associated with an issue type screen scheme.
// Docs: N/A
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-issuetypescreenschemeid-project-get
func (i *IssueTypeScreenSchemeService) SchemesByProject(ctx context.Context, issueTypeScreenSchemeID int, startAt,
	maxResults int) (result *models2.IssueTypeScreenSchemeByProjectPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("rest/api/3/issuetypescreenscheme/%v/project?%v", issueTypeScreenSchemeID, params.Encode())

	request, err := i.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = i.client.call(request, &result)
	if err != nil {
		return
	}

	return
}
