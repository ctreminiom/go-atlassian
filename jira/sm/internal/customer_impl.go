package internal

import (
	"context"
	"fmt"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/sm"
	"net/http"
	"net/url"
	"strconv"
)

func NewCustomerService(client service.Client, version string) (*CustomerService, error) {

	if version == "" {
		return nil, model.ErrNoVersionProvided
	}

	return &CustomerService{
		internalClient: &internalCustomerImpl{c: client, version: version},
	}, nil
}

type CustomerService struct {
	internalClient sm.CustomerConnector
}

// Create adds a customer to the Jira Service Management
//
// instance by passing a JSON file including an email address and display name.
//
// The display name does not need to be unique. The record's identifiers,
//
// name and key, are automatically generated from the request details.
//
// POST /rest/servicedeskapi/customer
//
// https://docs.go-atlassian.io/jira-service-management-cloud/customer#create-customer
func (c *CustomerService) Create(ctx context.Context, email, displayName string) (*model.CustomerScheme, *model.ResponseScheme, error) {
	return c.internalClient.Create(ctx, email, displayName)
}

// Gets  returns a list of the customers on a service desk.
//
// GET /rest/servicedeskapi/servicedesk/{serviceDeskId}/customer
//
// https://docs.go-atlassian.io/jira-service-management-cloud/customer#get-customers
func (c *CustomerService) Gets(ctx context.Context, serviceDeskID int, query string, start, limit int) (*model.CustomerPageScheme, *model.ResponseScheme, error) {
	return c.internalClient.Gets(ctx, serviceDeskID, query, start, limit)
}

// Add adds one or more customers to a service desk.
//
// POST /rest/servicedeskapi/servicedesk/{serviceDeskId}/customer
//
// https://docs.go-atlassian.io/jira-service-management-cloud/customer#add-customers
func (c *CustomerService) Add(ctx context.Context, serviceDeskID int, accountIDs []string) (*model.ResponseScheme, error) {
	return c.internalClient.Add(ctx, serviceDeskID, accountIDs)
}

// Remove removes one or more customers from a service desk. The service desk must have closed access
//
// DELETE /rest/servicedeskapi/servicedesk/{serviceDeskId}/customer
//
// https://docs.go-atlassian.io/jira-service-management-cloud/customer#remove-customers
func (c *CustomerService) Remove(ctx context.Context, serviceDeskID int, accountIDs []string) (*model.ResponseScheme, error) {
	return c.internalClient.Remove(ctx, serviceDeskID, accountIDs)
}

type internalCustomerImpl struct {
	c       service.Client
	version string
}

func (i *internalCustomerImpl) Create(ctx context.Context, email, displayName string) (*model.CustomerScheme, *model.ResponseScheme, error) {

	payload := struct {
		DisplayName string `json:"displayName,omitempty"`
		Email       string `json:"email,omitempty"`
	}{
		DisplayName: displayName,
		Email:       email,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, nil, err
	}

	endpoint := "rest/servicedeskapi/customer"

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, nil, err
	}

	customer := new(model.CustomerScheme)
	response, err := i.c.Call(request, customer)
	if err != nil {
		return nil, response, err
	}

	return customer, response, nil
}

func (i *internalCustomerImpl) Gets(ctx context.Context, serviceDeskID int, query string, start, limit int) (*model.CustomerPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if query != "" {
		params.Add("query", query)
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer?%v", serviceDeskID, params.Encode())

	request, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.CustomerPageScheme)
	response, err := i.c.Call(request, page)
	if err != nil {
		return nil, response, err
	}

	return page, response, nil
}

func (i *internalCustomerImpl) Add(ctx context.Context, serviceDeskID int, accountIDs []string) (*model.ResponseScheme, error) {

	if serviceDeskID == 0 {
		return nil, model.ErrNoServiceDeskIDError
	}

	if len(accountIDs) == 0 {
		return nil, model.ErrNoAccountSliceError
	}

	payload := struct {
		AccountIds []string `json:"accountIds"`
	}{
		AccountIds: accountIDs,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer", serviceDeskID)

	request, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}

func (i *internalCustomerImpl) Remove(ctx context.Context, serviceDeskID int, accountIDs []string) (*model.ResponseScheme, error) {

	if serviceDeskID == 0 {
		return nil, model.ErrNoServiceDeskIDError
	}

	if len(accountIDs) == 0 {
		return nil, model.ErrNoAccountSliceError
	}

	payload := struct {
		AccountIds []string `json:"accountIds"`
	}{
		AccountIds: accountIDs,
	}

	reader, err := i.c.TransformStructToReader(&payload)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer", serviceDeskID)

	request, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, reader)
	if err != nil {
		return nil, err
	}

	return i.c.Call(request, nil)
}
