package sm

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

type CustomerService struct{ client *Client }

// Create adds a customer to the Jira Service Management
// instance by passing a JSON file including an email address and display name.
// The display name does not need to be unique. The record's identifiers,
// name and key, are automatically generated from the request details.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/customer#create-customer
func (c *CustomerService) Create(ctx context.Context, email, displayName string) (result *CustomerScheme,
	response *ResponseScheme, err error) {

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

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = "rest/servicedeskapi/customer"

	request, err := c.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Call(request, &result)
	if err != nil {
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
	return regex.MatchString(email)
}

// Gets  returns a list of the customers on a service desk.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/customer#get-customers
func (c *CustomerService) Gets(ctx context.Context, serviceDeskID int, query string, start, limit int) (result *CustomerPageScheme,
	response *ResponseScheme, err error) {

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

	response, err = c.client.Call(request, &result)
	if err != nil {
		return
	}

	return
}

// Add adds one or more customers to a service desk.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/customer#add-customers
func (c *CustomerService) Add(ctx context.Context, serviceDeskID int, accountIDs []string) (response *ResponseScheme, err error) {

	if len(accountIDs) == 0 {
		return nil, fmt.Errorf("error, please provide a valid accountIDs slice value")
	}

	payload := struct {
		AccountIds []string `json:"accountIds"`
	}{
		AccountIds: accountIDs,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer", serviceDeskID)

	request, err := c.client.newRequest(ctx, http.MethodPost, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	response, err = c.client.Call(request, nil)
	if err != nil {
		return
	}

	return
}

// Remove removes one or more customers from a service desk. The service desk must have closed access
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/customer#remove-customers
func (c *CustomerService) Remove(ctx context.Context, serviceDeskID int, accountIDs []string) (response *ResponseScheme, err error) {

	if len(accountIDs) == 0 {
		return nil, fmt.Errorf("error, please provide a valid accountIDs slice value")
	}

	payload := struct {
		AccountIds []string `json:"accountIds"`
	}{
		AccountIds: accountIDs,
	}

	payloadAsReader, _ := transformStructToReader(&payload)

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer", serviceDeskID)

	request, err := c.client.newRequest(ctx, http.MethodDelete, endpoint, payloadAsReader)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", "application/json")

	/*
		This API is experimental.
		Experimental APIs are not guaranteed to be stable within the preview period.
	*/
	request.Header.Set("X-ExperimentalApi", "opt-in")

	response, err = c.client.Call(request, nil)
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
