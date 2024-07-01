package sm

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type CustomerConnector interface {

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
	Create(ctx context.Context, email, displayName string) (*model.CustomerScheme, *model.ResponseScheme, error)

	// Gets  returns a list of the customers on a service desk.
	//
	// GET /rest/servicedeskapi/servicedesk/{serviceDeskId}/customer
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/customer#get-customers
	Gets(ctx context.Context, serviceDeskID string, query string, start, limit int) (*model.CustomerPageScheme, *model.ResponseScheme, error)

	// Add adds one or more customers to a service desk.
	//
	// POST /rest/servicedeskapi/servicedesk/{serviceDeskId}/customer
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/customer#add-customers
	Add(ctx context.Context, serviceDeskID string, accountIDs []string) (*model.ResponseScheme, error)

	// Remove removes one or more customers from a service desk. The service desk must have closed access
	//
	// DELETE /rest/servicedeskapi/servicedesk/{serviceDeskId}/customer
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/customer#remove-customers
	Remove(ctx context.Context, serviceDeskID string, accountIDs []string) (*model.ResponseScheme, error)
}
