package sm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type OrganizationService struct{ client *Client }

// This method returns a list of organizations in the Jira Service Management instance.
// Use this method when you want to present a list of organizations or want to locate an organization by name.
func (o *OrganizationService) Gets(ctx context.Context, accountID string, start, limit int) (result *OrganizationPageScheme, response *Response, err error) {

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

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// This method returns details of an organization.
// Use this method to get organization details whenever your application component is passed an organization ID
// but needs to display other organization details.
func (o *OrganizationService) Get(ctx context.Context, organizationID int) (result *OrganizationScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/servicedeskapi/organization/%v", organizationID)

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// This method deletes an organization.
// Note that the organization is deleted regardless of other associations it may have.
// For example, associations with service desks.
func (o *OrganizationService) Delete(ctx context.Context, organizationID int) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/servicedeskapi/organization/%v", organizationID)

	request, err := o.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	return
}

// This method creates an organization by passing the name of the organization.
func (o *OrganizationService) Create(ctx context.Context, name string) (result *OrganizationScheme, response *Response, err error) {

	if len(name) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid name value")
	}

	payload := struct {
		Name string `json:"name"`
	}{Name: name}

	var endpoint = "rest/servicedeskapi/organization"

	request, err := o.client.newRequest(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// This method returns all the users associated with an organization.
// Use this method where you want to provide a list of users for an
// organization or determine if a user is associated with an organization.
func (o *OrganizationService) Users(ctx context.Context, organizationID, start, limit int) (result *OrganizationUsersPageScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	var endpoint = fmt.Sprintf("rest/servicedeskapi/organization/%v/user?%v", organizationID, params.Encode())

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationUsersPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// This method adds users to an organization.
func (o *OrganizationService) Add(ctx context.Context, organizationID int, accountIDs []string) (response *Response, err error) {

	if len(accountIDs) == 0 {
		return nil, fmt.Errorf("error, please provide a valid accountIDs list of values")
	}

	payload := struct {
		AccountIds []string `json:"accountIds"`
	}{AccountIds: accountIDs}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/organization/%v/user", organizationID)

	request, err := o.client.newRequest(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	return
}

// This method removes users from an organization.
func (o *OrganizationService) Remove(ctx context.Context, organizationID int, accountIDs []string) (response *Response, err error) {

	if len(accountIDs) == 0 {
		return nil, fmt.Errorf("error, please provide a valid accountIDs list of values")
	}

	payload := struct {
		AccountIds []string `json:"accountIds"`
	}{AccountIds: accountIDs}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/organization/%v/user", organizationID)

	request, err := o.client.newRequest(ctx, http.MethodDelete, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	return
}

// This method returns a list of all organizations associated with a service desk.
func (o *OrganizationService) Project(ctx context.Context, accountID string, projectID, start, limit int) (result *OrganizationPageScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if len(accountID) != 0 {
		params.Add("accountId", accountID)
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/organization?%v", projectID, params.Encode())

	request, err := o.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	result = new(OrganizationPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// This method adds an organization to a service desk.
// If the organization ID is already associated with the service desk,
// no change is made and the resource returns a 204 success code.
func (o *OrganizationService) Associate(ctx context.Context, projectID, organizationID int) (response *Response, err error) {

	payload := struct {
		OrganizationID int `json:"organizationId"`
	}{OrganizationID: organizationID}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/organization", projectID)

	request, err := o.client.newRequest(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	return
}

// This method removes an organization from a service desk.
// If the organization ID does not match an organization associated with the service desk,
// no change is made and the resource returns a 204 success code.
func (o *OrganizationService) Detach(ctx context.Context, projectID, organizationID int) (response *Response, err error) {

	payload := struct {
		OrganizationID int `json:"organizationId"`
	}{OrganizationID: organizationID}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/organization", projectID)

	request, err := o.client.newRequest(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = o.client.Do(request)
	if err != nil {
		return
	}

	return
}

//detach
type OrganizationUsersPageScheme struct {
	Size       int  `json:"size"`
	Start      int  `json:"start"`
	Limit      int  `json:"limit"`
	IsLastPage bool `json:"isLastPage"`
	Values     []struct {
		AccountID    string `json:"accountId"`
		Name         string `json:"name"`
		Key          string `json:"key"`
		EmailAddress string `json:"emailAddress"`
		DisplayName  string `json:"displayName"`
		Active       bool   `json:"active"`
		TimeZone     string `json:"timeZone"`
		Links        struct {
			Self       string `json:"self"`
			JiraRest   string `json:"jiraRest"`
			AvatarUrls struct {
			} `json:"avatarUrls"`
		} `json:"_links"`
	} `json:"values"`
	Expands []string `json:"_expands"`
	Links   struct {
		Self    string `json:"self"`
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
	} `json:"_links"`
}

type OrganizationPageScheme struct {
	Size       int                   `json:"size"`
	Start      int                   `json:"start"`
	Limit      int                   `json:"limit"`
	IsLastPage bool                  `json:"isLastPage"`
	Values     []*OrganizationScheme `json:"values"`
	Expands    []string              `json:"_expands"`
	Links      struct {
		Self    string `json:"self"`
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
	} `json:"_links"`
}

type OrganizationScheme struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Links struct {
		Self string `json:"self"`
	} `json:"_links"`
}
