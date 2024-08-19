package internal

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
	"github.com/ctreminiom/go-atlassian/service"
	"github.com/ctreminiom/go-atlassian/service/sm"
)

// NewCustomerService creates a new instance of CustomerService.
// It takes a service.Connector and a version string as input and returns a pointer to CustomerService.
func NewCustomerService(client service.Connector, version string) *CustomerService {
	return &CustomerService{
		internalClient: &internalCustomerImpl{c: client, version: version},
	}
}

// CustomerService provides methods to interact with customer operations in Jira Service Management.
type CustomerService struct {
	// internalClient is the connector interface for customer operations.
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
func (c *CustomerService) Gets(ctx context.Context, serviceDeskID string, query string, start, limit int) (*model.CustomerPageScheme, *model.ResponseScheme, error) {
	return c.internalClient.Gets(ctx, serviceDeskID, query, start, limit)
}

// Add adds one or more customers to a service desk.
//
// POST /rest/servicedeskapi/servicedesk/{serviceDeskId}/customer
//
// https://docs.go-atlassian.io/jira-service-management-cloud/customer#add-customers
func (c *CustomerService) Add(ctx context.Context, serviceDeskID string, accountIDs []string) (*model.ResponseScheme, error) {
	return c.internalClient.Add(ctx, serviceDeskID, accountIDs)
}

// Remove removes one or more customers from a service desk. The service desk must have closed access
//
// DELETE /rest/servicedeskapi/servicedesk/{serviceDeskId}/customer
//
// https://docs.go-atlassian.io/jira-service-management-cloud/customer#remove-customers
func (c *CustomerService) Remove(ctx context.Context, serviceDeskID string, accountIDs []string) (*model.ResponseScheme, error) {
	return c.internalClient.Remove(ctx, serviceDeskID, accountIDs)
}

type internalCustomerImpl struct {
	c       service.Connector
	version string
}

func (i *internalCustomerImpl) Create(ctx context.Context, email, displayName string) (*model.CustomerScheme, *model.ResponseScheme, error) {

	payload := map[string]interface{}{
		"displayName": displayName,
		"email":       email,
	}

	endpoint := "rest/servicedeskapi/customer"

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, nil, err
	}

	customer := new(model.CustomerScheme)
	res, err := i.c.Call(req, customer)
	if err != nil {
		return nil, res, err
	}

	return customer, res, nil
}

func (i *internalCustomerImpl) Gets(ctx context.Context, serviceDeskID string, query string, start, limit int) (*model.CustomerPageScheme, *model.ResponseScheme, error) {

	params := url.Values{}
	params.Add("start", strconv.Itoa(start))
	params.Add("limit", strconv.Itoa(limit))

	if query != "" {
		params.Add("query", query)
	}

	var endpoint = fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer?%v", serviceDeskID, params.Encode())

	req, err := i.c.NewRequest(ctx, http.MethodGet, endpoint, "", nil)
	if err != nil {
		return nil, nil, err
	}

	page := new(model.CustomerPageScheme)
	res, err := i.c.Call(req, page)
	if err != nil {
		return nil, res, err
	}

	return page, res, nil
}

func (i *internalCustomerImpl) Add(ctx context.Context, serviceDeskID string, accountIDs []string) (*model.ResponseScheme, error) {

	if serviceDeskID == "" {
		return nil, model.ErrNoServiceDeskID
	}

	if len(accountIDs) == 0 {
		return nil, model.ErrNoAccountSlice
	}

	payload := map[string]interface{}{
		"accountIds": accountIDs,
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer", serviceDeskID)

	req, err := i.c.NewRequest(ctx, http.MethodPost, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}

func (i *internalCustomerImpl) Remove(ctx context.Context, serviceDeskID string, accountIDs []string) (*model.ResponseScheme, error) {

	if serviceDeskID == "" {
		return nil, model.ErrNoServiceDeskID
	}

	if len(accountIDs) == 0 {
		return nil, model.ErrNoAccountSlice
	}

	payload := map[string]interface{}{
		"accountIds": accountIDs,
	}

	endpoint := fmt.Sprintf("rest/servicedeskapi/servicedesk/%v/customer", serviceDeskID)

	req, err := i.c.NewRequest(ctx, http.MethodDelete, endpoint, "", payload)
	if err != nil {
		return nil, err
	}

	return i.c.Call(req, nil)
}
