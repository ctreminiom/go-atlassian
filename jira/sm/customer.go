package sm

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"net/http"
	"net/url"
	"strconv"
)

type CustomerService struct{ client *Client }

// Create adds a customer to the Jira Service Management
// instance by passing a JSON file including an email address and display name.
// The display name does not need to be unique. The record's identifiers,
// name and key, are automatically generated from the request details.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/customer#create-customer
func (c *CustomerService) Create(ctx context.Context, email, displayName string) (result *model.CustomerScheme,
	response *ResponseScheme, err error) {

	if len(email) == 0 {
		return nil, nil, model.ErrNoCustomerMailError
	}

	if len(displayName) == 0 {
		return nil, nil, model.ErrNoCustomerDisplayNameError
	}

	payload := struct {
		DisplayName string `json:"displayName"`
		Email       string `json:"email"`
	}{
		DisplayName: displayName,
		Email:       email,
	}

	var (
		payloadAsReader, _ = transformStructToReader(&payload)
		endpoint           = "rest/servicedeskapi/customer"
	)

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

// Gets  returns a list of the customers on a service desk.
// Docs: https://docs.go-atlassian.io/jira-service-management-cloud/customer#get-customers
func (c *CustomerService) Gets(ctx context.Context, serviceDeskID int, query string, start, limit int) (result *model.CustomerPageScheme,
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
		return nil, model.ErrNoAccountSliceError
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
		return nil, model.ErrNoAccountSliceError
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
