package sm

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type OrganizationService struct{ client *Client }

// Gets returns a list of organizations in the Jira Service Management instance.
// Use this method when you want to present a list of organizations or want to locate an organization by name.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/organization#get-organizations
func (o *OrganizationService) Gets(ctx context.Context, accountID string, start, limit int) (
	result *OrganizationPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if len(accountID) != 0 {
		params.Add("accountId", accountID)
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/organization?%v", params.Encode())

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Get returns details of an organization.
// Use this method to get organization details whenever your application component is passed an organization ID
// but needs to display other organization details.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/organization#get-organization
func (o *OrganizationService) Get(ctx context.Context, organizationID int) (result *OrganizationScheme,
	response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/servicedeskapi/organization/%v", organizationID)

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Delete deletes an organization.
// Note that the organization is deleted regardless of other associations it may have.
// For example, associations with service desks.
func (o *OrganizationService) Delete(ctx context.Context, organizationID int) (response *ResponseScheme, err error) {

	var endpoint = fmt.Sprintf("rest/servicedeskapi/organization/%v", organizationID)

	request, err := o.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = o.client.Call(request, nil)
	if err != nil {
		return
	}

	return
}

// Create creates an organization by passing the name of the organization.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/organization#create-organization
func (o *OrganizationService) Create(ctx context.Context, name string) (result *OrganizationScheme,
	response *ResponseScheme, err error) {

	if len(name) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid name value")
	}

	payload := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = "rest/servicedeskapi/organization"

	request, err := o.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = o.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Users returns all the users associated with an organization.
// Use this method where you want to provide a list of users for an
// organization or determine if a user is associated with an organization.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/organization#get-users-in-organization
func (o *OrganizationService) Users(ctx context.Context, organizationID, start, limit int) (
	result *OrganizationUsersPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/organization/%v/user?%v", organizationID, params.Encode())

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Add adds users to an organization.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/organization#add-users-to-organization
func (o *OrganizationService) Add(ctx context.Context, organizationID int, accountIDs []string) (response *ResponseScheme, err error) {

	if len(accountIDs) == 0 {
		return nil, fmt.Errorf("error, please provide a valid accountIDs list of values")
	}

	payload := struct {
		AccountIds []string `json:"accountIds"`
	}{
		AccountIds: accountIDs,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = fmt.Sprintf("rest/servicedeskapi/organization/%v/user", organizationID)

	request, err := o.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = o.client.Call(request, nil)
	if err != nil {
		return
	}

	return
}

// Remove removes users from an organization.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/organization#remove-users-from-organization
func (o *OrganizationService) Remove(ctx context.Context, organizationID int, accountIDs []string) (response *ResponseScheme, err error) {

	if len(accountIDs) == 0 {
		return nil, fmt.Errorf("error, please provide a valid accountIDs list of values")
	}

	payload := struct {
		AccountIds []string `json:"accountIds"`
	}{
		AccountIds: accountIDs,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = fmt.Sprintf("rest/servicedeskapi/organization/%v/user", organizationID)

	request, err := o.client.newRequest(ctx, http.MethodDelete, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = o.client.Call(request, nil)
	if err != nil {
		return
	}

	return
}

// Project returns a list of all organizations associated with a service desk.
func (o *OrganizationService) Project(ctx context.Context, accountID string, serviceDeskPortalID, start,
	limit int) (result *OrganizationPageScheme, response *ResponseScheme, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if len(accountID) != 0 {
		params.Add("accountId", accountID)
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/organization?%v", serviceDeskPortalID, params.Encode())

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Associate adds an organization to a service desk.
// If the organization ID is already associated with the service desk,
// no change is made and the resource returns a 204 success code.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/organization#associate-organization
func (o *OrganizationService) Associate(ctx context.Context, serviceDeskPortalID, organizationID int) (
	response *ResponseScheme, err error) {

	payload := struct {
		OrganizationID int `json:"organizationId"`
	}{
		OrganizationID: organizationID,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/organization", serviceDeskPortalID)

	request, err := o.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = o.client.Call(request, nil)
	if err != nil {
		return
	}

	return
}

// Detach removes an organization from a service desk.
// If the organization ID does not match an organization associated with the service desk,
// no change is made and the resource returns a 204 success code.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/organization#detach-organization
func (o *OrganizationService) Detach(ctx context.Context, serviceDeskPortalID, organizationID int) (
	response *ResponseScheme, err error) {

	payload := struct {
		OrganizationID int `json:"organizationId"`
	}{
		OrganizationID: organizationID,
	}

	payloadAsReader, _ := transformStructToReader(&payload)
	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/organization", serviceDeskPortalID)

	request, err := o.client.newRequest(ctx, http.MethodDelete, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = o.client.Call(request, nil)
	if err != nil {
		return
	}

	return
}

type OrganizationUsersPageScheme struct {
	Size       int                              `json:"size"`
	Start      int                              `json:"start"`
	Limit      int                              `json:"limit"`
	IsLastPage bool                             `json:"isLastPage"`
	Values     []*OrganizationUserScheme        `json:"values"`
	Expands    []string                         `json:"_expands"`
	Links      *OrganizationUsersPageLinkScheme `json:"_links"`
}

type OrganizationUsersPageLinkScheme struct {
	Self    string `json:"self"`
	Base    string `json:"base"`
	Context string `json:"context"`
	Next    string `json:"next"`
	Prev    string `json:"prev"`
}

type OrganizationUserScheme struct {
	AccountID    string                      `json:"accountId"`
	Name         string                      `json:"name"`
	Key          string                      `json:"key"`
	EmailAddress string                      `json:"emailAddress"`
	DisplayName  string                      `json:"displayName"`
	Active       bool                        `json:"active"`
	TimeZone     string                      `json:"timeZone"`
	Links        *OrganizationUserLinkScheme `json:"_links"`
}

type OrganizationUserLinkScheme struct {
	Self     string `json:"self"`
	JiraRest string `json:"jiraRest"`
}

type OrganizationPageScheme struct {
	Size       int                         `json:"size"`
	Start      int                         `json:"start"`
	Limit      int                         `json:"limit"`
	IsLastPage bool                        `json:"isLastPage"`
	Values     []*OrganizationScheme       `json:"values"`
	Expands    []string                    `json:"_expands"`
	Links      *OrganizationPageLinkScheme `json:"_links"`
}

type OrganizationPageLinkScheme struct {
	Self    string `json:"self"`
	Base    string `json:"base"`
	Context string `json:"context"`
	Next    string `json:"next"`
	Prev    string `json:"prev"`
}

type OrganizationScheme struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Links struct {
		Self string `json:"self"`
	} `json:"_links"`
}
