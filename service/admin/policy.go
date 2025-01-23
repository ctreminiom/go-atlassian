package admin

import (
	"context"

	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

// OrganizationPolicyConnector represents the cloud admin organization policy endpoints.
// Use it to search, get, create, delete, and change policies.
type OrganizationPolicyConnector interface {

	// Gets returns information about org policies.
	//
	// GET /admin/v1/orgs/{organizationID}/policies
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/policy#get-list-of-policies
	Gets(ctx context.Context, organizationID, policyType, cursor string) (*model.OrganizationPolicyPageScheme, *model.ResponseScheme, error)

	// Get returns information about a single policy by ID.
	//
	// GET /admin/v1/orgs/{organizationID}/policies/{policyID}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/policy#get-a-policy-by-id
	Get(ctx context.Context, organizationID, policyID string) (*model.OrganizationPolicyScheme, *model.ResponseScheme, error)

	// Create creates a policy for an org
	//
	// POST /admin/v1/orgs/{organizationID}/policies
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/policy#create-a-policy
	Create(ctx context.Context, organizationID string, payload *model.OrganizationPolicyData) (*model.OrganizationPolicyScheme, *model.ResponseScheme, error)

	// Update updates a policy for an org
	//
	// PUT /admin/v1/orgs/{organizationID}/policies/{policyID}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/policy#update-a-policy
	Update(ctx context.Context, organizationID, policyID string, payload *model.OrganizationPolicyData) (*model.OrganizationPolicyScheme, *model.ResponseScheme, error)

	// Delete deletes a policy for an org.
	//
	// DELETE /admin/v1/orgs/{organizationID}/policies/{policyID}
	//
	// https://docs.go-atlassian.io/atlassian-admin-cloud/organization/policy#delete-a-policy
	Delete(ctx context.Context, organizationID, policyID string) (*model.ResponseScheme, error)
}
