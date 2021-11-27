package v2

import (
	"context"
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type IssueTypeSchemeService struct{ client *Client }

// Gets returns a paginated list of issue type schemes.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#get-all-issue-type-schemes
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-type-schemes/#api-rest-api-2-issuetypescheme-get
func (i *IssueTypeSchemeService) Gets(ctx context.Context, issueTypeSchemeIDs []int, startAt, maxResults int) (
	result *models.IssueTypeSchemePageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range issueTypeSchemeIDs {
		params.Add("id", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/2/issuetypescheme?%v", params.Encode())

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

// Create creates an issue type scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#create-issue-type-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-type-schemes/#api-rest-api-2-issuetypescheme-post
// NOTE: Experimental endpoint
func (i *IssueTypeSchemeService) Create(ctx context.Context, payload *models.IssueTypeSchemePayloadScheme) (result *models.NewIssueTypeSchemeScheme,
	response *ResponseScheme, err error) {

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	/*
		Validation considerations for the Atlassian Documentation.

		-------------------------
		value: defaultIssueTypeId
		validation: This ID must be included in issueTypeIds.
		-------------------------

		-------------------------
		value: issueTypeIds
		validation: At least one standard issue type ID is required
	*/

	var containsTheIssueType bool
	for _, issueType := range payload.IssueTypeIds {

		// The DefaultIssueTypeID value is not required
		if payload.DefaultIssueTypeID == "" {
			containsTheIssueType = true
			break
		}

		if issueType == payload.DefaultIssueTypeID {
			containsTheIssueType = true
			break
		}
	}

	if !containsTheIssueType {
		return nil, nil, notDefaultIssueTypeError
	}

	var endpoint = "rest/api/2/issuetypescheme"

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

// Items returns a paginated list of issue type scheme items.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#get-issue-type-scheme-items
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-type-schemes/#api-rest-api-2-issuetypescheme-mapping-get
func (i *IssueTypeSchemeService) Items(ctx context.Context, issueTypeSchemeIDs []int, startAt, maxResults int) (
	result *models.IssueTypeSchemeItemPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, id := range issueTypeSchemeIDs {
		params.Add("issueTypeSchemeId", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/2/issuetypescheme/mapping?%v", params.Encode())

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

// Projects returns a paginated list of issue type schemes and, for each issue type scheme, a list of the projects that use it.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#get-issue-type-schemes-for-projects
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-type-schemes/#api-rest-api-2-issuetypescheme-project-get
func (i *IssueTypeSchemeService) Projects(ctx context.Context, projectIDs []int, startAt, maxResults int) (
	result *models.ProjectIssueTypeSchemePageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(projectIDs) == 0 {
		return nil, nil, notProjectParamValueError
	}

	for _, id := range projectIDs {
		params.Add("projectId", strconv.Itoa(id))
	}

	var endpoint = fmt.Sprintf("rest/api/2/issuetypescheme/project?%v", params.Encode())

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

// Assign assigns an issue type scheme to a project.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#assign-issue-type-scheme-to-project
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-type-schemes/#api-rest-api-2-issuetypescheme-project-put
func (i *IssueTypeSchemeService) Assign(ctx context.Context, issueTypeSchemeID, projectID string) (response *ResponseScheme, err error) {

	if len(issueTypeSchemeID) == 0 {
		return nil, notIssueTypeSchemeIDError
	}

	if len(projectID) == 0 {
		return nil, models.ErrNoProjectIDError
	}

	payload := struct {
		IssueTypeSchemeID string `json:"issueTypeSchemeId"`
		ProjectID         string `json:"projectId"`
	}{
		IssueTypeSchemeID: issueTypeSchemeID,
		ProjectID:         projectID,
	}

	var endpoint = "rest/api/2/issuetypescheme/project"

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

// Update updates an issue type scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#update-issue-type-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-type-schemes/#api-rest-api-2-issuetypescheme-issuetypeschemeid-put
// NOTE: Experimental Method
func (i *IssueTypeSchemeService) Update(ctx context.Context, issueTypeSchemeID int, payload *models.IssueTypeSchemePayloadScheme) (
	response *ResponseScheme, err error) {

	if issueTypeSchemeID == 0 {
		return nil, models.ErrNoIssueTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/issuetypescheme/%v", issueTypeSchemeID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, err
	}

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

// Delete deletes an issue type scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#delete-issue-type-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-type-schemes/#api-rest-api-2-issuetypescheme-issuetypeschemeid-delete
// NOTE: Experimental Method
func (i *IssueTypeSchemeService) Delete(ctx context.Context, issueTypeSchemeID int) (response *ResponseScheme, err error) {

	if issueTypeSchemeID == 0 {
		return nil, models.ErrNoIssueTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/2/issuetypescheme/%v", issueTypeSchemeID)

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

// Append adds issue types to an issue type scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#add-issue-types-to-issue-type-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-type-schemes/#api-rest-api-2-issuetypescheme-issuetypeschemeid-issuetype-put
func (i *IssueTypeSchemeService) Append(ctx context.Context, issueTypeSchemeID int, issueTypeIDs []int) (response *ResponseScheme, err error) {

	if len(issueTypeIDs) == 0 {
		return nil, models.ErrNoIssueTypesError
	}

	var issueTypesIDsAsStringSlice []string
	for _, issueTypeID := range issueTypeIDs {
		issueTypesIDsAsStringSlice = append(issueTypesIDsAsStringSlice, strconv.Itoa(issueTypeID))
	}

	payload := struct {
		IssueTypeIds []string `json:"issueTypeIds"`
	}{
		IssueTypeIds: issueTypesIDsAsStringSlice,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = fmt.Sprintf("rest/api/2/issuetypescheme/%v/issuetype", issueTypeSchemeID)

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

// Remove removes an issue type from an issue type scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/scheme#remove-issue-type-from-issue-type-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v2/api-group-issue-type-schemes/#api-rest-api-2-issuetypescheme-issuetypeschemeid-issuetype-issuetypeid-delete
func (i *IssueTypeSchemeService) Remove(ctx context.Context, issueTypeSchemeID, issueTypeID int) (response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/api/2/issuetypescheme/%v/issuetype/%v", issueTypeSchemeID, issueTypeID)

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

var (
	notDefaultIssueTypeError  = fmt.Errorf("error, please add the DefaultIssueTypeID value on the IssueTypeIds value")
	notProjectParamValueError = fmt.Errorf("error, please provide values on the projectIDs param")
	notIssueTypeSchemeIDError = fmt.Errorf("error!, please provide a valid issueTypeSchemeID value")
)
