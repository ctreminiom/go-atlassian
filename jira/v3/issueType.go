package v3

import (
	"context"
	"fmt"
	"net/http"
)

type IssueTypeService struct {
	client       *Client
	Scheme       *IssueTypeSchemeService
	ScreenScheme *IssueTypeScreenSchemeService
}

type IssueTypeScheme struct {
	Self           string                `json:"self,omitempty"`
	ID             string                `json:"id,omitempty"`
	Description    string                `json:"description,omitempty"`
	IconURL        string                `json:"iconUrl,omitempty"`
	Name           string                `json:"name,omitempty"`
	Subtask        bool                  `json:"subtask,omitempty"`
	AvatarID       int                   `json:"avatarId,omitempty"`
	EntityID       string                `json:"entityId,omitempty"`
	HierarchyLevel int                   `json:"hierarchyLevel,omitempty"`
	Scope          *IssueTypeScopeScheme `json:"scope,omitempty"`
}

type IssueTypeScopeScheme struct {
	Type    string         `json:"type,omitempty"`
	Project *ProjectScheme `json:"project,omitempty"`
}

// Gets returns all issue types.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/type#get-all-issue-types-for-user
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-types/#api-rest-api-3-issuetype-get
func (i *IssueTypeService) Gets(ctx context.Context) (result []*IssueTypeScheme, response *ResponseScheme, err error) {

	var endpoint = "rest/api/3/issuetype"

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

type IssueTypePayloadScheme struct {
	Name           string `json:"name,omitempty"`
	Description    string `json:"description,omitempty"`
	Type           string `json:"type,omitempty"`
	HierarchyLevel int    `json:"hierarchyLevel,omitempty"`
	AvatarID       int    `json:"avatarId,omitempty"`
}

// Create creates an issue type and adds it to the default issue type scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/type#create-issue-type
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-types/#api-rest-api-3-issuetype-post
func (i *IssueTypeService) Create(ctx context.Context, payload *IssueTypePayloadScheme) (result *IssueTypeScheme,
	response *ResponseScheme, err error) {

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	var endpoint = "rest/api/3/issuetype"

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

// Get returns an issue type.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/type#get-issue-type
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-types/#api-rest-api-3-issuetype-id-get
func (i *IssueTypeService) Get(ctx context.Context, issueTypeID string) (result *IssueTypeScheme, response *ResponseScheme,
	err error) {

	if len(issueTypeID) == 0 {
		return nil, nil, notIssueTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetype/%v", issueTypeID)

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

// Update updates the issue type.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/type#update-issue-type
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-types/#api-rest-api-3-issuetype-id-put
func (i *IssueTypeService) Update(ctx context.Context, issueTypeID string, payload *IssueTypePayloadScheme) (
	result *IssueTypeScheme, response *ResponseScheme, err error) {

	if len(issueTypeID) == 0 {
		return nil, nil, notIssueTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetype/%v", issueTypeID)

	payloadAsReader, err := transformStructToReader(payload)
	if err != nil {
		return nil, nil, err
	}

	request, err := i.client.newRequest(ctx, http.MethodPut, endpoint, payloadAsReader)
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

// Delete deletes the issue type.
// If the issue type is in use, all uses are updated with the alternative issue type (alternativeIssueTypeId).
// A list of alternative issue types are obtained from the Get alternative issue types resource.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/type#delete-issue-type
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-types/#api-rest-api-3-issuetype-id-delete
func (i *IssueTypeService) Delete(ctx context.Context, issueTypeID string) (response *ResponseScheme, err error) {

	if len(issueTypeID) == 0 {
		return nil, notIssueTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetype/%v", issueTypeID)

	request, err := i.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = i.client.call(request, nil)
	if err != nil {
		return
	}

	return
}

// Alternatives returns a list of issue types that can be used to replace the issue type.
// The alternative issue types are those assigned to the same workflow scheme, field configuration scheme, and screen scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/type#get-alternative-issue-types
// Atlassian Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-issue-types/#api-rest-api-3-issuetype-id-alternatives-get
func (i *IssueTypeService) Alternatives(ctx context.Context, issueTypeID string) (result []*IssueTypeScheme,
	response *ResponseScheme, err error) {

	if len(issueTypeID) == 0 {
		return nil, nil, notIssueTypeIDError
	}

	var endpoint = fmt.Sprintf("rest/api/3/issuetype/%v/alternatives", issueTypeID)

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

var (
	notIssueTypeIDError = fmt.Errorf("error, please provide a valid issue type ID")
)
