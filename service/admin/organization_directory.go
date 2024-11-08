package admin

import (
	"context"

	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
)

// OrganizationDirectoryConnector represents the cloud admin organization directory endpoints.
// Use it to remove, restore, and suspend users from the organization.
type OrganizationDirectoryConnector interface {

	// Activity returns a user’s last active date for each product listed in Atlassian Administration.
	//
	// Active is defined as viewing a product's page for a minimum of 2 seconds.
	//
	// Last activity data can be delayed by up to 4 hours.
	//
	// If the user has not accessed a product, the product_access response field will be empty.
	//
	// The added_to_org date field is available only to customers using the new user management experience.
	//
	// GET /admin/v1/orgs/{organizationID}/directory/users/{accountID}/last-active-dates
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/directory#users-last-active-dates
	Activity(ctx context.Context, organizationID, accountID string) (*models.UserProductAccessScheme, *models.ResponseScheme, error)

	// Remove removes user access to products listed in Atlassian Administration.
	//
	// -- The API is available for customers using the new user management experience only. --
	//
	// Note: Users with emails whose domain is claimed can still be found in Managed accounts in Directory.
	//
	// DELETE /admin/v1/orgs/{organizationID}/directory/users/{accountID}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/directory#remove-user-access
	Remove(ctx context.Context, organizationID, accountID string) (*models.ResponseScheme, error)

	// Suspend suspends user access to products listed in Atlassian Administration.
	//
	// -- The API is available for customers using the new user management experience only. --
	//
	// Note: Users with emails whose domain is claimed can still be found in Managed accounts in Directory.
	//
	// POST /admin/v1/orgs/{organizationID}/directory/users/{accountID}/suspend-access
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/directory#suspend-user-access
	Suspend(ctx context.Context, organizationID, accountID string) (*models.GenericActionSuccessScheme, *models.ResponseScheme, error)

	// Restore restores user access to products listed in Atlassian Administration.
	//
	// -- The API is available for customers using the new user management experience only. --
	//
	// Note: Users with emails whose domain is claimed can still be found in Managed accounts in Directory.
	//
	// POST /admin/v1/orgs/{organizationID}/directory/users/{accountID}/restore-access
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/directory#restore-user-access
	Restore(ctx context.Context, organizationID, accountID string) (*models.GenericActionSuccessScheme, *models.ResponseScheme, error)
}
