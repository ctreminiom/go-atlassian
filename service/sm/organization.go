package sm

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type OrganizationConnector interface {

	// Gets returns a list of organizations in the Jira Service Management instance.
	//
	// Use this method when you want to present a list of organizations or want to locate an organization by name.
	//
	// GET /rest/servicedeskapi/organization
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/organization#get-organizations
	Gets(ctx context.Context, accountID string, start, limit int) (*model.OrganizationPageScheme, *model.ResponseScheme, error)

	// Get returns details of an organization.
	//
	// Use this method to get organization details whenever your application component is passed an organization ID
	//
	// but needs to display other organization details.
	//
	// GET /rest/servicedeskapi/organization/{organizationId}
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/organization#get-organization
	Get(ctx context.Context, organizationID int) (*model.OrganizationScheme, *model.ResponseScheme, error)

	// Delete deletes an organization.
	//
	// Note that the organization is deleted regardless of other associations it may have.
	//
	// For example, associations with service desks.
	//
	// DELETE /rest/servicedeskapi/organization/{organizationId}
	//
	// https://docs.go-atlassian.io/jira-service-management/organization#delete-organization
	Delete(ctx context.Context, organizationID int) (*model.ResponseScheme, error)

	// Create creates an organization by passing the name of the organization.
	//
	// POST /rest/servicedeskapi/organization
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/organization#create-organization
	Create(ctx context.Context, name string) (*model.OrganizationScheme, *model.ResponseScheme, error)

	// Users returns all the users associated with an organization.
	//
	// Use this method where you want to provide a list of users for an
	//
	// organization or determine if a user is associated with an organization.
	//
	// GET /rest/servicedeskapi/organization/{organizationId}/user
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/organization#get-users-in-organization
	Users(ctx context.Context, organizationID, start, limit int) (*model.OrganizationUsersPageScheme, *model.ResponseScheme, error)

	// Add adds users to an organization.
	//
	// POST /rest/servicedeskapi/organization/{organizationId}/user
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/organization#add-users-to-organization
	Add(ctx context.Context, organizationID int, accountIDs []string) (*model.ResponseScheme, error)

	// Remove removes users from an organization.
	//
	// DELETE /rest/servicedeskapi/organization/{organizationId}/user
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/organization#remove-users-from-organization
	Remove(ctx context.Context, organizationID int, accountIDs []string) (*model.ResponseScheme, error)

	// Project returns a list of all organizations associated with a service desk.
	//
	// GET /rest/servicedeskapi/servicedesk/{serviceDeskId}/organization
	//
	// https://docs.go-atlassian.io/jira-service-management/organization#get-project-organizations
	Project(ctx context.Context, accountID string, serviceDeskID, start, limit int) (*model.OrganizationPageScheme, *model.ResponseScheme, error)

	// Associate adds an organization to a service desk.
	//
	// If the organization ID is already associated with the service desk,
	//
	// no change is made and the resource returns a 204 success code.
	//
	// POST /rest/servicedeskapi/servicedesk/{serviceDeskId}/organization
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/organization#associate-organization
	Associate(ctx context.Context, serviceDeskID, organizationID int) (*model.ResponseScheme, error)

	// Detach removes an organization from a service desk.
	//
	// If the organization ID does not match an organization associated with the service desk,
	//
	// no change is made and the resource returns a 204 success code.
	//
	// DELETE /rest/servicedeskapi/servicedesk/{serviceDeskId}/organization
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/organization#detach-organization
	Detach(ctx context.Context, serviceDeskID, organizationID int) (*model.ResponseScheme, error)
}
