package sm

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

type CustomerService struct{ client *Client }

// This method adds a customer to the Jira Service Management
// instance by passing a JSON file including an email address and display name.
// The display name does not need to be unique. The record's identifiers,
// name and key, are automatically generated from the request details.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/customer#create-customer
func (c *CustomerService) Create(ctx context.Context, email, displayName string) (result *CustomerScheme, response *Response, err error) {

	if len(email) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid email value")
	}

	//Check the email
	if !isEmailValid(email) {
		return nil, nil, fmt.Errorf("error, the email (%v) is not valid mail", email)
	}

	if len(displayName) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid displayName value")
	}

	payload := struct {
		DisplayName string `json:"displayName"`
		Email       string `json:"email"`
	}{
		DisplayName: displayName,
		Email:       email,
	}

	var endpoint = "rest/servicedeskapi/customer"

	request, err := c.client.newRequest(ctx, http.MethodPost, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Do(request)
	if err != nil {
		return
	}

	result = new(CustomerScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type CustomerScheme struct {
	AccountID    string `json:"accountId"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	EmailAddress string `json:"emailAddress"`
	DisplayName  string `json:"displayName"`
	Active       bool   `json:"active"`
	TimeZone     string `json:"timeZone"`
	Links        struct {
		JiraRest   string `json:"jiraRest"`
		AvatarUrls struct {
			Four8X48  string `json:"48x48"`
			Two4X24   string `json:"24x24"`
			One6X16   string `json:"16x16"`
			Three2X32 string `json:"32x32"`
		} `json:"avatarUrls"`
		Self string `json:"self"`
	} `json:"_links"`
}

func isEmailValid(email string) bool {
	const emailRegexPattern = "^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

	var regex = regexp.MustCompile(emailRegexPattern)
	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return regex.MatchString(email)
}

// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/customer#get-customers
func (c *CustomerService) Gets(ctx context.Context, serviceDeskID int, query string, start, limit int) (result *CustomerPageScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if query != "" {
		params.Add("query", query)
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer?%v", serviceDeskID, params.Encode())

	request, err := c.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	/*
		This API is experimental.
		Experimental APIs are not guaranteed to be stable within the preview period.
	*/
	request.Header.Set("X-ExperimentalApi", "opt-in")

	response, err = c.client.Do(request)
	if err != nil {
		return
	}

	result = new(CustomerPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/customer#add-customers
func (c *CustomerService) Add(ctx context.Context, serviceDeskID int, accountIDs []string) (response *Response, err error) {

	if len(accountIDs) == 0 {
		return nil, fmt.Errorf("error, please provide a valid accountIDs slice value")
	}

	payload := struct {
		AccountIds []string `json:"accountIds"`
	}{
		AccountIds: accountIDs,
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer", serviceDeskID)

	request, err := c.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/customer#remove-customers
func (c *CustomerService) Remove(ctx context.Context, serviceDeskID int, accountIDs []string) (response *Response, err error) {

	if len(accountIDs) == 0 {
		return nil, fmt.Errorf("error, please provide a valid accountIDs slice value")
	}

	payload := struct {
		AccountIds []string `json:"accountIds"`
	}{
		AccountIds: accountIDs,
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer", serviceDeskID)

	request, err := c.client.newRequest(ctx, http.MethodDelete, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	/*
		This API is experimental.
		Experimental APIs are not guaranteed to be stable within the preview period.
	*/
	request.Header.Set("X-ExperimentalApi", "opt-in")

	response, err = c.client.Do(request)
	if err != nil {
		return
	}

	return
}

type CustomerPageScheme struct {
	Expands    []interface{} `json:"_expands"`
	Size       int           `json:"size"`
	Start      int           `json:"start"`
	Limit      int           `json:"limit"`
	IsLastPage bool          `json:"isLastPage"`
	Links      struct {
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
		Prev    string `json:"prev"`
	} `json:"_links"`
	Values []*CustomerScheme `json:"values"`
}
