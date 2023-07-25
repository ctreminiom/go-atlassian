package sm

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type TypeConnector interface {

	// Search returns all customer request types used in the Jira Service Management instance,
	// optionally filtered by a query string.
	//
	// GET /rest/servicedeskapi/requesttype
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/types#get-all-request-types
	Search(ctx context.Context, query string, start, limit int) (*model.RequestTypePageScheme, *model.ResponseScheme, error)

	// Gets returns all customer request types from a service desk.
	//
	// GET /rest/servicedeskapi/servicedesk/{serviceDeskId}/requesttype
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/types#get-request-types
	Gets(ctx context.Context, serviceDeskID, groupID, start, limit int) (*model.ProjectRequestTypePageScheme, *model.ResponseScheme, error)

	// Create enables a customer request type to be added to a service desk based on an issue type.
	//
	// POST /rest/servicedeskapi/servicedesk/{serviceDeskId}/requesttype
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/types#create-request-type
	Create(ctx context.Context, serviceDeskID int, payload *model.RequestTypePayloadScheme) (*model.RequestTypeScheme, *model.ResponseScheme, error)

	// Get returns a customer request type from a service desk.
	//
	// GET /rest/servicedeskapi/servicedesk/{serviceDeskId}/requesttype/{requestTypeId}
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/types#get-request-type-by-id
	Get(ctx context.Context, serviceDeskID, requestTypeID int) (*model.RequestTypeScheme, *model.ResponseScheme, error)

	// Delete deletes a customer request type from a service desk, and removes it from all customer requests.
	//
	// DELETE /rest/servicedeskapi/servicedesk/{serviceDeskId}/requesttype/{requestTypeId}
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/types#delete-request-type
	Delete(ctx context.Context, serviceDeskID, requestTypeID int) (*model.ResponseScheme, error)

	// Fields returns the fields for a service desk's customer request type.
	//
	// GET /rest/servicedeskapi/servicedesk/{serviceDeskId}/requesttype/{requestTypeId}/field
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/types#get-request-type-fields
	Fields(ctx context.Context, serviceDeskID, requestTypeID int) (*model.RequestTypeFieldsScheme, *model.ResponseScheme, error)
}
