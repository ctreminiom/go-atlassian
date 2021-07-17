package jira

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type IssueTypeScreenSchemeService struct{ client *Client }

type IssueTypeScreenSchemePageScheme struct {
	Self       string                         `json:"self,omitempty"`
	NextPage   string                         `json:"nextPage,omitempty"`
	MaxResults int                            `json:"maxResults,omitempty"`
	StartAt    int                            `json:"startAt,omitempty"`
	Total      int                            `json:"total,omitempty"`
	IsLast     bool                           `json:"isLast,omitempty"`
	Values     []*IssueTypeScreenSchemeScheme `json:"values,omitempty"`
}

type IssueTypeScreenSchemeScheme struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// Gets returns a paginated list of issue type screen schemes.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#get-issue-type-screen-schemes
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-get
func (i *IssueTypeScreenSchemeService) Gets(ctx context.Context, ids []int, startAt, maxResults int) (
	result *IssueTypeScreenSchemePageScheme, response *ResponseScheme, err error) {

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

type IssueTypeScreenSchemePayloadScheme struct {
	Name              string                                       `json:"name,omitempty"`
	IssueTypeMappings []*IssueTypeScreenSchemeMappingPayloadScheme `json:"issueTypeMappings,omitempty"`
}

type IssueTypeScreenSchemeMappingPayloadScheme struct {
	IssueTypeID    string `json:"issueTypeId,omitempty"`
	ScreenSchemeID string `json:"screenSchemeId,omitempty"`
}

type IssueTypeScreenScreenCreatedScheme struct {
	ID string `json:"id"`
}

// Create creates an issue type screen scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#create-issue-type-screen-scheme
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-post
func (i *IssueTypeScreenSchemeService) Create(ctx context.Context, payload *IssueTypeScreenSchemePayloadScheme) (
	result *IssueTypeScreenScreenCreatedScheme, response *ResponseScheme, err error) {

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
		return nil, notIssueTypeScreenSchemeIDError
	}

	if len(projectID) == 0 {
		return nil, notProjectIDError
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
	result *IssueTypeProjectScreenSchemePageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if len(projectIDs) == 0 {
		return nil, nil, notProjectsError
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

type IssueTypeProjectScreenSchemePageScheme struct {
	Self       string                                 `json:"self,omitempty"`
	NextPage   string                                 `json:"nextPage,omitempty"`
	MaxResults int                                    `json:"maxResults,omitempty"`
	StartAt    int                                    `json:"startAt,omitempty"`
	Total      int                                    `json:"total,omitempty"`
	IsLast     bool                                   `json:"isLast,omitempty"`
	Values     []*IssueTypeScreenSchemesProjectScheme `json:"values,omitempty"`
}

type IssueTypeScreenSchemesProjectScheme struct {
	IssueTypeScreenScheme *IssueTypeScreenSchemeScheme `json:"issueTypeScreenScheme,omitempty"`
	ProjectIds            []string                     `json:"projectIds,omitempty"`
}

// Mapping returns a paginated list of issue type screen scheme items.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/types/screen-scheme#get-issue-type-screen-scheme-items
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-mapping-get
func (i *IssueTypeScreenSchemeService) Mapping(ctx context.Context, issueTypeScreenSchemeIDs []int, startAt, maxResults int) (
	result *IssueTypeScreenSchemeMappingScheme, response *ResponseScheme, err error) {

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

type IssueTypeScreenSchemeMappingScheme struct {
	Self       string                             `json:"self,omitempty"`
	NextPage   string                             `json:"nextPage,omitempty"`
	MaxResults int                                `json:"maxResults,omitempty"`
	StartAt    int                                `json:"startAt,omitempty"`
	Total      int                                `json:"total,omitempty"`
	IsLast     bool                               `json:"isLast,omitempty"`
	Values     []*IssueTypeScreenSchemeItemScheme `json:"values,omitempty"`
}

type IssueTypeScreenSchemeItemScheme struct {
	IssueTypeScreenSchemeID string `json:"issueTypeScreenSchemeId,omitempty"`
	IssueTypeID             string `json:"issueTypeId,omitempty"`
	ScreenSchemeID          string `json:"screenSchemeId,omitempty"`
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
func (i *IssueTypeScreenSchemeService) Delete(ctx context.Context, issueTypeScreenSchemeID string) (response *ResponseScheme, err error) {

	if len(issueTypeScreenSchemeID) == 0 {
		return nil, notIssueTypeScreenSchemeIDError
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
	payload *IssueTypeScreenSchemePayloadScheme) (response *ResponseScheme, err error) {

	if len(issueTypeScreenSchemeID) == 0 {
		return nil, notIssueTypeScreenSchemeIDError
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
		return nil, notIssueTypeScreenSchemeIDError
	}

	if len(screenSchemeID) == 0 {
		return nil, notScreenSchemeIDError
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
// Atlassina Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-type-screen-schemes/#api-rest-api-3-issuetypescreenscheme-issuetypescreenschemeid-mapping-remove-post
func (i *IssueTypeScreenSchemeService) Remove(ctx context.Context, issueTypeScreenSchemeID string, issueTypeIDs []string) (response *ResponseScheme, err error) {

	if len(issueTypeScreenSchemeID) == 0 {
		return nil, notIssueTypeScreenSchemeIDError
	}

	if len(issueTypeIDs) == 0 {
		return nil, notIssueTypesError
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

var (
	notIssueTypeScreenSchemeIDError = fmt.Errorf("error, please provide a issueTypeScreenSchemeID value")
	notScreenSchemeIDError          = fmt.Errorf("error, please provide a screenSchemeID value")
)
