package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"net/url"
	"strconv"
)

type UserService struct {
	client *Client
	Search *UserSearchService
}

type UserScheme struct {
	Self         string `json:"self"`
	Key          string `json:"key"`
	AccountID    string `json:"accountId"`
	AccountType  string `json:"accountType"`
	Name         string `json:"name"`
	EmailAddress string `json:"emailAddress"`
	AvatarUrls   struct {
		One6X16   string `json:"16x16"`
		Two4X24   string `json:"24x24"`
		Three2X32 string `json:"32x32"`
		Four8X48  string `json:"48x48"`
	} `json:"avatarUrls"`
	DisplayName string `json:"displayName"`
	Active      bool   `json:"active"`
	TimeZone    string `json:"timeZone"`
	Locale      string `json:"locale"`
	Groups      struct {
		Size       int               `json:"size"`
		Items      []UserGroupScheme `json:"items"`
		MaxResults int               `json:"max-results"`
	} `json:"groups"`
	ApplicationRoles struct {
		Size  int `json:"size"`
		Items []struct {
			Key                  string   `json:"key"`
			Groups               []string `json:"groups"`
			Name                 string   `json:"name"`
			DefaultGroups        []string `json:"defaultGroups"`
			SelectedByDefault    bool     `json:"selectedByDefault"`
			Defined              bool     `json:"defined"`
			NumberOfSeats        int      `json:"numberOfSeats"`
			RemainingSeats       int      `json:"remainingSeats"`
			UserCount            int      `json:"userCount"`
			UserCountDescription string   `json:"userCountDescription"`
			HasUnlimitedSeats    bool     `json:"hasUnlimitedSeats"`
			Platform             bool     `json:"platform"`
		} `json:"items"`
		MaxResults int `json:"max-results"`
	} `json:"applicationRoles"`
	Expand string `json:"expand"`
}

// Returns a user.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-users/#api-rest-api-3-user-get
func (u *UserService) Get(ctx context.Context, accountID string, expands []string) (result *UserScheme, response *Response, err error) {

	if len(accountID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid accountID value")
	}

	params := url.Values{}

	var expand string
	for index, value := range expands {

		if index == 0 {
			expand = value
			continue
		}

		expand += "," + value
	}

	if len(expand) != 0 {
		params.Add("expand", expand)
	}

	params.Add("accountId", accountID)

	var endpoint = fmt.Sprintf("rest/api/3/user?%v", params.Encode())

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.Do(request)
	if err != nil {
		return
	}

	result = new(UserScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type UserPayloadScheme struct {
	Password     string `json:"password,omitempty"`
	EmailAddress string `json:"emailAddress" validate:"required"`
	DisplayName  string `json:"displayName" validate:"required"`
	Notification bool   `json:"notification,omitempty"`
}

// Creates a user. This resource is retained for legacy compatibility.
// As soon as a more suitable alternative is available this resource will be deprecated.
// The option is provided to set or generate a password for the user.
// When using the option to generate a password, by omitting password from the request, include "notification": "true" to ensure the user is
// sent an email advising them that their account is created.
// This email includes a link for the user to set their password. If the notification isn't sent for a generated password,
// the user will need to be sent a reset password request from Jira.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-users/#api-rest-api-3-user-post
func (u *UserService) Create(ctx context.Context, payload *UserPayloadScheme) (result *UserScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, please provide a valid UserPayloadScheme pointer")
	}

	validate := validator.New()
	if err = validate.Struct(payload); err != nil {
		err = fmt.Errorf("error: issuetype type payload invalid: %v", err.Error())
		return
	}

	var endpoint = "rest/api/3/user"

	request, err := u.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = u.client.Do(request)
	if err != nil {
		return
	}

	result = new(UserScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes a user.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-users/#api-rest-api-3-user-delete
func (u *UserService) Delete(ctx context.Context, accountID string) (response *Response, err error) {

	if len(accountID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid accountID value")
	}

	params := url.Values{}
	params.Add("accountId", accountID)
	var endpoint = fmt.Sprintf("rest/api/3/user?%v", params.Encode())

	request, err := u.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = u.client.Do(request)
	if err != nil {
		return
	}

	return
}

type UserSearchPageScheme struct {
	MaxResults int          `json:"maxResults"`
	StartAt    int          `json:"startAt"`
	Total      int          `json:"total"`
	IsLast     bool         `json:"isLast"`
	Values     []UserScheme `json:"values"`
}

// Returns a paginated list of the users specified by one or more account IDs.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-users/#api-rest-api-3-user-bulk-get
func (u *UserService) Find(ctx context.Context, accountIDs []string, startAt, maxResults int) (result *UserSearchPageScheme, response *Response, err error) {

	if len(accountIDs) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid accountIDs list")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, accountID := range accountIDs {
		params.Add("accountId", accountID)
	}

	var endpoint = fmt.Sprintf("rest/api/3/user/bulk?%v", params.Encode())

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.Do(request)
	if err != nil {
		return
	}

	result = new(UserSearchPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type UserGroupScheme struct {
	Name string `json:"name"`
	Self string `json:"self"`
}

// Returns the groups to which a user belongs.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-users/#api-rest-api-3-user-groups-get
func (u *UserService) Groups(ctx context.Context, accountID string) (result *[]UserGroupScheme, response *Response, err error) {

	if len(accountID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid accountID value")
	}

	params := url.Values{}
	params.Add("accountId", accountID)

	var endpoint = fmt.Sprintf("rest/api/3/user/groups?%v", accountID)

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.Do(request)
	if err != nil {
		return
	}

	result = new([]UserGroupScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Returns a list of all (active and inactive) users.
// Docs: https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-users/#api-rest-api-3-users-search-get
func (u *UserService) Gets(ctx context.Context, startAt, maxResults int) (result *[]UserScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("rest/api/3/users/search?%v", params.Encode())

	request, err := u.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = u.client.Do(request)
	if err != nil {
		return
	}

	result = new([]UserScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
